// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: authentication/service/v1/authentication.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-admin/api/gen/go/user/service/v1"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 授权类型
type GrantType int32

const (
	GrantType_password           GrantType = 0 // 密码模式（Resource Owner Password Credentials Grant）
	GrantType_client_credentials GrantType = 1 // 客户端模式（Client Credentials Grant）
	GrantType_authorization_code GrantType = 2 // 授权码模式（Authorization Code Grant）
	GrantType_refresh_token      GrantType = 3 // 刷新令牌 (Refresh Token)
	GrantType_implicit           GrantType = 4 // 简化模式 (Implicit Grant)
)

// Enum value maps for GrantType.
var (
	GrantType_name = map[int32]string{
		0: "password",
		1: "client_credentials",
		2: "authorization_code",
		3: "refresh_token",
		4: "implicit",
	}
	GrantType_value = map[string]int32{
		"password":           0,
		"client_credentials": 1,
		"authorization_code": 2,
		"refresh_token":      3,
		"implicit":           4,
	}
)

func (x GrantType) Enum() *GrantType {
	p := new(GrantType)
	*p = x
	return p
}

func (x GrantType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrantType) Descriptor() protoreflect.EnumDescriptor {
	return file_authentication_service_v1_authentication_proto_enumTypes[0].Descriptor()
}

func (GrantType) Type() protoreflect.EnumType {
	return &file_authentication_service_v1_authentication_proto_enumTypes[0]
}

func (x GrantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrantType.Descriptor instead.
func (GrantType) EnumDescriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{0}
}

// 令牌类型
type TokenType int32

const (
	TokenType_bearer TokenType = 0 //
	TokenType_mac    TokenType = 1 //
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "bearer",
		1: "mac",
	}
	TokenType_value = map[string]int32{
		"bearer": 0,
		"mac":    1,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_authentication_service_v1_authentication_proto_enumTypes[1].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_authentication_service_v1_authentication_proto_enumTypes[1]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{1}
}

// 客户端类型
type ClientType int32

const (
	ClientType_admin ClientType = 0 // 管理端
	ClientType_app   ClientType = 1 // APP
)

// Enum value maps for ClientType.
var (
	ClientType_name = map[int32]string{
		0: "admin",
		1: "app",
	}
	ClientType_value = map[string]int32{
		"admin": 0,
		"app":   1,
	}
)

func (x ClientType) Enum() *ClientType {
	p := new(ClientType)
	*p = x
	return p
}

func (x ClientType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientType) Descriptor() protoreflect.EnumDescriptor {
	return file_authentication_service_v1_authentication_proto_enumTypes[2].Descriptor()
}

func (ClientType) Type() protoreflect.EnumType {
	return &file_authentication_service_v1_authentication_proto_enumTypes[2]
}

func (x ClientType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientType.Descriptor instead.
func (ClientType) EnumDescriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{2}
}

// 用户后台登录 - 请求
type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GrantType     string                 `protobuf:"bytes,1,opt,name=grant_type,proto3" json:"grant_type,omitempty"` // 授权类型，此处的值固定为"password"，必选项。
	ClientId      *string                `protobuf:"bytes,2,opt,name=client_id,proto3,oneof" json:"client_id,omitempty"`
	ClientSecret  *string                `protobuf:"bytes,3,opt,name=client_secret,proto3,oneof" json:"client_secret,omitempty"`
	Scope         *string                `protobuf:"bytes,4,opt,name=scope,proto3,oneof" json:"scope,omitempty"`                  // 以空格分隔的范围列表。如果未提供，scope则授权任何范围，默认为空列表。
	RedirectUri   *string                `protobuf:"bytes,5,opt,name=redirect_uri,proto3,oneof" json:"redirect_uri,omitempty"`    // 跳转链接
	Username      *string                `protobuf:"bytes,10,opt,name=username,proto3,oneof" json:"username,omitempty"`           // 用户名，必选项。
	Password      *string                `protobuf:"bytes,11,opt,name=password,proto3,oneof" json:"password,omitempty"`           // 用户的密码，必选项。
	RefreshToken  *string                `protobuf:"bytes,20,opt,name=refresh_token,proto3,oneof" json:"refresh_token,omitempty"` // 更新令牌，用来获取下一次的访问令牌，必选项。
	Code          *string                `protobuf:"bytes,30,opt,name=code,proto3,oneof" json:"code,omitempty"`                   // 授权请求中收到的一次性验证/认证码。(当使用授权码模式时)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

