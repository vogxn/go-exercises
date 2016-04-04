FROM golang

RUN mkdir -p /go/src/github.com/vogxn 
WORKDIR /go/src/github.com/vogxn

# release
RUN git clone https://github.com/vogxn/go-exercises.git && cd go-exercises/simple-web-service && git checkout v3
WORKDIR /go/src/github.com/vogxn/go-exercises/simple-web-service

# start the service
CMD ["go", "run", "simple.go"]

# open port 3222 where a net/http server will run
EXPOSE 3222

