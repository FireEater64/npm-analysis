#! /bin/sh

apk add jq curl

curl -L https://replicate.npmjs.com/_all_docs | jq -c [.rows[]] > /npm.json