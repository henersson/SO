// Generado manualmente desde audio.proto
// package servicios

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.servicios.AlmacenarAudioRequest', null, global);
goog.exportSymbol('proto.servicios.AlmacenarAudioResponse', null, global);
goog.exportSymbol('proto.servicios.AudioChunk', null, global);
goog.exportSymbol('proto.servicios.AudioResumen', null, global);
goog.exportSymbol('proto.servicios.CampoMetadata', null, global);
goog.exportSymbol('proto.servicios.DetalleAudioRequest', null, global);
goog.exportSymbol('proto.servicios.DetalleAudioResponse', null, global);
goog.exportSymbol('proto.servicios.Empty', null, global);
goog.exportSymbol('proto.servicios.LoginRequest', null, global);
goog.exportSymbol('proto.servicios.LoginResponse', null, global);
goog.exportSymbol('proto.servicios.RegistrarUsuarioRequest', null, global);
goog.exportSymbol('proto.servicios.RegistrarUsuarioResponse', null, global);
goog.exportSymbol('proto.servicios.AudiosPorTipoRequest', null, global);
goog.exportSymbol('proto.servicios.AudiosPorTipoResponse', null, global);
goog.exportSymbol('proto.servicios.ReproducirAudioRequest', null, global);
goog.exportSymbol('proto.servicios.TipoAudio', null, global);
goog.exportSymbol('proto.servicios.TiposAudioResponse', null, global);

// ---- Empty ----
proto.servicios.Empty = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.servicios.Empty, jspb.Message);
proto.servicios.Empty.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.Empty.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.Empty.serializeBinaryToWriter = function(m, w) {};
proto.servicios.Empty.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.Empty(); return proto.servicios.Empty.deserializeBinaryFromReader(m, r);
};
proto.servicios.Empty.deserializeBinaryFromReader = function(m, r) { return m; };
proto.servicios.Empty.prototype.toObject = function(opt_includeInstance) { return {}; };

// ---- AudioChunk ----
proto.servicios.AudioChunk = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.servicios.AudioChunk, jspb.Message);
proto.servicios.AudioChunk.prototype.getData = function() { return jspb.Message.getFieldWithDefault(this, 1, ''); };
proto.servicios.AudioChunk.prototype.getData_asB64 = function() { return jspb.Message.bytesAsB64(this.getData()); };
proto.servicios.AudioChunk.prototype.getData_asU8 = function() { return jspb.Message.bytesAsU8(this.getData()); };
proto.servicios.AudioChunk.prototype.setData = function(v) { return jspb.Message.setProto3BytesField(this, 1, v); };
proto.servicios.AudioChunk.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.AudioChunk.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.AudioChunk.serializeBinaryToWriter = function(m, w) {
  var f = m.getData_asU8(); if (f.length > 0) w.writeBytes(1, f);
};
proto.servicios.AudioChunk.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.AudioChunk(); return proto.servicios.AudioChunk.deserializeBinaryFromReader(m, r);
};
proto.servicios.AudioChunk.deserializeBinaryFromReader = function(m, r) {
  while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); if (f === 1) { var v = r.readBytes(); m.setData(v); } } return m;
};
proto.servicios.AudioChunk.prototype.toObject = function() { return { data: this.getData_asB64() }; };

// ---- TipoAudio ----
proto.servicios.TipoAudio = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.servicios.TipoAudio, jspb.Message);
proto.servicios.TipoAudio.prototype.getId = function() { return jspb.Message.getFieldWithDefault(this, 1, 0); };
proto.servicios.TipoAudio.prototype.setId = function(v) { return jspb.Message.setProto3IntField(this, 1, v); };
proto.servicios.TipoAudio.prototype.getNombre = function() { return jspb.Message.getFieldWithDefault(this, 2, ''); };
proto.servicios.TipoAudio.prototype.setNombre = function(v) { return jspb.Message.setProto3StringField(this, 2, v); };
proto.servicios.TipoAudio.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.TipoAudio.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.TipoAudio.serializeBinaryToWriter = function(m, w) {
  var f = m.getId(); if (f !== 0) w.writeInt32(1, f);
  var f = m.getNombre(); if (f.length > 0) w.writeString(2, f);
};
proto.servicios.TipoAudio.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.TipoAudio(); return proto.servicios.TipoAudio.deserializeBinaryFromReader(m, r);
};
proto.servicios.TipoAudio.deserializeBinaryFromReader = function(m, r) {
  while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); switch(f) { case 1: m.setId(r.readInt32()); break; case 2: m.setNombre(r.readString()); break; default: r.skipField(); } } return m;
};
proto.servicios.TipoAudio.prototype.toObject = function() { return { id: this.getId(), nombre: this.getNombre() }; };

