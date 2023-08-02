// @generated by protoc-gen-es v1.3.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/worktime.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, FieldMask, proto3, Timestamp } from "@bufbuild/protobuf";

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

