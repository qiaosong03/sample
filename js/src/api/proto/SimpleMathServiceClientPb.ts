/**
 * @fileoverview gRPC-Web generated client stub for simpleMath
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_simpleMath_pb from '../proto/simpleMath_pb';


export class SimpleMathServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoadd = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_simpleMath_pb.addResponse,
    (request: proto_simpleMath_pb.addRequest) => {
      return request.serializeBinary();
    },
    proto_simpleMath_pb.addResponse.deserializeBinary
  );

  add(
    request: proto_simpleMath_pb.addRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_simpleMath_pb.addResponse>;

  add(
    request: proto_simpleMath_pb.addRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_simpleMath_pb.addResponse) => void): grpcWeb.ClientReadableStream<proto_simpleMath_pb.addResponse>;

  add(
    request: proto_simpleMath_pb.addRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_simpleMath_pb.addResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/simpleMath.SimpleMathService/add',
        request,
        metadata || {},
        this.methodInfoadd,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/simpleMath.SimpleMathService/add',
    request,
    metadata || {},
    this.methodInfoadd);
  }

  methodInfosubtract = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_simpleMath_pb.subtractResponse,
    (request: proto_simpleMath_pb.subtractRequest) => {
      return request.serializeBinary();
    },
    proto_simpleMath_pb.subtractResponse.deserializeBinary
  );

  subtract(
    request: proto_simpleMath_pb.subtractRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_simpleMath_pb.subtractResponse>;

  subtract(
    request: proto_simpleMath_pb.subtractRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_simpleMath_pb.subtractResponse) => void): grpcWeb.ClientReadableStream<proto_simpleMath_pb.subtractResponse>;

  subtract(
    request: proto_simpleMath_pb.subtractRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_simpleMath_pb.subtractResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/simpleMath.SimpleMathService/subtract',
        request,
        metadata || {},
        this.methodInfosubtract,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/simpleMath.SimpleMathService/subtract',
    request,
    metadata || {},
    this.methodInfosubtract);
  }

  methodInfomultiply = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_simpleMath_pb.multiplyResponse,
    (request: proto_simpleMath_pb.multiplyRequest) => {
      return request.serializeBinary();
    },
    proto_simpleMath_pb.multiplyResponse.deserializeBinary
  );

  multiply(
    request: proto_simpleMath_pb.multiplyRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_simpleMath_pb.multiplyResponse>;

  multiply(
    request: proto_simpleMath_pb.multiplyRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_simpleMath_pb.multiplyResponse) => void): grpcWeb.ClientReadableStream<proto_simpleMath_pb.multiplyResponse>;

  multiply(
    request: proto_simpleMath_pb.multiplyRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_simpleMath_pb.multiplyResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/simpleMath.SimpleMathService/multiply',
        request,
        metadata || {},
        this.methodInfomultiply,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/simpleMath.SimpleMathService/multiply',
    request,
    metadata || {},
    this.methodInfomultiply);
  }

  methodInfodivide = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_simpleMath_pb.divideResponse,
    (request: proto_simpleMath_pb.divideRequest) => {
      return request.serializeBinary();
    },
    proto_simpleMath_pb.divideResponse.deserializeBinary
  );

  divide(
    request: proto_simpleMath_pb.divideRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_simpleMath_pb.divideResponse>;

  divide(
    request: proto_simpleMath_pb.divideRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: proto_simpleMath_pb.divideResponse) => void): grpcWeb.ClientReadableStream<proto_simpleMath_pb.divideResponse>;

  divide(
    request: proto_simpleMath_pb.divideRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: proto_simpleMath_pb.divideResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/simpleMath.SimpleMathService/divide',
        request,
        metadata || {},
        this.methodInfodivide,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/simpleMath.SimpleMathService/divide',
    request,
    metadata || {},
    this.methodInfodivide);
  }

  methodInfoserverStream = new grpcWeb.AbstractClientBase.MethodInfo(
    proto_simpleMath_pb.squareResponse,
    (request: proto_simpleMath_pb.squareRequest) => {
      return request.serializeBinary();
    },
    proto_simpleMath_pb.squareResponse.deserializeBinary
  );

  serverStream(
    request: proto_simpleMath_pb.squareRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/simpleMath.SimpleMathService/serverStream',
      request,
      metadata || {},
      this.methodInfoserverStream);
  }

}

