FROM golang:1.18

RUN mkdir /sbercloud_test
WORKDIR /sbercloud_test

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY app/main.go app/
COPY pkg/models/*.go pkg/models/
COPY pkg/handler/*.go pkg/handler/
COPY pkg/repositories/*.go pkg/repositories/
COPY pkg/server/*.go pkg/server/
COPY pkg/services/*.go pkg/services/

RUN go build -o sbercloud_test app/main.go

EXPOSE 3228

CMD [ "./sbercloud_test" ]