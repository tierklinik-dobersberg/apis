// @generated by protoc-gen-es v2.0.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/self_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";
import type { FieldMask, Timestamp, Value } from "@bufbuild/protobuf/wkt";
import type { Address, EMail, PhoneNumber, User } from "./user_pb";

/**
 * Describes the file tkd/idm/v1/self_service.proto.
 */
export declare const file_tkd_idm_v1_self_service: GenFile;

/**
 * @generated from message tkd.idm.v1.ChangePasswordRequest
 */
export declare type ChangePasswordRequest = Message<"tkd.idm.v1.ChangePasswordRequest"> & {
  /**
   * @generated from field: string old_password = 1;
   */
  oldPassword: string;

  /**
   * @generated from field: string new_password = 2;
   */
  newPassword: string;
};

/**
 * Describes the message tkd.idm.v1.ChangePasswordRequest.
 * Use `create(ChangePasswordRequestSchema)` to create a new message.
 */
export declare const ChangePasswordRequestSchema: GenMessage<ChangePasswordRequest>;

/**
 * @generated from message tkd.idm.v1.ChangePasswordResponse
 */
export declare type ChangePasswordResponse = Message<"tkd.idm.v1.ChangePasswordResponse"> & {
};

/**
 * Describes the message tkd.idm.v1.ChangePasswordResponse.
 * Use `create(ChangePasswordResponseSchema)` to create a new message.
 */
export declare const ChangePasswordResponseSchema: GenMessage<ChangePasswordResponse>;

/**
 * @generated from message tkd.idm.v1.UpdateProfileRequest
 */
export declare type UpdateProfileRequest = Message<"tkd.idm.v1.UpdateProfileRequest"> & {
  /**
   * @generated from field: string username = 1;
   */
  username: string;

  /**
   * @generated from field: string display_name = 2;
   */
  displayName: string;

  /**
   * @generated from field: string first_name = 3;
   */
  firstName: string;

  /**
   * @generated from field: string last_name = 4;
   */
  lastName: string;

  /**
   * @generated from field: map<string, google.protobuf.Value> extra = 5;
   */
  extra: { [key: string]: Value };

  /**
   * @generated from field: string avatar = 6;
   */
  avatar: string;

  /**
   * @generated from field: string birthday = 7;
   */
  birthday: string;

  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 8;
   */
  fieldMask?: FieldMask;
};

/**
 * Describes the message tkd.idm.v1.UpdateProfileRequest.
 * Use `create(UpdateProfileRequestSchema)` to create a new message.
 */
export declare const UpdateProfileRequestSchema: GenMessage<UpdateProfileRequest>;

/**
 * @generated from message tkd.idm.v1.UpdateProfileResponse
 */
export declare type UpdateProfileResponse = Message<"tkd.idm.v1.UpdateProfileResponse"> & {
  /**
   * @generated from field: tkd.idm.v1.User user = 1;
   */
  user?: User;
};

/**
 * Describes the message tkd.idm.v1.UpdateProfileResponse.
 * Use `create(UpdateProfileResponseSchema)` to create a new message.
 */
export declare const UpdateProfileResponseSchema: GenMessage<UpdateProfileResponse>;

/**
 * @generated from message tkd.idm.v1.ValidateEmailRequest
 */
