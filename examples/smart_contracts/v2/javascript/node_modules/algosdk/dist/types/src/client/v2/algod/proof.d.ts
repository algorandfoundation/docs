import JSONRequest from '../jsonrequest';
import HTTPClient from '../../client';
import IntDecoding from '../../../types/intDecoding';
export default class Proof extends JSONRequest {
    private round;
    private txID;
    constructor(c: HTTPClient, intDecoding: IntDecoding, round: number, txID: string);
    path(): string;
}
