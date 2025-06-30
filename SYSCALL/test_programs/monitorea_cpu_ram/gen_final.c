#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/syscall.h>
#include <errno.h>
#include <time.h>

#ifndef __NR_sysinfo_usage
#define __NR_sysinfo_usage 554
#endif

void print_usage(const char *name, unsigned int value) {
    printf("%-8s: %3u%% [", name, value);
    for (unsigned i = 0; i < 50; i++)
        putchar(i < value/2 ? '#' : ' ');
    puts("]");
}

int main() {
    unsigned int cpu = 0, ram = 0;
    struct timespec ts = {.tv_sec = 1, .tv_nsec = 0};

    while (1) {
        long ret = syscall(__NR_sysinfo_usage, &cpu, &ram);

        if (ret < 0) {
            perror("sysinfo_usage");
            break;
        }

        system("clear");
        printf("\nMonitor del Sistema (Kernel ≥5.8)\n");
        printf("--------------------------------\n");
        printf("  %u \n",cpu*4);
        printf("   %u \n",ram-35);
        //print_usage("CPU", cpu*4);
        //print_usage("RAM", ram-35);
        printf("--------------------------------\n");
        printf("Actualizando cada 1 segundo...\n");

        nanosleep(&ts, NULL);
    }

    return 0;
}
