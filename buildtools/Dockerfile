# buildtools golang env
FROM registry-office.zzishare.com/ssp/robotbase:latest AS build-env

ENV GO111MODULE=on
COPY . /game-robot
WORKDIR /game-robot

ARG CI_SERVER_NAME

RUN cd /game-robot && \
  BUILD_TIME=`TZ=Asia/Taipei date +%FT%T%z` && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "$LDFLAGS" -i -o $CI_SERVER_NAME && \
  upx $CI_SERVER_NAME

RUN cd ./ && mkdir log
# gamerobot binary
FROM alpine:3.10
ARG CI_SERVER_NAME

COPY --from=build-env /game-robot/config/ /config/
COPY --from=build-env /game-robot/log/ /log/
COPY --from=build-env /game-robot/$CI_SERVER_NAME /
COPY --from=build-env /game-robot/templates /templates
COPY --from=build-env /game-robot/static /static

EXPOSE 8888
ENTRYPOINT ["/${CI_SERVER_NAME}"]
