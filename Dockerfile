# build binary
FROM golang:1.12.6 AS build

RUN go get github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/saharshsingh/gophinoff
ADD . .

RUN dep ensure && \
        cd taskmaster && \
        GOOS=linux CGO_ENABLED=0 go build

# create runnable image
FROM scratch

LABEL maintainer="Saharsh Singh"

COPY --from=build /go/src/github.com/saharshsingh/gophinoff/taskmaster/taskmaster /

EXPOSE 8080
ENTRYPOINT [ "/taskmaster" ]
