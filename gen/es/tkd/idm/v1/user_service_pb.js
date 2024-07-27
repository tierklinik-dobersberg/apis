// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/user_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FieldMask, proto3, Struct, Value } from "@bufbuild/protobuf";
import { Profile } from "./user_pb.js";

/**
 * @generated from message tkd.idm.v1.ImpersonateRequest
 */
export const ImpersonateRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ImpersonateRequest",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.ImpersonateResponse
 */
export const ImpersonateResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ImpersonateResponse",
  () => [
    { no: 2, name: "access_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.GetUserRequest
 */
export const GetUserRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.GetUserRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "search" },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "search" },
    { no: 3, name: "mail", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "search" },
    { no: 4, name: "field_mask", kind: "message", T: FieldMask },
    { no: 5, name: "exclude_fields", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.GetUserResponse
 */
export const GetUserResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.GetUserResponse",
  () => [
    { no: 1, name: "profile", kind: "message", T: Profile },
  ],
);

/**
 * @generated from message tkd.idm.v1.ListUsersRequest
 */
export const ListUsersRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ListUsersRequest",
  () => [
    { no: 1, name: "field_mask", kind: "message", T: FieldMask },
    { no: 2, name: "exclude_fields", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "filter_by_roles", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.ListUsersResponse
 */
export const ListUsersResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ListUsersResponse",
  () => [
    { no: 1, name: "users", kind: "message", T: Profile, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.CreateUserRequest
 */
export const CreateUserRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.CreateUserRequest",
  () => [
    { no: 1, name: "profile", kind: "message", T: Profile },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "password_is_bcrypt", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.CreateUserResponse
 */
export const CreateUserResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.CreateUserResponse",
  () => [
    { no: 1, name: "profile", kind: "message", T: Profile },
    { no: 2, name: "generated_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.UpdateUserRequest
 */
export const UpdateUserRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.UpdateUserRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "first_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "last_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "extra", kind: "message", T: Struct },
    { no: 7, name: "avatar", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "birthday", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "field_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.idm.v1.UpdateUserResponse
 */
export const UpdateUserResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.UpdateUserResponse",
  () => [
    { no: 1, name: "profile", kind: "message", T: Profile },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeleteUserRequest
 */
export const DeleteUserRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeleteUserRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeleteUserResponse
 */
export const DeleteUserResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeleteUserResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.UserInvite
 */
export const UserInvite = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.UserInvite",
  () => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.InviteUserRequest
 */
export const InviteUserRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.InviteUserRequest",
  () => [
    { no: 1, name: "invite", kind: "message", T: UserInvite, repeated: true },
    { no: 2, name: "initial_roles", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.InviteUserResponse
 */
export const InviteUserResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.InviteUserResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.SetUserExtraKeyRequest
 */
export const SetUserExtraKeyRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.SetUserExtraKeyRequest",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "value", kind: "message", T: Value },
  ],
);

/**
 * @generated from message tkd.idm.v1.SetUserExtraKeyResponse
 */
export const SetUserExtraKeyResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.SetUserExtraKeyResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.DeleteUserExtraKeyRequest
 */
export const DeleteUserExtraKeyRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeleteUserExtraKeyRequest",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "path", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeleteUserExtraKeyResponse
 */
export const DeleteUserExtraKeyResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeleteUserExtraKeyResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.SendAccountCreationNoticeRequest
 */
export const SendAccountCreationNoticeRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.SendAccountCreationNoticeRequest",
  () => [
    { no: 1, name: "user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.SendAccountCreationNoticeResponse
 */
export const SendAccountCreationNoticeResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.SendAccountCreationNoticeResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.SetUserPasswordRequest
 */
export const SetUserPasswordRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.SetUserPasswordRequest",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.SetUserPasswordResponse
 */
export const SetUserPasswordResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.SetUserPasswordResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.ResolveUserPermissionsRequest
 */
export const ResolveUserPermissionsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ResolveUserPermissionsRequest",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.ResolveUserPermissionsResponse
 */
export const ResolveUserPermissionsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ResolveUserPermissionsResponse",
  () => [
    { no: 1, name: "permissions", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

