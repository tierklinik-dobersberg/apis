// @generated by protoc-gen-es v1.5.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/workshift.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, Duration, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message tkd.roster.v1.Daytime
 */
export declare class Daytime extends Message<Daytime> {
  /**
   * @generated from field: int64 hour = 1;
   */
  hour: bigint;

  /**
   * @generated from field: int64 minute = 2;
   */
  minute: bigint;

  constructor(data?: PartialMessage<Daytime>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.Daytime";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Daytime;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Daytime;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Daytime;

  static equals(a: Daytime | PlainMessage<Daytime> | undefined, b: Daytime | PlainMessage<Daytime> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.WorkShift
 */
export declare class WorkShift extends Message<WorkShift> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: tkd.roster.v1.Daytime from = 2;
   */
  from?: Daytime;

  /**
   * @generated from field: google.protobuf.Duration duration = 3;
   */
  duration?: Duration;

  /**
   * 0 = Sunday, 6 = Saturday
   *
   * @generated from field: repeated int32 days = 4;
   */
  days: number[];

  /**
   * @generated from field: string name = 5;
   */
  name: string;

  /**
   * @generated from field: string display_name = 6;
   */
  displayName: string;

  /**
   * @generated from field: bool on_holiday = 7;
   */
  onHoliday: boolean;

  /**
   * @generated from field: repeated string eligible_role_ids = 8;
   */
  eligibleRoleIds: string[];

  /**
   * @generated from field: google.protobuf.Duration time_worth = 9;
   */
  timeWorth?: Duration;

  /**
   * @generated from field: int64 required_staff_count = 10;
   */
  requiredStaffCount: bigint;

  /**
   * @generated from field: string color = 11;
   */
  color: string;

  /**
   * @generated from field: string description = 12;
   */
  description: string;

  /**
   * @generated from field: int64 order = 13;
   */
  order: bigint;

  /**
   * @generated from field: repeated string tags = 14;
   */
  tags: string[];

  constructor(data?: PartialMessage<WorkShift>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.WorkShift";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkShift;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkShift;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkShift;

  static equals(a: WorkShift | PlainMessage<WorkShift> | undefined, b: WorkShift | PlainMessage<WorkShift> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.ListWorkShiftsRequest
 */
export declare class ListWorkShiftsRequest extends Message<ListWorkShiftsRequest> {
  /**
   * @generated from field: google.protobuf.FieldMask read_mask = 1;
   */
  readMask?: FieldMask;

  constructor(data?: PartialMessage<ListWorkShiftsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.ListWorkShiftsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListWorkShiftsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListWorkShiftsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListWorkShiftsRequest;

  static equals(a: ListWorkShiftsRequest | PlainMessage<ListWorkShiftsRequest> | undefined, b: ListWorkShiftsRequest | PlainMessage<ListWorkShiftsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.ListWorkShiftsResponse
 */
export declare class ListWorkShiftsResponse extends Message<ListWorkShiftsResponse> {
  /**
   * @generated from field: repeated tkd.roster.v1.WorkShift work_shifts = 1;
   */
  workShifts: WorkShift[];

  constructor(data?: PartialMessage<ListWorkShiftsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.ListWorkShiftsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListWorkShiftsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListWorkShiftsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListWorkShiftsResponse;

  static equals(a: ListWorkShiftsResponse | PlainMessage<ListWorkShiftsResponse> | undefined, b: ListWorkShiftsResponse | PlainMessage<ListWorkShiftsResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.CreateWorkShiftRequest
 */
export declare class CreateWorkShiftRequest extends Message<CreateWorkShiftRequest> {
  /**
   * @generated from field: tkd.roster.v1.Daytime from = 1;
   */
  from?: Daytime;

  /**
   * @generated from field: google.protobuf.Duration duration = 2;
   */
  duration?: Duration;

  /**
   * @generated from field: repeated int32 days = 3;
   */
  days: number[];

  /**
   * @generated from field: string name = 4;
   */
  name: string;

  /**
   * @generated from field: string display_name = 5;
   */
  displayName: string;

  /**
   * @generated from field: bool on_holiday = 6;
   */
  onHoliday: boolean;

  /**
   * @generated from field: repeated string eligible_role_ids = 7;
   */
  eligibleRoleIds: string[];

  /**
   * @generated from field: google.protobuf.Duration time_worth = 8;
   */
  timeWorth?: Duration;

  /**
   * @generated from field: int64 required_staff_count = 9;
   */
  requiredStaffCount: bigint;

  /**
   * @generated from field: string color = 10;
   */
  color: string;

  /**
   * @generated from field: string description = 11;
   */
  description: string;

  /**
   * @generated from field: int64 order = 12;
   */
  order: bigint;

  /**
   * @generated from field: repeated string tags = 13;
   */
  tags: string[];

  constructor(data?: PartialMessage<CreateWorkShiftRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.CreateWorkShiftRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateWorkShiftRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateWorkShiftRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateWorkShiftRequest;

  static equals(a: CreateWorkShiftRequest | PlainMessage<CreateWorkShiftRequest> | undefined, b: CreateWorkShiftRequest | PlainMessage<CreateWorkShiftRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.CreateWorkShiftResponse
 */
export declare class CreateWorkShiftResponse extends Message<CreateWorkShiftResponse> {
  /**
   * @generated from field: tkd.roster.v1.WorkShift work_shift = 1;
   */
  workShift?: WorkShift;

  constructor(data?: PartialMessage<CreateWorkShiftResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.CreateWorkShiftResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateWorkShiftResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateWorkShiftResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateWorkShiftResponse;

  static equals(a: CreateWorkShiftResponse | PlainMessage<CreateWorkShiftResponse> | undefined, b: CreateWorkShiftResponse | PlainMessage<CreateWorkShiftResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.WorkShiftUpdate
 */
export declare class WorkShiftUpdate extends Message<WorkShiftUpdate> {
  /**
   * @generated from field: tkd.roster.v1.Daytime from = 1;
   */
  from?: Daytime;

  /**
   * @generated from field: google.protobuf.Duration duration = 2;
   */
  duration?: Duration;

  /**
   * @generated from field: repeated int32 days = 3;
   */
  days: number[];

  /**
   * @generated from field: string name = 4;
   */
  name: string;

  /**
   * @generated from field: string display_name = 5;
   */
  displayName: string;

  /**
   * @generated from field: bool on_holiday = 6;
   */
  onHoliday: boolean;

  /**
   * @generated from field: repeated string eligible_role_ids = 7;
   */
  eligibleRoleIds: string[];

  /**
   * @generated from field: google.protobuf.Duration time_worth = 8;
   */
  timeWorth?: Duration;

  /**
   * @generated from field: int64 required_staff_count = 9;
   */
  requiredStaffCount: bigint;

  /**
   * @generated from field: string color = 10;
   */
  color: string;

  /**
   * @generated from field: string description = 11;
   */
  description: string;

  /**
   * @generated from field: int64 order = 12;
   */
  order: bigint;

  /**
   * @generated from field: repeated string tags = 13;
   */
  tags: string[];

  constructor(data?: PartialMessage<WorkShiftUpdate>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.WorkShiftUpdate";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkShiftUpdate;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkShiftUpdate;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkShiftUpdate;

  static equals(a: WorkShiftUpdate | PlainMessage<WorkShiftUpdate> | undefined, b: WorkShiftUpdate | PlainMessage<WorkShiftUpdate> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UpdateWorkShiftRequest
 */
export declare class UpdateWorkShiftRequest extends Message<UpdateWorkShiftRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: tkd.roster.v1.WorkShiftUpdate update = 2;
   */
  update?: WorkShiftUpdate;

  /**
   * @generated from field: bool update_in_place = 3;
   */
  updateInPlace: boolean;

  /**
   * @generated from field: google.protobuf.FieldMask write_mask = 4;
   */
  writeMask?: FieldMask;

  constructor(data?: PartialMessage<UpdateWorkShiftRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UpdateWorkShiftRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateWorkShiftRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateWorkShiftRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateWorkShiftRequest;

  static equals(a: UpdateWorkShiftRequest | PlainMessage<UpdateWorkShiftRequest> | undefined, b: UpdateWorkShiftRequest | PlainMessage<UpdateWorkShiftRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UpdateWorkShiftResponse
 */
export declare class UpdateWorkShiftResponse extends Message<UpdateWorkShiftResponse> {
  /**
   * @generated from field: tkd.roster.v1.WorkShift work_shift = 1;
   */
  workShift?: WorkShift;

  constructor(data?: PartialMessage<UpdateWorkShiftResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UpdateWorkShiftResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateWorkShiftResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateWorkShiftResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateWorkShiftResponse;

  static equals(a: UpdateWorkShiftResponse | PlainMessage<UpdateWorkShiftResponse> | undefined, b: UpdateWorkShiftResponse | PlainMessage<UpdateWorkShiftResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteWorkShiftRequest
 */
export declare class DeleteWorkShiftRequest extends Message<DeleteWorkShiftRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DeleteWorkShiftRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteWorkShiftRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteWorkShiftRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteWorkShiftRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteWorkShiftRequest;

  static equals(a: DeleteWorkShiftRequest | PlainMessage<DeleteWorkShiftRequest> | undefined, b: DeleteWorkShiftRequest | PlainMessage<DeleteWorkShiftRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteWorkShiftResponse
 */
export declare class DeleteWorkShiftResponse extends Message<DeleteWorkShiftResponse> {
  constructor(data?: PartialMessage<DeleteWorkShiftResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteWorkShiftResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteWorkShiftResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteWorkShiftResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteWorkShiftResponse;

  static equals(a: DeleteWorkShiftResponse | PlainMessage<DeleteWorkShiftResponse> | undefined, b: DeleteWorkShiftResponse | PlainMessage<DeleteWorkShiftResponse> | undefined): boolean;
}