// ---- AudioResumen ----
proto.servicios.AudioResumen = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.servicios.AudioResumen, jspb.Message);
proto.servicios.AudioResumen.prototype.getId = function() { return jspb.Message.getFieldWithDefault(this, 1, 0); };
proto.servicios.AudioResumen.prototype.setId = function(v) { return jspb.Message.setProto3IntField(this, 1, v); };
proto.servicios.AudioResumen.prototype.getTitulo = function() { return jspb.Message.getFieldWithDefault(this, 2, ''); };
proto.servicios.AudioResumen.prototype.setTitulo = function(v) { return jspb.Message.setProto3StringField(this, 2, v); };
proto.servicios.AudioResumen.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.AudioResumen.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.AudioResumen.serializeBinaryToWriter = function(m, w) {
  var f = m.getId(); if (f !== 0) w.writeInt32(1, f);
  var f = m.getTitulo(); if (f.length > 0) w.writeString(2, f);
};
proto.servicios.AudioResumen.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.AudioResumen(); return proto.servicios.AudioResumen.deserializeBinaryFromReader(m, r);
};
proto.servicios.AudioResumen.deserializeBinaryFromReader = function(m, r) {
  while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); switch(f) { case 1: m.setId(r.readInt32()); break; case 2: m.setTitulo(r.readString()); break; default: r.skipField(); } } return m;
};
proto.servicios.AudioResumen.prototype.toObject = function() { return { id: this.getId(), titulo: this.getTitulo() }; };

// ---- CampoMetadata ----
proto.servicios.CampoMetadata = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.servicios.CampoMetadata, jspb.Message);
proto.servicios.CampoMetadata.prototype.getClave = function() { return jspb.Message.getFieldWithDefault(this, 1, ''); };
proto.servicios.CampoMetadata.prototype.setClave = function(v) { return jspb.Message.setProto3StringField(this, 1, v); };
proto.servicios.CampoMetadata.prototype.getValor = function() { return jspb.Message.getFieldWithDefault(this, 2, ''); };
proto.servicios.CampoMetadata.prototype.setValor = function(v) { return jspb.Message.setProto3StringField(this, 2, v); };
proto.servicios.CampoMetadata.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.CampoMetadata.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.CampoMetadata.serializeBinaryToWriter = function(m, w) {
  var f = m.getClave(); if (f.length > 0) w.writeString(1, f);
  var f = m.getValor(); if (f.length > 0) w.writeString(2, f);
};
proto.servicios.CampoMetadata.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.CampoMetadata(); return proto.servicios.CampoMetadata.deserializeBinaryFromReader(m, r);
};
proto.servicios.CampoMetadata.deserializeBinaryFromReader = function(m, r) {
  while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); switch(f) { case 1: m.setClave(r.readString()); break; case 2: m.setValor(r.readString()); break; default: r.skipField(); } } return m;
};
proto.servicios.CampoMetadata.prototype.toObject = function() { return { clave: this.getClave(), valor: this.getValor() }; };

