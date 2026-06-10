package co.edu.unicauca.DTOs;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class AudioAlmacenarDTO {

    @JsonProperty("titulo")
    private String titulo;

    @JsonProperty("artista")
    private String artista;

    @JsonProperty("genero")
    private String genero;

    @JsonProperty("album")
    private String album;

    @JsonProperty("anio")
    private String anio;

    @JsonProperty("idTipo")
    private int idTipo;
}
