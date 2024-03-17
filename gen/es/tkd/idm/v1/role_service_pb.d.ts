// @generated by protoc-gen-es v1.8.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/role_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message tkd.idm.v1.Role
 */
export declare class Role extends Message<Role> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: string description = 3;
   */
  description: string;

  /**
   * @generated from field: bool delete_protected = 4;
   */
  deleteProtected: boolean;

  constructor(data?: PartialMessage<Role>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.Role";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Role;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Role;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Role;

  static equals(a: Role | PlainMessage<Role> | undefined, b: Role | PlainMessage<Role> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ListRolesRequest
 */
export declare class ListRolesRequest extends Message<ListRolesRequest> {
  constructor(data?: PartialMessage<ListRolesRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ListRolesRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRolesRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRolesRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRolesRequest;

  static equals(a: ListRolesRequest | PlainMessage<ListRolesRequest> | undefined, b: ListRolesRequest | PlainMessage<ListRolesRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ListRolesResponse
 */
export declare class ListRolesResponse extends Message<ListRolesResponse> {
  /**
   * @generated from field: repeated tkd.idm.v1.Role roles = 1;
   */
  roles: Role[];

  constructor(data?: PartialMessage<ListRolesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ListRolesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListRolesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListRolesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListRolesResponse;

  static equals(a: ListRolesResponse | PlainMessage<ListRolesResponse> | undefined, b: ListRolesResponse | PlainMessage<ListRolesResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.GetRoleRequest
 */
export declare class GetRoleRequest extends Message<GetRoleRequest> {
  /**
   * @generated from oneof tkd.idm.v1.GetRoleRequest.search
   */
  search: {
    /**
     * @generated from field: string name = 1;
     */
    value: string;
    case: "name";
  } | {
    /**
     * @generated from field: string id = 2;
     */
    value: string;
    case: "id";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<GetRoleRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.GetRoleRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRoleRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRoleRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRoleRequest;

  static equals(a: GetRoleRequest | PlainMessage<GetRoleRequest> | undefined, b: GetRoleRequest | PlainMessage<GetRoleRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.GetRoleResponse
 */
export declare class GetRoleResponse extends Message<GetRoleResponse> {
  /**
   * @generated from field: tkd.idm.v1.Role role = 1;
   */
  role?: Role;

  constructor(data?: PartialMessage<GetRoleResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.GetRoleResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRoleResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRoleResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRoleResponse;

  static equals(a: GetRoleResponse | PlainMessage<GetRoleResponse> | undefined, b: GetRoleResponse | PlainMessage<GetRoleResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.AssignRoleToUserRequest
 */
export declare class AssignRoleToUserRequest extends Message<AssignRoleToUserRequest> {
  /**
   * @generated from field: string role_id = 1;
   */
  roleId: string;

  /**
   * @generated from field: repeated string user_id = 2;
   */
  userId: string[];

  constructor(data?: PartialMessage<AssignRoleToUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.AssignRoleToUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssignRoleToUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssignRoleToUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssignRoleToUserRequest;

  static equals(a: AssignRoleToUserRequest | PlainMessage<AssignRoleToUserRequest> | undefined, b: AssignRoleToUserRequest | PlainMessage<AssignRoleToUserRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.AssignRoleToUserResponse
 */
export declare class AssignRoleToUserResponse extends Message<AssignRoleToUserResponse> {
  constructor(data?: PartialMessage<AssignRoleToUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.AssignRoleToUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssignRoleToUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssignRoleToUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssignRoleToUserResponse;

  static equals(a: AssignRoleToUserResponse | PlainMessage<AssignRoleToUserResponse> | undefined, b: AssignRoleToUserResponse | PlainMessage<AssignRoleToUserResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.UnassignRoleFromUserRequest
 */
export declare class UnassignRoleFromUserRequest extends Message<UnassignRoleFromUserRequest> {
  /**
   * @generated from field: string role_id = 1;
   */
  roleId: string;

  /**
   * @generated from field: repeated string user_id = 2;
   */
  userId: string[];

  constructor(data?: PartialMessage<UnassignRoleFromUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.UnassignRoleFromUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UnassignRoleFromUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UnassignRoleFromUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UnassignRoleFromUserRequest;

  static equals(a: UnassignRoleFromUserRequest | PlainMessage<UnassignRoleFromUserRequest> | undefined, b: UnassignRoleFromUserRequest | PlainMessage<UnassignRoleFromUserRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.UnassignRoleFromUserResponse
 */
export declare class UnassignRoleFromUserResponse extends Message<UnassignRoleFromUserResponse> {
  constructor(data?: PartialMessage<UnassignRoleFromUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.UnassignRoleFromUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UnassignRoleFromUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UnassignRoleFromUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UnassignRoleFromUserResponse;

  static equals(a: UnassignRoleFromUserResponse | PlainMessage<UnassignRoleFromUserResponse> | undefined, b: UnassignRoleFromUserResponse | PlainMessage<UnassignRoleFromUserResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.CreateRoleRequest
 */
export declare class CreateRoleRequest extends Message<CreateRoleRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string description = 2;
   */
  description: string;

  /**
   * @generated from field: bool delete_protection = 3;
   */
  deleteProtection: boolean;

  /**
   * @generated from field: string id = 4;
   */
  id: string;

  constructor(data?: PartialMessage<CreateRoleRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.CreateRoleRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRoleRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRoleRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRoleRequest;

  static equals(a: CreateRoleRequest | PlainMessage<CreateRoleRequest> | undefined, b: CreateRoleRequest | PlainMessage<CreateRoleRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.CreateRoleResponse
 */
export declare class CreateRoleResponse extends Message<CreateRoleResponse> {
  /**
   * @generated from field: tkd.idm.v1.Role role = 1;
   */
  role?: Role;

  constructor(data?: PartialMessage<CreateRoleResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.CreateRoleResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRoleResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRoleResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRoleResponse;

  static equals(a: CreateRoleResponse | PlainMessage<CreateRoleResponse> | undefined, b: CreateRoleResponse | PlainMessage<CreateRoleResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.UpdateRoleRequest
 */
export declare class UpdateRoleRequest extends Message<UpdateRoleRequest> {
  /**
   * @generated from field: string role_id = 1;
   */
  roleId: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: string description = 3;
   */
  description: string;

  /**
   * @generated from field: bool delete_protection = 4;
   */
  deleteProtection: boolean;

  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 5;
   */
  fieldMask?: FieldMask;

  constructor(data?: PartialMessage<UpdateRoleRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.UpdateRoleRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateRoleRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateRoleRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateRoleRequest;

  static equals(a: UpdateRoleRequest | PlainMessage<UpdateRoleRequest> | undefined, b: UpdateRoleRequest | PlainMessage<UpdateRoleRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.UpdateRoleResponse
 */
export declare class UpdateRoleResponse extends Message<UpdateRoleResponse> {
  /**
   * @generated from field: tkd.idm.v1.Role role = 1;
   */
  role?: Role;

  constructor(data?: PartialMessage<UpdateRoleResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.UpdateRoleResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateRoleResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateRoleResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateRoleResponse;

  static equals(a: UpdateRoleResponse | PlainMessage<UpdateRoleResponse> | undefined, b: UpdateRoleResponse | PlainMessage<UpdateRoleResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.DeleteRoleRequest
 */
export declare class DeleteRoleRequest extends Message<DeleteRoleRequest> {
  /**
   * @generated from field: string role_id = 1;
   */
  roleId: string;

  constructor(data?: PartialMessage<DeleteRoleRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.DeleteRoleRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRoleRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRoleRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRoleRequest;

  static equals(a: DeleteRoleRequest | PlainMessage<DeleteRoleRequest> | undefined, b: DeleteRoleRequest | PlainMessage<DeleteRoleRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.DeleteRoleResponse
 */
export declare class DeleteRoleResponse extends Message<DeleteRoleResponse> {
  constructor(data?: PartialMessage<DeleteRoleResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.DeleteRoleResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRoleResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRoleResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRoleResponse;

  static equals(a: DeleteRoleResponse | PlainMessage<DeleteRoleResponse> | undefined, b: DeleteRoleResponse | PlainMessage<DeleteRoleResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ResolveRolePermissionsRequest
 */
export declare class ResolveRolePermissionsRequest extends Message<ResolveRolePermissionsRequest> {
  /**
   * @generated from field: string role_id = 1;
   */
  roleId: string;

  constructor(data?: PartialMessage<ResolveRolePermissionsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ResolveRolePermissionsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResolveRolePermissionsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResolveRolePermissionsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResolveRolePermissionsRequest;

  static equals(a: ResolveRolePermissionsRequest | PlainMessage<ResolveRolePermissionsRequest> | undefined, b: ResolveRolePermissionsRequest | PlainMessage<ResolveRolePermissionsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ResolveRolePermissionsResponse
 */
export declare class ResolveRolePermissionsResponse extends Message<ResolveRolePermissionsResponse> {
  /**
   * @generated from field: repeated string permissions = 1;
   */
  permissions: string[];

  constructor(data?: PartialMessage<ResolveRolePermissionsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ResolveRolePermissionsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResolveRolePermissionsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResolveRolePermissionsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResolveRolePermissionsResponse;

  static equals(a: ResolveRolePermissionsResponse | PlainMessage<ResolveRolePermissionsResponse> | undefined, b: ResolveRolePermissionsResponse | PlainMessage<ResolveRolePermissionsResponse> | undefined): boolean;
}

