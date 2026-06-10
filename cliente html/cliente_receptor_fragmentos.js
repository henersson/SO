// cliente_receptor_fragmentos.js
const proto = require('./stubs_generados/servicios_pb.js');
const { StreamingServiceClient } = require('./stubs_generados/servicios_grpc_web_pb.js');

const ENVOY_HOST = 'http://localhost:8080';
const client = new StreamingServiceClient(ENVOY_HOST);

let audioChunks = [];
const audioPlayer = document.getElementById('audio-player');
let _firstFragmentReceived = false;

function iniciar_streaming_cancion(idAudio) {
    console.log('Iniciando streaming para idAudio:', idAudio);

    audioChunks = [];
    _firstFragmentReceived = false;

    const request = new proto.ReproducirAudioRequest();
    request.setIdaudio(idAudio);

    console.log('Solicitud preparada:', request.toObject());

    const stream = client.reproducirAudio(request, {});

    stream.on('data', function(response) {
        const data = response.getData_asU8();
        console.log('Fragmento recibido. Tamaño: ' + data.length + ' bytes');
        audioChunks.push(data);

        if (!_firstFragmentReceived) {
            _firstFragmentReceived = true;
            if (audioPlayer) {
                audioPlayer.style.display = '';
            }
            if (typeof window !== 'undefined' && window.__writeAudioLog) {
                window.__writeAudioLog('Primer fragmento recibido. Reproductor mostrado.', 'success');
            }
        }
    });

    stream.on('end', function() {
        console.log('Transmisión finalizada. Preparando reproducción...');
        const audioBlob = new Blob(audioChunks, { type: 'audio/mpeg' });
        const audioUrl = URL.createObjectURL(audioBlob);

        if (audioPlayer) {
            audioPlayer.src = audioUrl;
            audioPlayer.play()
                .then(function() { console.log('Reproducción iniciada.'); })
                .catch(function(e) { console.error('Error al reproducir:', e); });
        }

        audioChunks = [];
        _firstFragmentReceived = false;
    });

    stream.on('error', function(err) {
        console.error('ERROR gRPC-Web:', err);
        if (typeof window.__writeAudioLog === 'function') {
            window.__writeAudioLog('Error de streaming: ' + err.message, 'error');
        }
    });
}

if (typeof window !== 'undefined') {
    window.iniciar_streaming_cancion = iniciar_streaming_cancion;
    window.iniciarStreamGRPCImpl = iniciar_streaming_cancion;
    console.log('iniciar_streaming_cancion exportado a window (acepta idAudio numérico)');
}
