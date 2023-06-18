"use strict";
// @generated by protoc-gen-connect-es v0.9.1 with parameter "target=js+ts"
// @generated from file tkd/idm/v1/user_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck
Object.defineProperty(exports, "__esModule", { value: true });
exports.UserService = void 0;
const user_service_pb_js_1 = require("./user_service_pb.js");
const protobuf_1 = require("@bufbuild/protobuf");
/**
 * @generated from service tkd.idm.v1.UserService
 */
exports.UserService = {
    typeName: "tkd.idm.v1.UserService",
    methods: {
        /**
         * @generated from rpc tkd.idm.v1.UserService.GetUser
         */
        getUser: {
            name: "GetUser",
            I: user_service_pb_js_1.GetUserRequest,
            O: user_service_pb_js_1.GetUserResponse,
            kind: protobuf_1.MethodKind.Unary,
            idempotency: protobuf_1.MethodIdempotency.NoSideEffects,
        },
        /**
         * @generated from rpc tkd.idm.v1.UserService.ListUsers
         */
        listUsers: {
            name: "ListUsers",
            I: user_service_pb_js_1.ListUsersRequest,
            O: user_service_pb_js_1.ListUsersResponse,
            kind: protobuf_1.MethodKind.Unary,
            idempotency: protobuf_1.MethodIdempotency.NoSideEffects,
        },
        /**
         * @generated from rpc tkd.idm.v1.UserService.CreateUser
         */
        createUser: {
            name: "CreateUser",
            I: user_service_pb_js_1.CreateUserRequest,
            O: user_service_pb_js_1.CreateUserResponse,
            kind: protobuf_1.MethodKind.Unary,
        },
        /**
         * @generated from rpc tkd.idm.v1.UserService.UpdateUser
         */
        updateUser: {
            name: "UpdateUser",
            I: user_service_pb_js_1.UpdateUserRequest,
            O: user_service_pb_js_1.UpdateUserResponse,
            kind: protobuf_1.MethodKind.Unary,
        },
        /**
         * @generated from rpc tkd.idm.v1.UserService.DeleteUser
         */
        deleteUser: {
            name: "DeleteUser",
            I: user_service_pb_js_1.DeleteUserRequest,
            O: user_service_pb_js_1.DeleteUserResponse,
            kind: protobuf_1.MethodKind.Unary,
        },
    }
};
