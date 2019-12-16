FROM golang:alpine
WORKDIR /
RUN apk update && apk add make git
COPY . /
RUN make bench # Run whatever you want here
CMD ./btree