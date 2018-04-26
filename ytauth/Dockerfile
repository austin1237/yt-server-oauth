FROM golang:1.10.1-alpine
RUN apk update && apk upgrade

RUN apk add curl
RUN apk add git


# Instal go dep 
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

# Add Project Source Code
RUN mkdir -p /go/src/github.com/user/ytauth
COPY . /go/src/github.com/user/ytauth
WORKDIR /go/src/github.com/user/ytauth

# Install go dependencies
RUN dep ensure 

RUN go install .

CMD ytauth