// @generated by protoc-gen-es v1.3.3 with parameter "target=js+dts"
// @generated from file tkd/calendar/v1/event.proto (package tkd.calendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { Any, BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * Calendar describes a event calendar.
 *
 * @generated from message tkd.calendar.v1.Calendar
 */
export declare class Calendar extends Message<Calendar> {
  /**
   * ID is a unique ID for the calendar. The format of the ID may depend
   * on the application and the actual calendar backend.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * Name is a human friendly name of the calendar.
   *
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * Timezone holds the timezone information for the calendar.
   * In the format of Europe/Vienna.
   *
   * @generated from field: string timezone = 3;
   */
  timezone: string;

  constructor(data?: PartialMessage<Calendar>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.Calendar";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Calendar;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Calendar;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Calendar;

  static equals(a: Calendar | PlainMessage<Calendar> | undefined, b: Calendar | PlainMessage<Calendar> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.CalendarEvent
 */
export declare class CalendarEvent extends Message<CalendarEvent> {
  /**
   * ID is a unique ID for the calendar event.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * CalendarId is the unique ID of the calendar this event
   * belongs to.
   *
   * @generated from field: string calendar_id = 2;
   */
  calendarId: string;

  /**
   * StartTime holds the time the event begins.
   *
   * @generated from field: google.protobuf.Timestamp start_time = 3;
   */
  startTime?: Timestamp;

  /**
   * EndTime optionally holds the time the event ends.
   *
   * @generated from field: google.protobuf.Timestamp end_time = 4;
   */
  endTime?: Timestamp;

  /**
   * FullDay may be set to true for events the last the whole day.
   * In this case, the time part (HH:MM:SS) of StartTime is not important!
   *
   * @generated from field: bool full_day = 5;
   */
  fullDay: boolean;

  /**
   * Summary of the calendar event.
   *
   * @generated from field: string summary = 6;
   */
  summary: string;

  /**
   * Description of the calendar event.
   *
   * @generated from field: string description = 7;
   */
  description: string;

  /**
   * ExtraData may holds additional information about the calendar entry.
   * In most cases, this should be CustomerAnnotation
   *
   * @generated from field: google.protobuf.Any extra_data = 8;
   */
  extraData?: Any;

  constructor(data?: PartialMessage<CalendarEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.CalendarEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CalendarEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CalendarEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CalendarEvent;

  static equals(a: CalendarEvent | PlainMessage<CalendarEvent> | undefined, b: CalendarEvent | PlainMessage<CalendarEvent> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.CustomerAnnotation
 */
export declare class CustomerAnnotation extends Message<CustomerAnnotation> {
  /**
   * CustomerSource may be set to the source of the customer record.
   *
   * @generated from field: string customer_source = 1;
   */
  customerSource: string;

  /**
   * CustomerId is the ID of the customer within the specified source.
   *
   * @generated from field: string customer_id = 2;
   */
  customerId: string;

  /**
   * AnimalIds is a list of animals that are expected to show up during
   * the event and need treatment.
   *
   * @generated from field: repeated string animal_ids = 3;
   */
  animalIds: string[];

  /**
   * CreatedByUserId holds the ID of the user that created the event.
   *
   * @generated from field: string created_by_user_id = 4;
   */
  createdByUserId: string;

  constructor(data?: PartialMessage<CustomerAnnotation>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.CustomerAnnotation";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CustomerAnnotation;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CustomerAnnotation;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CustomerAnnotation;

  static equals(a: CustomerAnnotation | PlainMessage<CustomerAnnotation> | undefined, b: CustomerAnnotation | PlainMessage<CustomerAnnotation> | undefined): boolean;
}

