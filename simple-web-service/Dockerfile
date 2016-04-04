FROM golang

RUN mkdir -p /go/src/github.com/vogxn && cd /go/src/github.com/vogxn && git clone https://github.com/vogxn/go-exercises.git
WORKDIR /go/src/github.com/vogxn/go-exercises/simple-web-service

# start the service
CMD ["go", "run", "simple.go"]
RUN ps

# open port 3000 where a net/http server will run
EXPOSE 3000

