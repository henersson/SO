const grpc = {};
grpc.web = require('grpc-web');
const proto = require('./servicios_pb.js');

const StreamingServiceClient = function(hostport, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';
  this.client_ = new grpc.web.GrpcWebClientBase(options);
  this.hostname_ = hostport;
};

StreamingServiceClient.prototype.reproducirAudio = function(request, metadata) {
  return this.client_.serverStreaming(
    this.hostname_ + '/servicios.StreamingService/ReproducirAudio',
    request,
    metadata || {},
    new grpc.web.MethodDescriptor(
      '/servicios.StreamingService/ReproducirAudio',
      grpc.web.MethodType.SERVER_STREAMING,
      proto.ReproducirAudioRequest,
      proto.AudioChunk,
      function(r) { return r.serializeBinary(); },
      proto.AudioChunk.deserializeBinary
    )
  );
};

const MetadataServiceClient = function(hostport, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';
  this.client_ = new grpc.web.GrpcWebClientBase(options);
  this.hostname_ = hostport;
};

MetadataServiceClient.prototype.obtenerTiposAudio = function(request, metadata, callback) {
  return this.client_.rpcCall(
    this.hostname_ + '/servicios.MetadataService/ObtenerTiposAudio',
    request, metadata || {},
    new grpc.web.MethodDescriptor('/servicios.MetadataService/ObtenerTiposAudio', grpc.web.MethodType.UNARY, proto.Empty, proto.TiposAudioResponse, function(r){return r.serializeBinary();}, proto.TiposAudioResponse.deserializeBinary),
    callback
  );
};

MetadataServiceClient.prototype.obtenerAudiosPorTipo = function(request, metadata, callback) {
  return this.client_.rpcCall(
    this.hostname_ + '/servicios.MetadataService/ObtenerAudiosPorTipo',
    request, metadata || {},
    new grpc.web.MethodDescriptor('/servicios.MetadataService/ObtenerAudiosPorTipo', grpc.web.MethodType.UNARY, proto.AudiosPorTipoRequest, proto.AudiosPorTipoResponse, function(r){return r.serializeBinary();}, proto.AudiosPorTipoResponse.deserializeBinary),
    callback
  );
};

MetadataServiceClient.prototype.obtenerDetalleAudio = function(request, metadata, callback) {
  return this.client_.rpcCall(
    this.hostname_ + '/servicios.MetadataService/ObtenerDetalleAudio',
    request, metadata || {},
    new grpc.web.MethodDescriptor('/servicios.MetadataService/ObtenerDetalleAudio', grpc.web.MethodType.UNARY, proto.DetalleAudioRequest, proto.DetalleAudioResponse, function(r){return r.serializeBinary();}, proto.DetalleAudioResponse.deserializeBinary),
    callback
  );
};

MetadataServiceClient.prototype.login = function(request, metadata, callback) {
  return this.client_.rpcCall(
    this.hostname_ + '/servicios.MetadataService/Login',
    request, metadata || {},
    new grpc.web.MethodDescriptor('/servicios.MetadataService/Login', grpc.web.MethodType.UNARY, proto.LoginRequest, proto.LoginResponse, function(r){return r.serializeBinary();}, proto.LoginResponse.deserializeBinary),
    callback
  );
};

MetadataServiceClient.prototype.registrarUsuario = function(request, metadata, callback) {
  return this.client_.rpcCall(
    this.hostname_ + '/servicios.MetadataService/RegistrarUsuario',
    request, metadata || {},
    new grpc.web.MethodDescriptor('/servicios.MetadataService/RegistrarUsuario', grpc.web.MethodType.UNARY, proto.RegistrarUsuarioRequest, proto.RegistrarUsuarioResponse, function(r){return r.serializeBinary();}, proto.RegistrarUsuarioResponse.deserializeBinary),
    callback
  );
};

MetadataServiceClient.prototype.almacenarAudio = function(request, metadata, callback) {
  return this.client_.rpcCall(
    this.hostname_ + '/servicios.MetadataService/AlmacenarAudio',
    request, metadata || {},
    new grpc.web.MethodDescriptor('/servicios.MetadataService/AlmacenarAudio', grpc.web.MethodType.UNARY, proto.AlmacenarAudioRequest, proto.AlmacenarAudioResponse, function(r){return r.serializeBinary();}, proto.AlmacenarAudioResponse.deserializeBinary),
    callback
  );
};

module.exports.StreamingServiceClient = StreamingServiceClient;
module.exports.MetadataServiceClient = MetadataServiceClient;
