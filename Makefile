setup:
ifeq (,$(wildcard ./.env))
	cp ./.env-local ./.env
else
	$(info Env already exists.  Not overwriting.)
endif
	dep ensure

test:
	# run all tests verbose
	go test -cover ./...

build:
	# macosx compile
	go build -o app ./cmd/boulder
	# linux compile
	env GOOS=linux GOARCH=arm go build -o app.linux ./cmd/boulder

basic:
	go build -o app ./cmd/boulder
	env GOOS=linux GOARCH=arm go build -o app.linux ./cmd/boulder
	sudo ./app