func (x *LoginRequest) GetClientId() string {
	if x != nil && x.ClientId != nil {
		return *x.ClientId
	}
	return ""
}

func (x *LoginRequest) GetClientSecret() string {
	if x != nil && x.ClientSecret != nil {
		return *x.ClientSecret
	}
	return ""
}

func (x *LoginRequest) GetScope() string {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return ""
}

func (x *LoginRequest) GetRedirectUri() string {
	if x != nil && x.RedirectUri != nil {
		return *x.RedirectUri
	}
	return ""
}

func (x *LoginRequest) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *LoginRequest) GetRefreshToken() string {
	if x != nil && x.RefreshToken != nil {
		return *x.RefreshToken
	}
	return ""
}

func (x *LoginRequest) GetCode() string {
	if x != nil && x.Code != nil {
		return *x.Code
	}
	return ""
}

// 用户后台登录 - 回应
type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,proto3" json:"access_token,omitempty"`    // 访问令牌，必选项。
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,proto3" json:"refresh_token,omitempty"`  // 更新令牌，用来获取下一次的访问令牌，可选项。
	TokenType     string                 `protobuf:"bytes,3,opt,name=token_type,proto3" json:"token_type,omitempty"`        // 令牌类型，该值大小写不敏感，必选项，可以是bearer类型或mac类型。
	ExpiresIn     *int64                 `protobuf:"varint,4,opt,name=expires_in,proto3,oneof" json:"expires_in,omitempty"` // 令牌有效时间，单位为秒。如果访问令牌过期，服务器应回复授予访问令牌的持续时间。如果省略该参数，必须其他方式设置过期时间。
	Scope         *string                `protobuf:"bytes,5,opt,name=scope,proto3,oneof" json:"scope,omitempty"`            // 以空格分隔的用户授予范围列表。如果未提供，scope则授权任何范围，默认为空列表。
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *LoginResponse) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *LoginResponse) GetExpiresIn() int64 {
	if x != nil && x.ExpiresIn != nil {
		return *x.ExpiresIn
	}
	return 0
}

func (x *LoginResponse) GetScope() string {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return ""
}

// 验证令牌 - 请求
type ValidateTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,json=isValid,proto3" json:"token,omitempty"`                                                           // 令牌
	ClientType    ClientType             `protobuf:"varint,2,opt,name=client_type,json=clientType,proto3,enum=authentication.service.v1.ClientType" json:"client_type,omitempty"` // 客戶端類型
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ValidateTokenRequest) GetClientType() ClientType {
	if x != nil {
		return x.ClientType
	}
	return ClientType_admin
}

// 验证令牌 - 回应
type ValidateTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsValid       bool                   `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"` // 令牌是否有效
	Claim         *UserTokenPayload      `protobuf:"bytes,2,opt,name=claim,proto3,oneof" json:"claim,omitempty"`               // 用戶令牌载体
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateTokenResponse) Reset() {
	*x = ValidateTokenResponse{}
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenResponse) ProtoMessage() {}

func (x *ValidateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateTokenResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *ValidateTokenResponse) GetClaim() *UserTokenPayload {
	if x != nil {
		return x.Claim
	}
	return nil
}

type RegisterUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`                       // 用户名
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`                       // 登入密码
	TenantCode    string                 `protobuf:"bytes,3,opt,name=tenant_code,json=tenantCode,proto3" json:"tenant_code,omitempty"` // 租户代码
	Email         *string                `protobuf:"bytes,4,opt,name=email,proto3,oneof" json:"email,omitempty"`                       // 电子邮件地址
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterUserRequest) Reset() {
	*x = RegisterUserRequest{}
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserRequest) ProtoMessage() {}

func (x *RegisterUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserRequest.ProtoReflect.Descriptor instead.
func (*RegisterUserRequest) Descriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterUserRequest) GetTenantCode() string {
	if x != nil {
		return x.TenantCode
	}
	return ""
}

func (x *RegisterUserRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

type RegisterUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterUserResponse) Reset() {
	*x = RegisterUserResponse{}
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterUserResponse) ProtoMessage() {}

