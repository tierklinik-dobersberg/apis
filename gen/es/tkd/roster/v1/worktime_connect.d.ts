// @generated by protoc-gen-connect-es v0.12.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/worktime.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetVacationCreditsLeftRequest, GetVacationCreditsLeftResponse, GetWorkTimeRequest, GetWorkTimeResponse, SetWorkTimeRequest, SetWorkTimeResponse } from "./worktime_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service tkd.roster.v1.WorkTimeService
 */
export declare const WorkTimeService: {
  readonly typeName: "tkd.roster.v1.WorkTimeService",
  readonly methods: {
    /**
     * @generated from rpc tkd.roster.v1.WorkTimeService.SetWorkTime
     */
    readonly setWorkTime: {
      readonly name: "SetWorkTime",
      readonly I: typeof SetWorkTimeRequest,
      readonly O: typeof SetWorkTimeResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.WorkTimeService.GetWorkTime
     */
    readonly getWorkTime: {
      readonly name: "GetWorkTime",
      readonly I: typeof GetWorkTimeRequest,
      readonly O: typeof GetWorkTimeResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.WorkTimeService.GetVacationCreditsLeft
     */
    readonly getVacationCreditsLeft: {
      readonly name: "GetVacationCreditsLeft",
      readonly I: typeof GetVacationCreditsLeftRequest,
      readonly O: typeof GetVacationCreditsLeftResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

