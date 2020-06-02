#!/usr/bin/env bash

# SETUP
GO_ALGORAND_SRC=$1
INDEXER_SRC=$2
mkdir -p ../docs/reference-docs/rest-apis/algod/v1
mkdir -p ../docs/reference-docs/rest-apis/algod/v2

# CLI GOAL
./reformat.py -doc-dir ../docs/reference-docs/goal/ -cmd ~/go/bin/goal

# CLI ALGOKEY
./reformat.py -doc-dir ../docs/reference-docs/algokey/ -cmd ~/go/bin/algokey

# CLI KMD
./reformat.py -doc-dir ../docs/reference-docs/kmd/ -cmd ~/go/bin/kmd

# REST KMD
./convert_swagger.py -target ../docs/reference-docs/rest-apis/kmd.md -specfile $GO_ALGORAND_SRC/daemon/kmd/api/swagger.json

# REST ALGOD V1
./convert_swagger.py -target ../docs/reference-docs/rest-apis/algod/v1/algod.md -specfile $GO_ALGORAND_SRC/daemon/algod/api/swagger.json -processors algod_specfile_processors

# REST ALGOD V2 (from algod.oas2.json)
./convert_swagger.py -target ../docs/reference-docs/rest-apis/algod/v2/algod.md -specfile $GO_ALGORAND_SRC/daemon/algod/api/algod.oas2.json

# REST INDEXER
./convert_swagger.py -target ../docs/reference-docs/rest-apis/indexer.md -specfile $INDEXER_SRC/api/indexer.oas2.json 

# CLEANUP
rm swagger2markup-cli.jar
rm swagger2markup.properties
