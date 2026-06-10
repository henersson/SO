# Proyecto SegundoCorte — Guía de prueba (Windows)

Este README describe cómo ejecutar localmente los servicios y clientes del proyecto en Windows y una lista de verificación para validar el comportamiento (streaming, metadatos, reacciones WebSocket, subida de audios, panel admin).

## Requisitos
- Go (>=1.20) en PATH
- Envoy (ejecutable `envoyServer.exe`) si usas gRPC-Web/Envoy
- Node.js (opcional, para servir archivos estáticos con `npx http-server`) o Python (opcional)
- Puertos usados por defecto:
  - `:50051` — Servidor gRPC de streaming audio
   - `:50053` — Servidor gRPC de metadatos interno
  - `:8080` — Envoy (gRPC-Web gateway)
  - `:8082` — Servidor de metadatos (REST)
  - `:8083` — Servidor de subida de audios (HTTP)
  - `:3001` — Servidor de reacciones WebSocket (nuevo servicio)

Si algún puerto está ocupado, ajusta la configuración correspondiente antes de lanzar los servicios.

## Archivos relevantes
- Cliente (web): [cliente html/index.html](cliente%20html/index.html)
- Cliente admin: [cliente html/admin.html](cliente%20html/admin.html)
- Lógica cliente auxiliar: `cliente html/funciones.js` ([cliente html/funciones.js](cliente%20html/funciones.js))
- Bundle gRPC-web (no tocar): `cliente html/bundle.js` ([cliente html/bundle.js](cliente%20html/bundle.js))
- Servidor reacciones (Go): `servidor-reacciones` (módulo con `main/` y `capaControladores/`)
- Servidor metadatos: `servidorMetadatos` (REST)
- Servidor streaming: `servidor` (gRPC streaming server)
- Servidor de subida: `ServidorStreaming` (endpoints para subir archivos)

## Ejecución (PowerShell)
Abre varias terminales PowerShell (una por servicio). Ejecuta en este orden recomendado:

1) Iniciar `servidorMetadatos` (catálogo REST)

```powershell
cd .\servidorMetadatos
go run .\main
```

2) Iniciar `servidor` (streaming gRPC)

```powershell
cd ..\servidor
go run .\main\servidor.go
```

3) Iniciar Envoy (si usas gRPC-Web):

```powershell
cd ..
\.\envoyServer -c .\envoyConfig2.yaml --disable-hot-restart
```

4) Iniciar `ServidorStreaming` (endpoint HTTP de subida si aplica)

```powershell
cd .\ServidorStreaming
go run .\main\servidor.go
```

5) Iniciar `servidor-reacciones` (WebSocket)

```powershell
cd .\servidor-reacciones
go run .\main
```

Nota: Si `go run` falla por dependencias, ejecuta `go mod tidy` en la carpeta del módulo correspondiente.

## Servir los clientes (evitar file:// CORS)
Es recomendable servir la carpeta `cliente html` con un servidor estático para evitar problemas CORS al abrir archivos directamente:

Con `npx http-server` (Node.js):

```powershell
cd .\cliente html
npx http-server -p 8081
# Abrir http://localhost:8081/index.html y http://localhost:8081/admin.html
```

O con Python 3:

```powershell
cd .\cliente html
python -m http.server 8081
# Abrir http://localhost:8081/index.html y http://localhost:8081/admin.html
```

## Flujo de prueba paso a paso (checklist)
Sigue estos pasos para validar las funcionalidades solicitadas:

1. Iniciar los servicios en el orden indicado en la sección anterior.
2. Abrir el cliente en `http://localhost:8081/index.html`.
3. Verificar que el catálogo carga desde `http://localhost:8082/audios`.
   - Si el catálogo no carga, revisa la consola del navegador y los logs del `servidorMetadatos`.
4. Seleccionar una canción y presionar "Reproducir".
   - El reproductor debe llamar a `window.iniciar_streaming_cancion(titulo, formato)` (definido en `bundle.js`).
5. Al reproducir, el cliente debe abrir/usar la conexión WebSocket a `ws://localhost:3001` (expuesto por `servidor-reacciones`).
   - En la consola del servidor de reacciones verás logs de join/leave/reacción.
6. Con otro navegador o pestaña (o con otro usuario en admin), unirse al mismo `audioId` y enviar reacciones (emoji). Debes ver las animaciones en el cliente y los mensajes propagados.
7. Pausar la reproducción y confirmar que el cliente envía `leave` y que el servidor lo limpia correctamente.
8. Desde `admin.html` (http://localhost:8081/admin.html): registrar metadatos (`POST http://localhost:8082/audios/almacenar`) y subir un archivo (`POST http://localhost:8083/audios/subir`).
   - Confirmar que el nuevo audio aparece en el listado del catálogo.

## Checklist de aceptación rápida
- [ ] Catálogo cargado correctamente (GET :8082/audios)
- [ ] Reproducción por gRPC-Web funcionando (Envoy + :50051)
- [ ] Conexión WebSocket a :3001 y broadcast por canal `audioId`
- [ ] Mensajes `join`/`leave`/`reaccion` procesados y logs visibles
- [ ] Subida de audio y registro de metadatos exitosos
- [ ] Animaciones de reacciones visibles en el cliente

## Solución de problemas
- Si ves errores CORS en el navegador, sirve `cliente html` con http-server o python -m http.server (ver sección "Servir los clientes").
- Si Envoy no arranca o está apuntando a otro puerto, revisa `envoyConfig2.yaml` y ajusta `ENVOY_HOST` en `cliente html/bundle.js` si fuera necesario.
- Si `go run` falla por dependencias: en la carpeta del servicio ejecuta `go mod tidy`.

```powershell
cd .\servidor-reacciones
go mod tidy
go run .\main
```

- Logs: revisa las salidas de cada terminal; los handlers modificados incluyen mensajes de log para join/leave/reacción.

## Notas y consideraciones
- El `bundle.js` no fue modificado: el cliente llama a `window.iniciar_streaming_cancion(...)` y depende de Envoy para gRPC-Web.
- El servidor de reacciones fue implementado en Go (módulo `servidor-reacciones`) y expone WS en `:3001` en las rutas `/` y `/ws`.
- Si prefieres usar la versión Node.js del servidor de reacciones, hay artefactos iniciales en el repo pero la implementación final actual es la versión Go por consistencia con la arquitectura del proyecto.

---

Si quieres, puedo:
- Añadir scripts PowerShell para lanzar todos los servicios en pestañas separadas.
- Generar un pequeño Docker Compose para orquestar los servicios localmente.

Solicitud realizada: archivo con instrucciones para probar todo — creado.
