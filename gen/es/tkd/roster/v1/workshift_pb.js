// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/workshift.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, FieldMask, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message tkd.roster.v1.Daytime
 */
export const Daytime = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.Daytime",
  () => [
    { no: 1, name: "hour", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 2, name: "minute", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.WorkShift
 */
export const WorkShift = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.WorkShift",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "from", kind: "message", T: Daytime },
    { no: 3, name: "duration", kind: "message", T: Duration },
    { no: 4, name: "days", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 5, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "on_holiday", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 8, name: "eligible_role_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 9, name: "time_worth", kind: "message", T: Duration },
    { no: 10, name: "required_staff_count", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 11, name: "color", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 13, name: "order", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 14, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.ListWorkShiftsRequest
 */
export const ListWorkShiftsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.ListWorkShiftsRequest",
  () => [
    { no: 1, name: "read_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.roster.v1.ListWorkShiftsResponse
 */
export const ListWorkShiftsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.ListWorkShiftsResponse",
  () => [
    { no: 1, name: "work_shifts", kind: "message", T: WorkShift, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.CreateWorkShiftRequest
 */
export const CreateWorkShiftRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.CreateWorkShiftRequest",
  () => [
    { no: 1, name: "from", kind: "message", T: Daytime },
    { no: 2, name: "duration", kind: "message", T: Duration },
    { no: 3, name: "days", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 4, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "on_holiday", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 7, name: "eligible_role_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 8, name: "time_worth", kind: "message", T: Duration },
    { no: 9, name: "required_staff_count", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 10, name: "color", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "order", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 13, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.CreateWorkShiftResponse
 */
export const CreateWorkShiftResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.CreateWorkShiftResponse",
  () => [
    { no: 1, name: "work_shift", kind: "message", T: WorkShift },
  ],
);

/**
 * @generated from message tkd.roster.v1.WorkShiftUpdate
 */
export const WorkShiftUpdate = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.WorkShiftUpdate",
  () => [
    { no: 1, name: "from", kind: "message", T: Daytime },
    { no: 2, name: "duration", kind: "message", T: Duration },
    { no: 3, name: "days", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
    { no: 4, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "on_holiday", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 7, name: "eligible_role_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 8, name: "time_worth", kind: "message", T: Duration },
    { no: 9, name: "required_staff_count", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 10, name: "color", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 11, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "order", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 13, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.UpdateWorkShiftRequest
 */
export const UpdateWorkShiftRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.UpdateWorkShiftRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "update", kind: "message", T: WorkShiftUpdate },
    { no: 3, name: "update_in_place", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "write_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.roster.v1.UpdateWorkShiftResponse
 */
export const UpdateWorkShiftResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.UpdateWorkShiftResponse",
  () => [
    { no: 1, name: "work_shift", kind: "message", T: WorkShift },
  ],
);

/**
 * @generated from message tkd.roster.v1.DeleteWorkShiftRequest
 */
export const DeleteWorkShiftRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.DeleteWorkShiftRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.DeleteWorkShiftResponse
 */
export const DeleteWorkShiftResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.roster.v1.DeleteWorkShiftResponse",
  [],
);

