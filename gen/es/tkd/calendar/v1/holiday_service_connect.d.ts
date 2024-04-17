// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=js+dts"
// @generated from file tkd/calendar/v1/holiday_service.proto (package tkd.calendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetHolidayRequest, GetHolidayResponse, NumberOfWorkDaysRequest, NumberOfWorkDaysResponse } from "./holiday_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * HolidayService allows to query a list of public holidays.
 *
 * @generated from service tkd.calendar.v1.HolidayService
 */
export declare const HolidayService: {
  readonly typeName: "tkd.calendar.v1.HolidayService",
  readonly methods: {
    /**
     * GetHoliday returns a list of public holidays at a specifed year.
     *
     * @generated from rpc tkd.calendar.v1.HolidayService.GetHoliday
     */
    readonly getHoliday: {
      readonly name: "GetHoliday",
      readonly I: typeof GetHolidayRequest,
      readonly O: typeof GetHolidayResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * NumberOfWorkDays calculates the number of working days within a specified
     * time range taking weekends and public-holidays into account.
     *
     * @generated from rpc tkd.calendar.v1.HolidayService.NumberOfWorkDays
     */
    readonly numberOfWorkDays: {
      readonly name: "NumberOfWorkDays",
      readonly I: typeof NumberOfWorkDaysRequest,
      readonly O: typeof NumberOfWorkDaysResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

