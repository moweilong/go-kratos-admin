// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: admin/service/v1/i_router.proto

package servicev1

import (
	_ "github.com/google/gnostic/openapiv3"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// 路由项
type RouteItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Children      []*RouteItem           `protobuf:"bytes,1,rep,name=children,proto3" json:"children,omitempty"`          // 子节点树
	Path          *string                `protobuf:"bytes,10,opt,name=path,proto3,oneof" json:"path,omitempty"`           // 路由路径
	Redirect      *string                `protobuf:"bytes,11,opt,name=redirect,proto3,oneof" json:"redirect,omitempty"`   // 重定向地址
	Alias         *string                `protobuf:"bytes,12,opt,name=alias,proto3,oneof" json:"alias,omitempty"`         // 路由别名
	Name          *string                `protobuf:"bytes,13,opt,name=name,proto3,oneof" json:"name,omitempty"`           // 路由命名，然后我们可以使用 name 而不是 path 来传递 to 属性给 <router-link>。
	Component     *string                `protobuf:"bytes,14,opt,name=component,proto3,oneof" json:"component,omitempty"` // 指向的组件
	Meta          *RouteMeta             `protobuf:"bytes,15,opt,name=meta,proto3,oneof" json:"meta,omitempty"`           // 路由元信息
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RouteItem) Reset() {
	*x = RouteItem{}
	mi := &file_admin_service_v1_i_router_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteItem) ProtoMessage() {}

func (x *RouteItem) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_router_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteItem.ProtoReflect.Descriptor instead.
func (*RouteItem) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_router_proto_rawDescGZIP(), []int{0}
}

func (x *RouteItem) GetChildren() []*RouteItem {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *RouteItem) GetPath() string {
	if x != nil && x.Path != nil {
		return *x.Path
	}
	return ""
}

func (x *RouteItem) GetRedirect() string {
	if x != nil && x.Redirect != nil {
		return *x.Redirect
	}
	return ""
}

func (x *RouteItem) GetAlias() string {
	if x != nil && x.Alias != nil {
		return *x.Alias
	}
	return ""
}

func (x *RouteItem) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *RouteItem) GetComponent() string {
	if x != nil && x.Component != nil {
		return *x.Component
	}
	return ""
}

