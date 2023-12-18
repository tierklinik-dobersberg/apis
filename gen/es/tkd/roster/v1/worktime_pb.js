// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/worktime.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, FieldMask, proto3, Timestamp } from "@bufbuild/protobuf";
import { OffTimeCosts } from "./offtime_pb.js";

/**
 * WorkTime describes the regular work time for a given user.
 *
 * @generated from message tkd.roster.v1.WorkTime
 */
export const WorkTime = proto3.makeMessageType(
  "tkd.roster.v1.WorkTime",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "time_per_week", kind: "message", T: Duration },
    { no: 4, name: "applicable_after", kind: "message", T: Timestamp },
    { no: 5, name: "vacation_weeks_per_year", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 6, name: "overtime_allowance_per_month", kind: "message", T: Duration },
  ],
);

/**
 * @generated from message tkd.roster.v1.SetWorkTimeRequest
 */
export const SetWorkTimeRequest = proto3.makeMessageType(
  "tkd.roster.v1.SetWorkTimeRequest",
  () => [
    { no: 1, name: "work_times", kind: "message", T: WorkTime, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.SetWorkTimeResponse
 */
export const SetWorkTimeResponse = proto3.makeMessageType(
  "tkd.roster.v1.SetWorkTimeResponse",
  () => [
    { no: 1, name: "work_times", kind: "message", T: WorkTime, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetWorkTimeRequest
 */
export const GetWorkTimeRequest = proto3.makeMessageType(
  "tkd.roster.v1.GetWorkTimeRequest",
  () => [
    { no: 1, name: "user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 2, name: "read_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.roster.v1.UserWorkTime
 */
export const UserWorkTime = proto3.makeMessageType(
  "tkd.roster.v1.UserWorkTime",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "current", kind: "message", T: WorkTime },
    { no: 3, name: "history", kind: "message", T: WorkTime, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetWorkTimeResponse
 */
export const GetWorkTimeResponse = proto3.makeMessageType(
  "tkd.roster.v1.GetWorkTimeResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: UserWorkTime, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.SumForUsers
 */
export const SumForUsers = proto3.makeMessageType(
  "tkd.roster.v1.SumForUsers",
  () => [
    { no: 1, name: "user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetVacationCreditsLeftRequest
 */
export const GetVacationCreditsLeftRequest = proto3.makeMessageType(
  "tkd.roster.v1.GetVacationCreditsLeftRequest",
  () => [
    { no: 1, name: "for_users", kind: "message", T: SumForUsers },
    { no: 2, name: "until", kind: "message", T: Timestamp },
    { no: 3, name: "analyze", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.AnalyzeVacationSum
 */
export const AnalyzeVacationSum = proto3.makeMessageType(
  "tkd.roster.v1.AnalyzeVacationSum",
  () => [
    { no: 1, name: "work_time", kind: "message", T: WorkTime },
    { no: 2, name: "ends_at", kind: "message", T: Timestamp },
    { no: 3, name: "number_of_days", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 4, name: "vacation_weeks_per_day", kind: "scalar", T: 2 /* ScalarType.FLOAT */ },
    { no: 5, name: "vacation_per_work_time", kind: "message", T: Duration },
    { no: 6, name: "costs", kind: "message", T: OffTimeCosts, repeated: true },
    { no: 7, name: "costs_sum", kind: "message", T: Duration },
  ],
);

/**
 * @generated from message tkd.roster.v1.AnalyzeVacation
 */
export const AnalyzeVacation = proto3.makeMessageType(
  "tkd.roster.v1.AnalyzeVacation",
  () => [
    { no: 1, name: "slices", kind: "message", T: AnalyzeVacationSum, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.UserVacationSum
 */
export const UserVacationSum = proto3.makeMessageType(
  "tkd.roster.v1.UserVacationSum",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "vacation_credits_left", kind: "message", T: Duration },
    { no: 3, name: "time_off_credits", kind: "message", T: Duration },
    { no: 4, name: "analysis", kind: "message", T: AnalyzeVacation },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetVacationCreditsLeftResponse
 */
export const GetVacationCreditsLeftResponse = proto3.makeMessageType(
  "tkd.roster.v1.GetVacationCreditsLeftResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: UserVacationSum, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.DeleteWorkTimeRequest
 */
export const DeleteWorkTimeRequest = proto3.makeMessageType(
  "tkd.roster.v1.DeleteWorkTimeRequest",
  () => [
    { no: 1, name: "ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.DeleteWorkTimeResponse
 */
export const DeleteWorkTimeResponse = proto3.makeMessageType(
  "tkd.roster.v1.DeleteWorkTimeResponse",
  [],
);

