test-integration:
	go test -v ./tests/integration

test-integration-html: test-cover
	go tool cover -html=./tests/coverage/coverage.out -o ./tests/coverage/coverage.html
	open ./tests/coverage/coverage.html  # (Mac) Automatically opens the report
	# xdg-open coverage.html  # (Linux)
	# start coverage.html  # (Windows)