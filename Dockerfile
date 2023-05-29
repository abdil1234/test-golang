FROM golang:alpine

# Add packages to set timezone to WIB instead of default UTC value
RUN apk update && \
    apk add build-base && \
    apk add --no-cache tzdata && \
    apk add --no-cache openssh && \
    apk add --no-cache git && \
    apk add curl && \
    apk add bash

RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && echo "Asia/Jakarta" >  /etc/timezone

ENV GOARM=7 \
    GOOS=linux \
    GOARCH=amd64\
    GOFLAGS=-buildvcs=false

ENV WAIT_VERSION 2.7.2
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/$WAIT_VERSION/wait /wait
RUN chmod +x /wait

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

# Set working directory for docker container to /app
WORKDIR /app

COPY . .
# Create directory for storing application logs

# Config .netrc
RUN go env -w GO111MODULE=on
RUN go get -v

# Run the app
CMD ["air"]