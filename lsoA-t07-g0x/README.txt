Taller 7 - Laboratorio de Sistemas Operativos A
Universidad del Cauca - Facultad de Ingeniería Electrónica y Telecomunicaciones

Integrantes:
- Nombre: Henersson
  Correo: henersson@unicauca.edu.co

Descripción:
Este proyecto implementa un kernel básico en modo protegido que escribe
directamente en la memoria de video (0xB8000) para limpiar la pantalla
e imprimir un mensaje de bienvenida al usuario.

Estructura del proyecto:
- src/video.h: Definiciones y prototipos para manejo de video
- src/video.c: Implementación de funciones de video
- src/kernel.c: Punto de entrada del kernel (cmain)
- src/start.S: Código de arranque en ensamblador (basado en 00_enter_pmode)
- Makefile: Automatización de compilación y ejecución

Compilación:
$ make

Ejecución:
$ make run

Limpieza:
$ make clean

Requisitos:
- GCC con soporte para 32 bits
- NASM (ensamblador)
- QEMU (emulador)
- GNU Make

Fecha de entrega: 3 de noviembre de 2025