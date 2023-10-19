FROM golang:alpine as build-env

ARG SERVICE_NAME=traffic-farm-crawler

RUN mkdir /service
ADD . /service/

WORKDIR /service
RUN apk add --no-cache git
RUN go build  -o ${SERVICE_NAME} .


FROM debian:stable-slim
WORKDIR /service
COPY --from=build-env /service/${SERVICE_NAME}    /service/${SERVICE_NAME}
# RUN apt-get update \
#       && apt-get install -y wget gnupg \
#       && wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
#       && sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list' \
#       && apt-get update \
#       && apt-get -y install xvfb google-chrome-stable 

EXPOSE 7778

# CMD ["sh", "-c", "Xvfb :99 -screen 0 1024x768x16 & DISPLAY=:99 /service/traffic-farm-crawler"]

ENTRYPOINT ["/service/traffic-farm-crawler"]