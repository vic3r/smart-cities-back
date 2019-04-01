FROM golang:alpine AS build-env

ENV GOBIN /go/bin
ARG GITHUB_TOKEN
ENV GITHUB_TOKEN ${GITHUB_TOKEN}
RUN apk add git

ADD . /go/src/github.com/vic3r/smart-cities-back/
WORKDIR /go/src/github.com/vic3r/smart-cities-back/

RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/" \
    && wget -q -O - https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure -vendor-only

WORKDIR /go/src/github.com/vic3r/smart-cities-back
RUN CGO_ENABLED=1 go build -i ./cmd/mibici

FROM alpine

ARG ENVIRONMENT	
ENV ENVIRONMENT ${ENVIRONMENT}

EXPOSE 8080

WORKDIR /app
COPY --from=build-env /go/src/github.com/vic3r/smart-cities-back /app/
COPY config/* /app/config/

ENTRYPOINT ./mibici -ConfigDir=config -ConfigType=yaml
