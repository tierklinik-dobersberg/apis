// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/auth_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, Duration, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Profile, User } from "./user_pb.js";

/**
 * @generated from enum tkd.idm.v1.AuthType
 */
export declare enum AuthType {
  /**
   * @generated from enum value: AUTH_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: AUTH_TYPE_PASSWORD = 1;
   */
  PASSWORD = 1,

  /**
   * @generated from enum value: AUTH_TYPE_TOTP = 2;
   */
  TOTP = 2,
}

/**
 * @generated from enum tkd.idm.v1.RequiredMFAKind
 */
export declare enum RequiredMFAKind {
  /**
   * @generated from enum value: REQUIRED_MFA_KIND_UNSPECIFIED = 0;
   */
  REQUIRED_MFA_KIND_UNSPECIFIED = 0,

  /**
   * @generated from enum value: REQUIRED_MFA_KIND_TOTP = 1;
   */
  REQUIRED_MFA_KIND_TOTP = 1,
}

/**
 * @generated from message tkd.idm.v1.PasswordAuth
 */
export declare class PasswordAuth extends Message<PasswordAuth> {
  /**
   * @generated from field: string username = 1;
   */
  username: string;

  /**
   * @generated from field: string password = 2;
   */
  password: string;

  constructor(data?: PartialMessage<PasswordAuth>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.PasswordAuth";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PasswordAuth;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PasswordAuth;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PasswordAuth;

  static equals(a: PasswordAuth | PlainMessage<PasswordAuth> | undefined, b: PasswordAuth | PlainMessage<PasswordAuth> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.TotpAuth
 */
export declare class TotpAuth extends Message<TotpAuth> {
  /**
   * @generated from field: string code = 1;
   */
  code: string;

  /**
   * @generated from field: string state = 2;
   */
  state: string;

  constructor(data?: PartialMessage<TotpAuth>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.TotpAuth";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TotpAuth;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TotpAuth;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TotpAuth;

  static equals(a: TotpAuth | PlainMessage<TotpAuth> | undefined, b: TotpAuth | PlainMessage<TotpAuth> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.LoginRequest
 */
export declare class LoginRequest extends Message<LoginRequest> {
  /**
   * Deprecated: this field is ignored since the oneof below
   * acutally provides the same.
   *
   * @generated from field: tkd.idm.v1.AuthType auth_type = 1;
   */
  authType: AuthType;

  /**
   * NoRefreshToken may be set to true if the user only want's to
   * receive an access token.
   * Note that without a refresh token, it's not possible to use the
   * RefreshToken RPC, instead, the whole authentication flow of the
   * Login RPC must be performed.
   *
   * @generated from field: bool no_refresh_token = 2;
   */
  noRefreshToken: boolean;

  /**
   * @generated from oneof tkd.idm.v1.LoginRequest.auth
   */
  auth: {
    /**
     * @generated from field: tkd.idm.v1.PasswordAuth password = 3;
     */
    value: PasswordAuth;
    case: "password";
  } | {
    /**
     * @generated from field: tkd.idm.v1.TotpAuth totp = 4;
     */
    value: TotpAuth;
    case: "totp";
  } | { case: undefined; value?: undefined };

  /**
   * The URL that should be redirected to
   * once the authentication is successfull.
   *
   * @generated from field: string requested_redirect = 5;
   */
  requestedRedirect: string;

  /**
   * @generated from field: google.protobuf.Duration ttl = 6;
   */
  ttl?: Duration;

  constructor(data?: PartialMessage<LoginRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.LoginRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginRequest;

  static equals(a: LoginRequest | PlainMessage<LoginRequest> | undefined, b: LoginRequest | PlainMessage<LoginRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.AccessTokenResponse
 */
export declare class AccessTokenResponse extends Message<AccessTokenResponse> {
  /**
   * @generated from field: string token = 1;
   */
  token: string;

  /**
   * @generated from field: google.protobuf.Timestamp expires_at = 2;
   */
  expiresAt?: Timestamp;

  /**
   * @generated from field: tkd.idm.v1.User user = 3;
   */
  user?: User;

  constructor(data?: PartialMessage<AccessTokenResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.AccessTokenResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AccessTokenResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AccessTokenResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AccessTokenResponse;

  static equals(a: AccessTokenResponse | PlainMessage<AccessTokenResponse> | undefined, b: AccessTokenResponse | PlainMessage<AccessTokenResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.MFARequiredResponse
 */
export declare class MFARequiredResponse extends Message<MFARequiredResponse> {
  /**
   * @generated from field: tkd.idm.v1.RequiredMFAKind kind = 1;
   */
  kind: RequiredMFAKind;

  /**
   * @generated from field: string state = 2;
   */
  state: string;

  constructor(data?: PartialMessage<MFARequiredResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.MFARequiredResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MFARequiredResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MFARequiredResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MFARequiredResponse;

  static equals(a: MFARequiredResponse | PlainMessage<MFARequiredResponse> | undefined, b: MFARequiredResponse | PlainMessage<MFARequiredResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.LoginResponse
 */
export declare class LoginResponse extends Message<LoginResponse> {
  /**
   * @generated from oneof tkd.idm.v1.LoginResponse.response
   */
  response: {
    /**
     * @generated from field: tkd.idm.v1.AccessTokenResponse access_token = 1;
     */
    value: AccessTokenResponse;
    case: "accessToken";
  } | {
    /**
     * @generated from field: tkd.idm.v1.MFARequiredResponse mfa_required = 2;
     */
    value: MFARequiredResponse;
    case: "mfaRequired";
  } | { case: undefined; value?: undefined };

  /**
   * If set, the client is expected to redirect the
   * user to the returned URL. The url is sanitized and
   * validated to not allow open-redirects.
   *
   * @generated from field: string redirect_to = 5;
   */
  redirectTo: string;

  constructor(data?: PartialMessage<LoginResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.LoginResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LoginResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LoginResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LoginResponse;

  static equals(a: LoginResponse | PlainMessage<LoginResponse> | undefined, b: LoginResponse | PlainMessage<LoginResponse> | undefined): boolean;
}

/**
 * empty on purpose
 *
 * @generated from message tkd.idm.v1.LogoutRequest
 */
export declare class LogoutRequest extends Message<LogoutRequest> {
  constructor(data?: PartialMessage<LogoutRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.LogoutRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogoutRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogoutRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogoutRequest;

  static equals(a: LogoutRequest | PlainMessage<LogoutRequest> | undefined, b: LogoutRequest | PlainMessage<LogoutRequest> | undefined): boolean;
}

/**
 * empty on purpose
 *
 * @generated from message tkd.idm.v1.LogoutResponse
 */
export declare class LogoutResponse extends Message<LogoutResponse> {
  constructor(data?: PartialMessage<LogoutResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.LogoutResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogoutResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogoutResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogoutResponse;

  static equals(a: LogoutResponse | PlainMessage<LogoutResponse> | undefined, b: LogoutResponse | PlainMessage<LogoutResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.RefreshTokenRequest
 */
export declare class RefreshTokenRequest extends Message<RefreshTokenRequest> {
  /**
   * @generated from field: google.protobuf.Duration ttl = 1;
   */
  ttl?: Duration;

  /**
   * @generated from field: string requested_redirect = 2;
   */
  requestedRedirect: string;

  constructor(data?: PartialMessage<RefreshTokenRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.RefreshTokenRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RefreshTokenRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RefreshTokenRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RefreshTokenRequest;

  static equals(a: RefreshTokenRequest | PlainMessage<RefreshTokenRequest> | undefined, b: RefreshTokenRequest | PlainMessage<RefreshTokenRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.RefreshTokenResponse
 */
export declare class RefreshTokenResponse extends Message<RefreshTokenResponse> {
  /**
   * @generated from field: tkd.idm.v1.AccessTokenResponse access_token = 1;
   */
  accessToken?: AccessTokenResponse;

  /**
   * @generated from field: string redirect_to = 2;
   */
  redirectTo: string;

  constructor(data?: PartialMessage<RefreshTokenResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.RefreshTokenResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RefreshTokenResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RefreshTokenResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RefreshTokenResponse;

  static equals(a: RefreshTokenResponse | PlainMessage<RefreshTokenResponse> | undefined, b: RefreshTokenResponse | PlainMessage<RefreshTokenResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.IntrospectRequest
 */
export declare class IntrospectRequest extends Message<IntrospectRequest> {
  /**
   * @generated from field: google.protobuf.FieldMask read_mask = 1;
   */
  readMask?: FieldMask;

  /**
   * @generated from field: bool exclude_fields = 2;
   */
  excludeFields: boolean;

  constructor(data?: PartialMessage<IntrospectRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.IntrospectRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IntrospectRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IntrospectRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IntrospectRequest;

  static equals(a: IntrospectRequest | PlainMessage<IntrospectRequest> | undefined, b: IntrospectRequest | PlainMessage<IntrospectRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.IntrospectResponse
 */
export declare class IntrospectResponse extends Message<IntrospectResponse> {
  /**
   * @generated from field: tkd.idm.v1.Profile profile = 1;
   */
  profile?: Profile;

  /**
   * ValidTime holds the time until the provided access
   * token is valid.
   *
   * @generated from field: google.protobuf.Timestamp valid_time = 2;
   */
  validTime?: Timestamp;

  constructor(data?: PartialMessage<IntrospectResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.IntrospectResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): IntrospectResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): IntrospectResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): IntrospectResponse;

  static equals(a: IntrospectResponse | PlainMessage<IntrospectResponse> | undefined, b: IntrospectResponse | PlainMessage<IntrospectResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.GenerateRegistrationTokenRequest
 */
export declare class GenerateRegistrationTokenRequest extends Message<GenerateRegistrationTokenRequest> {
  /**
   * @generated from field: google.protobuf.Duration ttl = 1;
   */
  ttl?: Duration;

  /**
   * @generated from field: uint64 max_count = 2;
   */
  maxCount: bigint;

  /**
   * @generated from field: repeated string initial_roles = 3;
   */
  initialRoles: string[];

  constructor(data?: PartialMessage<GenerateRegistrationTokenRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.GenerateRegistrationTokenRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GenerateRegistrationTokenRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GenerateRegistrationTokenRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GenerateRegistrationTokenRequest;

  static equals(a: GenerateRegistrationTokenRequest | PlainMessage<GenerateRegistrationTokenRequest> | undefined, b: GenerateRegistrationTokenRequest | PlainMessage<GenerateRegistrationTokenRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.GenerateRegistrationTokenResponse
 */
export declare class GenerateRegistrationTokenResponse extends Message<GenerateRegistrationTokenResponse> {
  /**
   * @generated from field: string token = 1;
   */
  token: string;

  constructor(data?: PartialMessage<GenerateRegistrationTokenResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.GenerateRegistrationTokenResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GenerateRegistrationTokenResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GenerateRegistrationTokenResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GenerateRegistrationTokenResponse;

  static equals(a: GenerateRegistrationTokenResponse | PlainMessage<GenerateRegistrationTokenResponse> | undefined, b: GenerateRegistrationTokenResponse | PlainMessage<GenerateRegistrationTokenResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ValidateRegistrationTokenRequest
 */
export declare class ValidateRegistrationTokenRequest extends Message<ValidateRegistrationTokenRequest> {
  /**
   * @generated from field: string token = 1;
   */
  token: string;

  constructor(data?: PartialMessage<ValidateRegistrationTokenRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ValidateRegistrationTokenRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ValidateRegistrationTokenRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ValidateRegistrationTokenRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ValidateRegistrationTokenRequest;

  static equals(a: ValidateRegistrationTokenRequest | PlainMessage<ValidateRegistrationTokenRequest> | undefined, b: ValidateRegistrationTokenRequest | PlainMessage<ValidateRegistrationTokenRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.ValidateRegistrationTokenResponse
 */
export declare class ValidateRegistrationTokenResponse extends Message<ValidateRegistrationTokenResponse> {
  constructor(data?: PartialMessage<ValidateRegistrationTokenResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.ValidateRegistrationTokenResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ValidateRegistrationTokenResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ValidateRegistrationTokenResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ValidateRegistrationTokenResponse;

  static equals(a: ValidateRegistrationTokenResponse | PlainMessage<ValidateRegistrationTokenResponse> | undefined, b: ValidateRegistrationTokenResponse | PlainMessage<ValidateRegistrationTokenResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.RegisterUserRequest
 */
export declare class RegisterUserRequest extends Message<RegisterUserRequest> {
  /**
   * @generated from field: string registration_token = 1;
   */
  registrationToken: string;

  /**
   * @generated from field: string username = 2;
   */
  username: string;

  /**
   * @generated from field: string password = 3;
   */
  password: string;

  /**
   * @generated from field: string email = 4;
   */
  email: string;

  constructor(data?: PartialMessage<RegisterUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.RegisterUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterUserRequest;

  static equals(a: RegisterUserRequest | PlainMessage<RegisterUserRequest> | undefined, b: RegisterUserRequest | PlainMessage<RegisterUserRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.RegisterUserResponse
 */
export declare class RegisterUserResponse extends Message<RegisterUserResponse> {
  /**
   * @generated from field: tkd.idm.v1.AccessTokenResponse access_token = 1;
   */
  accessToken?: AccessTokenResponse;

  constructor(data?: PartialMessage<RegisterUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.RegisterUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RegisterUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RegisterUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RegisterUserResponse;

  static equals(a: RegisterUserResponse | PlainMessage<RegisterUserResponse> | undefined, b: RegisterUserResponse | PlainMessage<RegisterUserResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.PasswordReset
 */
export declare class PasswordReset extends Message<PasswordReset> {
  /**
   * @generated from field: string token = 1;
   */
  token: string;

  /**
   * @generated from field: string new_password = 2;
   */
  newPassword: string;

  constructor(data?: PartialMessage<PasswordReset>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.PasswordReset";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PasswordReset;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PasswordReset;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PasswordReset;

  static equals(a: PasswordReset | PlainMessage<PasswordReset> | undefined, b: PasswordReset | PlainMessage<PasswordReset> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.RequestPasswordResetRequest
 */
export declare class RequestPasswordResetRequest extends Message<RequestPasswordResetRequest> {
  /**
   * @generated from oneof tkd.idm.v1.RequestPasswordResetRequest.kind
   */
  kind: {
    /**
     * @generated from field: string email = 1;
     */
    value: string;
    case: "email";
  } | {
    /**
     * @generated from field: tkd.idm.v1.PasswordReset password_reset = 2;
     */
    value: PasswordReset;
    case: "passwordReset";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<RequestPasswordResetRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.RequestPasswordResetRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RequestPasswordResetRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RequestPasswordResetRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RequestPasswordResetRequest;

  static equals(a: RequestPasswordResetRequest | PlainMessage<RequestPasswordResetRequest> | undefined, b: RequestPasswordResetRequest | PlainMessage<RequestPasswordResetRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.RequestPasswordResetResponse
 */
export declare class RequestPasswordResetResponse extends Message<RequestPasswordResetResponse> {
  constructor(data?: PartialMessage<RequestPasswordResetResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.RequestPasswordResetResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RequestPasswordResetResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RequestPasswordResetResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RequestPasswordResetResponse;

  static equals(a: RequestPasswordResetResponse | PlainMessage<RequestPasswordResetResponse> | undefined, b: RequestPasswordResetResponse | PlainMessage<RequestPasswordResetResponse> | undefined): boolean;
}

