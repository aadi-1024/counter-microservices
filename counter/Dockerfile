FROM ubuntu

RUN mkdir app
WORKDIR /app
COPY ./counter/build .
RUN chmod +x build

ENTRYPOINT [ "./build" ]