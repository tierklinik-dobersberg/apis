// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file tkd/calendar/v1/event_service.proto (package tkd.calendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { Any, BinaryReadOptions, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Calendar, CalendarEvent } from "./event_pb.js";
import type { TimeRange } from "../../common/v1/time_range_pb.js";

/**
 * ListCalendarsRequest is used by the ListCalendars RPC. There are now message
 * fields for now.
 *
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
 * ListCalendarsResponse is the response of the ListCalendars RPC and contains a
 * list of available calendars.
 *
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
 * EventSource is used to query different event sources during the ListEvents
 * RPC:
 *
 * @generated from message tkd.calendar.v1.EventSource
 */
export declare class EventSource extends Message<EventSource> {
  /**
   * CalendarIds is a list of calendar ids that should be searched for events.
   *
   * @generated from field: repeated string calendar_ids = 1;
   */
  calendarIds: string[];

  /**
   * UserIds is a list of user IDs for which calendar event should be loaded.
   * When specified, cis-cal will lookup the users calendar id from cis-idms
   * additional user fields as specified in the configuration.
   *
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
 * ListEventsRequest is the request message for the ListEvents RPC and supports
 * querying calendar events over a list of sources.
 *
 * @generated from message tkd.calendar.v1.ListEventsRequest
 */
export declare class ListEventsRequest extends Message<ListEventsRequest> {
  /**
   * Source describes where cis-cal should look for calendar events. At least
   * on of the oneof members must be set.
   *
   * @generated from oneof tkd.calendar.v1.ListEventsRequest.source
   */
  source: {
    /**
     * Sources can be used to specify a dedicated list of user and/or
     * calendar ids.
     *
     * @generated from field: tkd.calendar.v1.EventSource sources = 1;
     */
    value: EventSource;
    case: "sources";
  } | {
    /**
     * AllCalendars might be set to true to query all available calendars.
     * To get a list of available calendars, user the ListCalendars RPC.
     *
     * @generated from field: bool all_calendars = 2;
     */
    value: boolean;
    case: "allCalendars";
  } | {
    /**
     * AllUsers might be set to true to query events from all user
     * calendars.
     *
     * @generated from field: bool all_users = 3;
     */
    value: boolean;
    case: "allUsers";
  } | { case: undefined; value?: undefined };

  /**
   * SearchTime allows to specify a search time for calendar events. At least
   * one of the oneof member fields must be set.
   *
   * @generated from oneof tkd.calendar.v1.ListEventsRequest.search_time
   */
  searchTime: {
    /**
     * TimeRange might be set to a start and end time. Only events that
     * start/end within the specified time range (inclusive) will be
     * returned.
     *
     * @generated from field: tkd.common.v1.TimeRange time_range = 4;
     */
    value: TimeRange;
    case: "timeRange";
  } | {
    /**
     * Date might be set to a date in the format of YYYY-MM-DD  or
     * YYYY/MM/DD. Only events that start OR end at the specified date will
     * be returned.
     *
     * @generated from field: string date = 5;
     */
    value: string;
    case: "date";
  } | { case: undefined; value?: undefined };

  /**
   * ReadMask can be used to limit the number of fields returned in the
   * response.
   *
   * @generated from field: google.protobuf.FieldMask read_mask = 6;
   */
  readMask?: FieldMask;

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
 * CalendarEventList holds a list of events along the calendar definition. Use
 * ReadMask from ListEventsRequest if not all fields are required in the
 * response.
 *
 * @generated from message tkd.calendar.v1.CalendarEventList
 */
export declare class CalendarEventList extends Message<CalendarEventList> {
  /**
   * Calendar holds the calendar definition.
   *
   * @generated from field: tkd.calendar.v1.Calendar calendar = 1;
   */
  calendar?: Calendar;

