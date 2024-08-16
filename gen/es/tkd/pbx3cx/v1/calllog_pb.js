// @generated by protoc-gen-es v2.0.0 with parameter "target=js+dts"
// @generated from file tkd/pbx3cx/v1/calllog.proto (package tkd.pbx3cx.v1, syntax proto3)
/* eslint-disable */

import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_descriptor, file_google_protobuf_duration, file_google_protobuf_empty, file_google_protobuf_field_mask, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import { file_tkd_common_v1_descriptor } from "../../common/v1/descriptor_pb";
import { file_tkd_common_v1_time_range } from "../../common/v1/time_range_pb";
import { file_tkd_customer_v1_customer } from "../../customer/v1/customer_pb";
import { file_tkd_idm_v1_user } from "../../idm/v1/user_pb";
import { file_buf_validate_validate } from "../../../buf/validate/validate_pb";

/**
 * Describes the file tkd/pbx3cx/v1/calllog.proto.
 */
export const file_tkd_pbx3cx_v1_calllog = /*@__PURE__*/
  fileDesc("Cht0a2QvcGJ4M2N4L3YxL2NhbGxsb2cucHJvdG8SDXRrZC5wYngzY3gudjEitQIKCUNhbGxFbnRyeRIKCgJpZBgBIAEoCRIOCgZjYWxsZXIYAiABKAkSFgoOaW5ib3VuZF9udW1iZXIYAyABKAkSLwoLcmVjZWl2ZWRfYXQYBCABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEisKCGR1cmF0aW9uGAUgASgLMhkuZ29vZ2xlLnByb3RvYnVmLkR1cmF0aW9uEhEKCWNhbGxfdHlwZRgGIAEoCRIVCg1hZ2VudF91c2VyX2lkGAcgASgJEhMKC2N1c3RvbWVyX2lkGAggASgJEhcKD2N1c3RvbWVyX3NvdXJjZRgJIAEoCRINCgVlcnJvchgKIAEoCBIXCg90cmFuc2Zlcl90YXJnZXQYCyABKAkSFgoOYWNjZXB0ZWRfYWdlbnQYDCABKAkicgoGT25DYWxsEiQKB3Byb2ZpbGUYASABKAsyEy50a2QuaWRtLnYxLlByb2ZpbGUSFwoPdHJhbnNmZXJfdGFyZ2V0GAIgASgJEikKBXVudGlsGAMgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcCJACg9DdXN0b21PdmVyd3JpdGUSFwoPdHJhbnNmZXJfdGFyZ2V0GAEgASgJEhQKDGRpc3BsYXlfbmFtZRgCIAEoCSK6AgoJT3ZlcndyaXRlEgoKAmlkGAEgASgJEigKBGZyb20YAiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEiYKAnRvGAMgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcBIwCgZjdXN0b20YBCABKAsyHi50a2QucGJ4M2N4LnYxLkN1c3RvbU92ZXJ3cml0ZUgAEhEKB3VzZXJfaWQYBSABKAlIABIuCgpjcmVhdGVkX2F0GAYgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcBIaChJjcmVhdGVkX2J5X3VzZXJfaWQYByABKAkSNAoOaW5ib3VuZF9udW1iZXIYCCABKAsyHC50a2QucGJ4M2N4LnYxLkluYm91bmROdW1iZXJCCAoGdGFyZ2V0ImoKDUluYm91bmROdW1iZXISDgoGbnVtYmVyGAEgASgJEhQKDGRpc3BsYXlfbmFtZRgCIAEoCRIZChFyb3N0ZXJfc2hpZnRfdGFncxgDIAMoCRIYChByb3N0ZXJfdHlwZV9uYW1lGAQgASgJIpgBChFSZWNvcmRDYWxsUmVxdWVzdBIQCghkdXJhdGlvbhgBIAEoCRIOCgZudW1iZXIYAiABKAkSDQoFYWdlbnQYAyABKAkSEQoJY2FsbF90eXBlGAQgASgJEhEKCWRhdGVfdGltZRgFIAEoCRITCgtjdXN0b21lcl9pZBgGIAEoCRIXCg9jdXN0b21lcl9zb3VyY2UYByABKAkiQAoZR2V0TG9nc0ZvckN1c3RvbWVyUmVxdWVzdBIOCgZzb3VyY2UYASABKAkSEwoCaWQYAiABKAlCB/r3GAPIAQEiRwoaR2V0TG9nc0ZvckN1c3RvbWVyUmVzcG9uc2USKQoHcmVzdWx0cxgBIAMoCzIYLnRrZC5wYngzY3gudjEuQ2FsbEVudHJ5Ii4KFUdldExvZ3NGb3JEYXRlUmVxdWVzdBIVCgRkYXRlGAEgASgJQgf69xgDyAEBIkMKFkdldExvZ3NGb3JEYXRlUmVzcG9uc2USKQoHcmVzdWx0cxgBIAMoCzIYLnRrZC5wYngzY3gudjEuQ2FsbEVudHJ5IocBChVTZWFyY2hDYWxsTG9nc1JlcXVlc3QSDAoEZGF0ZRgBIAEoCRIsCgp0aW1lX3JhbmdlGAIgASgLMhgudGtkLmNvbW1vbi52MS5UaW1lUmFuZ2USMgoMY3VzdG9tZXJfcmVmGAMgASgLMhwudGtkLmN1c3RvbWVyLnYxLkN1c3RvbWVyUmVmInEKFlNlYXJjaENhbGxMb2dzUmVzcG9uc2USKQoHcmVzdWx0cxgBIAMoCzIYLnRrZC5wYngzY3gudjEuQ2FsbEVudHJ5EiwKCWN1c3RvbWVycxgCIAMoCzIZLnRrZC5jdXN0b21lci52MS5DdXN0b21lciJTChBHZXRPbkNhbGxSZXF1ZXN0EgwKBGRhdGUYASABKAkSGQoRaWdub3JlX292ZXJ3cml0ZXMYAiABKAgSFgoOaW5ib3VuZF9udW1iZXIYAyABKAkihwEKEUdldE9uQ2FsbFJlc3BvbnNlEiYKB29uX2NhbGwYASADKAsyFS50a2QucGJ4M2N4LnYxLk9uQ2FsbBITCgtyb3N0ZXJfZGF0ZRgCIAEoCRIUCgxpc19vdmVyd3JpdGUYAyABKAgSHwoXcHJpbWFyeV90cmFuc2Zlcl90YXJnZXQYBCABKAki9AEKFkNyZWF0ZU92ZXJ3cml0ZVJlcXVlc3QSMAoGY3VzdG9tGAEgASgLMh4udGtkLnBieDNjeC52MS5DdXN0b21PdmVyd3JpdGVIABIRCgd1c2VyX2lkGAIgASgJSAASMQoEZnJvbRgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBCB/r3GAPIAQESLwoCdG8YBCABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wQgf69xgDyAEBEhYKDmluYm91bmRfbnVtYmVyGAUgASgJQhkKD3RyYW5zZmVyX3RhcmdldBIG+vcYAggBIkYKF0NyZWF0ZU92ZXJ3cml0ZVJlc3BvbnNlEisKCW92ZXJ3cml0ZRgBIAEoCzIYLnRrZC5wYngzY3gudjEuT3ZlcndyaXRlIrABChZEZWxldGVPdmVyd3JpdGVSZXF1ZXN0EhYKDG92ZXJ3cml0ZV9pZBgBIAEoCUgAEi8KCWFjdGl2ZV9hdBgCIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBIABI5Cg9pbmJvdW5kX251bWJlcnMYAyABKAsyIC50a2QucGJ4M2N4LnYxLkluYm91bmROdW1iZXJMaXN0QhIKCHNlbGVjdG9yEgb69xgCCAEiGQoXRGVsZXRlT3ZlcndyaXRlUmVzcG9uc2Ui3QEKE0dldE92ZXJ3cml0ZVJlcXVlc3QSFgoMb3ZlcndyaXRlX2lkGAEgASgJSAASLwoJYWN0aXZlX2F0GAIgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEgAEi4KCnRpbWVfcmFuZ2UYAyABKAsyGC50a2QuY29tbW9uLnYxLlRpbWVSYW5nZUgAEjkKD2luYm91bmRfbnVtYmVycxgEIAEoCzIgLnRrZC5wYngzY3gudjEuSW5ib3VuZE51bWJlckxpc3RCEgoIc2VsZWN0b3ISBvr3GAIIASIkChFJbmJvdW5kTnVtYmVyTGlzdBIPCgdudW1iZXJzGAEgAygJIkQKFEdldE92ZXJ3cml0ZVJlc3BvbnNlEiwKCm92ZXJ3cml0ZXMYASADKAsyGC50a2QucGJ4M2N4LnYxLk92ZXJ3cml0ZSKAAQoaQ3JlYXRlSW5ib3VuZE51bWJlclJlcXVlc3QSFwoGbnVtYmVyGAEgASgJQgf69xgDyAEBEhQKDGRpc3BsYXlfbmFtZRgCIAEoCRIZChFyb3N0ZXJfc2hpZnRfdGFncxgDIAMoCRIYChByb3N0ZXJfdHlwZV9uYW1lGAQgASgJIlMKG0NyZWF0ZUluYm91bmROdW1iZXJSZXNwb25zZRI0Cg5pbmJvdW5kX251bWJlchgBIAEoCzIcLnRrZC5wYngzY3gudjEuSW5ib3VuZE51bWJlciI1ChpEZWxldGVJbmJvdW5kTnVtYmVyUmVxdWVzdBIXCgZudW1iZXIYASABKAlCB/r3GAPIAQEiHQobRGVsZXRlSW5ib3VuZE51bWJlclJlc3BvbnNlIhoKGExpc3RJbmJvdW5kTnVtYmVyUmVxdWVzdCJSChlMaXN0SW5ib3VuZE51bWJlclJlc3BvbnNlEjUKD2luYm91bmRfbnVtYmVycxgBIAMoCzIcLnRrZC5wYngzY3gudjEuSW5ib3VuZE51bWJlciK1AQoaVXBkYXRlSW5ib3VuZE51bWJlclJlcXVlc3QSFwoGbnVtYmVyGAEgASgJQgf69xgDyAEBEhgKEG5ld19kaXNwbGF5X25hbWUYAiABKAkSGQoRcm9zdGVyX3NoaWZ0X3RhZ3MYAyADKAkSGAoQcm9zdGVyX3R5cGVfbmFtZRgEIAEoCRIvCgt1cGRhdGVfbWFzaxgKIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE1hc2siUwobVXBkYXRlSW5ib3VuZE51bWJlclJlc3BvbnNlEjQKDmluYm91bmRfbnVtYmVyGAEgASgLMhwudGtkLnBieDNjeC52MS5JbmJvdW5kTnVtYmVyMp0KCgtDYWxsU2VydmljZRJNCgpSZWNvcmRDYWxsEiAudGtkLnBieDNjeC52MS5SZWNvcmRDYWxsUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eSIFsn4CCAESWAoJR2V0T25DYWxsEh8udGtkLnBieDNjeC52MS5HZXRPbkNhbGxSZXF1ZXN0GiAudGtkLnBieDNjeC52MS5HZXRPbkNhbGxSZXNwb25zZSIIkAIBsn4CCAEScwoTQ3JlYXRlSW5ib3VuZE51bWJlchIpLnRrZC5wYngzY3gudjEuQ3JlYXRlSW5ib3VuZE51bWJlclJlcXVlc3QaKi50a2QucGJ4M2N4LnYxLkNyZWF0ZUluYm91bmROdW1iZXJSZXNwb25zZSIFsn4CCAIScwoTVXBkYXRlSW5ib3VuZE51bWJlchIpLnRrZC5wYngzY3gudjEuVXBkYXRlSW5ib3VuZE51bWJlclJlcXVlc3QaKi50a2QucGJ4M2N4LnYxLlVwZGF0ZUluYm91bmROdW1iZXJSZXNwb25zZSIFsn4CCAIScwoTRGVsZXRlSW5ib3VuZE51bWJlchIpLnRrZC5wYngzY3gudjEuRGVsZXRlSW5ib3VuZE51bWJlclJlcXVlc3QaKi50a2QucGJ4M2N4LnYxLkRlbGV0ZUluYm91bmROdW1iZXJSZXNwb25zZSIFsn4CCAISbQoRTGlzdEluYm91bmROdW1iZXISJy50a2QucGJ4M2N4LnYxLkxpc3RJbmJvdW5kTnVtYmVyUmVxdWVzdBooLnRrZC5wYngzY3gudjEuTGlzdEluYm91bmROdW1iZXJSZXNwb25zZSIFsn4CCAISZwoPQ3JlYXRlT3ZlcndyaXRlEiUudGtkLnBieDNjeC52MS5DcmVhdGVPdmVyd3JpdGVSZXF1ZXN0GiYudGtkLnBieDNjeC52MS5DcmVhdGVPdmVyd3JpdGVSZXNwb25zZSIFsn4CCAESZwoPRGVsZXRlT3ZlcndyaXRlEiUudGtkLnBieDNjeC52MS5EZWxldGVPdmVyd3JpdGVSZXF1ZXN0GiYudGtkLnBieDNjeC52MS5EZWxldGVPdmVyd3JpdGVSZXNwb25zZSIFsn4CCAESXgoMR2V0T3ZlcndyaXRlEiIudGtkLnBieDNjeC52MS5HZXRPdmVyd3JpdGVSZXF1ZXN0GiMudGtkLnBieDNjeC52MS5HZXRPdmVyd3JpdGVSZXNwb25zZSIFsn4CCAESZAoOR2V0TG9nc0ZvckRhdGUSJC50a2QucGJ4M2N4LnYxLkdldExvZ3NGb3JEYXRlUmVxdWVzdBolLnRrZC5wYngzY3gudjEuR2V0TG9nc0ZvckRhdGVSZXNwb25zZSIFsn4CCAESZAoOU2VhcmNoQ2FsbExvZ3MSJC50a2QucGJ4M2N4LnYxLlNlYXJjaENhbGxMb2dzUmVxdWVzdBolLnRrZC5wYngzY3gudjEuU2VhcmNoQ2FsbExvZ3NSZXNwb25zZSIFsn4CCAEScAoSR2V0TG9nc0ZvckN1c3RvbWVyEigudGtkLnBieDNjeC52MS5HZXRMb2dzRm9yQ3VzdG9tZXJSZXF1ZXN0GikudGtkLnBieDNjeC52MS5HZXRMb2dzRm9yQ3VzdG9tZXJSZXNwb25zZSIFsn4CCAEaJ7J+EjNjeC5kb2JlcnNiZXJnLnZldLp+DwoNaWRtX3N1cGVydXNlckK8AQoRY29tLnRrZC5wYngzY3gudjFCDENhbGxsb2dQcm90b1ABWkNnaXRodWIuY29tL3RpZXJrbGluaWstZG9iZXJzYmVyZy9hcGlzL2dlbi9nby90a2QvcGJ4M2N4L3YxO3BieDNjeHYxogIDVFBYqgINVGtkLlBieDNjeC5WMcoCDVRrZFxQYngzY3hcVjHiAhlUa2RcUGJ4M2N4XFYxXEdQQk1ldGFkYXRh6gIPVGtkOjpQYngzY3g6OlYxYgZwcm90bzM", [file_google_protobuf_descriptor, file_google_protobuf_timestamp, file_google_protobuf_duration, file_google_protobuf_field_mask, file_tkd_common_v1_descriptor, file_tkd_common_v1_time_range, file_tkd_customer_v1_customer, file_tkd_idm_v1_user, file_google_protobuf_empty, file_buf_validate_validate]);

