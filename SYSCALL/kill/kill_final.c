#define _GNU_SOURCE
#include <stdio.h>
#include <unistd.h>
#include <sys/syscall.h>
#include <errno.h>

#define SYS_kill_process_by_pid 552

int main() {
    pid_t pid = 1787; // Cambia esto al PID real que quieres matar
    long result;

    result = syscall(SYS_kill_process_by_pid, pid);
    if (result == -1) {
        perror("Error en syscall kill_process_by_pid");
        return 1;
    }

    printf("Proceso con PID %d terminado exitosamente.\n", pid);
    return 0;
}
