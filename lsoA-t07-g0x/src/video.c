/**
 * @file video.c
 * @brief Implementación de funciones para manejo de video en modo texto
 */

#include "video.h"

/**
 * @brief Calcula la longitud de un string
 */
int strlen(const char *str) {
    int len = 0;
    while (str[len] != '\0') {
        len++;
    }
    return len;
}

/**
 * @brief Limpia la pantalla con un color específico
 */
void clear_screen(unsigned char attr) {
    unsigned char *video_mem = (unsigned char *)VIDEO_MEMORY;
    int i;
    
    /* Recorrer toda la pantalla (80x25 = 2000 caracteres) */
    for (i = 0; i < SCREEN_WIDTH * SCREEN_HEIGHT * 2; i += 2) {
        video_mem[i] = ' ';        /* Caracter: espacio */
        video_mem[i + 1] = attr;   /* Atributo de color */
    }
}

/**
 * @brief Imprime un string en una posición específica
 */
void print_at(const char *str, int x, int y, unsigned char attr) {
    unsigned char *video_mem = (unsigned char *)VIDEO_MEMORY;
    int offset = (y * SCREEN_WIDTH + x) * 2;
    int i = 0;
    
    /* Escribir cada caracter del string */
    while (str[i] != '\0' && x < SCREEN_WIDTH) {
        video_mem[offset] = str[i];         /* Caracter ASCII */
        video_mem[offset + 1] = attr;       /* Atributo de color */
        offset += 2;
        i++;
        x++;
    }
}

/**
 * @brief Imprime un string centrado en una fila
 */
void print_centered(const char *str, int y, unsigned char attr) {
    int len = strlen(str);
    int x = (SCREEN_WIDTH - len) / 2;  /* Calcular posición centrada */
    
    if (x < 0) x = 0;  /* Asegurar que no sea negativo */
    
    print_at(str, x, y, attr);
}