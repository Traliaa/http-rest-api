/**
 * @fileoverview gRPC-Web generated client stub for api
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.api = require('./api_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.LoginClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.api.LoginPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.api.LoginRequest,
 *   !proto.api.LoginResponse>}
 */
const methodDescriptor_Login_Login = new grpc.web.MethodDescriptor(
  '/api.Login/Login',
  grpc.web.MethodType.UNARY,
  proto.api.LoginRequest,
  proto.api.LoginResponse,
  /**
   * @param {!proto.api.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.LoginResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.api.LoginRequest,
 *   !proto.api.LoginResponse>}
 */
const methodInfo_Login_Login = new grpc.web.AbstractClientBase.MethodInfo(
  proto.api.LoginResponse,
  /**
   * @param {!proto.api.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.api.LoginResponse.deserializeBinary
);


/**
 * @param {!proto.api.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.api.LoginResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.api.LoginResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.api.LoginClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/api.Login/Login',
      request,
      metadata || {},
      methodDescriptor_Login_Login,
      callback);
};


/**
 * @param {!proto.api.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.api.LoginResponse>}
 *     A native promise that resolves to the response
 */
proto.api.LoginPromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/api.Login/Login',
      request,
      metadata || {},
      methodDescriptor_Login_Login);
};


module.exports = proto.api;

