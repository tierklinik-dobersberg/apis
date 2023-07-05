// @generated by protoc-gen-es v1.2.1 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/user.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Value } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message tkd.idm.v1.EMail
 */
export declare class EMail extends Message<EMail> {
  /**
   * The unique ID of the email address record required for validation
   * editing and deletion of the records.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The actual email address.
   *
   * @generated from field: string address = 2;
   */
  address: string;

  /**
   * Whether or not the email has been verified by the user.
   * This field may only be set by using the SelfServiceService
   * or by an administrator using the UserService.
   *
   * @generated from field: bool verified = 3;
   */
  verified: boolean;

  /**
   * Whether or not this email address is the user's primary e-mail.
   * The primary e-mail address is also part of the User message so
   * users may choose to allow public access to their primary e-mail address
   * while keeping other e-mail records private.
   *
   * @generated from field: bool primary = 4;
   */
  primary: boolean;

  constructor(data?: PartialMessage<EMail>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.EMail";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EMail;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EMail;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EMail;

  static equals(a: EMail | PlainMessage<EMail> | undefined, b: EMail | PlainMessage<EMail> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.PhoneNumber
 */
export declare class PhoneNumber extends Message<PhoneNumber> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string number = 2;
   */
  number: string;

  /**
   * @generated from field: bool verified = 3;
   */
  verified: boolean;

  /**
   * @generated from field: bool primary = 4;
   */
  primary: boolean;

  constructor(data?: PartialMessage<PhoneNumber>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.PhoneNumber";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PhoneNumber;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PhoneNumber;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PhoneNumber;

  static equals(a: PhoneNumber | PlainMessage<PhoneNumber> | undefined, b: PhoneNumber | PlainMessage<PhoneNumber> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.Address
 */
export declare class Address extends Message<Address> {
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

  constructor(data?: PartialMessage<Address>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.Address";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Address;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Address;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Address;

  static equals(a: Address | PlainMessage<Address> | undefined, b: Address | PlainMessage<Address> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.User
 */
export declare class User extends Message<User> {
  /**
   * The primary and unique identifier of the user. This is normally a generated UUIDv4
   * but may be set to any arbitrary string. Applications should treat the user ID is an
   * opaque string and avoid any assumptions on the format.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The name of the user for login purposes. Depending on the CISIDM feature set this
   * field may be changed by the user.
   *
   * @generated from field: string username = 2;
   */
  username: string;

  /**
   * The name of the user as it should be displayed in all applications. If unset, applications should
   * fallback to the username. 
   *
   * @generated from field: string display_name = 3;
   */
  displayName: string;

  /**
   * The first name of the user.
   *
   * @generated from field: string first_name = 4;
   */
  firstName: string;

  /**
   * The last name of the user.
   *
   * @generated from field: string last_name = 5;
   */
  lastName: string;

  /**
   * Additional data for the user. The available keys in the extra map and their meaning is application
   * specific. CISIDM may provide validation and schema definitions for the user-extra data any some point
   * in the future. For now, only administrators may set extra data on user profiles.
   *
   * @generated from field: map<string, google.protobuf.Value> extra = 9;
   */
  extra: { [key: string]: Value };

  /**
   * The URL (or data-URL) of the user avatar.
   *
   * @generated from field: string avatar = 10;
   */
  avatar: string;

  /**
   * The birthday of the user. Should follow the format YYYY-MM-DD. Applications may choose to
   * display the birthday of a user in a different format depending on the viewer's locale.
   *
   * @generated from field: string birthday = 11;
   */
  birthday: string;

  /**
   * The primary email address of the user.
   *
   * @generated from field: tkd.idm.v1.EMail primary_mail = 12;
   */
  primaryMail?: EMail;

  /**
   * The primary phone number of the user.
   *
   * @generated from field: tkd.idm.v1.PhoneNumber primary_phone_number = 13;
   */
  primaryPhoneNumber?: PhoneNumber;

  constructor(data?: PartialMessage<User>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.User";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): User;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): User;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): User;

  static equals(a: User | PlainMessage<User> | undefined, b: User | PlainMessage<User> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.Profile
 */
export declare class Profile extends Message<Profile> {
  /**
   * The actual user object.
   *
   * @generated from field: tkd.idm.v1.User user = 1;
   */
  user?: User;

  /**
   * A list of all addresses the user has configured.
   * This may be set using the SelfServiceService. Administrators
   * may change these values using the UserService.
   *
   * @generated from field: repeated tkd.idm.v1.Address addresses = 6;
   */
  addresses: Address[];

  /**
   * A list of all phone numbers configured by the user.
   * This may be set using the SelfServiceService. Administrators
   * may change these values using the UserService.
   *
   * @generated from field: repeated tkd.idm.v1.PhoneNumber phone_numbers = 7;
   */
  phoneNumbers: PhoneNumber[];

  /**
   * A list of all email addresses configured by the user.
   * This may be set using the SelfServiceService. Administrators
   * may change these values using the UserService.
   *
   * @generated from field: repeated tkd.idm.v1.EMail email_addresses = 8;
   */
  emailAddresses: EMail[];

  /**
   * Whether or not this user has TOTP enabled.
   *
   * @generated from field: bool totp_enabled = 9;
   */
  totpEnabled: boolean;

  /**
   * Whether or not there are MFA recovery codes for this user.
   *
   * @generated from field: bool recovery_codes_generated = 10;
   */
  recoveryCodesGenerated: boolean;

  /**
   * A field-mask to defined publically viewable fields from the
   * user profile. If set, the configured field-mask overwrites
   * the field mask of the (tkd.common.v1.readable) option.
   * Note that the privacy_mask is not applied to users that either
   * match the owner_field_name or allowed_roles field of the (tkd.common.v1.readable)
   * option.
   *
   * @generated from field: google.protobuf.FieldMask privacy_mask = 20;
   */
  privacyMask?: FieldMask;

  constructor(data?: PartialMessage<Profile>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.Profile";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Profile;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Profile;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Profile;

  static equals(a: Profile | PlainMessage<Profile> | undefined, b: Profile | PlainMessage<Profile> | undefined): boolean;
}