  /**
   * Events holds a list of calendar events that matched the search query for
   * this calendar.
   *
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
 * ListEventsResponse is the response of the ListEvents RPC and contains a list
 * of CalendarEventList messages that contains the calendar definition as well
 * as the list of events that matched the search query.
 *
 * @generated from message tkd.calendar.v1.ListEventsResponse
 */
export declare class ListEventsResponse extends Message<ListEventsResponse> {
  /**
   * Results is a list of CalendarEventList containing events that matched the
   * search query.
   *
   * @generated from field: repeated tkd.calendar.v1.CalendarEventList results = 1;
   */
  results: CalendarEventList[];

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
 * CreateEventRequest is used to create a new calendar event in the specified
 * calendar.
 *
 * @generated from message tkd.calendar.v1.CreateEventRequest
 */
export declare class CreateEventRequest extends Message<CreateEventRequest> {
  /**
   * CalendarId is the ID of the calendar where the new event should be
   * created.
   *
   * @generated from field: string calendar_id = 1;
   */
  calendarId: string;

  /**
   * Name is the name or summary of the calendar event. This field is
   * required.
   *
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * Description may hold an optional description of the calendar event.
   *
   * @generated from field: string description = 3;
   */
  description: string;

  /**
   * Start holds the start-time of the calendar event (inclusive).
   *
   * @generated from field: google.protobuf.Timestamp start = 4;
   */
  start?: Timestamp;

  /**
   * End holds the end-time of the calendar-event (inclusive).
   *
   * @generated from field: google.protobuf.Timestamp end = 5;
   */
  end?: Timestamp;

  /**
   * ExtraData might be set to an abritrary protobuf message that should be
   * attached to the calendar event.
   *
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
 * CreateEventResponse is the response message of the CreateEvent RPC and
 * contains the created event.
 *
 * @generated from message tkd.calendar.v1.CreateEventResponse
 */
export declare class CreateEventResponse extends Message<CreateEventResponse> {
  /**
   * Event holds the newly created calendar event.
   *
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
 * UpdateEventRequest is the request message for the UpdateEvent RPC and might
 * be used to partitially update a calendar event.
 *
 * @generated from message tkd.calendar.v1.UpdateEventRequest
 */
export declare class UpdateEventRequest extends Message<UpdateEventRequest> {
  /**
   * CalendarId is the id of the calendar that contains the event.
   *
   * @generated from field: string calendar_id = 1;
   */
  calendarId: string;

  /**
   * EventId is the unique (per-calendar) id of the event to update.
   *
   * @generated from field: string event_id = 2;
   */
  eventId: string;

  /**
   * Name can be set to the new name/summary of the event.
   *
   * @generated from field: string name = 3;
   */
  name: string;

  /**
   * Description can be set to a new description of the event.
   *
   * @generated from field: string description = 4;
   */
  description: string;

  /**
   * Start can be set to the new start time of the event.
   *
   * @generated from field: google.protobuf.Timestamp start = 5;
   */
  start?: Timestamp;

  /**
   * End can be set to the new end-time of the event.
   *
   * @generated from field: google.protobuf.Timestamp end = 6;
   */
  end?: Timestamp;

  /**
   * ExtraData can be set to replace the extra-data associated with the event.
   *
   * @generated from field: google.protobuf.Any extra_data = 7;
   */
  extraData?: Any;

  /**
   * UpdateMask specifies which fields of the original event should be
   * updated. If left empty, all fields will be replaced with the values from
   * the UpdateEventRequest message potentially clearing out fields if unset.
   *
   * @generated from field: google.protobuf.FieldMask update_mask = 20;
   */
  updateMask?: FieldMask;

  constructor(data?: PartialMessage<UpdateEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.UpdateEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateEventRequest;

  static equals(a: UpdateEventRequest | PlainMessage<UpdateEventRequest> | undefined, b: UpdateEventRequest | PlainMessage<UpdateEventRequest> | undefined): boolean;
}

/**
 * UpdateEventResponse is the response message of the UpdateEvent RPC and
 * contains the updated event definition.
 *
 * @generated from message tkd.calendar.v1.UpdateEventResponse
 */
export declare class UpdateEventResponse extends Message<UpdateEventResponse> {
  /**
   * Event holds the updated calendar event.
   *
   * @generated from field: tkd.calendar.v1.CalendarEvent event = 1;
   */
  event?: CalendarEvent;

