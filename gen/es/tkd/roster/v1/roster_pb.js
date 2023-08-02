// @generated by protoc-gen-es v1.3.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/roster.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, FieldMask, proto3, Timestamp } from "@bufbuild/protobuf";
import { WorkShift } from "./workshift_pb.js";

/**
 * @generated from message tkd.roster.v1.PlannedShift
 */
export const PlannedShift = proto3.makeMessageType(
  "tkd.roster.v1.PlannedShift",
  () => [
    { no: 1, name: "assigned_user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "work_shift_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "is_holiday", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "is_weekend", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "from", kind: "message", T: Timestamp },
    { no: 6, name: "to", kind: "message", T: Timestamp },
    { no: 7, name: "time_worth", kind: "message", T: Duration },
    { no: 8, name: "shift", kind: "message", T: WorkShift },
    { no: 9, name: "eligible_user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 10, name: "read_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.roster.v1.RosterMeta
 */
export const RosterMeta = proto3.makeMessageType(
  "tkd.roster.v1.RosterMeta",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "month", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "year", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "approved", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 5, name: "approved_at", kind: "message", T: Timestamp },
    { no: 6, name: "last_modified_by", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "created_at", kind: "message", T: Timestamp },
    { no: 8, name: "updated_at", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message tkd.roster.v1.WorkTimeStatus
 */
export const WorkTimeStatus = proto3.makeMessageType(
  "tkd.roster.v1.WorkTimeStatus",
  () => [
    { no: 1, name: "time_per_week", kind: "message", T: Duration },
    { no: 2, name: "expected_work_time", kind: "message", T: Duration },
    { no: 3, name: "planned_work_time", kind: "message", T: Duration },
  ],
);

/**
 * @generated from message tkd.roster.v1.StartSessionRequest
 */
export const StartSessionRequest = proto3.makeMessageType(
  "tkd.roster.v1.StartSessionRequest",
  () => [
    { no: 1, name: "year", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "month", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.StartSessionResponse
 */
export const StartSessionResponse = proto3.makeMessageType(
  "tkd.roster.v1.StartSessionResponse",
  () => [
    { no: 1, name: "meta", kind: "message", T: RosterMeta },
    { no: 2, name: "planned_shifts", kind: "message", T: PlannedShift, repeated: true },
    { no: 3, name: "active_users", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.AssignUserToShift
 */
export const AssignUserToShift = proto3.makeMessageType(
  "tkd.roster.v1.AssignUserToShift",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "shift_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "from", kind: "message", T: Timestamp },
    { no: 4, name: "to", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message tkd.roster.v1.UnassignUserFromShift
 */
export const UnassignUserFromShift = proto3.makeMessageType(
  "tkd.roster.v1.UnassignUserFromShift",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "shift_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "from", kind: "message", T: Timestamp },
    { no: 4, name: "to", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message tkd.roster.v1.ShiftUpdateEvent
 */
export const ShiftUpdateEvent = proto3.makeMessageType(
  "tkd.roster.v1.ShiftUpdateEvent",
  () => [
    { no: 1, name: "shift", kind: "message", T: PlannedShift },
    { no: 2, name: "changed_by", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.WorkTimeUpdateEvent
 */
export const WorkTimeUpdateEvent = proto3.makeMessageType(
  "tkd.roster.v1.WorkTimeUpdateEvent",
  () => [
    { no: 1, name: "work_times", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: WorkTimeStatus} },
  ],
);

/**
 * @generated from message tkd.roster.v1.SessionPing
 */
export const SessionPing = proto3.makeMessageType(
  "tkd.roster.v1.SessionPing",
  [],
);

/**
 * @generated from message tkd.roster.v1.SessionRequest
 */
export const SessionRequest = proto3.makeMessageType(
  "tkd.roster.v1.SessionRequest",
  () => [
    { no: 1, name: "start_session", kind: "message", T: StartSessionRequest, oneof: "kind" },
    { no: 2, name: "assign_user", kind: "message", T: AssignUserToShift, oneof: "kind" },
    { no: 3, name: "unassign_user", kind: "message", T: UnassignUserFromShift, oneof: "kind" },
    { no: 4, name: "ping", kind: "message", T: SessionPing, oneof: "kind" },
  ],
);

/**
 * @generated from message tkd.roster.v1.SessionResponse
 */
export const SessionResponse = proto3.makeMessageType(
  "tkd.roster.v1.SessionResponse",
  () => [
    { no: 1, name: "start_session", kind: "message", T: StartSessionResponse, oneof: "kind" },
    { no: 2, name: "shift_update", kind: "message", T: ShiftUpdateEvent, oneof: "kind" },
    { no: 3, name: "work_time_update", kind: "message", T: WorkTimeUpdateEvent, oneof: "kind" },
    { no: 4, name: "user_joined", kind: "message", T: UserJoinedSessionEvent, oneof: "kind" },
    { no: 5, name: "user_left", kind: "message", T: UserLeftSessionEvent, oneof: "kind" },
    { no: 6, name: "message", kind: "message", T: UserChatMessageEvent, oneof: "kind" },
    { no: 7, name: "ping", kind: "message", T: SessionPing, oneof: "kind" },
  ],
);

/**
 * @generated from message tkd.roster.v1.UserJoinedSessionEvent
 */
export const UserJoinedSessionEvent = proto3.makeMessageType(
  "tkd.roster.v1.UserJoinedSessionEvent",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.UserLeftSessionEvent
 */
export const UserLeftSessionEvent = proto3.makeMessageType(
  "tkd.roster.v1.UserLeftSessionEvent",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.UserChatMessageEvent
 */
export const UserChatMessageEvent = proto3.makeMessageType(
  "tkd.roster.v1.UserChatMessageEvent",
  () => [
    { no: 1, name: "sender_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.ApproveRosterRequest
 */
export const ApproveRosterRequest = proto3.makeMessageType(
  "tkd.roster.v1.ApproveRosterRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.ApproveRosterResponse
 */
export const ApproveRosterResponse = proto3.makeMessageType(
  "tkd.roster.v1.ApproveRosterResponse",
  [],
);

/**
 * @generated from message tkd.roster.v1.DeleteRosterRequest
 */
export const DeleteRosterRequest = proto3.makeMessageType(
  "tkd.roster.v1.DeleteRosterRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.DeleteRosterResponse
 */
export const DeleteRosterResponse = proto3.makeMessageType(
  "tkd.roster.v1.DeleteRosterResponse",
  [],
);

/**
 * @generated from message tkd.roster.v1.GetByDate
 */
export const GetByDate = proto3.makeMessageType(
  "tkd.roster.v1.GetByDate",
  () => [
    { no: 1, name: "year", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "month", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetRosterRequest
 */
export const GetRosterRequest = proto3.makeMessageType(
  "tkd.roster.v1.GetRosterRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "search" },
    { no: 2, name: "date", kind: "message", T: GetByDate, oneof: "search" },
    { no: 3, name: "read_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetRosterResponse
 */
export const GetRosterResponse = proto3.makeMessageType(
  "tkd.roster.v1.GetRosterResponse",
  () => [
    { no: 1, name: "meta", kind: "message", T: RosterMeta },
    { no: 2, name: "shifts", kind: "message", T: PlannedShift, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetWorkingStaffRequest
 */
export const GetWorkingStaffRequest = proto3.makeMessageType(
  "tkd.roster.v1.GetWorkingStaffRequest",
  () => [
    { no: 1, name: "time", kind: "message", T: Timestamp },
    { no: 2, name: "read_maks", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetWorkingStaffResponse
 */
export const GetWorkingStaffResponse = proto3.makeMessageType(
  "tkd.roster.v1.GetWorkingStaffResponse",
  () => [
    { no: 1, name: "user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "current_shifts", kind: "message", T: PlannedShift, repeated: true },
  ],
);

