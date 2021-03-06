package main

const envTemplate = `# Basic env configuration
ENVIRONMENT=development #development,staging,production
DEBUG_MODE=true

# Service Handlers
## Server
USE_REST={{.RestHandler}}
USE_GRPC={{.GRPCHandler}}
USE_GRAPHQL={{.GraphQLHandler}}
## Worker
USE_KAFKA_CONSUMER={{.KafkaHandler}}
USE_CRON_SCHEDULER={{.SchedulerHandler}}
USE_REDIS_SUBSCRIBER={{.RedisSubsHandler}}
USE_TASK_QUEUE_WORKER={{.TaskQueueHandler}}

HTTP_PORT=8000
GRPC_PORT=8002

BASIC_AUTH_USERNAME=user
BASIC_AUTH_PASS=pass

MONGODB_HOST_WRITE=mongodb://localhost:27017
MONGODB_HOST_READ=mongodb://localhost:27017
MONGODB_DATABASE_NAME={{.ServiceName}}

SQL_DRIVER_NAME={{.SQLDriver}}
SQL_DB_READ_HOST=[string]
SQL_DB_READ_USER=[string]
SQL_DB_READ_PASSWORD=[string]
SQL_DB_WRITE_HOST=[string]
SQL_DB_WRITE_USER=[string]
SQL_DB_WRITE_PASSWORD=[string]
SQL_DATABASE_NAME=[string]

REDIS_READ_HOST=localhost
REDIS_READ_PORT=6379
REDIS_READ_AUTH=
REDIS_WRITE_TLS=false
REDIS_READ_DB=0
REDIS_WRITE_HOST=localhost
REDIS_WRITE_PORT=6379
REDIS_WRITE_AUTH=
REDIS_WRITE_TLS=false
REDIS_WRITE_DB=0

KAFKA_BROKERS=localhost:9092
KAFKA_CLIENT_VERSION=2.0.0
KAFKA_CLIENT_ID={{.ServiceName}}
KAFKA_CONSUMER_GROUP={{.ServiceName}}

JAEGER_TRACING_HOST=127.0.0.1:5775
GRAPHQL_SCHEMA_DIR="api/graphql"
JSON_SCHEMA_DIR="api/jsonschema"

MAX_GOROUTINES=100

# Additional env
`
