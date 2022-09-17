FROM golang:1.18 AS build
#Unset the GOPATH
ENV GOPATH=

COPY . . 
RUN go mod download

RUN CGO_ENABLED=0 go build -o /go/test-app

# Use empty image for running the app
FROM scratch
COPY --from=build /go/test-app /bin/test-app

ENTRYPOINT [ "./bin/test-app" ]

