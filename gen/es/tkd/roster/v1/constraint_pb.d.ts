// @generated by protoc-gen-es v1.4.1 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/constraint.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { OffTimeEntry } from "./offtime_pb.js";

/**
 * OffTimeViolation is a constraint type due to a off-time request.
 *
 * @generated from message tkd.roster.v1.OffTimeViolation
 */
export declare class OffTimeViolation extends Message<OffTimeViolation> {
  /**
   * Entry holds the OffTimeEntry.
   *
   * @generated from field: tkd.roster.v1.OffTimeEntry entry = 1;
   */
  entry?: OffTimeEntry;

  constructor(data?: PartialMessage<OffTimeViolation>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.OffTimeViolation";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OffTimeViolation;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OffTimeViolation;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OffTimeViolation;

  static equals(a: OffTimeViolation | PlainMessage<OffTimeViolation> | undefined, b: OffTimeViolation | PlainMessage<OffTimeViolation> | undefined): boolean;
}

/**
 * ConstraintEvaluationViolation is a constraint type that is the result
 * of evaluating a constraint expression for a user or roster.
 *
 * @generated from message tkd.roster.v1.ConstraintEvaluationViolation
 */
export declare class ConstraintEvaluationViolation extends Message<ConstraintEvaluationViolation> {
  /**
   * Id is the unique identifier of the Constraint that caused the violation.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * Description is the description of the Constraint and copied to the
   * ConstraintEvaluationViolation for convenience.
   *
   * @generated from field: string description = 2;
   */
  description: string;

  constructor(data?: PartialMessage<ConstraintEvaluationViolation>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.ConstraintEvaluationViolation";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConstraintEvaluationViolation;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConstraintEvaluationViolation;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConstraintEvaluationViolation;

  static equals(a: ConstraintEvaluationViolation | PlainMessage<ConstraintEvaluationViolation> | undefined, b: ConstraintEvaluationViolation | PlainMessage<ConstraintEvaluationViolation> | undefined): boolean;
}

/**
 * ConstraintViolation is a violation of a constraint for a given
 * user or roster.
 *
 * @generated from message tkd.roster.v1.ConstraintViolation
 */
export declare class ConstraintViolation extends Message<ConstraintViolation> {
  /**
   * @generated from field: bool hard = 1;
   */
  hard: boolean;

  /**
   * @generated from oneof tkd.roster.v1.ConstraintViolation.kind
   */
  kind: {
    /**
     * @generated from field: tkd.roster.v1.OffTimeViolation off_time = 2;
     */
    value: OffTimeViolation;
    case: "offTime";
  } | {
    /**
     * @generated from field: tkd.roster.v1.ConstraintEvaluationViolation evaluation = 3;
     */
    value: ConstraintEvaluationViolation;
    case: "evaluation";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<ConstraintViolation>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.ConstraintViolation";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConstraintViolation;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConstraintViolation;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConstraintViolation;

  static equals(a: ConstraintViolation | PlainMessage<ConstraintViolation> | undefined, b: ConstraintViolation | PlainMessage<ConstraintViolation> | undefined): boolean;
}

/**
 * ConstraintViolationList holds all ConstraintViolations for a
 * user.
 *
 * @generated from message tkd.roster.v1.ConstraintViolationList
 */
export declare class ConstraintViolationList extends Message<ConstraintViolationList> {
  /**
   * @generated from field: string user_id = 1;
   */
  userId: string;

  /**
   * @generated from field: repeated tkd.roster.v1.ConstraintViolation violations = 2;
   */
  violations: ConstraintViolation[];

