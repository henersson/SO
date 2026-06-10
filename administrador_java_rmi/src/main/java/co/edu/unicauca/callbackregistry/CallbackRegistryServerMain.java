package co.edu.unicauca.callbackregistry;

import java.rmi.registry.LocateRegistry;
import java.rmi.Naming;

public class CallbackRegistryServerMain {
    public static void start(int httpPort) {
        new Thread(() -> {
            try {
                // Start RMI registry on 1099 if not present
                try { LocateRegistry.createRegistry(1099); System.out.println("[Registry] Created RMI registry on 1099"); } catch (Exception e) { System.out.println("[Registry] RMI registry already present or could not create: " + e.getMessage()); }
                CallbackRegistryImpl impl = new CallbackRegistryImpl();
                Naming.rebind("rmi://localhost/CallbackRegistry", impl);
                System.out.println("[Registry] CallbackRegistry bound to rmi://localhost/CallbackRegistry");
                impl.startHttpListener(httpPort);
            } catch (Exception e) {
                System.err.println("[Registry] Error starting registry: " + e.getMessage());
                e.printStackTrace();
            }
        }, "CallbackRegistryMain").start();
    }
}
