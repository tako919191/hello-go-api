# ----------------------------------------- #
# Build Stage
# ----------------------------------------- #
FROM golang:1.17-alpine as builder

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY . ${ROOT}
RUN CGO_ENABLED=0 GOOS=linux go build -o $ROOT/binary

# ----------------------------------------- #
# Prod Stage
# ----------------------------------------- #
FROM scratch

ENV ROOT=/go/src/app
WORKDIR ${ROOT}
COPY --from=builder ${ROOT}/binary ${ROOT}

EXPOSE 3000
CMD ["/go/src/app/binary"]