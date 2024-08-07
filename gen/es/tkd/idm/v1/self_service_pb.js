// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/self_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FieldMask, proto3, Timestamp, Value } from "@bufbuild/protobuf";
import { Address, EMail, PhoneNumber, User } from "./user_pb.js";

/**
 * @generated from message tkd.idm.v1.ChangePasswordRequest
 */
export const ChangePasswordRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ChangePasswordRequest",
  () => [
    { no: 1, name: "old_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "new_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.ChangePasswordResponse
 */
export const ChangePasswordResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ChangePasswordResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.UpdateProfileRequest
 */
export const UpdateProfileRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.UpdateProfileRequest",
  () => [
    { no: 1, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "first_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "last_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "extra", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Value} },
    { no: 6, name: "avatar", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "birthday", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "field_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.idm.v1.UpdateProfileResponse
 */
export const UpdateProfileResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.UpdateProfileResponse",
  () => [
    { no: 1, name: "user", kind: "message", T: User },
  ],
);

/**
 * @generated from message tkd.idm.v1.ValidateEmailRequest
 */
export const ValidateEmailRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ValidateEmailRequest",
  () => [
    { no: 1, name: "email_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "kind" },
    { no: 2, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "kind" },
  ],
);

/**
 * @generated from message tkd.idm.v1.ValidateEmailResponse
 */
export const ValidateEmailResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ValidateEmailResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.AddEmailAddressRequest
 */
export const AddEmailAddressRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.AddEmailAddressRequest",
  () => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.AddEmailAddressResponse
 */
export const AddEmailAddressResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.AddEmailAddressResponse",
  () => [
    { no: 1, name: "emails", kind: "message", T: EMail, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeleteEmailAddressRequest
 */
export const DeleteEmailAddressRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeleteEmailAddressRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeleteEmailAddressResponse
 */
export const DeleteEmailAddressResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeleteEmailAddressResponse",
  () => [
    { no: 1, name: "emails", kind: "message", T: EMail, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.AddAddressRequest
 */
export const AddAddressRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.AddAddressRequest",
  () => [
    { no: 2, name: "city_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "city_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "street", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "extra", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.AddAddressResponse
 */
export const AddAddressResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.AddAddressResponse",
  () => [
    { no: 1, name: "addresses", kind: "message", T: Address, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.UpdateAddressRequest
 */
export const UpdateAddressRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.UpdateAddressRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "city_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "city_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "street", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "extra", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "field_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.idm.v1.UpdateAddressResponse
 */
export const UpdateAddressResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.UpdateAddressResponse",
  () => [
    { no: 1, name: "addresses", kind: "message", T: Address, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeleteAddressRequest
 */
export const DeleteAddressRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeleteAddressRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeleteAddressResponse
 */
export const DeleteAddressResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeleteAddressResponse",
  () => [
    { no: 1, name: "addresses", kind: "message", T: Address, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.MarkEmailAsPrimaryRequest
 */
export const MarkEmailAsPrimaryRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.MarkEmailAsPrimaryRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * TODO(ppacher): return all emails here?
 *
 * @generated from message tkd.idm.v1.MarkEmailAsPrimaryResponse
 */
export const MarkEmailAsPrimaryResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.MarkEmailAsPrimaryResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.AddPhoneNumberRequest
 */
export const AddPhoneNumberRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.AddPhoneNumberRequest",
  () => [
    { no: 1, name: "number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.AddPhoneNumberResponse
 */
export const AddPhoneNumberResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.AddPhoneNumberResponse",
  () => [
    { no: 1, name: "phone_number", kind: "message", T: PhoneNumber },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeletePhoneNumberRequest
 */
export const DeletePhoneNumberRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeletePhoneNumberRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeletePhoneNumberResponse
 */
export const DeletePhoneNumberResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeletePhoneNumberResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.MarkPhoneNumberAsPrimaryRequest
 */
export const MarkPhoneNumberAsPrimaryRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.MarkPhoneNumberAsPrimaryRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.MarkPhoneNumberAsPrimaryResponse
 */
export const MarkPhoneNumberAsPrimaryResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.MarkPhoneNumberAsPrimaryResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.ValidatePhoneNumberRequest
 */
export const ValidatePhoneNumberRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ValidatePhoneNumberRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "step" },
    { no: 2, name: "code", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "step" },
  ],
);

/**
 * @generated from message tkd.idm.v1.ValidatePhoneNumberResponse
 */
export const ValidatePhoneNumberResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ValidatePhoneNumberResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.EnrollTOTPRequestStep1
 */
export const EnrollTOTPRequestStep1 = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.EnrollTOTPRequestStep1",
  [],
);

/**
 * @generated from message tkd.idm.v1.EnrollTOTPRequestStep2
 */
export const EnrollTOTPRequestStep2 = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.EnrollTOTPRequestStep2",
  () => [
    { no: 1, name: "verify_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "secret", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "secret_hmac", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.Enroll2FARequest
 */
export const Enroll2FARequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.Enroll2FARequest",
  () => [
    { no: 1, name: "totp_step1", kind: "message", T: EnrollTOTPRequestStep1, oneof: "kind" },
    { no: 2, name: "totp_step2", kind: "message", T: EnrollTOTPRequestStep2, oneof: "kind" },
  ],
);

/**
 * @generated from message tkd.idm.v1.EnrollTOTPResponseStep1
 */
export const EnrollTOTPResponseStep1 = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.EnrollTOTPResponseStep1",
  () => [
    { no: 2, name: "secret", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "qr_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "secret_hmac", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.EnrollTOTPResponseStep2
 */
export const EnrollTOTPResponseStep2 = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.EnrollTOTPResponseStep2",
  [],
);

/**
 * @generated from message tkd.idm.v1.Enroll2FAResponse
 */
export const Enroll2FAResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.Enroll2FAResponse",
  () => [
    { no: 1, name: "totp_step1", kind: "message", T: EnrollTOTPResponseStep1, oneof: "kind" },
    { no: 2, name: "totp_step2", kind: "message", T: EnrollTOTPResponseStep2, oneof: "kind" },
  ],
);

/**
 * @generated from message tkd.idm.v1.Remove2FARequest
 */
export const Remove2FARequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.Remove2FARequest",
  () => [
    { no: 1, name: "totp_code", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "kind" },
  ],
);

/**
 * @generated from message tkd.idm.v1.Remove2FAResponse
 */
export const Remove2FAResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.Remove2FAResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.GenerateRecoveryCodesRequest
 */
export const GenerateRecoveryCodesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.GenerateRecoveryCodesRequest",
  [],
);

/**
 * @generated from message tkd.idm.v1.GenerateRecoveryCodesResponse
 */
export const GenerateRecoveryCodesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.GenerateRecoveryCodesResponse",
  () => [
    { no: 1, name: "recovery_codes", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.RegisteredPasskey
 */
export const RegisteredPasskey = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.RegisteredPasskey",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "client_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "client_os", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "client_device", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "cred_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "passkey_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.GetRegisteredPasskeysRequest
 */
export const GetRegisteredPasskeysRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.GetRegisteredPasskeysRequest",
  [],
);

/**
 * @generated from message tkd.idm.v1.GetRegisteredPasskeysResponse
 */
export const GetRegisteredPasskeysResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.GetRegisteredPasskeysResponse",
  () => [
    { no: 1, name: "passkeys", kind: "message", T: RegisteredPasskey, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.RemovePasskeyRequest
 */
export const RemovePasskeyRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.RemovePasskeyRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.RemovePasskeyResponse
 */
export const RemovePasskeyResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.RemovePasskeyResponse",
  [],
);

/**
 * @generated from message tkd.idm.v1.GenerateAPITokenRequest
 */
export const GenerateAPITokenRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.GenerateAPITokenRequest",
  () => [
    { no: 1, name: "roles", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "expires", kind: "message", T: Timestamp },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.GenerateAPITokenResponse
 */
export const GenerateAPITokenResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.GenerateAPITokenResponse",
  () => [
    { no: 1, name: "token", kind: "message", T: APIToken },
  ],
);

/**
 * @generated from message tkd.idm.v1.ListAPITokensRequest
 */
export const ListAPITokensRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ListAPITokensRequest",
  [],
);

/**
 * @generated from message tkd.idm.v1.APIToken
 */
export const APIToken = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.APIToken",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "redacted_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "expires", kind: "message", T: Timestamp },
    { no: 5, name: "created_at", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message tkd.idm.v1.ListAPITokensResponse
 */
export const ListAPITokensResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ListAPITokensResponse",
  () => [
    { no: 1, name: "tokens", kind: "message", T: APIToken, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.RemoveAPITokenRequest
 */
export const RemoveAPITokenRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.RemoveAPITokenRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.RemoveAPITokenResponse
 */
export const RemoveAPITokenResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.RemoveAPITokenResponse",
  [],
);

