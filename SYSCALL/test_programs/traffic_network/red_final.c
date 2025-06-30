#include <stdio.h>
#include <unistd.h>
#include <sys/syscall.h>
#include <linux/kernel.h>

// Definir el número de syscall (debe coincidir con syscall_64.tbl)
#define SYS_get_network_stats 553

// Wrapper para la syscall
long get_network_stats(unsigned long *rx_kb, unsigned long *tx_kb)
{
    return syscall(SYS_get_network_stats, rx_kb, tx_kb);
}

int main()
{
    unsigned long rx_kb, tx_kb;
    int ret;

    while(1) {
        ret = get_network_stats(&rx_kb, &tx_kb);

        if(ret < 0) {
            perror("Error en syscall");
            return 1;
        }

        printf("Tráfico de red:\n");
        printf("  Recibido: %lu KB\n", rx_kb);
        printf("  Enviado: %lu KB\n", tx_kb);
        printf("-------------------\n");

        sleep(1); // Esperar 1 segundo entre mediciones
    }

    return 0;
}