// ---- TiposAudioResponse ----
proto.servicios.TiposAudioResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.servicios.TiposAudioResponse.repeatedFields_, null);
};
goog.inherits(proto.servicios.TiposAudioResponse, jspb.Message);
proto.servicios.TiposAudioResponse.repeatedFields_ = [1];
proto.servicios.TiposAudioResponse.prototype.getTiposList = function() { return jspb.Message.getRepeatedWrapperField(this, proto.servicios.TipoAudio, 1); };
proto.servicios.TiposAudioResponse.prototype.setTiposList = function(v) { return jspb.Message.setRepeatedWrapperField(this, 1, v); };
proto.servicios.TiposAudioResponse.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.TiposAudioResponse.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.TiposAudioResponse.serializeBinaryToWriter = function(m, w) {
  var f = m.getTiposList(); if (f.length > 0) w.writeRepeatedMessage(1, f, proto.servicios.TipoAudio.serializeBinaryToWriter);
};
proto.servicios.TiposAudioResponse.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.TiposAudioResponse(); return proto.servicios.TiposAudioResponse.deserializeBinaryFromReader(m, r);
};
proto.servicios.TiposAudioResponse.deserializeBinaryFromReader = function(m, r) {
  while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); if (f === 1) { var v = new proto.servicios.TipoAudio(); r.readMessage(v, proto.servicios.TipoAudio.deserializeBinaryFromReader); m.addTipos(v); } else r.skipField(); } return m;
};
proto.servicios.TiposAudioResponse.prototype.addTipos = function(v, opt_i) { return jspb.Message.addToRepeatedWrapperField(this, 1, v, proto.servicios.TipoAudio, opt_i); };
proto.servicios.TiposAudioResponse.prototype.toObject = function() { return { tiposList: this.getTiposList().map(function(i) { return i.toObject(); }) }; };

// ---- AudiosPorTipoRequest ----
proto.servicios.AudiosPorTipoRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.servicios.AudiosPorTipoRequest, jspb.Message);
proto.servicios.AudiosPorTipoRequest.prototype.getIdtipo = function() { return jspb.Message.getFieldWithDefault(this, 1, 0); };
proto.servicios.AudiosPorTipoRequest.prototype.setIdtipo = function(v) { return jspb.Message.setProto3IntField(this, 1, v); };
proto.servicios.AudiosPorTipoRequest.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.AudiosPorTipoRequest.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.AudiosPorTipoRequest.serializeBinaryToWriter = function(m, w) {
  var f = m.getIdtipo(); if (f !== 0) w.writeInt32(1, f);
};
proto.servicios.AudiosPorTipoRequest.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.AudiosPorTipoRequest(); return proto.servicios.AudiosPorTipoRequest.deserializeBinaryFromReader(m, r);
};
proto.servicios.AudiosPorTipoRequest.deserializeBinaryFromReader = function(m, r) {
  while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); if (f === 1) m.setIdtipo(r.readInt32()); else r.skipField(); } return m;
};
proto.servicios.AudiosPorTipoRequest.prototype.toObject = function() { return { idtipo: this.getIdtipo() }; };

// ---- AudiosPorTipoResponse ----
proto.servicios.AudiosPorTipoResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.servicios.AudiosPorTipoResponse.repeatedFields_, null);
};
goog.inherits(proto.servicios.AudiosPorTipoResponse, jspb.Message);
proto.servicios.AudiosPorTipoResponse.repeatedFields_ = [1];
proto.servicios.AudiosPorTipoResponse.prototype.getAudiosList = function() { return jspb.Message.getRepeatedWrapperField(this, proto.servicios.AudioResumen, 1); };
proto.servicios.AudiosPorTipoResponse.prototype.setAudiosList = function(v) { return jspb.Message.setRepeatedWrapperField(this, 1, v); };
proto.servicios.AudiosPorTipoResponse.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.AudiosPorTipoResponse.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.AudiosPorTipoResponse.serializeBinaryToWriter = function(m, w) {
  var f = m.getAudiosList(); if (f.length > 0) w.writeRepeatedMessage(1, f, proto.servicios.AudioResumen.serializeBinaryToWriter);
};
proto.servicios.AudiosPorTipoResponse.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.AudiosPorTipoResponse(); return proto.servicios.AudiosPorTipoResponse.deserializeBinaryFromReader(m, r);
};
proto.servicios.AudiosPorTipoResponse.deserializeBinaryFromReader = function(m, r) {
  while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); if (f === 1) { var v = new proto.servicios.AudioResumen(); r.readMessage(v, proto.servicios.AudioResumen.deserializeBinaryFromReader); m.addAudios(v); } else r.skipField(); } return m;
};
proto.servicios.AudiosPorTipoResponse.prototype.addAudios = function(v, opt_i) { return jspb.Message.addToRepeatedWrapperField(this, 1, v, proto.servicios.AudioResumen, opt_i); };
proto.servicios.AudiosPorTipoResponse.prototype.toObject = function() { return { audiosList: this.getAudiosList().map(function(i) { return i.toObject(); }) }; };

