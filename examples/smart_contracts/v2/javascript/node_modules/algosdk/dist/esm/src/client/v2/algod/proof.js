import JSONRequest from '../jsonrequest';
export default class Proof extends JSONRequest {
    constructor(c, intDecoding, round, txID) {
        super(c, intDecoding);
        this.round = round;
        this.txID = txID;
        this.round = round;
        this.txID = txID;
    }
    path() {
        return `/v2/blocks/${this.round}/transactions/${this.txID}/proof`;
    }
}
//# sourceMappingURL=proof.js.map