/**
 * @file video.h
 * @brief Definiciones para manejo de video en modo texto
 */

#ifndef VIDEO_H
#define VIDEO_H

/* Dirección base de la memoria de video en modo texto */
#define VIDEO_MEMORY 0xB8000

/* Dimensiones de la pantalla en modo texto */
#define SCREEN_WIDTH 80
#define SCREEN_HEIGHT 25

/* Códigos de color para el atributo de video */
#define COLOR_BLACK 0x0
#define COLOR_BLUE 0x1
#define COLOR_GREEN 0x2
#define COLOR_CYAN 0x3
#define COLOR_RED 0x4
#define COLOR_MAGENTA 0x5
#define COLOR_BROWN 0x6
#define COLOR_LIGHT_GRAY 0x7
#define COLOR_DARK_GRAY 0x8
#define COLOR_LIGHT_BLUE 0x9
#define COLOR_LIGHT_GREEN 0xA
#define COLOR_LIGHT_CYAN 0xB
#define COLOR_LIGHT_RED 0xC
#define COLOR_LIGHT_MAGENTA 0xD
#define COLOR_YELLOW 0xE
#define COLOR_WHITE 0xF

/* Macro para crear un atributo de color (fondo y frente) */
#define VIDEO_ATTR(fg, bg) ((bg << 4) | (fg & 0x0F))

/* Prototipos de funciones */

/**
 * @brief Limpia la pantalla con un color de fondo específico
 * @param attr Atributo de color (byte más significativo)
 */
void clear_screen(unsigned char attr);

/**
 * @brief Imprime un string en una posición específica de la pantalla
 * @param str String a imprimir (terminado en null)
 * @param x Columna (0-79)
 * @param y Fila (0-24)
 * @param attr Atributo de color
 */
void print_at(const char *str, int x, int y, unsigned char attr);

/**
 * @brief Imprime un string centrado en una fila específica
 * @param str String a imprimir
 * @param y Fila donde centrar el texto
 * @param attr Atributo de color
 */
void print_centered(const char *str, int y, unsigned char attr);

/**
 * @brief Calcula la longitud de un string
 * @param str String
 * @return Longitud del string
 */
int strlen(const char *str);

#endif /* VIDEO_H */