func (x *RegisterUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterUserResponse.ProtoReflect.Descriptor instead.
func (*RegisterUserResponse) Descriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterUserResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// 用户令牌载体
type UserTokenPayload struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=uid,proto3" json:"user_id,omitempty"`                                        // 用户ID
	TenantId      *uint32                `protobuf:"varint,2,opt,name=tenant_id,json=tid,proto3,oneof" json:"tenant_id,omitempty"`                              // 租户ID
	Username      *string                `protobuf:"bytes,3,opt,name=username,json=sub,proto3,oneof" json:"username,omitempty"`                                 // 用户名
	ClientId      *string                `protobuf:"bytes,4,opt,name=client_id,json=cid,proto3,oneof" json:"client_id,omitempty"`                               // 客户端ID
	Authority     v1.UserAuthority       `protobuf:"varint,5,opt,name=authority,json=aut,proto3,enum=user.service.v1.UserAuthority" json:"authority,omitempty"` // 用户权限
	Roles         []string               `protobuf:"bytes,6,rep,name=roles,json=roc,proto3" json:"roles,omitempty"`                                             // 用户角色列表，以逗号分隔
	RoleId        *uint32                `protobuf:"varint,7,opt,name=role_id,json=rid,proto3,oneof" json:"role_id,omitempty"`                                  // 角色ID
	DeviceId      *string                `protobuf:"bytes,8,opt,name=device_id,json=did,proto3,oneof" json:"device_id,omitempty"`                               // 设备ID
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserTokenPayload) Reset() {
	*x = UserTokenPayload{}
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserTokenPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserTokenPayload) ProtoMessage() {}

func (x *UserTokenPayload) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserTokenPayload.ProtoReflect.Descriptor instead.
func (*UserTokenPayload) Descriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{6}
}

func (x *UserTokenPayload) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserTokenPayload) GetTenantId() uint32 {
	if x != nil && x.TenantId != nil {
		return *x.TenantId
	}
	return 0
}

func (x *UserTokenPayload) GetUsername() string {
	if x != nil && x.Username != nil {
		return *x.Username
	}
	return ""
}

func (x *UserTokenPayload) GetClientId() string {
	if x != nil && x.ClientId != nil {
		return *x.ClientId
	}
	return ""
}

func (x *UserTokenPayload) GetAuthority() v1.UserAuthority {
	if x != nil {
		return x.Authority
	}
	return v1.UserAuthority(0)
}

func (x *UserTokenPayload) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UserTokenPayload) GetRoleId() uint32 {
	if x != nil && x.RoleId != nil {
		return *x.RoleId
	}
	return 0
}

func (x *UserTokenPayload) GetDeviceId() string {
	if x != nil && x.DeviceId != nil {
		return *x.DeviceId
	}
	return ""
}

// 获取当前用户身份信息 - 响应
type WhoAmIResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint32                 `protobuf:"varint,1,opt,name=user_id,json=uid,proto3" json:"user_id,omitempty"`                               // 用户ID
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`                                       // 当前用户的用户名
	Authority     v1.UserAuthority       `protobuf:"varint,3,opt,name=authority,proto3,enum=user.service.v1.UserAuthority" json:"authority,omitempty"` // 用户权限
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WhoAmIResponse) Reset() {
	*x = WhoAmIResponse{}
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WhoAmIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIResponse) ProtoMessage() {}

func (x *WhoAmIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIResponse.ProtoReflect.Descriptor instead.
func (*WhoAmIResponse) Descriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{7}
}

func (x *WhoAmIResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *WhoAmIResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *WhoAmIResponse) GetAuthority() v1.UserAuthority {
	if x != nil {
		return x.Authority
	}
	return v1.UserAuthority(0)
}

// 修改用户密码 - 请求
type ChangePasswordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`                          // 用户名
	OldPassword   string                 `protobuf:"bytes,2,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"` // 旧密码
	NewPassword   string                 `protobuf:"bytes,3,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"` // 新密码
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChangePasswordRequest) Reset() {
	*x = ChangePasswordRequest{}
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordRequest) ProtoMessage() {}

func (x *ChangePasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authentication_service_v1_authentication_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordRequest.ProtoReflect.Descriptor instead.
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) {
	return file_authentication_service_v1_authentication_proto_rawDescGZIP(), []int{8}
}

func (x *ChangePasswordRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChangePasswordRequest) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangePasswordRequest) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

