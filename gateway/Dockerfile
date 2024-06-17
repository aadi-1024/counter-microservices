FROM ubuntu

RUN mkdir app
WORKDIR /app
COPY ./gateway/build .
RUN chmod +x build

ENTRYPOINT [ "./build" ]