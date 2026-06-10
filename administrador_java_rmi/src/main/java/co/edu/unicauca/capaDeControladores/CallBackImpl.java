package co.edu.unicauca.capaDeControladores;

import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;
import co.edu.unicauca.DTOs.NotificacionReproduccionDTO;

public class CallBackImpl extends UnicastRemoteObject implements CallBackInt {

    public CallBackImpl() throws RemoteException {
        super();
    }

    @Override
    public void notificar(NotificacionReproduccionDTO mensaje) throws RemoteException {
        mostrarNotificacion(mensaje);
    }

    private void mostrarNotificacion(NotificacionReproduccionDTO notificacion) {
        System.out.println();
        System.out.println("╔══════════════════════════════════════════════╗");
        System.out.println("║   🔔  CALLBACK — AUDIO REPRODUCIDO           ║");
        System.out.println("╠══════════════════════════════════════════════╣");
        System.out.printf( "║  ID Audio    : %-30d║%n", notificacion.getIdAudio());
        System.out.printf( "║  Fecha/Hora  : %-30s║%n", notificacion.getFechaHoraReproduccion());
        System.out.println("╚══════════════════════════════════════════════╝");
        System.out.println();
    }
}
