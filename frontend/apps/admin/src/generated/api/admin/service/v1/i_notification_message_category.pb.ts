// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.3
//   protoc               unknown
// source: admin/service/v1/i_notification_message_category.proto

/* eslint-disable */
import type { Empty } from "../../../google/protobuf/empty.pb";
import type {
  CreateNotificationMessageCategoryRequest,
  DeleteNotificationMessageCategoryRequest,
  GetNotificationMessageCategoryRequest,
  ListNotificationMessageCategoryResponse,
  NotificationMessageCategory,
  UpdateNotificationMessageCategoryRequest,
} from "../../../internal_message/service/v1/notification_message_category.pb";
import type { PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 通知消息分类管理服务 */
export interface NotificationMessageCategoryService {
  /** 查询通知消息分类列表 */
  List(request: PagingRequest): Promise<ListNotificationMessageCategoryResponse>;
  /** 查询通知消息分类详情 */
  Get(request: GetNotificationMessageCategoryRequest): Promise<NotificationMessageCategory>;
  /** 创建通知消息分类 */
  Create(request: CreateNotificationMessageCategoryRequest): Promise<Empty>;
  /** 更新通知消息分类 */
  Update(request: UpdateNotificationMessageCategoryRequest): Promise<Empty>;
  /** 删除通知消息分类 */
  Delete(request: DeleteNotificationMessageCategoryRequest): Promise<Empty>;
}
