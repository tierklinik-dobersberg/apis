// @generated by protoc-gen-es v1.5.0 with parameter "target=js+dts"
// @generated from file tkd/pbx3cx/v1/calllog.proto (package tkd.pbx3cx.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, proto3, Timestamp } from "@bufbuild/protobuf";
import { Profile } from "../../idm/v1/user_pb.js";
import { TimeRange } from "../../common/v1/time_range_pb.js";
import { CustomerRef } from "../../customer/v1/customer_pb.js";

/**
 * @generated from message tkd.pbx3cx.v1.CallEntry
 */
export const CallEntry = proto3.makeMessageType(
  "tkd.pbx3cx.v1.CallEntry",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "caller", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "inbound_number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "received_at", kind: "message", T: Timestamp },
    { no: 5, name: "duration", kind: "message", T: Duration },
    { no: 6, name: "call_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "agent_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "customer_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "customer_source", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 10, name: "error", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 11, name: "transfer_target", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "accepted_agent", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.OnCall
 */
export const OnCall = proto3.makeMessageType(
  "tkd.pbx3cx.v1.OnCall",
  () => [
    { no: 1, name: "profile", kind: "message", T: Profile },
    { no: 2, name: "transfer_target", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "until", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.CustomOverwrite
 */
export const CustomOverwrite = proto3.makeMessageType(
  "tkd.pbx3cx.v1.CustomOverwrite",
  () => [
    { no: 1, name: "transfer_target", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.Overwrite
 */
export const Overwrite = proto3.makeMessageType(
  "tkd.pbx3cx.v1.Overwrite",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "from", kind: "message", T: Timestamp },
    { no: 3, name: "to", kind: "message", T: Timestamp },
    { no: 4, name: "custom", kind: "message", T: CustomOverwrite, oneof: "target" },
    { no: 5, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "target" },
    { no: 6, name: "created_at", kind: "message", T: Timestamp },
    { no: 7, name: "created_by_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.RecordCallRequest
 */
export const RecordCallRequest = proto3.makeMessageType(
  "tkd.pbx3cx.v1.RecordCallRequest",
  () => [
    { no: 1, name: "duration", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "agent", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "call_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "date_time", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "customer_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "customer_source", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForCustomerRequest
 */
export const GetLogsForCustomerRequest = proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetLogsForCustomerRequest",
  () => [
    { no: 1, name: "source", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForCustomerResponse
 */
export const GetLogsForCustomerResponse = proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetLogsForCustomerResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: CallEntry, repeated: true },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForDateRequest
 */
export const GetLogsForDateRequest = proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetLogsForDateRequest",
  () => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForDateResponse
 */
export const GetLogsForDateResponse = proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetLogsForDateResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: CallEntry, repeated: true },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.SearchCallLogsRequest
 */
export const SearchCallLogsRequest = proto3.makeMessageType(
  "tkd.pbx3cx.v1.SearchCallLogsRequest",
  () => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "time_range", kind: "message", T: TimeRange },
    { no: 3, name: "customer_ref", kind: "message", T: CustomerRef },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.SearchCallLogsResponse
 */
export const SearchCallLogsResponse = proto3.makeMessageType(
  "tkd.pbx3cx.v1.SearchCallLogsResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: CallEntry, repeated: true },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetOnCallRequest
 */
export const GetOnCallRequest = proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetOnCallRequest",
  () => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "ignore_overwrites", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetOnCallResponse
 */
export const GetOnCallResponse = proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetOnCallResponse",
  () => [
    { no: 1, name: "on_call", kind: "message", T: OnCall, repeated: true },
    { no: 2, name: "roster_date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "is_overwrite", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 4, name: "primary_transfer_target", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.CreateOverwriteRequest
 */
export const CreateOverwriteRequest = proto3.makeMessageType(
  "tkd.pbx3cx.v1.CreateOverwriteRequest",
  () => [
    { no: 1, name: "custom", kind: "message", T: CustomOverwrite, oneof: "transfer_target" },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "transfer_target" },
    { no: 3, name: "from", kind: "message", T: Timestamp },
    { no: 4, name: "to", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.CreateOverwriteResponse
 */
export const CreateOverwriteResponse = proto3.makeMessageType(
  "tkd.pbx3cx.v1.CreateOverwriteResponse",
  () => [
    { no: 1, name: "overwrite", kind: "message", T: Overwrite },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.DeleteOverwriteRequest
 */
export const DeleteOverwriteRequest = proto3.makeMessageType(
  "tkd.pbx3cx.v1.DeleteOverwriteRequest",
  () => [
    { no: 1, name: "overwrite_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "selector" },
    { no: 2, name: "active_at", kind: "message", T: Timestamp, oneof: "selector" },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.DeleteOverwriteResponse
 */
export const DeleteOverwriteResponse = proto3.makeMessageType(
  "tkd.pbx3cx.v1.DeleteOverwriteResponse",
  [],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetOverwriteRequest
 */
export const GetOverwriteRequest = proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetOverwriteRequest",
  () => [
    { no: 1, name: "overwrite_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "selector" },
    { no: 2, name: "active_at", kind: "message", T: Timestamp, oneof: "selector" },
    { no: 3, name: "time_range", kind: "message", T: TimeRange, oneof: "selector" },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetOverwriteResponse
 */
export const GetOverwriteResponse = proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetOverwriteResponse",
  () => [
    { no: 1, name: "overwrites", kind: "message", T: Overwrite, repeated: true },
  ],
);

