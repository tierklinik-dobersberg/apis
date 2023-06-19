// @generated by protoc-gen-connect-es v0.9.1 with parameter "target=js+ts"
// @generated from file tkd/idm/v1/auth_service.proto (package tkd.idm.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { IntrospectRequest, IntrospectResponse, LoginRequest, LoginResponse, LogoutRequest, LogoutResponse, RefreshTokenRequest, RefreshTokenResponse } from "./auth_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * AuthService provides methods for authentication to the
 * services hosted and managed by Tierklinik Dobersberg (TKD).
 *
 * @generated from service tkd.idm.v1.AuthService
 */
export const AuthService = {
  typeName: "tkd.idm.v1.AuthService",
  methods: {
    /**
     * Login requests authentication. The authentication type (flow) is
     * determined by the initial request and may require sub-sequent calls
     * to full-fill the requirements of the choosen authentication flow.
     *
     * Upon success, a LoginResponse with a AccessTokenResponse is returned to
     * the caller containing a short lived access token (typically about ~24h).
     * In addition, a "Set-Cookie" header is appended to the response that contains
     * a HttpOnly, Secure (if not in dev-mode) cookie with a long-lived refresh token
     * (~ about a month).
     *
     * In case the access token expires the client is expected to call the RefreshToken
     * endpoint to retrieve a new access token.
     *
     * Refresh tokens cannot be re-newed like this but require a full re-authentication
     * using the Login method.
     *
     * @generated from rpc tkd.idm.v1.AuthService.Login
     */
    login: {
      name: "Login",
      I: LoginRequest,
      O: LoginResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Logout invalidates the current access token that was used to call the 
     * method. If a refresh token is appended in the logout request, it will be invalidated
     * as well.
     *
     * @generated from rpc tkd.idm.v1.AuthService.Logout
     */
    logout: {
      name: "Logout",
      I: LogoutRequest,
      O: LogoutResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.AuthService.RefreshToken
     */
    refreshToken: {
      name: "RefreshToken",
      I: RefreshTokenRequest,
      O: RefreshTokenResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc tkd.idm.v1.AuthService.Introspect
     */
    introspect: {
      name: "Introspect",
      I: IntrospectRequest,
      O: IntrospectResponse,
      kind: MethodKind.Unary,
    },
  }
};

