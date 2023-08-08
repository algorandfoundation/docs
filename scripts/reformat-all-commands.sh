#!/usr/bin/env bash

# SETUP
GO_ALGORAND_SRC=$1
INDEXER_SRC=$2
CLI_TOOLS="~/go/bin/" # path to goal, algorand-indexer, algokey, etc.

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

# CLI CONDUIT
#./reformat.py -doc-dir ../docs/clis/conduit/ -cmd $CLI_TOOLS/conduit

# REST KMD
./convert_swagger.py -target ../docs/rest-apis/kmd.md -specfile $GO_ALGORAND_SRC/daemon/kmd/api/swagger.json

# REST ALGOD V2 (from algod.oas2.json)
./convert_swagger.py -target ../docs/rest-apis/algod.md -specfile $GO_ALGORAND_SRC/daemon/algod/api/algod.oas2.json

# REST INDEXER
./convert_swagger.py -target ../docs/rest-apis/indexer.md -specfile $INDEXER_SRC/api/indexer.oas2.json 

# TEAL
cp $GO_ALGORAND_SRC/data/transactions/logic/jsonspec.md ../docs/get-details/dapps/avm/teal/jsonspec.md
cp $GO_ALGORAND_SRC/data/transactions/logic/README.md ../docs/get-details/dapps/avm/teal/specification.md
sed -i.bak '1s/#/title:/' ../docs/get-details/dapps/avm/teal/specification.md
sed -i.bak 's/TEAL_opcodes.md/..\/opcodes/' ../docs/get-details/dapps/avm/teal/specification.md

# TEAL - OPCODES DIRECTORY
rm -rf ../docs/get-details/dapps/avm/teal/opcodes
mkdir ../docs/get-details/dapps/avm/teal/opcodes

pages_file="../docs/get-details/dapps/avm/teal/opcodes/.pages"
cat <<EOF > $pages_file
title: Opcodes
arrange:
  - index.md
EOF

index_file="../docs/get-details/dapps/avm/teal/opcodes/index.md"
cat <<EOF > $index_file
title: Opcodes By Version

This page lists the TEAL opcodes by version. For more information on TEAL, see the [TEAL specification](../specification/).

Opcodes by version:

EOF

for file in ${GO_ALGORAND_SRC}/data/transactions/logic/TEAL_opcodes_v*.md; do
    if [ -f "$file" ]; then
        filename=$(basename "$file")  # Extract the filename without path
        version=${filename#TEAL_opcodes_v}  # Remove the prefix
        version=${version%.md}  # Remove the suffix

        cp "$file" "../docs/get-details/dapps/avm/teal/opcodes/v${version}.md"
        sed -i.bak '1s/#/title:/' "../docs/get-details/dapps/avm/teal/opcodes/v${version}.md"
        echo "  - v${version}.md" >> $pages_file
        echo "- [v${version}](v${version}.md)" >> $index_file
    fi
done

# CLEANUP
rm swagger2markup-cli.jar
rm swagger2markup.properties
rm ../docs/get-details/dapps/avm/teal/*.bak
rm ../docs/get-details/dapps/avm/teal/opcodes/*.bak
