const { HelloRequest } = require('./api/proto/simpleMath_pb');
const { HelloWorldServiceClient } = require('./api/v1/hello_grpc_web_pb');

// 注意这个端口是代理服务器的端口，不是grpc的端口
var client = new HelloWorldServiceClient('http://localhost:8199',
    null, null);

// simple unary call
var request = new HelloRequest();
request.setName('World');

client.sayHello(request, {}, (err, response) => {
    document.getElementById("response").innerHTML = response.getMessage();
});