// ---- DetalleAudioRequest ----
proto.servicios.DetalleAudioRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.servicios.DetalleAudioRequest, jspb.Message);
proto.servicios.DetalleAudioRequest.prototype.getIdaudio = function() { return jspb.Message.getFieldWithDefault(this, 1, 0); };
proto.servicios.DetalleAudioRequest.prototype.setIdaudio = function(v) { return jspb.Message.setProto3IntField(this, 1, v); };
proto.servicios.DetalleAudioRequest.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.DetalleAudioRequest.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.DetalleAudioRequest.serializeBinaryToWriter = function(m, w) {
  var f = m.getIdaudio(); if (f !== 0) w.writeInt32(1, f);
};
proto.servicios.DetalleAudioRequest.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.DetalleAudioRequest(); return proto.servicios.DetalleAudioRequest.deserializeBinaryFromReader(m, r);
};
proto.servicios.DetalleAudioRequest.deserializeBinaryFromReader = function(m, r) {
  while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); if (f === 1) m.setIdaudio(r.readInt32()); else r.skipField(); } return m;
};
proto.servicios.DetalleAudioRequest.prototype.toObject = function() { return { idaudio: this.getIdaudio() }; };

// ---- ReproducirAudioRequest ----
proto.servicios.ReproducirAudioRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.servicios.ReproducirAudioRequest, jspb.Message);
proto.servicios.ReproducirAudioRequest.prototype.getIdaudio = function() { return jspb.Message.getFieldWithDefault(this, 1, 0); };
proto.servicios.ReproducirAudioRequest.prototype.setIdaudio = function(v) { return jspb.Message.setProto3IntField(this, 1, v); };
proto.servicios.ReproducirAudioRequest.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.ReproducirAudioRequest.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.ReproducirAudioRequest.serializeBinaryToWriter = function(m, w) {
  var f = m.getIdaudio(); if (f !== 0) w.writeInt32(1, f);
};
proto.servicios.ReproducirAudioRequest.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.ReproducirAudioRequest(); return proto.servicios.ReproducirAudioRequest.deserializeBinaryFromReader(m, r);
};
proto.servicios.ReproducirAudioRequest.deserializeBinaryFromReader = function(m, r) {
  while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); if (f === 1) m.setIdaudio(r.readInt32()); else r.skipField(); } return m;
};
proto.servicios.ReproducirAudioRequest.prototype.toObject = function() { return { idaudio: this.getIdaudio() }; };

