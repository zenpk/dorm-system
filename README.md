# dorm-system

A dormitory selecting system in microservice architecture which supports high concurrency, written in Go.

## Architecture

![architecture](./assets/img/architecture.png)

## Dependencies

### Backend (Go)

| Name   | Usage           | URL                          |
| :----- | --------------- | ---------------------------- |
| Gin    | HTTP Framework  | github.com/gin-gonic/gin     |
| gRPC   | gRPC            | google.golang.org/grpc       |
| sarama | Kafka Connector | github.com/Shopify/sarama    |
| GORM   | MySQL ORM       | gorm.io/gorm                 |
| Viper  | Configuration   | github.com/spf13/viper       |
| zap    | Logging         | go.uber.org/zap              |
| JWT    | JWT             | github.com/golang-jwt/jwt/v4 |

### Frontend (TypeScript)

| Name      | Usage           | URL                               |
| --------- | --------------- | --------------------------------- |
| Next.js   | React Framework | <https://github.com/vercel/next.js> |
| React     | React           | <https://github.com/facebook/react> |
| Bootstrap | CSS             | <https://github.com/twbs/bootstrap> |

## Getting started

### Setup the databases

This project uses two databases: MySQL and Redis, Redis is for caching. By default they run in Docker containers, you can use a docker compose command to set them both up.

```shell
cd scripts/docker_db
sudo docker compose up -d
```

 The default configurations are as below, change them in `docker-compose.yml` and `redis.conf`:

| Key                 | Value      |
| ------------------- | ---------- |
| MySQL root password | root       |
| MySQL port          | 3306       |
| Redis auth password | redis_pass |
| Redis port          | 6379       |

### Create dummy data

Use Python scripts to create dummy data, change the MySQL configuration in `util.py` and run:

```shell
cd scripts/python
python3 main.py
```

This will create the data below:

| Data      | Value                                                        |
| --------- | ------------------------------------------------------------ |
| buildings | 5 enabled, 1 not enabled                                     |
| dorms     | 100 dorms with random gender and bed counts                  |
| users     | ~1000 users with random gender and name, along with the correspondent account |
| test user | username: temp; password: temp                               |
| teams     | 1 team: test user + first four users                         |

### Setup the message queue

Kafka and ZooKeeper also run in Docker containers

```shell
cd scripts/docker_kafka
sudo docker compose up -d
```

Default configurations are as below, change them in `docker-compose.yml`

| Key                     | Value     |
| ----------------------- | --------- |
| Kafka port              | 19092     |
| Kafka security protocol | PLAINTEXT |

### Build the microservices

Use Makefile to build all executable files in one command

```shell
make -C cmd build-all
```

All binary executable files are located in `bin/`

The program read configuration files at runtime, modify them in `configs/`

### Run microservices

```shell
sudo bin/main -mode=[development|production|testing]
```

The mode decides which configuration file to read

### Run frontend

```shell
cd web/dorm-system-frontend
npm run start
```

Change any configuration you need in `next.config.js`

Now you should be able to explore the whole system in your browser at `localhost:3000`, enjoy!
