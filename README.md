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

## Deployment
If you want to deploy this project on kubernetes, just use the following command:
```shell
helm dep up ./deployments/stan-gee
helm install ./deployments/stan-gee
```

With providing the _application.yaml_ so you can set for cluster **ArgoCD**

protoc -I proto/ proto/[file].proto --go_out=plugins=grpc:proto