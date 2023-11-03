# Simple Golang Kafka Implementation

## Technologies

![Go](https://go.dev/)
![Kafka](https://kafka.apache.org/)
![Chi](https://go-chi.io/#/)
![MySQL](https://www.mysql.com/)
![Docker](https://www.docker.com/)

## How to run the project

### Start Docker Compose Containers

```bash
  docker-compose up -d
```

### Create Product DB

```bash
  docker-compose exec mysql bash
```

```bash
  # Password: root
  mysql -u root -p products
```

```bash
  create table products (id varchar(255), name varchar(255), price float);
```

### Create Product Kafka Topic

```bash
  docker-compose exec kafka bash
```

```bash
  kafka topics --bootstrap-server=localhost:9092 --topic=product --create
```

### Run Application

```bash
  docker-compose exec goapp bash
```

```bash
  go run cmd/app/main.go
```

## How to test the project

### HTTP

- Use test.http file to test the HTTP endpoints

### Kafka

```bash
  docker-compose exec kafka bash
```

```bash
  kafka-console-producer --bootstrap-server=localhost:9092 --topic=product
```

- Paste the following JSON to the console

```JSON
  {"name": "My Product 2", "price": 200}
```

- Try to send http request to the find all endpoint and see the My Product 2 in the response
