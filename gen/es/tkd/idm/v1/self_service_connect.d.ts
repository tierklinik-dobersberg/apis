// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/self_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddAddressRequest, AddAddressResponse, AddEmailAddressRequest, AddEmailAddressResponse, AddPhoneNumberRequest, AddPhoneNumberResponse, ChangePasswordRequest, ChangePasswordResponse, DeleteAddressRequest, DeleteAddressResponse, DeleteEmailAddressRequest, DeleteEmailAddressResponse, DeletePhoneNumberRequest, DeletePhoneNumberResponse, Enroll2FARequest, Enroll2FAResponse, GenerateAPITokenRequest, GenerateAPITokenResponse, GenerateRecoveryCodesRequest, GenerateRecoveryCodesResponse, GetRegisteredPasskeysRequest, GetRegisteredPasskeysResponse, ListAPITokensRequest, ListAPITokensResponse, MarkEmailAsPrimaryRequest, MarkEmailAsPrimaryResponse, MarkPhoneNumberAsPrimaryRequest, MarkPhoneNumberAsPrimaryResponse, Remove2FARequest, Remove2FAResponse, RemoveAPITokenRequest, RemoveAPITokenResponse, RemovePasskeyRequest, RemovePasskeyResponse, UpdateAddressRequest, UpdateAddressResponse, UpdateProfileRequest, UpdateProfileResponse, ValidateEmailRequest, ValidateEmailResponse, ValidatePhoneNumberRequest, ValidatePhoneNumberResponse } from "./self_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service tkd.idm.v1.SelfServiceService
 */
export declare const SelfServiceService: {
  readonly typeName: "tkd.idm.v1.SelfServiceService",
  readonly methods: {
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.ChangePassword
     */
    readonly changePassword: {
      readonly name: "ChangePassword",
      readonly I: typeof ChangePasswordRequest,
      readonly O: typeof ChangePasswordResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.UpdateProfile
     */
    readonly updateProfile: {
      readonly name: "UpdateProfile",
      readonly I: typeof UpdateProfileRequest,
      readonly O: typeof UpdateProfileResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * E-Mails
     *
     * @generated from rpc tkd.idm.v1.SelfServiceService.AddEmailAddress
     */
    readonly addEmailAddress: {
      readonly name: "AddEmailAddress",
      readonly I: typeof AddEmailAddressRequest,
      readonly O: typeof AddEmailAddressResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.DeleteEmailAddress
     */
    readonly deleteEmailAddress: {
      readonly name: "DeleteEmailAddress",
      readonly I: typeof DeleteEmailAddressRequest,
      readonly O: typeof DeleteEmailAddressResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.MarkEmailAsPrimary
     */
    readonly markEmailAsPrimary: {
      readonly name: "MarkEmailAsPrimary",
      readonly I: typeof MarkEmailAsPrimaryRequest,
      readonly O: typeof MarkEmailAsPrimaryResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.ValidateEmail
     */
    readonly validateEmail: {
      readonly name: "ValidateEmail",
      readonly I: typeof ValidateEmailRequest,
      readonly O: typeof ValidateEmailResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Addresses
     *
     * @generated from rpc tkd.idm.v1.SelfServiceService.AddAddress
     */
    readonly addAddress: {
      readonly name: "AddAddress",
      readonly I: typeof AddAddressRequest,
      readonly O: typeof AddAddressResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.DeleteAddress
     */
    readonly deleteAddress: {
      readonly name: "DeleteAddress",
      readonly I: typeof DeleteAddressRequest,
      readonly O: typeof DeleteAddressResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.UpdateAddress
     */
    readonly updateAddress: {
      readonly name: "UpdateAddress",
      readonly I: typeof UpdateAddressRequest,
      readonly O: typeof UpdateAddressResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Phone Numbers
     *
     * @generated from rpc tkd.idm.v1.SelfServiceService.AddPhoneNumber
     */
    readonly addPhoneNumber: {
      readonly name: "AddPhoneNumber",
      readonly I: typeof AddPhoneNumberRequest,
      readonly O: typeof AddPhoneNumberResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.DeletePhoneNumber
     */
    readonly deletePhoneNumber: {
      readonly name: "DeletePhoneNumber",
      readonly I: typeof DeletePhoneNumberRequest,
      readonly O: typeof DeletePhoneNumberResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.MarkPhoneNumberAsPrimary
     */
    readonly markPhoneNumberAsPrimary: {
      readonly name: "MarkPhoneNumberAsPrimary",
      readonly I: typeof MarkPhoneNumberAsPrimaryRequest,
      readonly O: typeof MarkPhoneNumberAsPrimaryResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.ValidatePhoneNumber
     */
    readonly validatePhoneNumber: {
      readonly name: "ValidatePhoneNumber",
      readonly I: typeof ValidatePhoneNumberRequest,
      readonly O: typeof ValidatePhoneNumberResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * WebAuthN
     *
     * @generated from rpc tkd.idm.v1.SelfServiceService.GetRegisteredPasskeys
     */
    readonly getRegisteredPasskeys: {
      readonly name: "GetRegisteredPasskeys",
      readonly I: typeof GetRegisteredPasskeysRequest,
      readonly O: typeof GetRegisteredPasskeysResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.RemovePasskey
     */
    readonly removePasskey: {
      readonly name: "RemovePasskey",
      readonly I: typeof RemovePasskeyRequest,
      readonly O: typeof RemovePasskeyResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Two-Factor authentication
     *
     * @generated from rpc tkd.idm.v1.SelfServiceService.Enroll2FA
     */
    readonly enroll2FA: {
      readonly name: "Enroll2FA",
      readonly I: typeof Enroll2FARequest,
      readonly O: typeof Enroll2FAResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.Remove2FA
     */
    readonly remove2FA: {
      readonly name: "Remove2FA",
      readonly I: typeof Remove2FARequest,
      readonly O: typeof Remove2FAResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.GenerateRecoveryCodes
     */
    readonly generateRecoveryCodes: {
      readonly name: "GenerateRecoveryCodes",
      readonly I: typeof GenerateRecoveryCodesRequest,
      readonly O: typeof GenerateRecoveryCodesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.GenerateAPIToken
     */
    readonly generateAPIToken: {
      readonly name: "GenerateAPIToken",
      readonly I: typeof GenerateAPITokenRequest,
      readonly O: typeof GenerateAPITokenResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.ListAPITokens
     */
    readonly listAPITokens: {
      readonly name: "ListAPITokens",
      readonly I: typeof ListAPITokensRequest,
      readonly O: typeof ListAPITokensResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.SelfServiceService.RemoveAPIToken
     */
    readonly removeAPIToken: {
      readonly name: "RemoveAPIToken",
      readonly I: typeof RemoveAPITokenRequest,
      readonly O: typeof RemoveAPITokenResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

