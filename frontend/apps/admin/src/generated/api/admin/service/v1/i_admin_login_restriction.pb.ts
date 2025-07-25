// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.3
//   protoc               unknown
// source: admin/service/v1/i_admin_login_restriction.proto

/* eslint-disable */
import type { Empty } from "../../../google/protobuf/empty.pb";
import type { Timestamp } from "../../../google/protobuf/timestamp.pb";
import type { PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 后台登录限制类型 */
export enum AdminLoginRestrictionType {
  /** LOGIN_RESTRICTION_TYPE_UNSPECIFIED - 未知 */
  LOGIN_RESTRICTION_TYPE_UNSPECIFIED = "LOGIN_RESTRICTION_TYPE_UNSPECIFIED",
  /** BLACKLIST - 黑名单 */
  BLACKLIST = "BLACKLIST",
  /** WHITELIST - 白名单 */
  WHITELIST = "WHITELIST",
}

/** 后台登录限制方式 */
export enum AdminLoginRestrictionMethod {
  /** LOGIN_RESTRICTION_METHOD_UNSPECIFIED - 未知 */
  LOGIN_RESTRICTION_METHOD_UNSPECIFIED = "LOGIN_RESTRICTION_METHOD_UNSPECIFIED",
  /** IP - IP地址限制 */
  IP = "IP",
  /** MAC - MAC地址限制，绑定设备的MAC地址。 */
  MAC = "MAC",
  /** REGION - 地区限制，根据地理位置（如国家、城市）限制登录。 */
  REGION = "REGION",
  /** TIME - 时间限制，限制登录的时间段，例如只允许工作时间登录。 */
  TIME = "TIME",
  /** DEVICE - 设备限制，限制登录设备的类型（如PC、手机）或特定设备ID。 */
  DEVICE = "DEVICE",
}

/** 后台登录限制 */
export interface AdminLoginRestriction {
  /** 后台登录限制ID */
  id?:
    | number
    | null
    | undefined;
  /** 目标用户ID */
  targetId?:
    | number
    | null
    | undefined;
  /** 限制类型 */
  type?:
    | AdminLoginRestrictionType
    | null
    | undefined;
  /** 限制方式 */
  method?:
    | AdminLoginRestrictionMethod
    | null
    | undefined;
  /** 限制值（如IP地址、MAC地址或地区代码） */
  value?:
    | string
    | null
    | undefined;
  /** 限制原因 */
  reason?:
    | string
    | null
    | undefined;
  /** 创建者ID */
  createBy?:
    | number
    | null
    | undefined;
  /** 更新者ID */
  updateBy?:
    | number
    | null
    | undefined;
  /** 创建时间 */
  createTime?:
    | Timestamp
    | null
    | undefined;
  /** 更新时间 */
  updateTime?:
    | Timestamp
    | null
    | undefined;
  /** 删除时间 */
  deleteTime?: Timestamp | null | undefined;
}

/** 查询后台登录限制列表 - 回应 */
export interface ListAdminLoginRestrictionResponse {
  items: AdminLoginRestriction[];
  total: number;
}

/** 查询后台登录限制详情 - 请求 */
export interface GetAdminLoginRestrictionRequest {
  id: number;
}

/** 创建后台登录限制 - 请求 */
export interface CreateAdminLoginRestrictionRequest {
  data: AdminLoginRestriction | null;
}

/** 更新后台登录限制 - 请求 */
export interface UpdateAdminLoginRestrictionRequest {
  data:
    | AdminLoginRestriction
    | null;
  /** 要更新的字段列表 */
  updateMask:
    | string[]
    | null;
  /** 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。 */
  allowMissing?: boolean | null | undefined;
}

/** 删除后台登录限制 - 请求 */
export interface DeleteAdminLoginRestrictionRequest {
  id: number;
}

/** 后台登录限制管理服务 */
export interface AdminLoginRestrictionService {
  /** 查询后台登录限制列表 */
  List(request: PagingRequest): Promise<ListAdminLoginRestrictionResponse>;
  /** 查询后台登录限制详情 */
  Get(request: GetAdminLoginRestrictionRequest): Promise<AdminLoginRestriction>;
  /** 创建后台登录限制 */
  Create(request: CreateAdminLoginRestrictionRequest): Promise<Empty>;
  /** 更新后台登录限制 */
  Update(request: UpdateAdminLoginRestrictionRequest): Promise<Empty>;
  /** 删除后台登录限制 */
  Delete(request: DeleteAdminLoginRestrictionRequest): Promise<Empty>;
}
