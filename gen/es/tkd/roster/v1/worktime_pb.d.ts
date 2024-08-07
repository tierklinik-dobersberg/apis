// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/worktime.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, Duration, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { OffTimeCosts } from "./offtime_pb.js";

/**
 * WorkTime describes the regular work time for a given user.
 *
 * @generated from message tkd.roster.v1.WorkTime
 */
export declare class WorkTime extends Message<WorkTime> {
  /**
   * Id is a unique identifier for this work-time entry.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * UserId holds the ID of the user this work-time entry belongs to.
   *
   * @generated from field: string user_id = 2;
   */
  userId: string;

  /**
   * TimePerWeek defines the regular working time per week.
   *
   * @generated from field: google.protobuf.Duration time_per_week = 3;
   */
  timePerWeek?: Duration;

  /**
   * ApplicableAfter defines the date at which this work-time entry
   * is considered active.
   *
   * Format: YYYY-MM-DD
   *
   * @generated from field: string applicable_after = 4;
   */
  applicableAfter: string;

  /**
   * VacationWeeksPerYear defines how many weeks of vacation should
   * be granted to UserId in a full-year of work time.
   *
   * @generated from field: float vacation_weeks_per_year = 5;
   */
  vacationWeeksPerYear: number;

  /**
   * Overtime Allowance per month.
   *
   * @generated from field: google.protobuf.Duration overtime_allowance_per_month = 6;
   */
  overtimeAllowancePerMonth?: Duration;

  /**
   * Whether or not the user should be excluded from time-tracking.
   *
   * @generated from field: bool exclude_from_time_tracking = 7;
   */
  excludeFromTimeTracking: boolean;

  /**
   * An optional timestamp _AFTER_ which this work-time entry ends.
   *
   * Format: YYYY-MM-DD
   *
   * @generated from field: string ends_with = 8;
   */
  endsWith: string;

