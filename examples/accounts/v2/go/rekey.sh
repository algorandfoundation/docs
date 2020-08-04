#!/bin/bash

# origianl code supplied by @algobolson
# https://github.com/justicz/go-algorand/blob/maxj/applications/test/scripts/e2e_subs/rekey.sh

set -e
set -x
set -o pipefail
export SHELLOPTS

# ensure you pass the wallet name as arg1 when executing the script
WALLET=$1

# set the path for your temporary directory used to store account keys
TEMPDIR="/TEMPDIR"

# Set the appropriate algod values for your environment. These are the defaults for Sandbox on TestNet
ALGOD_HOST_PORT="localhost:4001"
ALGOD_HEADER="x-algo-api-token:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

# ensure your PATH is set properly, else modify the following and set /path/to/goal and /path/to/algokey
GOAL_CMD="goal -d $DATA -w ${WALLET}"
ALGOKEY_CMD="algokey"

# load *Default account from your wallet as ACCOUNT_MASTER. It must be funded and will dispense funds to newly accounts for testing. MASTER_ACCOUNT will NOT be rekeyed.
ACCOUNT_MASTER=$(${GOAL_CMD} account list | awk '$6 == "*Default" { print $3 }')

# create a new account as ACCOUNT_A using goal
ACCOUNT_A=$(${GOAL_CMD} account new | awk '{ print $6 }')

echo "Fund ACCOUNT_A from ACCOUNT_MASTER"
${GOAL_CMD} clerk send --from ${ACCOUNT_MASTER} --to ${ACCOUNT_A} --amount 1000000

# generate a new account as ACCOUNT_B using algokey
${ALGOKEY_CMD} generate > ${TEMPDIR}/account_b
MNEMONIC_ACCOUNT_B=$(grep 'Private key mnemonic:' < ${TEMPDIR}/account_b | sed 's/Private key mnemonic: //')
ACCOUNT_B=$(grep 'Public key:' < ${TEMPDIR}/account_b | sed 's/Public key: //')

# import ACCOUNT_B into your wallet using the mnemonic passphrase
${GOAL_CMD} account import -m "${MNEMONIC_ACCOUNT_B}"

# create a new account as ACCOUNT_C using goal
ACCOUNT_C=$(${GOAL_CMD} account new | awk '{ print $6 }')

echo "Initial account object for ACCOUNT_A:\n"
curl "$ALGOD_HOST_PORT/v2/accounts/$ACCOUNT_A" -H $ALGOD_HEADER

# TEST_1: send from ACCOUNT_A to ACCOUNT_B
echo "TEST_1:\n"
${GOAL_CMD} clerk send --from ${ACCOUNT_A} --to ${ACCOUNT_B} --amount 100000
# EXPECTED RESULT: Pass, the transaction will confirm.

# REKEY_TO_1: ACCOUNT_A rekey-to ACCOUNT_B
echo "Rekeying ACCOUNT_A to ACCOUNT_B\n"
${GOAL_CMD} clerk send --from ${ACCOUNT_A} --to ${ACCOUNT_B} --amount 0 --rekey-to ${ACCOUNT_B}
echo "Updated account object for ACCOUNT_A:\n"
curl "$ALGOD_HOST_PORT/v2/accounts/$ACCOUNT_A" -H $ALGOD_HEADER

# TEST_2: send from ACCOUNT_A to ACCOUNT_B signed by ACCOUNT_A (default behavior)
echo "TEST_2:\n"
# create, sign and send transaction
${GOAL_CMD} clerk send --from ${ACCOUNT_A} --to ${ACCOUNT_B} --amount 100000
# EXPECTED RESULT: Fail, the transaction will not confirm because should have been authorized by ACCOUNT_B

# TEST_3: send from ACCOUNT_A to ACCOUNT_B signed by ACCOUNT_B
echo "TEST_3:\n"
# create unsigned transaction
${GOAL_CMD} clerk send --from ${ACCOUNT_A} --to ${ACCOUNT_B} --amount 100000 --out ${TEMPDIR}/test3.utxn
# sign transaction
${GOAL_CMD} clerk sign --signer ${ACCOUNT_B} --infile ${TEMPDIR}/test3.utxn --outfile ${TEMPDIR}/test3.stxn
# send signed transaction
${GOAL_CMD} clerk rawsend -f ${TEMPDIR}/test3.stxn
# EXPECTED RESULT: Pass, the transaction will confirm because ACCOUNT_B is the authorized address.

# Create a multisig account MSIG_BC_T_1 with threshold 1
MSIG_BC_T_1=$(${GOAL_CMD} account multisig new $ACCOUNT_B $ACCOUNT_C --threshold 1 | awk '{ print $6 }')

# rekey ACCOUNT_A to MSIG_BC_T_1
echo "Rekeying ACCOUNT_A to MSIG_BC_T_1\n"
${GOAL_CMD} clerk send --from ${ACCOUNT_A} --to ${ACCOUNT_B} --amount 0 --rekey-to ${MSIG_BC_T_1} --out ${TEMPDIR}/rekey-to-msig_bc.utxn
${GOAL_CMD} clerk sign --signer ${ACCOUNT_B} --infile ${TEMPDIR}/rekey-to-msig_bc.utxn --outfile ${TEMPDIR}/rekey-to-msig_bc.stxn
${GOAL_CMD} clerk rawsend -f ${TEMPDIR}/rekey-to-msig_bc.stxn
echo "Updated account object for ACCOUNT_A:\n"
curl "$ALGOD_HOST_PORT/v2/accounts/$ACCOUNT_A" -H $ALGOD_HEADER

# TEST_4: send from ACCOUNT_A to ACCOUNT_B signed by ACCOUNT_B
echo "TEST_4:\n"
${GOAL_CMD} clerk send --from ${ACCOUNT_A} --to ${ACCOUNT_B} --amount 100000 --out ${TEMPDIR}/test4.utxn
${GOAL_CMD} clerk multisig sign --address ${ACCOUNT_B} --tx ${TEMPDIR}/test4.utxn
${GOAL_CMD} clerk rawsend -f ${TEMPDIR}/test4.utxn
# EXPECTED RESULT: Pass, the transaction will confirm because ACCOUNT_B is suficiant threshold of MSIG_BC_T_1.


# change key on every transaction


