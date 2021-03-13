#! /bin/sh

apk add python3 jq

# NOT the way to do this
python3 -c "import xmlrpc.client; packages = xmlrpc.client.ServerProxy('https://pypi.python.org/pypi').list_packages(); print(\"\n\".join(packages))" | jq -R -n -c '[inputs | split(",") | {id: .[0]}]' > pypi.json