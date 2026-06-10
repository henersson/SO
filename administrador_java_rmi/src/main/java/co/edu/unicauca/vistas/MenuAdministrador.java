package co.edu.unicauca.vistas;

import co.edu.unicauca.DTOs.AudioAlmacenarDTO;
import co.edu.unicauca.DTOs.AudioAlmacenarResponseDTO;
import co.edu.unicauca.capaFachada.FachadaAdministrador;

import java.io.File;
import java.util.Scanner;

public class MenuAdministrador {

    private final FachadaAdministrador fachada;
    private final Scanner scanner = new Scanner(System.in);

    public MenuAdministrador(FachadaAdministrador fachada) {
        this.fachada = fachada;
    }

    public void mostrar() {
        boolean continuar = true;
        while (continuar) {
            System.out.println("\n===== Menú Administrador =====");
            System.out.println("1. Registrar nuevo audio");
            System.out.println("2. Listar audios");
            System.out.println("0. Salir");
            System.out.print("Opción: ");

            String opcion;
            try {
                if (!scanner.hasNextLine()) {
                    System.out.println("\n[Menu] Entrada estándar cerrada. Saliendo...");
                    break;
                }
                opcion = scanner.nextLine().trim();
            } catch (java.util.NoSuchElementException e) {
                System.out.println("\n[Menu] No hay más entrada. Saliendo...");
                break;
            }
            switch (opcion) {
                case "1" -> registrarAudio();
                case "2" -> listarAudios();
                case "0" -> { continuar = false; System.out.println("Cerrando administrador..."); }
                default  -> System.out.println("Opción inválida. Ingrese 1 o 0.");
            }
        }
    }

    private void listarAudios() {
        System.out.println("\n--- Listado de audios (ServidorMetadatos) ---");
        String json = fachada.listarAudios();
        System.out.println(json);
    }

    private void registrarAudio() {
        System.out.println("\n--- Registrar Nuevo Audio ---");
        System.out.print("Título  : "); String titulo  = scanner.nextLine().trim();
        System.out.print("Artista : "); String artista = scanner.nextLine().trim();
        System.out.print("Género  : "); String genero  = scanner.nextLine().trim();
        System.out.print("Álbum   : "); String album   = scanner.nextLine().trim();
        System.out.print("Año     : "); String anio    = scanner.nextLine().trim();
        System.out.println("Tipo de audio:");
        System.out.println("  1 = Música  |  2 = Podcast  |  3 = Audiolibro  |  4 = Ruido Blanco");
        System.out.print("ID Tipo : ");
        int idTipo;
        try {
            idTipo = Integer.parseInt(scanner.nextLine().trim());
        } catch (NumberFormatException e) {
            System.out.println("ID de tipo inválido. Operación cancelada.");
            return;
        }

        System.out.print("Ruta del archivo de audio (mp3): ");
        String rutaArchivo = scanner.nextLine().trim();
        if (rutaArchivo.isEmpty() || !new File(rutaArchivo).exists()) {
            System.out.println("Ruta inválida o archivo no existe. Operación cancelada.");
            return;
        }

        AudioAlmacenarDTO dto = new AudioAlmacenarDTO(titulo, artista, genero, album, anio, idTipo);
        AudioAlmacenarResponseDTO respuesta = fachada.registrarNuevoAudio(dto);

        if (respuesta == null || !respuesta.isGuardado()) {
            System.out.println("✗ Error al registrar el audio. Revise que ServidorMetadatos esté corriendo.");
            return;
        }

        boolean subido = fachada.subirArchivoAudio(respuesta.getIdGenerado(), rutaArchivo);
        if (subido) {
            System.out.println("✓ Audio registrado y archivo subido exitosamente.");
            System.out.println("  El ServidorCorreos recibirá la notificación por RabbitMQ.");
        } else {
            System.out.println("✗ Audio registrado, pero falló la subida del archivo.");
            System.out.println("  Verifique que ServidorStreaming esté corriendo en :8083.");
        }
    }
}
