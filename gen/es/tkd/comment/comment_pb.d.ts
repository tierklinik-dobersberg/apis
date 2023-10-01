// @generated by protoc-gen-es v1.3.1 with parameter "target=js+dts"
// @generated from file tkd/comment/comment.proto (package tkd.comment.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum tkd.comment.v1.NotificationType
 */
export declare enum NotificationType {
  /**
   * @generated from enum value: NOTIFICATION_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: NOTIFICATION_TYPE_MAIL = 1;
   */
  MAIL = 1,

  /**
   * @generated from enum value: NOTIFICATION_TYPE_SMS = 2;
   */
  SMS = 2,
}

/**
 * @generated from message tkd.comment.v1.Scope
 */
export declare class Scope extends Message<Scope> {
  /**
   * The unique identifier of the scope. This must not be changed
   * once created and referenced by comments.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The display name of the comment scope. This is used
   * for notifications.
   *
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * If set, the type of notification that will be sent for answers and
   * mentions.
   *
   * @generated from field: tkd.comment.v1.NotificationType notifcation_type = 3;
   */
  notifcationType: NotificationType;

  constructor(data?: PartialMessage<Scope>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.Scope";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Scope;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Scope;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Scope;

  static equals(a: Scope | PlainMessage<Scope> | undefined, b: Scope | PlainMessage<Scope> | undefined): boolean;
}

/**
 * @generated from message tkd.comment.v1.Comment
 */
export declare class Comment extends Message<Comment> {
  /**
   * The scope of the comment.
   *
   * @generated from field: string scope = 1;
   */
  scope: string;

  /**
   * The scope-unique ID of the comment.
   *
   * @generated from field: string id = 2;
   */
  id: string;

  /**
   * The content of the comment in markdown format.
   *
   * @generated from field: string content = 3;
   */
  content: string;

  /**
   * An optional Id of the parent comment. This is only set if the
   * comment is an answer to another comment.
   *
   * @generated from field: string parent_id = 4;
   */
  parentId: string;

  /**
   * The timestamp at which the comment has been created.
   *
   * @generated from field: google.protobuf.Timestamp created_at = 5;
   */
  createdAt?: Timestamp;

  /**
   * The user identifier that created the comment.
   *
   * @generated from field: string creator_id = 6;
   */
  creatorId: string;

  constructor(data?: PartialMessage<Comment>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.Comment";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Comment;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Comment;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Comment;

  static equals(a: Comment | PlainMessage<Comment> | undefined, b: Comment | PlainMessage<Comment> | undefined): boolean;
}

/**
 * @generated from message tkd.comment.v1.CommentTree
 */
export declare class CommentTree extends Message<CommentTree> {
  /**
   * The actual comment.
   *
   * @generated from field: tkd.comment.v1.Comment comment = 1;
   */
  comment?: Comment;

  /**
   * A list of answers to this comment.
   *
   * @generated from field: repeated tkd.comment.v1.CommentTree answers = 2;
   */
  answers: CommentTree[];

  constructor(data?: PartialMessage<CommentTree>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.CommentTree";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CommentTree;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CommentTree;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CommentTree;

  static equals(a: CommentTree | PlainMessage<CommentTree> | undefined, b: CommentTree | PlainMessage<CommentTree> | undefined): boolean;
}

/**
 * CreateScopeRequest is used to create a new comment scope.
 *
 * @generated from message tkd.comment.v1.CreateScopeRequest
 */
export declare class CreateScopeRequest extends Message<CreateScopeRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: tkd.comment.v1.NotificationType notifcation_type = 3;
   */
  notifcationType: NotificationType;

  constructor(data?: PartialMessage<CreateScopeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.CreateScopeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateScopeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateScopeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateScopeRequest;

  static equals(a: CreateScopeRequest | PlainMessage<CreateScopeRequest> | undefined, b: CreateScopeRequest | PlainMessage<CreateScopeRequest> | undefined): boolean;
}

/**
 * CreateScopeResponse is returned when a new comment scope has been created
 * and includes the created scope definition.
 *
 * @generated from message tkd.comment.v1.CreateScopeResponse
 */
export declare class CreateScopeResponse extends Message<CreateScopeResponse> {
  /**
   * @generated from field: tkd.comment.v1.Scope scope = 1;
   */
  scope?: Scope;

