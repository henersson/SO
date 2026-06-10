package co.edu.unicauca.capaFachada;

import co.edu.unicauca.DTOs.AudioAlmacenarDTO;
import co.edu.unicauca.DTOs.AudioAlmacenarResponseDTO;
import co.edu.unicauca.capaDeControladores.ControladorAdminRest;

public class FachadaAdministrador {

    private final ControladorAdminRest controladorRest;

    public FachadaAdministrador(ControladorAdminRest controladorRest) {
        this.controladorRest = controladorRest;
    }

    public AudioAlmacenarResponseDTO registrarNuevoAudio(AudioAlmacenarDTO dto) {
        return controladorRest.almacenarAudio(dto);
    }

    public boolean subirArchivoAudio(int idAudio, String rutaArchivo) {
        return controladorRest.subirArchivoAudio(idAudio, rutaArchivo);
    }

    public boolean registrarParaNotificaciones(String host, int puerto) {
        return controladorRest.registrarCallbackEnServidorStreaming(host, puerto);
    }
    public String listarAudios() {
        return controladorRest.listarAudios();
    }

}

