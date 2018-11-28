FROM alpine

WORKDIR /app

COPY gamelink-apns ./

RUN apk update && apk add --no-cache ca-certificates

ENTRYPOINT [ "./gamelink-apns" ]


