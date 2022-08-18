# Собираем в гошке
FROM golang:1.17.3 as build

ENV BIN_FILE /opt/system_monitoring_daemon/daemon
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/daemon/*

# На выходе тонкий образ
FROM alpine:3.15

LABEL ORGANIZATION="Cyber inc."
LABEL SERVICE="system monitoring daemon"
LABEL MAINTAINERS="cybertmt@gmail.com"

ENV BIN_FILE "/opt/system_monitoring_daemon/daemon"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/system_monitoring_daemon/config.yaml
COPY ./configs/config.yaml ${CONFIG_FILE}

CMD ${BIN_FILE} -config ${CONFIG_FILE}
