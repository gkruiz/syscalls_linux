#include <stdio.h>
#include <unistd.h>
#include <sys/syscall.h>
#include <errno.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#define SYS_get_proc_info 551  // Número correcto según syscall_64.tbl

struct proc_info {
    pid_t pid;
    char name[16];
    uid_t uid;
    unsigned long ram_usage_kb;
    int priority;
    unsigned long cpu_usage;
    unsigned long start_time;
};


time_t get_boot_time() {
    FILE *fp = fopen("/proc/uptime", "r");
    if (!fp) return -1;

    double uptime_seconds;
    if (fscanf(fp, "%lf", &uptime_seconds) != 1) {
        fclose(fp);
        return -1;
    }

    fclose(fp);

    time_t now = time(NULL);
    return now - (time_t)uptime_seconds;
}


int main() {
    struct proc_info *infos;
    int max_procs = 512; // máximo número de procesos que esperamos (puedes ajustar)
    int num_procs = 0;
    long result;

    infos = malloc(max_procs * sizeof(struct proc_info));
    if (!infos) {
        perror("malloc");
        return 1;
    }

    result = syscall(SYS_get_proc_info, infos, &num_procs);
    if (result == -1) {
        perror("syscall get_proc_info failed");
        free(infos);
        return 1;
    }

    printf("Total processes: %d\n", num_procs);
    //long ticks_per_sec = sysconf(_SC_CLK_TCK);
    time_t boot_time = get_boot_time();
    for (int i = 0; i < num_procs; i++) {

        time_t start_time = boot_time + infos[i].start_time;
        struct tm *tm_info = localtime(&start_time);
        char time_str[64];

        strftime(time_str, sizeof(time_str), "%Y-%m-%d %H:%M:%S", tm_info);



        printf("PID: %d, Name: %s, UID: %d, RAM: %lu KB, Priority: %d, CPU time: %.2f\n, Inicio: %s\n",
               infos[i].pid, infos[i].name, infos[i].uid, infos[i].ram_usage_kb, infos[i].priority, infos[i].cpu_usage/ (float)1e9,time_str);
    }

    free(infos);
    return 0;
}
