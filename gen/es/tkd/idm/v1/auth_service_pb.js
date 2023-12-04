// @generated by protoc-gen-es v1.5.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/auth_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, FieldMask, proto3, Timestamp } from "@bufbuild/protobuf";
import { Profile, User } from "./user_pb.js";

/**
 * @generated from enum tkd.idm.v1.AuthType
 */
export const AuthType = proto3.makeEnum(
  "tkd.idm.v1.AuthType",
  [
    {no: 0, name: "AUTH_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "AUTH_TYPE_PASSWORD", localName: "PASSWORD"},
    {no: 2, name: "AUTH_TYPE_TOTP", localName: "TOTP"},
  ],
);

/**
 * @generated from enum tkd.idm.v1.RequiredMFAKind
 */
export const RequiredMFAKind = proto3.makeEnum(
  "tkd.idm.v1.RequiredMFAKind",
  [
    {no: 0, name: "REQUIRED_MFA_KIND_UNSPECIFIED"},
    {no: 1, name: "REQUIRED_MFA_KIND_TOTP"},
  ],
);

/**
 * @generated from message tkd.idm.v1.PasswordAuth
 */
export const PasswordAuth = proto3.makeMessageType(
  "tkd.idm.v1.PasswordAuth",
  () => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.TotpAuth
 */
export const TotpAuth = proto3.makeMessageType(
  "tkd.idm.v1.TotpAuth",
  () => [
    { no: 1, name: "code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "state", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.LoginRequest
 */
export const LoginRequest = proto3.makeMessageType(
  "tkd.idm.v1.LoginRequest",
  () => [
    { no: 1, name: "auth_type", kind: "enum", T: proto3.getEnumType(AuthType) },
    { no: 2, name: "no_refresh_token", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "password", kind: "message", T: PasswordAuth, oneof: "auth" },
    { no: 4, name: "totp", kind: "message", T: TotpAuth, oneof: "auth" },
    { no: 5, name: "requested_redirect", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "ttl", kind: "message", T: Duration },
  ],
);

/**
 * @generated from message tkd.idm.v1.AccessTokenResponse
 */
export const AccessTokenResponse = proto3.makeMessageType(
  "tkd.idm.v1.AccessTokenResponse",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expires_at", kind: "message", T: Timestamp },
    { no: 3, name: "user", kind: "message", T: User },
  ],
);

/**
 * @generated from message tkd.idm.v1.MFARequiredResponse
 */
export const MFARequiredResponse = proto3.makeMessageType(
  "tkd.idm.v1.MFARequiredResponse",
  () => [
    { no: 1, name: "kind", kind: "enum", T: proto3.getEnumType(RequiredMFAKind) },
    { no: 2, name: "state", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.LoginResponse
 */
export const LoginResponse = proto3.makeMessageType(
  "tkd.idm.v1.LoginResponse",
  () => [
    { no: 1, name: "access_token", kind: "message", T: AccessTokenResponse, oneof: "response" },
    { no: 2, name: "mfa_required", kind: "message", T: MFARequiredResponse, oneof: "response" },
    { no: 5, name: "redirect_to", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * empty on purpose
 *
 * @generated from message tkd.idm.v1.LogoutRequest
 */
export const LogoutRequest = proto3.makeMessageType(
  "tkd.idm.v1.LogoutRequest",
  [],
);

/**
 * empty on purpose
 *
 * @generated from message tkd.idm.v1.LogoutResponse
 */
export const LogoutResponse = proto3.makeMessageType(
  "tkd.idm.v1.LogoutResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.RefreshTokenRequest
 */
export const RefreshTokenRequest = proto3.makeMessageType(
  "tkd.idm.v1.RefreshTokenRequest",
  () => [
    { no: 1, name: "ttl", kind: "message", T: Duration },
    { no: 2, name: "requested_redirect", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.RefreshTokenResponse
 */
export const RefreshTokenResponse = proto3.makeMessageType(
  "tkd.idm.v1.RefreshTokenResponse",
  () => [
    { no: 1, name: "access_token", kind: "message", T: AccessTokenResponse },
    { no: 2, name: "redirect_to", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.IntrospectRequest
 */
export const IntrospectRequest = proto3.makeMessageType(
  "tkd.idm.v1.IntrospectRequest",
  () => [
    { no: 1, name: "read_mask", kind: "message", T: FieldMask },
    { no: 2, name: "exclude_fields", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.IntrospectResponse
 */
export const IntrospectResponse = proto3.makeMessageType(
  "tkd.idm.v1.IntrospectResponse",
  () => [
    { no: 1, name: "profile", kind: "message", T: Profile },
  ],
);

/**
 * @generated from message tkd.idm.v1.GenerateRegistrationTokenRequest
 */
export const GenerateRegistrationTokenRequest = proto3.makeMessageType(
  "tkd.idm.v1.GenerateRegistrationTokenRequest",
  () => [
    { no: 1, name: "ttl", kind: "message", T: Duration },
    { no: 2, name: "max_count", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "initial_roles", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.GenerateRegistrationTokenResponse
 */
export const GenerateRegistrationTokenResponse = proto3.makeMessageType(
  "tkd.idm.v1.GenerateRegistrationTokenResponse",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.ValidateRegistrationTokenRequest
 */
export const ValidateRegistrationTokenRequest = proto3.makeMessageType(
  "tkd.idm.v1.ValidateRegistrationTokenRequest",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.ValidateRegistrationTokenResponse
 */
export const ValidateRegistrationTokenResponse = proto3.makeMessageType(
  "tkd.idm.v1.ValidateRegistrationTokenResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.RegisterUserRequest
 */
export const RegisterUserRequest = proto3.makeMessageType(
  "tkd.idm.v1.RegisterUserRequest",
  () => [
    { no: 1, name: "registration_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.RegisterUserResponse
 */
export const RegisterUserResponse = proto3.makeMessageType(
  "tkd.idm.v1.RegisterUserResponse",
  () => [
    { no: 1, name: "access_token", kind: "message", T: AccessTokenResponse },
  ],
);

/**
 * @generated from message tkd.idm.v1.PasswordReset
 */
export const PasswordReset = proto3.makeMessageType(
  "tkd.idm.v1.PasswordReset",
  () => [
    { no: 1, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "new_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.RequestPasswordResetRequest
 */
export const RequestPasswordResetRequest = proto3.makeMessageType(
  "tkd.idm.v1.RequestPasswordResetRequest",
  () => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "kind" },
    { no: 2, name: "password_reset", kind: "message", T: PasswordReset, oneof: "kind" },
  ],
);

/**
 * @generated from message tkd.idm.v1.RequestPasswordResetResponse
 */
export const RequestPasswordResetResponse = proto3.makeMessageType(
  "tkd.idm.v1.RequestPasswordResetResponse",
  [],
);

