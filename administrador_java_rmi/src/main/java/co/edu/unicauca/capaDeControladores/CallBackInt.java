package co.edu.unicauca.capaDeControladores;

import java.rmi.Remote;
import java.rmi.RemoteException;
import co.edu.unicauca.DTOs.NotificacionReproduccionDTO;

public interface CallBackInt extends Remote {
    void notificar(NotificacionReproduccionDTO mensaje) throws RemoteException;
}
