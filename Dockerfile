FROM scratch

WORKDIR /app

COPY ./client .
COPY ./config ./config
COPY ./files ./files
COPY ./localtime /etc/localtime

ENTRYPOINT ["./client"]
