FROM ubuntu

RUN mkdir app
WORKDIR /app
COPY ./auth/build .
RUN chmod +x build

ENTRYPOINT [ "./build" ]