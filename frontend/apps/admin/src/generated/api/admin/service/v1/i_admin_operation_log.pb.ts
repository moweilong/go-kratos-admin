// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.3
//   protoc               unknown
// source: admin/service/v1/i_admin_operation_log.proto

/* eslint-disable */
import type { Duration } from "../../../google/protobuf/duration.pb";
import type { Timestamp } from "../../../google/protobuf/timestamp.pb";
import type { PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 后台操作日志 */
export interface AdminOperationLog {
  /** 后台操作日志ID */
  id?:
    | number
    | null
    | undefined;
  /** 操作耗时 */
  costTime?:
    | Duration
    | null
    | undefined;
  /** 操作是否成功 */
  success?:
    | boolean
    | null
    | undefined;
  /** 请求ID */
  requestId?:
    | string
    | null
    | undefined;
  /** 状态码 */
  statusCode?:
    | number
    | null
    | undefined;
  /** 操作失败原因 */
  reason?:
    | string
    | null
    | undefined;
  /** 操作地理位置 */
  location?:
    | string
    | null
    | undefined;
  /** 操作方法 */
  operation?:
    | string
    | null
    | undefined;
  /** 请求方法 */
  method?:
    | string
    | null
    | undefined;
  /** 请求路径 */
  path?:
    | string
    | null
    | undefined;
  /** API所属模块 */
  apiModule?:
    | string
    | null
    | undefined;
  /** API操作描述 */
  apiDescription?:
    | string
    | null
    | undefined;
  /** 请求源 */
  referer?:
    | string
    | null
    | undefined;
  /** 请求URI */
  requestUri?:
    | string
    | null
    | undefined;
  /** 请求头 */
  requestHeader?:
    | string
    | null
    | undefined;
  /** 请求体 */
  requestBody?:
    | string
    | null
    | undefined;
  /** 响应信息 */
  response?:
    | string
    | null
    | undefined;
  /** 操作者用户ID */
  userId?:
    | number
    | null
    | undefined;
  /** 操作者账号名 */
  username?:
    | string
    | null
    | undefined;
  /** 操作者IP */
  clientIp?:
    | string
    | null
    | undefined;
  /** 浏览器的用户代理信息 */
  userAgent?:
    | string
    | null
    | undefined;
  /** 浏览器名称 */
  browserName?:
    | string
    | null
    | undefined;
  /** 浏览器版本 */
  browserVersion?:
    | string
    | null
    | undefined;
  /** 客户端ID */
  clientId?:
    | string
    | null
    | undefined;
  /** 客户端名称 */
  clientName?:
    | string
    | null
    | undefined;
  /** 操作系统名称 */
  osName?:
    | string
    | null
    | undefined;
  /** 操作系统版本 */
  osVersion?:
    | string
    | null
    | undefined;
  /** 创建时间 */
  createTime?: Timestamp | null | undefined;
}

/** 查询后台操作日志列表 - 回应 */
export interface ListAdminOperationLogResponse {
  items: AdminOperationLog[];
  total: number;
}

/** 查询后台操作日志详情 - 请求 */
export interface GetAdminOperationLogRequest {
  id: number;
}

/** 创建后台操作日志 - 请求 */
export interface CreateAdminOperationLogRequest {
  data: AdminOperationLog | null;
}

/** 更新后台操作日志 - 请求 */
export interface UpdateAdminOperationLogRequest {
  data:
    | AdminOperationLog
    | null;
  /** 要更新的字段列表 */
  updateMask:
    | string[]
    | null;
  /** 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。 */
  allowMissing?: boolean | null | undefined;
}

/** 删除后台操作日志 - 请求 */
export interface DeleteAdminOperationLogRequest {
  id: number;
}

/** 后台操作日志管理服务 */
export interface AdminOperationLogService {
  /** 查询后台操作日志列表 */
  List(request: PagingRequest): Promise<ListAdminOperationLogResponse>;
  /** 查询后台操作日志详情 */
  Get(request: GetAdminOperationLogRequest): Promise<AdminOperationLog>;
}