/**
 * Describes the message tkd.pbx3cx.v1.CallEntry.
 * Use `create(CallEntrySchema)` to create a new message.
 */
export const CallEntrySchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 0);

/**
 * Describes the message tkd.pbx3cx.v1.OnCall.
 * Use `create(OnCallSchema)` to create a new message.
 */
export const OnCallSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 1);

/**
 * Describes the message tkd.pbx3cx.v1.CustomOverwrite.
 * Use `create(CustomOverwriteSchema)` to create a new message.
 */
export const CustomOverwriteSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 2);

/**
 * Describes the message tkd.pbx3cx.v1.Overwrite.
 * Use `create(OverwriteSchema)` to create a new message.
 */
export const OverwriteSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 3);

/**
 * Describes the message tkd.pbx3cx.v1.InboundNumber.
 * Use `create(InboundNumberSchema)` to create a new message.
 */
export const InboundNumberSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 4);

/**
 * Describes the message tkd.pbx3cx.v1.RecordCallRequest.
 * Use `create(RecordCallRequestSchema)` to create a new message.
 */
export const RecordCallRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 5);

/**
 * Describes the message tkd.pbx3cx.v1.GetLogsForCustomerRequest.
 * Use `create(GetLogsForCustomerRequestSchema)` to create a new message.
 */
