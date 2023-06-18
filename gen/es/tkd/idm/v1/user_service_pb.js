"use strict";
// @generated by protoc-gen-es v1.2.1 with parameter "target=js+ts"
// @generated from file tkd/idm/v1/user_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck
Object.defineProperty(exports, "__esModule", { value: true });
exports.DeleteUserResponse = exports.DeleteUserRequest = exports.UpdateUserResponse = exports.UpdateUserRequest = exports.CreateUserResponse = exports.CreateUserRequest = exports.ListUsersResponse = exports.ListUsersRequest = exports.GetUserResponse = exports.GetUserRequest = void 0;
const protobuf_1 = require("@bufbuild/protobuf");
const user_pb_js_1 = require("./user_pb.js");
/**
 * @generated from message tkd.idm.v1.GetUserRequest
 */
class GetUserRequest extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: string id = 1;
         */
        this.id = "";
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new GetUserRequest().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new GetUserRequest().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new GetUserRequest().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(GetUserRequest, a, b);
    }
}
exports.GetUserRequest = GetUserRequest;
GetUserRequest.runtime = protobuf_1.proto3;
GetUserRequest.typeName = "tkd.idm.v1.GetUserRequest";
GetUserRequest.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
]);
/**
 * @generated from message tkd.idm.v1.GetUserResponse
 */
class GetUserResponse extends protobuf_1.Message {
    constructor(data) {
        super();
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new GetUserResponse().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new GetUserResponse().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new GetUserResponse().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(GetUserResponse, a, b);
    }
}
exports.GetUserResponse = GetUserResponse;
GetUserResponse.runtime = protobuf_1.proto3;
GetUserResponse.typeName = "tkd.idm.v1.GetUserResponse";
GetUserResponse.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "user", kind: "message", T: user_pb_js_1.User },
]);
/**
 * @generated from message tkd.idm.v1.ListUsersRequest
 */
class ListUsersRequest extends protobuf_1.Message {
    constructor(data) {
        super();
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new ListUsersRequest().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new ListUsersRequest().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new ListUsersRequest().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(ListUsersRequest, a, b);
    }
}
exports.ListUsersRequest = ListUsersRequest;
ListUsersRequest.runtime = protobuf_1.proto3;
ListUsersRequest.typeName = "tkd.idm.v1.ListUsersRequest";
ListUsersRequest.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "field_mask", kind: "message", T: protobuf_1.FieldMask },
]);
/**
 * @generated from message tkd.idm.v1.ListUsersResponse
 */
class ListUsersResponse extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: repeated tkd.idm.v1.Profile users = 1;
         */
        this.users = [];
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new ListUsersResponse().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new ListUsersResponse().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new ListUsersResponse().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(ListUsersResponse, a, b);
    }
}
exports.ListUsersResponse = ListUsersResponse;
ListUsersResponse.runtime = protobuf_1.proto3;
ListUsersResponse.typeName = "tkd.idm.v1.ListUsersResponse";
ListUsersResponse.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "users", kind: "message", T: user_pb_js_1.Profile, repeated: true },
]);
/**
 * @generated from message tkd.idm.v1.CreateUserRequest
 */
class CreateUserRequest extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: string password = 2;
         */
        this.password = "";
        /**
         * @generated from field: bool password_is_bcrypt = 3;
         */
        this.passwordIsBcrypt = false;
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new CreateUserRequest().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new CreateUserRequest().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new CreateUserRequest().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(CreateUserRequest, a, b);
    }
}
exports.CreateUserRequest = CreateUserRequest;
CreateUserRequest.runtime = protobuf_1.proto3;
CreateUserRequest.typeName = "tkd.idm.v1.CreateUserRequest";
CreateUserRequest.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "profile", kind: "message", T: user_pb_js_1.Profile },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "password_is_bcrypt", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
]);
/**
 * @generated from message tkd.idm.v1.CreateUserResponse
 */
class CreateUserResponse extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: string generated_password = 2;
         */
        this.generatedPassword = "";
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new CreateUserResponse().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new CreateUserResponse().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new CreateUserResponse().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(CreateUserResponse, a, b);
    }
}
exports.CreateUserResponse = CreateUserResponse;
CreateUserResponse.runtime = protobuf_1.proto3;
CreateUserResponse.typeName = "tkd.idm.v1.CreateUserResponse";
CreateUserResponse.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "profile", kind: "message", T: user_pb_js_1.Profile },
    { no: 2, name: "generated_password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
]);
/**
 * @generated from message tkd.idm.v1.UpdateUserRequest
 */
class UpdateUserRequest extends protobuf_1.Message {
    constructor(data) {
        super();
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new UpdateUserRequest().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new UpdateUserRequest().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new UpdateUserRequest().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(UpdateUserRequest, a, b);
    }
}
exports.UpdateUserRequest = UpdateUserRequest;
UpdateUserRequest.runtime = protobuf_1.proto3;
UpdateUserRequest.typeName = "tkd.idm.v1.UpdateUserRequest";
UpdateUserRequest.fields = protobuf_1.proto3.util.newFieldList(() => []);
/**
 * @generated from message tkd.idm.v1.UpdateUserResponse
 */
class UpdateUserResponse extends protobuf_1.Message {
    constructor(data) {
        super();
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new UpdateUserResponse().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new UpdateUserResponse().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new UpdateUserResponse().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(UpdateUserResponse, a, b);
    }
}
exports.UpdateUserResponse = UpdateUserResponse;
UpdateUserResponse.runtime = protobuf_1.proto3;
UpdateUserResponse.typeName = "tkd.idm.v1.UpdateUserResponse";
UpdateUserResponse.fields = protobuf_1.proto3.util.newFieldList(() => []);
/**
 * @generated from message tkd.idm.v1.DeleteUserRequest
 */
class DeleteUserRequest extends protobuf_1.Message {
    constructor(data) {
        super();
        /**
         * @generated from field: string id = 1;
         */
        this.id = "";
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new DeleteUserRequest().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new DeleteUserRequest().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new DeleteUserRequest().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(DeleteUserRequest, a, b);
    }
}
exports.DeleteUserRequest = DeleteUserRequest;
DeleteUserRequest.runtime = protobuf_1.proto3;
DeleteUserRequest.typeName = "tkd.idm.v1.DeleteUserRequest";
DeleteUserRequest.fields = protobuf_1.proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
]);
/**
 * @generated from message tkd.idm.v1.DeleteUserResponse
 */
class DeleteUserResponse extends protobuf_1.Message {
    constructor(data) {
        super();
        protobuf_1.proto3.util.initPartial(data, this);
    }
    static fromBinary(bytes, options) {
        return new DeleteUserResponse().fromBinary(bytes, options);
    }
    static fromJson(jsonValue, options) {
        return new DeleteUserResponse().fromJson(jsonValue, options);
    }
    static fromJsonString(jsonString, options) {
        return new DeleteUserResponse().fromJsonString(jsonString, options);
    }
    static equals(a, b) {
        return protobuf_1.proto3.util.equals(DeleteUserResponse, a, b);
    }
}
exports.DeleteUserResponse = DeleteUserResponse;
DeleteUserResponse.runtime = protobuf_1.proto3;
DeleteUserResponse.typeName = "tkd.idm.v1.DeleteUserResponse";
DeleteUserResponse.fields = protobuf_1.proto3.util.newFieldList(() => []);