// ---- DetalleAudioResponse ----
proto.servicios.DetalleAudioResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.servicios.DetalleAudioResponse.repeatedFields_, null);
};
goog.inherits(proto.servicios.DetalleAudioResponse, jspb.Message);
proto.servicios.DetalleAudioResponse.repeatedFields_ = [4];
proto.servicios.DetalleAudioResponse.prototype.getId = function() { return jspb.Message.getFieldWithDefault(this, 1, 0); };
proto.servicios.DetalleAudioResponse.prototype.setId = function(v) { return jspb.Message.setProto3IntField(this, 1, v); };
proto.servicios.DetalleAudioResponse.prototype.getIdtipo = function() { return jspb.Message.getFieldWithDefault(this, 2, 0); };
proto.servicios.DetalleAudioResponse.prototype.setIdtipo = function(v) { return jspb.Message.setProto3IntField(this, 2, v); };
proto.servicios.DetalleAudioResponse.prototype.getTitulo = function() { return jspb.Message.getFieldWithDefault(this, 3, ''); };
proto.servicios.DetalleAudioResponse.prototype.setTitulo = function(v) { return jspb.Message.setProto3StringField(this, 3, v); };
proto.servicios.DetalleAudioResponse.prototype.getMetadatosList = function() { return jspb.Message.getRepeatedWrapperField(this, proto.servicios.CampoMetadata, 4); };
proto.servicios.DetalleAudioResponse.prototype.serializeBinary = function() {
  var w = new jspb.BinaryWriter(); proto.servicios.DetalleAudioResponse.serializeBinaryToWriter(this, w); return w.getResultBuffer();
};
proto.servicios.DetalleAudioResponse.serializeBinaryToWriter = function(m, w) {
  var f = m.getId(); if (f !== 0) w.writeInt32(1, f);
  var f = m.getIdtipo(); if (f !== 0) w.writeInt32(2, f);
  var f = m.getTitulo(); if (f.length > 0) w.writeString(3, f);
  var f = m.getMetadatosList(); if (f.length > 0) w.writeRepeatedMessage(4, f, proto.servicios.CampoMetadata.serializeBinaryToWriter);
};
proto.servicios.DetalleAudioResponse.deserializeBinary = function(b) {
  var r = new jspb.BinaryReader(b); var m = new proto.servicios.DetalleAudioResponse(); return proto.servicios.DetalleAudioResponse.deserializeBinaryFromReader(m, r);
};
proto.servicios.DetalleAudioResponse.deserializeBinaryFromReader = function(m, r) {
  while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); switch(f) { case 1: m.setId(r.readInt32()); break; case 2: m.setIdtipo(r.readInt32()); break; case 3: m.setTitulo(r.readString()); break; case 4: { var v = new proto.servicios.CampoMetadata(); r.readMessage(v, proto.servicios.CampoMetadata.deserializeBinaryFromReader); m.addMetadatos(v); break; } default: r.skipField(); } } return m;
};
proto.servicios.DetalleAudioResponse.prototype.addMetadatos = function(v, opt_i) { return jspb.Message.addToRepeatedWrapperField(this, 4, v, proto.servicios.CampoMetadata, opt_i); };
proto.servicios.DetalleAudioResponse.prototype.toObject = function() { return { id: this.getId(), idtipo: this.getIdtipo(), titulo: this.getTitulo(), metadatosList: this.getMetadatosList().map(function(i){return i.toObject();}) }; };

// ---- LoginRequest ----
proto.servicios.LoginRequest = function(opt_data) { jspb.Message.initialize(this, opt_data, 0, -1, null, null); };
goog.inherits(proto.servicios.LoginRequest, jspb.Message);
proto.servicios.LoginRequest.prototype.getNickname = function() { return jspb.Message.getFieldWithDefault(this, 1, ''); };
proto.servicios.LoginRequest.prototype.setNickname = function(v) { return jspb.Message.setProto3StringField(this, 1, v); };
proto.servicios.LoginRequest.prototype.getPassword = function() { return jspb.Message.getFieldWithDefault(this, 2, ''); };
proto.servicios.LoginRequest.prototype.setPassword = function(v) { return jspb.Message.setProto3StringField(this, 2, v); };
proto.servicios.LoginRequest.prototype.serializeBinary = function() { var w = new jspb.BinaryWriter(); proto.servicios.LoginRequest.serializeBinaryToWriter(this, w); return w.getResultBuffer(); };
proto.servicios.LoginRequest.serializeBinaryToWriter = function(m, w) { var f = m.getNickname(); if (f.length > 0) w.writeString(1, f); var f = m.getPassword(); if (f.length > 0) w.writeString(2, f); };
proto.servicios.LoginRequest.deserializeBinary = function(b) { var r = new jspb.BinaryReader(b); var m = new proto.servicios.LoginRequest(); return proto.servicios.LoginRequest.deserializeBinaryFromReader(m, r); };
proto.servicios.LoginRequest.deserializeBinaryFromReader = function(m, r) { while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); switch(f) { case 1: m.setNickname(r.readString()); break; case 2: m.setPassword(r.readString()); break; default: r.skipField(); } } return m; };
proto.servicios.LoginRequest.prototype.toObject = function() { return { nickname: this.getNickname(), password: this.getPassword() }; };

