build:
	CGO_ENABLED=0 go mod download && go build -ldflags="-s -w" -o ./dist/cliprip ./cmd/cliprip
