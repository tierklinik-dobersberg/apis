// @generated by protoc-gen-es v1.3.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/offtime.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { BoolValue, Duration, FieldMask, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum tkd.roster.v1.OffTimeType
 */
export const OffTimeType = proto3.makeEnum(
  "tkd.roster.v1.OffTimeType",
  [
    {no: 0, name: "OFF_TIME_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "OFF_TIME_TYPE_VACATION", localName: "VACATION"},
    {no: 2, name: "OFF_TIME_TYPE_TIME_OFF", localName: "TIME_OFF"},
  ],
);

/**
 * @generated from enum tkd.roster.v1.ApprovalRequestType
 */
export const ApprovalRequestType = proto3.makeEnum(
  "tkd.roster.v1.ApprovalRequestType",
  [
    {no: 0, name: "APPROVAL_REQUEST_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "APPROVAL_REQUEST_TYPE_APPROVED", localName: "APPROVED"},
    {no: 2, name: "APPROVAL_REQUEST_TYPE_REJECTED", localName: "REJECTED"},
  ],
);

/**
 * @generated from message tkd.roster.v1.OffTimeApproval
 */
export const OffTimeApproval = proto3.makeMessageType(
  "tkd.roster.v1.OffTimeApproval",
  () => [
    { no: 2, name: "approved", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "approved_at", kind: "message", T: Timestamp },
    { no: 4, name: "approver_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "comment", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.OffTimeCosts
 */
export const OffTimeCosts = proto3.makeMessageType(
  "tkd.roster.v1.OffTimeCosts",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "offtime_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "roster_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "created_at", kind: "message", T: Timestamp },
    { no: 6, name: "creator_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "costs", kind: "message", T: Duration },
    { no: 8, name: "is_vacation", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 9, name: "date", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message tkd.roster.v1.OffTimeEntry
 */
export const OffTimeEntry = proto3.makeMessageType(
  "tkd.roster.v1.OffTimeEntry",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "requestor_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "from", kind: "message", T: Timestamp },
    { no: 4, name: "to", kind: "message", T: Timestamp },
    { no: 5, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "type", kind: "enum", T: proto3.getEnumType(OffTimeType) },
    { no: 7, name: "approval", kind: "message", T: OffTimeApproval },
    { no: 8, name: "created_at", kind: "message", T: Timestamp },
    { no: 9, name: "creator_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetOffTimeEntryRequest
 */
export const GetOffTimeEntryRequest = proto3.makeMessageType(
  "tkd.roster.v1.GetOffTimeEntryRequest",
  () => [
    { no: 1, name: "ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetOffTimeEntryResponse
 */
export const GetOffTimeEntryResponse = proto3.makeMessageType(
  "tkd.roster.v1.GetOffTimeEntryResponse",
  () => [
    { no: 1, name: "entry", kind: "message", T: OffTimeEntry, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.CreateOffTimeRequestRequest
 */
export const CreateOffTimeRequestRequest = proto3.makeMessageType(
  "tkd.roster.v1.CreateOffTimeRequestRequest",
  () => [
    { no: 1, name: "from", kind: "message", T: Timestamp },
    { no: 2, name: "to", kind: "message", T: Timestamp },
    { no: 3, name: "requestor_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "request_type", kind: "enum", T: proto3.getEnumType(OffTimeType) },
  ],
);

/**
 * @generated from message tkd.roster.v1.CreateOffTimeRequestResponse
 */
export const CreateOffTimeRequestResponse = proto3.makeMessageType(
  "tkd.roster.v1.CreateOffTimeRequestResponse",
  () => [
    { no: 1, name: "entry", kind: "message", T: OffTimeEntry },
  ],
);

/**
 * Request deletion of one or more off-time-request.
 *
 * @generated from message tkd.roster.v1.DeleteOffTimeRequestRequest
 */
export const DeleteOffTimeRequestRequest = proto3.makeMessageType(
  "tkd.roster.v1.DeleteOffTimeRequestRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.DeleteOffTimeRequestResponse
 */
export const DeleteOffTimeRequestResponse = proto3.makeMessageType(
  "tkd.roster.v1.DeleteOffTimeRequestResponse",
  [],
);

/**
 * @generated from message tkd.roster.v1.FindOffTimeRequestsRequest
 */
export const FindOffTimeRequestsRequest = proto3.makeMessageType(
  "tkd.roster.v1.FindOffTimeRequestsRequest",
  () => [
    { no: 1, name: "from", kind: "message", T: Timestamp },
    { no: 2, name: "to", kind: "message", T: Timestamp },
    { no: 3, name: "user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 4, name: "approved", kind: "message", T: BoolValue },
    { no: 5, name: "read_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.roster.v1.FindOffTimeRequestsResponse
 */
export const FindOffTimeRequestsResponse = proto3.makeMessageType(
  "tkd.roster.v1.FindOffTimeRequestsResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: OffTimeEntry, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.ApproveOrRejectRequest
 */
export const ApproveOrRejectRequest = proto3.makeMessageType(
  "tkd.roster.v1.ApproveOrRejectRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "comment", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "type", kind: "enum", T: proto3.getEnumType(ApprovalRequestType) },
  ],
);

/**
 * @generated from message tkd.roster.v1.ApproveOrRejectResponse
 */
export const ApproveOrRejectResponse = proto3.makeMessageType(
  "tkd.roster.v1.ApproveOrRejectResponse",
  () => [
    { no: 1, name: "entry", kind: "message", T: OffTimeEntry },
  ],
);

/**
 * @generated from message tkd.roster.v1.AddOffTimeCostsRequest
 */
export const AddOffTimeCostsRequest = proto3.makeMessageType(
  "tkd.roster.v1.AddOffTimeCostsRequest",
  () => [
    { no: 1, name: "add_costs", kind: "message", T: OffTimeCosts, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.AddOffTimeCostsResponse
 */
export const AddOffTimeCostsResponse = proto3.makeMessageType(
  "tkd.roster.v1.AddOffTimeCostsResponse",
  [],
);

/**
 * @generated from message tkd.roster.v1.CostsForUsers
 */
export const CostsForUsers = proto3.makeMessageType(
  "tkd.roster.v1.CostsForUsers",
  () => [
    { no: 1, name: "user_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetOffTimeCostsRequest
 */
export const GetOffTimeCostsRequest = proto3.makeMessageType(
  "tkd.roster.v1.GetOffTimeCostsRequest",
  () => [
    { no: 1, name: "for_users", kind: "message", T: CostsForUsers },
    { no: 2, name: "read_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.roster.v1.OffTimeCostSummary
 */
export const OffTimeCostSummary = proto3.makeMessageType(
  "tkd.roster.v1.OffTimeCostSummary",
  () => [
    { no: 1, name: "vacation", kind: "message", T: Duration },
    { no: 2, name: "time_off", kind: "message", T: Duration },
  ],
);

/**
 * @generated from message tkd.roster.v1.UserOffTimeCosts
 */
export const UserOffTimeCosts = proto3.makeMessageType(
  "tkd.roster.v1.UserOffTimeCosts",
  () => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "costs", kind: "message", T: OffTimeCosts, repeated: true },
    { no: 3, name: "summary", kind: "message", T: OffTimeCostSummary },
  ],
);

/**
 * @generated from message tkd.roster.v1.GetOffTimeCostsResponse
 */
export const GetOffTimeCostsResponse = proto3.makeMessageType(
  "tkd.roster.v1.GetOffTimeCostsResponse",
  () => [
    { no: 1, name: "results", kind: "message", T: UserOffTimeCosts, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.DeleteOffTimeCostsRequest
 */
export const DeleteOffTimeCostsRequest = proto3.makeMessageType(
  "tkd.roster.v1.DeleteOffTimeCostsRequest",
  () => [
    { no: 1, name: "ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message tkd.roster.v1.DeleteOffTimeCostsResponse
 */
export const DeleteOffTimeCostsResponse = proto3.makeMessageType(
  "tkd.roster.v1.DeleteOffTimeCostsResponse",
  [],
);

