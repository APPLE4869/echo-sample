FROM golang:1.13.0-alpine

# Set workdir
WORKDIR /go/src/app/
COPY . ./

ENV GO111MODULE=on

# Update packages && install necessary ones.
RUN apk update && apk add --no-cache git

# Install "realize" for live reload.
RUN GO111MODULE=off go get -u github.com/oxequa/realize

CMD ["realize", "start"]