export const GetLogsForCustomerRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 6);

/**
 * Describes the message tkd.pbx3cx.v1.GetLogsForCustomerResponse.
 * Use `create(GetLogsForCustomerResponseSchema)` to create a new message.
 */
export const GetLogsForCustomerResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 7);

/**
 * Describes the message tkd.pbx3cx.v1.GetLogsForDateRequest.
 * Use `create(GetLogsForDateRequestSchema)` to create a new message.
 */
export const GetLogsForDateRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 8);

/**
 * Describes the message tkd.pbx3cx.v1.GetLogsForDateResponse.
 * Use `create(GetLogsForDateResponseSchema)` to create a new message.
 */
export const GetLogsForDateResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 9);

/**
 * Describes the message tkd.pbx3cx.v1.SearchCallLogsRequest.
 * Use `create(SearchCallLogsRequestSchema)` to create a new message.
 */
export const SearchCallLogsRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 10);

/**
 * Describes the message tkd.pbx3cx.v1.SearchCallLogsResponse.
 * Use `create(SearchCallLogsResponseSchema)` to create a new message.
 */
export const SearchCallLogsResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 11);

/**
 * Describes the message tkd.pbx3cx.v1.GetOnCallRequest.
 * Use `create(GetOnCallRequestSchema)` to create a new message.
 */
