import * as request from 'superagent';
/**
 * Implementation of BaseHTTPClient that uses a URL and a token
 * and make the REST queries using superagent.
 * This is the default implementation of BaseHTTPClient.
 */
export class URLTokenBaseHTTPClient {
    constructor(tokenHeader, baseServer, port, defaultHeaders = {}) {
        this.defaultHeaders = defaultHeaders;
        // Append a trailing slash so we can use relative paths. Without the trailing
        // slash, the last path segment will be replaced by the relative path. See
        // usage in `addressWithPath`.
        const fixedBaseServer = baseServer.endsWith('/')
            ? baseServer
            : `${baseServer}/`;
        const baseServerURL = new URL(fixedBaseServer);
        if (typeof port !== 'undefined') {
            baseServerURL.port = port.toString();
        }
        if (baseServerURL.protocol.length === 0) {
            throw new Error('Invalid base server URL, protocol must be defined.');
        }
        this.baseURL = baseServerURL;
        this.tokenHeader = tokenHeader;
    }
    /**
     * Compute the URL for a path relative to the instance's address
     * @param relativePath - A path string
     * @returns A URL string
     */
    addressWithPath(relativePath) {
        let fixedRelativePath;
        if (relativePath.startsWith('./')) {
            fixedRelativePath = relativePath;
        }
        else if (relativePath.startsWith('/')) {
            fixedRelativePath = `.${relativePath}`;
        }
        else {
            fixedRelativePath = `./${relativePath}`;
        }
        const address = new URL(fixedRelativePath, this.baseURL);
        return address.toString();
    }
    /**
     * Convert a superagent response to a valid BaseHTTPClientResponse
     * Modify the superagent response
     * @private
     */
    static superagentToHTTPClientResponse(res) {
        if (res.body instanceof ArrayBuffer) {
            // Handle the case where the body is an arraybuffer which happens in the browser
            res.body = new Uint8Array(res.body);
        }
        return res;
    }
    /**
     * Make a superagent error more readable. For more info, see https://github.com/visionmedia/superagent/issues/1074
     */
    static formatSuperagentError(err) {
        if (err.response) {
            try {
                const decoded = JSON.parse(Buffer.from(err.response.body).toString());
                // eslint-disable-next-line no-param-reassign
                err.message = `Network request error. Received status ${err.response.status}: ${decoded.message}`;
            }
            catch (err2) {
                // ignore any error that happened while we are formatting the original error
            }
        }
        return err;
    }
    async get(relativePath, query, requestHeaders = {}) {
        const r = request
            .get(this.addressWithPath(relativePath))
            .set(this.tokenHeader)
            .set(this.defaultHeaders)
            .set(requestHeaders)
            .responseType('arraybuffer')
            .query(query);
        try {
            const res = await r;
            return URLTokenBaseHTTPClient.superagentToHTTPClientResponse(res);
        }
        catch (err) {
            throw URLTokenBaseHTTPClient.formatSuperagentError(err);
        }
    }
    async post(relativePath, data, query, requestHeaders = {}) {
        const r = request
            .post(this.addressWithPath(relativePath))
            .set(this.tokenHeader)
            .set(this.defaultHeaders)
            .set(requestHeaders)
            .query(query)
            .serialize((o) => o) // disable serialization from superagent
            .responseType('arraybuffer')
            .send(Buffer.from(data)); // Buffer.from necessary for superagent
        try {
            const res = await r;
            return URLTokenBaseHTTPClient.superagentToHTTPClientResponse(res);
        }
        catch (err) {
            throw URLTokenBaseHTTPClient.formatSuperagentError(err);
        }
    }
    async delete(relativePath, data, query, requestHeaders = {}) {
        const r = request
            .delete(this.addressWithPath(relativePath))
            .set(this.tokenHeader)
            .set(this.defaultHeaders)
            .set(requestHeaders)
            .query(query)
            .serialize((o) => o) // disable serialization from superagent
            .responseType('arraybuffer')
            .send(Buffer.from(data)); // Buffer.from necessary for superagent
        try {
            const res = await r;
            return URLTokenBaseHTTPClient.superagentToHTTPClientResponse(res);
        }
        catch (err) {
            throw URLTokenBaseHTTPClient.formatSuperagentError(err);
        }
    }
}
//# sourceMappingURL=urlTokenBaseHTTPClient.js.map