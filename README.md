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
- movie-service (golang,grpc,mysql)
- auth-service (golang,grpc,mysql)
<img src="https://www.img.in.th/images/5de6052467964c69d2ddb4c66a570c88.png" alt="5de6052467964c69d2ddb4c66a570c88.png" border="0" />

### Dashboard, UI
- frontend http://localhost/
- Jaeger Tracing http://localhost:16686/
- Swagger http://localhost:5000/swagger/index.html

### Tracing
- Opentelemetry
- Jaeger
<img src="https://sv1.picz.in.th/images/2022/06/14/VBxIpP.png" alt="VBxIpP.png" border="0" />

### Deployment
- docker
- docker-compose

### Document
- Swaggo API Doc [link](https://github.com/swaggo/swag) 
<img src="https://sv1.picz.in.th/images/2022/06/14/VBxUWv.png" alt="VBxUWv.png" border="0" />

### Resource
- Tracing Example [link](https://github.com/TonPC64/distributed-tracing-in-golang)
- Microservice example [link](https://levelup.gitconnected.com/microservices-with-go-grpc-api-gateway-and-authentication-part-1-2-393ad9fc9d30)
- Docker volumn error [link](https://stackoverflow.com/questions/30604846/docker-error-no-space-left-on-device)
