FROM golang:1.16-alpine as builder

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary


FROM scratch as prod

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 8080
CMD ["/go/src/app/binary"]

# ---------for development--------------

FROM golang:1.16-alpine as dev

WORKDIR /app

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go get -u github.com/cosmtrek/air
EXPOSE 8080

CMD ["air"]
