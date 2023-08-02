// @generated by protoc-gen-connect-es v0.12.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/workshift.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateWorkShiftRequest, CreateWorkShiftResponse, DeleteWorkShiftRequest, DeleteWorkShiftResponse, ListWorkShiftsRequest, ListWorkShiftsResponse, UpdateWorkShiftRequest, UpdateWorkShiftResponse } from "./workshift_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service tkd.roster.v1.WorkShiftService
 */
export const WorkShiftService = {
  typeName: "tkd.roster.v1.WorkShiftService",
  methods: {
    /**
     * @generated from rpc tkd.roster.v1.WorkShiftService.ListWorkShifts
     */
    listWorkShifts: {
      name: "ListWorkShifts",
      I: ListWorkShiftsRequest,
      O: ListWorkShiftsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.WorkShiftService.CreateWorkShift
     */
    createWorkShift: {
      name: "CreateWorkShift",
      I: CreateWorkShiftRequest,
      O: CreateWorkShiftResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.WorkShiftService.UpdateWorkShift
     */
    updateWorkShift: {
      name: "UpdateWorkShift",
      I: UpdateWorkShiftRequest,
      O: UpdateWorkShiftResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.WorkShiftService.DeleteWorkShift
     */
    deleteWorkShift: {
      name: "DeleteWorkShift",
      I: DeleteWorkShiftRequest,
      O: DeleteWorkShiftResponse,
      kind: MethodKind.Unary,
    },
  }
};
