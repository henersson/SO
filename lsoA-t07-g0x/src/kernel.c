/**
 * @file kernel.c
 * @brief Punto de entrada del kernel en modo protegido
 * @author Henersson Cobo
 */

#include "video.h"

/**
 * @brief Función principal del kernel
 * Esta función se ejecuta después de entrar en modo protegido
 * @param kernel_size Tamaño del kernel en sectores
 * @param message Mensaje del bootloader
 */
void cmain(int kernel_size, char *message) {
    unsigned char attr_background;
    unsigned char attr_title;
    unsigned char attr_text;
    unsigned char attr_info;
    
    /* Definir atributos de color */
    /* Fondo azul con texto blanco para el título */
    attr_title = VIDEO_ATTR(COLOR_WHITE, COLOR_BLUE);
    
    /* Fondo negro con texto verde claro para el texto normal */
    attr_text = VIDEO_ATTR(COLOR_LIGHT_GREEN, COLOR_BLACK);
    
    /* Fondo negro con texto cian para información adicional */
    attr_info = VIDEO_ATTR(COLOR_LIGHT_CYAN, COLOR_BLACK);
    
    /* Fondo negro para limpiar la pantalla */
    attr_background = VIDEO_ATTR(COLOR_WHITE, COLOR_BLACK);
    
    /* Limpiar la pantalla */
    clear_screen(attr_background);
    
    /* Imprimir mensaje de bienvenida */
    print_centered("========================================================", 6, attr_title);
    print_centered("      BIENVENIDO AL SISTEMA OPERATIVO", 7, attr_title);
    print_centered("      Universidad del Cauca - FIET", 8, attr_title);
    print_centered("========================================================", 9, attr_title);
    
    print_centered("Laboratorio de Sistemas Operativos A", 11, attr_text);
    print_centered("Taller 7: Acceso directo a memoria de video", 12, attr_text);
    
    print_centered("Estudiante: Henersson Cobo", 14, attr_info);
    print_centered("Correo: henerssoncobo@unicauca.edu.co", 15, attr_info);
    
    /* Imprimir información técnica */
    print_at("Direccion memoria video: 0xB8000", 5, 18, 
             VIDEO_ATTR(COLOR_YELLOW, COLOR_BLACK));
    print_at("Modo: Texto 80x25", 5, 19, 
             VIDEO_ATTR(COLOR_YELLOW, COLOR_BLACK));
    
    print_centered("Sistema iniciado correctamente en modo protegido", 22, 
                   VIDEO_ATTR(COLOR_LIGHT_GREEN, COLOR_BLACK));
    
    /* Bucle infinito - el kernel no debe terminar */
    while(1) {
        /* halt - detiene el CPU hasta la siguiente interrupción */
        __asm__ __volatile__("hlt");
    }
}