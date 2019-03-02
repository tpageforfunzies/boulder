setup:
ifeq (,$(wildcard ./.env))
	cp ./.env-local ./.env
else
	$(info Env already exists.  Not overwriting.)
endif
	dep ensure

basic:
	go build -o app ./cmd/boulder
	./app