// ---- LoginResponse ----
proto.servicios.LoginResponse = function(opt_data) { jspb.Message.initialize(this, opt_data, 0, -1, null, null); };
goog.inherits(proto.servicios.LoginResponse, jspb.Message);
proto.servicios.LoginResponse.prototype.getExitoso = function() { return jspb.Message.getBooleanFieldWithDefault(this, 1, false); };
proto.servicios.LoginResponse.prototype.setExitoso = function(v) { return jspb.Message.setProto3BooleanField(this, 1, v); };
proto.servicios.LoginResponse.prototype.getMensaje = function() { return jspb.Message.getFieldWithDefault(this, 2, ''); };
proto.servicios.LoginResponse.prototype.setMensaje = function(v) { return jspb.Message.setProto3StringField(this, 2, v); };
proto.servicios.LoginResponse.prototype.serializeBinary = function() { var w = new jspb.BinaryWriter(); proto.servicios.LoginResponse.serializeBinaryToWriter(this, w); return w.getResultBuffer(); };
proto.servicios.LoginResponse.serializeBinaryToWriter = function(m, w) { var f = m.getExitoso(); if (f) w.writeBool(1, f); var f = m.getMensaje(); if (f.length > 0) w.writeString(2, f); };
proto.servicios.LoginResponse.deserializeBinary = function(b) { var r = new jspb.BinaryReader(b); var m = new proto.servicios.LoginResponse(); return proto.servicios.LoginResponse.deserializeBinaryFromReader(m, r); };
proto.servicios.LoginResponse.deserializeBinaryFromReader = function(m, r) { while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); switch(f) { case 1: m.setExitoso(r.readBool()); break; case 2: m.setMensaje(r.readString()); break; default: r.skipField(); } } return m; };
proto.servicios.LoginResponse.prototype.toObject = function() { return { exitoso: this.getExitoso(), mensaje: this.getMensaje() }; };

// ---- RegistrarUsuarioRequest ----
proto.servicios.RegistrarUsuarioRequest = function(opt_data) { jspb.Message.initialize(this, opt_data, 0, -1, null, null); };
goog.inherits(proto.servicios.RegistrarUsuarioRequest, jspb.Message);
proto.servicios.RegistrarUsuarioRequest.prototype.getNickname = function() { return jspb.Message.getFieldWithDefault(this, 1, ''); };
proto.servicios.RegistrarUsuarioRequest.prototype.setNickname = function(v) { return jspb.Message.setProto3StringField(this, 1, v); };
proto.servicios.RegistrarUsuarioRequest.prototype.getPassword = function() { return jspb.Message.getFieldWithDefault(this, 2, ''); };
proto.servicios.RegistrarUsuarioRequest.prototype.setPassword = function(v) { return jspb.Message.setProto3StringField(this, 2, v); };
proto.servicios.RegistrarUsuarioRequest.prototype.serializeBinary = function() { var w = new jspb.BinaryWriter(); proto.servicios.RegistrarUsuarioRequest.serializeBinaryToWriter(this, w); return w.getResultBuffer(); };
proto.servicios.RegistrarUsuarioRequest.serializeBinaryToWriter = function(m, w) { var f = m.getNickname(); if (f.length > 0) w.writeString(1, f); var f = m.getPassword(); if (f.length > 0) w.writeString(2, f); };
proto.servicios.RegistrarUsuarioRequest.deserializeBinary = function(b) { var r = new jspb.BinaryReader(b); var m = new proto.servicios.RegistrarUsuarioRequest(); return proto.servicios.RegistrarUsuarioRequest.deserializeBinaryFromReader(m, r); };
proto.servicios.RegistrarUsuarioRequest.deserializeBinaryFromReader = function(m, r) { while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); switch(f) { case 1: m.setNickname(r.readString()); break; case 2: m.setPassword(r.readString()); break; default: r.skipField(); } } return m; };
proto.servicios.RegistrarUsuarioRequest.prototype.toObject = function() { return { nickname: this.getNickname(), password: this.getPassword() }; };

