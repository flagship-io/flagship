#!/bin/bash

name=( $(jq -r '.[].name' $1) )
type=( $(jq -r '.[].type' $1) )
readarray -t description <<< $(jq -r '.[].description' $1)


for (( i=0; i<${#name[@]}; i++))
do 
    eval $(echo flagship tk create -d \'{\"name\":\"${name[$i]}\",\"type\":\"${type[$i]}\", \"description\":\"${description[$i]}\"}\')
done
