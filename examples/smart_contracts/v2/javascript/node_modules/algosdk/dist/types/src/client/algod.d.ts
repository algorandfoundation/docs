export function Algod(token?: string, baseServer?: string, port?: number, headers?: {}): void;
export class Algod {
    constructor(token?: string, baseServer?: string, port?: number, headers?: {});
    /**
     * status retrieves the StatusResponse from the running node
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    status: (headerObj?: {}) => Promise<any>;
    /**
     * healthCheck returns an empty object iff the node is running
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    healthCheck: (headerObj?: {}) => Promise<any>;
    /**
     * statusAfterBlock waits for round roundNumber to occur then returns the StatusResponse for this round.
     * This call blocks
     * @param roundNumber
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    statusAfterBlock: (roundNumber: any, headerObj?: {}) => Promise<any>;
    /**
     * pendingTransactions asks algod for a snapshot of current pending txns on the node, bounded by maxTxns.
     * If maxTxns = 0, fetches as many transactions as possible.
     * @param maxTxns - number
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    pendingTransactions: (maxTxns: any, headerObj?: {}) => Promise<any>;
    /**
     * versions retrieves the VersionResponse from the running node
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    versions: (headerObj?: {}) => Promise<any>;
    /**
     * LedgerSupply gets the supply details for the specified node's Ledger
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    ledgerSupply: (headerObj?: {}) => Promise<any>;
    /**
     * transactionsByAddress returns all transactions for a PK [addr] in the [first, last] rounds range.
     * @param addr - string
     * @param first - number, optional
     * @param last - number, optional
     * @param maxTxns - number, optional
     * @param headers, optional
     * @returns {Promise<*>}
     */
    transactionByAddress: (addr: any, first?: any, last?: any, maxTxns?: any, headerObj?: {}) => Promise<any>;
    /**
     * transactionsByAddressAndDate returns all transactions for a PK [addr] in the [fromDate, toDate] date range.
     * The date is a string in the YYYY-MM-DD format.
     * @param addr - string
     * @param fromDate - string
     * @param toDate - string
     * @param maxTxns - number, optional
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    transactionByAddressAndDate: (addr: any, fromDate: any, toDate: any, maxTxns?: any, headerObj?: {}) => Promise<any>;
    /**
     * transactionById returns the a transaction information of a specific txid [txId]
     * Note - This method is allowed only when Indexer is enabled.
     * @param txid
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    transactionById: (txid: any, headerObj?: {}) => Promise<any>;
    /**
     * transactionInformation returns the transaction information of a specific txid and an address
     * @param addr
     * @param txid
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    transactionInformation: (addr: any, txid: any, headerObj?: {}) => Promise<any>;
    /**
     * pendingTransactionInformation returns the transaction information for a specific txid of a pending transaction
     * @param txid
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    pendingTransactionInformation: (txid: any, headerObj?: {}) => Promise<any>;
    /**
     * accountInformation returns the passed account's information
     * @param addr - string
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    accountInformation: (addr: any, headerObj?: {}) => Promise<any>;
    /**
     * assetInformation returns the information for the asset with the passed creator and index
     * @param index - number
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    assetInformation: (index: any, headerObj?: {}) => Promise<any>;
    /**
     * suggestedFee gets the recommended transaction fee from the node
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    suggestedFee: (headerObj?: {}) => Promise<any>;
    /**
     * sendRawTransaction gets an encoded SignedTxn and broadcasts it to the network
     * @param txn - Uin8Array
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    sendRawTransaction: (txn: any, headerObj?: {}) => Promise<any>;
    /**
     * sendRawTransactions gets a list of encoded SignedTxns and broadcasts it to the network
     * @param txn - Array of Uin8Array
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    sendRawTransactions: (txns: any, headerObj?: {}) => Promise<any>;
    /**
     * getTransactionParams returns to common needed parameters for a new transaction
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    getTransactionParams: (headerObj?: {}) => Promise<any>;
    /**
     * suggestParams returns to common needed parameters for a new transaction, in a format the transaction builder expects
     * @param headerObj, optional
     * @returns {Object}
     */
    suggestParams: (headerObj?: {}) => any;
    /**
     * block gets the block info for the given round This call blocks
     * @param roundNumber
     * @param headerObj, optional
     * @returns {Promise<*>}
     */
    block: (roundNumber: any, headerObj?: {}) => Promise<any>;
}
