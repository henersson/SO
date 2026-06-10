// funciones.js - wrapper simple para llamar a la implementación del streaming

function pedirCancion(titulo, formato) {
    if (typeof window.iniciar_streaming_cancion === 'function') {
        return window.iniciar_streaming_cancion(titulo, formato);
    }
    if (typeof window.iniciar_streaming_cancion_impl === 'function') {
        return window.iniciar_streaming_cancion_impl(titulo, formato);
    }
    if (typeof window.iniciarStreamGRPCImpl === 'function') {
        return window.iniciarStreamGRPCImpl(titulo, formato);
    }
    if (typeof window.iniciarStreamGRPC === 'function') {
        return window.iniciarStreamGRPC(titulo, formato);
    }

    console.error('No se encontró ninguna implementación de iniciar_streaming_cancion.');
    const d = document.getElementById('log');
    if (d) {
        const p = document.createElement('div');
        p.className = 'error';
        p.textContent = 'No se encontró ninguna implementación de iniciar_streaming_cancion.';
        d.appendChild(p);
    }
}

// Export para compatibilidad con cargas como módulo
if (typeof module !== 'undefined' && module.exports) {
    module.exports = { pedirCancion };
}

// --- Helpers y listeners para logging de audio ---
(function () {
    function writeLog(message, level) {
        const d = document.getElementById('log');
        if (!d) return;
        const p = document.createElement('div');
        p.className = level || '';
        const ts = new Date().toLocaleTimeString();
        p.textContent = `[${ts}] ${message}`;
        d.appendChild(p);
        d.scrollTop = d.scrollHeight;
    }

    function attachAudioListeners() {
        const audio = document.getElementById('audio-player');
        if (!audio) {
            writeLog('No se encontró el elemento audio#audio-player.', 'error');
            return;
        }

        // CAMBIO: usar 'playing' en lugar de 'play'
        // 'playing' se dispara cuando el audio realmente empieza a sonar (ya cargó datos)
        // 'play' se dispara apenas se llama a .play(), antes de que haya audio
        audio.addEventListener('playing', function () {
            writeLog('Reproducción iniciada (playing).', 'success');
            if (typeof window !== 'undefined' && typeof window.unirseAlCanal === 'function' && window.audioIdActual && window.nicknameActual) {
                window.unirseAlCanal(window.audioIdActual);
            }
        });

        audio.addEventListener('pause', function () {
            writeLog('Reproducción pausada (pause).', 'error');
            if (typeof window !== 'undefined' && typeof window.salirDelCanal === 'function' && window.audioIdActual && window.nicknameActual) {
                window.salirDelCanal(window.audioIdActual);
            }
        });
    }

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', attachAudioListeners);
    } else {
        attachAudioListeners();
    }

    if (typeof window !== 'undefined') {
        window.__writeAudioLog = writeLog;
        window.__attachAudioListeners = attachAudioListeners;
    }

    function conectarWebSocket() {
        if (typeof WebSocket === 'undefined') return;
        if (typeof window !== 'undefined' && window._ws && (window._ws.readyState === WebSocket.OPEN || window._ws.readyState === WebSocket.CONNECTING)) {
            return;
        }

        const ws = new WebSocket('ws://localhost:3001');
        if (typeof window !== 'undefined') {
            window._ws = ws;
        }

        ws.onopen = () => console.log('[WS] Conectado al servidor de reacciones :3001');

        ws.onmessage = (event) => {
            try {
                const msg = JSON.parse(event.data);
                if (typeof window !== 'undefined' && typeof window.manejarMensajeServidor === 'function') {
                    window.manejarMensajeServidor(msg);
                }
            } catch (error) {
                console.error('[WS] Error parseando mensaje:', error);
            }
        };

        ws.onclose = () => {
            console.log('[WS] Desconectado. Reintentando en 3s...');
            setTimeout(conectarWebSocket, 3000);
        };

        ws.onerror = (err) => console.error('[WS] Error:', err);
    }

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', conectarWebSocket);
    } else {
        conectarWebSocket();
    }

    if (typeof window !== 'undefined') {
        window.conectarWebSocket = conectarWebSocket;
        window.writeAudioLog = writeLog;
    }
})();