import { BaseHTTPClient, BaseHTTPClientResponse, Query } from './baseHTTPClient';
export interface AlgodTokenHeader {
    'X-Algo-API-Token': string;
}
export interface IndexerTokenHeader {
    'X-Indexer-API-Token': string;
}
export interface KMDTokenHeader {
    'X-KMD-API-Token': string;
}
export interface CustomTokenHeader {
    [headerName: string]: string;
}
export declare type TokenHeader = AlgodTokenHeader | IndexerTokenHeader | KMDTokenHeader | CustomTokenHeader;
/**
 * Implementation of BaseHTTPClient that uses a URL and a token
 * and make the REST queries using superagent.
 * This is the default implementation of BaseHTTPClient.
 */
export declare class URLTokenBaseHTTPClient implements BaseHTTPClient {
    private defaultHeaders;
    private readonly baseURL;
    private readonly tokenHeader;
    constructor(tokenHeader: TokenHeader, baseServer: string, port?: string | number, defaultHeaders?: Record<string, any>);
    /**
     * Compute the URL for a path relative to the instance's address
     * @param relativePath - A path string
     * @returns A URL string
     */
    private addressWithPath;
    /**
     * Convert a superagent response to a valid BaseHTTPClientResponse
     * Modify the superagent response
     * @private
     */
    private static superagentToHTTPClientResponse;
    /**
     * Make a superagent error more readable. For more info, see https://github.com/visionmedia/superagent/issues/1074
     */
    private static formatSuperagentError;
    get(relativePath: string, query?: Query<string>, requestHeaders?: Record<string, string>): Promise<BaseHTTPClientResponse>;
    post(relativePath: string, data: Uint8Array, query?: Query<string>, requestHeaders?: Record<string, string>): Promise<BaseHTTPClientResponse>;
    delete(relativePath: string, data: Uint8Array, query?: Query<string>, requestHeaders?: Record<string, string>): Promise<BaseHTTPClientResponse>;
}
