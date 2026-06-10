package co.edu.unicauca.callbackregistry;

import co.edu.unicauca.capaDeControladores.CallBackInt;
import co.edu.unicauca.DTOs.NotificacionReproduccionDTO;
import java.rmi.Remote;
import java.rmi.RemoteException;
import java.util.List;

public interface CallbackRegistry extends Remote {
    void registerCallback(String name, CallBackInt callback) throws RemoteException;
    void unregisterCallback(String name) throws RemoteException;
    void notifyAllCallbacks(NotificacionReproduccionDTO payload) throws RemoteException;
    List<String> listRegistered() throws RemoteException;
}
