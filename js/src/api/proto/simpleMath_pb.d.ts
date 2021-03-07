import * as jspb from 'google-protobuf'



export class addRequest extends jspb.Message {
  getA(): number;
  setA(value: number): addRequest;

  getB(): number;
  setB(value: number): addRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): addRequest.AsObject;
  static toObject(includeInstance: boolean, msg: addRequest): addRequest.AsObject;
  static serializeBinaryToWriter(message: addRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): addRequest;
  static deserializeBinaryFromReader(message: addRequest, reader: jspb.BinaryReader): addRequest;
}

export namespace addRequest {
  export type AsObject = {
    a: number,
    b: number,
  }
}

export class addResponse extends jspb.Message {
  getRes(): number;
  setRes(value: number): addResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): addResponse.AsObject;
  static toObject(includeInstance: boolean, msg: addResponse): addResponse.AsObject;
  static serializeBinaryToWriter(message: addResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): addResponse;
  static deserializeBinaryFromReader(message: addResponse, reader: jspb.BinaryReader): addResponse;
}

export namespace addResponse {
  export type AsObject = {
    res: number,
  }
}

export class subtractRequest extends jspb.Message {
  getA(): number;
  setA(value: number): subtractRequest;

  getB(): number;
  setB(value: number): subtractRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): subtractRequest.AsObject;
  static toObject(includeInstance: boolean, msg: subtractRequest): subtractRequest.AsObject;
  static serializeBinaryToWriter(message: subtractRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): subtractRequest;
  static deserializeBinaryFromReader(message: subtractRequest, reader: jspb.BinaryReader): subtractRequest;
}

export namespace subtractRequest {
  export type AsObject = {
    a: number,
    b: number,
  }
}

export class subtractResponse extends jspb.Message {
  getRes(): number;
  setRes(value: number): subtractResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): subtractResponse.AsObject;
  static toObject(includeInstance: boolean, msg: subtractResponse): subtractResponse.AsObject;
  static serializeBinaryToWriter(message: subtractResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): subtractResponse;
  static deserializeBinaryFromReader(message: subtractResponse, reader: jspb.BinaryReader): subtractResponse;
}

export namespace subtractResponse {
  export type AsObject = {
    res: number,
  }
}

export class multiplyRequest extends jspb.Message {
  getA(): number;
  setA(value: number): multiplyRequest;

  getB(): number;
  setB(value: number): multiplyRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): multiplyRequest.AsObject;
  static toObject(includeInstance: boolean, msg: multiplyRequest): multiplyRequest.AsObject;
  static serializeBinaryToWriter(message: multiplyRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): multiplyRequest;
  static deserializeBinaryFromReader(message: multiplyRequest, reader: jspb.BinaryReader): multiplyRequest;
}

export namespace multiplyRequest {
  export type AsObject = {
    a: number,
    b: number,
  }
}

export class multiplyResponse extends jspb.Message {
  getRes(): number;
  setRes(value: number): multiplyResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): multiplyResponse.AsObject;
  static toObject(includeInstance: boolean, msg: multiplyResponse): multiplyResponse.AsObject;
  static serializeBinaryToWriter(message: multiplyResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): multiplyResponse;
  static deserializeBinaryFromReader(message: multiplyResponse, reader: jspb.BinaryReader): multiplyResponse;
}

export namespace multiplyResponse {
  export type AsObject = {
    res: number,
  }
}

export class divideRequest extends jspb.Message {
  getA(): number;
  setA(value: number): divideRequest;

  getB(): number;
  setB(value: number): divideRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): divideRequest.AsObject;
  static toObject(includeInstance: boolean, msg: divideRequest): divideRequest.AsObject;
  static serializeBinaryToWriter(message: divideRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): divideRequest;
  static deserializeBinaryFromReader(message: divideRequest, reader: jspb.BinaryReader): divideRequest;
}

export namespace divideRequest {
  export type AsObject = {
    a: number,
    b: number,
  }
}

export class divideResponse extends jspb.Message {
  getRes(): number;
  setRes(value: number): divideResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): divideResponse.AsObject;
  static toObject(includeInstance: boolean, msg: divideResponse): divideResponse.AsObject;
  static serializeBinaryToWriter(message: divideResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): divideResponse;
  static deserializeBinaryFromReader(message: divideResponse, reader: jspb.BinaryReader): divideResponse;
}

export namespace divideResponse {
  export type AsObject = {
    res: number,
  }
}

export class squareRequest extends jspb.Message {
  getA(): number;
  setA(value: number): squareRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): squareRequest.AsObject;
  static toObject(includeInstance: boolean, msg: squareRequest): squareRequest.AsObject;
  static serializeBinaryToWriter(message: squareRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): squareRequest;
  static deserializeBinaryFromReader(message: squareRequest, reader: jspb.BinaryReader): squareRequest;
}

export namespace squareRequest {
  export type AsObject = {
    a: number,
  }
}

export class squareResponse extends jspb.Message {
  getStr(): string;
  setStr(value: string): squareResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): squareResponse.AsObject;
  static toObject(includeInstance: boolean, msg: squareResponse): squareResponse.AsObject;
  static serializeBinaryToWriter(message: squareResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): squareResponse;
  static deserializeBinaryFromReader(message: squareResponse, reader: jspb.BinaryReader): squareResponse;
}

export namespace squareResponse {
  export type AsObject = {
    str: string,
  }
}

export class helloRequest extends jspb.Message {
  getName(): string;
  setName(value: string): helloRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): helloRequest.AsObject;
  static toObject(includeInstance: boolean, msg: helloRequest): helloRequest.AsObject;
  static serializeBinaryToWriter(message: helloRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): helloRequest;
  static deserializeBinaryFromReader(message: helloRequest, reader: jspb.BinaryReader): helloRequest;
}

export namespace helloRequest {
  export type AsObject = {
    name: string,
  }
}

export class helloResponse extends jspb.Message {
  getHello(): string;
  setHello(value: string): helloResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): helloResponse.AsObject;
  static toObject(includeInstance: boolean, msg: helloResponse): helloResponse.AsObject;
  static serializeBinaryToWriter(message: helloResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): helloResponse;
  static deserializeBinaryFromReader(message: helloResponse, reader: jspb.BinaryReader): helloResponse;
}

export namespace helloResponse {
  export type AsObject = {
    hello: string,
  }
}

