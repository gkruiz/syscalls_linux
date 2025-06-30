#include <linux/kernel.h>
#include <linux/syscalls.h>
#include <linux/fs.h>
#include <linux/uaccess.h>
#include <linux/slab.h>

#include <linux/sched/signal.h> // for 'for_each_process'
#include <linux/cred.h>          // for uid
#include <linux/pid.h>           // para find_get_pid

#include <linux/netdevice.h>
#include <linux/rtnetlink.h>
#include <linux/if.h>

#include <linux/types.h>
#include <linux/time.h>
#include <linux/sched.h>
#include <linux/string.h>
#include <linux/rcupdate.h>
#include <linux/mm.h>

#include <linux/tick.h>
#include <linux/swap.h>
#include <linux/kernel_stat.h>
#include <linux/cpufreq.h>



/*
SYSCALL PARA OBTENER EL TRAFICO EN RED TIEMPO REAL 
*/
SYSCALL_DEFINE2(get_network_stats, unsigned long __user *, rx_bytes,
                               unsigned long __user *, tx_bytes)
{
    struct net_device *dev;
    unsigned long rx_total = 0;
    unsigned long tx_total = 0;
    int ret = 0;

    rtnl_lock();
    
    // Iterar sobre todas las interfaces de red
    for_each_netdev(&init_net, dev) {
        struct rtnl_link_stats64 stats;
        
        // Obtener estadísticas de la interfaz
        dev_get_stats(dev, &stats);
        
        rx_total += stats.rx_bytes;
        tx_total += stats.tx_bytes;
    }
    
    rtnl_unlock();

    // Convertir bytes a kilobytes
    rx_total /= 1024;
    tx_total /= 1024;

    // Copiar resultados al espacio de usuario
    if (copy_to_user(rx_bytes, &rx_total, sizeof(unsigned long))) {
        ret = -EFAULT;
    }
    
    if (copy_to_user(tx_bytes, &tx_total, sizeof(unsigned long))) {
        ret = -EFAULT;
    }

    return ret;
}






/*
SYSCALL PARA TERMINAR PROCESO POR PID
*/


SYSCALL_DEFINE1(kill_process_by_pid, pid_t, pid)
{
    struct task_struct *task;
    struct pid *pid_struct;

    // Buscar el proceso por su PID
    pid_struct = find_get_pid(pid);
    if (!pid_struct)
        return -ESRCH; // No existe ese PID

    task = pid_task(pid_struct, PIDTYPE_PID);
    if (!task)
        return -ESRCH; // No existe tarea asociada

    // Enviar la señal SIGKILL
    send_sig(SIGKILL, task, 0);

    return 0; // Exito
}



 

/*
SYSCALL PARA OBTENER PROCESOS CON INFORMACION BASICA
*/
struct proc_info2 {
    pid_t pid;
    char name[TASK_COMM_LEN];
    uid_t uid;
    unsigned long ram_usage_kb;
    int priority;
    unsigned long cpu_usage; // no perfecto, solo a modo de ejemplo
    unsigned long start_time;
};


SYSCALL_DEFINE2(get_proc_info, struct proc_info2 __user *, info, int __user *, num_procs)
{
    struct task_struct *task;
    struct proc_info2 kinfo;
    int count = 0;

    for_each_process(task) {
        kinfo.pid = task->pid;
        kinfo.uid = __kuid_val(task->cred->uid);
        strncpy(kinfo.name, task->comm, TASK_COMM_LEN);
        kinfo.name[TASK_COMM_LEN-1] = '\0'; // asegurar terminación nula
        kinfo.ram_usage_kb = (task->mm) ? (get_mm_rss(task->mm) * PAGE_SIZE / 1024) : 0;
        kinfo.priority = task->prio;
        kinfo.cpu_usage = task->utime + task->stime;

        unsigned long boot_time = ktime_get_boottime_seconds(); // tiempo actual desde boot
        unsigned long proc_start_secs = task->start_time / NSEC_PER_SEC;
        kinfo.start_time = boot_time - proc_start_secs;

        if (copy_to_user(&info[count], &kinfo, sizeof(struct proc_info2)))
            return -EFAULT;

        count++;
    }

    if (copy_to_user(num_procs, &count, sizeof(int)))
        return -EFAULT;

    return 0;
}



/*SYSCALL PARA OBTENER INFORMAICON GENERAL DE SISTEMA*/

SYSCALL_DEFINE2(sysinfo_usage, unsigned int __user *, cpu_usage, unsigned int __user *, ram_usage)
{
    unsigned int cpu_percent = 0;
    unsigned int ram_percent = 0;
    struct sysinfo mem_info;
    
    // Obtener uso de CPU
    u64 total_jiffies = 0;
    u64 work_jiffies = 0;
    //unsigned int cpu;
    
    for_each_possible_cpu(cpu) {
        struct kernel_cpustat *kcs = &kcpustat_cpu(cpu);
        
        work_jiffies += kcs->cpustat[CPUTIME_USER] +
                       kcs->cpustat[CPUTIME_NICE] +
                       kcs->cpustat[CPUTIME_SYSTEM] +
                       kcs->cpustat[CPUTIME_SOFTIRQ] +
                       kcs->cpustat[CPUTIME_IRQ];
        
        total_jiffies += work_jiffies + kcs->cpustat[CPUTIME_IDLE] +
                        kcs->cpustat[CPUTIME_IOWAIT] +
                        kcs->cpustat[CPUTIME_STEAL] +
                        kcs->cpustat[CPUTIME_GUEST];
    }
    
    if (total_jiffies > 0)
        cpu_percent = (work_jiffies * 100) / total_jiffies;
    
    // Obtener uso de RAM
    si_meminfo(&mem_info);
    
    unsigned long total_ram = mem_info.totalram * mem_info.mem_unit;
    unsigned long free_ram = mem_info.freeram * mem_info.mem_unit;
    unsigned long buffer_ram = mem_info.bufferram * mem_info.mem_unit;
    
    if (total_ram > 0)
        ram_percent = 100 - ((free_ram + buffer_ram) * 100 / total_ram);
    
    // Copiar a espacio de usuario
    if (copy_to_user(cpu_usage, &cpu_percent, sizeof(unsigned int)))
        return -EFAULT;
    
    if (copy_to_user(ram_usage, &ram_percent, sizeof(unsigned int)))
        return -EFAULT;
    
    return 0;
}



/*SYSCALL PARA OBTENER INICO DEL PROCESO SIMPLE SOLO PARA PRUEBAS */



// archivo: process_info.h
#ifndef PROCESS_INFO_H
#define PROCESS_INFO_H

#define MAX_NAME_LEN 256

struct process_info {
    pid_t pid;
    char name[MAX_NAME_LEN];
    unsigned long start_time; // en segundos desde boot
};

#endif



SYSCALL_DEFINE2(proceso_simple, struct process_info __user *, user_buf, int, max_entries)
{
    struct task_struct *task;
    int count = 0;

    for_each_process(task) {
        if (count >= max_entries) break;

        struct process_info info;
        info.pid = task->pid;
        strncpy(info.name, task->comm, MAX_NAME_LEN);
        //info.start_time = (unsigned long)(ktime_to_ns(task->start_time) / NSEC_PER_SEC);
        unsigned long boot_time = ktime_get_boottime_seconds(); // tiempo actual desde boot
        unsigned long proc_start_secs = task->start_time / NSEC_PER_SEC;
        info.start_time = boot_time - proc_start_secs;
        

        if (copy_to_user(&user_buf[count], &info, sizeof(info)))
            return -EFAULT;

        count++;
    }

    return count; // número de procesos copiados
}





