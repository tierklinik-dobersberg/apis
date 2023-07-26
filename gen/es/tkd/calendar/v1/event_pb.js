// @generated by protoc-gen-es v1.3.0 with parameter "target=js+dts"
// @generated from file tkd/calendar/v1/event.proto (package tkd.calendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Any, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * Calendar describes a event calendar.
 *
 * @generated from message tkd.calendar.v1.Calendar
 */
export const Calendar = proto3.makeMessageType(
  "tkd.calendar.v1.Calendar",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "timezone", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.calendar.v1.CalendarEvent
 */
export const CalendarEvent = proto3.makeMessageType(
  "tkd.calendar.v1.CalendarEvent",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "calendar_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "start_time", kind: "message", T: Timestamp },
    { no: 4, name: "end_time", kind: "message", T: Timestamp },
    { no: 5, name: "full_day", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "summary", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "extra_data", kind: "message", T: Any },
  ],
);

/**
 * @generated from message tkd.calendar.v1.CustomerAnnotation
 */
export const CustomerAnnotation = proto3.makeMessageType(
  "tkd.calendar.v1.CustomerAnnotation",
  () => [
    { no: 1, name: "customer_source", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "customer_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "animal_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "created_by_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

