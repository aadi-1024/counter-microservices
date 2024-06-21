## Per User Counter but using microservices

### Architecture:
- **Gateway:**
	- Exposes HTTP endpoints
	- Checks for and verifies JWT token
	- Communicates with Auth and Counter using gRPC
- **Auth:**
	- Handles Login and Registration
	- Uses PostgreSQL as database and gRPC to communicate with Gateway
- **Counter:**
	- The main service
	- Communicates with Gateway using gRPC
	- Uses PostgreSQL as database and Memcached as a caching layer
- **Logger:**
	- Sends logs to an Elasticsearch instance running along with Kibana for visualisation
	- Collects logs from a RabbitMQ instance before sending to Elasticsearch
	- Uses a multithreading-safe worker pool to send several requests concurrently
- **Logger Client:**
	- Client library that provides an API for different services to send logs to the Logger service
	- Passes on the Log messages to the RabbitMQ instance to be picked up by the Logger

Logger can easily be scaled horizontally by multiplying the instances. Other services might require a bit more work.
