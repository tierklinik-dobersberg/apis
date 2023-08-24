// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/constraint.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateConstraintRequest, CreateConstraintResponse, DeleteConstraintRequest, DeleteConstraintResponse, FindConstraintsRequest, FindConstraintsResponse, UpdateConstraintRequest, UpdateConstraintResponse } from "./constraint_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service tkd.roster.v1.ConstraintService
 */
export declare const ConstraintService: {
  readonly typeName: "tkd.roster.v1.ConstraintService",
  readonly methods: {
    /**
     * @generated from rpc tkd.roster.v1.ConstraintService.CreateConstraint
     */
    readonly createConstraint: {
      readonly name: "CreateConstraint",
      readonly I: typeof CreateConstraintRequest,
      readonly O: typeof CreateConstraintResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.ConstraintService.UpdateConstraint
     */
    readonly updateConstraint: {
      readonly name: "UpdateConstraint",
      readonly I: typeof UpdateConstraintRequest,
      readonly O: typeof UpdateConstraintResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.ConstraintService.DeleteConstraint
     */
    readonly deleteConstraint: {
      readonly name: "DeleteConstraint",
      readonly I: typeof DeleteConstraintRequest,
      readonly O: typeof DeleteConstraintResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.ConstraintService.FindConstraints
     */
    readonly findConstraints: {
      readonly name: "FindConstraints",
      readonly I: typeof FindConstraintsRequest,
      readonly O: typeof FindConstraintsResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

