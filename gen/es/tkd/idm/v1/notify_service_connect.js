// @generated by protoc-gen-connect-es v0.11.0 with parameter "target=js+dts"
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
export const NotifyService = {
  typeName: "tkd.idm.v1.NotifyService",
  methods: {
    /**
     * @generated from rpc tkd.idm.v1.NotifyService.SendNotification
     */
    sendNotification: {
      name: "SendNotification",
      I: SendNotificationRequest,
      O: SendNotificationResponse,
      kind: MethodKind.Unary,
    },
  }
};

