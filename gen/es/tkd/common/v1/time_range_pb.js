// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file tkd/common/v1/time_range.proto (package tkd.common.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * TimeRange describes a time range with a start and end-time.
 * If both, start and end time is unset, the time range is not valid
 * and no times will match.
 *
 * @generated from message tkd.common.v1.TimeRange
 */
export const TimeRange = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.common.v1.TimeRange",
  () => [
    { no: 1, name: "from", kind: "message", T: Timestamp },
    { no: 2, name: "to", kind: "message", T: Timestamp },
  ],
);

