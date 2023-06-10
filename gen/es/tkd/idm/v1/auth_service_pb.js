// @generated by protoc-gen-es v1.2.1 with parameter "target=js+ts"
// @generated from file tkd/idm/v1/auth_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";
import { Profile, User } from "./user_pb.js";

/**
 * @generated from enum tkd.idm.v1.AuthType
 */
export const AuthType = proto3.makeEnum(
  "tkd.idm.v1.AuthType",
  [
    {no: 0, name: "AUTH_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "AUTH_TYPE_PASSWORD", localName: "PASSWORD"},
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
 * @generated from message tkd.idm.v1.LoginRequest
 */
export const LoginRequest = proto3.makeMessageType(
  "tkd.idm.v1.LoginRequest",
  () => [
    { no: 1, name: "auth_type", kind: "enum", T: proto3.getEnumType(AuthType) },
    { no: 2, name: "no_refresh_token", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "password", kind: "message", T: PasswordAuth, oneof: "auth" },
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
 * @generated from message tkd.idm.v1.LoginResponse
 */
export const LoginResponse = proto3.makeMessageType(
  "tkd.idm.v1.LoginResponse",
  () => [
    { no: 1, name: "access_token", kind: "message", T: AccessTokenResponse, oneof: "response" },
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
  [],
);

/**
 * @generated from message tkd.idm.v1.RefreshTokenResponse
 */
export const RefreshTokenResponse = proto3.makeMessageType(
  "tkd.idm.v1.RefreshTokenResponse",
  () => [
    { no: 1, name: "access_token", kind: "message", T: AccessTokenResponse },
  ],
);

/**
 * empty on purpose
 *
 * @generated from message tkd.idm.v1.IntrospectRequest
 */
export const IntrospectRequest = proto3.makeMessageType(
  "tkd.idm.v1.IntrospectRequest",
  [],
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

