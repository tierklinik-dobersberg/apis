// @generated by protoc-gen-es v1.3.0 with parameter "target=js+dts"
// @generated from file tkd/calendar/v1/event_service.proto (package tkd.calendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { Any, BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Calendar, CalendarEvent } from "./event_pb.js";

/**
 * @generated from message tkd.calendar.v1.ListCalendarsRequest
 */
export declare class ListCalendarsRequest extends Message<ListCalendarsRequest> {
  constructor(data?: PartialMessage<ListCalendarsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.ListCalendarsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListCalendarsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListCalendarsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListCalendarsRequest;

  static equals(a: ListCalendarsRequest | PlainMessage<ListCalendarsRequest> | undefined, b: ListCalendarsRequest | PlainMessage<ListCalendarsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.ListCalendarsResponse
 */
export declare class ListCalendarsResponse extends Message<ListCalendarsResponse> {
  /**
   * @generated from field: repeated tkd.calendar.v1.Calendar calendars = 1;
   */
  calendars: Calendar[];

  constructor(data?: PartialMessage<ListCalendarsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.ListCalendarsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListCalendarsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListCalendarsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListCalendarsResponse;

  static equals(a: ListCalendarsResponse | PlainMessage<ListCalendarsResponse> | undefined, b: ListCalendarsResponse | PlainMessage<ListCalendarsResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.TimeRange
 */
export declare class TimeRange extends Message<TimeRange> {
  /**
   * @generated from field: google.protobuf.Timestamp from = 1;
   */
  from?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp to = 2;
   */
  to?: Timestamp;

  constructor(data?: PartialMessage<TimeRange>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.TimeRange";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TimeRange;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TimeRange;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TimeRange;

  static equals(a: TimeRange | PlainMessage<TimeRange> | undefined, b: TimeRange | PlainMessage<TimeRange> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.EventSource
 */
export declare class EventSource extends Message<EventSource> {
  /**
   * @generated from field: repeated string calendar_ids = 1;
   */
  calendarIds: string[];

  /**
   * @generated from field: repeated string user_ids = 2;
   */
  userIds: string[];

  constructor(data?: PartialMessage<EventSource>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.EventSource";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): EventSource;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): EventSource;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): EventSource;

  static equals(a: EventSource | PlainMessage<EventSource> | undefined, b: EventSource | PlainMessage<EventSource> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.ListEventsRequest
 */
export declare class ListEventsRequest extends Message<ListEventsRequest> {
  /**
   * @generated from field: tkd.calendar.v1.EventSource sources = 1;
   */
  sources?: EventSource;

  /**
   * @generated from oneof tkd.calendar.v1.ListEventsRequest.search_time
   */
  searchTime: {
    /**
     * @generated from field: tkd.calendar.v1.TimeRange time_range = 2;
     */
    value: TimeRange;
    case: "timeRange";
  } | {
    /**
     * Format is YYYY/MM/DD
     *
     * @generated from field: string date = 4;
     */
    value: string;
    case: "date";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<ListEventsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.ListEventsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEventsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEventsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEventsRequest;

  static equals(a: ListEventsRequest | PlainMessage<ListEventsRequest> | undefined, b: ListEventsRequest | PlainMessage<ListEventsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.CalendarEventList
 */
export declare class CalendarEventList extends Message<CalendarEventList> {
  /**
   * @generated from field: string calendar_id = 1;
   */
  calendarId: string;

  /**
   * @generated from field: repeated tkd.calendar.v1.CalendarEvent events = 2;
   */
  events: CalendarEvent[];

  constructor(data?: PartialMessage<CalendarEventList>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.CalendarEventList";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CalendarEventList;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CalendarEventList;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CalendarEventList;

  static equals(a: CalendarEventList | PlainMessage<CalendarEventList> | undefined, b: CalendarEventList | PlainMessage<CalendarEventList> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.ListEventsResponse
 */
export declare class ListEventsResponse extends Message<ListEventsResponse> {
  /**
   * @generated from field: repeated tkd.calendar.v1.CalendarEventList calendar_events = 1;
   */
  calendarEvents: CalendarEventList[];

  constructor(data?: PartialMessage<ListEventsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.ListEventsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEventsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEventsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEventsResponse;

  static equals(a: ListEventsResponse | PlainMessage<ListEventsResponse> | undefined, b: ListEventsResponse | PlainMessage<ListEventsResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.CreateEventRequest
 */
export declare class CreateEventRequest extends Message<CreateEventRequest> {
  /**
   * @generated from field: string calendar_id = 1;
   */
  calendarId: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: string description = 3;
   */
  description: string;

  /**
   * @generated from field: google.protobuf.Timestamp start = 4;
   */
  start?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp end = 5;
   */
  end?: Timestamp;

  /**
   * @generated from field: google.protobuf.Any extra_data = 6;
   */
  extraData?: Any;

  constructor(data?: PartialMessage<CreateEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.CreateEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateEventRequest;

  static equals(a: CreateEventRequest | PlainMessage<CreateEventRequest> | undefined, b: CreateEventRequest | PlainMessage<CreateEventRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.CreateEventResponse
 */
export declare class CreateEventResponse extends Message<CreateEventResponse> {
  /**
   * @generated from field: tkd.calendar.v1.CalendarEvent event = 1;
   */
  event?: CalendarEvent;

  constructor(data?: PartialMessage<CreateEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.CreateEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateEventResponse;

  static equals(a: CreateEventResponse | PlainMessage<CreateEventResponse> | undefined, b: CreateEventResponse | PlainMessage<CreateEventResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.DeleteEventRequest
 */
export declare class DeleteEventRequest extends Message<DeleteEventRequest> {
  /**
   * @generated from field: string calendar_id = 1;
   */
  calendarId: string;

  /**
   * @generated from field: string event_id = 2;
   */
  eventId: string;

  constructor(data?: PartialMessage<DeleteEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.DeleteEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteEventRequest;

  static equals(a: DeleteEventRequest | PlainMessage<DeleteEventRequest> | undefined, b: DeleteEventRequest | PlainMessage<DeleteEventRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.DeleteEventResponse
 */
export declare class DeleteEventResponse extends Message<DeleteEventResponse> {
  constructor(data?: PartialMessage<DeleteEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.DeleteEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteEventResponse;

  static equals(a: DeleteEventResponse | PlainMessage<DeleteEventResponse> | undefined, b: DeleteEventResponse | PlainMessage<DeleteEventResponse> | undefined): boolean;
}

