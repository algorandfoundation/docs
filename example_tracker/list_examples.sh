#!/bin/bash

echo "Looking for example names in docs"
grep -r 'JSSDK_' ../docs | sort | uniq  |awk -F= '{print $4}'  | cut -f 2- -d '_'  > js_examples.txt
grep -r 'PYSDK_' ../docs | sort | uniq  |awk -F= '{print $4}'  | cut -f 2- -d '_' > py_examples.txt
grep -r 'JAVASDK_' ../docs | sort | uniq  |awk -F= '{print $4}'  | cut -f 2- -d '_' > java_examples.txt
grep -r 'GOSDK_' ../docs | sort | uniq  |awk -F= '{print $4}'  | cut -f 2- -d '_' > go_examples.txt


echo "Looking at counts of each"
cat ./*.txt | sort | uniq -c | sort -n -r
