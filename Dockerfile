FROM golang:alpine AS build
WORKDIR /go/src/myapp
COPY . .
RUN go build -o main.go

FROM scratch
COPY --from=build myapp myapp
ENTRYPOINT ["/go/bin/myapp"]