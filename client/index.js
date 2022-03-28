// import packages
let grpc = require('@grpc/grpc-js');
let protoLoader = require('@grpc/proto-loader');

// out proto file path
let PROTO_PATH = __dirname + '/../proto/service.proto';

// Suggested options for similarity to existing grpc.load behavior
let packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
    }
);

let protoDescriptor = grpc.loadPackageDefinition(packageDefinition)
// The protoDescriptor object has the full package hierarchy
let service = protoDescriptor.constructor;
