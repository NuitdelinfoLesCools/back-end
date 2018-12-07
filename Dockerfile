FROM alpine:3.7

# Smooth 10Mo image

ADD server.o .

ADD config.yaml .

EXPOSE 9000

ENTRYPOINT [ "./server.o" ]