export const GetOnCallRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 12);

/**
 * Describes the message tkd.pbx3cx.v1.GetOnCallResponse.
 * Use `create(GetOnCallResponseSchema)` to create a new message.
 */
export const GetOnCallResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 13);

/**
 * Describes the message tkd.pbx3cx.v1.CreateOverwriteRequest.
 * Use `create(CreateOverwriteRequestSchema)` to create a new message.
 */
export const CreateOverwriteRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 14);

/**
 * Describes the message tkd.pbx3cx.v1.CreateOverwriteResponse.
 * Use `create(CreateOverwriteResponseSchema)` to create a new message.
 */
export const CreateOverwriteResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 15);

/**
 * Describes the message tkd.pbx3cx.v1.DeleteOverwriteRequest.
 * Use `create(DeleteOverwriteRequestSchema)` to create a new message.
 */
export const DeleteOverwriteRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 16);

/**
 * Describes the message tkd.pbx3cx.v1.DeleteOverwriteResponse.
 * Use `create(DeleteOverwriteResponseSchema)` to create a new message.
 */
export const DeleteOverwriteResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 17);

/**
 * Describes the message tkd.pbx3cx.v1.GetOverwriteRequest.
 * Use `create(GetOverwriteRequestSchema)` to create a new message.
 */
