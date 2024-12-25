// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file tkd/calendar/v1/holiday_service.proto (package tkd.calendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";
import { Date } from "../../common/v1/date_pb.js";

/**
 * HolidayType specifies the type of a public holiday.
 *
 * @generated from enum tkd.calendar.v1.HolidayType
 */
export const HolidayType = /*@__PURE__*/ proto3.makeEnum(
  "tkd.calendar.v1.HolidayType",
  [
    {no: 0, name: "HOLIDAY_TYPE_UNSPECIFIED"},
    {no: 1, name: "PUBLIC"},
    {no: 2, name: "BANK"},
    {no: 3, name: "SCHOOL"},
    {no: 4, name: "AUTHORITIES"},
    {no: 5, name: "OPTIONAL"},
    {no: 6, name: "OBSERVANCE"},
  ],
);

/**
 * PublicHoliday describes a public holiday at a specified country.
 *
 * @generated from message tkd.calendar.v1.PublicHoliday
 */
export const PublicHoliday = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.calendar.v1.PublicHoliday",
  () => [
    { no: 1, name: "date", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "local_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "country_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "fixed", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 6, name: "global", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 7, name: "type", kind: "enum", T: proto3.getEnumType(HolidayType) },
  ],
);

/**
 * GetHolidayRequest is the request message for the GetHoliday RPC.
 *
 * @generated from message tkd.calendar.v1.GetHolidayRequest
 */
export const GetHolidayRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.calendar.v1.GetHolidayRequest",
  () => [
    { no: 1, name: "year", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "month", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "country_code", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * GetHolidayResponse is the response message of the GetHoliday RPC and contains
 * a list of public holidays.
 *
 * @generated from message tkd.calendar.v1.GetHolidayResponse
 */
export const GetHolidayResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.calendar.v1.GetHolidayResponse",
  () => [
    { no: 1, name: "holidays", kind: "message", T: PublicHoliday, repeated: true },
  ],
);

/**
 * NumberOfWorkDaysRequest is the request message for the NumberOfWorkDays RPC.
 *
 * @generated from message tkd.calendar.v1.NumberOfWorkDaysRequest
 */
export const NumberOfWorkDaysRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.calendar.v1.NumberOfWorkDaysRequest",
  () => [
    { no: 1, name: "country", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "from", kind: "message", T: Timestamp },
    { no: 3, name: "to", kind: "message", T: Timestamp },
  ],
);

/**
 * NumberOfWorkDaysResponse is the response message of the NumberOfWorkDays RPC.
 *
 * @generated from message tkd.calendar.v1.NumberOfWorkDaysResponse
 */
export const NumberOfWorkDaysResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.calendar.v1.NumberOfWorkDaysResponse",
  () => [
    { no: 1, name: "number_of_work_days", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 2, name: "number_of_weekend_days", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 3, name: "number_of_holidays", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
  ],
);

/**
 * @generated from message tkd.calendar.v1.IsHolidayRequest
 */
export const IsHolidayRequest = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.calendar.v1.IsHolidayRequest",
  () => [
    { no: 1, name: "date", kind: "message", T: Date },
  ],
);

/**
 * @generated from message tkd.calendar.v1.IsHolidayResponse
 */
export const IsHolidayResponse = /*@__PURE__*/ proto3.makeMessageType(
  "tkd.calendar.v1.IsHolidayResponse",
  () => [
    { no: 1, name: "is_holiday", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "holiday", kind: "message", T: PublicHoliday },
    { no: 3, name: "queried_date", kind: "message", T: Date },
  ],
);

