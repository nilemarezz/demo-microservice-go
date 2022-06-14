## Demo Microservice

start service
```cmd
cd deployment && make start
```

stop service
```cmd
cd deployment && make stop
```

### Service
- front-end (react)
- api-gateway (golang,REST,grpc)
- movie-service (golang,REST,grpc,mysql)

### Dashboard, UI
- frontend http://localhost/
- Jaeger Tracing http://localhost:16686/
- Swagger http://localhost:5000/swagger/index.html

### Tracing
- Opentelemetry
- Jaeger

### Deployment
- docker
- docker-compose

### Document
- Swaggo API Doc [link](https://github.com/swaggo/swag) 

### Resource
- Tracing Example [link](https://github.com/TonPC64/distributed-tracing-in-golang)
- Microservice example [link](https://levelup.gitconnected.com/microservices-with-go-grpc-api-gateway-and-authentication-part-1-2-393ad9fc9d30)
