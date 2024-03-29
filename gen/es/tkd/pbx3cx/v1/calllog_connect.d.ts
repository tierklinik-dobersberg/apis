// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=js+dts"
// @generated from file tkd/pbx3cx/v1/calllog.proto (package tkd.pbx3cx.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateOverwriteRequest, CreateOverwriteResponse, DeleteOverwriteRequest, DeleteOverwriteResponse, GetLogsForCustomerRequest, GetLogsForCustomerResponse, GetLogsForDateRequest, GetLogsForDateResponse, GetOnCallRequest, GetOnCallResponse, GetOverwriteRequest, GetOverwriteResponse, RecordCallRequest, SearchCallLogsRequest, SearchCallLogsResponse } from "./calllog_pb.js";
import { Empty, MethodIdempotency, MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service tkd.pbx3cx.v1.CallService
 */
export declare const CallService: {
  readonly typeName: "tkd.pbx3cx.v1.CallService",
  readonly methods: {
    /**
     * Recording APIS 
     *
     * @generated from rpc tkd.pbx3cx.v1.CallService.RecordCall
     */
    readonly recordCall: {
      readonly name: "RecordCall",
      readonly I: typeof RecordCallRequest,
      readonly O: typeof Empty,
      readonly kind: MethodKind.Unary,
    },
    /**
     * On-Duty/Call APIs
     *
     * @generated from rpc tkd.pbx3cx.v1.CallService.GetOnCall
     */
    readonly getOnCall: {
      readonly name: "GetOnCall",
      readonly I: typeof GetOnCallRequest,
      readonly O: typeof GetOnCallResponse,
      readonly kind: MethodKind.Unary,
      readonly idempotency: MethodIdempotency.NoSideEffects,
    },
    /**
     * Overwrite APIS
     *
     * @generated from rpc tkd.pbx3cx.v1.CallService.CreateOverwrite
     */
    readonly createOverwrite: {
      readonly name: "CreateOverwrite",
      readonly I: typeof CreateOverwriteRequest,
      readonly O: typeof CreateOverwriteResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.pbx3cx.v1.CallService.DeleteOverwrite
     */
    readonly deleteOverwrite: {
      readonly name: "DeleteOverwrite",
      readonly I: typeof DeleteOverwriteRequest,
      readonly O: typeof DeleteOverwriteResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.pbx3cx.v1.CallService.GetOverwrite
     */
    readonly getOverwrite: {
      readonly name: "GetOverwrite",
      readonly I: typeof GetOverwriteRequest,
      readonly O: typeof GetOverwriteResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * Calllog APIs
     *
     * @generated from rpc tkd.pbx3cx.v1.CallService.GetLogsForDate
     */
    readonly getLogsForDate: {
      readonly name: "GetLogsForDate",
      readonly I: typeof GetLogsForDateRequest,
      readonly O: typeof GetLogsForDateResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.pbx3cx.v1.CallService.SearchCallLogs
     */
    readonly searchCallLogs: {
      readonly name: "SearchCallLogs",
      readonly I: typeof SearchCallLogsRequest,
      readonly O: typeof SearchCallLogsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.pbx3cx.v1.CallService.GetLogsForCustomer
     */
    readonly getLogsForCustomer: {
      readonly name: "GetLogsForCustomer",
      readonly I: typeof GetLogsForCustomerRequest,
      readonly O: typeof GetLogsForCustomerResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

