// @generated by protoc-gen-es v1.5.0 with parameter "target=js+dts"
// @generated from file tkd/common/v1/time_range.proto (package tkd.common.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message tkd.common.v1.TimeRange
 */
export const TimeRange = proto3.makeMessageType(
  "tkd.common.v1.TimeRange",
  () => [
    { no: 1, name: "from", kind: "message", T: Timestamp },
    { no: 2, name: "to", kind: "message", T: Timestamp },
  ],
);

