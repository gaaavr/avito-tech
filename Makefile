test:
	go test ./...

run:
	# Export .env
	set -o allexport
	. "./.env"
	set +o allexport
	go run cmd/main.go