format:
	gofmt -w .

send_example_request:
	curl -s -X POST -d @scripts/example_webhook_request.json 0.0.0.0:8080/slack_tz_webhook | jq