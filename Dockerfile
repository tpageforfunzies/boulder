FROM iron/go:dev

RUN mkdir /app
WORKDIR /app

ENV SRC_DIR=/go/src/github.com/tpageforfunzies/boulder/

# Add the source code and env vars:
ADD . $SRC_DIR
ADD ./.env $SRC_DIR

# Build it, move binary and env:
RUN cd $SRC_DIR; go build -o boulder ./cmd/boulder; cp boulder /app/; cp .env /app/
RUN ["chmod", "+x", "./boulder"]
ENTRYPOINT ["./boulder"]