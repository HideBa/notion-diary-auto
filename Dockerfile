FROM golang:1.16-alpine as build
ARG TAG=production

WORKDIR /app

RUN apk update --no-cache \
  && apk add --no-cache \
    git \
    gcc \
    musl-dev

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN GOOS=linux go build -o app main.go

FROM scratch as prod

WORKDIR /app

# RUN apk update --no-cache \
#   && apk add --no-cache ca-certificates
# RUN update-ca-certificates

COPY --from=build /app/app .

CMD ["./app"]

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
