// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/offtime.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddOffTimeCostsRequest, AddOffTimeCostsResponse, ApproveOrRejectRequest, ApproveOrRejectResponse, CreateOffTimeRequestRequest, CreateOffTimeRequestResponse, DeleteOffTimeCostsRequest, DeleteOffTimeCostsResponse, DeleteOffTimeRequestRequest, DeleteOffTimeRequestResponse, FindOffTimeRequestsRequest, FindOffTimeRequestsResponse, GetOffTimeCostsRequest, GetOffTimeCostsResponse, GetOffTimeEntryRequest, GetOffTimeEntryResponse, UpdateOffTimeRequestRequest, UpdateOffTimeRequestResponse } from "./offtime_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service tkd.roster.v1.OffTimeService
 */
export declare const OffTimeService: {
  readonly typeName: "tkd.roster.v1.OffTimeService",
  readonly methods: {
    /**
     * @generated from rpc tkd.roster.v1.OffTimeService.GetOffTimeEntry
     */
    readonly getOffTimeEntry: {
      readonly name: "GetOffTimeEntry",
      readonly I: typeof GetOffTimeEntryRequest,
      readonly O: typeof GetOffTimeEntryResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.OffTimeService.CreateOffTimeRequest
     */
    readonly createOffTimeRequest: {
      readonly name: "CreateOffTimeRequest",
      readonly I: typeof CreateOffTimeRequestRequest,
      readonly O: typeof CreateOffTimeRequestResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.OffTimeService.UpdateOffTimeRequest
     */
    readonly updateOffTimeRequest: {
      readonly name: "UpdateOffTimeRequest",
      readonly I: typeof UpdateOffTimeRequestRequest,
      readonly O: typeof UpdateOffTimeRequestResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.OffTimeService.DeleteOffTimeRequest
     */
    readonly deleteOffTimeRequest: {
      readonly name: "DeleteOffTimeRequest",
      readonly I: typeof DeleteOffTimeRequestRequest,
      readonly O: typeof DeleteOffTimeRequestResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.OffTimeService.FindOffTimeRequests
     */
    readonly findOffTimeRequests: {
      readonly name: "FindOffTimeRequests",
      readonly I: typeof FindOffTimeRequestsRequest,
      readonly O: typeof FindOffTimeRequestsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.OffTimeService.ApproveOrReject
     */
    readonly approveOrReject: {
      readonly name: "ApproveOrReject",
      readonly I: typeof ApproveOrRejectRequest,
      readonly O: typeof ApproveOrRejectResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.OffTimeService.AddOffTimeCosts
     */
    readonly addOffTimeCosts: {
      readonly name: "AddOffTimeCosts",
      readonly I: typeof AddOffTimeCostsRequest,
      readonly O: typeof AddOffTimeCostsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.OffTimeService.GetOffTimeCosts
     */
    readonly getOffTimeCosts: {
      readonly name: "GetOffTimeCosts",
      readonly I: typeof GetOffTimeCostsRequest,
      readonly O: typeof GetOffTimeCostsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.OffTimeService.DeleteOffTimeCosts
     */
    readonly deleteOffTimeCosts: {
      readonly name: "DeleteOffTimeCosts",
      readonly I: typeof DeleteOffTimeCostsRequest,
      readonly O: typeof DeleteOffTimeCostsResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

