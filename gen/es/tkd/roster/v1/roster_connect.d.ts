// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=js+dts"
// @generated from file tkd/roster/v1/roster.proto (package tkd.roster.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AnalyzeWorkTimeRequest, AnalyzeWorkTimeResponse, ApproveRosterRequest, ApproveRosterResponse, CreateRosterTypeRequest, CreateRosterTypeResponse, DeleteRosterRequest, DeleteRosterResponse, DeleteRosterTypeRequest, DeleteRosterTypeResponse, GetRequiredShiftsRequest, GetRequiredShiftsResponse, GetRosterRequest, GetRosterResponse, GetWorkingStaffRequest, GetWorkingStaffResponse, ListRosterTypesRequest, ListRosterTypesResponse, ListShiftTagsRequest, ListShiftTagsResponse, SaveRosterRequest, SaveRosterResponse, SendRosterPreviewRequest, SendRosterPreviewResponse } from "./roster_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service tkd.roster.v1.RosterService
 */
export declare const RosterService: {
  readonly typeName: "tkd.roster.v1.RosterService",
  readonly methods: {
    /**
     * @generated from rpc tkd.roster.v1.RosterService.CreateRosterType
     */
    readonly createRosterType: {
      readonly name: "CreateRosterType",
      readonly I: typeof CreateRosterTypeRequest,
      readonly O: typeof CreateRosterTypeResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.RosterService.DeleteRosterType
     */
    readonly deleteRosterType: {
      readonly name: "DeleteRosterType",
      readonly I: typeof DeleteRosterTypeRequest,
      readonly O: typeof DeleteRosterTypeResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.RosterService.ListRosterTypes
     */
    readonly listRosterTypes: {
      readonly name: "ListRosterTypes",
      readonly I: typeof ListRosterTypesRequest,
      readonly O: typeof ListRosterTypesResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.RosterService.ListShiftTags
     */
    readonly listShiftTags: {
      readonly name: "ListShiftTags",
      readonly I: typeof ListShiftTagsRequest,
      readonly O: typeof ListShiftTagsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * SaveRoster saves a duty roster. It may be used to initially create a new
     * roster or to save subsequent changes.
     *
     * @generated from rpc tkd.roster.v1.RosterService.SaveRoster
     */
    readonly saveRoster: {
      readonly name: "SaveRoster",
      readonly I: typeof SaveRosterRequest,
      readonly O: typeof SaveRosterResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * DeleteRoster deletes a roster from the internal storage. This operation
     * cannot be undone!
     *
     * @generated from rpc tkd.roster.v1.RosterService.DeleteRoster
     */
    readonly deleteRoster: {
      readonly name: "DeleteRoster",
      readonly I: typeof DeleteRosterRequest,
      readonly O: typeof DeleteRosterResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * AnalyzeWorkTime can be used to analyze the work time of users to determine
     * undertime or overtime in a given time-range.
     *
     * @generated from rpc tkd.roster.v1.RosterService.AnalyzeWorkTime
     */
    readonly analyzeWorkTime: {
      readonly name: "AnalyzeWorkTime",
      readonly I: typeof AnalyzeWorkTimeRequest,
      readonly O: typeof AnalyzeWorkTimeResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * ApproveRoster marks a roster as approved by management.
     *
     * @generated from rpc tkd.roster.v1.RosterService.ApproveRoster
     */
    readonly approveRoster: {
      readonly name: "ApproveRoster",
      readonly I: typeof ApproveRosterRequest,
      readonly O: typeof ApproveRosterResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * GetRoster returns a previously saved roster.
     *
     * @generated from rpc tkd.roster.v1.RosterService.GetRoster
     */
    readonly getRoster: {
      readonly name: "GetRoster",
      readonly I: typeof GetRosterRequest,
      readonly O: typeof GetRosterResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * GetWorkingStaff returns a list of user_ids that are working at the
     * date specified in GetWorkingStaffRequest. If date is unset, it defaults
     * to NOW.
     *
     * @generated from rpc tkd.roster.v1.RosterService.GetWorkingStaff
     */
    readonly getWorkingStaff: {
      readonly name: "GetWorkingStaff",
      readonly I: typeof GetWorkingStaffRequest,
      readonly O: typeof GetWorkingStaffResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.RosterService.GetRequiredShifts
     */
    readonly getRequiredShifts: {
      readonly name: "GetRequiredShifts",
      readonly I: typeof GetRequiredShiftsRequest,
      readonly O: typeof GetRequiredShiftsResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.roster.v1.RosterService.SendRosterPreview
     */
    readonly sendRosterPreview: {
      readonly name: "SendRosterPreview",
      readonly I: typeof SendRosterPreviewRequest,
      readonly O: typeof SendRosterPreviewResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

