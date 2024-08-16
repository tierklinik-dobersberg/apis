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
  fileDesc("Cht0a2QvcGJ4M2N4L3YxL2NhbGxsb2cucHJvdG8SDXRrZC5wYngzY3gudjEi4gIKCUNhbGxFbnRyeRIKCgJpZBgBIAEoCRIOCgZjYWxsZXIYAiABKAkSFgoOaW5ib3VuZF9udW1iZXIYAyABKAkSLwoLcmVjZWl2ZWRfYXQYBCABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEisKCGR1cmF0aW9uGAUgASgLMhkuZ29vZ2xlLnByb3RvYnVmLkR1cmF0aW9uEhEKCWNhbGxfdHlwZRgGIAEoCRIVCg1hZ2VudF91c2VyX2lkGAcgASgJEhMKC2N1c3RvbWVyX2lkGAggASgJEhcKD2N1c3RvbWVyX3NvdXJjZRgJIAEoCRINCgVlcnJvchgKIAEoCBIXCg90cmFuc2Zlcl90YXJnZXQYCyABKAkSFgoOYWNjZXB0ZWRfYWdlbnQYDCABKAkSKwoIY3VzdG9tZXIYDSABKAsyGS50a2QuY3VzdG9tZXIudjEuQ3VzdG9tZXIicgoGT25DYWxsEiQKB3Byb2ZpbGUYASABKAsyEy50a2QuaWRtLnYxLlByb2ZpbGUSFwoPdHJhbnNmZXJfdGFyZ2V0GAIgASgJEikKBXVudGlsGAMgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcCJACg9DdXN0b21PdmVyd3JpdGUSFwoPdHJhbnNmZXJfdGFyZ2V0GAEgASgJEhQKDGRpc3BsYXlfbmFtZRgCIAEoCSK6AgoJT3ZlcndyaXRlEgoKAmlkGAEgASgJEigKBGZyb20YAiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEiYKAnRvGAMgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcBIwCgZjdXN0b20YBCABKAsyHi50a2QucGJ4M2N4LnYxLkN1c3RvbU92ZXJ3cml0ZUgAEhEKB3VzZXJfaWQYBSABKAlIABIuCgpjcmVhdGVkX2F0GAYgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcBIaChJjcmVhdGVkX2J5X3VzZXJfaWQYByABKAkSNAoOaW5ib3VuZF9udW1iZXIYCCABKAsyHC50a2QucGJ4M2N4LnYxLkluYm91bmROdW1iZXJCCAoGdGFyZ2V0ImoKDUluYm91bmROdW1iZXISDgoGbnVtYmVyGAEgASgJEhQKDGRpc3BsYXlfbmFtZRgCIAEoCRIZChFyb3N0ZXJfc2hpZnRfdGFncxgDIAMoCRIYChByb3N0ZXJfdHlwZV9uYW1lGAQgASgJIpgBChFSZWNvcmRDYWxsUmVxdWVzdBIQCghkdXJhdGlvbhgBIAEoCRIOCgZudW1iZXIYAiABKAkSDQoFYWdlbnQYAyABKAkSEQoJY2FsbF90eXBlGAQgASgJEhEKCWRhdGVfdGltZRgFIAEoCRITCgtjdXN0b21lcl9pZBgGIAEoCRIXCg9jdXN0b21lcl9zb3VyY2UYByABKAkiQAoZR2V0TG9nc0ZvckN1c3RvbWVyUmVxdWVzdBIOCgZzb3VyY2UYASABKAkSEwoCaWQYAiABKAlCB/r3GAPIAQEiRwoaR2V0TG9nc0ZvckN1c3RvbWVyUmVzcG9uc2USKQoHcmVzdWx0cxgBIAMoCzIYLnRrZC5wYngzY3gudjEuQ2FsbEVudHJ5Ii4KFUdldExvZ3NGb3JEYXRlUmVxdWVzdBIVCgRkYXRlGAEgASgJQgf69xgDyAEBIkMKFkdldExvZ3NGb3JEYXRlUmVzcG9uc2USKQoHcmVzdWx0cxgBIAMoCzIYLnRrZC5wYngzY3gudjEuQ2FsbEVudHJ5IocBChVTZWFyY2hDYWxsTG9nc1JlcXVlc3QSDAoEZGF0ZRgBIAEoCRIsCgp0aW1lX3JhbmdlGAIgASgLMhgudGtkLmNvbW1vbi52MS5UaW1lUmFuZ2USMgoMY3VzdG9tZXJfcmVmGAMgASgLMhwudGtkLmN1c3RvbWVyLnYxLkN1c3RvbWVyUmVmIkMKFlNlYXJjaENhbGxMb2dzUmVzcG9uc2USKQoHcmVzdWx0cxgBIAMoCzIYLnRrZC5wYngzY3gudjEuQ2FsbEVudHJ5IlMKEEdldE9uQ2FsbFJlcXVlc3QSDAoEZGF0ZRgBIAEoCRIZChFpZ25vcmVfb3ZlcndyaXRlcxgCIAEoCBIWCg5pbmJvdW5kX251bWJlchgDIAEoCSKHAQoRR2V0T25DYWxsUmVzcG9uc2USJgoHb25fY2FsbBgBIAMoCzIVLnRrZC5wYngzY3gudjEuT25DYWxsEhMKC3Jvc3Rlcl9kYXRlGAIgASgJEhQKDGlzX292ZXJ3cml0ZRgDIAEoCBIfChdwcmltYXJ5X3RyYW5zZmVyX3RhcmdldBgEIAEoCSL0AQoWQ3JlYXRlT3ZlcndyaXRlUmVxdWVzdBIwCgZjdXN0b20YASABKAsyHi50a2QucGJ4M2N4LnYxLkN1c3RvbU92ZXJ3cml0ZUgAEhEKB3VzZXJfaWQYAiABKAlIABIxCgRmcm9tGAMgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEIH+vcYA8gBARIvCgJ0bxgEIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXBCB/r3GAPIAQESFgoOaW5ib3VuZF9udW1iZXIYBSABKAlCGQoPdHJhbnNmZXJfdGFyZ2V0Egb69xgCCAEiRgoXQ3JlYXRlT3ZlcndyaXRlUmVzcG9uc2USKwoJb3ZlcndyaXRlGAEgASgLMhgudGtkLnBieDNjeC52MS5PdmVyd3JpdGUisAEKFkRlbGV0ZU92ZXJ3cml0ZVJlcXVlc3QSFgoMb3ZlcndyaXRlX2lkGAEgASgJSAASLwoJYWN0aXZlX2F0GAIgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEgAEjkKD2luYm91bmRfbnVtYmVycxgDIAEoCzIgLnRrZC5wYngzY3gudjEuSW5ib3VuZE51bWJlckxpc3RCEgoIc2VsZWN0b3ISBvr3GAIIASIZChdEZWxldGVPdmVyd3JpdGVSZXNwb25zZSLdAQoTR2V0T3ZlcndyaXRlUmVxdWVzdBIWCgxvdmVyd3JpdGVfaWQYASABKAlIABIvCglhY3RpdmVfYXQYAiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wSAASLgoKdGltZV9yYW5nZRgDIAEoCzIYLnRrZC5jb21tb24udjEuVGltZVJhbmdlSAASOQoPaW5ib3VuZF9udW1iZXJzGAQgASgLMiAudGtkLnBieDNjeC52MS5JbmJvdW5kTnVtYmVyTGlzdEISCghzZWxlY3RvchIG+vcYAggBIiQKEUluYm91bmROdW1iZXJMaXN0Eg8KB251bWJlcnMYASADKAkiRAoUR2V0T3ZlcndyaXRlUmVzcG9uc2USLAoKb3ZlcndyaXRlcxgBIAMoCzIYLnRrZC5wYngzY3gudjEuT3ZlcndyaXRlIoABChpDcmVhdGVJbmJvdW5kTnVtYmVyUmVxdWVzdBIXCgZudW1iZXIYASABKAlCB/r3GAPIAQESFAoMZGlzcGxheV9uYW1lGAIgASgJEhkKEXJvc3Rlcl9zaGlmdF90YWdzGAMgAygJEhgKEHJvc3Rlcl90eXBlX25hbWUYBCABKAkiUwobQ3JlYXRlSW5ib3VuZE51bWJlclJlc3BvbnNlEjQKDmluYm91bmRfbnVtYmVyGAEgASgLMhwudGtkLnBieDNjeC52MS5JbmJvdW5kTnVtYmVyIjUKGkRlbGV0ZUluYm91bmROdW1iZXJSZXF1ZXN0EhcKBm51bWJlchgBIAEoCUIH+vcYA8gBASIdChtEZWxldGVJbmJvdW5kTnVtYmVyUmVzcG9uc2UiGgoYTGlzdEluYm91bmROdW1iZXJSZXF1ZXN0IlIKGUxpc3RJbmJvdW5kTnVtYmVyUmVzcG9uc2USNQoPaW5ib3VuZF9udW1iZXJzGAEgAygLMhwudGtkLnBieDNjeC52MS5JbmJvdW5kTnVtYmVyIrUBChpVcGRhdGVJbmJvdW5kTnVtYmVyUmVxdWVzdBIXCgZudW1iZXIYASABKAlCB/r3GAPIAQESGAoQbmV3X2Rpc3BsYXlfbmFtZRgCIAEoCRIZChFyb3N0ZXJfc2hpZnRfdGFncxgDIAMoCRIYChByb3N0ZXJfdHlwZV9uYW1lGAQgASgJEi8KC3VwZGF0ZV9tYXNrGAogASgLMhouZ29vZ2xlLnByb3RvYnVmLkZpZWxkTWFzayJTChtVcGRhdGVJbmJvdW5kTnVtYmVyUmVzcG9uc2USNAoOaW5ib3VuZF9udW1iZXIYASABKAsyHC50a2QucGJ4M2N4LnYxLkluYm91bmROdW1iZXIynQoKC0NhbGxTZXJ2aWNlEk0KClJlY29yZENhbGwSIC50a2QucGJ4M2N4LnYxLlJlY29yZENhbGxSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5IgWyfgIIARJYCglHZXRPbkNhbGwSHy50a2QucGJ4M2N4LnYxLkdldE9uQ2FsbFJlcXVlc3QaIC50a2QucGJ4M2N4LnYxLkdldE9uQ2FsbFJlc3BvbnNlIgiQAgGyfgIIARJzChNDcmVhdGVJbmJvdW5kTnVtYmVyEikudGtkLnBieDNjeC52MS5DcmVhdGVJbmJvdW5kTnVtYmVyUmVxdWVzdBoqLnRrZC5wYngzY3gudjEuQ3JlYXRlSW5ib3VuZE51bWJlclJlc3BvbnNlIgWyfgIIAhJzChNVcGRhdGVJbmJvdW5kTnVtYmVyEikudGtkLnBieDNjeC52MS5VcGRhdGVJbmJvdW5kTnVtYmVyUmVxdWVzdBoqLnRrZC5wYngzY3gudjEuVXBkYXRlSW5ib3VuZE51bWJlclJlc3BvbnNlIgWyfgIIAhJzChNEZWxldGVJbmJvdW5kTnVtYmVyEikudGtkLnBieDNjeC52MS5EZWxldGVJbmJvdW5kTnVtYmVyUmVxdWVzdBoqLnRrZC5wYngzY3gudjEuRGVsZXRlSW5ib3VuZE51bWJlclJlc3BvbnNlIgWyfgIIAhJtChFMaXN0SW5ib3VuZE51bWJlchInLnRrZC5wYngzY3gudjEuTGlzdEluYm91bmROdW1iZXJSZXF1ZXN0GigudGtkLnBieDNjeC52MS5MaXN0SW5ib3VuZE51bWJlclJlc3BvbnNlIgWyfgIIAhJnCg9DcmVhdGVPdmVyd3JpdGUSJS50a2QucGJ4M2N4LnYxLkNyZWF0ZU92ZXJ3cml0ZVJlcXVlc3QaJi50a2QucGJ4M2N4LnYxLkNyZWF0ZU92ZXJ3cml0ZVJlc3BvbnNlIgWyfgIIARJnCg9EZWxldGVPdmVyd3JpdGUSJS50a2QucGJ4M2N4LnYxLkRlbGV0ZU92ZXJ3cml0ZVJlcXVlc3QaJi50a2QucGJ4M2N4LnYxLkRlbGV0ZU92ZXJ3cml0ZVJlc3BvbnNlIgWyfgIIARJeCgxHZXRPdmVyd3JpdGUSIi50a2QucGJ4M2N4LnYxLkdldE92ZXJ3cml0ZVJlcXVlc3QaIy50a2QucGJ4M2N4LnYxLkdldE92ZXJ3cml0ZVJlc3BvbnNlIgWyfgIIARJkCg5HZXRMb2dzRm9yRGF0ZRIkLnRrZC5wYngzY3gudjEuR2V0TG9nc0ZvckRhdGVSZXF1ZXN0GiUudGtkLnBieDNjeC52MS5HZXRMb2dzRm9yRGF0ZVJlc3BvbnNlIgWyfgIIARJkCg5TZWFyY2hDYWxsTG9ncxIkLnRrZC5wYngzY3gudjEuU2VhcmNoQ2FsbExvZ3NSZXF1ZXN0GiUudGtkLnBieDNjeC52MS5TZWFyY2hDYWxsTG9nc1Jlc3BvbnNlIgWyfgIIARJwChJHZXRMb2dzRm9yQ3VzdG9tZXISKC50a2QucGJ4M2N4LnYxLkdldExvZ3NGb3JDdXN0b21lclJlcXVlc3QaKS50a2QucGJ4M2N4LnYxLkdldExvZ3NGb3JDdXN0b21lclJlc3BvbnNlIgWyfgIIARonsn4SM2N4LmRvYmVyc2JlcmcudmV0un4PCg1pZG1fc3VwZXJ1c2VyQrwBChFjb20udGtkLnBieDNjeC52MUIMQ2FsbGxvZ1Byb3RvUAFaQ2dpdGh1Yi5jb20vdGllcmtsaW5pay1kb2JlcnNiZXJnL2FwaXMvZ2VuL2dvL3RrZC9wYngzY3gvdjE7cGJ4M2N4djGiAgNUUFiqAg1Ua2QuUGJ4M2N4LlYxygINVGtkXFBieDNjeFxWMeICGVRrZFxQYngzY3hcVjFcR1BCTWV0YWRhdGHqAg9Ua2Q6OlBieDNjeDo6VjFiBnByb3RvMw", [file_google_protobuf_descriptor, file_google_protobuf_timestamp, file_google_protobuf_duration, file_google_protobuf_field_mask, file_tkd_common_v1_descriptor, file_tkd_common_v1_time_range, file_tkd_customer_v1_customer, file_tkd_idm_v1_user, file_google_protobuf_empty, file_buf_validate_validate]);

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