export const GetOverwriteRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 18);

/**
 * Describes the message tkd.pbx3cx.v1.InboundNumberList.
 * Use `create(InboundNumberListSchema)` to create a new message.
 */
export const InboundNumberListSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 19);

/**
 * Describes the message tkd.pbx3cx.v1.GetOverwriteResponse.
 * Use `create(GetOverwriteResponseSchema)` to create a new message.
 */
export const GetOverwriteResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 20);

/**
 * Describes the message tkd.pbx3cx.v1.CreateInboundNumberRequest.
 * Use `create(CreateInboundNumberRequestSchema)` to create a new message.
 */
export const CreateInboundNumberRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 21);

/**
 * Describes the message tkd.pbx3cx.v1.CreateInboundNumberResponse.
 * Use `create(CreateInboundNumberResponseSchema)` to create a new message.
 */
export const CreateInboundNumberResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 22);

/**
 * Describes the message tkd.pbx3cx.v1.DeleteInboundNumberRequest.
 * Use `create(DeleteInboundNumberRequestSchema)` to create a new message.
 */
export const DeleteInboundNumberRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 23);

/**
 * Describes the message tkd.pbx3cx.v1.DeleteInboundNumberResponse.
 * Use `create(DeleteInboundNumberResponseSchema)` to create a new message.
 */
export const DeleteInboundNumberResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 24);

/**
 * Describes the message tkd.pbx3cx.v1.ListInboundNumberRequest.
 * Use `create(ListInboundNumberRequestSchema)` to create a new message.
 */
export const ListInboundNumberRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 25);

/**
 * Describes the message tkd.pbx3cx.v1.ListInboundNumberResponse.
 * Use `create(ListInboundNumberResponseSchema)` to create a new message.
 */
export const ListInboundNumberResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 26);

/**
 * Describes the message tkd.pbx3cx.v1.UpdateInboundNumberRequest.
 * Use `create(UpdateInboundNumberRequestSchema)` to create a new message.
 */
export const UpdateInboundNumberRequestSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 27);

/**
 * Describes the message tkd.pbx3cx.v1.UpdateInboundNumberResponse.
 * Use `create(UpdateInboundNumberResponseSchema)` to create a new message.
 */
export const UpdateInboundNumberResponseSchema = /*@__PURE__*/
  messageDesc(file_tkd_pbx3cx_v1_calllog, 28);

/**
 * @generated from service tkd.pbx3cx.v1.CallService
 */
export const CallService = /*@__PURE__*/
  serviceDesc(file_tkd_pbx3cx_v1_calllog, 0);

