// @generated by protoc-gen-es v1.4.1 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/notify_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Struct } from "@bufbuild/protobuf";

/**
 * @generated from enum tkd.idm.v1.AttachmentType
 */
export const AttachmentType = proto3.makeEnum(
  "tkd.idm.v1.AttachmentType",
  [
    {no: 0, name: "UNSPECIFIED"},
    {no: 1, name: "INLINE"},
    {no: 2, name: "ATTACHEMNT"},
    {no: 3, name: "ALTERNATIVE_BODY"},
  ],
);

/**
 * @generated from message tkd.idm.v1.Attachment
 */
export const Attachment = proto3.makeMessageType(
  "tkd.idm.v1.Attachment",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "media_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "content", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 4, name: "content_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "attachment_type", kind: "enum", T: proto3.getEnumType(AttachmentType) },
    { no: 6, name: "for_user", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.EMailMessage
 */
export const EMailMessage = proto3.makeMessageType(
  "tkd.idm.v1.EMailMessage",
  () => [
    { no: 1, name: "subject", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "attachments", kind: "message", T: Attachment, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.SMS
 */
export const SMS = proto3.makeMessageType(
  "tkd.idm.v1.SMS",
  () => [
    { no: 1, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.SendNotificationRequest
 */
export const SendNotificationRequest = proto3.makeMessageType(
  "tkd.idm.v1.SendNotificationRequest",
  () => [
    { no: 1, name: "sms", kind: "message", T: SMS, oneof: "message" },
    { no: 2, name: "email", kind: "message", T: EMailMessage, oneof: "message" },
    { no: 3, name: "target_users", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "target_roles", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "per_user_template_context", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Struct} },
    { no: 6, name: "sender_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeliveryNotification
 */
export const DeliveryNotification = proto3.makeMessageType(
  "tkd.idm.v1.DeliveryNotification",
  () => [
    { no: 1, name: "target_user", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "error", kind: "scalar", T: 9 /* ScalarType.STRING */ },
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

