// @generated by protoc-gen-es v1.2.1 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/user_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Profile } from "./user_pb.js";

/**
 * @generated from message tkd.idm.v1.GetUserRequest
 */
export declare class GetUserRequest extends Message<GetUserRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<GetUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.GetUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserRequest;

  static equals(a: GetUserRequest | PlainMessage<GetUserRequest> | undefined, b: GetUserRequest | PlainMessage<GetUserRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.GetUserResponse
 */
export declare class GetUserResponse extends Message<GetUserResponse> {
  /**
   * @generated from field: tkd.idm.v1.Profile profile = 1;
   */
  profile?: Profile;

  constructor(data?: PartialMessage<GetUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.GetUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserResponse;

  static equals(a: GetUserResponse | PlainMessage<GetUserResponse> | undefined, b: GetUserResponse | PlainMessage<GetUserResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ListUsersRequest
 */
export declare class ListUsersRequest extends Message<ListUsersRequest> {
  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 1;
   */
  fieldMask?: FieldMask;

  constructor(data?: PartialMessage<ListUsersRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ListUsersRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUsersRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUsersRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUsersRequest;

  static equals(a: ListUsersRequest | PlainMessage<ListUsersRequest> | undefined, b: ListUsersRequest | PlainMessage<ListUsersRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ListUsersResponse
 */
export declare class ListUsersResponse extends Message<ListUsersResponse> {
  /**
   * @generated from field: repeated tkd.idm.v1.Profile users = 1;
   */
  users: Profile[];

  constructor(data?: PartialMessage<ListUsersResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ListUsersResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUsersResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUsersResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUsersResponse;

  static equals(a: ListUsersResponse | PlainMessage<ListUsersResponse> | undefined, b: ListUsersResponse | PlainMessage<ListUsersResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.CreateUserRequest
 */
export declare class CreateUserRequest extends Message<CreateUserRequest> {
  /**
   * @generated from field: tkd.idm.v1.Profile profile = 1;
   */
  profile?: Profile;

  /**
   * @generated from field: string password = 2;
   */
  password: string;

  /**
   * @generated from field: bool password_is_bcrypt = 3;
   */
  passwordIsBcrypt: boolean;

  constructor(data?: PartialMessage<CreateUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.CreateUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserRequest;

  static equals(a: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined, b: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.CreateUserResponse
 */
export declare class CreateUserResponse extends Message<CreateUserResponse> {
  /**
   * @generated from field: tkd.idm.v1.Profile profile = 1;
   */
  profile?: Profile;

  /**
   * @generated from field: string generated_password = 2;
   */
  generatedPassword: string;

  constructor(data?: PartialMessage<CreateUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.CreateUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserResponse;

  static equals(a: CreateUserResponse | PlainMessage<CreateUserResponse> | undefined, b: CreateUserResponse | PlainMessage<CreateUserResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.UpdateUserRequest
 */
export declare class UpdateUserRequest extends Message<UpdateUserRequest> {
  constructor(data?: PartialMessage<UpdateUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.UpdateUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserRequest;

  static equals(a: UpdateUserRequest | PlainMessage<UpdateUserRequest> | undefined, b: UpdateUserRequest | PlainMessage<UpdateUserRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.UpdateUserResponse
 */
export declare class UpdateUserResponse extends Message<UpdateUserResponse> {
  constructor(data?: PartialMessage<UpdateUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.UpdateUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserResponse;

  static equals(a: UpdateUserResponse | PlainMessage<UpdateUserResponse> | undefined, b: UpdateUserResponse | PlainMessage<UpdateUserResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.DeleteUserRequest
 */
export declare class DeleteUserRequest extends Message<DeleteUserRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DeleteUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.DeleteUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteUserRequest;

  static equals(a: DeleteUserRequest | PlainMessage<DeleteUserRequest> | undefined, b: DeleteUserRequest | PlainMessage<DeleteUserRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.DeleteUserResponse
 */
export declare class DeleteUserResponse extends Message<DeleteUserResponse> {
  constructor(data?: PartialMessage<DeleteUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.DeleteUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteUserResponse;

  static equals(a: DeleteUserResponse | PlainMessage<DeleteUserResponse> | undefined, b: DeleteUserResponse | PlainMessage<DeleteUserResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.UserInvite
 */
export declare class UserInvite extends Message<UserInvite> {
  /**
   * @generated from field: string email = 1;
   */
  email: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  constructor(data?: PartialMessage<UserInvite>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.UserInvite";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserInvite;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserInvite;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserInvite;

  static equals(a: UserInvite | PlainMessage<UserInvite> | undefined, b: UserInvite | PlainMessage<UserInvite> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.InviteUserRequest
 */
export declare class InviteUserRequest extends Message<InviteUserRequest> {
  /**
   * @generated from field: repeated tkd.idm.v1.UserInvite invite = 1;
   */
  invite: UserInvite[];

  /**
   * @generated from field: repeated string initial_roles = 2;
   */
  initialRoles: string[];

  constructor(data?: PartialMessage<InviteUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.InviteUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InviteUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InviteUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InviteUserRequest;

  static equals(a: InviteUserRequest | PlainMessage<InviteUserRequest> | undefined, b: InviteUserRequest | PlainMessage<InviteUserRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.InviteUserResponse
 */
export declare class InviteUserResponse extends Message<InviteUserResponse> {
  constructor(data?: PartialMessage<InviteUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.InviteUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InviteUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InviteUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InviteUserResponse;

  static equals(a: InviteUserResponse | PlainMessage<InviteUserResponse> | undefined, b: InviteUserResponse | PlainMessage<InviteUserResponse> | undefined): boolean;
}

