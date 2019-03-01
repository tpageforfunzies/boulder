setup:
	dep ensure

basic:
	go build -o app ./cmd/boulder
	./app