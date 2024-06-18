go build -C auth -o build
go build -C gateway -o build
go build -C counter -o build

docker compose down
docker image rm counter-microservices-auth
docker image rm counter-microservices-gateway
docker image rm counter-microservices-counter
docker compose up -d

# wait for postgres to start
sleep 2
cd auth/migrations && goose postgres "postgres://postgres:password@localhost:5432/auth" up
cd ../../counter/migrations && goose postgres "postgres://postgres:password@localhost:5433/counter" up