## Per User Counter but using microservices

### Scope:
- Gateway/Broker service
	- Handles authentication
	- Pass on messages to rabbitMQ
- Service A:
	- Authentication microservice
- Service B:
	- Logger microservice
- Service C:
	- Simple counter service.
	- Memcached or some other cache as data store
- Prometheus for monitoring
- Loki for logs

probably more later on