var File_authentication_service_v1_authentication_proto protoreflect.FileDescriptor

const file_authentication_service_v1_authentication_proto_rawDesc = "" +
	"\n" +
	".authentication/service/v1/authentication.proto\x12\x19authentication.service.v1\x1a$gnostic/openapi/v3/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x1auser/service/v1/user.proto\"\xb6\b\n" +
	"\fLoginRequest\x12s\n" +
	"\n" +
	"grant_type\x18\x01 \x01(\tBS\xe0A\x02\xbaGM\x8a\x02\n" +
	"\x1a\bpassword\x92\x02=授权类型，此处的值固定为\"password\"，必选项。R\n" +
	"grant_type\x124\n" +
	"\tclient_id\x18\x02 \x01(\tB\x11\xbaG\x0e\x92\x02\v客户端IDH\x00R\tclient_id\x88\x01\x01\x12@\n" +
	"\rclient_secret\x18\x03 \x01(\tB\x15\xbaG\x12\x92\x02\x0f客户端密钥H\x01R\rclient_secret\x88\x01\x01\x12\x92\x01\n" +
	"\x05scope\x18\x04 \x01(\tBw\xbaGt\x92\x02q以空格分隔的用户授予范围列表。如果未提供，scope则授权任何范围，默认为空列表。H\x02R\x05scope\x88\x01\x01\x12;\n" +
	"\fredirect_uri\x18\x05 \x01(\tB\x12\xbaG\x0f\x92\x02\f跳转链接H\x03R\fredirect_uri\x88\x01\x01\x12\x1f\n" +
	"\busername\x18\n" +
	" \x01(\tH\x04R\busername\x88\x01\x01\x12\x1f\n" +
	"\bpassword\x18\v \x01(\tH\x05R\bpassword\x88\x01\x01\x12\xc2\x02\n" +
	"\rrefresh_token\x18\x14 \x01(\tB\x96\x02\xbaG\x92\x02\x92\x02\x8e\x02更新令牌，用来获取下一次的访问令牌，可选项。如果访问令牌将过期，则返回刷新令牌很有用，应用程序可以使用该刷新令牌来获取另一个访问令牌。但是，通过隐式授予颁发的令牌不能颁发刷新令牌。H\x06R\rrefresh_token\x88\x01\x01\x12p\n" +
	"\x04code\x18\x1e \x01(\tBW\xbaGT\x92\x02Q授权请求中收到的一次性验证/认证码。(当使用授权码模式时)H\aR\x04code\x88\x01\x01B\f\n" +
	"\n" +
	"_client_idB\x10\n" +
	"\x0e_client_secretB\b\n" +
	"\x06_scopeB\x0f\n" +
	"\r_redirect_uriB\v\n" +
	"\t_usernameB\v\n" +
	"\t_passwordB\x10\n" +
	"\x0e_refresh_tokenB\a\n" +
	"\x05_code\"\x91\b\n" +
	"\rLoginResponse\x12u\n" +
	"\faccess_token\x18\x01 \x01(\tBQ\xbaGN\x92\x02K访问令牌，必选项。授权服务器颁发的访问令牌字符串。R\faccess_token\x12\xbd\x02\n" +
	"\rrefresh_token\x18\x02 \x01(\tB\x96\x02\xbaG\x92\x02\x92\x02\x8e\x02更新令牌，用来获取下一次的访问令牌，可选项。如果访问令牌将过期，则返回刷新令牌很有用，应用程序可以使用该刷新令牌来获取另一个访问令牌。但是，通过隐式授予颁发的令牌不能颁发刷新令牌。R\rrefresh_token\x12\xb5\x01\n" +
	"\n" +
	"token_type\x18\x03 \x01(\tB\x94\x01\xbaG\x90\x01\x8a\x02\b\x1a\x06Bearer\x92\x02\x81\x01令牌的类型，该值大小写不敏感，必选项，可以是bearer类型或mac类型，通常只是字符串“Bearer”。R\n" +
	"token_type\x12\xe2\x01\n" +
	"\n" +
	"expires_in\x18\x04 \x01(\x03B\xbc\x01\xbaG\xb8\x01\x92\x02\xb4\x01令牌有效时间，单位为秒。如果访问令牌过期，服务器应回复授予访问令牌的持续时间。如果省略该参数，必须其他方式设置过期时间。H\x00R\n" +
	"expires_in\x88\x01\x01\x12\x92\x01\n" +
	"\x05scope\x18\x05 \x01(\tBw\xbaGt\x92\x02q以空格分隔的用户授予范围列表。如果未提供，scope则授权任何范围，默认为空列表。H\x01R\x05scope\x88\x01\x01B\r\n" +
	"\v_expires_inB\b\n" +
	"\x06_scope\"\x9b\x01\n" +
	"\x14ValidateTokenRequest\x12$\n" +
	"\x05token\x18\x01 \x01(\tB\f\xbaG\t\x92\x02\x06令牌R\aisValid\x12]\n" +
	"\vclient_type\x18\x02 \x01(\x0e2%.authentication.service.v1.ClientTypeB\x15\xbaG\x12\x92\x02\x0f客戶端類型R\n" +
	"clientType\"\xb8\x01\n" +
	"\x15ValidateTokenResponse\x123\n" +
	"\bis_valid\x18\x01 \x01(\bB\x18\xbaG\x15\x92\x02\x12令牌是否有效R\aisValid\x12`\n" +
	"\x05claim\x18\x02 \x01(\v2+.authentication.service.v1.UserTokenPayloadB\x18\xbaG\x15\x92\x02\x12用戶令牌载体H\x00R\x05claim\x88\x01\x01B\b\n" +
	"\x06_claim\"\xe6\x01\n" +
	"\x13RegisterUserRequest\x12+\n" +
	"\busername\x18\x01 \x01(\tB\x0f\xbaG\f\x92\x02\t用户名R\busername\x12.\n" +
	"\bpassword\x18\x02 \x01(\tB\x12\xbaG\x0f\x92\x02\f登入密码R\bpassword\x123\n" +
	"\vtenant_code\x18\x03 \x01(\tB\x12\xbaG\x0f\x92\x02\f租户代码R\n" +
	"tenantCode\x123\n" +
	"\x05email\x18\x04 \x01(\tB\x18\xbaG\x15\x92\x02\x12电子邮件地址H\x00R\x05email\x88\x01\x01B\b\n" +
	"\x06_email\"/\n" +
	"\x14RegisterUserResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\rR\x06userId\"\xe9\x03\n" +
	"\x10UserTokenPayload\x12$\n" +
	"\auser_id\x18\x01 \x01(\rB\x0e\xbaG\v\x92\x02\b用户IDR\x03uid\x12+\n" +
	"\ttenant_id\x18\x02 \x01(\rB\x0e\xbaG\v\x92\x02\b租户IDH\x00R\x03tid\x88\x01\x01\x12+\n" +
	"\busername\x18\x03 \x01(\tB\x0f\xbaG\f\x92\x02\t用户名H\x01R\x03sub\x88\x01\x01\x12.\n" +
	"\tclient_id\x18\x04 \x01(\tB\x11\xbaG\x0e\x92\x02\v客户端IDH\x02R\x03cid\x88\x01\x01\x12J\n" +
	"\tauthority\x18\x05 \x01(\x0e2\x1e.user.service.v1.UserAuthorityB\x12\xbaG\x0f\x92\x02\f用户权限R\x03aut\x12>\n" +
	"\x05roles\x18\x06 \x03(\tB*\xbaG'\x92\x02$用户角色列表，以逗号分隔R\x03roc\x12)\n" +
	"\arole_id\x18\a \x01(\rB\x0e\xbaG\v\x92\x02\b角色IDH\x03R\x03rid\x88\x01\x01\x12+\n" +
	"\tdevice_id\x18\b \x01(\tB\x0e\xbaG\v\x92\x02\b设备IDH\x04R\x03did\x88\x01\x01B\f\n" +
	"\n" +
	"_tenant_idB\v\n" +
	"\t_usernameB\f\n" +
	"\n" +
	"_client_idB\n" +
	"\n" +
	"\b_role_idB\f\n" +
	"\n" +
	"_device_id\"\xc4\x01\n" +
	"\x0eWhoAmIResponse\x12$\n" +
	"\auser_id\x18\x01 \x01(\rB\x0e\xbaG\v\x92\x02\b用户IDR\x03uid\x12:\n" +
	"\busername\x18\x02 \x01(\tB\x1e\xbaG\x1b\x92\x02\x18当前用户的用户名R\busername\x12P\n" +
	"\tauthority\x18\x03 \x01(\x0e2\x1e.user.service.v1.UserAuthorityB\x12\xbaG\x0f\x92\x02\f用户权限R\tauthority\"\xac\x01\n" +
	"\x15ChangePasswordRequest\x12+\n" +
	"\busername\x18\x01 \x01(\tB\x0f\xbaG\f\x92\x02\t用户名R\busername\x122\n" +
	"\fold_password\x18\x02 \x01(\tB\x0f\xbaG\f\x92\x02\t旧密码R\voldPassword\x122\n" +
	"\fnew_password\x18\x03 \x01(\tB\x0f\xbaG\f\x92\x02\t新密码R\vnewPassword*j\n" +
	"\tGrantType\x12\f\n" +
	"\bpassword\x10\x00\x12\x16\n" +
	"\x12client_credentials\x10\x01\x12\x16\n" +
	"\x12authorization_code\x10\x02\x12\x11\n" +
	"\rrefresh_token\x10\x03\x12\f\n" +
	"\bimplicit\x10\x04* \n" +
	"\tTokenType\x12\n" +
	"\n" +
	"\x06bearer\x10\x00\x12\a\n" +
	"\x03mac\x10\x01* \n" +
	"\n" +
	"ClientType\x12\t\n" +
	"\x05admin\x10\x00\x12\a\n" +
	"\x03app\x10\x012\xaa\x05\n" +
	"\x15AuthenticationService\x12\\\n" +
	"\x05Login\x12'.authentication.service.v1.LoginRequest\x1a(.authentication.service.v1.LoginResponse\"\x00\x12:\n" +
	"\x06Logout\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\"\x00\x12q\n" +
	"\fRegisterUser\x12..authentication.service.v1.RegisterUserRequest\x1a/.authentication.service.v1.RegisterUserResponse\"\x00\x12c\n" +
	"\fRefreshToken\x12'.authentication.service.v1.LoginRequest\x1a(.authentication.service.v1.LoginResponse\"\x00\x12t\n" +
	"\rValidateToken\x12/.authentication.service.v1.ValidateTokenRequest\x1a0.authentication.service.v1.ValidateTokenResponse\"\x00\x12Z\n" +
	"\x0eChangePassword\x120.authentication.service.v1.ChangePasswordRequest\x1a\x16.google.protobuf.Empty\x12M\n" +
	"\x06WhoAmI\x12\x16.google.protobuf.Empty\x1a).authentication.service.v1.WhoAmIResponse\"\x00B\xf7\x01\n" +
	"\x1dcom.authentication.service.v1B\x13AuthenticationProtoP\x01Z;kratos-admin/api/gen/go/authentication/service/v1;servicev1\xa2\x02\x03ASX\xaa\x02\x19Authentication.Service.V1\xca\x02\x19Authentication\\Service\\V1\xe2\x02%Authentication\\Service\\V1\\GPBMetadata\xea\x02\x1bAuthentication::Service::V1b\x06proto3"

