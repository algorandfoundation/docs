"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function () { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function () { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getAccounts = void 0;
var algosdk_1 = require("algosdk");
var kmd_token = "a".repeat(64);
var kmd_host = "http://localhost";
var kmd_port = "4002";
var kmd_wallet = "unencrypted-default-wallet";
var kmd_password = "";
function getAccounts() {
    return __awaiter(this, void 0, Promise, function () {
        var kmdClient, wallets, walletId, _i, _a, wallet, handleResp, handle, addresses, acctPromises, _b, _c, addr, keys;
        return __generator(this, function (_d) {
            switch (_d.label) {
                case 0:
                    kmdClient = new algosdk_1.default.Kmd(kmd_token, kmd_host, kmd_port);
                    return [4 /*yield*/, kmdClient.listWallets()];
                case 1:
                    wallets = _d.sent();
                    for (_i = 0, _a = wallets['wallets']; _i < _a.length; _i++) {
                        wallet = _a[_i];
                        if (wallet['name'] === kmd_wallet)
                            walletId = wallet['id'];
                    }
                    if (walletId === undefined)
                        throw Error("No wallet named: " + kmd_wallet);
                    return [4 /*yield*/, kmdClient.initWalletHandle(walletId, kmd_password)];
                case 2:
                    handleResp = _d.sent();
                    handle = handleResp['wallet_handle_token'];
                    return [4 /*yield*/, kmdClient.listKeys(handle)];
                case 3:
                    addresses = _d.sent();
                    acctPromises = [];
                    for (_b = 0, _c = addresses['addresses']; _b < _c.length; _b++) {
                        addr = _c[_b];
                        acctPromises.push(kmdClient.exportKey(handle, kmd_password, addr));
                    }
                    return [4 /*yield*/, Promise.all(acctPromises)
                        // Don't need to wait for it
                    ];
                case 4:
                    keys = _d.sent();
                    // Don't need to wait for it
                    kmdClient.releaseWalletHandle(handle);
                    return [2 /*return*/, keys.map(function (k) {
                        var addr = algosdk_1.default.encodeAddress(k.private_key.slice(32));
                        var acct = { sk: k.private_key, addr: addr };
                        return acct;
                    })];
            }
        });
    });
}
exports.getAccounts = getAccounts;
