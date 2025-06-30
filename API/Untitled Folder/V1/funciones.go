package main

import (
	"syscall"
	"unsafe"
)

const SYS_KILL_PROCESS_BY_PID = 552

func KillProcessByPID(pid int) error {
	_, _, errno := syscall.Syscall(uintptr(SYS_KILL_PROCESS_BY_PID), uintptr(pid), 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

const SYS_GET_NETWORK_STATS = 553

func GetNetworkStats() (rxKB, txKB uint64, err error) {
	var rx, tx uint64

	_, _, errno := syscall.Syscall(
		uintptr(SYS_GET_NETWORK_STATS),
		uintptr(unsafe.Pointer(&rx)),
		uintptr(unsafe.Pointer(&tx)),
		0,
	)

	if errno != 0 {
		return 0, 0, errno
	}

	return rx, tx, nil
}

/*SYSCALL PARA OBTENER INFORMACION DE PROCESOS*/
const (
	SYS_GET_PROC_INFO = 551
	TASK_COMM_LEN     = 16
)

type ProcInfo struct {
	PID        int32
	Name       [TASK_COMM_LEN]byte
	UID        uint32
	RamUsageKB uint64
	Priority   int32
	CPUUsage   uint64
	StartTime  uint64
}

func GetProcInfo() ([]ProcInfo, error) {
	const maxProcs = 1024
	infos := make([]ProcInfo, maxProcs)
	var count int32

	_, _, errno := syscall.Syscall(
		uintptr(SYS_GET_PROC_INFO),
		uintptr(unsafe.Pointer(&infos[0])),
		uintptr(unsafe.Pointer(&count)),
		0,
	)

	if errno != 0 {
		return nil, errno
	}

	return infos[:count], nil
}

/*
//implementacion para obtrener procesos
func main() {
    procs, err := GetProcInfo()
    if err != nil {
        fmt.Printf("Error al obtener procesos: %v\n", err)
        return
    }

    for _, p := range procs {
        name := string(p.Name[:])
        nullIndex := len(name)
        for i, b := range p.Name {
            if b == 0 {
                nullIndex = i
                break
            }
        }
        name = name[:nullIndex] // remover bytes basura tras el null

        fmt.Printf("PID: %d | Nombre: %s | UID: %d | RAM: %d KB | Prio: %d | CPU: %d | Inició hace: %d s\n",
            p.PID, name, p.UID, p.RamUsageKB, p.Priority, p.CPUUsage, p.StartTime)
    }
}

*/

/*
//IMPLEMENTACION PARA TRAFICO DE RED
func main() {
    rx, tx, err := GetNetworkStats()
    if err != nil {
        fmt.Printf("Error al obtener estadísticas de red: %v\n", err)
        return
    }

    fmt.Printf("Tráfico recibido: %d KB\n", rx)
    fmt.Printf("Tráfico transmitido: %d KB\n", tx)
}


*/

/*
//PRUEBA PARA SYSCALL KILL PROCESS
func main() {
    pid := 1234 // PID del proceso a matar
    err := killProcessByPID(pid)
    if err != nil {
        fmt.Printf("Error al matar el proceso: %v\n", err)
    } else {
        fmt.Println("Proceso terminado con éxito.")
    }
}
*/
