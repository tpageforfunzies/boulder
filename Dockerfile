FROM iron/go:dev

RUN mkdir /app
WORKDIR /app
# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/tpageforfunzies/boulder/
# Add the source code:
ADD . $SRC_DIR
ADD ./.env $SRC_DIR
# Build it:
RUN cd $SRC_DIR; go build -o boulder ./cmd/boulder; cp boulder /app/; cp .env /app/
RUN ["chmod", "+x", "./boulder"]
ENTRYPOINT ["./boulder"]