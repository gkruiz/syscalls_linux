/* SPDX-License-Identifier: GPL-2.0-only */
/*
 * syscalls.h - Linux syscall interfaces (arch-specific)
 *
 * Copyright (c) 2008 Jaswinder Singh Rajput
 */
#ifndef _ASM_X86_SYSCALLS_H
#define _ASM_X86_SYSCALLS_H

/* Common in X86_32 and X86_64 */
/* kernel/ioport.c */
long ksys_ioperm(unsigned long from, unsigned long num, int turn_on);

//asmlinkage long sys_hola_mundo(void);


asmlinkage long sys_get_proc_info(struct proc_info __user *info, int __user *num_procs);

asmlinkage long sys_kill_process_by_pid(pid_t pid);

asmlinkage long sys_get_network_stats(unsigned long __user *rx_bytes,
                                     unsigned long __user *tx_bytes);

asmlinkage long sys_sysinfo_usage(unsigned int __user *cpu_usage, unsigned int __user *ram_usage);

asmlinkage long sys_proceso_simple(struct process_info __user *user_buf, int max_entries);

#endif /* _ASM_X86_SYSCALLS_H */
