FROM golang

RUN mkdir -p /go/src/github.com/vogxn 
RUN cd /go/src/github.com/vogxn && git clone https://github.com/vogxn/go-exercises.git 
WORKDIR /go/src/github.com/vogxn/go-exercises/simple-web-service

# release
RUN git checkout v3

# start the service
CMD ["go", "run", "simple.go"]

# open port 3222 where a net/http server will run
EXPOSE 3222

