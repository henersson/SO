package co.edu.unicauca.main;

import co.edu.unicauca.capaDeControladores.ControladorAdminRest;
import co.edu.unicauca.capaDeControladores.ServidorHttpCallback;
import co.edu.unicauca.capaFachada.FachadaAdministrador;
import co.edu.unicauca.vistas.MenuAdministrador;

public class Main {

    private static final String HOST_ADMIN = "localhost";

    public static void main(String[] args) throws Exception {
        System.out.println("=== Administrador del Sistema de Audios ===");

        int puertoCb = 8081;
        String envPuerto = System.getenv("CALLBACK_PUERTO");
        if (envPuerto != null && !envPuerto.isEmpty()) {
            try {
                puertoCb = Integer.parseInt(envPuerto);
            } catch (NumberFormatException e) {
                System.out.println("CALLBACK_PUERTO inválido, usando 8081.");
            }
        }
        final int puertoCallback = puertoCb;

        // 1. Start local Callback Registry (RMI + HTTP relay)
        int registryHttpPort = 8090;
        String envRegistryPort = System.getenv("CALLBACK_REGISTRY_PUERTO");
        if (envRegistryPort != null && !envRegistryPort.isEmpty()) {
            try { registryHttpPort = Integer.parseInt(envRegistryPort); } catch (NumberFormatException ignored) {}
        }
        boolean iniciarRegistry = true;
        String envStartRegistry = System.getenv("CALLBACK_REGISTRY_START");
        if (envStartRegistry != null && !envStartRegistry.isEmpty()) {
            iniciarRegistry = !("0".equals(envStartRegistry) || "false".equalsIgnoreCase(envStartRegistry) || "no".equalsIgnoreCase(envStartRegistry));
        }
        if (iniciarRegistry) {
            System.out.println("Iniciando CallbackRegistry (RMI + HTTP) en puerto " + registryHttpPort);
            co.edu.unicauca.callbackregistry.CallbackRegistryServerMain.start(registryHttpPort);
        } else {
            System.out.println("CallbackRegistry no iniciado (CALLBACK_REGISTRY_START=" + envStartRegistry + ")");
        }

        // 2. Iniciar servidor HTTP de callbacks
        ServidorHttpCallback servidorCallback = new ServidorHttpCallback(puertoCallback);
        int puertoCallbackReal = servidorCallback.iniciar();

        // 3. Registrar en ServidorStreaming (HTTP) y en CallbackRegistry (RMI)
        ControladorAdminRest controladorRest = new ControladorAdminRest();
        boolean registrado = controladorRest.registrarCallbackEnServidorStreaming(HOST_ADMIN, puertoCallbackReal);
        if (registrado) {
            System.out.println("✓ Registrado para notificaciones (" + HOST_ADMIN + ":" + puertoCallbackReal + ")");
        } else {
            System.out.println("! Advertencia: no se pudo registrar para callbacks.");
            System.out.println("  Verifique que ServidorStreaming esté corriendo en :8083");
        }

        // 3. Mostrar menú
        FachadaAdministrador fachada = new FachadaAdministrador(controladorRest);
        // Register RMI callback with local CallbackRegistry
        try {
            java.rmi.registry.Registry reg = java.rmi.registry.LocateRegistry.getRegistry();
            co.edu.unicauca.callbackregistry.CallbackRegistry registry = (co.edu.unicauca.callbackregistry.CallbackRegistry) java.rmi.Naming.lookup("rmi://localhost/CallbackRegistry");
            String callbackName = "adminCallback-" + java.util.UUID.randomUUID();
            co.edu.unicauca.capaDeControladores.CallBackImpl cbImpl = new co.edu.unicauca.capaDeControladores.CallBackImpl();
            registry.registerCallback(callbackName, cbImpl);
            System.out.println("✓ RMI callback registrado como: " + callbackName);
        } catch (Exception e) {
            System.err.println("! No se pudo registrar callback RMI: " + e.getMessage());
        }
        MenuAdministrador menu = new MenuAdministrador(fachada);
        menu.mostrar();
    }
}
