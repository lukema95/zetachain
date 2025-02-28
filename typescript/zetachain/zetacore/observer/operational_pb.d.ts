// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file zetachain/zetacore/observer/operational.proto (package zetachain.zetacore.observer, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, Duration, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * Flags for the top-level operation of zetaclient.
 *
 * @generated from message zetachain.zetacore.observer.OperationalFlags
 */
export declare class OperationalFlags extends Message<OperationalFlags> {
  /**
   * Height for a coordinated zetaclient restart.
   * Will be ignored if missed.
   *
   * @generated from field: int64 restart_height = 1;
   */
  restartHeight: bigint;

  /**
   * Offset from the zetacore block time to initiate signing.
   * Should be calculated and set based on max(zetaclient_core_block_latency).
   *
   * @generated from field: google.protobuf.Duration signer_block_time_offset = 2;
   */
  signerBlockTimeOffset?: Duration;

  /**
   * Minimum version of zetaclient that is allowed to run. This must be either
   * a valid semver string (v23.0.1) or empty. If empty, all versions are
   * allowed.
   *
   * @generated from field: string minimum_version = 3;
   */
  minimumVersion: string;

  constructor(data?: PartialMessage<OperationalFlags>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.OperationalFlags";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OperationalFlags;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OperationalFlags;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OperationalFlags;

  static equals(a: OperationalFlags | PlainMessage<OperationalFlags> | undefined, b: OperationalFlags | PlainMessage<OperationalFlags> | undefined): boolean;
}

