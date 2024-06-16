## Per User Counter but using microservices

### Scope:
- Gateway/Broker service
	- Handles authentication
	- ~~Pass on messages to rabbitMQ~~
	- Use gRPC to communicate with services
- Service A:
	- Authentication microservice
	- Was trying CassandraDB but the Go library is unmaintained so postgres it is
- Service B:
	- Logger microservice
	- Use rabbitMQ to receive logs
- Service C:
	- Simple counter service.
	- Memcached or some other cache as data store
- Prometheus for monitoring
- Loki for logs

probably more later on
