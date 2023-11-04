// @generated by protoc-gen-es v1.4.1 with parameter "target=js+dts"
// @generated from file tkd/comment/v1/comment.proto (package tkd.comment.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FieldMask, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum tkd.comment.v1.NotificationType
 */
export const NotificationType = proto3.makeEnum(
  "tkd.comment.v1.NotificationType",
  [
    {no: 0, name: "NOTIFICATION_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "NOTIFICATION_TYPE_MAIL", localName: "MAIL"},
    {no: 2, name: "NOTIFICATION_TYPE_SMS", localName: "SMS"},
  ],
);

/**
 * @generated from message tkd.comment.v1.Scope
 */
export const Scope = proto3.makeMessageType(
  "tkd.comment.v1.Scope",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "notifcation_type", kind: "enum", T: proto3.getEnumType(NotificationType) },
    { no: 4, name: "view_comment_url_template", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.comment.v1.Comment
 */
export const Comment = proto3.makeMessageType(
  "tkd.comment.v1.Comment",
  () => [
    { no: 1, name: "scope", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "parent_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "created_at", kind: "message", T: Timestamp },
    { no: 6, name: "creator_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.comment.v1.CommentTree
 */
export const CommentTree = proto3.makeMessageType(
  "tkd.comment.v1.CommentTree",
  () => [
    { no: 1, name: "comment", kind: "message", T: Comment },
    { no: 2, name: "answers", kind: "message", T: CommentTree, repeated: true },
  ],
);

/**
 * CreateScopeRequest is used to create a new comment scope.
 *
 * @generated from message tkd.comment.v1.CreateScopeRequest
 */
export const CreateScopeRequest = proto3.makeMessageType(
  "tkd.comment.v1.CreateScopeRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "notifcation_type", kind: "enum", T: proto3.getEnumType(NotificationType) },
    { no: 4, name: "view_comment_url_template", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * CreateScopeResponse is returned when a new comment scope has been created
 * and includes the created scope definition.
 *
 * @generated from message tkd.comment.v1.CreateScopeResponse
 */
export const CreateScopeResponse = proto3.makeMessageType(
  "tkd.comment.v1.CreateScopeResponse",
  () => [
    { no: 1, name: "scope", kind: "message", T: Scope },
  ],
);

/**
 *
 *
 * @generated from message tkd.comment.v1.ListScopeRequest
 */
export const ListScopeRequest = proto3.makeMessageType(
  "tkd.comment.v1.ListScopeRequest",
  [],
);

/**
 * @generated from message tkd.comment.v1.ListScopeResponse
 */
export const ListScopeResponse = proto3.makeMessageType(
  "tkd.comment.v1.ListScopeResponse",
  () => [
    { no: 1, name: "scopes", kind: "message", T: Scope, repeated: true },
  ],
);

/**
 * Delete Scope
 *
 * @generated from message tkd.comment.v1.DeleteScopeRequest
 */
export const DeleteScopeRequest = proto3.makeMessageType(
  "tkd.comment.v1.DeleteScopeRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message tkd.comment.v1.DeleteScopeResponse
 */
export const DeleteScopeResponse = proto3.makeMessageType(
  "tkd.comment.v1.DeleteScopeResponse",
  [],
);

/**
 * Create Comment
 *
 * @generated from message tkd.comment.v1.CreateCommentRequest
 */
export const CreateCommentRequest = proto3.makeMessageType(
  "tkd.comment.v1.CreateCommentRequest",
  () => [
    { no: 1, name: "scope", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "kind" },
    { no: 2, name: "parent_id", kind: "scalar", T: 9 /* ScalarType.STRING */, oneof: "kind" },
    { no: 3, name: "content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "read_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message tkd.comment.v1.CreateCommentResponse
 */
export const CreateCommentResponse = proto3.makeMessageType(
  "tkd.comment.v1.CreateCommentResponse",
  () => [
    { no: 1, name: "comment", kind: "message", T: Comment },
  ],
);

/**
 * List Comments
 *
 * @generated from message tkd.comment.v1.ListCommentsRequest
 */
export const ListCommentsRequest = proto3.makeMessageType(
  "tkd.comment.v1.ListCommentsRequest",
  () => [
    { no: 1, name: "scope", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "recurse", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "render_html", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message tkd.comment.v1.ListCommentsResponse
 */
export const ListCommentsResponse = proto3.makeMessageType(
  "tkd.comment.v1.ListCommentsResponse",
  () => [
    { no: 2, name: "result", kind: "message", T: CommentTree, repeated: true },
  ],
);

/**
 * Get Comments
 *
 * @generated from message tkd.comment.v1.GetCommentRequest
 */
export const GetCommentRequest = proto3.makeMessageType(
  "tkd.comment.v1.GetCommentRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "recurse", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "render_html", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message tkd.comment.v1.GetCommentResponse
 */
export const GetCommentResponse = proto3.makeMessageType(
  "tkd.comment.v1.GetCommentResponse",
  () => [
    { no: 2, name: "result", kind: "message", T: CommentTree },
  ],
);

