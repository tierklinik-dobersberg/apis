// @generated by protoc-gen-es v1.3.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/roster.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, Duration, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { WorkShift } from "./workshift_pb.js";

/**
 * @generated from message tkd.roster.v1.PlannedShift
 */
export declare class PlannedShift extends Message<PlannedShift> {
  /**
   * @generated from field: repeated string assigned_user_ids = 1;
   */
  assignedUserIds: string[];

  /**
   * @generated from field: string work_shift_id = 2;
   */
  workShiftId: string;

  /**
   * @generated from field: bool is_holiday = 3;
   */
  isHoliday: boolean;

  /**
   * @generated from field: bool is_weekend = 4;
   */
  isWeekend: boolean;

  /**
   * @generated from field: google.protobuf.Timestamp from = 5;
   */
  from?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp to = 6;
   */
  to?: Timestamp;

  /**
   * @generated from field: google.protobuf.Duration time_worth = 7;
   */
  timeWorth?: Duration;

  /**
   * @generated from field: tkd.roster.v1.WorkShift shift = 8;
   */
  shift?: WorkShift;

  /**
   * @generated from field: repeated string eligible_user_ids = 9;
   */
  eligibleUserIds: string[];

  /**
   * @generated from field: google.protobuf.FieldMask read_mask = 10;
   */
  readMask?: FieldMask;

