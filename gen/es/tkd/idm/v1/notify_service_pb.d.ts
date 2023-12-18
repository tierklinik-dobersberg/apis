// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/notify_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Struct } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum tkd.idm.v1.AttachmentType
 */
export declare enum AttachmentType {
  /**
   * @generated from enum value: UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: INLINE = 1;
   */
  INLINE = 1,

  /**
   * @generated from enum value: ATTACHEMNT = 2;
   */
  ATTACHEMNT = 2,

  /**
   * @generated from enum value: ALTERNATIVE_BODY = 3;
   */
  ALTERNATIVE_BODY = 3,
}

/**
 * @generated from message tkd.idm.v1.Attachment
 */
export declare class Attachment extends Message<Attachment> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string media_type = 2;
   */
  mediaType: string;

  /**
   * @generated from field: bytes content = 3;
   */
  content: Uint8Array;

  /**
   * @generated from field: string content_id = 4;
   */
  contentId: string;

  /**
   * @generated from field: tkd.idm.v1.AttachmentType attachment_type = 5;
   */
  attachmentType: AttachmentType;

  /**
   * A list of user IDs this attachment is for.
   * if left empty, the attachment will be sent to
   * all users.
   *
   * @generated from field: repeated string for_user = 6;
   */
  forUser: string[];

  constructor(data?: PartialMessage<Attachment>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.Attachment";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Attachment;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Attachment;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Attachment;

  static equals(a: Attachment | PlainMessage<Attachment> | undefined, b: Attachment | PlainMessage<Attachment> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.EMailMessage
 */
export declare class EMailMessage extends Message<EMailMessage> {
  /**
   * @generated from field: string subject = 1;
   */
  subject: string;

  /**
   * @generated from field: string body = 2;
   */
  body: string;

  /**
   * @generated from field: repeated tkd.idm.v1.Attachment attachments = 3;
   */
  attachments: Attachment[];

  constructor(data?: PartialMessage<EMailMessage>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.EMailMessage";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EMailMessage;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EMailMessage;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EMailMessage;

  static equals(a: EMailMessage | PlainMessage<EMailMessage> | undefined, b: EMailMessage | PlainMessage<EMailMessage> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.SMS
 */
export declare class SMS extends Message<SMS> {
  /**
   * @generated from field: string body = 1;
   */
  body: string;

  constructor(data?: PartialMessage<SMS>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.SMS";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SMS;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SMS;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SMS;

  static equals(a: SMS | PlainMessage<SMS> | undefined, b: SMS | PlainMessage<SMS> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.SendNotificationRequest
 */
export declare class SendNotificationRequest extends Message<SendNotificationRequest> {
  /**
   * @generated from oneof tkd.idm.v1.SendNotificationRequest.message
   */
  message: {
    /**
     * @generated from field: tkd.idm.v1.SMS sms = 1;
     */
    value: SMS;
    case: "sms";
  } | {
    /**
     * @generated from field: tkd.idm.v1.EMailMessage email = 2;
     */
    value: EMailMessage;
    case: "email";
  } | { case: undefined; value?: undefined };

  /**
   * @generated from field: repeated string target_users = 3;
   */
  targetUsers: string[];

  /**
   * @generated from field: repeated string target_roles = 4;
   */
  targetRoles: string[];

  /**
   * @generated from field: map<string, google.protobuf.Struct> per_user_template_context = 5;
   */
  perUserTemplateContext: { [key: string]: Struct };

  /**
   * @generated from field: string sender_user_id = 6;
   */
  senderUserId: string;

  constructor(data?: PartialMessage<SendNotificationRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.SendNotificationRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendNotificationRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendNotificationRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendNotificationRequest;

  static equals(a: SendNotificationRequest | PlainMessage<SendNotificationRequest> | undefined, b: SendNotificationRequest | PlainMessage<SendNotificationRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.DeliveryNotification
 */
export declare class DeliveryNotification extends Message<DeliveryNotification> {
  /**
   * @generated from field: string target_user = 1;
   */
  targetUser: string;

  /**
   * @generated from field: string error = 2;
   */
  error: string;

  constructor(data?: PartialMessage<DeliveryNotification>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.DeliveryNotification";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeliveryNotification;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeliveryNotification;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeliveryNotification;

  static equals(a: DeliveryNotification | PlainMessage<DeliveryNotification> | undefined, b: DeliveryNotification | PlainMessage<DeliveryNotification> | undefined): boolean;
}

/**
 * @generated from message tkd.idm.v1.SendNotificationResponse
 */
export declare class SendNotificationResponse extends Message<SendNotificationResponse> {
  /**
   * @generated from field: repeated tkd.idm.v1.DeliveryNotification deliveries = 1;
   */
  deliveries: DeliveryNotification[];

  constructor(data?: PartialMessage<SendNotificationResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.idm.v1.SendNotificationResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SendNotificationResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SendNotificationResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SendNotificationResponse;

  static equals(a: SendNotificationResponse | PlainMessage<SendNotificationResponse> | undefined, b: SendNotificationResponse | PlainMessage<SendNotificationResponse> | undefined): boolean;
}

