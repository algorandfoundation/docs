#!/usr/bin/env bash
rm -rf ../docs/Reference-Docs/goal && mkdir ../docs/Reference-Docs/goal && goal generate-docs ../docs/Reference-Docs/goal && ./reformat.py -path ../docs/Reference-Docs/goal -cmd goal
