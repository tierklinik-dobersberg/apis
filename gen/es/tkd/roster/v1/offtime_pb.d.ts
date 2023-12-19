// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/offtime.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, Duration, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum tkd.roster.v1.OffTimeType
 */
export declare enum OffTimeType {
  /**
   * Unspecified means that the requestor does not care whether the off-time
   * is based on vacation credits of time compenstation.
   *
   * @generated from enum value: OFF_TIME_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * Vacation is used to tell the manager that this off-time is highly required
   * and the requestor is fine with spending vacation credits.
   *
   * @generated from enum value: OFF_TIME_TYPE_VACATION = 1;
   */
  VACATION = 1,

  /**
   * TimeOff is used to tell the manager that this off-time is not mandatory
   * and that the requestor is only willing to spend time-compensation.
   *
   * @generated from enum value: OFF_TIME_TYPE_TIME_OFF = 2;
   */
  TIME_OFF = 2,
}

/**
 * @generated from enum tkd.roster.v1.ApprovalRequestType
 */
export declare enum ApprovalRequestType {
  /**
   * @generated from enum value: APPROVAL_REQUEST_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: APPROVAL_REQUEST_TYPE_APPROVED = 1;
   */
  APPROVED = 1,

  /**
   * @generated from enum value: APPROVAL_REQUEST_TYPE_REJECTED = 2;
   */
  REJECTED = 2,
}

/**
 * @generated from message tkd.roster.v1.OffTimeApproval
 */
export declare class OffTimeApproval extends Message<OffTimeApproval> {
  /**
   * Approved is set to true when the off-time request has been approved
   * by a manager.
   *
   * @generated from field: bool approved = 2;
   */
  approved: boolean;

  /**
   * ApprovedAt is set to the time-stamp the off-time request has been
   * approved.
   *
   * @generated from field: google.protobuf.Timestamp approved_at = 3;
   */
  approvedAt?: Timestamp;

  /**
   * ApproverId holds the ID of the user that approved the request.
   *
   * @generated from field: string approver_id = 4;
   */
  approverId: string;

  /**
   * Comment may hold an additional comment from management.
   *
   * @generated from field: string comment = 5;
   */
  comment: string;

