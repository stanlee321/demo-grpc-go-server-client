FROM golang:latest 

LABEL Name=profile-service Version=0.0.1 

RUN apt-get update && apt-get install -y ca-certificates git-core ssh

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH 

# now copy your app to the proper build path
RUN mkdir -p $GOPATH/src/app 
ADD . $GOPATH/src/app

WORKDIR $GOPATH/src/app
RUN go build -o main . 
CMD ["/go/src/app/main"]

EXPOSE 50051