// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file tkd/pbx3cx/v1/calllog.proto (package tkd.pbx3cx.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, Duration, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Profile } from "../../idm/v1/user_pb.js";
import type { TimeRange } from "../../common/v1/time_range_pb.js";
import type { Customer, CustomerRef } from "../../customer/v1/customer_pb.js";

/**
 * @generated from message tkd.pbx3cx.v1.CallEntry
 */
export declare class CallEntry extends Message<CallEntry> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string caller = 2;
   */
  caller: string;

  /**
   * @generated from field: string inbound_number = 3;
   */
  inboundNumber: string;

  /**
   * @generated from field: google.protobuf.Timestamp received_at = 4;
   */
  receivedAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Duration duration = 5;
   */
  duration?: Duration;

  /**
   * @generated from field: string call_type = 6;
   */
  callType: string;

  /**
   * @generated from field: string agent_user_id = 7;
   */
  agentUserId: string;

  /**
   * @generated from field: string customer_id = 8;
   */
  customerId: string;

  /**
   * CustomerSource is unused.
   * Deprecated: CustomerSource is not used anymore.
   *
   * @generated from field: string customer_source = 9;
   */
  customerSource: string;

  /**
   * @generated from field: bool error = 10;
   */
  error: boolean;

  /**
   * @generated from field: string transfer_target = 11;
   */
  transferTarget: string;

  /**
   * @generated from field: string accepted_agent = 12;
   */
  acceptedAgent: string;

  constructor(data?: PartialMessage<CallEntry>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.CallEntry";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CallEntry;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CallEntry;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CallEntry;

  static equals(a: CallEntry | PlainMessage<CallEntry> | undefined, b: CallEntry | PlainMessage<CallEntry> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.OnCall
 */
export declare class OnCall extends Message<OnCall> {
  /**
   * @generated from field: tkd.idm.v1.Profile profile = 1;
   */
  profile?: Profile;

  /**
   * @generated from field: string transfer_target = 2;
   */
  transferTarget: string;

  /**
   * @generated from field: google.protobuf.Timestamp until = 3;
   */
  until?: Timestamp;

  constructor(data?: PartialMessage<OnCall>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.OnCall";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OnCall;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OnCall;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OnCall;

  static equals(a: OnCall | PlainMessage<OnCall> | undefined, b: OnCall | PlainMessage<OnCall> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.CustomOverwrite
 */
export declare class CustomOverwrite extends Message<CustomOverwrite> {
  /**
   * @generated from field: string transfer_target = 1;
   */
  transferTarget: string;

  /**
   * @generated from field: string display_name = 2;
   */
  displayName: string;

  constructor(data?: PartialMessage<CustomOverwrite>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.CustomOverwrite";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CustomOverwrite;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CustomOverwrite;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CustomOverwrite;

  static equals(a: CustomOverwrite | PlainMessage<CustomOverwrite> | undefined, b: CustomOverwrite | PlainMessage<CustomOverwrite> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.Overwrite
 */
export declare class Overwrite extends Message<Overwrite> {
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
   * @generated from oneof tkd.pbx3cx.v1.Overwrite.target
   */
  target: {
    /**
     * @generated from field: tkd.pbx3cx.v1.CustomOverwrite custom = 4;
     */
    value: CustomOverwrite;
    case: "custom";
  } | {
    /**
     * @generated from field: string user_id = 5;
     */
    value: string;
    case: "userId";
  } | { case: undefined; value?: undefined };

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 6;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: string created_by_user_id = 7;
   */
  createdByUserId: string;

  /**
   * @generated from field: tkd.pbx3cx.v1.InboundNumber inbound_number = 8;
   */
  inboundNumber?: InboundNumber;

  constructor(data?: PartialMessage<Overwrite>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.Overwrite";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Overwrite;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Overwrite;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Overwrite;

  static equals(a: Overwrite | PlainMessage<Overwrite> | undefined, b: Overwrite | PlainMessage<Overwrite> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.InboundNumber
 */
export declare class InboundNumber extends Message<InboundNumber> {
  /**
   * Number is the inbound number.
   *
   * @generated from field: string number = 1;
   */
  number: string;

  /**
   * DisplayName is an optional display-name or description for
   * the inbound number.
   *
   * @generated from field: string display_name = 2;
   */
  displayName: string;

  /**
   * RosterShiftTags is a list of roster shift tags that should
   * be queried when resolving the employees that are currently on-duty.
   *
   * @generated from field: repeated string roster_shift_tags = 3;
   */
  rosterShiftTags: string[];

  /**
   * An optional roster type name to query when resolving the employees
   * that are currently on duty.
   * If unset, CallService will use the default RosterTypeName from it's
   * configuration.
   *
   * @generated from field: string roster_type_name = 4;
   */
  rosterTypeName: string;

  constructor(data?: PartialMessage<InboundNumber>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.InboundNumber";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InboundNumber;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InboundNumber;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InboundNumber;

  static equals(a: InboundNumber | PlainMessage<InboundNumber> | undefined, b: InboundNumber | PlainMessage<InboundNumber> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.RecordCallRequest
 */
export declare class RecordCallRequest extends Message<RecordCallRequest> {
  /**
   * @generated from field: string duration = 1;
   */
  duration: string;

  /**
   * @generated from field: string number = 2;
   */
  number: string;

  /**
   * @generated from field: string agent = 3;
   */
  agent: string;

  /**
   * @generated from field: string call_type = 4;
   */
  callType: string;

  /**
   * @generated from field: string date_time = 5;
   */
  dateTime: string;

  /**
   * @generated from field: string customer_id = 6;
   */
  customerId: string;

  /**
   * CustomerSource used to hold the source of the customer.
   * Deprecated: this field is only set for old records
   *
   * @generated from field: string customer_source = 7;
   */
  customerSource: string;

  constructor(data?: PartialMessage<RecordCallRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.RecordCallRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RecordCallRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RecordCallRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RecordCallRequest;

  static equals(a: RecordCallRequest | PlainMessage<RecordCallRequest> | undefined, b: RecordCallRequest | PlainMessage<RecordCallRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForCustomerRequest
 */
export declare class GetLogsForCustomerRequest extends Message<GetLogsForCustomerRequest> {
  /**
   * Deprecated
   *
   * @generated from field: string source = 1;
   */
  source: string;

  /**
   * @generated from field: string id = 2;
   */
  id: string;

  constructor(data?: PartialMessage<GetLogsForCustomerRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.GetLogsForCustomerRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLogsForCustomerRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLogsForCustomerRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLogsForCustomerRequest;

  static equals(a: GetLogsForCustomerRequest | PlainMessage<GetLogsForCustomerRequest> | undefined, b: GetLogsForCustomerRequest | PlainMessage<GetLogsForCustomerRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForCustomerResponse
 */
export declare class GetLogsForCustomerResponse extends Message<GetLogsForCustomerResponse> {
  /**
   * @generated from field: repeated tkd.pbx3cx.v1.CallEntry results = 1;
   */
  results: CallEntry[];

  constructor(data?: PartialMessage<GetLogsForCustomerResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.GetLogsForCustomerResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLogsForCustomerResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLogsForCustomerResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLogsForCustomerResponse;

  static equals(a: GetLogsForCustomerResponse | PlainMessage<GetLogsForCustomerResponse> | undefined, b: GetLogsForCustomerResponse | PlainMessage<GetLogsForCustomerResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForDateRequest
 */
export declare class GetLogsForDateRequest extends Message<GetLogsForDateRequest> {
  /**
   * YYYY-MM-DD
   *
   * @generated from field: string date = 1;
   */
  date: string;

  constructor(data?: PartialMessage<GetLogsForDateRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.GetLogsForDateRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLogsForDateRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLogsForDateRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLogsForDateRequest;

  static equals(a: GetLogsForDateRequest | PlainMessage<GetLogsForDateRequest> | undefined, b: GetLogsForDateRequest | PlainMessage<GetLogsForDateRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.GetLogsForDateResponse
 */
export declare class GetLogsForDateResponse extends Message<GetLogsForDateResponse> {
  /**
   * @generated from field: repeated tkd.pbx3cx.v1.CallEntry results = 1;
   */
  results: CallEntry[];

  constructor(data?: PartialMessage<GetLogsForDateResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.GetLogsForDateResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLogsForDateResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLogsForDateResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLogsForDateResponse;

  static equals(a: GetLogsForDateResponse | PlainMessage<GetLogsForDateResponse> | undefined, b: GetLogsForDateResponse | PlainMessage<GetLogsForDateResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.SearchCallLogsRequest
 */
export declare class SearchCallLogsRequest extends Message<SearchCallLogsRequest> {
  /**
   * YYYY-MM-DD
   *
   * @generated from field: string date = 1;
   */
  date: string;

  /**
   * @generated from field: tkd.common.v1.TimeRange time_range = 2;
   */
  timeRange?: TimeRange;

  /**
   * @generated from field: tkd.customer.v1.CustomerRef customer_ref = 3;
   */
  customerRef?: CustomerRef;

  constructor(data?: PartialMessage<SearchCallLogsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.SearchCallLogsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchCallLogsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchCallLogsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchCallLogsRequest;

  static equals(a: SearchCallLogsRequest | PlainMessage<SearchCallLogsRequest> | undefined, b: SearchCallLogsRequest | PlainMessage<SearchCallLogsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.SearchCallLogsResponse
 */
export declare class SearchCallLogsResponse extends Message<SearchCallLogsResponse> {
  /**
   * @generated from field: repeated tkd.pbx3cx.v1.CallEntry results = 1;
   */
  results: CallEntry[];

  /**
   * @generated from field: repeated tkd.customer.v1.Customer customers = 2;
   */
  customers: Customer[];

  constructor(data?: PartialMessage<SearchCallLogsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.SearchCallLogsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchCallLogsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchCallLogsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchCallLogsResponse;

  static equals(a: SearchCallLogsResponse | PlainMessage<SearchCallLogsResponse> | undefined, b: SearchCallLogsResponse | PlainMessage<SearchCallLogsResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.GetOnCallRequest
 */
export declare class GetOnCallRequest extends Message<GetOnCallRequest> {
  /**
   * YYYY-MM-DD
   *
   * @generated from field: string date = 1;
   */
  date: string;

  /**
   * @generated from field: bool ignore_overwrites = 2;
   */
  ignoreOverwrites: boolean;

  /**
   * @generated from field: string inbound_number = 3;
   */
  inboundNumber: string;

  constructor(data?: PartialMessage<GetOnCallRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.GetOnCallRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetOnCallRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetOnCallRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetOnCallRequest;

  static equals(a: GetOnCallRequest | PlainMessage<GetOnCallRequest> | undefined, b: GetOnCallRequest | PlainMessage<GetOnCallRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.GetOnCallResponse
 */
export declare class GetOnCallResponse extends Message<GetOnCallResponse> {
  /**
   * @generated from field: repeated tkd.pbx3cx.v1.OnCall on_call = 1;
   */
  onCall: OnCall[];

  /**
   * YYYY-MM-DD
   *
   * @generated from field: string roster_date = 2;
   */
  rosterDate: string;

  /**
   * @generated from field: bool is_overwrite = 3;
   */
  isOverwrite: boolean;

  /**
   * @generated from field: string primary_transfer_target = 4;
   */
  primaryTransferTarget: string;

  constructor(data?: PartialMessage<GetOnCallResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.GetOnCallResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetOnCallResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetOnCallResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetOnCallResponse;

  static equals(a: GetOnCallResponse | PlainMessage<GetOnCallResponse> | undefined, b: GetOnCallResponse | PlainMessage<GetOnCallResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.CreateOverwriteRequest
 */
export declare class CreateOverwriteRequest extends Message<CreateOverwriteRequest> {
  /**
   * @generated from oneof tkd.pbx3cx.v1.CreateOverwriteRequest.transfer_target
   */
  transferTarget: {
    /**
     * @generated from field: tkd.pbx3cx.v1.CustomOverwrite custom = 1;
     */
    value: CustomOverwrite;
    case: "custom";
  } | {
    /**
     * @generated from field: string user_id = 2;
     */
    value: string;
    case: "userId";
  } | { case: undefined; value?: undefined };

  /**
   * @generated from field: google.protobuf.Timestamp from = 3;
   */
  from?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp to = 4;
   */
  to?: Timestamp;

  /**
   * @generated from field: string inbound_number = 5;
   */
  inboundNumber: string;

  constructor(data?: PartialMessage<CreateOverwriteRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.CreateOverwriteRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateOverwriteRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateOverwriteRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateOverwriteRequest;

  static equals(a: CreateOverwriteRequest | PlainMessage<CreateOverwriteRequest> | undefined, b: CreateOverwriteRequest | PlainMessage<CreateOverwriteRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.CreateOverwriteResponse
 */
export declare class CreateOverwriteResponse extends Message<CreateOverwriteResponse> {
  /**
   * @generated from field: tkd.pbx3cx.v1.Overwrite overwrite = 1;
   */
  overwrite?: Overwrite;

  constructor(data?: PartialMessage<CreateOverwriteResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.CreateOverwriteResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateOverwriteResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateOverwriteResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateOverwriteResponse;

  static equals(a: CreateOverwriteResponse | PlainMessage<CreateOverwriteResponse> | undefined, b: CreateOverwriteResponse | PlainMessage<CreateOverwriteResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.DeleteOverwriteRequest
 */
export declare class DeleteOverwriteRequest extends Message<DeleteOverwriteRequest> {
  /**
   * @generated from oneof tkd.pbx3cx.v1.DeleteOverwriteRequest.selector
   */
  selector: {
    /**
     * @generated from field: string overwrite_id = 1;
     */
    value: string;
    case: "overwriteId";
  } | {
    /**
     * @generated from field: google.protobuf.Timestamp active_at = 2;
     */
    value: Timestamp;
    case: "activeAt";
  } | { case: undefined; value?: undefined };

  /**
   * @generated from field: tkd.pbx3cx.v1.InboundNumberList inbound_numbers = 3;
   */
  inboundNumbers?: InboundNumberList;

  constructor(data?: PartialMessage<DeleteOverwriteRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.DeleteOverwriteRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteOverwriteRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteOverwriteRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteOverwriteRequest;

  static equals(a: DeleteOverwriteRequest | PlainMessage<DeleteOverwriteRequest> | undefined, b: DeleteOverwriteRequest | PlainMessage<DeleteOverwriteRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.DeleteOverwriteResponse
 */
export declare class DeleteOverwriteResponse extends Message<DeleteOverwriteResponse> {
  constructor(data?: PartialMessage<DeleteOverwriteResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.DeleteOverwriteResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteOverwriteResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteOverwriteResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteOverwriteResponse;

  static equals(a: DeleteOverwriteResponse | PlainMessage<DeleteOverwriteResponse> | undefined, b: DeleteOverwriteResponse | PlainMessage<DeleteOverwriteResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.GetOverwriteRequest
 */
export declare class GetOverwriteRequest extends Message<GetOverwriteRequest> {
  /**
   * @generated from oneof tkd.pbx3cx.v1.GetOverwriteRequest.selector
   */
  selector: {
    /**
     * @generated from field: string overwrite_id = 1;
     */
    value: string;
    case: "overwriteId";
  } | {
    /**
     * @generated from field: google.protobuf.Timestamp active_at = 2;
     */
    value: Timestamp;
    case: "activeAt";
  } | {
    /**
     * @generated from field: tkd.common.v1.TimeRange time_range = 3;
     */
    value: TimeRange;
    case: "timeRange";
  } | { case: undefined; value?: undefined };

  /**
   * @generated from field: tkd.pbx3cx.v1.InboundNumberList inbound_numbers = 4;
   */
  inboundNumbers?: InboundNumberList;

  constructor(data?: PartialMessage<GetOverwriteRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.GetOverwriteRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetOverwriteRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetOverwriteRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetOverwriteRequest;

  static equals(a: GetOverwriteRequest | PlainMessage<GetOverwriteRequest> | undefined, b: GetOverwriteRequest | PlainMessage<GetOverwriteRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.InboundNumberList
 */
export declare class InboundNumberList extends Message<InboundNumberList> {
  /**
   * @generated from field: repeated string numbers = 1;
   */
  numbers: string[];

  constructor(data?: PartialMessage<InboundNumberList>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.InboundNumberList";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): InboundNumberList;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): InboundNumberList;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): InboundNumberList;

  static equals(a: InboundNumberList | PlainMessage<InboundNumberList> | undefined, b: InboundNumberList | PlainMessage<InboundNumberList> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.GetOverwriteResponse
 */
export declare class GetOverwriteResponse extends Message<GetOverwriteResponse> {
  /**
   * @generated from field: repeated tkd.pbx3cx.v1.Overwrite overwrites = 1;
   */
  overwrites: Overwrite[];

  constructor(data?: PartialMessage<GetOverwriteResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.GetOverwriteResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetOverwriteResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetOverwriteResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetOverwriteResponse;

  static equals(a: GetOverwriteResponse | PlainMessage<GetOverwriteResponse> | undefined, b: GetOverwriteResponse | PlainMessage<GetOverwriteResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.CreateInboundNumberRequest
 */
export declare class CreateInboundNumberRequest extends Message<CreateInboundNumberRequest> {
  /**
   * @generated from field: string number = 1;
   */
  number: string;

  /**
   * @generated from field: string display_name = 2;
   */
  displayName: string;

  /**
   * RosterShiftTags is a list of roster shift tags that should
   * be queried when resolving the employees that are currently on-duty.
   *
   * @generated from field: repeated string roster_shift_tags = 3;
   */
  rosterShiftTags: string[];

  /**
   * An optional roster type name to query when resolving the employees
   * that are currently on duty.
   * If unset, CallService will use the default RosterTypeName from it's
   * configuration.
   *
   * @generated from field: string roster_type_name = 4;
   */
  rosterTypeName: string;

  constructor(data?: PartialMessage<CreateInboundNumberRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.CreateInboundNumberRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateInboundNumberRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateInboundNumberRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateInboundNumberRequest;

  static equals(a: CreateInboundNumberRequest | PlainMessage<CreateInboundNumberRequest> | undefined, b: CreateInboundNumberRequest | PlainMessage<CreateInboundNumberRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.CreateInboundNumberResponse
 */
export declare class CreateInboundNumberResponse extends Message<CreateInboundNumberResponse> {
  /**
   * @generated from field: tkd.pbx3cx.v1.InboundNumber inbound_number = 1;
   */
  inboundNumber?: InboundNumber;

  constructor(data?: PartialMessage<CreateInboundNumberResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.CreateInboundNumberResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateInboundNumberResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateInboundNumberResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateInboundNumberResponse;

  static equals(a: CreateInboundNumberResponse | PlainMessage<CreateInboundNumberResponse> | undefined, b: CreateInboundNumberResponse | PlainMessage<CreateInboundNumberResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.DeleteInboundNumberRequest
 */
export declare class DeleteInboundNumberRequest extends Message<DeleteInboundNumberRequest> {
  /**
   * @generated from field: string number = 1;
   */
  number: string;

  constructor(data?: PartialMessage<DeleteInboundNumberRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.DeleteInboundNumberRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteInboundNumberRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteInboundNumberRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteInboundNumberRequest;

  static equals(a: DeleteInboundNumberRequest | PlainMessage<DeleteInboundNumberRequest> | undefined, b: DeleteInboundNumberRequest | PlainMessage<DeleteInboundNumberRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.DeleteInboundNumberResponse
 */
export declare class DeleteInboundNumberResponse extends Message<DeleteInboundNumberResponse> {
  constructor(data?: PartialMessage<DeleteInboundNumberResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.DeleteInboundNumberResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteInboundNumberResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteInboundNumberResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteInboundNumberResponse;

  static equals(a: DeleteInboundNumberResponse | PlainMessage<DeleteInboundNumberResponse> | undefined, b: DeleteInboundNumberResponse | PlainMessage<DeleteInboundNumberResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.ListInboundNumberRequest
 */
export declare class ListInboundNumberRequest extends Message<ListInboundNumberRequest> {
  constructor(data?: PartialMessage<ListInboundNumberRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.ListInboundNumberRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListInboundNumberRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListInboundNumberRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListInboundNumberRequest;

  static equals(a: ListInboundNumberRequest | PlainMessage<ListInboundNumberRequest> | undefined, b: ListInboundNumberRequest | PlainMessage<ListInboundNumberRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.ListInboundNumberResponse
 */
export declare class ListInboundNumberResponse extends Message<ListInboundNumberResponse> {
  /**
   * @generated from field: repeated tkd.pbx3cx.v1.InboundNumber inbound_numbers = 1;
   */
  inboundNumbers: InboundNumber[];

  constructor(data?: PartialMessage<ListInboundNumberResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.ListInboundNumberResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListInboundNumberResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListInboundNumberResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListInboundNumberResponse;

  static equals(a: ListInboundNumberResponse | PlainMessage<ListInboundNumberResponse> | undefined, b: ListInboundNumberResponse | PlainMessage<ListInboundNumberResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.UpdateInboundNumberRequest
 */
export declare class UpdateInboundNumberRequest extends Message<UpdateInboundNumberRequest> {
  /**
   * Number is the number to update.
   *
   * @generated from field: string number = 1;
   */
  number: string;

  /**
   * NewDisplayName is the new display name for the number.
   *
   * @generated from field: string new_display_name = 2;
   */
  newDisplayName: string;

  /**
   * RosterShiftTags is a list of roster shift tags that should
   * be queried when resolving the employees that are currently on-duty.
   *
   * @generated from field: repeated string roster_shift_tags = 3;
   */
  rosterShiftTags: string[];

  /**
   * An optional roster type name to query when resolving the employees
   * that are currently on duty.
   * If unset, CallService will use the default RosterTypeName from it's
   * configuration.
   *
   * @generated from field: string roster_type_name = 4;
   */
  rosterTypeName: string;

  /**
   * UpdateMask specifies which fields of the InboundNumber should be updated.
   *
   * @generated from field: google.protobuf.FieldMask update_mask = 10;
   */
  updateMask?: FieldMask;

  constructor(data?: PartialMessage<UpdateInboundNumberRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.UpdateInboundNumberRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateInboundNumberRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateInboundNumberRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateInboundNumberRequest;

  static equals(a: UpdateInboundNumberRequest | PlainMessage<UpdateInboundNumberRequest> | undefined, b: UpdateInboundNumberRequest | PlainMessage<UpdateInboundNumberRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.pbx3cx.v1.UpdateInboundNumberResponse
 */
export declare class UpdateInboundNumberResponse extends Message<UpdateInboundNumberResponse> {
  /**
   * @generated from field: tkd.pbx3cx.v1.InboundNumber inbound_number = 1;
   */
  inboundNumber?: InboundNumber;

  constructor(data?: PartialMessage<UpdateInboundNumberResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.pbx3cx.v1.UpdateInboundNumberResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateInboundNumberResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateInboundNumberResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateInboundNumberResponse;

  static equals(a: UpdateInboundNumberResponse | PlainMessage<UpdateInboundNumberResponse> | undefined, b: UpdateInboundNumberResponse | PlainMessage<UpdateInboundNumberResponse> | undefined): boolean;
}

