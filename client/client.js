export function CreateClient(address) {
    // import grpc package
    const grpc = require('@grpc/grpc-js');

    // out proto file path
    const PROTO_PATH = '../proto/service.proto';

    // import proto loader
    let protoLoader = require('@grpc/proto-loader');

    // grpc loader options
    const options = {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    };

    // Suggested options for similarity to existing grpc.load behavior
    let packageDefinition = protoLoader.loadSync(
        PROTO_PATH,
        options
    );

    // The protoDescriptor object has the full package hierarchy
    const stanG = grpc.loadPackageDefinition(packageDefinition).service.StanG;

    // create a client
    return new stanG(
        address,
        grpc.credentials.createInsecure()
    )
}
