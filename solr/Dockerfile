FROM solr:latest

# Download latest npm index and convert to importable json
RUN curl -L https://replicate.npmjs.com/_all_docs | sed '1c[' | sed '$c]' > /tmp/npm.json

# Create core and import npm packages
RUN precreate-core packages && start-local-solr && post -c packages /tmp/npm.json && stop-local-solr