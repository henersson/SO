package co.edu.unicauca.capaDeControladores;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpServer;
import co.edu.unicauca.DTOs.NotificacionReproduccionDTO;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.net.InetSocketAddress;

public class ServidorHttpCallback {

    private final int puertoConfigurado;
    private int puertoReal;
    private HttpServer servidor;
    private final ObjectMapper mapper = new ObjectMapper();

    public ServidorHttpCallback(int puerto) {
        this.puertoConfigurado = puerto;
        this.puertoReal = puerto;
    }

    public int iniciar() throws IOException {
        servidor = HttpServer.create(new InetSocketAddress(puertoConfigurado), 10);
        servidor.createContext("/callback/reproduccion", this::procesarCallback);
        servidor.start();
        puertoReal = servidor.getAddress().getPort();
        System.out.println("[Callback] Servidor HTTP iniciado en puerto " + puertoReal);
            System.out.println("[Callback] Esperando notificaciones de reproducción...");
        return puertoReal;
    }

    private void procesarCallback(HttpExchange intercambio) throws IOException {
        if (!"POST".equals(intercambio.getRequestMethod())) {
            intercambio.sendResponseHeaders(405, -1);
            return;
        }
        try (InputStream cuerpo = intercambio.getRequestBody()) {
            NotificacionReproduccionDTO notificacion = mapper.readValue(cuerpo, NotificacionReproduccionDTO.class);
            mostrarNotificacion(notificacion);
        } catch (Exception e) {
            System.err.println("[Callback] Error procesando notificación: " + e.getMessage());
        }
        byte[] respuesta = "OK".getBytes();
        intercambio.sendResponseHeaders(200, respuesta.length);
        try (OutputStream os = intercambio.getResponseBody()) {
            os.write(respuesta);
        }
    }

    private void mostrarNotificacion(NotificacionReproduccionDTO notificacion) {
        System.out.println();
        System.out.println("╔══════════════════════════════════════════════╗");
        System.out.println("║   🔔  CALLBACK — AUDIO REPRODUCIDO           ║");
        System.out.println("╠══════════════════════════════════════════════╣");
        System.out.printf( "║  ID Audio    : %-30d║%n", notificacion.getIdAudio());
        System.out.printf( "║  Fecha/Hora  : %-30s║%n", notificacion.getFechaHoraReproduccion());
        System.out.println("╚══════════════════════════════════════════════╝");
        System.out.println();
    }
}