  constructor(data?: PartialMessage<CreateScopeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.CreateScopeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateScopeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateScopeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateScopeResponse;

  static equals(a: CreateScopeResponse | PlainMessage<CreateScopeResponse> | undefined, b: CreateScopeResponse | PlainMessage<CreateScopeResponse> | undefined): boolean;
}

/**
 *
 *
 * @generated from message tkd.comment.v1.ListScopeRequest
 */
export declare class ListScopeRequest extends Message<ListScopeRequest> {
  constructor(data?: PartialMessage<ListScopeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.ListScopeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListScopeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListScopeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListScopeRequest;

  static equals(a: ListScopeRequest | PlainMessage<ListScopeRequest> | undefined, b: ListScopeRequest | PlainMessage<ListScopeRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.comment.v1.ListScopeResponse
 */
export declare class ListScopeResponse extends Message<ListScopeResponse> {
  /**
   * @generated from field: repeated tkd.comment.v1.Scope scopes = 1;
   */
  scopes: Scope[];

  constructor(data?: PartialMessage<ListScopeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.ListScopeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListScopeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListScopeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListScopeResponse;

  static equals(a: ListScopeResponse | PlainMessage<ListScopeResponse> | undefined, b: ListScopeResponse | PlainMessage<ListScopeResponse> | undefined): boolean;
}

/**
 * Delete Scope
 *
 * @generated from message tkd.comment.v1.DeleteScopeRequest
 */
export declare class DeleteScopeRequest extends Message<DeleteScopeRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DeleteScopeRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.DeleteScopeRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteScopeRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteScopeRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteScopeRequest;

  static equals(a: DeleteScopeRequest | PlainMessage<DeleteScopeRequest> | undefined, b: DeleteScopeRequest | PlainMessage<DeleteScopeRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.comment.v1.DeleteScopeResponse
 */
export declare class DeleteScopeResponse extends Message<DeleteScopeResponse> {
  constructor(data?: PartialMessage<DeleteScopeResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.DeleteScopeResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteScopeResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteScopeResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteScopeResponse;

  static equals(a: DeleteScopeResponse | PlainMessage<DeleteScopeResponse> | undefined, b: DeleteScopeResponse | PlainMessage<DeleteScopeResponse> | undefined): boolean;
}

/**
 * Create Comment
 *
 * @generated from message tkd.comment.v1.CreateCommentRequest
 */
export declare class CreateCommentRequest extends Message<CreateCommentRequest> {
  /**
   * @generated from field: string scope = 1;
   */
  scope: string;

  /**
   * @generated from field: string parent_id = 2;
   */
  parentId: string;

  /**
   * @generated from field: string content = 3;
   */
  content: string;

  /**
   * @generated from field: google.protobuf.FieldMask read_mask = 4;
   */
  readMask?: FieldMask;

  constructor(data?: PartialMessage<CreateCommentRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.CreateCommentRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateCommentRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateCommentRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateCommentRequest;

  static equals(a: CreateCommentRequest | PlainMessage<CreateCommentRequest> | undefined, b: CreateCommentRequest | PlainMessage<CreateCommentRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.comment.v1.CreateCommentResponse
 */
export declare class CreateCommentResponse extends Message<CreateCommentResponse> {
  /**
   * @generated from field: tkd.comment.v1.Comment comment = 1;
   */
  comment?: Comment;

  constructor(data?: PartialMessage<CreateCommentResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.CreateCommentResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateCommentResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateCommentResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateCommentResponse;

  static equals(a: CreateCommentResponse | PlainMessage<CreateCommentResponse> | undefined, b: CreateCommentResponse | PlainMessage<CreateCommentResponse> | undefined): boolean;
}

/**
 * List Comments
 *
 * @generated from message tkd.comment.v1.ListCommentsRequest
 */
export declare class ListCommentsRequest extends Message<ListCommentsRequest> {
  /**
   * @generated from field: string scope = 1;
   */
  scope: string;

  /**
   * @generated from field: bool recurse = 2;
   */
  recurse: boolean;

  constructor(data?: PartialMessage<ListCommentsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.ListCommentsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListCommentsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListCommentsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListCommentsRequest;

  static equals(a: ListCommentsRequest | PlainMessage<ListCommentsRequest> | undefined, b: ListCommentsRequest | PlainMessage<ListCommentsRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.comment.v1.ListCommentsResponse
 */
export declare class ListCommentsResponse extends Message<ListCommentsResponse> {
  /**
   * @generated from field: repeated tkd.comment.v1.CommentTree result = 2;
   */
  result: CommentTree[];

  constructor(data?: PartialMessage<ListCommentsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.ListCommentsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListCommentsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListCommentsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListCommentsResponse;

  static equals(a: ListCommentsResponse | PlainMessage<ListCommentsResponse> | undefined, b: ListCommentsResponse | PlainMessage<ListCommentsResponse> | undefined): boolean;
}

/**
 * Get Comments
 *
 * @generated from message tkd.comment.v1.GetCommentRequest
 */
export declare class GetCommentRequest extends Message<GetCommentRequest> {
  /**
   * @generated from field: string scope = 1;
   */
  scope: string;

  /**
   * @generated from field: string id = 2;
   */
  id: string;

  /**
   * @generated from field: bool recurse = 3;
   */
  recurse: boolean;

  constructor(data?: PartialMessage<GetCommentRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.GetCommentRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCommentRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCommentRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCommentRequest;

  static equals(a: GetCommentRequest | PlainMessage<GetCommentRequest> | undefined, b: GetCommentRequest | PlainMessage<GetCommentRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.comment.v1.GetCommentResponse
 */
export declare class GetCommentResponse extends Message<GetCommentResponse> {
  /**
   * @generated from field: tkd.comment.v1.CommentTree result = 2;
   */
  result?: CommentTree;

  constructor(data?: PartialMessage<GetCommentResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.comment.v1.GetCommentResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetCommentResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetCommentResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetCommentResponse;

  static equals(a: GetCommentResponse | PlainMessage<GetCommentResponse> | undefined, b: GetCommentResponse | PlainMessage<GetCommentResponse> | undefined): boolean;
}

