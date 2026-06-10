package co.edu.unicauca.callbackregistry;

import com.fasterxml.jackson.databind.ObjectMapper;
import co.edu.unicauca.capaDeControladores.CallBackInt;
import co.edu.unicauca.DTOs.NotificacionReproduccionDTO;
import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpExchange;

import java.io.InputStream;
import java.io.OutputStream;
import java.net.InetSocketAddress;
import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public class CallbackRegistryImpl extends UnicastRemoteObject implements CallbackRegistry {

    private final ConcurrentHashMap<String, CallBackInt> callbacks = new ConcurrentHashMap<>();
    private final ObjectMapper mapper = new ObjectMapper();

    protected CallbackRegistryImpl() throws RemoteException {
        super();
    }

    @Override
    public void registerCallback(String name, CallBackInt callback) throws RemoteException {
        callbacks.put(name, callback);
        System.out.println("[Registry] Registered callback: " + name);
    }

    @Override
    public void unregisterCallback(String name) throws RemoteException {
        callbacks.remove(name);
        System.out.println("[Registry] Unregistered callback: " + name);
    }

    @Override
    public void notifyAllCallbacks(NotificacionReproduccionDTO payload) throws RemoteException {
        System.out.println("[Registry] Notifying " + callbacks.size() + " callbacks");
        for (Map.Entry<String, CallBackInt> e : callbacks.entrySet()) {
            try {
                e.getValue().notificar(payload);
            } catch (Exception ex) {
                System.err.println("[Registry] Error notifying " + e.getKey() + ": " + ex.getMessage());
            }
        }
    }

    @Override
    public List<String> listRegistered() throws RemoteException {
        return new ArrayList<>(callbacks.keySet());
    }

    // Starts a simple HTTP server that accepts POST /notify with a JSON body
    public void startHttpListener(int port) throws Exception {
        HttpServer server = HttpServer.create(new InetSocketAddress(port), 0);
        server.createContext("/notify", new HttpHandler() {
            @Override
            public void handle(HttpExchange exchange) {
                try (InputStream is = exchange.getRequestBody()) {
                    NotificacionReproduccionDTO dto = mapper.readValue(is, NotificacionReproduccionDTO.class);
                    CallbackRegistryImpl.this.notifyAllCallbacks(dto);
                    String resp = "ok";
                    exchange.sendResponseHeaders(200, resp.length());
                    try (OutputStream os = exchange.getResponseBody()) {
                        os.write(resp.getBytes());
                    }
                } catch (Exception e) {
                    try {
                        String resp = "error: " + e.getMessage();
                        exchange.sendResponseHeaders(500, resp.length());
                        try (OutputStream os = exchange.getResponseBody()) { os.write(resp.getBytes()); }
                    } catch (Exception ex) { /* ignore */ }
                }
            }
        });
        System.out.println("[Registry] HTTP listener started on :" + port + " (/notify)");
        server.start();
    }
}
