package co.edu.unicauca.DTOs;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.io.Serializable;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class NotificacionReproduccionDTO implements Serializable {

    private static final long serialVersionUID = 1L;

    @JsonProperty("fechaHoraReproduccion")
    private String fechaHoraReproduccion;

    @JsonProperty("idAudio")
    private int idAudio;
}
