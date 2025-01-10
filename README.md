# go-boilerplate-app

This application effort adopted from [go-boilerplate](https://github.com/codoworks/go-boilerplate) to understand the go app structures.


Main branch test module 

```
# Run 
go run .

# Run test
go test -v ./...
```

# Start a service and do the application health check 

```
# start the service 
go run main.go  --dev service


# run health check 
curl --header "Accept: application/json" --header "Content-Type: application/json" http://localhost:8080/health/alive
{"code":"20001","message":"success","errors":[],"payload":{"message":"ok","version":""}}


curl --header "Accept: application/json" --header "Content-Type: application/json" http://localhost:8080/health/ready
{"code":"20001","message":"success","errors":[],"payload":{"message":"ready"}}
```



# test random number generator 

```
# Index (capped at 6 by default)
curl --header "Accept: application/json" --header "Content-Type: application/json" http://localhost:8080/rolldice

# GET with max (capped at 200)
curl --header "Accept: application/json" --header "Content-Type: application/json" http://localhost:8080/rolldice/200  

# POST with min and max 
curl -d '{"min":"100", "max":"200"}' -X POST --header "Accept: application/json" --header "Content-Type: application/json" http://localhost:8080/rolldice

```


# Build docker container 
```
# point to minikube image registary 
eval $(minikube -p minikube docker-env)

# build and push in local image registary
docker build -t go-boilerplate-app .
```


### TODO
- Introduce Viper for configuration file management
- Log level parameter checks
- adjust log levels through cmd line param
