// @generated by protoc-gen-es v2.0.0 with parameter "target=js+dts"
// @generated from file tkd/common/v1/descriptor.proto (package tkd.common.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenExtension, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";
import type { FieldMask, MessageOptions, MethodOptions, ServiceOptions } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file tkd/common/v1/descriptor.proto.
 */
export declare const file_tkd_common_v1_descriptor: GenFile;

/**
 * @generated from message tkd.common.v1.ServiceAuthDecorator
 */
export declare type ServiceAuthDecorator = Message<"tkd.common.v1.ServiceAuthDecorator"> & {
  /**
   * @generated from field: repeated string admin_roles = 1;
   */
  adminRoles: string[];
};

/**
 * Describes the message tkd.common.v1.ServiceAuthDecorator.
 * Use `create(ServiceAuthDecoratorSchema)` to create a new message.
 */
export declare const ServiceAuthDecoratorSchema: GenMessage<ServiceAuthDecorator>;

/**
 * @generated from message tkd.common.v1.AuthDecorator
 */
export declare type AuthDecorator = Message<"tkd.common.v1.AuthDecorator"> & {
  /**
   * @generated from field: tkd.common.v1.AuthRequirement require = 1;
   */
  require: AuthRequirement;

  /**
   * @generated from field: repeated string allowed_roles = 2;
   */
  allowedRoles: string[];

  /**
   * @generated from field: repeated string admin_roles = 3;
   */
  adminRoles: string[];
};

/**
 * Describes the message tkd.common.v1.AuthDecorator.
 * Use `create(AuthDecoratorSchema)` to create a new message.
 */
export declare const AuthDecoratorSchema: GenMessage<AuthDecorator>;

/**
 * @generated from message tkd.common.v1.PrivacyACL
 */
export declare type PrivacyACL = Message<"tkd.common.v1.PrivacyACL"> & {
  /**
   * The field mask defined which field are accessible to
   * anyone (maybe authenticated) if none of the other
   * settings allow access to all fields.
   *
   * @generated from field: google.protobuf.FieldMask field_mask = 1;
   */
  fieldMask?: FieldMask;

  /**
   * The name of the owner ID field. If set, and the authenticated
   * user has the same ID, access to all fields is permitted.
   *
   * @generated from field: string owner_field_name = 2;
   */
  ownerFieldName: string;

  /**
   * A list of roles that have access to all fields.
   *
   * @generated from field: repeated string allowed_roles = 3;
   */
  allowedRoles: string[];
};

/**
 * Describes the message tkd.common.v1.PrivacyACL.
 * Use `create(PrivacyACLSchema)` to create a new message.
 */
export declare const PrivacyACLSchema: GenMessage<PrivacyACL>;

/**
 * @generated from enum tkd.common.v1.AuthRequirement
 */
export enum AuthRequirement {
  /**
   * @generated from enum value: AUTH_REQ_UNSPECIFIED = 0;
   */
  AUTH_REQ_UNSPECIFIED = 0,

  /**
   * @generated from enum value: AUTH_REQ_REQUIRED = 1;
   */
  AUTH_REQ_REQUIRED = 1,

  /**
   * @generated from enum value: AUTH_REQ_ADMIN = 2;
   */
  AUTH_REQ_ADMIN = 2,
}

/**
 * Describes the enum tkd.common.v1.AuthRequirement.
 */
export declare const AuthRequirementSchema: GenEnum<AuthRequirement>;

/**
 * @generated from extension: string default_host = 2022;
 */
export declare const default_host: GenExtension<ServiceOptions, string>;

/**
 * @generated from extension: tkd.common.v1.ServiceAuthDecorator service_auth = 2023;
 */
export declare const service_auth: GenExtension<ServiceOptions, ServiceAuthDecorator>;

/**
 * @generated from extension: tkd.common.v1.AuthDecorator auth = 2022;
 */
export declare const auth: GenExtension<MethodOptions, AuthDecorator>;

/**
 * @generated from extension: tkd.common.v1.PrivacyACL readable = 2026;
 */
export declare const readable: GenExtension<MessageOptions, PrivacyACL>;

/**
 * @generated from extension: tkd.common.v1.PrivacyACL self_service = 2027;
 */
export declare const self_service: GenExtension<MessageOptions, PrivacyACL>;

