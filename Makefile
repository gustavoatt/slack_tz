lint:
	gofmt -w .
	goimports -w .
	golangci-lint run -E golint -E gofmt -E godot -E goimports

# Runs the server and uses ngrok to forward requests to it from the internet.
tunnel:
	go run cmd/slack_tz/main.go
	ngrok http 8080

send_example_request:
	curl -s -X POST -d @scripts/example_webhook_request.txt 0.0.0.0:8080/slack_tz_webhook
