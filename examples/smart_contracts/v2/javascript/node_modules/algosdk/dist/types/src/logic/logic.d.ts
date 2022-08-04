/**
 * Utilities for working with program bytes.
 */
export declare function parseUvarint(array: Uint8Array): [numberFound: number, size: number];
/** readProgram validates program for length and running cost,
 * and additionally provides the found int variables and byte blocks
 * @param program - Program to check
 * @param args - Program arguments as array of Uint8Array arrays
 * @throws
 * @returns
 */
export declare function readProgram(program: Uint8Array, args?: Uint8Array[]): [ints: number[], byteArrays: Uint8Array[], valid: boolean];
/**
 * checkProgram validates program for length and running cost
 * @param program - Program to check
 * @param args - Program arguments as array of Uint8Array arrays
 * @throws
 * @returns true if success
 */
export declare function checkProgram(program: Uint8Array, args?: Uint8Array[]): boolean;
export declare function checkIntConstBlock(program: Uint8Array, pc: number): number;
export declare function checkByteConstBlock(program: Uint8Array, pc: number): number;
export declare function checkPushIntOp(program: Uint8Array, pc: number): number;
export declare function checkPushByteOp(program: Uint8Array, pc: number): number;
export declare const langspecEvalMaxVersion: number;
export declare const langspecLogicSigVersion: number;
