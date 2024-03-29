// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/user_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateUserRequest, CreateUserResponse, DeleteUserExtraKeyRequest, DeleteUserExtraKeyResponse, DeleteUserRequest, DeleteUserResponse, GetUserRequest, GetUserResponse, ImpersonateRequest, ImpersonateResponse, InviteUserRequest, InviteUserResponse, ListUsersRequest, ListUsersResponse, ResolveUserPermissionsRequest, ResolveUserPermissionsResponse, SendAccountCreationNoticeRequest, SendAccountCreationNoticeResponse, SetUserExtraKeyRequest, SetUserExtraKeyResponse, SetUserPasswordRequest, SetUserPasswordResponse, UpdateUserRequest, UpdateUserResponse } from "./user_service_pb.js";
import { MethodIdempotency, MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service tkd.idm.v1.UserService
 */
export declare const UserService: {
  readonly typeName: "tkd.idm.v1.UserService",
  readonly methods: {
    /**
     * @generated from rpc tkd.idm.v1.UserService.Impersonate
     */
    readonly impersonate: {
      readonly name: "Impersonate",
      readonly I: typeof ImpersonateRequest,
      readonly O: typeof ImpersonateResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.GetUser
     */
    readonly getUser: {
      readonly name: "GetUser",
      readonly I: typeof GetUserRequest,
      readonly O: typeof GetUserResponse,
      readonly kind: MethodKind.Unary,
      readonly idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.InviteUser
     */
    readonly inviteUser: {
      readonly name: "InviteUser",
      readonly I: typeof InviteUserRequest,
      readonly O: typeof InviteUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.ListUsers
     */
    readonly listUsers: {
      readonly name: "ListUsers",
      readonly I: typeof ListUsersRequest,
      readonly O: typeof ListUsersResponse,
      readonly kind: MethodKind.Unary,
      readonly idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.CreateUser
     */
    readonly createUser: {
      readonly name: "CreateUser",
      readonly I: typeof CreateUserRequest,
      readonly O: typeof CreateUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.UpdateUser
     */
    readonly updateUser: {
      readonly name: "UpdateUser",
      readonly I: typeof UpdateUserRequest,
      readonly O: typeof UpdateUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.DeleteUser
     */
    readonly deleteUser: {
      readonly name: "DeleteUser",
      readonly I: typeof DeleteUserRequest,
      readonly O: typeof DeleteUserResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.SetUserExtraKey
     */
    readonly setUserExtraKey: {
      readonly name: "SetUserExtraKey",
      readonly I: typeof SetUserExtraKeyRequest,
      readonly O: typeof SetUserExtraKeyResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.DeleteUserExtraKey
     */
    readonly deleteUserExtraKey: {
      readonly name: "DeleteUserExtraKey",
      readonly I: typeof DeleteUserExtraKeyRequest,
      readonly O: typeof DeleteUserExtraKeyResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.SendAccountCreationNotice
     */
    readonly sendAccountCreationNotice: {
      readonly name: "SendAccountCreationNotice",
      readonly I: typeof SendAccountCreationNoticeRequest,
      readonly O: typeof SendAccountCreationNoticeResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.SetUserPassword
     */
    readonly setUserPassword: {
      readonly name: "SetUserPassword",
      readonly I: typeof SetUserPasswordRequest,
      readonly O: typeof SetUserPasswordResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.UserService.ResolveUserPermissions
     */
    readonly resolveUserPermissions: {
      readonly name: "ResolveUserPermissions",
      readonly I: typeof ResolveUserPermissionsRequest,
      readonly O: typeof ResolveUserPermissionsResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