  constructor(data?: PartialMessage<PlannedShift>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.PlannedShift";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PlannedShift;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PlannedShift;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PlannedShift;

  static equals(a: PlannedShift | PlainMessage<PlannedShift> | undefined, b: PlannedShift | PlainMessage<PlannedShift> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.RosterMeta
 */
export declare class RosterMeta extends Message<RosterMeta> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: int64 month = 2;
   */
  month: bigint;

  /**
   * @generated from field: int64 year = 3;
   */
  year: bigint;

  /**
   * @generated from field: bool approved = 4;
   */
  approved: boolean;

  /**
   * @generated from field: google.protobuf.Timestamp approved_at = 5;
   */
  approvedAt?: Timestamp;

  /**
   * @generated from field: string last_modified_by = 6;
   */
  lastModifiedBy: string;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 7;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 8;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<RosterMeta>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.RosterMeta";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RosterMeta;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RosterMeta;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RosterMeta;

  static equals(a: RosterMeta | PlainMessage<RosterMeta> | undefined, b: RosterMeta | PlainMessage<RosterMeta> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.WorkTimeStatus
 */
export declare class WorkTimeStatus extends Message<WorkTimeStatus> {
  /**
   * @generated from field: google.protobuf.Duration time_per_week = 1;
   */
  timePerWeek?: Duration;

  /**
   * @generated from field: google.protobuf.Duration expected_work_time = 2;
   */
  expectedWorkTime?: Duration;

  /**
   * @generated from field: google.protobuf.Duration planned_work_time = 3;
   */
  plannedWorkTime?: Duration;

  constructor(data?: PartialMessage<WorkTimeStatus>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.WorkTimeStatus";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkTimeStatus;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkTimeStatus;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkTimeStatus;

  static equals(a: WorkTimeStatus | PlainMessage<WorkTimeStatus> | undefined, b: WorkTimeStatus | PlainMessage<WorkTimeStatus> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.StartSessionRequest
 */
export declare class StartSessionRequest extends Message<StartSessionRequest> {
  /**
   * @generated from field: int32 year = 1;
   */
  year: number;

  /**
   * @generated from field: int32 month = 3;
   */
  month: number;

  constructor(data?: PartialMessage<StartSessionRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.StartSessionRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartSessionRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartSessionRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartSessionRequest;

  static equals(a: StartSessionRequest | PlainMessage<StartSessionRequest> | undefined, b: StartSessionRequest | PlainMessage<StartSessionRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.StartSessionResponse
 */
export declare class StartSessionResponse extends Message<StartSessionResponse> {
  /**
   * @generated from field: tkd.roster.v1.RosterMeta meta = 1;
   */
  meta?: RosterMeta;

  /**
   * @generated from field: repeated tkd.roster.v1.PlannedShift planned_shifts = 2;
   */
  plannedShifts: PlannedShift[];

  /**
   * @generated from field: repeated string active_users = 3;
   */
  activeUsers: string[];

  constructor(data?: PartialMessage<StartSessionResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.StartSessionResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StartSessionResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StartSessionResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StartSessionResponse;

  static equals(a: StartSessionResponse | PlainMessage<StartSessionResponse> | undefined, b: StartSessionResponse | PlainMessage<StartSessionResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.AssignUserToShift
 */
export declare class AssignUserToShift extends Message<AssignUserToShift> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * @generated from field: string shift_id = 2;
   */
  shiftId: string;

  /**
   * @generated from field: google.protobuf.Timestamp from = 3;
   */
  from?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp to = 4;
   */
  to?: Timestamp;

  constructor(data?: PartialMessage<AssignUserToShift>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.AssignUserToShift";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AssignUserToShift;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AssignUserToShift;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AssignUserToShift;

  static equals(a: AssignUserToShift | PlainMessage<AssignUserToShift> | undefined, b: AssignUserToShift | PlainMessage<AssignUserToShift> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UnassignUserFromShift
 */
export declare class UnassignUserFromShift extends Message<UnassignUserFromShift> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * @generated from field: string shift_id = 2;
   */
  shiftId: string;

  /**
   * @generated from field: google.protobuf.Timestamp from = 3;
   */
  from?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp to = 4;
   */
  to?: Timestamp;

  constructor(data?: PartialMessage<UnassignUserFromShift>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UnassignUserFromShift";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UnassignUserFromShift;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UnassignUserFromShift;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UnassignUserFromShift;

  static equals(a: UnassignUserFromShift | PlainMessage<UnassignUserFromShift> | undefined, b: UnassignUserFromShift | PlainMessage<UnassignUserFromShift> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.ShiftUpdateEvent
 */
export declare class ShiftUpdateEvent extends Message<ShiftUpdateEvent> {
  /**
   * @generated from field: tkd.roster.v1.PlannedShift shift = 1;
   */
  shift?: PlannedShift;

  /**
   * @generated from field: string changed_by = 2;
   */
  changedBy: string;

  constructor(data?: PartialMessage<ShiftUpdateEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.ShiftUpdateEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ShiftUpdateEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ShiftUpdateEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ShiftUpdateEvent;

  static equals(a: ShiftUpdateEvent | PlainMessage<ShiftUpdateEvent> | undefined, b: ShiftUpdateEvent | PlainMessage<ShiftUpdateEvent> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.WorkTimeUpdateEvent
 */
export declare class WorkTimeUpdateEvent extends Message<WorkTimeUpdateEvent> {
  /**
   * @generated from field: map<string, tkd.roster.v1.WorkTimeStatus> work_times = 1;
   */
  workTimes: { [key: string]: WorkTimeStatus };

  constructor(data?: PartialMessage<WorkTimeUpdateEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.WorkTimeUpdateEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): WorkTimeUpdateEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): WorkTimeUpdateEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): WorkTimeUpdateEvent;

  static equals(a: WorkTimeUpdateEvent | PlainMessage<WorkTimeUpdateEvent> | undefined, b: WorkTimeUpdateEvent | PlainMessage<WorkTimeUpdateEvent> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.SessionPing
 */
export declare class SessionPing extends Message<SessionPing> {
  constructor(data?: PartialMessage<SessionPing>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.SessionPing";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionPing;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionPing;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionPing;

  static equals(a: SessionPing | PlainMessage<SessionPing> | undefined, b: SessionPing | PlainMessage<SessionPing> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.SessionRequest
 */
export declare class SessionRequest extends Message<SessionRequest> {
  /**
   * @generated from oneof tkd.roster.v1.SessionRequest.kind
   */
  kind: {
    /**
     * @generated from field: tkd.roster.v1.StartSessionRequest start_session = 1;
     */
    value: StartSessionRequest;
    case: "startSession";
  } | {
    /**
     * @generated from field: tkd.roster.v1.AssignUserToShift assign_user = 2;
     */
    value: AssignUserToShift;
    case: "assignUser";
  } | {
    /**
     * @generated from field: tkd.roster.v1.UnassignUserFromShift unassign_user = 3;
     */
    value: UnassignUserFromShift;
    case: "unassignUser";
  } | {
    /**
     * @generated from field: tkd.roster.v1.SessionPing ping = 4;
     */
    value: SessionPing;
    case: "ping";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<SessionRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.SessionRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionRequest;

  static equals(a: SessionRequest | PlainMessage<SessionRequest> | undefined, b: SessionRequest | PlainMessage<SessionRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.SessionResponse
 */
export declare class SessionResponse extends Message<SessionResponse> {
  /**
   * @generated from oneof tkd.roster.v1.SessionResponse.kind
   */
  kind: {
    /**
     * @generated from field: tkd.roster.v1.StartSessionResponse start_session = 1;
     */
    value: StartSessionResponse;
    case: "startSession";
  } | {
    /**
     * @generated from field: tkd.roster.v1.ShiftUpdateEvent shift_update = 2;
     */
    value: ShiftUpdateEvent;
    case: "shiftUpdate";
  } | {
    /**
     * @generated from field: tkd.roster.v1.WorkTimeUpdateEvent work_time_update = 3;
     */
    value: WorkTimeUpdateEvent;
    case: "workTimeUpdate";
  } | {
    /**
     * @generated from field: tkd.roster.v1.UserJoinedSessionEvent user_joined = 4;
     */
    value: UserJoinedSessionEvent;
    case: "userJoined";
  } | {
    /**
     * @generated from field: tkd.roster.v1.UserLeftSessionEvent user_left = 5;
     */
    value: UserLeftSessionEvent;
    case: "userLeft";
  } | {
    /**
     * @generated from field: tkd.roster.v1.UserChatMessageEvent message = 6;
     */
    value: UserChatMessageEvent;
    case: "message";
  } | {
    /**
     * @generated from field: tkd.roster.v1.SessionPing ping = 7;
     */
    value: SessionPing;
    case: "ping";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<SessionResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.SessionResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionResponse;

  static equals(a: SessionResponse | PlainMessage<SessionResponse> | undefined, b: SessionResponse | PlainMessage<SessionResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UserJoinedSessionEvent
 */
export declare class UserJoinedSessionEvent extends Message<UserJoinedSessionEvent> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  constructor(data?: PartialMessage<UserJoinedSessionEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UserJoinedSessionEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserJoinedSessionEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserJoinedSessionEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserJoinedSessionEvent;

  static equals(a: UserJoinedSessionEvent | PlainMessage<UserJoinedSessionEvent> | undefined, b: UserJoinedSessionEvent | PlainMessage<UserJoinedSessionEvent> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UserLeftSessionEvent
 */
export declare class UserLeftSessionEvent extends Message<UserLeftSessionEvent> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  constructor(data?: PartialMessage<UserLeftSessionEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UserLeftSessionEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserLeftSessionEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserLeftSessionEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserLeftSessionEvent;

  static equals(a: UserLeftSessionEvent | PlainMessage<UserLeftSessionEvent> | undefined, b: UserLeftSessionEvent | PlainMessage<UserLeftSessionEvent> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UserChatMessageEvent
 */
export declare class UserChatMessageEvent extends Message<UserChatMessageEvent> {
  /**
   * @generated from field: string sender_user_id = 1;
   */
  senderUserId: string;

  /**
   * @generated from field: string message = 2;
   */
  message: string;

  constructor(data?: PartialMessage<UserChatMessageEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UserChatMessageEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UserChatMessageEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UserChatMessageEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UserChatMessageEvent;

  static equals(a: UserChatMessageEvent | PlainMessage<UserChatMessageEvent> | undefined, b: UserChatMessageEvent | PlainMessage<UserChatMessageEvent> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.ApproveRosterRequest
 */
export declare class ApproveRosterRequest extends Message<ApproveRosterRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<ApproveRosterRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.ApproveRosterRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApproveRosterRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApproveRosterRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApproveRosterRequest;

  static equals(a: ApproveRosterRequest | PlainMessage<ApproveRosterRequest> | undefined, b: ApproveRosterRequest | PlainMessage<ApproveRosterRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.ApproveRosterResponse
 */
export declare class ApproveRosterResponse extends Message<ApproveRosterResponse> {
  constructor(data?: PartialMessage<ApproveRosterResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.ApproveRosterResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ApproveRosterResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ApproveRosterResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ApproveRosterResponse;

  static equals(a: ApproveRosterResponse | PlainMessage<ApproveRosterResponse> | undefined, b: ApproveRosterResponse | PlainMessage<ApproveRosterResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteRosterRequest
 */
export declare class DeleteRosterRequest extends Message<DeleteRosterRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DeleteRosterRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteRosterRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRosterRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRosterRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRosterRequest;

  static equals(a: DeleteRosterRequest | PlainMessage<DeleteRosterRequest> | undefined, b: DeleteRosterRequest | PlainMessage<DeleteRosterRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteRosterResponse
 */
export declare class DeleteRosterResponse extends Message<DeleteRosterResponse> {
  constructor(data?: PartialMessage<DeleteRosterResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteRosterResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteRosterResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteRosterResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteRosterResponse;

  static equals(a: DeleteRosterResponse | PlainMessage<DeleteRosterResponse> | undefined, b: DeleteRosterResponse | PlainMessage<DeleteRosterResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetByDate
 */
export declare class GetByDate extends Message<GetByDate> {
  /**
   * @generated from field: int32 year = 1;
   */
  year: number;

  /**
   * @generated from field: int32 month = 3;
   */
  month: number;

  constructor(data?: PartialMessage<GetByDate>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetByDate";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetByDate;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetByDate;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetByDate;

  static equals(a: GetByDate | PlainMessage<GetByDate> | undefined, b: GetByDate | PlainMessage<GetByDate> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetRosterRequest
 */
export declare class GetRosterRequest extends Message<GetRosterRequest> {
  /**
   * @generated from oneof tkd.roster.v1.GetRosterRequest.search
   */
  search: {
    /**
     * @generated from field: string id = 1;
     */
    value: string;
    case: "id";
  } | {
    /**
     * @generated from field: tkd.roster.v1.GetByDate date = 2;
     */
    value: GetByDate;
    case: "date";
  } | { case: undefined; value?: undefined };

  /**
   * @generated from field: google.protobuf.FieldMask read_mask = 3;
   */
  readMask?: FieldMask;

  constructor(data?: PartialMessage<GetRosterRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetRosterRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRosterRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRosterRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRosterRequest;

  static equals(a: GetRosterRequest | PlainMessage<GetRosterRequest> | undefined, b: GetRosterRequest | PlainMessage<GetRosterRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetRosterResponse
 */
export declare class GetRosterResponse extends Message<GetRosterResponse> {
  /**
   * @generated from field: tkd.roster.v1.RosterMeta meta = 1;
   */
  meta?: RosterMeta;

  /**
   * @generated from field: repeated tkd.roster.v1.PlannedShift shifts = 2;
   */
  shifts: PlannedShift[];

  constructor(data?: PartialMessage<GetRosterResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetRosterResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetRosterResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetRosterResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetRosterResponse;

  static equals(a: GetRosterResponse | PlainMessage<GetRosterResponse> | undefined, b: GetRosterResponse | PlainMessage<GetRosterResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetWorkingStaffRequest
 */
export declare class GetWorkingStaffRequest extends Message<GetWorkingStaffRequest> {
  /**
   * @generated from field: google.protobuf.Timestamp time = 1;
   */
  time?: Timestamp;

  /**
   * @generated from field: google.protobuf.FieldMask read_maks = 2;
   */
  readMaks?: FieldMask;

  constructor(data?: PartialMessage<GetWorkingStaffRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetWorkingStaffRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetWorkingStaffRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetWorkingStaffRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetWorkingStaffRequest;

  static equals(a: GetWorkingStaffRequest | PlainMessage<GetWorkingStaffRequest> | undefined, b: GetWorkingStaffRequest | PlainMessage<GetWorkingStaffRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.GetWorkingStaffResponse
 */
export declare class GetWorkingStaffResponse extends Message<GetWorkingStaffResponse> {
  /**
   * @generated from field: repeated string user_ids = 1;
   */
  userIds: string[];

  /**
   * @generated from field: repeated tkd.roster.v1.PlannedShift current_shifts = 2;
   */
  currentShifts: PlannedShift[];

  constructor(data?: PartialMessage<GetWorkingStaffResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.GetWorkingStaffResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetWorkingStaffResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetWorkingStaffResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetWorkingStaffResponse;

  static equals(a: GetWorkingStaffResponse | PlainMessage<GetWorkingStaffResponse> | undefined, b: GetWorkingStaffResponse | PlainMessage<GetWorkingStaffResponse> | undefined): boolean;
}

