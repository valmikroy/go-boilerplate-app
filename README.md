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



### TODO
- Introduce Viper for configuration file management
- Log level parameter checks
- adjust log levels through cmd line param
