build-MonolithFunction:
	GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o ./bootstrap ./cmd/main.go
	cp ./bootstrap $(ARTIFACTS_DIR)/.
