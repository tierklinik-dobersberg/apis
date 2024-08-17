// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file tkd/pbx3cx/v1/calllog.proto (package tkd.pbx3cx.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Duration, FieldMask, proto3, Timestamp } from "@bufbuild/protobuf";
import { Profile } from "../../idm/v1/user_pb.js";
import { TimeRange } from "../../common/v1/time_range_pb.js";
import { Customer, CustomerRef } from "../../customer/v1/customer_pb.js";

/**
 * @generated from message tkd.pbx3cx.v1.CallEntry
 */
export const CallEntry = /*@__PURE__*/ proto3.makeMessageType(
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
export const OnCall = /*@__PURE__*/ proto3.makeMessageType(
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
export const CustomOverwrite = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.CustomOverwrite",
  () => [
    { no: 1, name: "transfer_target", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.Overwrite
 */
export const Overwrite = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.Overwrite",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "from", kind: "message", T: Timestamp },
    { no: 3, name: "to", kind: "message", T: Timestamp },
    { no: 4, name: "custom", kind: "message", T: CustomOverwrite, oneof: "target" },
    { no: 5, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "target" },
    { no: 6, name: "created_at", kind: "message", T: Timestamp },
    { no: 7, name: "created_by_user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "inbound_number", kind: "message", T: InboundNumber },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.InboundNumber
 */
export const InboundNumber = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.InboundNumber",
  () => [
    { no: 1, name: "number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "roster_shift_tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "roster_type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.RecordCallRequest
 */
export const RecordCallRequest = /*@__PURE__*/ proto3.makeMessageType(
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
export const GetLogsForCustomerRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetLogsForCustomerRequest",
  () => [
    { no: 1, name: "source", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForCustomerResponse
 */
export const GetLogsForCustomerResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetLogsForCustomerResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: CallEntry, repeated: true },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForDateRequest
 */
export const GetLogsForDateRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetLogsForDateRequest",
  () => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForDateResponse
 */
export const GetLogsForDateResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetLogsForDateResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: CallEntry, repeated: true },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.SearchCallLogsRequest
 */
export const SearchCallLogsRequest = /*@__PURE__*/ proto3.makeMessageType(
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
export const SearchCallLogsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.SearchCallLogsResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: CallEntry, repeated: true },
    { no: 2, name: "customers", kind: "message", T: Customer, repeated: true },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetOnCallRequest
 */
export const GetOnCallRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetOnCallRequest",
  () => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "ignore_overwrites", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "inbound_number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetOnCallResponse
 */
export const GetOnCallResponse = /*@__PURE__*/ proto3.makeMessageType(
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
export const CreateOverwriteRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.CreateOverwriteRequest",
  () => [
    { no: 1, name: "custom", kind: "message", T: CustomOverwrite, oneof: "transfer_target" },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "transfer_target" },
    { no: 3, name: "from", kind: "message", T: Timestamp },
    { no: 4, name: "to", kind: "message", T: Timestamp },
    { no: 5, name: "inbound_number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.CreateOverwriteResponse
 */
export const CreateOverwriteResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.CreateOverwriteResponse",
  () => [
    { no: 1, name: "overwrite", kind: "message", T: Overwrite },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.DeleteOverwriteRequest
 */
export const DeleteOverwriteRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.DeleteOverwriteRequest",
  () => [
    { no: 1, name: "overwrite_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "selector" },
    { no: 2, name: "active_at", kind: "message", T: Timestamp, oneof: "selector" },
    { no: 3, name: "inbound_numbers", kind: "message", T: InboundNumberList },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.DeleteOverwriteResponse
 */
export const DeleteOverwriteResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.DeleteOverwriteResponse",
  [],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetOverwriteRequest
 */
export const GetOverwriteRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetOverwriteRequest",
  () => [
    { no: 1, name: "overwrite_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "selector" },
    { no: 2, name: "active_at", kind: "message", T: Timestamp, oneof: "selector" },
    { no: 3, name: "time_range", kind: "message", T: TimeRange, oneof: "selector" },
    { no: 4, name: "inbound_numbers", kind: "message", T: InboundNumberList },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.InboundNumberList
 */
export const InboundNumberList = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.InboundNumberList",
  () => [
    { no: 1, name: "numbers", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.GetOverwriteResponse
 */
export const GetOverwriteResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.GetOverwriteResponse",
  () => [
    { no: 1, name: "overwrites", kind: "message", T: Overwrite, repeated: true },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.CreateInboundNumberRequest
 */
export const CreateInboundNumberRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.CreateInboundNumberRequest",
  () => [
    { no: 1, name: "number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "roster_shift_tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "roster_type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.CreateInboundNumberResponse
 */
export const CreateInboundNumberResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.CreateInboundNumberResponse",
  () => [
    { no: 1, name: "inbound_number", kind: "message", T: InboundNumber },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.DeleteInboundNumberRequest
 */
export const DeleteInboundNumberRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.DeleteInboundNumberRequest",
  () => [
    { no: 1, name: "number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.DeleteInboundNumberResponse
 */
export const DeleteInboundNumberResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.DeleteInboundNumberResponse",
  [],
);

/**
 * @generated from message tkd.pbx3cx.v1.ListInboundNumberRequest
 */
export const ListInboundNumberRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.ListInboundNumberRequest",
  [],
);

/**
 * @generated from message tkd.pbx3cx.v1.ListInboundNumberResponse
 */
export const ListInboundNumberResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.ListInboundNumberResponse",
  () => [
    { no: 1, name: "inbound_numbers", kind: "message", T: InboundNumber, repeated: true },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.UpdateInboundNumberRequest
 */
export const UpdateInboundNumberRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.UpdateInboundNumberRequest",
  () => [
    { no: 1, name: "number", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "new_display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "roster_shift_tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "roster_type_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 10, name: "update_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.pbx3cx.v1.UpdateInboundNumberResponse
 */
export const UpdateInboundNumberResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.pbx3cx.v1.UpdateInboundNumberResponse",
  () => [
    { no: 1, name: "inbound_number", kind: "message", T: InboundNumber },
  ],
);

