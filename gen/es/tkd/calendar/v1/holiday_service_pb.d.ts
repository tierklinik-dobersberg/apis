// @generated by protoc-gen-es v1.3.0 with parameter "target=js+dts"
// @generated from file tkd/calendar/v1/holiday_service.proto (package tkd.calendar.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum tkd.calendar.v1.HolidayType
 */
export declare enum HolidayType {
  /**
   * @generated from enum value: HOLIDAY_TYPE_UNSPECIFIED = 0;
   */
  HOLIDAY_TYPE_UNSPECIFIED = 0,

  /**
   * @generated from enum value: PUBLIC = 1;
   */
  PUBLIC = 1,

  /**
   * @generated from enum value: BANK = 2;
   */
  BANK = 2,

  /**
   * @generated from enum value: SCHOOL = 3;
   */
  SCHOOL = 3,

  /**
   * @generated from enum value: AUTHORITIES = 4;
   */
  AUTHORITIES = 4,

  /**
   * @generated from enum value: OPTIONAL = 5;
   */
  OPTIONAL = 5,

  /**
   * @generated from enum value: OBSERVANCE = 6;
   */
  OBSERVANCE = 6,
}

/**
 * @generated from message tkd.calendar.v1.PublicHoliday
 */
export declare class PublicHoliday extends Message<PublicHoliday> {
  /**
   * YYYY-MM-DD
   *
   * @generated from field: string date = 1;
   */
  date: string;

  /**
   * @generated from field: string local_name = 2;
   */
  localName: string;

  /**
   * @generated from field: string name = 3;
   */
  name: string;

  /**
   * @generated from field: string country_code = 4;
   */
  countryCode: string;

  /**
   * @generated from field: bool fixed = 5;
   */
  fixed: boolean;

  /**
   * @generated from field: bool global = 6;
   */
  global: boolean;

  /**
   * @generated from field: tkd.calendar.v1.HolidayType type = 7;
   */
  type: HolidayType;

  constructor(data?: PartialMessage<PublicHoliday>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.PublicHoliday";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PublicHoliday;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PublicHoliday;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PublicHoliday;

  static equals(a: PublicHoliday | PlainMessage<PublicHoliday> | undefined, b: PublicHoliday | PlainMessage<PublicHoliday> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.GetHolidayRequest
 */
export declare class GetHolidayRequest extends Message<GetHolidayRequest> {
  /**
   * @generated from field: uint64 year = 1;
   */
  year: bigint;

  /**
   * @generated from field: uint64 month = 2;
   */
  month: bigint;

  constructor(data?: PartialMessage<GetHolidayRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.GetHolidayRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetHolidayRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetHolidayRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetHolidayRequest;

  static equals(a: GetHolidayRequest | PlainMessage<GetHolidayRequest> | undefined, b: GetHolidayRequest | PlainMessage<GetHolidayRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.GetHolidayResponse
 */
export declare class GetHolidayResponse extends Message<GetHolidayResponse> {
  /**
   * @generated from field: repeated tkd.calendar.v1.PublicHoliday holidays = 1;
   */
  holidays: PublicHoliday[];

  constructor(data?: PartialMessage<GetHolidayResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.GetHolidayResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetHolidayResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetHolidayResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetHolidayResponse;

  static equals(a: GetHolidayResponse | PlainMessage<GetHolidayResponse> | undefined, b: GetHolidayResponse | PlainMessage<GetHolidayResponse> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.NumberOfWorkDaysRequest
 */
export declare class NumberOfWorkDaysRequest extends Message<NumberOfWorkDaysRequest> {
  /**
   * @generated from field: string country = 1;
   */
  country: string;

  /**
   * @generated from field: google.protobuf.Timestamp from = 2;
   */
  from?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp to = 3;
   */
  to?: Timestamp;

  constructor(data?: PartialMessage<NumberOfWorkDaysRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.NumberOfWorkDaysRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NumberOfWorkDaysRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NumberOfWorkDaysRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NumberOfWorkDaysRequest;

  static equals(a: NumberOfWorkDaysRequest | PlainMessage<NumberOfWorkDaysRequest> | undefined, b: NumberOfWorkDaysRequest | PlainMessage<NumberOfWorkDaysRequest> | undefined): boolean;
}

/**
 * @generated from message tkd.calendar.v1.NumberOfWorkDaysResponse
 */
export declare class NumberOfWorkDaysResponse extends Message<NumberOfWorkDaysResponse> {
  /**
   * @generated from field: uint32 number_of_work_days = 1;
   */
  numberOfWorkDays: number;

  /**
   * @generated from field: uint32 number_of_weekend_days = 2;
   */
  numberOfWeekendDays: number;

  /**
   * @generated from field: uint32 number_of_holidays = 3;
   */
  numberOfHolidays: number;

  constructor(data?: PartialMessage<NumberOfWorkDaysResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.calendar.v1.NumberOfWorkDaysResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): NumberOfWorkDaysResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): NumberOfWorkDaysResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): NumberOfWorkDaysResponse;

  static equals(a: NumberOfWorkDaysResponse | PlainMessage<NumberOfWorkDaysResponse> | undefined, b: NumberOfWorkDaysResponse | PlainMessage<NumberOfWorkDaysResponse> | undefined): boolean;
}

