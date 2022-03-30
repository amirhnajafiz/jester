# Stan-Gee
Nats streaming service in Golang, With a JS client.

## What is this service for?
This is a gRPC service to send and receive messages over a nats-streaming
server. The main idea behind this service is to make a communication system
between clients of different programming languages by using proto-buffers 
and a gRPC server.

## How does it work?

## How to use?

## Building other clients yourself
The **service.proto** file in _./proto_ directory describes the logic of our application,
you can make your own clients in different programming languages.

#### Golang
```shell
protoc -I proto/ proto/service.proto --go_out=plugins=grpc:proto
```

#### Javascript
```shell
protoc -I proto/ proto/service.proto --js_out=library=grpc:proto
```

#### Java
```shell
protoc -I proto/ proto/service.proto --java_out=build/proto
```

#### Python
```shell
protoc -I proto/ proto/service.proto --python_out=build/proto
```

And ....

## Deployment
If you want to deploy this project on kubernetes, just use the following command:
```shell
helm dep up ./deployments/stan-gee
helm install ./deployments/stan-gee
```

With providing the _application.yaml_ so you can set for cluster **ArgoCD**
