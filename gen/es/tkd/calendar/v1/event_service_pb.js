// @generated by protoc-gen-es v1.3.0 with parameter "target=js+dts"
// @generated from file tkd/calendar/v1/event_service.proto (package tkd.calendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Any, FieldMask, proto3, Timestamp } from "@bufbuild/protobuf";
import { Calendar, CalendarEvent } from "./event_pb.js";
import { TimeRange } from "../../common/v1/time_range_pb.js";

/**
 * @generated from message tkd.calendar.v1.ListCalendarsRequest
 */
export const ListCalendarsRequest = proto3.makeMessageType(
  "tkd.calendar.v1.ListCalendarsRequest",
  [],
);

/**
 * @generated from message tkd.calendar.v1.ListCalendarsResponse
 */
export const ListCalendarsResponse = proto3.makeMessageType(
  "tkd.calendar.v1.ListCalendarsResponse",
  () => [
    { no: 1, name: "calendars", kind: "message", T: Calendar, repeated: true },
  ],
);

/**
 * @generated from message tkd.calendar.v1.EventSource
 */
export const EventSource = proto3.makeMessageType(
  "tkd.calendar.v1.EventSource",
  () => [
    { no: 1, name: "calendar_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.calendar.v1.ListEventsRequest
 */
export const ListEventsRequest = proto3.makeMessageType(
  "tkd.calendar.v1.ListEventsRequest",
  () => [
    { no: 1, name: "sources", kind: "message", T: EventSource, oneof: "source" },
    { no: 2, name: "all_calendars", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "source" },
    { no: 3, name: "all_users", kind: "scalar", T: 8 /* ScalarType.BOOL */, oneof: "source" },
    { no: 4, name: "time_range", kind: "message", T: TimeRange, oneof: "search_time" },
    { no: 5, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "search_time" },
    { no: 6, name: "read_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.calendar.v1.CalendarEventList
 */
export const CalendarEventList = proto3.makeMessageType(
  "tkd.calendar.v1.CalendarEventList",
  () => [
    { no: 1, name: "calendar", kind: "message", T: Calendar },
    { no: 2, name: "events", kind: "message", T: CalendarEvent, repeated: true },
  ],
);

/**
 * @generated from message tkd.calendar.v1.ListEventsResponse
 */
export const ListEventsResponse = proto3.makeMessageType(
  "tkd.calendar.v1.ListEventsResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: CalendarEventList, repeated: true },
  ],
);

/**
 * @generated from message tkd.calendar.v1.CreateEventRequest
 */
export const CreateEventRequest = proto3.makeMessageType(
  "tkd.calendar.v1.CreateEventRequest",
  () => [
    { no: 1, name: "calendar_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "start", kind: "message", T: Timestamp },
    { no: 5, name: "end", kind: "message", T: Timestamp },
    { no: 6, name: "extra_data", kind: "message", T: Any },
  ],
);

/**
 * @generated from message tkd.calendar.v1.CreateEventResponse
 */
export const CreateEventResponse = proto3.makeMessageType(
  "tkd.calendar.v1.CreateEventResponse",
  () => [
    { no: 1, name: "event", kind: "message", T: CalendarEvent },
  ],
);

/**
 * @generated from message tkd.calendar.v1.DeleteEventRequest
 */
export const DeleteEventRequest = proto3.makeMessageType(
  "tkd.calendar.v1.DeleteEventRequest",
  () => [
    { no: 1, name: "calendar_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "event_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.calendar.v1.DeleteEventResponse
 */
export const DeleteEventResponse = proto3.makeMessageType(
  "tkd.calendar.v1.DeleteEventResponse",
  [],
);

