// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: flyteidl/service/auth.proto

package service

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OAuth2MetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OAuth2MetadataRequest) Reset() {
	*x = OAuth2MetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_service_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2MetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2MetadataRequest) ProtoMessage() {}

func (x *OAuth2MetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_service_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2MetadataRequest.ProtoReflect.Descriptor instead.
func (*OAuth2MetadataRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_service_auth_proto_rawDescGZIP(), []int{0}
}

// OAuth2MetadataResponse defines an RFC-Compliant response for /.well-known/oauth-authorization-server metadata
// as defined in https://tools.ietf.org/html/rfc8414
type OAuth2MetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Defines the issuer string in all JWT tokens this server issues. The issuer can be admin itself or an external
	// issuer.
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// URL of the authorization server's authorization endpoint [RFC6749]. This is REQUIRED unless no grant types are
	// supported that use the authorization endpoint.
	AuthorizationEndpoint string `protobuf:"bytes,2,opt,name=authorization_endpoint,json=authorizationEndpoint,proto3" json:"authorization_endpoint,omitempty"`
	// URL of the authorization server's token endpoint [RFC6749].
	TokenEndpoint string `protobuf:"bytes,3,opt,name=token_endpoint,json=tokenEndpoint,proto3" json:"token_endpoint,omitempty"`
	// Array containing a list of the OAuth 2.0 response_type values that this authorization server supports.
	ResponseTypesSupported []string `protobuf:"bytes,4,rep,name=response_types_supported,json=responseTypesSupported,proto3" json:"response_types_supported,omitempty"`
	// JSON array containing a list of the OAuth 2.0 [RFC6749] scope values that this authorization server supports.
	ScopesSupported []string `protobuf:"bytes,5,rep,name=scopes_supported,json=scopesSupported,proto3" json:"scopes_supported,omitempty"`
	// JSON array containing a list of client authentication methods supported by this token endpoint.
	TokenEndpointAuthMethodsSupported []string `protobuf:"bytes,6,rep,name=token_endpoint_auth_methods_supported,json=tokenEndpointAuthMethodsSupported,proto3" json:"token_endpoint_auth_methods_supported,omitempty"`
	// URL of the authorization server's JWK Set [JWK] document. The referenced document contains the signing key(s) the
	// client uses to validate signatures from the authorization server.
	JwksUri string `protobuf:"bytes,7,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
	// JSON array containing a list of Proof Key for Code Exchange (PKCE) [RFC7636] code challenge methods supported by
	// this authorization server.
	CodeChallengeMethodsSupported []string `protobuf:"bytes,8,rep,name=code_challenge_methods_supported,json=codeChallengeMethodsSupported,proto3" json:"code_challenge_methods_supported,omitempty"`
	// JSON array containing a list of the OAuth 2.0 grant type values that this authorization server supports.
	GrantTypesSupported []string `protobuf:"bytes,9,rep,name=grant_types_supported,json=grantTypesSupported,proto3" json:"grant_types_supported,omitempty"`
}

func (x *OAuth2MetadataResponse) Reset() {
	*x = OAuth2MetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_service_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuth2MetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuth2MetadataResponse) ProtoMessage() {}

func (x *OAuth2MetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_service_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuth2MetadataResponse.ProtoReflect.Descriptor instead.
func (*OAuth2MetadataResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_service_auth_proto_rawDescGZIP(), []int{1}
}

func (x *OAuth2MetadataResponse) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *OAuth2MetadataResponse) GetAuthorizationEndpoint() string {
	if x != nil {
		return x.AuthorizationEndpoint
	}
	return ""
}

func (x *OAuth2MetadataResponse) GetTokenEndpoint() string {
	if x != nil {
		return x.TokenEndpoint
	}
	return ""
}

func (x *OAuth2MetadataResponse) GetResponseTypesSupported() []string {
	if x != nil {
		return x.ResponseTypesSupported
	}
	return nil
}

func (x *OAuth2MetadataResponse) GetScopesSupported() []string {
	if x != nil {
		return x.ScopesSupported
	}
	return nil
}

func (x *OAuth2MetadataResponse) GetTokenEndpointAuthMethodsSupported() []string {
	if x != nil {
		return x.TokenEndpointAuthMethodsSupported
	}
	return nil
}

func (x *OAuth2MetadataResponse) GetJwksUri() string {
	if x != nil {
		return x.JwksUri
	}
	return ""
}

func (x *OAuth2MetadataResponse) GetCodeChallengeMethodsSupported() []string {
	if x != nil {
		return x.CodeChallengeMethodsSupported
	}
	return nil
}

func (x *OAuth2MetadataResponse) GetGrantTypesSupported() []string {
	if x != nil {
		return x.GrantTypesSupported
	}
	return nil
}

type PublicClientAuthConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublicClientAuthConfigRequest) Reset() {
	*x = PublicClientAuthConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_service_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicClientAuthConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicClientAuthConfigRequest) ProtoMessage() {}

func (x *PublicClientAuthConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_service_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicClientAuthConfigRequest.ProtoReflect.Descriptor instead.
func (*PublicClientAuthConfigRequest) Descriptor() ([]byte, []int) {
	return file_flyteidl_service_auth_proto_rawDescGZIP(), []int{2}
}

// FlyteClientResponse encapsulates public information that flyte clients (CLIs... etc.) can use to authenticate users.
type PublicClientAuthConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// client_id to use when initiating OAuth2 authorization requests.
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// redirect uri to use when initiating OAuth2 authorization requests.
	RedirectUri string `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	// scopes to request when initiating OAuth2 authorization requests.
	Scopes []string `protobuf:"bytes,3,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// Authorization Header to use when passing Access Tokens to the server. If not provided, the client should use the
	// default http `Authorization` header.
	AuthorizationMetadataKey string `protobuf:"bytes,4,opt,name=authorization_metadata_key,json=authorizationMetadataKey,proto3" json:"authorization_metadata_key,omitempty"`
	// ServiceHttpEndpoint points to the http endpoint for the backend. If empty, clients can assume the endpoint used
	// to configure the gRPC connection can be used for the http one respecting the insecure flag to choose between
	// SSL or no SSL connections.
	ServiceHttpEndpoint string `protobuf:"bytes,5,opt,name=service_http_endpoint,json=serviceHttpEndpoint,proto3" json:"service_http_endpoint,omitempty"`
}

func (x *PublicClientAuthConfigResponse) Reset() {
	*x = PublicClientAuthConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyteidl_service_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicClientAuthConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicClientAuthConfigResponse) ProtoMessage() {}

func (x *PublicClientAuthConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flyteidl_service_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicClientAuthConfigResponse.ProtoReflect.Descriptor instead.
func (*PublicClientAuthConfigResponse) Descriptor() ([]byte, []int) {
	return file_flyteidl_service_auth_proto_rawDescGZIP(), []int{3}
}

func (x *PublicClientAuthConfigResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *PublicClientAuthConfigResponse) GetRedirectUri() string {
	if x != nil {
		return x.RedirectUri
	}
	return ""
}

func (x *PublicClientAuthConfigResponse) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *PublicClientAuthConfigResponse) GetAuthorizationMetadataKey() string {
	if x != nil {
		return x.AuthorizationMetadataKey
	}
	return ""
}

func (x *PublicClientAuthConfigResponse) GetServiceHttpEndpoint() string {
	if x != nil {
		return x.ServiceHttpEndpoint
	}
	return ""
}

var File_flyteidl_service_auth_proto protoreflect.FileDescriptor

var file_flyteidl_service_auth_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x32, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xdd, 0x03, 0x0a, 0x16, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x16, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x18, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12,
	0x29, 0x0a, 0x10, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x50, 0x0a, 0x25, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x21, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x6a, 0x77, 0x6b, 0x73, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6a, 0x77, 0x6b, 0x73, 0x55, 0x72, 0x69, 0x12, 0x47, 0x0a, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x1d, 0x63, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x12, 0x32, 0x0a, 0x15, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x13, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x22, 0x1f, 0x0a, 0x1d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xea, 0x01, 0x0a, 0x1e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73,
	0x12, 0x3c, 0x0a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x12, 0x32,
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x74, 0x74, 0x70, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x32, 0xfc, 0x03, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xf5, 0x01, 0x0a, 0x11, 0x47,
	0x65, 0x74, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x27, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x66, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x41, 0x75,
	0x74, 0x68, 0x32, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x8c, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x2e,
	0x77, 0x65, 0x6c, 0x6c, 0x2d, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x92, 0x41, 0x5a, 0x1a, 0x58, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x73, 0x20, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x20, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x69, 0x73, 0x20, 0x61, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x6c, 0x79, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x2e, 0x12, 0xec, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f, 0x2e, 0x66,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x70, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x92, 0x41, 0x4e, 0x1a, 0x4c, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x73, 0x20, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x20, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x20, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x20, 0x54, 0x68, 0x69, 0x73, 0x20, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x69, 0x73, 0x20, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x6c, 0x79, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x2e, 0x42, 0xbb, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x69,
	0x64, 0x6c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x74, 0x65, 0x6f, 0x72, 0x67, 0x2f, 0x66, 0x6c, 0x79,
	0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2d, 0x67, 0x6f, 0x2f,
	0x66, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0xa2, 0x02, 0x03, 0x46, 0x53, 0x58, 0xaa, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x65, 0x69, 0x64,
	0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xca, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74,
	0x65, 0x69, 0x64, 0x6c, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xe2, 0x02, 0x1c, 0x46,
	0x6c, 0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x46, 0x6c,
	0x79, 0x74, 0x65, 0x69, 0x64, 0x6c, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyteidl_service_auth_proto_rawDescOnce sync.Once
	file_flyteidl_service_auth_proto_rawDescData = file_flyteidl_service_auth_proto_rawDesc
)

func file_flyteidl_service_auth_proto_rawDescGZIP() []byte {
	file_flyteidl_service_auth_proto_rawDescOnce.Do(func() {
		file_flyteidl_service_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyteidl_service_auth_proto_rawDescData)
	})
	return file_flyteidl_service_auth_proto_rawDescData
}

var file_flyteidl_service_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_flyteidl_service_auth_proto_goTypes = []interface{}{
	(*OAuth2MetadataRequest)(nil),          // 0: flyteidl.service.OAuth2MetadataRequest
	(*OAuth2MetadataResponse)(nil),         // 1: flyteidl.service.OAuth2MetadataResponse
	(*PublicClientAuthConfigRequest)(nil),  // 2: flyteidl.service.PublicClientAuthConfigRequest
	(*PublicClientAuthConfigResponse)(nil), // 3: flyteidl.service.PublicClientAuthConfigResponse
}
var file_flyteidl_service_auth_proto_depIdxs = []int32{
	0, // 0: flyteidl.service.AuthMetadataService.GetOAuth2Metadata:input_type -> flyteidl.service.OAuth2MetadataRequest
	2, // 1: flyteidl.service.AuthMetadataService.GetPublicClientConfig:input_type -> flyteidl.service.PublicClientAuthConfigRequest
	1, // 2: flyteidl.service.AuthMetadataService.GetOAuth2Metadata:output_type -> flyteidl.service.OAuth2MetadataResponse
	3, // 3: flyteidl.service.AuthMetadataService.GetPublicClientConfig:output_type -> flyteidl.service.PublicClientAuthConfigResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flyteidl_service_auth_proto_init() }
func file_flyteidl_service_auth_proto_init() {
	if File_flyteidl_service_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flyteidl_service_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2MetadataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_service_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuth2MetadataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_service_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicClientAuthConfigRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_flyteidl_service_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicClientAuthConfigResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_flyteidl_service_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flyteidl_service_auth_proto_goTypes,
		DependencyIndexes: file_flyteidl_service_auth_proto_depIdxs,
		MessageInfos:      file_flyteidl_service_auth_proto_msgTypes,
	}.Build()
	File_flyteidl_service_auth_proto = out.File
	file_flyteidl_service_auth_proto_rawDesc = nil
	file_flyteidl_service_auth_proto_goTypes = nil
	file_flyteidl_service_auth_proto_depIdxs = nil
}