// ---- RegistrarUsuarioResponse ----
proto.servicios.RegistrarUsuarioResponse = function(opt_data) { jspb.Message.initialize(this, opt_data, 0, -1, null, null); };
goog.inherits(proto.servicios.RegistrarUsuarioResponse, jspb.Message);
proto.servicios.RegistrarUsuarioResponse.prototype.getRegistrado = function() { return jspb.Message.getBooleanFieldWithDefault(this, 1, false); };
proto.servicios.RegistrarUsuarioResponse.prototype.setRegistrado = function(v) { return jspb.Message.setProto3BooleanField(this, 1, v); };
proto.servicios.RegistrarUsuarioResponse.prototype.getMensaje = function() { return jspb.Message.getFieldWithDefault(this, 2, ''); };
proto.servicios.RegistrarUsuarioResponse.prototype.setMensaje = function(v) { return jspb.Message.setProto3StringField(this, 2, v); };
proto.servicios.RegistrarUsuarioResponse.prototype.serializeBinary = function() { var w = new jspb.BinaryWriter(); proto.servicios.RegistrarUsuarioResponse.serializeBinaryToWriter(this, w); return w.getResultBuffer(); };
proto.servicios.RegistrarUsuarioResponse.serializeBinaryToWriter = function(m, w) { var f = m.getRegistrado(); if (f) w.writeBool(1, f); var f = m.getMensaje(); if (f.length > 0) w.writeString(2, f); };
proto.servicios.RegistrarUsuarioResponse.deserializeBinary = function(b) { var r = new jspb.BinaryReader(b); var m = new proto.servicios.RegistrarUsuarioResponse(); return proto.servicios.RegistrarUsuarioResponse.deserializeBinaryFromReader(m, r); };
proto.servicios.RegistrarUsuarioResponse.deserializeBinaryFromReader = function(m, r) { while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); switch(f) { case 1: m.setRegistrado(r.readBool()); break; case 2: m.setMensaje(r.readString()); break; default: r.skipField(); } } return m; };
proto.servicios.RegistrarUsuarioResponse.prototype.toObject = function() { return { registrado: this.getRegistrado(), mensaje: this.getMensaje() }; };

// ---- AlmacenarAudioRequest ----
proto.servicios.AlmacenarAudioRequest = function(opt_data) { jspb.Message.initialize(this, opt_data, 0, -1, null, null); };
goog.inherits(proto.servicios.AlmacenarAudioRequest, jspb.Message);
proto.servicios.AlmacenarAudioRequest.prototype.getTitulo = function() { return jspb.Message.getFieldWithDefault(this, 1, ''); };
proto.servicios.AlmacenarAudioRequest.prototype.setTitulo = function(v) { return jspb.Message.setProto3StringField(this, 1, v); };
proto.servicios.AlmacenarAudioRequest.prototype.getArtista = function() { return jspb.Message.getFieldWithDefault(this, 2, ''); };
proto.servicios.AlmacenarAudioRequest.prototype.setArtista = function(v) { return jspb.Message.setProto3StringField(this, 2, v); };
proto.servicios.AlmacenarAudioRequest.prototype.getGenero = function() { return jspb.Message.getFieldWithDefault(this, 3, ''); };
proto.servicios.AlmacenarAudioRequest.prototype.setGenero = function(v) { return jspb.Message.setProto3StringField(this, 3, v); };
proto.servicios.AlmacenarAudioRequest.prototype.getAlbum = function() { return jspb.Message.getFieldWithDefault(this, 4, ''); };
proto.servicios.AlmacenarAudioRequest.prototype.setAlbum = function(v) { return jspb.Message.setProto3StringField(this, 4, v); };
proto.servicios.AlmacenarAudioRequest.prototype.getAnio = function() { return jspb.Message.getFieldWithDefault(this, 5, ''); };
proto.servicios.AlmacenarAudioRequest.prototype.setAnio = function(v) { return jspb.Message.setProto3StringField(this, 5, v); };
proto.servicios.AlmacenarAudioRequest.prototype.getIdtipo = function() { return jspb.Message.getFieldWithDefault(this, 6, 0); };
proto.servicios.AlmacenarAudioRequest.prototype.setIdtipo = function(v) { return jspb.Message.setProto3IntField(this, 6, v); };
proto.servicios.AlmacenarAudioRequest.prototype.serializeBinary = function() { var w = new jspb.BinaryWriter(); proto.servicios.AlmacenarAudioRequest.serializeBinaryToWriter(this, w); return w.getResultBuffer(); };
proto.servicios.AlmacenarAudioRequest.serializeBinaryToWriter = function(m, w) {
  var f = m.getTitulo(); if (f.length > 0) w.writeString(1, f);
  var f = m.getArtista(); if (f.length > 0) w.writeString(2, f);
  var f = m.getGenero(); if (f.length > 0) w.writeString(3, f);
  var f = m.getAlbum(); if (f.length > 0) w.writeString(4, f);
  var f = m.getAnio(); if (f.length > 0) w.writeString(5, f);
  var f = m.getIdtipo(); if (f !== 0) w.writeInt32(6, f);
};
proto.servicios.AlmacenarAudioRequest.deserializeBinary = function(b) { var r = new jspb.BinaryReader(b); var m = new proto.servicios.AlmacenarAudioRequest(); return proto.servicios.AlmacenarAudioRequest.deserializeBinaryFromReader(m, r); };
proto.servicios.AlmacenarAudioRequest.deserializeBinaryFromReader = function(m, r) { while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); switch(f) { case 1: m.setTitulo(r.readString()); break; case 2: m.setArtista(r.readString()); break; case 3: m.setGenero(r.readString()); break; case 4: m.setAlbum(r.readString()); break; case 5: m.setAnio(r.readString()); break; case 6: m.setIdtipo(r.readInt32()); break; default: r.skipField(); } } return m; };
proto.servicios.AlmacenarAudioRequest.prototype.toObject = function() { return { titulo: this.getTitulo(), artista: this.getArtista(), genero: this.getGenero(), album: this.getAlbum(), anio: this.getAnio(), idtipo: this.getIdtipo() }; };

