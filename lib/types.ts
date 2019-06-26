/**
 * A `Uint8Array` of binary data
 */
export type Bytes = Uint8Array;

/**
 * A `Uint8Array` of Amino disambiguation data (3 bytes)
 */
export type DisambBytes = AminoBytes;

/**
 * A `Uint8Array` of Amino prefix data (4 bytes)
 */
export type PrefixBytes = AminoBytes;

/**
 * A `Uint8Array` of JSON data
 */
export type JSONBytes = Bytes;

/**
 * A `Uint8Array` of Amino data
 */
export type AminoBytes = Bytes;

/**
 * Base64-encoded string
 */
export type Base64String = string;

/**
 * Bech32-encoded string
 */
export type Bech32String = string;

/**
 * Binary-encoded string
 */
export type BinaryString = string;

/**
 * Hex-encoded string
 */
export type HexString = string;

/**
 * Unicode Scalar Values-encoded string
 */
export type USVString = string;

/**
 * @TODO document
 */
export type Int8 = number;

/**
 * @TODO document
 */
export type Int16 = number;

/**
 * @TODO document
 */
export type Int32 = string;

/**
 * @TODO document
 */
export type Int64 = string;

/**
 * @TODO document
 */
export type Varint = string;

/**
 * @TODO document
 */
export type Uint8 = number;

/**
 * @TODO document
 */
export type Uint16 = number;

/**
 * @TODO document
 */
export type Uint32 = string;

/**
 * @TODO document
 */
export type Uint64 = string;

/**
 * @TODO document
 */
export type Uvarint = string;

/**
 * @TODO document
 */
export type Byte = Uint8;
