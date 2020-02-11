title: Renew Participation Keys

The process of renewing a participation key is simply creating a new participation key and registering it online before the previous key expires.

You can renew a participation key anytime before it expires, and we recommend to do it at least two weeks (about 38,400 rounds) in advance so as not to risk [having an account marked as online that is not participating](./overview.md#ensure-that-online-accounts-are-participating). 

The validity ranges of participation keys can overlap. For any account, at any time, at most one participation key is registered, namely the one included in the latest online key registration transaction for this account. 


# Step-by-Step
1. [Create a new participation key](./generate_keys.md) with a first voting round that is less than the last voting round of the current participation key. It should leave enough time to carry out this whole process (e.g. 40,000 rounds).
2. Place both keys in the participating node's ledger directory.
3. Once the network reaches the first voting round for the new key, [submit an online key registration transaction for the new participation key](./online.md).
4. Wait at least 320 rounds to [validate that the node is participating](./online.md#check-that-the-node-is-participating).
5. Once participation is confirmed, it is safe to delete the old participation key.

<center> ![Renewal Window](../../imgs/renew-0.png) </center>
<center>*Example key rotation window* </center>