var (
	file_authentication_service_v1_authentication_proto_rawDescOnce sync.Once
	file_authentication_service_v1_authentication_proto_rawDescData []byte
)

func file_authentication_service_v1_authentication_proto_rawDescGZIP() []byte {
	file_authentication_service_v1_authentication_proto_rawDescOnce.Do(func() {
		file_authentication_service_v1_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_authentication_service_v1_authentication_proto_rawDesc), len(file_authentication_service_v1_authentication_proto_rawDesc)))
	})
	return file_authentication_service_v1_authentication_proto_rawDescData
}

var file_authentication_service_v1_authentication_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_authentication_service_v1_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_authentication_service_v1_authentication_proto_goTypes = []any{
	(GrantType)(0),                // 0: authentication.service.v1.GrantType
	(TokenType)(0),                // 1: authentication.service.v1.TokenType
	(ClientType)(0),               // 2: authentication.service.v1.ClientType
	(*LoginRequest)(nil),          // 3: authentication.service.v1.LoginRequest
	(*LoginResponse)(nil),         // 4: authentication.service.v1.LoginResponse
	(*ValidateTokenRequest)(nil),  // 5: authentication.service.v1.ValidateTokenRequest
	(*ValidateTokenResponse)(nil), // 6: authentication.service.v1.ValidateTokenResponse
	(*RegisterUserRequest)(nil),   // 7: authentication.service.v1.RegisterUserRequest
	(*RegisterUserResponse)(nil),  // 8: authentication.service.v1.RegisterUserResponse
	(*UserTokenPayload)(nil),      // 9: authentication.service.v1.UserTokenPayload
	(*WhoAmIResponse)(nil),        // 10: authentication.service.v1.WhoAmIResponse
	(*ChangePasswordRequest)(nil), // 11: authentication.service.v1.ChangePasswordRequest
	(v1.UserAuthority)(0),         // 12: user.service.v1.UserAuthority
	(*emptypb.Empty)(nil),         // 13: google.protobuf.Empty
}
var file_authentication_service_v1_authentication_proto_depIdxs = []int32{
	2,  // 0: authentication.service.v1.ValidateTokenRequest.client_type:type_name -> authentication.service.v1.ClientType
	9,  // 1: authentication.service.v1.ValidateTokenResponse.claim:type_name -> authentication.service.v1.UserTokenPayload
	12, // 2: authentication.service.v1.UserTokenPayload.authority:type_name -> user.service.v1.UserAuthority
	12, // 3: authentication.service.v1.WhoAmIResponse.authority:type_name -> user.service.v1.UserAuthority
	3,  // 4: authentication.service.v1.AuthenticationService.Login:input_type -> authentication.service.v1.LoginRequest
	13, // 5: authentication.service.v1.AuthenticationService.Logout:input_type -> google.protobuf.Empty
	7,  // 6: authentication.service.v1.AuthenticationService.RegisterUser:input_type -> authentication.service.v1.RegisterUserRequest
	3,  // 7: authentication.service.v1.AuthenticationService.RefreshToken:input_type -> authentication.service.v1.LoginRequest
	5,  // 8: authentication.service.v1.AuthenticationService.ValidateToken:input_type -> authentication.service.v1.ValidateTokenRequest
	11, // 9: authentication.service.v1.AuthenticationService.ChangePassword:input_type -> authentication.service.v1.ChangePasswordRequest
	13, // 10: authentication.service.v1.AuthenticationService.WhoAmI:input_type -> google.protobuf.Empty
	4,  // 11: authentication.service.v1.AuthenticationService.Login:output_type -> authentication.service.v1.LoginResponse
	13, // 12: authentication.service.v1.AuthenticationService.Logout:output_type -> google.protobuf.Empty
	8,  // 13: authentication.service.v1.AuthenticationService.RegisterUser:output_type -> authentication.service.v1.RegisterUserResponse
	4,  // 14: authentication.service.v1.AuthenticationService.RefreshToken:output_type -> authentication.service.v1.LoginResponse
	6,  // 15: authentication.service.v1.AuthenticationService.ValidateToken:output_type -> authentication.service.v1.ValidateTokenResponse
	13, // 16: authentication.service.v1.AuthenticationService.ChangePassword:output_type -> google.protobuf.Empty
	10, // 17: authentication.service.v1.AuthenticationService.WhoAmI:output_type -> authentication.service.v1.WhoAmIResponse
	11, // [11:18] is the sub-list for method output_type
	4,  // [4:11] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_authentication_service_v1_authentication_proto_init() }
func file_authentication_service_v1_authentication_proto_init() {
	if File_authentication_service_v1_authentication_proto != nil {
		return
	}
	file_authentication_service_v1_authentication_proto_msgTypes[0].OneofWrappers = []any{}
	file_authentication_service_v1_authentication_proto_msgTypes[1].OneofWrappers = []any{}
	file_authentication_service_v1_authentication_proto_msgTypes[3].OneofWrappers = []any{}
	file_authentication_service_v1_authentication_proto_msgTypes[4].OneofWrappers = []any{}
	file_authentication_service_v1_authentication_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_authentication_service_v1_authentication_proto_rawDesc), len(file_authentication_service_v1_authentication_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authentication_service_v1_authentication_proto_goTypes,
		DependencyIndexes: file_authentication_service_v1_authentication_proto_depIdxs,
		EnumInfos:         file_authentication_service_v1_authentication_proto_enumTypes,
		MessageInfos:      file_authentication_service_v1_authentication_proto_msgTypes,
	}.Build()
	File_authentication_service_v1_authentication_proto = out.File
	file_authentication_service_v1_authentication_proto_goTypes = nil
	file_authentication_service_v1_authentication_proto_depIdxs = nil
}
