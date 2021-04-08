package runtime

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/flyteorg/flyteadmin/pkg/runtime/interfaces"
	"github.com/flyteorg/flytestdlib/config"
	"github.com/flyteorg/flytestdlib/logger"
)

const database = "database"
const flyteAdmin = "flyteadmin"
const scheduler = "scheduler"
const remoteData = "remoteData"
const notifications = "notifications"
const domains = "domains"
const externalEvents = "externalEvents"

const defaultAsyncEventsBufferSize = 100

var databaseConfig = config.MustRegisterSection(database, &interfaces.DbConfigSection{})
var flyteAdminConfig = config.MustRegisterSection(flyteAdmin, &interfaces.ApplicationConfig{})
var schedulerConfig = config.MustRegisterSection(scheduler, &interfaces.SchedulerConfig{})
var remoteDataConfig = config.MustRegisterSection(remoteData, &interfaces.RemoteDataConfig{})
var notificationsConfig = config.MustRegisterSection(notifications, &interfaces.NotificationsConfig{})
var domainsConfig = config.MustRegisterSection(domains, &interfaces.DomainsConfig{})
var externalEventsConfig = config.MustRegisterSection(externalEvents, &interfaces.ExternalEventsConfig{})

// Implementation of an interfaces.ApplicationConfiguration
type ApplicationConfigurationProvider struct{}

func (p *ApplicationConfigurationProvider) GetDbConfig() interfaces.DbConfig {
	dbConfigSection := databaseConfig.GetConfig().(*interfaces.DbConfigSection)
	password := dbConfigSection.Password
	if len(dbConfigSection.PasswordPath) > 0 {
		if _, err := os.Stat(dbConfigSection.PasswordPath); os.IsNotExist(err) {
			logger.Fatalf(context.Background(),
				"missing database password at specified path [%s]", dbConfigSection.PasswordPath)
		}
		passwordVal, err := ioutil.ReadFile(dbConfigSection.PasswordPath)
		if err != nil {
			logger.Fatalf(context.Background(), "failed to read database password from path [%s] with err: %v",
				dbConfigSection.PasswordPath, err)
		}
		password = string(passwordVal)
	}
	return interfaces.DbConfig{
		Host:         dbConfigSection.Host,
		Port:         dbConfigSection.Port,
		DbName:       dbConfigSection.DbName,
		User:         dbConfigSection.User,
		Password:     password,
		ExtraOptions: dbConfigSection.ExtraOptions,
		Debug:        dbConfigSection.Debug,
	}
}

func (p *ApplicationConfigurationProvider) GetTopLevelConfig() *interfaces.ApplicationConfig {
	applicationConfig := flyteAdminConfig.GetConfig().(*interfaces.ApplicationConfig)
	if applicationConfig.AsyncEventsBufferSize == 0 {
		applicationConfig.AsyncEventsBufferSize = defaultAsyncEventsBufferSize
	}
	return applicationConfig
}

func (p *ApplicationConfigurationProvider) GetSchedulerConfig() *interfaces.SchedulerConfig {
	return schedulerConfig.GetConfig().(*interfaces.SchedulerConfig)
}

func (p *ApplicationConfigurationProvider) GetRemoteDataConfig() *interfaces.RemoteDataConfig {
	return remoteDataConfig.GetConfig().(*interfaces.RemoteDataConfig)
}

func (p *ApplicationConfigurationProvider) GetNotificationsConfig() *interfaces.NotificationsConfig {
	return notificationsConfig.GetConfig().(*interfaces.NotificationsConfig)
}

func (p *ApplicationConfigurationProvider) GetDomainsConfig() *interfaces.DomainsConfig {
	return domainsConfig.GetConfig().(*interfaces.DomainsConfig)
}

func (p *ApplicationConfigurationProvider) GetExternalEventsConfig() *interfaces.ExternalEventsConfig {
	return externalEventsConfig.GetConfig().(*interfaces.ExternalEventsConfig)
}

func NewApplicationConfigurationProvider() interfaces.ApplicationConfiguration {
	return &ApplicationConfigurationProvider{}
}
