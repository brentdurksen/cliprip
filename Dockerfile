FROM golang:alpine AS build
WORKDIR /app
COPY . ./
RUN go mod download && go build -ldflags="-s -w" -o ./cliprip ./cmd/cliprip

FROM scratch
COPY --from=build /app/cliprip /cliprip
COPY --from=build /app/ui/ /ui/
CMD ["/cliprip"]