export declare type ValidateEmailRequest = Message<"tkd.idm.v1.ValidateEmailRequest"> & {
  /**
   * @generated from oneof tkd.idm.v1.ValidateEmailRequest.kind
   */
  kind: {
    /**
     * @generated from field: string email_id = 1;
     */
    value: string;
    case: "emailId";
  } | {
    /**
     * @generated from field: string token = 2;
     */
    value: string;
    case: "token";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message tkd.idm.v1.ValidateEmailRequest.
 * Use `create(ValidateEmailRequestSchema)` to create a new message.
 */
export declare const ValidateEmailRequestSchema: GenMessage<ValidateEmailRequest>;

/**
 * @generated from message tkd.idm.v1.ValidateEmailResponse
 */
export declare type ValidateEmailResponse = Message<"tkd.idm.v1.ValidateEmailResponse"> & {
};

/**
 * Describes the message tkd.idm.v1.ValidateEmailResponse.
 * Use `create(ValidateEmailResponseSchema)` to create a new message.
 */
export declare const ValidateEmailResponseSchema: GenMessage<ValidateEmailResponse>;

/**
 * @generated from message tkd.idm.v1.AddEmailAddressRequest
 */
export declare type AddEmailAddressRequest = Message<"tkd.idm.v1.AddEmailAddressRequest"> & {
  /**
   * @generated from field: string email = 1;
   */
  email: string;
};

/**
 * Describes the message tkd.idm.v1.AddEmailAddressRequest.
 * Use `create(AddEmailAddressRequestSchema)` to create a new message.
 */
export declare const AddEmailAddressRequestSchema: GenMessage<AddEmailAddressRequest>;

/**
 * @generated from message tkd.idm.v1.AddEmailAddressResponse
 */
export declare type AddEmailAddressResponse = Message<"tkd.idm.v1.AddEmailAddressResponse"> & {
  /**
   * @generated from field: repeated tkd.idm.v1.EMail emails = 1;
   */
  emails: EMail[];
};

/**
 * Describes the message tkd.idm.v1.AddEmailAddressResponse.
 * Use `create(AddEmailAddressResponseSchema)` to create a new message.
 */
export declare const AddEmailAddressResponseSchema: GenMessage<AddEmailAddressResponse>;

/**
 * @generated from message tkd.idm.v1.DeleteEmailAddressRequest
 */
export declare type DeleteEmailAddressRequest = Message<"tkd.idm.v1.DeleteEmailAddressRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message tkd.idm.v1.DeleteEmailAddressRequest.
 * Use `create(DeleteEmailAddressRequestSchema)` to create a new message.
 */
export declare const DeleteEmailAddressRequestSchema: GenMessage<DeleteEmailAddressRequest>;

/**
 * @generated from message tkd.idm.v1.DeleteEmailAddressResponse
 */
export declare type DeleteEmailAddressResponse = Message<"tkd.idm.v1.DeleteEmailAddressResponse"> & {
  /**
   * @generated from field: repeated tkd.idm.v1.EMail emails = 1;
   */
  emails: EMail[];
};

/**
 * Describes the message tkd.idm.v1.DeleteEmailAddressResponse.
 * Use `create(DeleteEmailAddressResponseSchema)` to create a new message.
 */
export declare const DeleteEmailAddressResponseSchema: GenMessage<DeleteEmailAddressResponse>;

/**
 * @generated from message tkd.idm.v1.AddAddressRequest
 */
export declare type AddAddressRequest = Message<"tkd.idm.v1.AddAddressRequest"> & {
  /**
   * @generated from field: string city_code = 2;
   */
  cityCode: string;

  /**
   * @generated from field: string city_name = 3;
   */
  cityName: string;

  /**
   * @generated from field: string street = 4;
   */
  street: string;

  /**
   * @generated from field: string extra = 5;
   */
  extra: string;
};

/**
 * Describes the message tkd.idm.v1.AddAddressRequest.
 * Use `create(AddAddressRequestSchema)` to create a new message.
 */
export declare const AddAddressRequestSchema: GenMessage<AddAddressRequest>;

/**
 * @generated from message tkd.idm.v1.AddAddressResponse
 */
export declare type AddAddressResponse = Message<"tkd.idm.v1.AddAddressResponse"> & {
  /**
   * @generated from field: repeated tkd.idm.v1.Address addresses = 1;
   */
  addresses: Address[];
};

/**
 * Describes the message tkd.idm.v1.AddAddressResponse.
 * Use `create(AddAddressResponseSchema)` to create a new message.
 */
export declare const AddAddressResponseSchema: GenMessage<AddAddressResponse>;

/**
 * @generated from message tkd.idm.v1.UpdateAddressRequest
 */
export declare type UpdateAddressRequest = Message<"tkd.idm.v1.UpdateAddressRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string city_code = 2;
   */
  cityCode: string;

  /**
   * @generated from field: string city_name = 3;
   */
  cityName: string;

  /**
   * @generated from field: string street = 4;
   */
  street: string;

  /**
   * @generated from field: string extra = 5;
   */
  extra: string;

  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 6;
   */
  fieldMask?: FieldMask;
};

/**
 * Describes the message tkd.idm.v1.UpdateAddressRequest.
 * Use `create(UpdateAddressRequestSchema)` to create a new message.
 */
export declare const UpdateAddressRequestSchema: GenMessage<UpdateAddressRequest>;

/**
 * @generated from message tkd.idm.v1.UpdateAddressResponse
 */
export declare type UpdateAddressResponse = Message<"tkd.idm.v1.UpdateAddressResponse"> & {
  /**
   * @generated from field: repeated tkd.idm.v1.Address addresses = 1;
   */
  addresses: Address[];
};

/**
 * Describes the message tkd.idm.v1.UpdateAddressResponse.
 * Use `create(UpdateAddressResponseSchema)` to create a new message.
 */
export declare const UpdateAddressResponseSchema: GenMessage<UpdateAddressResponse>;

/**
 * @generated from message tkd.idm.v1.DeleteAddressRequest
 */
export declare type DeleteAddressRequest = Message<"tkd.idm.v1.DeleteAddressRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message tkd.idm.v1.DeleteAddressRequest.
 * Use `create(DeleteAddressRequestSchema)` to create a new message.
 */
export declare const DeleteAddressRequestSchema: GenMessage<DeleteAddressRequest>;

/**
 * @generated from message tkd.idm.v1.DeleteAddressResponse
 */
export declare type DeleteAddressResponse = Message<"tkd.idm.v1.DeleteAddressResponse"> & {
  /**
   * @generated from field: repeated tkd.idm.v1.Address addresses = 1;
   */
  addresses: Address[];
};

/**
 * Describes the message tkd.idm.v1.DeleteAddressResponse.
 * Use `create(DeleteAddressResponseSchema)` to create a new message.
 */
export declare const DeleteAddressResponseSchema: GenMessage<DeleteAddressResponse>;

/**
 * @generated from message tkd.idm.v1.MarkEmailAsPrimaryRequest
 */
export declare type MarkEmailAsPrimaryRequest = Message<"tkd.idm.v1.MarkEmailAsPrimaryRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message tkd.idm.v1.MarkEmailAsPrimaryRequest.
 * Use `create(MarkEmailAsPrimaryRequestSchema)` to create a new message.
 */
export declare const MarkEmailAsPrimaryRequestSchema: GenMessage<MarkEmailAsPrimaryRequest>;

/**
 * TODO(ppacher): return all emails here?
 *
 * @generated from message tkd.idm.v1.MarkEmailAsPrimaryResponse
 */
export declare type MarkEmailAsPrimaryResponse = Message<"tkd.idm.v1.MarkEmailAsPrimaryResponse"> & {
};

/**
 * Describes the message tkd.idm.v1.MarkEmailAsPrimaryResponse.
 * Use `create(MarkEmailAsPrimaryResponseSchema)` to create a new message.
 */
export declare const MarkEmailAsPrimaryResponseSchema: GenMessage<MarkEmailAsPrimaryResponse>;

/**
 * @generated from message tkd.idm.v1.AddPhoneNumberRequest
 */
export declare type AddPhoneNumberRequest = Message<"tkd.idm.v1.AddPhoneNumberRequest"> & {
  /**
   * @generated from field: string number = 1;
   */
  number: string;
};

/**
 * Describes the message tkd.idm.v1.AddPhoneNumberRequest.
 * Use `create(AddPhoneNumberRequestSchema)` to create a new message.
 */
export declare const AddPhoneNumberRequestSchema: GenMessage<AddPhoneNumberRequest>;

/**
 * @generated from message tkd.idm.v1.AddPhoneNumberResponse
 */
export declare type AddPhoneNumberResponse = Message<"tkd.idm.v1.AddPhoneNumberResponse"> & {
  /**
   * @generated from field: tkd.idm.v1.PhoneNumber phone_number = 1;
   */
  phoneNumber?: PhoneNumber;
};

/**
 * Describes the message tkd.idm.v1.AddPhoneNumberResponse.
 * Use `create(AddPhoneNumberResponseSchema)` to create a new message.
 */
export declare const AddPhoneNumberResponseSchema: GenMessage<AddPhoneNumberResponse>;

/**
 * @generated from message tkd.idm.v1.DeletePhoneNumberRequest
 */
export declare type DeletePhoneNumberRequest = Message<"tkd.idm.v1.DeletePhoneNumberRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message tkd.idm.v1.DeletePhoneNumberRequest.
 * Use `create(DeletePhoneNumberRequestSchema)` to create a new message.
 */
export declare const DeletePhoneNumberRequestSchema: GenMessage<DeletePhoneNumberRequest>;

/**
 * @generated from message tkd.idm.v1.DeletePhoneNumberResponse
 */
export declare type DeletePhoneNumberResponse = Message<"tkd.idm.v1.DeletePhoneNumberResponse"> & {
};

/**
 * Describes the message tkd.idm.v1.DeletePhoneNumberResponse.
 * Use `create(DeletePhoneNumberResponseSchema)` to create a new message.
 */
export declare const DeletePhoneNumberResponseSchema: GenMessage<DeletePhoneNumberResponse>;

/**
 * @generated from message tkd.idm.v1.MarkPhoneNumberAsPrimaryRequest
 */
export declare type MarkPhoneNumberAsPrimaryRequest = Message<"tkd.idm.v1.MarkPhoneNumberAsPrimaryRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message tkd.idm.v1.MarkPhoneNumberAsPrimaryRequest.
 * Use `create(MarkPhoneNumberAsPrimaryRequestSchema)` to create a new message.
 */
export declare const MarkPhoneNumberAsPrimaryRequestSchema: GenMessage<MarkPhoneNumberAsPrimaryRequest>;

/**
 * @generated from message tkd.idm.v1.MarkPhoneNumberAsPrimaryResponse
 */
export declare type MarkPhoneNumberAsPrimaryResponse = Message<"tkd.idm.v1.MarkPhoneNumberAsPrimaryResponse"> & {
};

/**
 * Describes the message tkd.idm.v1.MarkPhoneNumberAsPrimaryResponse.
 * Use `create(MarkPhoneNumberAsPrimaryResponseSchema)` to create a new message.
 */
export declare const MarkPhoneNumberAsPrimaryResponseSchema: GenMessage<MarkPhoneNumberAsPrimaryResponse>;

/**
 * @generated from message tkd.idm.v1.ValidatePhoneNumberRequest
 */
export declare type ValidatePhoneNumberRequest = Message<"tkd.idm.v1.ValidatePhoneNumberRequest"> & {
  /**
   * @generated from oneof tkd.idm.v1.ValidatePhoneNumberRequest.step
   */
  step: {
    /**
     * @generated from field: string id = 1;
     */
    value: string;
    case: "id";
  } | {
    /**
     * @generated from field: string code = 2;
     */
    value: string;
    case: "code";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message tkd.idm.v1.ValidatePhoneNumberRequest.
 * Use `create(ValidatePhoneNumberRequestSchema)` to create a new message.
 */
export declare const ValidatePhoneNumberRequestSchema: GenMessage<ValidatePhoneNumberRequest>;

/**
 * @generated from message tkd.idm.v1.ValidatePhoneNumberResponse
 */
export declare type ValidatePhoneNumberResponse = Message<"tkd.idm.v1.ValidatePhoneNumberResponse"> & {
};

/**
 * Describes the message tkd.idm.v1.ValidatePhoneNumberResponse.
 * Use `create(ValidatePhoneNumberResponseSchema)` to create a new message.
 */
export declare const ValidatePhoneNumberResponseSchema: GenMessage<ValidatePhoneNumberResponse>;

/**
 * @generated from message tkd.idm.v1.EnrollTOTPRequestStep1
 */
export declare type EnrollTOTPRequestStep1 = Message<"tkd.idm.v1.EnrollTOTPRequestStep1"> & {
};

/**
 * Describes the message tkd.idm.v1.EnrollTOTPRequestStep1.
 * Use `create(EnrollTOTPRequestStep1Schema)` to create a new message.
 */
export declare const EnrollTOTPRequestStep1Schema: GenMessage<EnrollTOTPRequestStep1>;

/**
 * @generated from message tkd.idm.v1.EnrollTOTPRequestStep2
 */
export declare type EnrollTOTPRequestStep2 = Message<"tkd.idm.v1.EnrollTOTPRequestStep2"> & {
  /**
   * @generated from field: string verify_code = 1;
   */
  verifyCode: string;

  /**
   * @generated from field: string secret = 2;
   */
  secret: string;

  /**
   * @generated from field: string secret_hmac = 3;
   */
  secretHmac: string;
};

/**
 * Describes the message tkd.idm.v1.EnrollTOTPRequestStep2.
 * Use `create(EnrollTOTPRequestStep2Schema)` to create a new message.
 */
export declare const EnrollTOTPRequestStep2Schema: GenMessage<EnrollTOTPRequestStep2>;

/**
 * @generated from message tkd.idm.v1.Enroll2FARequest
 */
export declare type Enroll2FARequest = Message<"tkd.idm.v1.Enroll2FARequest"> & {
  /**
   * @generated from oneof tkd.idm.v1.Enroll2FARequest.kind
   */
  kind: {
    /**
     * @generated from field: tkd.idm.v1.EnrollTOTPRequestStep1 totp_step1 = 1;
     */
    value: EnrollTOTPRequestStep1;
    case: "totpStep1";
  } | {
    /**
     * @generated from field: tkd.idm.v1.EnrollTOTPRequestStep2 totp_step2 = 2;
     */
    value: EnrollTOTPRequestStep2;
    case: "totpStep2";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message tkd.idm.v1.Enroll2FARequest.
 * Use `create(Enroll2FARequestSchema)` to create a new message.
 */
export declare const Enroll2FARequestSchema: GenMessage<Enroll2FARequest>;

/**
 * @generated from message tkd.idm.v1.EnrollTOTPResponseStep1
 */
export declare type EnrollTOTPResponseStep1 = Message<"tkd.idm.v1.EnrollTOTPResponseStep1"> & {
  /**
   * @generated from field: string secret = 2;
   */
  secret: string;

  /**
   * @generated from field: string qr_code = 3;
   */
  qrCode: string;

  /**
   * @generated from field: string secret_hmac = 4;
   */
  secretHmac: string;

  /**
   * @generated from field: string url = 5;
   */
  url: string;
};

/**
 * Describes the message tkd.idm.v1.EnrollTOTPResponseStep1.
 * Use `create(EnrollTOTPResponseStep1Schema)` to create a new message.
 */
export declare const EnrollTOTPResponseStep1Schema: GenMessage<EnrollTOTPResponseStep1>;

/**
 * @generated from message tkd.idm.v1.EnrollTOTPResponseStep2
 */
export declare type EnrollTOTPResponseStep2 = Message<"tkd.idm.v1.EnrollTOTPResponseStep2"> & {
};

/**
 * Describes the message tkd.idm.v1.EnrollTOTPResponseStep2.
 * Use `create(EnrollTOTPResponseStep2Schema)` to create a new message.
 */
export declare const EnrollTOTPResponseStep2Schema: GenMessage<EnrollTOTPResponseStep2>;

/**
 * @generated from message tkd.idm.v1.Enroll2FAResponse
 */
export declare type Enroll2FAResponse = Message<"tkd.idm.v1.Enroll2FAResponse"> & {
  /**
   * @generated from oneof tkd.idm.v1.Enroll2FAResponse.kind
   */
  kind: {
    /**
     * @generated from field: tkd.idm.v1.EnrollTOTPResponseStep1 totp_step1 = 1;
     */
    value: EnrollTOTPResponseStep1;
    case: "totpStep1";
  } | {
    /**
     * @generated from field: tkd.idm.v1.EnrollTOTPResponseStep2 totp_step2 = 2;
     */
    value: EnrollTOTPResponseStep2;
    case: "totpStep2";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message tkd.idm.v1.Enroll2FAResponse.
 * Use `create(Enroll2FAResponseSchema)` to create a new message.
 */
export declare const Enroll2FAResponseSchema: GenMessage<Enroll2FAResponse>;

/**
 * @generated from message tkd.idm.v1.Remove2FARequest
 */
export declare type Remove2FARequest = Message<"tkd.idm.v1.Remove2FARequest"> & {
  /**
   * @generated from oneof tkd.idm.v1.Remove2FARequest.kind
   */
  kind: {
    /**
     * If set and valid, the TOTP 2FA will be disabled.
     *
     * @generated from field: string totp_code = 1;
     */
    value: string;
    case: "totpCode";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message tkd.idm.v1.Remove2FARequest.
 * Use `create(Remove2FARequestSchema)` to create a new message.
 */
export declare const Remove2FARequestSchema: GenMessage<Remove2FARequest>;

/**
 * @generated from message tkd.idm.v1.Remove2FAResponse
 */
export declare type Remove2FAResponse = Message<"tkd.idm.v1.Remove2FAResponse"> & {
};

/**
 * Describes the message tkd.idm.v1.Remove2FAResponse.
 * Use `create(Remove2FAResponseSchema)` to create a new message.
 */
export declare const Remove2FAResponseSchema: GenMessage<Remove2FAResponse>;

/**
 * @generated from message tkd.idm.v1.GenerateRecoveryCodesRequest
 */
export declare type GenerateRecoveryCodesRequest = Message<"tkd.idm.v1.GenerateRecoveryCodesRequest"> & {
};

/**
 * Describes the message tkd.idm.v1.GenerateRecoveryCodesRequest.
 * Use `create(GenerateRecoveryCodesRequestSchema)` to create a new message.
 */
export declare const GenerateRecoveryCodesRequestSchema: GenMessage<GenerateRecoveryCodesRequest>;

/**
 * @generated from message tkd.idm.v1.GenerateRecoveryCodesResponse
 */
export declare type GenerateRecoveryCodesResponse = Message<"tkd.idm.v1.GenerateRecoveryCodesResponse"> & {
  /**
   * @generated from field: repeated string recovery_codes = 1;
   */
  recoveryCodes: string[];
};

/**
 * Describes the message tkd.idm.v1.GenerateRecoveryCodesResponse.
 * Use `create(GenerateRecoveryCodesResponseSchema)` to create a new message.
 */
export declare const GenerateRecoveryCodesResponseSchema: GenMessage<GenerateRecoveryCodesResponse>;

/**
 * @generated from message tkd.idm.v1.RegisteredPasskey
 */
export declare type RegisteredPasskey = Message<"tkd.idm.v1.RegisteredPasskey"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string client_name = 2;
   */
  clientName: string;

  /**
   * @generated from field: string client_os = 3;
   */
  clientOs: string;

  /**
   * @generated from field: string client_device = 4;
   */
  clientDevice: string;

  /**
   * @generated from field: string cred_type = 5;
   */
  credType: string;

  /**
   * @generated from field: string passkey_name = 6;
   */
  passkeyName: string;
};

/**
 * Describes the message tkd.idm.v1.RegisteredPasskey.
 * Use `create(RegisteredPasskeySchema)` to create a new message.
 */
export declare const RegisteredPasskeySchema: GenMessage<RegisteredPasskey>;

/**
 * @generated from message tkd.idm.v1.GetRegisteredPasskeysRequest
 */
export declare type GetRegisteredPasskeysRequest = Message<"tkd.idm.v1.GetRegisteredPasskeysRequest"> & {
};

/**
 * Describes the message tkd.idm.v1.GetRegisteredPasskeysRequest.
 * Use `create(GetRegisteredPasskeysRequestSchema)` to create a new message.
 */
export declare const GetRegisteredPasskeysRequestSchema: GenMessage<GetRegisteredPasskeysRequest>;

/**
 * @generated from message tkd.idm.v1.GetRegisteredPasskeysResponse
 */
export declare type GetRegisteredPasskeysResponse = Message<"tkd.idm.v1.GetRegisteredPasskeysResponse"> & {
  /**
   * @generated from field: repeated tkd.idm.v1.RegisteredPasskey passkeys = 1;
   */
  passkeys: RegisteredPasskey[];
};

/**
 * Describes the message tkd.idm.v1.GetRegisteredPasskeysResponse.
 * Use `create(GetRegisteredPasskeysResponseSchema)` to create a new message.
 */
export declare const GetRegisteredPasskeysResponseSchema: GenMessage<GetRegisteredPasskeysResponse>;

/**
 * @generated from message tkd.idm.v1.RemovePasskeyRequest
 */
export declare type RemovePasskeyRequest = Message<"tkd.idm.v1.RemovePasskeyRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message tkd.idm.v1.RemovePasskeyRequest.
 * Use `create(RemovePasskeyRequestSchema)` to create a new message.
 */
export declare const RemovePasskeyRequestSchema: GenMessage<RemovePasskeyRequest>;

/**
 * @generated from message tkd.idm.v1.RemovePasskeyResponse
 */
export declare type RemovePasskeyResponse = Message<"tkd.idm.v1.RemovePasskeyResponse"> & {
};

/**
 * Describes the message tkd.idm.v1.RemovePasskeyResponse.
 * Use `create(RemovePasskeyResponseSchema)` to create a new message.
 */
export declare const RemovePasskeyResponseSchema: GenMessage<RemovePasskeyResponse>;

/**
 * @generated from message tkd.idm.v1.GenerateAPITokenRequest
 */
export declare type GenerateAPITokenRequest = Message<"tkd.idm.v1.GenerateAPITokenRequest"> & {
  /**
   * A list of roles that should be assigned to the token.
   * Note that it is not possible to assign roles that the
   * authenticated user doesn't have.
   *
   * @generated from field: repeated string roles = 1;
   */
  roles: string[];

  /**
   * The timestamp at which the api token should expire.
   *
   * @generated from field: google.protobuf.Timestamp expires = 2;
   */
  expires?: Timestamp;

  /**
   * A human readable description to identify the API token.
   *
   * @generated from field: string description = 3;
   */
  description: string;
};

/**
 * Describes the message tkd.idm.v1.GenerateAPITokenRequest.
 * Use `create(GenerateAPITokenRequestSchema)` to create a new message.
 */
export declare const GenerateAPITokenRequestSchema: GenMessage<GenerateAPITokenRequest>;

/**
 * @generated from message tkd.idm.v1.GenerateAPITokenResponse
 */
export declare type GenerateAPITokenResponse = Message<"tkd.idm.v1.GenerateAPITokenResponse"> & {
  /**
   * @generated from field: tkd.idm.v1.APIToken token = 1;
   */
  token?: APIToken;
};

/**
 * Describes the message tkd.idm.v1.GenerateAPITokenResponse.
 * Use `create(GenerateAPITokenResponseSchema)` to create a new message.
 */
export declare const GenerateAPITokenResponseSchema: GenMessage<GenerateAPITokenResponse>;

/**
 * @generated from message tkd.idm.v1.ListAPITokensRequest
 */
export declare type ListAPITokensRequest = Message<"tkd.idm.v1.ListAPITokensRequest"> & {
};

/**
 * Describes the message tkd.idm.v1.ListAPITokensRequest.
 * Use `create(ListAPITokensRequestSchema)` to create a new message.
 */
export declare const ListAPITokensRequestSchema: GenMessage<ListAPITokensRequest>;

/**
 * @generated from message tkd.idm.v1.APIToken
 */
export declare type APIToken = Message<"tkd.idm.v1.APIToken"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string description = 2;
   */
  description: string;

  /**
   * @generated from field: string redacted_token = 3;
   */
  redactedToken: string;

  /**
   * @generated from field: google.protobuf.Timestamp expires = 4;
   */
  expires?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 5;
   */
  createdAt?: Timestamp;
};

/**
 * Describes the message tkd.idm.v1.APIToken.
 * Use `create(APITokenSchema)` to create a new message.
 */
export declare const APITokenSchema: GenMessage<APIToken>;

/**
 * @generated from message tkd.idm.v1.ListAPITokensResponse
 */
export declare type ListAPITokensResponse = Message<"tkd.idm.v1.ListAPITokensResponse"> & {
  /**
   * @generated from field: repeated tkd.idm.v1.APIToken tokens = 1;
   */
  tokens: APIToken[];
};

/**
 * Describes the message tkd.idm.v1.ListAPITokensResponse.
 * Use `create(ListAPITokensResponseSchema)` to create a new message.
 */
export declare const ListAPITokensResponseSchema: GenMessage<ListAPITokensResponse>;

/**
 * @generated from message tkd.idm.v1.RemoveAPITokenRequest
 */
export declare type RemoveAPITokenRequest = Message<"tkd.idm.v1.RemoveAPITokenRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message tkd.idm.v1.RemoveAPITokenRequest.
 * Use `create(RemoveAPITokenRequestSchema)` to create a new message.
 */
export declare const RemoveAPITokenRequestSchema: GenMessage<RemoveAPITokenRequest>;

/**
 * @generated from message tkd.idm.v1.RemoveAPITokenResponse
 */
export declare type RemoveAPITokenResponse = Message<"tkd.idm.v1.RemoveAPITokenResponse"> & {
};

/**
 * Describes the message tkd.idm.v1.RemoveAPITokenResponse.
 * Use `create(RemoveAPITokenResponseSchema)` to create a new message.
 */
export declare const RemoveAPITokenResponseSchema: GenMessage<RemoveAPITokenResponse>;

/**
 * @generated from service tkd.idm.v1.SelfServiceService
 */
export declare const SelfServiceService: GenService<{
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.ChangePassword
   */
  changePassword: {
    methodKind: "unary";
    input: typeof ChangePasswordRequestSchema;
    output: typeof ChangePasswordResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.UpdateProfile
   */
  updateProfile: {
    methodKind: "unary";
    input: typeof UpdateProfileRequestSchema;
    output: typeof UpdateProfileResponseSchema;
  },
  /**
   * E-Mails
   *
   * @generated from rpc tkd.idm.v1.SelfServiceService.AddEmailAddress
   */
  addEmailAddress: {
    methodKind: "unary";
    input: typeof AddEmailAddressRequestSchema;
    output: typeof AddEmailAddressResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.DeleteEmailAddress
   */
  deleteEmailAddress: {
    methodKind: "unary";
    input: typeof DeleteEmailAddressRequestSchema;
    output: typeof DeleteEmailAddressResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.MarkEmailAsPrimary
   */
  markEmailAsPrimary: {
    methodKind: "unary";
    input: typeof MarkEmailAsPrimaryRequestSchema;
    output: typeof MarkEmailAsPrimaryResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.ValidateEmail
   */
  validateEmail: {
    methodKind: "unary";
    input: typeof ValidateEmailRequestSchema;
    output: typeof ValidateEmailResponseSchema;
  },
  /**
   * Addresses
   *
   * @generated from rpc tkd.idm.v1.SelfServiceService.AddAddress
   */
  addAddress: {
    methodKind: "unary";
    input: typeof AddAddressRequestSchema;
    output: typeof AddAddressResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.DeleteAddress
   */
  deleteAddress: {
    methodKind: "unary";
    input: typeof DeleteAddressRequestSchema;
    output: typeof DeleteAddressResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.UpdateAddress
   */
  updateAddress: {
    methodKind: "unary";
    input: typeof UpdateAddressRequestSchema;
    output: typeof UpdateAddressResponseSchema;
  },
  /**
   * Phone Numbers
   *
   * @generated from rpc tkd.idm.v1.SelfServiceService.AddPhoneNumber
   */
  addPhoneNumber: {
    methodKind: "unary";
    input: typeof AddPhoneNumberRequestSchema;
    output: typeof AddPhoneNumberResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.DeletePhoneNumber
   */
  deletePhoneNumber: {
    methodKind: "unary";
    input: typeof DeletePhoneNumberRequestSchema;
    output: typeof DeletePhoneNumberResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.MarkPhoneNumberAsPrimary
   */
  markPhoneNumberAsPrimary: {
    methodKind: "unary";
    input: typeof MarkPhoneNumberAsPrimaryRequestSchema;
    output: typeof MarkPhoneNumberAsPrimaryResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.ValidatePhoneNumber
   */
  validatePhoneNumber: {
    methodKind: "unary";
    input: typeof ValidatePhoneNumberRequestSchema;
    output: typeof ValidatePhoneNumberResponseSchema;
  },
  /**
   * WebAuthN
   *
   * @generated from rpc tkd.idm.v1.SelfServiceService.GetRegisteredPasskeys
   */
  getRegisteredPasskeys: {
    methodKind: "unary";
    input: typeof GetRegisteredPasskeysRequestSchema;
    output: typeof GetRegisteredPasskeysResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.RemovePasskey
   */
  removePasskey: {
    methodKind: "unary";
    input: typeof RemovePasskeyRequestSchema;
    output: typeof RemovePasskeyResponseSchema;
  },
  /**
   * Two-Factor authentication
   *
   * @generated from rpc tkd.idm.v1.SelfServiceService.Enroll2FA
   */
  enroll2FA: {
    methodKind: "unary";
    input: typeof Enroll2FARequestSchema;
    output: typeof Enroll2FAResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.Remove2FA
   */
  remove2FA: {
    methodKind: "unary";
    input: typeof Remove2FARequestSchema;
    output: typeof Remove2FAResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.GenerateRecoveryCodes
   */
  generateRecoveryCodes: {
    methodKind: "unary";
    input: typeof GenerateRecoveryCodesRequestSchema;
    output: typeof GenerateRecoveryCodesResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.GenerateAPIToken
   */
  generateAPIToken: {
    methodKind: "unary";
    input: typeof GenerateAPITokenRequestSchema;
    output: typeof GenerateAPITokenResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.ListAPITokens
   */
  listAPITokens: {
    methodKind: "unary";
    input: typeof ListAPITokensRequestSchema;
    output: typeof ListAPITokensResponseSchema;
  },
  /**
   * @generated from rpc tkd.idm.v1.SelfServiceService.RemoveAPIToken
   */
  removeAPIToken: {
    methodKind: "unary";
    input: typeof RemoveAPITokenRequestSchema;
    output: typeof RemoveAPITokenResponseSchema;
  },
}>;

