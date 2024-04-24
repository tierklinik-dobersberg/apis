// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/notify_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, proto3, Struct } from "@bufbuild/protobuf";

/**
 * @generated from enum tkd.idm.v1.AttachmentType
 */
export const AttachmentType = /*@__PURE__*/ proto3.makeEnum(
  "tkd.idm.v1.AttachmentType",
  [
    {no: 0, name: "UNSPECIFIED"},
    {no: 1, name: "INLINE"},
    {no: 2, name: "ATTACHEMNT"},
    {no: 3, name: "ALTERNATIVE_BODY"},
  ],
);

/**
 * @generated from enum tkd.idm.v1.Operation
 */
export const Operation = /*@__PURE__*/ proto3.makeEnum(
  "tkd.idm.v1.Operation",
  [
    {no: 0, name: "OPERATION_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "OPERATION_OPEN_WINDOW", localName: "OPEN_WINDOW"},
    {no: 2, name: "OPERATION_FOCUS_LAST_FOCUSED_OR_OPEN", localName: "FOCUS_LAST_FOCUSED_OR_OPEN"},
    {no: 3, name: "OPERATION_NAVIGATE_LAST_FOCUSED_OR_OPEN", localName: "NAVIGATE_LAST_FOCUSED_OR_OPEN"},
    {no: 4, name: "OPERATION_SEND_REQUEST", localName: "SEND_REQUEST"},
  ],
);

/**
 * @generated from enum tkd.idm.v1.ErrorKind
 */
export const ErrorKind = /*@__PURE__*/ proto3.makeEnum(
  "tkd.idm.v1.ErrorKind",
  [
    {no: 0, name: "ERROR_KIND_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "ERROR_KIND_NO_PRIMARY_MAIL", localName: "NO_PRIMARY_MAIL"},
    {no: 2, name: "ERROR_KIND_NO_PRIMARY_PHONE", localName: "NO_PRIMARY_PHONE"},
    {no: 3, name: "ERROR_KIND_NO_WEBPUSH_SUBSCRIPTION", localName: "NO_WEBPUSH_SUBSCRIPTION"},
    {no: 4, name: "ERROR_KIND_TRANSPORT", localName: "TRANSPORT"},
    {no: 5, name: "ERROR_KIND_TEMPLATE", localName: "TEMPLATE"},
    {no: 255, name: "ERROR_KIND_OTHER", localName: "OTHER"},
  ],
);

/**
 * @generated from message tkd.idm.v1.Attachment
 */
export const Attachment = /*@__PURE__*/ proto3.makeMessageType(
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
export const EMailMessage = /*@__PURE__*/ proto3.makeMessageType(
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
export const SMS = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.SMS",
  () => [
    { no: 1, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.NotificationAction
 */
export const NotificationAction = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.NotificationAction",
  () => [
    { no: 1, name: "action", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "operation", kind: "enum", T: proto3.getEnumType(Operation) },
    { no: 4, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.ServiceWorkerNotification
 */
export const ServiceWorkerNotification = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.ServiceWorkerNotification",
  () => [
    { no: 1, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "default_operation", kind: "enum", T: proto3.getEnumType(Operation) },
    { no: 4, name: "default_operation_url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "actions", kind: "message", T: NotificationAction, repeated: true },
    { no: 6, name: "data", kind: "message", T: Struct },
  ],
);

/**
 * @generated from message tkd.idm.v1.WebPushNotification
 */
export const WebPushNotification = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.WebPushNotification",
  () => [
    { no: 1, name: "binary", kind: "scalar", T: 12 /* ScalarType.BYTES */, oneof: "kind" },
    { no: 2, name: "notification", kind: "message", T: ServiceWorkerNotification, oneof: "kind" },
  ],
);

/**
 * @generated from message tkd.idm.v1.SendNotificationRequest
 */
export const SendNotificationRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.SendNotificationRequest",
  () => [
    { no: 1, name: "sms", kind: "message", T: SMS, oneof: "message" },
    { no: 2, name: "email", kind: "message", T: EMailMessage, oneof: "message" },
    { no: 7, name: "webpush", kind: "message", T: WebPushNotification, oneof: "message" },
    { no: 3, name: "target_users", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "target_roles", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "per_user_template_context", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Struct} },
    { no: 6, name: "sender_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.DeliveryNotification
 */
export const DeliveryNotification = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.DeliveryNotification",
  () => [
    { no: 1, name: "target_user", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "error", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "error_kind", kind: "enum", T: proto3.getEnumType(ErrorKind) },
  ],
);

/**
 * @generated from message tkd.idm.v1.SendNotificationResponse
 */
export const SendNotificationResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.SendNotificationResponse",
  () => [
    { no: 1, name: "deliveries", kind: "message", T: DeliveryNotification, repeated: true },
  ],
);

/**
 * @generated from message tkd.idm.v1.GetVAPIDPublicKeyRequest
 */
export const GetVAPIDPublicKeyRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.GetVAPIDPublicKeyRequest",
  [],
);

/**
 * @generated from message tkd.idm.v1.GetVAPIDPublicKeyResponse
 */
export const GetVAPIDPublicKeyResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.GetVAPIDPublicKeyResponse",
  () => [
    { no: 1, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.WebPushSubscription
 */
export const WebPushSubscription = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.WebPushSubscription",
  () => [
    { no: 1, name: "endpoint", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "expiration_type", kind: "message", T: Duration },
    { no: 3, name: "keys", kind: "message", T: WebPushKeys },
  ],
);

/**
 * @generated from message tkd.idm.v1.WebPushKeys
 */
export const WebPushKeys = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.WebPushKeys",
  () => [
    { no: 1, name: "auth", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "p256dh", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.idm.v1.AddWebPushSubscriptionRequest
 */
export const AddWebPushSubscriptionRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.AddWebPushSubscriptionRequest",
  () => [
    { no: 1, name: "subscription", kind: "message", T: WebPushSubscription },
  ],
);

/**
 * @generated from message tkd.idm.v1.AddWebPushSubscriptionResponse
 */
export const AddWebPushSubscriptionResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.idm.v1.AddWebPushSubscriptionResponse",
  [],
);