func (x *RouteItem) GetMeta() *RouteMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// 路由元数据
type RouteMeta struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	ActiveIcon               *string                `protobuf:"bytes,1,opt,name=activeIcon,proto3,oneof" json:"activeIcon,omitempty"`                               // 激活图标，用于：菜单、tab
	ActivePath               *string                `protobuf:"bytes,2,opt,name=activePath,proto3,oneof" json:"activePath,omitempty"`                               // 当前激活的菜单，有时候不想激活现有菜单，需要激活父级菜单时使用
	AffixTab                 *bool                  `protobuf:"varint,3,opt,name=affixTab,proto3,oneof" json:"affixTab,omitempty"`                                  // 是否固定标签页
	AffixTabOrder            *int32                 `protobuf:"varint,4,opt,name=affixTabOrder,proto3,oneof" json:"affixTabOrder,omitempty"`                        // 固定标签页的顺序
	Authority                []string               `protobuf:"bytes,5,rep,name=authority,proto3" json:"authority,omitempty"`                                       // 权限列表，需要特定的角色标识才可以访问
	Badge                    *string                `protobuf:"bytes,6,opt,name=badge,proto3,oneof" json:"badge,omitempty"`                                         // 徽标
	BadgeType                *string                `protobuf:"bytes,7,opt,name=badgeType,proto3,oneof" json:"badgeType,omitempty"`                                 // 徽标类型
	BadgeVariants            *string                `protobuf:"bytes,8,opt,name=badgeVariants,proto3,oneof" json:"badgeVariants,omitempty"`                         // 徽标颜色
	HideChildrenInMenu       *bool                  `protobuf:"varint,9,opt,name=hideChildrenInMenu,proto3,oneof" json:"hideChildrenInMenu,omitempty"`              // 当前路由的子级在菜单中不展现
	HideInBreadcrumb         *bool                  `protobuf:"varint,10,opt,name=hideInBreadcrumb,proto3,oneof" json:"hideInBreadcrumb,omitempty"`                 // 当前路由在面包屑中不展现
	HideInMenu               *bool                  `protobuf:"varint,11,opt,name=hideInMenu,proto3,oneof" json:"hideInMenu,omitempty"`                             // 当前路由在菜单中不展现
	HideInTab                *bool                  `protobuf:"varint,12,opt,name=hideInTab,proto3,oneof" json:"hideInTab,omitempty"`                               // 当前路由在标签页不展现
	Icon                     *string                `protobuf:"bytes,13,opt,name=icon,proto3,oneof" json:"icon,omitempty"`                                          // 图标，用于：菜单、标签页
	IframeSrc                *string                `protobuf:"bytes,14,opt,name=iframeSrc,proto3,oneof" json:"iframeSrc,omitempty"`                                // iframe 地址
	IgnoreAccess             *bool                  `protobuf:"varint,15,opt,name=ignoreAccess,proto3,oneof" json:"ignoreAccess,omitempty"`                         // 忽略权限，直接可以访问
	KeepAlive                *bool                  `protobuf:"varint,16,opt,name=keepAlive,proto3,oneof" json:"keepAlive,omitempty"`                               // 开启KeepAlive缓存
	Link                     *string                `protobuf:"bytes,17,opt,name=link,proto3,oneof" json:"link,omitempty"`                                          // 外链-跳转路径
	Loaded                   *bool                  `protobuf:"varint,18,opt,name=loaded,proto3,oneof" json:"loaded,omitempty"`                                     // 路由是否已经加载过
	MaxNumOfOpenTab          *int32                 `protobuf:"varint,19,opt,name=maxNumOfOpenTab,proto3,oneof" json:"maxNumOfOpenTab,omitempty"`                   // 标签页最大打开数量
	MenuVisibleWithForbidden *bool                  `protobuf:"varint,20,opt,name=menuVisibleWithForbidden,proto3,oneof" json:"menuVisibleWithForbidden,omitempty"` // 菜单可以看到，但是访问会被重定向到403
	OpenInNewWindow          *bool                  `protobuf:"varint,21,opt,name=openInNewWindow,proto3,oneof" json:"openInNewWindow,omitempty"`                   // 在新窗口打开
	Order                    *int32                 `protobuf:"varint,22,opt,name=order,proto3,oneof" json:"order,omitempty"`                                       // 排序编号，用于路由->菜单排序
	Title                    *string                `protobuf:"bytes,23,opt,name=title,proto3,oneof" json:"title,omitempty"`                                        // 标题名称，路由上显示的标题
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *RouteMeta) Reset() {
	*x = RouteMeta{}
	mi := &file_admin_service_v1_i_router_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteMeta) ProtoMessage() {}

func (x *RouteMeta) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_router_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteMeta.ProtoReflect.Descriptor instead.
func (*RouteMeta) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_router_proto_rawDescGZIP(), []int{1}
}

func (x *RouteMeta) GetActiveIcon() string {
	if x != nil && x.ActiveIcon != nil {
		return *x.ActiveIcon
	}
	return ""
}

func (x *RouteMeta) GetActivePath() string {
	if x != nil && x.ActivePath != nil {
		return *x.ActivePath
	}
	return ""
}

func (x *RouteMeta) GetAffixTab() bool {
	if x != nil && x.AffixTab != nil {
		return *x.AffixTab
	}
	return false
}

func (x *RouteMeta) GetAffixTabOrder() int32 {
	if x != nil && x.AffixTabOrder != nil {
		return *x.AffixTabOrder
	}
	return 0
}

func (x *RouteMeta) GetAuthority() []string {
	if x != nil {
		return x.Authority
	}
	return nil
}

func (x *RouteMeta) GetBadge() string {
	if x != nil && x.Badge != nil {
		return *x.Badge
	}
	return ""
}

func (x *RouteMeta) GetBadgeType() string {
	if x != nil && x.BadgeType != nil {
		return *x.BadgeType
	}
	return ""
}

func (x *RouteMeta) GetBadgeVariants() string {
	if x != nil && x.BadgeVariants != nil {
		return *x.BadgeVariants
	}
	return ""
}

func (x *RouteMeta) GetHideChildrenInMenu() bool {
	if x != nil && x.HideChildrenInMenu != nil {
		return *x.HideChildrenInMenu
	}
	return false
}

func (x *RouteMeta) GetHideInBreadcrumb() bool {
	if x != nil && x.HideInBreadcrumb != nil {
		return *x.HideInBreadcrumb
	}
	return false
}

