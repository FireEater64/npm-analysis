FROM alpine AS nuget

COPY scripts/nuget.sh /nuget.sh
RUN chmod +x /nuget.sh && ./nuget.sh

FROM alpine AS npm

COPY scripts/npm.sh /npm.sh
RUN chmod +x /npm.sh && ./npm.sh

FROM alpine AS pypi

COPY scripts/pypi.sh /pypi.sh
RUN chmod +x /pypi.sh && ./pypi.sh

FROM solr:latest

# Download latest npm index and convert to importable json
COPY --from=nuget /nuget.json /tmp/nuget.json
COPY --from=npm /npm.json /tmp/npm.json
COPY --from=pypi /pypi.json /tmp/pypi.json

RUN precreate-core npm && \
    precreate-core nuget && \
    precreate-core pypi && \
    start-local-solr && \
    post -c npm /tmp/npm.json && \
    post -c nuget /tmp/nuget.json && \
    post -c pypi /tmp/pypi.json && \
    stop-local-solr