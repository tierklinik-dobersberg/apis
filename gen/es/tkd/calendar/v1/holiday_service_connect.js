// @generated by protoc-gen-connect-es v0.12.0 with parameter "target=js+dts"
// @generated from file tkd/calendar/v1/holiday_service.proto (package tkd.calendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetRequest, GetResponse } from "./holiday_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service tkd.calendar.v1.HolidayService
 */
export const HolidayService = {
  typeName: "tkd.calendar.v1.HolidayService",
  methods: {
    /**
     * @generated from rpc tkd.calendar.v1.HolidayService.Get
     */
    get: {
      name: "Get",
      I: GetRequest,
      O: GetResponse,
      kind: MethodKind.Unary,
    },
  }
};

