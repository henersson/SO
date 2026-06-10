package co.edu.unicauca.DTOs;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class AudioAlmacenarResponseDTO {

    @JsonProperty("guardado")
    private boolean guardado;

    @JsonProperty("idGenerado")
    private int idGenerado;

    @JsonProperty("fechaHoraRegistro")
    private String fechaHoraRegistro;
}
