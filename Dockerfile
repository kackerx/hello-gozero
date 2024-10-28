FROM golang:1.23 as builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp rpc/add/add.go

From scratch

COPY --from=builder /app/myapp /myapp

ENTRYPOINT ["/myapp"]

CMD ["-f", "/add.yaml"]
# d run -it \
# --name add \
# -v /Users/apple/GolandProjects/hellozero/rpc/add/etc/add.yaml:/add.yaml \
# kacker/book-add