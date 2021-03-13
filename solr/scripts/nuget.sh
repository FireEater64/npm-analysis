#! /bin/sh

apk add jq curl parallel

URLS=$(curl https://api.nuget.org/v3/catalog0/index.json | jq '.items[] | .["@id"]' | tr -d '"')

parallel --eta -j 100 curl -s --compressed {} ::: $URLS | jq '.items[] | .["nuget:id"]' | tr -d '"' | sort | uniq > raw.txt

jq -R -n -c '[inputs | split(",") | {id: .[0]}]' raw.txt > nuget.json