package impl

import (
	"github.com/flyteorg/flyteadmin/pkg/errors"
	"github.com/flyteorg/flyteadmin/pkg/workflowengine/interfaces"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	"google.golang.org/grpc/codes"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func addMapValues(overrides map[string]string, defaultValues map[string]string) map[string]string {
	if defaultValues == nil {
		defaultValues = map[string]string{}
	}
	if overrides == nil {
		return defaultValues
	}
	for label, value := range overrides {
		defaultValues[label] = value
	}
	return defaultValues
}

func addPermissions(auth *admin.AuthRole, roleNameKey string, flyteWf *v1alpha1.FlyteWorkflow) {
	// Set role permissions based on launch plan Auth values.
	// The branched-ness of this check is due to the presence numerous deprecated fields
	if auth == nil {
		return
	}
	if len(auth.AssumableIamRole) > 0 {
		if flyteWf.Annotations == nil {
			flyteWf.Annotations = map[string]string{}
		}
		flyteWf.Annotations[roleNameKey] = auth.AssumableIamRole
	}
	if len(auth.KubernetesServiceAccount) > 0 {
		flyteWf.ServiceAccountName = auth.KubernetesServiceAccount
	}
}

func addExecutionOverrides(taskPluginOverrides []*admin.PluginOverride,
	workflowExecutionConfig *admin.WorkflowExecutionConfig, recoveryExecution *core.WorkflowExecutionIdentifier,
	taskResources *interfaces.TaskResources, flyteWf *v1alpha1.FlyteWorkflow) {
	executionConfig := v1alpha1.ExecutionConfig{
		TaskPluginImpls: make(map[string]v1alpha1.TaskPluginOverride),
		RecoveryExecution: v1alpha1.WorkflowExecutionIdentifier{
			WorkflowExecutionIdentifier: recoveryExecution,
		},
	}
	for _, override := range taskPluginOverrides {
		executionConfig.TaskPluginImpls[override.TaskType] = v1alpha1.TaskPluginOverride{
			PluginIDs:             override.PluginId,
			MissingPluginBehavior: override.MissingPluginBehavior,
		}

	}
	if workflowExecutionConfig != nil {
		executionConfig.MaxParallelism = uint32(workflowExecutionConfig.MaxParallelism)
	}
	if taskResources != nil {
		var requests = v1alpha1.TaskResourceSpec{}
		if !taskResources.Defaults.CPU.IsZero() {
			requests.CPU = taskResources.Defaults.CPU
		}
		if !taskResources.Defaults.Memory.IsZero() {
			requests.Memory = taskResources.Defaults.Memory
		}
		if !taskResources.Defaults.EphemeralStorage.IsZero() {
			requests.EphemeralStorage = taskResources.Defaults.EphemeralStorage
		}
		if !taskResources.Defaults.Storage.IsZero() {
			requests.Storage = taskResources.Defaults.Storage
		}
		if !taskResources.Defaults.GPU.IsZero() {
			requests.GPU = taskResources.Defaults.GPU
		}

		var limits = v1alpha1.TaskResourceSpec{}
		if !taskResources.Limits.CPU.IsZero() {
			limits.CPU = taskResources.Limits.CPU
		}
		if !taskResources.Limits.Memory.IsZero() {
			limits.Memory = taskResources.Limits.Memory
		}
		if !taskResources.Limits.EphemeralStorage.IsZero() {
			limits.EphemeralStorage = taskResources.Limits.EphemeralStorage
		}
		if !taskResources.Limits.Storage.IsZero() {
			limits.Storage = taskResources.Limits.Storage
		}
		if !taskResources.Limits.GPU.IsZero() {
			limits.GPU = taskResources.Limits.GPU
		}
		executionConfig.TaskResources = v1alpha1.TaskResources{
			Requests: requests,
			Limits:   limits,
		}
	}
	flyteWf.ExecutionConfig = executionConfig
}

func PrepareFlyteWorkflow(data interfaces.ExecutionData, flyteWorkflow *v1alpha1.FlyteWorkflow) error {
	if data.ExecutionID == nil {
		return errors.NewFlyteAdminErrorf(codes.Internal, "invalid execution id")
	}
	if flyteWorkflow == nil {
		return errors.NewFlyteAdminErrorf(codes.Internal, "missing Flyte Workflow")
	}

	// add the executionId so Propeller can send events back that are associated with the ID
	flyteWorkflow.ExecutionID = v1alpha1.WorkflowExecutionIdentifier{
		WorkflowExecutionIdentifier: data.ExecutionID,
	}
	// add the acceptedAt timestamp so propeller can emit latency metrics.
	acceptAtWrapper := v1.NewTime(data.ExecutionParameters.AcceptedAt)
	flyteWorkflow.AcceptedAt = &acceptAtWrapper

	addPermissions(data.ExecutionParameters.Auth, data.ExecutionParameters.RoleNameKey, flyteWorkflow)

	labels := addMapValues(data.ExecutionParameters.Labels, flyteWorkflow.Labels)
	flyteWorkflow.Labels = labels
	annotations := addMapValues(data.ExecutionParameters.Annotations, flyteWorkflow.Annotations)
	flyteWorkflow.Annotations = annotations
	if flyteWorkflow.WorkflowMeta == nil {
		flyteWorkflow.WorkflowMeta = &v1alpha1.WorkflowMeta{}
	}
	flyteWorkflow.WorkflowMeta.EventVersion = v1alpha1.EventVersion(data.ExecutionParameters.EventVersion)
	addExecutionOverrides(data.ExecutionParameters.TaskPluginOverrides, data.ExecutionParameters.ExecutionConfig,
		data.ExecutionParameters.RecoveryExecution, data.ExecutionParameters.TaskResources, flyteWorkflow)

	if data.ExecutionParameters.RawOutputDataConfig != nil {
		flyteWorkflow.RawOutputDataConfig = v1alpha1.RawOutputDataConfig{
			RawOutputDataConfig: data.ExecutionParameters.RawOutputDataConfig,
		}
	}

	return nil
}
