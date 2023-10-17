// @generated by protoc-gen-es v1.3.3 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/user.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FieldMask, proto3, Struct } from "@bufbuild/protobuf";
import { Role } from "./role_service_pb.js";

/**
 * @generated from message tkd.idm.v1.EMail
 */
export const EMail = proto3.makeMessageType(
  "tkd.idm.v1.EMail",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "verified", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "primary", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.PhoneNumber
 */
export const PhoneNumber = proto3.makeMessageType(
  "tkd.idm.v1.PhoneNumber",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "verified", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "primary", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.Address
 */
export const Address = proto3.makeMessageType(
  "tkd.idm.v1.Address",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "city_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "city_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "street", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "extra", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.User
 */
export const User = proto3.makeMessageType(
  "tkd.idm.v1.User",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "username", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "first_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "last_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "extra", kind: "message", T: Struct },
    { no: 10, name: "avatar", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "birthday", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "primary_mail", kind: "message", T: EMail },
    { no: 13, name: "primary_phone_number", kind: "message", T: PhoneNumber },
  ],
);

/**
 * @generated from message tkd.idm.v1.Profile
 */
export const Profile = proto3.makeMessageType(
  "tkd.idm.v1.Profile",
  () => [
    { no: 1, name: "user", kind: "message", T: User },
    { no: 6, name: "addresses", kind: "message", T: Address, repeated: true },
    { no: 7, name: "phone_numbers", kind: "message", T: PhoneNumber, repeated: true },
    { no: 8, name: "email_addresses", kind: "message", T: EMail, repeated: true },
    { no: 9, name: "totp_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 10, name: "recovery_codes_generated", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 11, name: "roles", kind: "message", T: Role, repeated: true },
    { no: 12, name: "password_auth_enabled", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 20, name: "privacy_mask", kind: "message", T: FieldMask },
  ],
);

