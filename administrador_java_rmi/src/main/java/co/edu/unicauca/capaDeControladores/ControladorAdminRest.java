package co.edu.unicauca.capaDeControladores;

import com.fasterxml.jackson.databind.ObjectMapper;
import co.edu.unicauca.DTOs.AudioAlmacenarDTO;
import co.edu.unicauca.DTOs.AudioAlmacenarResponseDTO;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Base64;
import java.util.HashMap;
import java.util.Map;

public class ControladorAdminRest {

    private static final String URL_ALMACENAR_AUDIO    = "http://localhost:8082/audios/almacenar";
    private static final String URL_REGISTRAR_CALLBACK = "http://localhost:8083/callbacks/registrar";
    private static final String URL_SUBIR_AUDIO        = "http://localhost:8083/audios/subir";

    private final HttpClient httpClient = HttpClient.newHttpClient();
    private final ObjectMapper mapper    = new ObjectMapper();

    public AudioAlmacenarResponseDTO almacenarAudio(AudioAlmacenarDTO dto) {
        try {
            String cuerpoJson = mapper.writeValueAsString(dto);
            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create(URL_ALMACENAR_AUDIO))
                    .header("Content-Type", "application/json")
                    .POST(HttpRequest.BodyPublishers.ofString(cuerpoJson))
                    .build();
            HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
            System.out.println("[REST] Respuesta ServidorMetadatos: " + response.body());
            if (response.statusCode() != 200) {
                return null;
            }
            return mapper.readValue(response.body(), AudioAlmacenarResponseDTO.class);
        } catch (Exception e) {
            System.err.println("[REST] Error al almacenar audio: " + e.getMessage());
            return null;
        }
    }

    public boolean subirArchivoAudio(int idAudio, String rutaArchivo) {
        try {
            Path ruta = Paths.get(rutaArchivo);
            byte[] datos = Files.readAllBytes(ruta);
            String base64 = Base64.getEncoder().encodeToString(datos);

            Map<String, Object> payload = new HashMap<>();
            payload.put("idAudio", idAudio);
            payload.put("fileName", ruta.getFileName().toString());
            payload.put("dataBase64", base64);

            String cuerpoJson = mapper.writeValueAsString(payload);
            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create(URL_SUBIR_AUDIO))
                    .header("Content-Type", "application/json")
                    .POST(HttpRequest.BodyPublishers.ofString(cuerpoJson))
                    .build();
            HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
            System.out.println("[REST] Respuesta ServidorStreaming: " + response.body());
            return response.statusCode() == 200;
        } catch (Exception e) {
            System.err.println("[REST] Error al subir audio: " + e.getMessage());
            System.err.println("       Verifique que ServidorStreaming esté corriendo en :8083");
            return false;
        }
    }

    public boolean registrarCallbackEnServidorStreaming(String host, int puerto) {
        try {
            String cuerpoJson = String.format("{\"host\":\"%s\",\"puerto\":%d}", host, puerto);
            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create(URL_REGISTRAR_CALLBACK))
                    .header("Content-Type", "application/json")
                    .POST(HttpRequest.BodyPublishers.ofString(cuerpoJson))
                    .build();
            HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
            System.out.println("[REST] Registro callback: " + response.body());
            return response.statusCode() == 200;
        } catch (Exception e) {
            System.err.println("[REST] Error al registrar callback: " + e.getMessage());
            System.err.println("       Verifique que ServidorStreaming esté corriendo en :8083");
            return false;
        }
    }

    public String listarAudios() {
        try {
            HttpRequest request = HttpRequest.newBuilder()
                    .uri(URI.create("http://localhost:8082/audios"))
                    .GET()
                    .build();
            HttpResponse<String> response = httpClient.send(request, HttpResponse.BodyHandlers.ofString());
            if (response.statusCode() == 200) {
                return response.body();
            }
            return "[]";
        } catch (Exception e) {
            System.err.println("[REST] Error al listar audios: " + e.getMessage());
            return "[]";
        }
    }
}
