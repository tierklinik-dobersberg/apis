// @generated by protoc-gen-es v1.4.1 with parameter "target=js+dts"
// @generated from file tkd/customer/v1/customer.proto (package tkd.customer.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message tkd.customer.v1.CustomerRef
 */
export declare class CustomerRef extends Message<CustomerRef> {
  /**
   * @generated from field: string source = 1;
   */
  source: string;

  /**
   * @generated from field: string id = 2;
   */
  id: string;

  constructor(data?: PartialMessage<CustomerRef>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "tkd.customer.v1.CustomerRef";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CustomerRef;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CustomerRef;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CustomerRef;

  static equals(a: CustomerRef | PlainMessage<CustomerRef> | undefined, b: CustomerRef | PlainMessage<CustomerRef> | undefined): boolean;
}