  constructor(data?: PartialMessage<WorkTime>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.WorkTime";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkTime;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkTime;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkTime;

  static equals(a: WorkTime | PlainMessage<WorkTime> | undefined, b: WorkTime | PlainMessage<WorkTime> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.SetWorkTimeRequest
 */
export declare class SetWorkTimeRequest extends Message<SetWorkTimeRequest> {
  /**
   * @generated from field: repeated tkd.roster.v1.WorkTime work_times = 1;
   */
  workTimes: WorkTime[];

  constructor(data?: PartialMessage<SetWorkTimeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.SetWorkTimeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetWorkTimeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetWorkTimeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetWorkTimeRequest;

  static equals(a: SetWorkTimeRequest | PlainMessage<SetWorkTimeRequest> | undefined, b: SetWorkTimeRequest | PlainMessage<SetWorkTimeRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.SetWorkTimeResponse
 */
export declare class SetWorkTimeResponse extends Message<SetWorkTimeResponse> {
  /**
   * @generated from field: repeated tkd.roster.v1.WorkTime work_times = 1;
   */
  workTimes: WorkTime[];

  constructor(data?: PartialMessage<SetWorkTimeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.SetWorkTimeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetWorkTimeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetWorkTimeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetWorkTimeResponse;

  static equals(a: SetWorkTimeResponse | PlainMessage<SetWorkTimeResponse> | undefined, b: SetWorkTimeResponse | PlainMessage<SetWorkTimeResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetWorkTimeRequest
 */
export declare class GetWorkTimeRequest extends Message<GetWorkTimeRequest> {
  /**
   * UserIds is a list of user IDs for which the work-time should
   * be returned.
   *
   * @generated from field: repeated string user_ids = 1;
   */
  userIds: string[];

  /**
   * ReadMask defines which fields of the response should be populated.
   * The field-mask is applied to the `result` field of the GetWorkTimeResponse.
   * For example, to only retrieve the user_id and the current work time specify
   * a read_mask like this:
   *
   * {
   *    read_mask: {
   *      paths: [
   *          "user_id",
   *          "current",
   *      ]
   *    }    
   * }
   *
   *
   * @generated from field: google.protobuf.FieldMask read_mask = 2;
   */
  readMask?: FieldMask;

  constructor(data?: PartialMessage<GetWorkTimeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetWorkTimeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetWorkTimeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetWorkTimeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetWorkTimeRequest;

  static equals(a: GetWorkTimeRequest | PlainMessage<GetWorkTimeRequest> | undefined, b: GetWorkTimeRequest | PlainMessage<GetWorkTimeRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UserWorkTime
 */
export declare class UserWorkTime extends Message<UserWorkTime> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * @generated from field: tkd.roster.v1.WorkTime current = 2;
   */
  current?: WorkTime;

  /**
   * @generated from field: repeated tkd.roster.v1.WorkTime history = 3;
   */
  history: WorkTime[];

  constructor(data?: PartialMessage<UserWorkTime>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UserWorkTime";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserWorkTime;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserWorkTime;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserWorkTime;

  static equals(a: UserWorkTime | PlainMessage<UserWorkTime> | undefined, b: UserWorkTime | PlainMessage<UserWorkTime> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetWorkTimeResponse
 */
export declare class GetWorkTimeResponse extends Message<GetWorkTimeResponse> {
  /**
   * @generated from field: repeated tkd.roster.v1.UserWorkTime results = 1;
   */
  results: UserWorkTime[];

  constructor(data?: PartialMessage<GetWorkTimeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetWorkTimeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetWorkTimeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetWorkTimeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetWorkTimeResponse;

  static equals(a: GetWorkTimeResponse | PlainMessage<GetWorkTimeResponse> | undefined, b: GetWorkTimeResponse | PlainMessage<GetWorkTimeResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.SumForUsers
 */
export declare class SumForUsers extends Message<SumForUsers> {
  /**
   * @generated from field: repeated string user_ids = 1;
   */
  userIds: string[];

  constructor(data?: PartialMessage<SumForUsers>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.SumForUsers";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SumForUsers;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SumForUsers;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SumForUsers;

  static equals(a: SumForUsers | PlainMessage<SumForUsers> | undefined, b: SumForUsers | PlainMessage<SumForUsers> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetVacationCreditsLeftRequest
 */
export declare class GetVacationCreditsLeftRequest extends Message<GetVacationCreditsLeftRequest> {
  /**
   * @generated from field: tkd.roster.v1.SumForUsers for_users = 1;
   */
  forUsers?: SumForUsers;

  /**
   * @generated from field: google.protobuf.Timestamp until = 2;
   */
  until?: Timestamp;

  /**
   * @generated from field: bool analyze = 3;
   */
  analyze: boolean;

  constructor(data?: PartialMessage<GetVacationCreditsLeftRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetVacationCreditsLeftRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetVacationCreditsLeftRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetVacationCreditsLeftRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetVacationCreditsLeftRequest;

  static equals(a: GetVacationCreditsLeftRequest | PlainMessage<GetVacationCreditsLeftRequest> | undefined, b: GetVacationCreditsLeftRequest | PlainMessage<GetVacationCreditsLeftRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.AnalyzeVacationSum
 */
export declare class AnalyzeVacationSum extends Message<AnalyzeVacationSum> {
  /**
   * @generated from field: tkd.roster.v1.WorkTime work_time = 1;
   */
  workTime?: WorkTime;

  /**
   * @generated from field: google.protobuf.Timestamp ends_at = 2;
   */
  endsAt?: Timestamp;

  /**
   * @generated from field: double number_of_days = 3;
   */
  numberOfDays: number;

  /**
   * @generated from field: double vacation_weeks_per_day = 4;
   */
  vacationWeeksPerDay: number;

  /**
   * @generated from field: google.protobuf.Duration vacation_per_work_time = 5;
   */
  vacationPerWorkTime?: Duration;

  /**
   * @generated from field: repeated tkd.roster.v1.OffTimeCosts costs = 6;
   */
  costs: OffTimeCosts[];

  /**
   * @generated from field: google.protobuf.Duration costs_sum = 7;
   */
  costsSum?: Duration;

  constructor(data?: PartialMessage<AnalyzeVacationSum>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.AnalyzeVacationSum";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyzeVacationSum;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyzeVacationSum;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyzeVacationSum;

  static equals(a: AnalyzeVacationSum | PlainMessage<AnalyzeVacationSum> | undefined, b: AnalyzeVacationSum | PlainMessage<AnalyzeVacationSum> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.AnalyzeVacation
 */
export declare class AnalyzeVacation extends Message<AnalyzeVacation> {
  /**
   * @generated from field: repeated tkd.roster.v1.AnalyzeVacationSum slices = 1;
   */
  slices: AnalyzeVacationSum[];

  constructor(data?: PartialMessage<AnalyzeVacation>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.AnalyzeVacation";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AnalyzeVacation;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AnalyzeVacation;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AnalyzeVacation;

  static equals(a: AnalyzeVacation | PlainMessage<AnalyzeVacation> | undefined, b: AnalyzeVacation | PlainMessage<AnalyzeVacation> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UserVacationSum
 */
export declare class UserVacationSum extends Message<UserVacationSum> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * @generated from field: google.protobuf.Duration vacation_credits_left = 2;
   */
  vacationCreditsLeft?: Duration;

  /**
   * @generated from field: google.protobuf.Duration time_off_credits = 3;
   */
  timeOffCredits?: Duration;

  /**
   * @generated from field: tkd.roster.v1.AnalyzeVacation analysis = 4;
   */
  analysis?: AnalyzeVacation;

  constructor(data?: PartialMessage<UserVacationSum>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UserVacationSum";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserVacationSum;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserVacationSum;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserVacationSum;

  static equals(a: UserVacationSum | PlainMessage<UserVacationSum> | undefined, b: UserVacationSum | PlainMessage<UserVacationSum> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetVacationCreditsLeftResponse
 */
export declare class GetVacationCreditsLeftResponse extends Message<GetVacationCreditsLeftResponse> {
  /**
   * @generated from field: repeated tkd.roster.v1.UserVacationSum results = 1;
   */
  results: UserVacationSum[];

  constructor(data?: PartialMessage<GetVacationCreditsLeftResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetVacationCreditsLeftResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetVacationCreditsLeftResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetVacationCreditsLeftResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetVacationCreditsLeftResponse;

  static equals(a: GetVacationCreditsLeftResponse | PlainMessage<GetVacationCreditsLeftResponse> | undefined, b: GetVacationCreditsLeftResponse | PlainMessage<GetVacationCreditsLeftResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteWorkTimeRequest
 */
export declare class DeleteWorkTimeRequest extends Message<DeleteWorkTimeRequest> {
  /**
   * @generated from field: repeated string ids = 1;
   */
  ids: string[];

  constructor(data?: PartialMessage<DeleteWorkTimeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteWorkTimeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteWorkTimeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteWorkTimeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteWorkTimeRequest;

  static equals(a: DeleteWorkTimeRequest | PlainMessage<DeleteWorkTimeRequest> | undefined, b: DeleteWorkTimeRequest | PlainMessage<DeleteWorkTimeRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteWorkTimeResponse
 */
export declare class DeleteWorkTimeResponse extends Message<DeleteWorkTimeResponse> {
  constructor(data?: PartialMessage<DeleteWorkTimeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteWorkTimeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteWorkTimeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteWorkTimeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteWorkTimeResponse;

  static equals(a: DeleteWorkTimeResponse | PlainMessage<DeleteWorkTimeResponse> | undefined, b: DeleteWorkTimeResponse | PlainMessage<DeleteWorkTimeResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UpdateWorkTimeRequest
 */
export declare class UpdateWorkTimeRequest extends Message<UpdateWorkTimeRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: bool exclude_from_time_tracking = 2;
   */
  excludeFromTimeTracking: boolean;

  /**
   * Format: YYYY-MM-DD
   *
   * @generated from field: string ends_with = 3;
   */
  endsWith: string;

  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 4;
   */
  fieldMask?: FieldMask;

  constructor(data?: PartialMessage<UpdateWorkTimeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UpdateWorkTimeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateWorkTimeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateWorkTimeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateWorkTimeRequest;

  static equals(a: UpdateWorkTimeRequest | PlainMessage<UpdateWorkTimeRequest> | undefined, b: UpdateWorkTimeRequest | PlainMessage<UpdateWorkTimeRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UpdateWorkTimeResponse
 */
export declare class UpdateWorkTimeResponse extends Message<UpdateWorkTimeResponse> {
  /**
   * @generated from field: tkd.roster.v1.WorkTime worktime = 1;
   */
  worktime?: WorkTime;

  constructor(data?: PartialMessage<UpdateWorkTimeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UpdateWorkTimeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateWorkTimeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateWorkTimeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateWorkTimeResponse;

  static equals(a: UpdateWorkTimeResponse | PlainMessage<UpdateWorkTimeResponse> | undefined, b: UpdateWorkTimeResponse | PlainMessage<UpdateWorkTimeResponse> | undefined): boolean;
}

