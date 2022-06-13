#!/usr/bin/env bash

# SETUP
GO_ALGORAND_SRC=$1
INDEXER_SRC=$2
CLI_TOOLS="~/go/bin/" # path to goal, algokey, etc.

# CLI GOAL
./reformat.py -doc-dir ../docs/clis/goal/ -cmd $CLI_TOOLS/goal

# CLI ALGOKEY
./reformat.py -doc-dir ../docs/clis/algokey/ -cmd $CLI_TOOLS/algokey

# CLI KMD
./reformat.py -doc-dir ../docs/clis/kmd/ -cmd $CLI_TOOLS/kmd

# CLI DIAGCFG
./reformat.py -doc-dir ../docs/clis/diagcfg/ -cmd $CLI_TOOLS/diagcfg

# CLI TEALDBG
./reformat.py -doc-dir ../docs/clis/tealdbg/ -cmd $CLI_TOOLS/tealdbg

# CLI INDEXER
./reformat.py -doc-dir ../docs/clis/indexer/ -cmd $CLI_TOOLS/algorand-indexer

# REST KMD
./convert_swagger.py -target ../docs/rest-apis/kmd.md -specfile $GO_ALGORAND_SRC/daemon/kmd/api/swagger.json

# REST ALGOD V1
#./convert_swagger.py -target ../docs/rest-apis/algod/v1.md -specfile $GO_ALGORAND_SRC/daemon/algod/api/swagger.json -processors algod_specfile_processors

# REST ALGOD V2 (from algod.oas2.json)
./convert_swagger.py -target ../docs/rest-apis/algod/v2.md -specfile $GO_ALGORAND_SRC/daemon/algod/api/algod.oas2.json

# REST INDEXER
./convert_swagger.py -target ../docs/rest-apis/indexer.md -specfile $INDEXER_SRC/api/indexer.oas2.json 

# TEAL
cp $GO_ALGORAND_SRC/data/transactions/logic/TEAL_opcodes.md ../docs/get-details/dapps/avm/teal/opcodes.md
cp $GO_ALGORAND_SRC/data/transactions/logic/README.md ../docs/get-details/dapps/avm/teal/specification.md
sed -i.bak '1s/#/title:/' ../docs/get-details/dapps/avm/teal/opcodes.md
sed -i.bak '1s/#/title:/' ../docs/get-details/dapps/avm/teal/specification.md
sed -i.bak 's/TEAL_opcodes.md/..\/opcodes/' ../docs/get-details/dapps/avm/teal/specification.md

# CLEANUP
rm swagger2markup-cli.jar
rm swagger2markup.properties
rm ../docs/get-details/dapps/avm/teal/*.bak
