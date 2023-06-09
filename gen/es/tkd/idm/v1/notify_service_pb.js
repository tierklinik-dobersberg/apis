// @generated by protoc-gen-es v1.2.1 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/notify_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum tkd.idm.v1.Channel
 */
export const Channel = proto3.makeEnum(
  "tkd.idm.v1.Channel",
  [
    {no: 0, name: "CHANNEL_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "CHANNEL_SMS", localName: "SMS"},
    {no: 2, name: "CHANNEL_MAIL", localName: "MAIL"},
  ],
);

/**
 * @generated from message tkd.idm.v1.SendNotificationRequest
 */
export const SendNotificationRequest = proto3.makeMessageType(
  "tkd.idm.v1.SendNotificationRequest",
  () => [
    { no: 1, name: "text", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "message" },
    { no: 2, name: "template_name", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "message" },
    { no: 3, name: "target_users", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "target_roles", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "preferred_channel", kind: "enum", T: proto3.getEnumType(Channel) },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeliveryNotification
 */
export const DeliveryNotification = proto3.makeMessageType(
  "tkd.idm.v1.DeliveryNotification",
  () => [
    { no: 1, name: "target_user", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "used_channel", kind: "enum", T: proto3.getEnumType(Channel) },
  ],
);

/**
 * @generated from message tkd.idm.v1.SendNotificationResponse
 */
export const SendNotificationResponse = proto3.makeMessageType(
  "tkd.idm.v1.SendNotificationResponse",
  () => [
    { no: 1, name: "deliveries", kind: "message", T: DeliveryNotification, repeated: true },
  ],
);