  constructor(data?: PartialMessage<UpdateEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.UpdateEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateEventResponse;

  static equals(a: UpdateEventResponse | PlainMessage<UpdateEventResponse> | undefined, b: UpdateEventResponse | PlainMessage<UpdateEventResponse> | undefined): boolean;
}

/**
 * MoveEventRequest is the request message of the MoveEvent RPC and allows to
 * move a calendar event from one calendar to another one. Note that the ID of
 * the event might change after a successfull move!
 *
 * @generated from message tkd.calendar.v1.MoveEventRequest
 */
export declare class MoveEventRequest extends Message<MoveEventRequest> {
  /**
   * Source describes the source of the event.
   *
   * @generated from oneof tkd.calendar.v1.MoveEventRequest.source
   */
  source: {
    /**
     * SourceCalendarId is the origin calendar that contains the event to be
     * moved.
     *
     * @generated from field: string source_calendar_id = 1;
     */
    value: string;
    case: "sourceCalendarId";
  } | {
    /**
     * SourceUserId is the id of the user which owns the event. cis-cal will
     * lookup the associated calendar id from cis-idm's additional user
     * fields.
     *
     * @generated from field: string source_user_id = 2;
     */
    value: string;
    case: "sourceUserId";
  } | { case: undefined; value?: undefined };

  /**
   * EventId is the unique (per-calendar) id of the event to move to a new
   * calendar.
   *
   * @generated from field: string event_id = 3;
   */
  eventId: string;

  /**
   * Target describes the target calendar where the event should be moved to.
   *
   * @generated from oneof tkd.calendar.v1.MoveEventRequest.target
   */
  target: {
    /**
     * TargetCalendarId is the target calendar that should contain the
     * event.
     *
     * @generated from field: string target_calendar_id = 4;
     */
    value: string;
    case: "targetCalendarId";
  } | {
    /**
     * TargetUserId is the id of the user to which the event should be
     * moved. cis-cal will lookup the associated calendar id from cis-idm's
     * additional user fields.
     *
     * @generated from field: string target_user_id = 5;
     */
    value: string;
    case: "targetUserId";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<MoveEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.MoveEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MoveEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MoveEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MoveEventRequest;

  static equals(a: MoveEventRequest | PlainMessage<MoveEventRequest> | undefined, b: MoveEventRequest | PlainMessage<MoveEventRequest> | undefined): boolean;
}

/**
 * MoveEventResponse is the response message of the MoveEvent RPC and contains
 * the moved event definition. Note that after a successful move, the ID of the
 * event might have changed!
 *
 * @generated from message tkd.calendar.v1.MoveEventResponse
 */
export declare class MoveEventResponse extends Message<MoveEventResponse> {
  /**
   * Event holds the moved event definition.
   *
   * @generated from field: tkd.calendar.v1.CalendarEvent event = 1;
   */
  event?: CalendarEvent;

  constructor(data?: PartialMessage<MoveEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.MoveEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MoveEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MoveEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MoveEventResponse;

  static equals(a: MoveEventResponse | PlainMessage<MoveEventResponse> | undefined, b: MoveEventResponse | PlainMessage<MoveEventResponse> | undefined): boolean;
}

/**
 * DeleteEventRequest is the request message for the DeleteEvent RPC and allows
 * to remove a calendar event.
 *
 * @generated from message tkd.calendar.v1.DeleteEventRequest
 */
export declare class DeleteEventRequest extends Message<DeleteEventRequest> {
  /**
   * CalendarId is the ID of the origin calendar that contains the event.
   *
   * @generated from field: string calendar_id = 1;
   */
  calendarId: string;

  /**
   * EventId is the ID of the event that should be removed from the calendar
   * specified by calendar_id.
   *
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
 * DeleteEventResponse is the response message of the DeleteEvent RPC. This
 * message does not contain any fields for now.
 *
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

