// @generated by protoc-gen-es v1.8.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/user_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Struct, Value } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Profile } from "./user_pb.js";

/**
 * @generated from message tkd.idm.v1.ImpersonateRequest
 */
export declare class ImpersonateRequest extends Message<ImpersonateRequest> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  constructor(data?: PartialMessage<ImpersonateRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ImpersonateRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ImpersonateRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ImpersonateRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ImpersonateRequest;

  static equals(a: ImpersonateRequest | PlainMessage<ImpersonateRequest> | undefined, b: ImpersonateRequest | PlainMessage<ImpersonateRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ImpersonateResponse
 */
export declare class ImpersonateResponse extends Message<ImpersonateResponse> {
  /**
   * @generated from field: string access_token = 2;
   */
  accessToken: string;

  constructor(data?: PartialMessage<ImpersonateResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ImpersonateResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ImpersonateResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ImpersonateResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ImpersonateResponse;

  static equals(a: ImpersonateResponse | PlainMessage<ImpersonateResponse> | undefined, b: ImpersonateResponse | PlainMessage<ImpersonateResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.GetUserRequest
 */
export declare class GetUserRequest extends Message<GetUserRequest> {
  /**
   * @generated from oneof tkd.idm.v1.GetUserRequest.search
   */
  search: {
    /**
     * @generated from field: string id = 1;
     */
    value: string;
    case: "id";
  } | {
    /**
     * @generated from field: string name = 2;
     */
    value: string;
    case: "name";
  } | {
    /**
     * @generated from field: string mail = 3;
     */
    value: string;
    case: "mail";
  } | { case: undefined; value?: undefined };

  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 4;
   */
  fieldMask?: FieldMask;

  /**
   * @generated from field: bool exclude_fields = 5;
   */
  excludeFields: boolean;

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

  /**
   * @generated from field: bool exclude_fields = 2;
   */
  excludeFields: boolean;

  /**
   * @generated from field: repeated string filter_by_roles = 3;
   */
  filterByRoles: string[];

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
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string username = 2;
   */
  username: string;

  /**
   * @generated from field: string display_name = 3;
   */
  displayName: string;

  /**
   * @generated from field: string first_name = 4;
   */
  firstName: string;

  /**
   * @generated from field: string last_name = 5;
   */
  lastName: string;

  /**
   * @generated from field: google.protobuf.Struct extra = 6;
   */
  extra?: Struct;

  /**
   * @generated from field: string avatar = 7;
   */
  avatar: string;

  /**
   * @generated from field: string birthday = 8;
   */
  birthday: string;

  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 9;
   */
  fieldMask?: FieldMask;

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
  /**
   * @generated from field: tkd.idm.v1.Profile profile = 1;
   */
  profile?: Profile;

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

/**
 * @generated from message tkd.idm.v1.SetUserExtraKeyRequest
 */
export declare class SetUserExtraKeyRequest extends Message<SetUserExtraKeyRequest> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * @generated from field: string path = 2;
   */
  path: string;

  /**
   * @generated from field: google.protobuf.Value value = 3;
   */
  value?: Value;

  constructor(data?: PartialMessage<SetUserExtraKeyRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.SetUserExtraKeyRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetUserExtraKeyRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetUserExtraKeyRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetUserExtraKeyRequest;

  static equals(a: SetUserExtraKeyRequest | PlainMessage<SetUserExtraKeyRequest> | undefined, b: SetUserExtraKeyRequest | PlainMessage<SetUserExtraKeyRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.SetUserExtraKeyResponse
 */
export declare class SetUserExtraKeyResponse extends Message<SetUserExtraKeyResponse> {
  constructor(data?: PartialMessage<SetUserExtraKeyResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.SetUserExtraKeyResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetUserExtraKeyResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetUserExtraKeyResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetUserExtraKeyResponse;

  static equals(a: SetUserExtraKeyResponse | PlainMessage<SetUserExtraKeyResponse> | undefined, b: SetUserExtraKeyResponse | PlainMessage<SetUserExtraKeyResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.DeleteUserExtraKeyRequest
 */
export declare class DeleteUserExtraKeyRequest extends Message<DeleteUserExtraKeyRequest> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * @generated from field: string path = 2;
   */
  path: string;

  constructor(data?: PartialMessage<DeleteUserExtraKeyRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.DeleteUserExtraKeyRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteUserExtraKeyRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteUserExtraKeyRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteUserExtraKeyRequest;

  static equals(a: DeleteUserExtraKeyRequest | PlainMessage<DeleteUserExtraKeyRequest> | undefined, b: DeleteUserExtraKeyRequest | PlainMessage<DeleteUserExtraKeyRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.DeleteUserExtraKeyResponse
 */
export declare class DeleteUserExtraKeyResponse extends Message<DeleteUserExtraKeyResponse> {
  constructor(data?: PartialMessage<DeleteUserExtraKeyResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.DeleteUserExtraKeyResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteUserExtraKeyResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteUserExtraKeyResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteUserExtraKeyResponse;

  static equals(a: DeleteUserExtraKeyResponse | PlainMessage<DeleteUserExtraKeyResponse> | undefined, b: DeleteUserExtraKeyResponse | PlainMessage<DeleteUserExtraKeyResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.SendAccountCreationNoticeRequest
 */
export declare class SendAccountCreationNoticeRequest extends Message<SendAccountCreationNoticeRequest> {
  /**
   * @generated from field: repeated string user_ids = 1;
   */
  userIds: string[];

  constructor(data?: PartialMessage<SendAccountCreationNoticeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.SendAccountCreationNoticeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendAccountCreationNoticeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendAccountCreationNoticeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendAccountCreationNoticeRequest;

  static equals(a: SendAccountCreationNoticeRequest | PlainMessage<SendAccountCreationNoticeRequest> | undefined, b: SendAccountCreationNoticeRequest | PlainMessage<SendAccountCreationNoticeRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.SendAccountCreationNoticeResponse
 */
export declare class SendAccountCreationNoticeResponse extends Message<SendAccountCreationNoticeResponse> {
  constructor(data?: PartialMessage<SendAccountCreationNoticeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.SendAccountCreationNoticeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendAccountCreationNoticeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendAccountCreationNoticeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendAccountCreationNoticeResponse;

  static equals(a: SendAccountCreationNoticeResponse | PlainMessage<SendAccountCreationNoticeResponse> | undefined, b: SendAccountCreationNoticeResponse | PlainMessage<SendAccountCreationNoticeResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.SetUserPasswordRequest
 */
export declare class SetUserPasswordRequest extends Message<SetUserPasswordRequest> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * @generated from field: string password = 2;
   */
  password: string;

  constructor(data?: PartialMessage<SetUserPasswordRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.SetUserPasswordRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetUserPasswordRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetUserPasswordRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetUserPasswordRequest;

  static equals(a: SetUserPasswordRequest | PlainMessage<SetUserPasswordRequest> | undefined, b: SetUserPasswordRequest | PlainMessage<SetUserPasswordRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.SetUserPasswordResponse
 */
export declare class SetUserPasswordResponse extends Message<SetUserPasswordResponse> {
  constructor(data?: PartialMessage<SetUserPasswordResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.SetUserPasswordResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetUserPasswordResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetUserPasswordResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetUserPasswordResponse;

  static equals(a: SetUserPasswordResponse | PlainMessage<SetUserPasswordResponse> | undefined, b: SetUserPasswordResponse | PlainMessage<SetUserPasswordResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ResolveUserPermissionsRequest
 */
export declare class ResolveUserPermissionsRequest extends Message<ResolveUserPermissionsRequest> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  constructor(data?: PartialMessage<ResolveUserPermissionsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ResolveUserPermissionsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResolveUserPermissionsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResolveUserPermissionsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResolveUserPermissionsRequest;

  static equals(a: ResolveUserPermissionsRequest | PlainMessage<ResolveUserPermissionsRequest> | undefined, b: ResolveUserPermissionsRequest | PlainMessage<ResolveUserPermissionsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ResolveUserPermissionsResponse
 */
export declare class ResolveUserPermissionsResponse extends Message<ResolveUserPermissionsResponse> {
  /**
   * @generated from field: repeated string permissions = 1;
   */
  permissions: string[];

  constructor(data?: PartialMessage<ResolveUserPermissionsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ResolveUserPermissionsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ResolveUserPermissionsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ResolveUserPermissionsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ResolveUserPermissionsResponse;

  static equals(a: ResolveUserPermissionsResponse | PlainMessage<ResolveUserPermissionsResponse> | undefined, b: ResolveUserPermissionsResponse | PlainMessage<ResolveUserPermissionsResponse> | undefined): boolean;
}

