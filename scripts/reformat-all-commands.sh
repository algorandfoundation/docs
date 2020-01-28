#!/usr/bin/env bash
./reformat.py -doc-dir ../docs/reference-docs/goal/ -cmd ~/go/bin/goal
./reformat.py -doc-dir ../docs/reference-docs/algokey/ -cmd ~/go/bin/algokey
./reformat.py -doc-dir ../docs/reference-docs/kmd/ -cmd ~/go/bin/kmd
./convert_swagger.py -target ../docs/reference-docs/rest-apis/algod.md -specfile ../swagger-algod.json -processors algod_specfile_processors
./convert_swagger.py -target ../docs/reference-docs/rest-apis/kmd.md -specfile ../swagger-kmd.json