func (x *RouteMeta) GetHideInMenu() bool {
	if x != nil && x.HideInMenu != nil {
		return *x.HideInMenu
	}
	return false
}

func (x *RouteMeta) GetHideInTab() bool {
	if x != nil && x.HideInTab != nil {
		return *x.HideInTab
	}
	return false
}

func (x *RouteMeta) GetIcon() string {
	if x != nil && x.Icon != nil {
		return *x.Icon
	}
	return ""
}

func (x *RouteMeta) GetIframeSrc() string {
	if x != nil && x.IframeSrc != nil {
		return *x.IframeSrc
	}
	return ""
}

func (x *RouteMeta) GetIgnoreAccess() bool {
	if x != nil && x.IgnoreAccess != nil {
		return *x.IgnoreAccess
	}
	return false
}

func (x *RouteMeta) GetKeepAlive() bool {
	if x != nil && x.KeepAlive != nil {
		return *x.KeepAlive
	}
	return false
}

func (x *RouteMeta) GetLink() string {
	if x != nil && x.Link != nil {
		return *x.Link
	}
	return ""
}

func (x *RouteMeta) GetLoaded() bool {
	if x != nil && x.Loaded != nil {
		return *x.Loaded
	}
	return false
}

func (x *RouteMeta) GetMaxNumOfOpenTab() int32 {
	if x != nil && x.MaxNumOfOpenTab != nil {
		return *x.MaxNumOfOpenTab
	}
	return 0
}

func (x *RouteMeta) GetMenuVisibleWithForbidden() bool {
	if x != nil && x.MenuVisibleWithForbidden != nil {
		return *x.MenuVisibleWithForbidden
	}
	return false
}

func (x *RouteMeta) GetOpenInNewWindow() bool {
	if x != nil && x.OpenInNewWindow != nil {
		return *x.OpenInNewWindow
	}
	return false
}

func (x *RouteMeta) GetOrder() int32 {
	if x != nil && x.Order != nil {
		return *x.Order
	}
	return 0
}

func (x *RouteMeta) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

type ListRouteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRouteRequest) Reset() {
	*x = ListRouteRequest{}
	mi := &file_admin_service_v1_i_router_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRouteRequest) ProtoMessage() {}

func (x *ListRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_router_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRouteRequest.ProtoReflect.Descriptor instead.
func (*ListRouteRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_router_proto_rawDescGZIP(), []int{2}
}

// 查询路由列表 - 回应
type ListRouteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*RouteItem           `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListRouteResponse) Reset() {
	*x = ListRouteResponse{}
	mi := &file_admin_service_v1_i_router_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRouteResponse) ProtoMessage() {}

func (x *ListRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_router_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRouteResponse.ProtoReflect.Descriptor instead.
func (*ListRouteResponse) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_router_proto_rawDescGZIP(), []int{3}
}

func (x *ListRouteResponse) GetItems() []*RouteItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListPermissionCodeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPermissionCodeRequest) Reset() {
	*x = ListPermissionCodeRequest{}
	mi := &file_admin_service_v1_i_router_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPermissionCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionCodeRequest) ProtoMessage() {}

func (x *ListPermissionCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_router_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionCodeRequest.ProtoReflect.Descriptor instead.
func (*ListPermissionCodeRequest) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_router_proto_rawDescGZIP(), []int{4}
}

// 查询权限码列表 - 回应
type ListPermissionCodeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Codes         []string               `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPermissionCodeResponse) Reset() {
	*x = ListPermissionCodeResponse{}
	mi := &file_admin_service_v1_i_router_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPermissionCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPermissionCodeResponse) ProtoMessage() {}

