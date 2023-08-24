// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=js+dts"
// @generated from file tkd/idm/v1/notify_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { SendNotificationRequest, SendNotificationResponse } from "./notify_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * Admin API only
 *
 *
 * @generated from service tkd.idm.v1.NotifyService
 */
export declare const NotifyService: {
  readonly typeName: "tkd.idm.v1.NotifyService",
  readonly methods: {
    /**
     * @generated from rpc tkd.idm.v1.NotifyService.SendNotification
     */
    readonly sendNotification: {
      readonly name: "SendNotification",
      readonly I: typeof SendNotificationRequest,
      readonly O: typeof SendNotificationResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

