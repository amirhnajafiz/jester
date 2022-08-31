<p align="center">
<img src="./assets/main.png" width="700" />
</p>

<h1 align="center">
Stan Gee
</h1>

Nats streaming service in Golang, With a JS client.

This is a gRPC service to send and receive messages over a nats-streaming
server. The main idea behind this service is to make a communication system
between clients of different programming languages by using proto-buffers 
and a gRPC server.

This service connects to three nats-streaming services.<br />
Clients will subscribe to our gRPC, which subscribes on the stan service. Then
the clients will send their data over the publish method on the stan service, waiting
for the other clients to receive their message on the topic that
they defined in the request.

## How to use this project?
Main service methods:
```protobuf
service StanG {
  rpc Sub(Send) returns (stream Catch) {}
  rpc Put(Data) returns (Response) {}
}
```

### Subscribe
Send data for subscribe:
```protobuf
message Send {
  string topic = 1; // NATS topic
  string type = 2;  // Subscribe type
}
```

Catch response for subscribe:
```protobuf
message Catch {
  string content = 1; // NATS sub content
}
```

### Publish
Request body for publish:
```protobuf
message Data {
  string topic = 1; // NATS topic
  string content = 2; // Publish data
}
```

Response of publish:
```protobuf
message Response {
  int32 status = 1; // Response status
  string message = 2; // Response message
}
```

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

You can also [check](client/README.md) my client implemented in **JavaScript** programming language.

## Deployment
If you want to deploy this project on kubernetes, just use the following command:
```shell
helm dep up ./deployments/stan-gee
helm install ./deployments/stan-gee
```

With providing the _application.yaml_ so you can set for cluster **ArgoCD**
