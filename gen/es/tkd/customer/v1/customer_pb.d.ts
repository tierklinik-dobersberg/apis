// @generated by protoc-gen-es v2.0.0 with parameter "target=js+dts"
// @generated from file tkd/customer/v1/customer.proto (package tkd.customer.v1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import type { JsonObject, Message } from "@bufbuild/protobuf";
import type { Timestamp } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file tkd/customer/v1/customer.proto.
 */
export declare const file_tkd_customer_v1_customer: GenFile;

/**
 * @generated from message tkd.customer.v1.CustomerRef
 */
export declare type CustomerRef = Message<"tkd.customer.v1.CustomerRef"> & {
  /**
   * @generated from field: string source = 1;
   */
  source: string;

  /**
   * @generated from field: string id = 2;
   */
  id: string;
};

/**
 * Describes the message tkd.customer.v1.CustomerRef.
 * Use `create(CustomerRefSchema)` to create a new message.
 */
export declare const CustomerRefSchema: GenMessage<CustomerRef>;

/**
 * @generated from message tkd.customer.v1.OwnedAttribute
 */
export declare type OwnedAttribute = Message<"tkd.customer.v1.OwnedAttribute"> & {
  /**
   * @generated from oneof tkd.customer.v1.OwnedAttribute.kind
   */
  kind: {
    /**
     * @generated from field: string first_name = 1;
     */
    value: string;
    case: "firstName";
  } | {
    /**
     * @generated from field: string last_name = 2;
     */
    value: string;
    case: "lastName";
  } | {
    /**
     * @generated from field: string phone_number = 3;
     */
    value: string;
    case: "phoneNumber";
  } | {
    /**
     * @generated from field: string email_address = 4;
     */
    value: string;
    case: "emailAddress";
  } | {
    /**
     * @generated from field: tkd.customer.v1.Address address = 5;
     */
    value: Address;
    case: "address";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message tkd.customer.v1.OwnedAttribute.
 * Use `create(OwnedAttributeSchema)` to create a new message.
 */
export declare const OwnedAttributeSchema: GenMessage<OwnedAttribute>;

/**
 * @generated from message tkd.customer.v1.ImportState
 */
export declare type ImportState = Message<"tkd.customer.v1.ImportState"> & {
  /**
   * Importer is a unique name of the importer that created or updated the
   * customer or patient record.
   *
   * @generated from field: string importer = 1;
   */
  importer: string;

  /**
   * LastSeen holds the timestamp the importer has seen the customer record the
   * last time.
   *
   * @generated from field: google.protobuf.Timestamp last_seen = 2;
   */
  lastSeen?: Timestamp;

  /**
   * ExtraData may hold additional data for the importer.
   *
   * @generated from field: google.protobuf.Struct extra_data = 3;
   */
  extraData?: JsonObject;

  /**
   * InternalReference may hold an internal customer reference from
   * the importer.
   *
   * @generated from field: string internal_reference = 4;
   */
  internalReference: string;

  /**
   * @generated from field: repeated tkd.customer.v1.OwnedAttribute owned_attributes = 5;
   */
  ownedAttributes: OwnedAttribute[];
};

/**
 * Describes the message tkd.customer.v1.ImportState.
 * Use `create(ImportStateSchema)` to create a new message.
 */
export declare const ImportStateSchema: GenMessage<ImportState>;

/**
 * @generated from message tkd.customer.v1.Address
 */
export declare type Address = Message<"tkd.customer.v1.Address"> & {
  /**
   * @generated from field: string postal_code = 1;
   */
  postalCode: string;

  /**
   * @generated from field: string city = 2;
   */
  city: string;

  /**
   * @generated from field: string street = 3;
   */
  street: string;

  /**
   * @generated from field: string extra = 4;
   */
  extra: string;
};

/**
 * Describes the message tkd.customer.v1.Address.
 * Use `create(AddressSchema)` to create a new message.
 */
export declare const AddressSchema: GenMessage<Address>;

/**
 * @generated from message tkd.customer.v1.Customer
 */
export declare type Customer = Message<"tkd.customer.v1.Customer"> & {
  /**
   * ID is a unique, randomly generated ID for this customer.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * FirstName is the first name of the customer.
   *
   * @generated from field: string first_name = 2;
   */
  firstName: string;

  /**
   * LastName is the last name of the customer
   *
   * @generated from field: string last_name = 3;
   */
  lastName: string;

  /**
   * Addresses is a list of addresses.
   *
   * @generated from field: repeated tkd.customer.v1.Address addresses = 4;
   */
  addresses: Address[];

  /**
   * PhoneNumbers is a list of phone numbers
   *
   * @generated from field: repeated string phone_numbers = 5;
   */
  phoneNumbers: string[];

  /**
   * EmailAddresses is a list of email addresses.
   *
   * @generated from field: repeated string email_addresses = 6;
   */
  emailAddresses: string[];

  /**
   * RecordCreatedAt is the timestamp at which the customer record has been created.
   *
   * @generated from field: google.protobuf.Timestamp record_created_at = 7;
   */
  recordCreatedAt?: Timestamp;
};

/**
 * Describes the message tkd.customer.v1.Customer.
 * Use `create(CustomerSchema)` to create a new message.
 */
export declare const CustomerSchema: GenMessage<Customer>;

/**
 * @generated from message tkd.customer.v1.CustomerResponse
 */
export declare type CustomerResponse = Message<"tkd.customer.v1.CustomerResponse"> & {
  /**
   * @generated from field: tkd.customer.v1.Customer customer = 1;
   */
  customer?: Customer;

  /**
   * @generated from field: repeated tkd.customer.v1.ImportState states = 2;
   */
  states: ImportState[];
};

/**
 * Describes the message tkd.customer.v1.CustomerResponse.
 * Use `create(CustomerResponseSchema)` to create a new message.
 */
export declare const CustomerResponseSchema: GenMessage<CustomerResponse>;

/**
 * @generated from message tkd.customer.v1.InternalReferenceQuery
 */
export declare type InternalReferenceQuery = Message<"tkd.customer.v1.InternalReferenceQuery"> & {
  /**
   * @generated from field: string importer = 1;
   */
  importer: string;

  /**
   * @generated from field: string ref = 2;
   */
  ref: string;
};

/**
 * Describes the message tkd.customer.v1.InternalReferenceQuery.
 * Use `create(InternalReferenceQuerySchema)` to create a new message.
 */
export declare const InternalReferenceQuerySchema: GenMessage<InternalReferenceQuery>;

/**
 * @generated from message tkd.customer.v1.NameQuery
 */
export declare type NameQuery = Message<"tkd.customer.v1.NameQuery"> & {
  /**
   * @generated from field: string first_name = 1;
   */
  firstName: string;

  /**
   * @generated from field: string last_name = 2;
   */
  lastName: string;
};

/**
 * Describes the message tkd.customer.v1.NameQuery.
 * Use `create(NameQuerySchema)` to create a new message.
 */
export declare const NameQuerySchema: GenMessage<NameQuery>;

/**
 * @generated from message tkd.customer.v1.CustomerQuery
 */
export declare type CustomerQuery = Message<"tkd.customer.v1.CustomerQuery"> & {
  /**
   * @generated from oneof tkd.customer.v1.CustomerQuery.query
   */
  query: {
    /**
     * @generated from field: string id = 1;
     */
    value: string;
    case: "id";
  } | {
    /**
     * @generated from field: tkd.customer.v1.InternalReferenceQuery internal_reference = 2;
     */
    value: InternalReferenceQuery;
    case: "internalReference";
  } | {
    /**
     * @generated from field: tkd.customer.v1.NameQuery name = 3;
     */
    value: NameQuery;
    case: "name";
  } | {
    /**
     * @generated from field: string phone_number = 4;
     */
    value: string;
    case: "phoneNumber";
  } | {
    /**
     * @generated from field: string email_address = 5;
     */
    value: string;
    case: "emailAddress";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message tkd.customer.v1.CustomerQuery.
 * Use `create(CustomerQuerySchema)` to create a new message.
 */
export declare const CustomerQuerySchema: GenMessage<CustomerQuery>;

/**
 * @generated from message tkd.customer.v1.SearchCustomerRequest
 */
export declare type SearchCustomerRequest = Message<"tkd.customer.v1.SearchCustomerRequest"> & {
  /**
   * @generated from field: repeated tkd.customer.v1.CustomerQuery queries = 1;
   */
  queries: CustomerQuery[];
};

/**
 * Describes the message tkd.customer.v1.SearchCustomerRequest.
 * Use `create(SearchCustomerRequestSchema)` to create a new message.
 */
export declare const SearchCustomerRequestSchema: GenMessage<SearchCustomerRequest>;

/**
 * @generated from message tkd.customer.v1.SearchCustomerResponse
 */
export declare type SearchCustomerResponse = Message<"tkd.customer.v1.SearchCustomerResponse"> & {
  /**
   * @generated from field: repeated tkd.customer.v1.CustomerResponse results = 1;
   */
  results: CustomerResponse[];
};

/**
 * Describes the message tkd.customer.v1.SearchCustomerResponse.
 * Use `create(SearchCustomerResponseSchema)` to create a new message.
 */
export declare const SearchCustomerResponseSchema: GenMessage<SearchCustomerResponse>;

/**
 * @generated from message tkd.customer.v1.UpdateCustomerRequest
 */
export declare type UpdateCustomerRequest = Message<"tkd.customer.v1.UpdateCustomerRequest"> & {
  /**
   * @generated from field: tkd.customer.v1.Customer customer = 1;
   */
  customer?: Customer;
};

/**
 * Describes the message tkd.customer.v1.UpdateCustomerRequest.
 * Use `create(UpdateCustomerRequestSchema)` to create a new message.
 */
export declare const UpdateCustomerRequestSchema: GenMessage<UpdateCustomerRequest>;

/**
 * @generated from message tkd.customer.v1.UpdateCustomerResponse
 */
export declare type UpdateCustomerResponse = Message<"tkd.customer.v1.UpdateCustomerResponse"> & {
  /**
   * @generated from field: tkd.customer.v1.CustomerResponse response = 1;
   */
  response?: CustomerResponse;
};

/**
 * Describes the message tkd.customer.v1.UpdateCustomerResponse.
 * Use `create(UpdateCustomerResponseSchema)` to create a new message.
 */
export declare const UpdateCustomerResponseSchema: GenMessage<UpdateCustomerResponse>;

/**
 * @generated from service tkd.customer.v1.CustomerService
 */
export declare const CustomerService: GenService<{
  /**
   * @generated from rpc tkd.customer.v1.CustomerService.SearchCustomer
   */
  searchCustomer: {
    methodKind: "unary";
    input: typeof SearchCustomerRequestSchema;
    output: typeof SearchCustomerResponseSchema;
  },
  /**
   * @generated from rpc tkd.customer.v1.CustomerService.SearchCustomerStream
   */
  searchCustomerStream: {
    methodKind: "bidi_streaming";
    input: typeof SearchCustomerRequestSchema;
    output: typeof SearchCustomerResponseSchema;
  },
  /**
   * @generated from rpc tkd.customer.v1.CustomerService.UpdateCustomer
   */
  updateCustomer: {
    methodKind: "unary";
    input: typeof UpdateCustomerRequestSchema;
    output: typeof UpdateCustomerResponseSchema;
  },
}>;