func (x *ListPermissionCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_service_v1_i_router_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPermissionCodeResponse.ProtoReflect.Descriptor instead.
func (*ListPermissionCodeResponse) Descriptor() ([]byte, []int) {
	return file_admin_service_v1_i_router_proto_rawDescGZIP(), []int{5}
}

func (x *ListPermissionCodeResponse) GetCodes() []string {
	if x != nil {
		return x.Codes
	}
	return nil
}

var File_admin_service_v1_i_router_proto protoreflect.FileDescriptor

const file_admin_service_v1_i_router_proto_rawDesc = "" +
	"\n" +
	"\x1fadmin/service/v1/i_router.proto\x12\x10admin.service.v1\x1a$gnostic/openapi/v3/annotations.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a google/protobuf/field_mask.proto\"\xca\x04\n" +
	"\tRouteItem\x12K\n" +
	"\bchildren\x18\x01 \x03(\v2\x1b.admin.service.v1.RouteItemB\x12\xbaG\x0f\x92\x02\f子节点树R\bchildren\x12.\n" +
	"\x04path\x18\n" +
	" \x01(\tB\x15\xe0A\x01\xbaG\x0f\x92\x02\f路由路径H\x00R\x04path\x88\x01\x01\x129\n" +
	"\bredirect\x18\v \x01(\tB\x18\xe0A\x01\xbaG\x12\x92\x02\x0f重定向地址H\x01R\bredirect\x88\x01\x01\x120\n" +
	"\x05alias\x18\f \x01(\tB\x15\xe0A\x01\xbaG\x0f\x92\x02\f路由别名H\x02R\x05alias\x88\x01\x01\x12\x85\x01\n" +
	"\x04name\x18\r \x01(\tBl\xe0A\x01\xbaGf\x92\x02c路由命名，然后我们可以使用 name 而不是 path 来传递 to 属性给 <router-link>。H\x03R\x04name\x88\x01\x01\x12;\n" +
	"\tcomponent\x18\x0e \x01(\tB\x18\xe0A\x01\xbaG\x12\x92\x02\x0f指向的组件H\x04R\tcomponent\x88\x01\x01\x12N\n" +
	"\x04meta\x18\x0f \x01(\v2\x1b.admin.service.v1.RouteMetaB\x18\xe0A\x01\xbaG\x12\x92\x02\x0f路由元信息H\x05R\x04meta\x88\x01\x01B\a\n" +
	"\x05_pathB\v\n" +
	"\t_redirectB\b\n" +
	"\x06_aliasB\a\n" +
	"\x05_nameB\f\n" +
	"\n" +
	"_componentB\a\n" +
	"\x05_meta\"\xec\x11\n" +
	"\tRouteMeta\x12R\n" +
	"\n" +
	"activeIcon\x18\x01 \x01(\tB-\xe0A\x01\xbaG'\x92\x02$激活图标，用于：菜单、tabH\x00R\n" +
	"activeIcon\x88\x01\x01\x12\x8b\x01\n" +
	"\n" +
	"activePath\x18\x02 \x01(\tBf\xe0A\x01\xbaG`\x92\x02]当前激活的菜单，有时候不想激活现有菜单，需要激活父级菜单时使用H\x01R\n" +
	"activePath\x88\x01\x01\x12?\n" +
	"\baffixTab\x18\x03 \x01(\bB\x1e\xe0A\x01\xbaG\x18\x92\x02\x15是否固定标签页H\x02R\baffixTab\x88\x01\x01\x12L\n" +
	"\raffixTabOrder\x18\x04 \x01(\x05B!\xe0A\x01\xbaG\x1b\x92\x02\x18固定标签页的顺序H\x03R\raffixTabOrder\x88\x01\x01\x12`\n" +
	"\tauthority\x18\x05 \x03(\tBB\xe0A\x01\xbaG<\x92\x029权限列表，需要特定的角色标识才可以访问R\tauthority\x12*\n" +
	"\x05badge\x18\x06 \x01(\tB\x0f\xe0A\x01\xbaG\t\x92\x02\x06徽标H\x04R\x05badge\x88\x01\x01\x12K\n" +
	"\tbadgeType\x18\a \x01(\tB(\xe0A\x01\xbaG\"\xc2\x01\x05\x12\x03dot\xc2\x01\b\x12\x06normal\x92\x02\f徽标类型H\x05R\tbadgeType\x88\x01\x01\x12\x80\x01\n" +
	"\rbadgeVariants\x18\b \x01(\tBU\xe0A\x01\xbaGO\xc2\x01\t\x12\adefault\xc2\x01\r\x12\vdestructive\xc2\x01\t\x12\aprimary\xc2\x01\t\x12\asuccess\xc2\x01\t\x12\awarning\x92\x02\f徽标颜色H\x06R\rbadgeVariants\x88\x01\x01\x12h\n" +
	"\x12hideChildrenInMenu\x18\t \x01(\bB3\xe0A\x01\xbaG-\x92\x02*当前路由的子级在菜单中不展现H\aR\x12hideChildrenInMenu\x88\x01\x01\x12^\n" +
	"\x10hideInBreadcrumb\x18\n" +
	" \x01(\bB-\xe0A\x01\xbaG'\x92\x02$当前路由在面包屑中不展现H\bR\x10hideInBreadcrumb\x88\x01\x01\x12O\n" +
	"\n" +
	"hideInMenu\x18\v \x01(\bB*\xe0A\x01\xbaG$\x92\x02!当前路由在菜单中不展现H\tR\n" +
	"hideInMenu\x88\x01\x01\x12M\n" +
	"\thideInTab\x18\f \x01(\bB*\xe0A\x01\xbaG$\x92\x02!当前路由在标签页不展现H\n" +
	"R\thideInTab\x88\x01\x01\x12F\n" +
	"\x04icon\x18\r \x01(\tB-\xe0A\x01\xbaG'\x92\x02$图标，用于：菜单、标签页H\vR\x04icon\x88\x01\x01\x129\n" +
	"\tiframeSrc\x18\x0e \x01(\tB\x16\xe0A\x01\xbaG\x10\x92\x02\riframe 地址H\fR\tiframeSrc\x88\x01\x01\x12S\n" +
	"\fignoreAccess\x18\x0f \x01(\bB*\xe0A\x01\xbaG$\x92\x02!忽略权限，直接可以访问H\rR\fignoreAccess\x88\x01\x01\x12A\n" +
	"\tkeepAlive\x18\x10 \x01(\bB\x1e\xe0A\x01\xbaG\x18\x92\x02\x15开启KeepAlive缓存H\x0eR\tkeepAlive\x88\x01\x01\x125\n" +
	"\x04link\x18\x11 \x01(\tB\x1c\xe0A\x01\xbaG\x16\x92\x02\x13外链-跳转路径H\x0fR\x04link\x88\x01\x01\x12A\n" +
	"\x06loaded\x18\x12 \x01(\bB$\xe0A\x01\xbaG\x1e\x92\x02\x1b路由是否已经加载过H\x10R\x06loaded\x88\x01\x01\x12S\n" +
	"\x0fmaxNumOfOpenTab\x18\x13 \x01(\x05B$\xe0A\x01\xbaG\x1e\x92\x02\x1b标签页最大打开数量H\x11R\x0fmaxNumOfOpenTab\x88\x01\x01\x12\x80\x01\n" +
	"\x18menuVisibleWithForbidden\x18\x14 \x01(\bB?\xe0A\x01\xbaG9\x92\x026菜单可以看到，但是访问会被重定向到403H\x12R\x18menuVisibleWithForbidden\x88\x01\x01\x12J\n" +
	"\x0fopenInNewWindow\x18\x15 \x01(\bB\x1b\xe0A\x01\xbaG\x15\x92\x02\x12在新窗口打开H\x13R\x0fopenInNewWindow\x88\x01\x01\x12M\n" +
	"\x05order\x18\x16 \x01(\x05B2\xe0A\x01\xbaG,\x92\x02)排序编号，用于路由->菜单排序H\x14R\x05order\x88\x01\x01\x12K\n" +
	"\x05title\x18\x17 \x01(\tB0\xe0A\x01\xbaG*\x92\x02'标题名称，路由上显示的标题H\x15R\x05title\x88\x01\x01B\r\n" +
	"\v_activeIconB\r\n" +
	"\v_activePathB\v\n" +
	"\t_affixTabB\x10\n" +
	"\x0e_affixTabOrderB\b\n" +
	"\x06_badgeB\f\n" +
	"\n" +
	"_badgeTypeB\x10\n" +
	"\x0e_badgeVariantsB\x15\n" +
	"\x13_hideChildrenInMenuB\x13\n" +
	"\x11_hideInBreadcrumbB\r\n" +
	"\v_hideInMenuB\f\n" +
	"\n" +
	"_hideInTabB\a\n" +
	"\x05_iconB\f\n" +
	"\n" +
	"_iframeSrcB\x0f\n" +
	"\r_ignoreAccessB\f\n" +
	"\n" +
	"_keepAliveB\a\n" +
	"\x05_linkB\t\n" +
	"\a_loadedB\x12\n" +
	"\x10_maxNumOfOpenTabB\x1b\n" +
	"\x19_menuVisibleWithForbiddenB\x12\n" +
	"\x10_openInNewWindowB\b\n" +
	"\x06_orderB\b\n" +
	"\x06_title\"\x12\n" +
	"\x10ListRouteRequest\"F\n" +
	"\x11ListRouteResponse\x121\n" +
	"\x05items\x18\x01 \x03(\v2\x1b.admin.service.v1.RouteItemR\x05items\"\x1b\n" +
	"\x19ListPermissionCodeRequest\"2\n" +
	"\x1aListPermissionCodeResponse\x12\x14\n" +
	"\x05codes\x18\x01 \x03(\tR\x05codes2\xed\x01\n" +
	"\rRouterService\x12b\n" +
	"\tListRoute\x12\x16.google.protobuf.Empty\x1a#.admin.service.v1.ListRouteResponse\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/admin/v1/routes\x12x\n" +
	"\x12ListPermissionCode\x12\x16.google.protobuf.Empty\x1a,.admin.service.v1.ListPermissionCodeResponse\"\x1c\x82\xd3\xe4\x93\x02\x16\x12\x14/admin/v1/perm-codesB\xba\x01\n" +
	"\x14com.admin.service.v1B\fIRouterProtoP\x01Z2kratos-admin/api/gen/go/admin/service/v1;servicev1\xa2\x02\x03ASX\xaa\x02\x10Admin.Service.V1\xca\x02\x10Admin\\Service\\V1\xe2\x02\x1cAdmin\\Service\\V1\\GPBMetadata\xea\x02\x12Admin::Service::V1b\x06proto3"

var (
	file_admin_service_v1_i_router_proto_rawDescOnce sync.Once
	file_admin_service_v1_i_router_proto_rawDescData []byte
)

func file_admin_service_v1_i_router_proto_rawDescGZIP() []byte {
	file_admin_service_v1_i_router_proto_rawDescOnce.Do(func() {
		file_admin_service_v1_i_router_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_admin_service_v1_i_router_proto_rawDesc), len(file_admin_service_v1_i_router_proto_rawDesc)))
	})
	return file_admin_service_v1_i_router_proto_rawDescData
}