// ---- AlmacenarAudioResponse ----
proto.servicios.AlmacenarAudioResponse = function(opt_data) { jspb.Message.initialize(this, opt_data, 0, -1, null, null); };
goog.inherits(proto.servicios.AlmacenarAudioResponse, jspb.Message);
proto.servicios.AlmacenarAudioResponse.prototype.getGuardado = function() { return jspb.Message.getBooleanFieldWithDefault(this, 1, false); };
proto.servicios.AlmacenarAudioResponse.prototype.setGuardado = function(v) { return jspb.Message.setProto3BooleanField(this, 1, v); };
proto.servicios.AlmacenarAudioResponse.prototype.getIdgenerado = function() { return jspb.Message.getFieldWithDefault(this, 2, 0); };
proto.servicios.AlmacenarAudioResponse.prototype.setIdgenerado = function(v) { return jspb.Message.setProto3IntField(this, 2, v); };
proto.servicios.AlmacenarAudioResponse.prototype.getFechahoraregistro = function() { return jspb.Message.getFieldWithDefault(this, 3, ''); };
proto.servicios.AlmacenarAudioResponse.prototype.setFechahoraregistro = function(v) { return jspb.Message.setProto3StringField(this, 3, v); };
proto.servicios.AlmacenarAudioResponse.prototype.serializeBinary = function() { var w = new jspb.BinaryWriter(); proto.servicios.AlmacenarAudioResponse.serializeBinaryToWriter(this, w); return w.getResultBuffer(); };
proto.servicios.AlmacenarAudioResponse.serializeBinaryToWriter = function(m, w) { var f = m.getGuardado(); if (f) w.writeBool(1, f); var f = m.getIdgenerado(); if (f !== 0) w.writeInt32(2, f); var f = m.getFechahoraregistro(); if (f.length > 0) w.writeString(3, f); };
proto.servicios.AlmacenarAudioResponse.deserializeBinary = function(b) { var r = new jspb.BinaryReader(b); var m = new proto.servicios.AlmacenarAudioResponse(); return proto.servicios.AlmacenarAudioResponse.deserializeBinaryFromReader(m, r); };
proto.servicios.AlmacenarAudioResponse.deserializeBinaryFromReader = function(m, r) { while (r.nextField()) { if (r.isEndGroup()) break; var f = r.getFieldNumber(); switch(f) { case 1: m.setGuardado(r.readBool()); break; case 2: m.setIdgenerado(r.readInt32()); break; case 3: m.setFechahoraregistro(r.readString()); break; default: r.skipField(); } } return m; };
proto.servicios.AlmacenarAudioResponse.prototype.toObject = function() { return { guardado: this.getGuardado(), idgenerado: this.getIdgenerado(), fechahoraregistro: this.getFechahoraregistro() }; };

module.exports = proto.servicios;
