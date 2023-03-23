#!/bin/bash

echo "Looking for SDK example names in docs"
grep -r 'JSSDK_' ../../docs    |awk -F= '{print $4}'  |cut -f 2- -d '_' |sort |uniq > js_examples.txt
grep -r 'PYSDK_' ../../docs    |awk -F= '{print $4}'  |cut -f 2- -d '_' |sort |uniq > py_examples.txt
grep -r 'JAVASDK_' ../../docs  |awk -F= '{print $4}'  |cut -f 2- -d '_' |sort |uniq > java_examples.txt
grep -r 'GOSDK_' ../../docs    |awk -F= '{print $4}'  |cut -f 2- -d '_' |sort |uniq > go_examples.txt


echo "Looking at counts of each (they should all be 4, one for each sdk)"
cat ./*.txt | sort | uniq -c | sort -n -r 