var file_admin_service_v1_i_router_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_admin_service_v1_i_router_proto_goTypes = []any{
	(*RouteItem)(nil),                  // 0: admin.service.v1.RouteItem
	(*RouteMeta)(nil),                  // 1: admin.service.v1.RouteMeta
	(*ListRouteRequest)(nil),           // 2: admin.service.v1.ListRouteRequest
	(*ListRouteResponse)(nil),          // 3: admin.service.v1.ListRouteResponse
	(*ListPermissionCodeRequest)(nil),  // 4: admin.service.v1.ListPermissionCodeRequest
	(*ListPermissionCodeResponse)(nil), // 5: admin.service.v1.ListPermissionCodeResponse
	(*emptypb.Empty)(nil),              // 6: google.protobuf.Empty
}
var file_admin_service_v1_i_router_proto_depIdxs = []int32{
	0, // 0: admin.service.v1.RouteItem.children:type_name -> admin.service.v1.RouteItem
	1, // 1: admin.service.v1.RouteItem.meta:type_name -> admin.service.v1.RouteMeta
	0, // 2: admin.service.v1.ListRouteResponse.items:type_name -> admin.service.v1.RouteItem
	6, // 3: admin.service.v1.RouterService.ListRoute:input_type -> google.protobuf.Empty
	6, // 4: admin.service.v1.RouterService.ListPermissionCode:input_type -> google.protobuf.Empty
	3, // 5: admin.service.v1.RouterService.ListRoute:output_type -> admin.service.v1.ListRouteResponse
	5, // 6: admin.service.v1.RouterService.ListPermissionCode:output_type -> admin.service.v1.ListPermissionCodeResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_admin_service_v1_i_router_proto_init() }
func file_admin_service_v1_i_router_proto_init() {
	if File_admin_service_v1_i_router_proto != nil {
		return
	}
	file_admin_service_v1_i_router_proto_msgTypes[0].OneofWrappers = []any{}
	file_admin_service_v1_i_router_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_admin_service_v1_i_router_proto_rawDesc), len(file_admin_service_v1_i_router_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_service_v1_i_router_proto_goTypes,
		DependencyIndexes: file_admin_service_v1_i_router_proto_depIdxs,
		MessageInfos:      file_admin_service_v1_i_router_proto_msgTypes,
	}.Build()
	File_admin_service_v1_i_router_proto = out.File
	file_admin_service_v1_i_router_proto_goTypes = nil
	file_admin_service_v1_i_router_proto_depIdxs = nil
}