  constructor(data?: PartialMessage<OffTimeApproval>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.OffTimeApproval";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OffTimeApproval;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OffTimeApproval;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OffTimeApproval;

  static equals(a: OffTimeApproval | PlainMessage<OffTimeApproval> | undefined, b: OffTimeApproval | PlainMessage<OffTimeApproval> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.OffTimeCosts
 */
export declare class OffTimeCosts extends Message<OffTimeCosts> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * OfftimeId is the ID of the off-time request, if any.
   *
   * @generated from field: string offtime_id = 2;
   */
  offtimeId: string;

  /**
   * RosterId is the ID of the roster this OffTimeCosts belong to.
   *
   * @generated from field: string roster_id = 3;
   */
  rosterId: string;

  /**
   * UserId is the ID of the user this off-time costs belong to.
   *
   * @generated from field: string user_id = 4;
   */
  userId: string;

  /**
   * CreatedAt is set to the time this off-time-costs entry has been
   * created. This field must not be set during OffTimeService.AddOffTimeCosts.
   *
   * @generated from field: google.protobuf.Timestamp created_at = 5;
   */
  createdAt?: Timestamp;

  /**
   * CreatorId holds the ID of the user that created this entry.
   * This field must not be set during OffTimeService.AddOffTimeCosts.
   *
   * @generated from field: string creator_id = 6;
   */
  creatorId: string;

  /**
   * The actual duration costs of this entry.
   *
   * @generated from field: google.protobuf.Duration costs = 7;
   */
  costs?: Duration;

  /**
   * IsVacation is set to true if the off-time-costs apply to the vacation
   * credits. If set to false, the off-time costs are for time compensation.
   *
   * @generated from field: bool is_vacation = 8;
   */
  isVacation: boolean;

  /**
   * Date holds the date of the off-time costs.
   *
   * @generated from field: google.protobuf.Timestamp date = 9;
   */
  date?: Timestamp;

  /**
   * Comment may holde an optional comment for the off-time costs.
   *
   * @generated from field: string comment = 10;
   */
  comment: string;

  constructor(data?: PartialMessage<OffTimeCosts>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.OffTimeCosts";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OffTimeCosts;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OffTimeCosts;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OffTimeCosts;

  static equals(a: OffTimeCosts | PlainMessage<OffTimeCosts> | undefined, b: OffTimeCosts | PlainMessage<OffTimeCosts> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.OffTimeEntry
 */
export declare class OffTimeEntry extends Message<OffTimeEntry> {
  /**
   * Id is a unique identifier for this off-time entry.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * RequestorId is the ID of the user that requested off-time.
   *
   * @generated from field: string requestor_id = 2;
   */
  requestorId: string;

  /**
   * The date/time (inclusive) at which the off-time should start.
   *
   * @generated from field: google.protobuf.Timestamp from = 3;
   */
  from?: Timestamp;

  /**
   * The date/time (inclusive) at which the off-time should end.
   *
   * @generated from field: google.protobuf.Timestamp to = 4;
   */
  to?: Timestamp;

  /**
   * An optional description for management in Markdown format.
   *
   * @generated from field: string description = 5;
   */
  description: string;

  /**
   * The type of the off-time request.
   *
   * @generated from field: tkd.roster.v1.OffTimeType type = 6;
   */
  type: OffTimeType;

  /**
   * If approved/rejected by management, this field will be poluated.
   *
   * @generated from field: tkd.roster.v1.OffTimeApproval approval = 7;
   */
  approval?: OffTimeApproval;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 8;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: string creator_id = 9;
   */
  creatorId: string;

  constructor(data?: PartialMessage<OffTimeEntry>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.OffTimeEntry";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OffTimeEntry;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OffTimeEntry;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OffTimeEntry;

  static equals(a: OffTimeEntry | PlainMessage<OffTimeEntry> | undefined, b: OffTimeEntry | PlainMessage<OffTimeEntry> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetOffTimeEntryRequest
 */
export declare class GetOffTimeEntryRequest extends Message<GetOffTimeEntryRequest> {
  /**
   * A list of OffTimeEntry IDs to load.
   *
   * @generated from field: repeated string ids = 1;
   */
  ids: string[];

  constructor(data?: PartialMessage<GetOffTimeEntryRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetOffTimeEntryRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetOffTimeEntryRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetOffTimeEntryRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetOffTimeEntryRequest;

  static equals(a: GetOffTimeEntryRequest | PlainMessage<GetOffTimeEntryRequest> | undefined, b: GetOffTimeEntryRequest | PlainMessage<GetOffTimeEntryRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetOffTimeEntryResponse
 */
export declare class GetOffTimeEntryResponse extends Message<GetOffTimeEntryResponse> {
  /**
   * @generated from field: repeated tkd.roster.v1.OffTimeEntry entry = 1;
   */
  entry: OffTimeEntry[];

  constructor(data?: PartialMessage<GetOffTimeEntryResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetOffTimeEntryResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetOffTimeEntryResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetOffTimeEntryResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetOffTimeEntryResponse;

  static equals(a: GetOffTimeEntryResponse | PlainMessage<GetOffTimeEntryResponse> | undefined, b: GetOffTimeEntryResponse | PlainMessage<GetOffTimeEntryResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.CreateOffTimeRequestRequest
 */
export declare class CreateOffTimeRequestRequest extends Message<CreateOffTimeRequestRequest> {
  /**
   * @generated from field: google.protobuf.Timestamp from = 1;
   */
  from?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp to = 2;
   */
  to?: Timestamp;

  /**
   * The ID of the user that requests off-time. Only administrators may
   * set this field.
   * For non-administrators, or if the field is empty, it defaults to the
   * user-id that performs the request.
   *
   * @generated from field: string requestor_id = 3;
   */
  requestorId: string;

  /**
   * An optional description of the off-time-request.
   *
   * @generated from field: string description = 4;
   */
  description: string;

  /**
   * The type of off-time request.
   *
   * @generated from field: tkd.roster.v1.OffTimeType request_type = 5;
   */
  requestType: OffTimeType;

  constructor(data?: PartialMessage<CreateOffTimeRequestRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.CreateOffTimeRequestRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateOffTimeRequestRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateOffTimeRequestRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateOffTimeRequestRequest;

  static equals(a: CreateOffTimeRequestRequest | PlainMessage<CreateOffTimeRequestRequest> | undefined, b: CreateOffTimeRequestRequest | PlainMessage<CreateOffTimeRequestRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.CreateOffTimeRequestResponse
 */
export declare class CreateOffTimeRequestResponse extends Message<CreateOffTimeRequestResponse> {
  /**
   * Holds the created off-time request entry.
   *
   * @generated from field: tkd.roster.v1.OffTimeEntry entry = 1;
   */
  entry?: OffTimeEntry;

  constructor(data?: PartialMessage<CreateOffTimeRequestResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.CreateOffTimeRequestResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateOffTimeRequestResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateOffTimeRequestResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateOffTimeRequestResponse;

  static equals(a: CreateOffTimeRequestResponse | PlainMessage<CreateOffTimeRequestResponse> | undefined, b: CreateOffTimeRequestResponse | PlainMessage<CreateOffTimeRequestResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UpdateOffTimeRequestRequest
 */
export declare class UpdateOffTimeRequestRequest extends Message<UpdateOffTimeRequestRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: google.protobuf.Timestamp from = 2;
   */
  from?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp to = 3;
   */
  to?: Timestamp;

  /**
   * @generated from field: string requestor_id = 4;
   */
  requestorId: string;

  /**
   * @generated from field: string description = 5;
   */
  description: string;

  /**
   * @generated from field: tkd.roster.v1.OffTimeType request_type = 6;
   */
  requestType: OffTimeType;

  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 7;
   */
  fieldMask?: FieldMask;

  constructor(data?: PartialMessage<UpdateOffTimeRequestRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UpdateOffTimeRequestRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOffTimeRequestRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOffTimeRequestRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOffTimeRequestRequest;

  static equals(a: UpdateOffTimeRequestRequest | PlainMessage<UpdateOffTimeRequestRequest> | undefined, b: UpdateOffTimeRequestRequest | PlainMessage<UpdateOffTimeRequestRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UpdateOffTimeRequestResponse
 */
export declare class UpdateOffTimeRequestResponse extends Message<UpdateOffTimeRequestResponse> {
  /**
   * @generated from field: tkd.roster.v1.OffTimeEntry entry = 1;
   */
  entry?: OffTimeEntry;

  constructor(data?: PartialMessage<UpdateOffTimeRequestResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UpdateOffTimeRequestResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateOffTimeRequestResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateOffTimeRequestResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateOffTimeRequestResponse;

  static equals(a: UpdateOffTimeRequestResponse | PlainMessage<UpdateOffTimeRequestResponse> | undefined, b: UpdateOffTimeRequestResponse | PlainMessage<UpdateOffTimeRequestResponse> | undefined): boolean;
}

/**
 * Request deletion of one or more off-time-request.
 *
 * @generated from message tkd.roster.v1.DeleteOffTimeRequestRequest
 */
export declare class DeleteOffTimeRequestRequest extends Message<DeleteOffTimeRequestRequest> {
  /**
   * @generated from field: repeated string id = 1;
   */
  id: string[];

  constructor(data?: PartialMessage<DeleteOffTimeRequestRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteOffTimeRequestRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteOffTimeRequestRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteOffTimeRequestRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteOffTimeRequestRequest;

  static equals(a: DeleteOffTimeRequestRequest | PlainMessage<DeleteOffTimeRequestRequest> | undefined, b: DeleteOffTimeRequestRequest | PlainMessage<DeleteOffTimeRequestRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteOffTimeRequestResponse
 */
export declare class DeleteOffTimeRequestResponse extends Message<DeleteOffTimeRequestResponse> {
  constructor(data?: PartialMessage<DeleteOffTimeRequestResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteOffTimeRequestResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteOffTimeRequestResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteOffTimeRequestResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteOffTimeRequestResponse;

  static equals(a: DeleteOffTimeRequestResponse | PlainMessage<DeleteOffTimeRequestResponse> | undefined, b: DeleteOffTimeRequestResponse | PlainMessage<DeleteOffTimeRequestResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.FindOffTimeRequestsRequest
 */
export declare class FindOffTimeRequestsRequest extends Message<FindOffTimeRequestsRequest> {
  /**
   * @generated from field: google.protobuf.Timestamp from = 1;
   */
  from?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp to = 2;
   */
  to?: Timestamp;

  /**
   * @generated from field: repeated string user_ids = 3;
   */
  userIds: string[];

  /**
   * @generated from field: google.protobuf.BoolValue approved = 4;
   */
  approved?: boolean;

  /**
   * @generated from field: google.protobuf.FieldMask read_mask = 5;
   */
  readMask?: FieldMask;

  constructor(data?: PartialMessage<FindOffTimeRequestsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.FindOffTimeRequestsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FindOffTimeRequestsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FindOffTimeRequestsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FindOffTimeRequestsRequest;

  static equals(a: FindOffTimeRequestsRequest | PlainMessage<FindOffTimeRequestsRequest> | undefined, b: FindOffTimeRequestsRequest | PlainMessage<FindOffTimeRequestsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.FindOffTimeRequestsResponse
 */
export declare class FindOffTimeRequestsResponse extends Message<FindOffTimeRequestsResponse> {
  /**
   * @generated from field: repeated tkd.roster.v1.OffTimeEntry results = 1;
   */
  results: OffTimeEntry[];

  constructor(data?: PartialMessage<FindOffTimeRequestsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.FindOffTimeRequestsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FindOffTimeRequestsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FindOffTimeRequestsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FindOffTimeRequestsResponse;

  static equals(a: FindOffTimeRequestsResponse | PlainMessage<FindOffTimeRequestsResponse> | undefined, b: FindOffTimeRequestsResponse | PlainMessage<FindOffTimeRequestsResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.ApproveOrRejectRequest
 */
export declare class ApproveOrRejectRequest extends Message<ApproveOrRejectRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string comment = 2;
   */
  comment: string;

  /**
   * @generated from field: tkd.roster.v1.ApprovalRequestType type = 3;
   */
  type: ApprovalRequestType;

  constructor(data?: PartialMessage<ApproveOrRejectRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.ApproveOrRejectRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApproveOrRejectRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApproveOrRejectRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApproveOrRejectRequest;

  static equals(a: ApproveOrRejectRequest | PlainMessage<ApproveOrRejectRequest> | undefined, b: ApproveOrRejectRequest | PlainMessage<ApproveOrRejectRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.ApproveOrRejectResponse
 */
export declare class ApproveOrRejectResponse extends Message<ApproveOrRejectResponse> {
  /**
   * @generated from field: tkd.roster.v1.OffTimeEntry entry = 1;
   */
  entry?: OffTimeEntry;

  constructor(data?: PartialMessage<ApproveOrRejectResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.ApproveOrRejectResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApproveOrRejectResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApproveOrRejectResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApproveOrRejectResponse;

  static equals(a: ApproveOrRejectResponse | PlainMessage<ApproveOrRejectResponse> | undefined, b: ApproveOrRejectResponse | PlainMessage<ApproveOrRejectResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.AddOffTimeCostsRequest
 */
export declare class AddOffTimeCostsRequest extends Message<AddOffTimeCostsRequest> {
  /**
   * @generated from field: repeated tkd.roster.v1.OffTimeCosts add_costs = 1;
   */
  addCosts: OffTimeCosts[];

  constructor(data?: PartialMessage<AddOffTimeCostsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.AddOffTimeCostsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddOffTimeCostsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddOffTimeCostsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddOffTimeCostsRequest;

  static equals(a: AddOffTimeCostsRequest | PlainMessage<AddOffTimeCostsRequest> | undefined, b: AddOffTimeCostsRequest | PlainMessage<AddOffTimeCostsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.AddOffTimeCostsResponse
 */
export declare class AddOffTimeCostsResponse extends Message<AddOffTimeCostsResponse> {
  constructor(data?: PartialMessage<AddOffTimeCostsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.AddOffTimeCostsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddOffTimeCostsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddOffTimeCostsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddOffTimeCostsResponse;

  static equals(a: AddOffTimeCostsResponse | PlainMessage<AddOffTimeCostsResponse> | undefined, b: AddOffTimeCostsResponse | PlainMessage<AddOffTimeCostsResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.CostsForUsers
 */
export declare class CostsForUsers extends Message<CostsForUsers> {
  /**
   * @generated from field: repeated string user_ids = 1;
   */
  userIds: string[];

  constructor(data?: PartialMessage<CostsForUsers>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.CostsForUsers";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CostsForUsers;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CostsForUsers;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CostsForUsers;

  static equals(a: CostsForUsers | PlainMessage<CostsForUsers> | undefined, b: CostsForUsers | PlainMessage<CostsForUsers> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetOffTimeCostsRequest
 */
export declare class GetOffTimeCostsRequest extends Message<GetOffTimeCostsRequest> {
  /**
   * @generated from field: tkd.roster.v1.CostsForUsers for_users = 1;
   */
  forUsers?: CostsForUsers;

  /**
   * @generated from field: google.protobuf.FieldMask read_mask = 2;
   */
  readMask?: FieldMask;

  constructor(data?: PartialMessage<GetOffTimeCostsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetOffTimeCostsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetOffTimeCostsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetOffTimeCostsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetOffTimeCostsRequest;

  static equals(a: GetOffTimeCostsRequest | PlainMessage<GetOffTimeCostsRequest> | undefined, b: GetOffTimeCostsRequest | PlainMessage<GetOffTimeCostsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.OffTimeCostSummary
 */
export declare class OffTimeCostSummary extends Message<OffTimeCostSummary> {
  /**
   * @generated from field: google.protobuf.Duration vacation = 1;
   */
  vacation?: Duration;

  /**
   * @generated from field: google.protobuf.Duration time_off = 2;
   */
  timeOff?: Duration;

  constructor(data?: PartialMessage<OffTimeCostSummary>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.OffTimeCostSummary";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OffTimeCostSummary;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OffTimeCostSummary;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OffTimeCostSummary;

  static equals(a: OffTimeCostSummary | PlainMessage<OffTimeCostSummary> | undefined, b: OffTimeCostSummary | PlainMessage<OffTimeCostSummary> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UserOffTimeCosts
 */
export declare class UserOffTimeCosts extends Message<UserOffTimeCosts> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * @generated from field: repeated tkd.roster.v1.OffTimeCosts costs = 2;
   */
  costs: OffTimeCosts[];

  /**
   * @generated from field: tkd.roster.v1.OffTimeCostSummary summary = 3;
   */
  summary?: OffTimeCostSummary;

  constructor(data?: PartialMessage<UserOffTimeCosts>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UserOffTimeCosts";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserOffTimeCosts;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserOffTimeCosts;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserOffTimeCosts;

  static equals(a: UserOffTimeCosts | PlainMessage<UserOffTimeCosts> | undefined, b: UserOffTimeCosts | PlainMessage<UserOffTimeCosts> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetOffTimeCostsResponse
 */
export declare class GetOffTimeCostsResponse extends Message<GetOffTimeCostsResponse> {
  /**
   * @generated from field: repeated tkd.roster.v1.UserOffTimeCosts results = 1;
   */
  results: UserOffTimeCosts[];

  constructor(data?: PartialMessage<GetOffTimeCostsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetOffTimeCostsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetOffTimeCostsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetOffTimeCostsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetOffTimeCostsResponse;

  static equals(a: GetOffTimeCostsResponse | PlainMessage<GetOffTimeCostsResponse> | undefined, b: GetOffTimeCostsResponse | PlainMessage<GetOffTimeCostsResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteOffTimeCostsRequest
 */
export declare class DeleteOffTimeCostsRequest extends Message<DeleteOffTimeCostsRequest> {
  /**
   * @generated from field: repeated string ids = 1;
   */
  ids: string[];

  constructor(data?: PartialMessage<DeleteOffTimeCostsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteOffTimeCostsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteOffTimeCostsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteOffTimeCostsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteOffTimeCostsRequest;

  static equals(a: DeleteOffTimeCostsRequest | PlainMessage<DeleteOffTimeCostsRequest> | undefined, b: DeleteOffTimeCostsRequest | PlainMessage<DeleteOffTimeCostsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteOffTimeCostsResponse
 */
export declare class DeleteOffTimeCostsResponse extends Message<DeleteOffTimeCostsResponse> {
  constructor(data?: PartialMessage<DeleteOffTimeCostsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteOffTimeCostsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteOffTimeCostsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteOffTimeCostsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteOffTimeCostsResponse;

  static equals(a: DeleteOffTimeCostsResponse | PlainMessage<DeleteOffTimeCostsResponse> | undefined, b: DeleteOffTimeCostsResponse | PlainMessage<DeleteOffTimeCostsResponse> | undefined): boolean;
}

