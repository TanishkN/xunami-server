const path = require("path");
const grpc = require("@grpc/grpc.js");
const protoLoader = required("@grpc/proto-loader");



const proto = protoLoader.loadSync(path.join(__dirname,"..","post_service.proto"));
const definition = grpc.loadPackageDefinition(proto); ////Load definitios from proto file into grpc frmwork

const postList = [//Example posts to send
    {id:1, title:'title1',test:'Text1'},
    {id:1, title:'title1',test:'Text1'},
    {id:1, title:'title1',test:'Text1'}
]

//Use call from function and callback
const getPosts= (call, callback) => {
    callback(null,{ posts: postList })
};

const serverURL = 'localhost:10000';
const server = grpc.Server();
server.addService(definition.Postservice.service)//Load protofile definition into grpcs framework
server.bindAsync(serverURL, grpc.ServerCredentials.createInsecure(),port =>{
    server.start()
})