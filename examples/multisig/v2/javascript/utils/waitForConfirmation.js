// Function used to wait for a tx confirmation
exports.waitForConfirmation = async function (algodClient, txId) {
    const response = await algodClient.status().do();
    let lastround = response["last-round"];
    while (true) {
        const pendingInfo = await algodClient.pendingTransactionInformation(txId).do();
        if (pendingInfo["confirmed-round"] !== null && pendingInfo["confirmed-round"] > 0) {
            //Got the completed Transaction
            console.log("Transaction " + txId + " confirmed in round " + pendingInfo["confirmed-round"]);
            return pendingInfo;
            break;
        }
        lastround++;
        await algodClient.statusAfterBlock(lastround).do();
    }
};