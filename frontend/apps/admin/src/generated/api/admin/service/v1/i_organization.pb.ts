// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.3
//   protoc               unknown
// source: admin/service/v1/i_organization.proto

/* eslint-disable */
import type { Empty } from "../../../google/protobuf/empty.pb";
import type { PagingRequest } from "../../../pagination/v1/pagination.pb";
import type {
  CreateOrganizationRequest,
  DeleteOrganizationRequest,
  GetOrganizationRequest,
  ListOrganizationResponse,
  Organization,
  UpdateOrganizationRequest,
} from "../../../user/service/v1/organization.pb";

/** 组织管理服务 */
export interface OrganizationService {
  /** 查询组织列表 */
  List(request: PagingRequest): Promise<ListOrganizationResponse>;
  /** 查询组织详情 */
  Get(request: GetOrganizationRequest): Promise<Organization>;
  /** 创建组织 */
  Create(request: CreateOrganizationRequest): Promise<Empty>;
  /** 更新组织 */
  Update(request: UpdateOrganizationRequest): Promise<Empty>;
  /** 删除组织 */
  Delete(request: DeleteOrganizationRequest): Promise<Empty>;
}
