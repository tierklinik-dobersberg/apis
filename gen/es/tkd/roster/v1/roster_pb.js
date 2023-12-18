// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/roster.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, FieldMask, proto3, Timestamp } from "@bufbuild/protobuf";
import { ConstraintViolationList } from "./constraint_pb.js";
import { WorkShift } from "./workshift_pb.js";
import { DeliveryNotification } from "../../idm/v1/notify_service_pb.js";
import { TimeRange } from "../../common/v1/time_range_pb.js";

/**
 * @generated from message tkd.roster.v1.RequiredShift
 */
export const RequiredShift = proto3.makeMessageType(
  "tkd.roster.v1.RequiredShift",
  () => [
    { no: 1, name: "from", kind: "message", T: Timestamp },
    { no: 2, name: "to", kind: "message", T: Timestamp },
    { no: 3, name: "work_shift_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "eligible_user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 5, name: "on_holiday", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "on_weekend", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 7, name: "violations_per_user_id", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: ConstraintViolationList} },
  ],
);

/**
 * PlannedShift is a planned work-shift in a roster.
 *
 * @generated from message tkd.roster.v1.PlannedShift
 */
export const PlannedShift = proto3.makeMessageType(
  "tkd.roster.v1.PlannedShift",
  () => [
    { no: 1, name: "from", kind: "message", T: Timestamp },
    { no: 2, name: "to", kind: "message", T: Timestamp },
    { no: 3, name: "assigned_user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "work_shift_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.RosterType
 */
export const RosterType = proto3.makeMessageType(
  "tkd.roster.v1.RosterType",
  () => [
    { no: 1, name: "unique_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "shift_tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 3, name: "on_call_tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * Roster is a planned roster for a given time period.
 *
 * @generated from message tkd.roster.v1.Roster
 */
export const Roster = proto3.makeMessageType(
  "tkd.roster.v1.Roster",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "from", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "to", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "shifts", kind: "message", T: PlannedShift, repeated: true },
    { no: 5, name: "approved", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "approved_at", kind: "message", T: Timestamp },
    { no: 7, name: "approver_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "last_modified_by", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "created_at", kind: "message", T: Timestamp },
    { no: 10, name: "updated_at", kind: "message", T: Timestamp },
    { no: 11, name: "roster_type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.SaveRosterRequest
 */
export const SaveRosterRequest = proto3.makeMessageType(
  "tkd.roster.v1.SaveRosterRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "from", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "to", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "shifts", kind: "message", T: PlannedShift, repeated: true },
    { no: 5, name: "read_mask", kind: "message", T: FieldMask },
    { no: 7, name: "shift_tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 8, name: "roster_type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.WorkTimeAnalysisWeek
 */
export const WorkTimeAnalysisWeek = proto3.makeMessageType(
  "tkd.roster.v1.WorkTimeAnalysisWeek",
  () => [
    { no: 1, name: "year", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "week", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "working_days", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "expected_work", kind: "message", T: Duration },
    { no: 5, name: "planned", kind: "message", T: Duration },
  ],
);

/**
 * @generated from message tkd.roster.v1.WorkTimeAnalysisStep
 */
export const WorkTimeAnalysisStep = proto3.makeMessageType(
  "tkd.roster.v1.WorkTimeAnalysisStep",
  () => [
    { no: 1, name: "work_time_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "weeks", kind: "message", T: WorkTimeAnalysisWeek, repeated: true },
    { no: 3, name: "expected_work_time", kind: "message", T: Duration },
    { no: 4, name: "work_time_per_week", kind: "message", T: Duration },
    { no: 5, name: "from", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "to", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "planned", kind: "message", T: Duration },
    { no: 8, name: "overtime", kind: "message", T: Duration },
  ],
);

/**
 * @generated from message tkd.roster.v1.WorkTimeAnalysis
 */
export const WorkTimeAnalysis = proto3.makeMessageType(
  "tkd.roster.v1.WorkTimeAnalysis",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "planned_time", kind: "message", T: Duration },
    { no: 3, name: "expected_time", kind: "message", T: Duration },
    { no: 4, name: "steps", kind: "message", T: WorkTimeAnalysisStep, repeated: true },
    { no: 5, name: "overtime", kind: "message", T: Duration },
  ],
);

/**
 * @generated from message tkd.roster.v1.SaveRosterResponse
 */
export const SaveRosterResponse = proto3.makeMessageType(
  "tkd.roster.v1.SaveRosterResponse",
  () => [
    { no: 1, name: "roster", kind: "message", T: Roster },
    { no: 2, name: "work_time_analysis", kind: "message", T: WorkTimeAnalysis, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.UsersToAnalyze
 */
export const UsersToAnalyze = proto3.makeMessageType(
  "tkd.roster.v1.UsersToAnalyze",
  () => [
    { no: 1, name: "user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "all_users", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.AnalyzeWorkTimeRequest
 */
export const AnalyzeWorkTimeRequest = proto3.makeMessageType(
  "tkd.roster.v1.AnalyzeWorkTimeRequest",
  () => [
    { no: 1, name: "users", kind: "message", T: UsersToAnalyze },
    { no: 3, name: "from", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "to", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.AnalyzeWorkTimeResponse
 */
export const AnalyzeWorkTimeResponse = proto3.makeMessageType(
  "tkd.roster.v1.AnalyzeWorkTimeResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: WorkTimeAnalysis, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.ApproveRosterWorkTimeSplit
 */
export const ApproveRosterWorkTimeSplit = proto3.makeMessageType(
  "tkd.roster.v1.ApproveRosterWorkTimeSplit",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "time_off", kind: "message", T: Duration },
    { no: 3, name: "vacation", kind: "message", T: Duration },
  ],
);

/**
 * @generated from message tkd.roster.v1.ApproveRosterRequest
 */
export const ApproveRosterRequest = proto3.makeMessageType(
  "tkd.roster.v1.ApproveRosterRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "work_time_split", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: ApproveRosterWorkTimeSplit} },
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
 * @generated from message tkd.roster.v1.GetRosterRequest
 */
export const GetRosterRequest = proto3.makeMessageType(
  "tkd.roster.v1.GetRosterRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "search" },
    { no: 2, name: "date", kind: "message", T: Timestamp, oneof: "search" },
    { no: 3, name: "roster_type_names", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "read_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetRosterResponse
 */
export const GetRosterResponse = proto3.makeMessageType(
  "tkd.roster.v1.GetRosterResponse",
  () => [
    { no: 1, name: "roster", kind: "message", T: Roster, repeated: true },
    { no: 2, name: "work_time_analysis", kind: "message", T: WorkTimeAnalysis, repeated: true },
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
    { no: 3, name: "roster_type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "on_call", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
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
    { no: 3, name: "roster_id", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetRequiredShiftsRequest
 */
export const GetRequiredShiftsRequest = proto3.makeMessageType(
  "tkd.roster.v1.GetRequiredShiftsRequest",
  () => [
    { no: 1, name: "from", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "to", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "roster_type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "read_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetRequiredShiftsResponse
 */
export const GetRequiredShiftsResponse = proto3.makeMessageType(
  "tkd.roster.v1.GetRequiredShiftsResponse",
  () => [
    { no: 1, name: "required_shifts", kind: "message", T: RequiredShift, repeated: true },
    { no: 2, name: "work_shift_definitions", kind: "message", T: WorkShift, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.SendRosterPreviewRequest
 */
export const SendRosterPreviewRequest = proto3.makeMessageType(
  "tkd.roster.v1.SendRosterPreviewRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.SendRosterPreviewResponse
 */
export const SendRosterPreviewResponse = proto3.makeMessageType(
  "tkd.roster.v1.SendRosterPreviewResponse",
  () => [
    { no: 1, name: "delivery", kind: "message", T: DeliveryNotification, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.CreateRosterTypeRequest
 */
export const CreateRosterTypeRequest = proto3.makeMessageType(
  "tkd.roster.v1.CreateRosterTypeRequest",
  () => [
    { no: 1, name: "roster_type", kind: "message", T: RosterType },
  ],
);

/**
 * @generated from message tkd.roster.v1.CreateRosterTypeResponse
 */
export const CreateRosterTypeResponse = proto3.makeMessageType(
  "tkd.roster.v1.CreateRosterTypeResponse",
  () => [
    { no: 1, name: "roster_type", kind: "message", T: RosterType },
  ],
);

/**
 * @generated from message tkd.roster.v1.DeleteRosterTypeRequest
 */
export const DeleteRosterTypeRequest = proto3.makeMessageType(
  "tkd.roster.v1.DeleteRosterTypeRequest",
  () => [
    { no: 1, name: "unique_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.DeleteRosterTypeResponse
 */
export const DeleteRosterTypeResponse = proto3.makeMessageType(
  "tkd.roster.v1.DeleteRosterTypeResponse",
  [],
);

/**
 * @generated from message tkd.roster.v1.ListRosterTypesRequest
 */
export const ListRosterTypesRequest = proto3.makeMessageType(
  "tkd.roster.v1.ListRosterTypesRequest",
  [],
);

/**
 * @generated from message tkd.roster.v1.ListRosterTypesResponse
 */
export const ListRosterTypesResponse = proto3.makeMessageType(
  "tkd.roster.v1.ListRosterTypesResponse",
  () => [
    { no: 1, name: "roster_types", kind: "message", T: RosterType, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.ListShiftTagsRequest
 */
export const ListShiftTagsRequest = proto3.makeMessageType(
  "tkd.roster.v1.ListShiftTagsRequest",
  [],
);

/**
 * @generated from message tkd.roster.v1.ListShiftTagsResponse
 */
export const ListShiftTagsResponse = proto3.makeMessageType(
  "tkd.roster.v1.ListShiftTagsResponse",
  () => [
    { no: 1, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetUserShiftsRequest
 */
export const GetUserShiftsRequest = proto3.makeMessageType(
  "tkd.roster.v1.GetUserShiftsRequest",
  () => [
    { no: 1, name: "timerange", kind: "message", T: TimeRange },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetUserShiftsResponse
 */
export const GetUserShiftsResponse = proto3.makeMessageType(
  "tkd.roster.v1.GetUserShiftsResponse",
  () => [
    { no: 1, name: "shifts", kind: "message", T: PlannedShift, repeated: true },
    { no: 2, name: "definitions", kind: "message", T: WorkShift, repeated: true },
  ],
);