  constructor(data?: PartialMessage<ConstraintViolationList>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.ConstraintViolationList";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ConstraintViolationList;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ConstraintViolationList;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ConstraintViolationList;

  static equals(a: ConstraintViolationList | PlainMessage<ConstraintViolationList> | undefined, b: ConstraintViolationList | PlainMessage<ConstraintViolationList> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.Constraint
 */
export declare class Constraint extends Message<Constraint> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string description = 2;
   */
  description: string;

  /**
   * @generated from field: string expression = 3;
   */
  expression: string;

  /**
   * @generated from field: repeated string role_ids = 4;
   */
  roleIds: string[];

  /**
   * @generated from field: repeated string user_ids = 5;
   */
  userIds: string[];

  /**
   * @generated from field: bool hard = 6;
   */
  hard: boolean;

  /**
   * @generated from field: float penalty = 7;
   */
  penalty: number;

  /**
   * @generated from field: bool deny = 8;
   */
  deny: boolean;

  /**
   * @generated from field: bool roster_only = 9;
   */
  rosterOnly: boolean;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 10;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: string creator_id = 11;
   */
  creatorId: string;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 12;
   */
  updatedAt?: Timestamp;

  /**
   * @generated from field: string last_updated_by = 13;
   */
  lastUpdatedBy: string;

  constructor(data?: PartialMessage<Constraint>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.Constraint";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Constraint;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Constraint;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Constraint;

  static equals(a: Constraint | PlainMessage<Constraint> | undefined, b: Constraint | PlainMessage<Constraint> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.CreateConstraintRequest
 */
export declare class CreateConstraintRequest extends Message<CreateConstraintRequest> {
  /**
   * @generated from field: string description = 1;
   */
  description: string;

  /**
   * @generated from field: string expression = 2;
   */
  expression: string;

  /**
   * @generated from field: repeated string role_ids = 3;
   */
  roleIds: string[];

  /**
   * @generated from field: repeated string user_ids = 4;
   */
  userIds: string[];

  /**
   * @generated from field: bool hard = 5;
   */
  hard: boolean;

  /**
   * @generated from field: float penalty = 6;
   */
  penalty: number;

  /**
   * @generated from field: bool deny = 7;
   */
  deny: boolean;

  /**
   * @generated from field: bool roster_only = 8;
   */
  rosterOnly: boolean;

  constructor(data?: PartialMessage<CreateConstraintRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.CreateConstraintRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateConstraintRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateConstraintRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateConstraintRequest;

  static equals(a: CreateConstraintRequest | PlainMessage<CreateConstraintRequest> | undefined, b: CreateConstraintRequest | PlainMessage<CreateConstraintRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.CreateConstraintResponse
 */
export declare class CreateConstraintResponse extends Message<CreateConstraintResponse> {
  /**
   * @generated from field: tkd.roster.v1.Constraint constraint = 1;
   */
  constraint?: Constraint;

  constructor(data?: PartialMessage<CreateConstraintResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.CreateConstraintResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateConstraintResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateConstraintResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateConstraintResponse;

  static equals(a: CreateConstraintResponse | PlainMessage<CreateConstraintResponse> | undefined, b: CreateConstraintResponse | PlainMessage<CreateConstraintResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UpdateConstraintRequest
 */
export declare class UpdateConstraintRequest extends Message<UpdateConstraintRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string description = 2;
   */
  description: string;

  /**
   * @generated from field: string expression = 3;
   */
  expression: string;

  /**
   * @generated from field: repeated string role_ids = 4;
   */
  roleIds: string[];

  /**
   * @generated from field: repeated string user_ids = 5;
   */
  userIds: string[];

  /**
   * @generated from field: bool hard = 6;
   */
  hard: boolean;

  /**
   * @generated from field: float penalty = 7;
   */
  penalty: number;

  /**
   * @generated from field: bool deny = 8;
   */
  deny: boolean;

  /**
   * @generated from field: bool roster_only = 9;
   */
  rosterOnly: boolean;

  /**
   * @generated from field: google.protobuf.FieldMask write_mask = 20;
   */
  writeMask?: FieldMask;

  constructor(data?: PartialMessage<UpdateConstraintRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UpdateConstraintRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateConstraintRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateConstraintRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateConstraintRequest;

  static equals(a: UpdateConstraintRequest | PlainMessage<UpdateConstraintRequest> | undefined, b: UpdateConstraintRequest | PlainMessage<UpdateConstraintRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.UpdateConstraintResponse
 */
export declare class UpdateConstraintResponse extends Message<UpdateConstraintResponse> {
  /**
   * @generated from field: tkd.roster.v1.Constraint constraint = 1;
   */
  constraint?: Constraint;

  constructor(data?: PartialMessage<UpdateConstraintResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.UpdateConstraintResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateConstraintResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateConstraintResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateConstraintResponse;

  static equals(a: UpdateConstraintResponse | PlainMessage<UpdateConstraintResponse> | undefined, b: UpdateConstraintResponse | PlainMessage<UpdateConstraintResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteConstraintRequest
 */
export declare class DeleteConstraintRequest extends Message<DeleteConstraintRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DeleteConstraintRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteConstraintRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteConstraintRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteConstraintRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteConstraintRequest;

  static equals(a: DeleteConstraintRequest | PlainMessage<DeleteConstraintRequest> | undefined, b: DeleteConstraintRequest | PlainMessage<DeleteConstraintRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.DeleteConstraintResponse
 */
export declare class DeleteConstraintResponse extends Message<DeleteConstraintResponse> {
  constructor(data?: PartialMessage<DeleteConstraintResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.DeleteConstraintResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteConstraintResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteConstraintResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteConstraintResponse;

  static equals(a: DeleteConstraintResponse | PlainMessage<DeleteConstraintResponse> | undefined, b: DeleteConstraintResponse | PlainMessage<DeleteConstraintResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.FindConstraintsRequest
 */
export declare class FindConstraintsRequest extends Message<FindConstraintsRequest> {
  /**
   * @generated from field: repeated string user_ids = 1;
   */
  userIds: string[];

  /**
   * @generated from field: repeated string role_ids = 2;
   */
  roleIds: string[];

  constructor(data?: PartialMessage<FindConstraintsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.FindConstraintsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FindConstraintsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FindConstraintsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FindConstraintsRequest;

  static equals(a: FindConstraintsRequest | PlainMessage<FindConstraintsRequest> | undefined, b: FindConstraintsRequest | PlainMessage<FindConstraintsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.roster.v1.FindConstraintsResponse
 */
export declare class FindConstraintsResponse extends Message<FindConstraintsResponse> {
  /**
   * @generated from field: repeated tkd.roster.v1.Constraint results = 1;
   */
  results: Constraint[];

  constructor(data?: PartialMessage<FindConstraintsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.roster.v1.FindConstraintsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FindConstraintsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FindConstraintsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FindConstraintsResponse;

  static equals(a: FindConstraintsResponse | PlainMessage<FindConstraintsResponse> | undefined, b: FindConstraintsResponse | PlainMessage<FindConstraintsResponse> | undefined): boolean;
}

