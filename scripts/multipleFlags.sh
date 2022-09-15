#!/bin/bash

readarray -t name <<< $(jq -r '.[].name' $1)
type=( $(jq -r '.[].type' $1) )
readarray -t description <<< $(jq -r '.[].description' $1)
source=( $(jq -r '.[].source' $1) )

for (( i=0; i<${#name[@]}; i++))
do 
    flagship flag create -d $"{\"name\":\"${name[$i]}\",\"type\":\"${type[$i]}\", \"description\":\"${description[$i]}\", \"source\":\"${source[$i]}\"}"
done
