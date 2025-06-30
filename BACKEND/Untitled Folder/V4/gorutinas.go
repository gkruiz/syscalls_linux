package main

import (
	"fmt"
	"time"
)

func getNetworkDataTime(datoRed *FixedQueue) error {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		//Obtiene la informacion de red
		rx, tx, err := GetNetworkStats()
		if err != nil {
			fmt.Printf("Error al obtener estadísticas de red: %v\n", err)
			return nil
		}

		temp := StructTraffic{
			RX: rx,
			TX: tx,
		}

		datoRed.Enqueue(temp)
		/*println("GUARDO UN VALOR")
		println(temp.RX)
		println(temp.TX)
		println(len(datoRed.Values()))*/

		//fmt.Printf("Tráfico recibido: %d KB\n", rx)
		//fmt.Printf("Tráfico transmitido: %d KB\n", tx)

	}

}

func getProcessInfo(DatoProc *FixedQueueP) error {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		procs, err := GetProcInfo()
		if err != nil {
			fmt.Printf("Error al obtener procesos: %v\n", err)
			return nil
		}

		//obtiene primero el lista de procesos , n cantidad de lecturas
		data := make([]ProcInfoR, 0, 1024)

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

			temp := ProcInfoR{
				PID:        p.PID,
				Name:       name,
				UID:        p.UID,
				RamUsageKB: p.RamUsageKB,
				Priority:   p.Priority,
				CPUUsage:   p.CPUUsage,
				StartTime:  p.StartTime,
				Energy:     p.RamUsageKB + p.CPUUsage/1000000,
			}

			//agrega el nuevo elemento
			data = append(data, temp)

			/*fmt.Printf("PID: %d | Nombre: %s | UID: %d | RAM: %d KB | Prio: %d | CPU: %d | Inició hace: %d s\n",
			p.PID, name, p.UID, p.RamUsageKB, p.Priority, p.CPUUsage, p.StartTime)*/

		}

		//se genera el array y se guarda en la otra estructura
		arrP := ElmentoArr{
			Data: data,
		}

		//se agrega el nuevo array a la cola general para ver procesos
		DatoProc.Enqueue(arrP)

	}

}
