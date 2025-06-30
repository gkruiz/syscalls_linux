package main

import (
	"sort"
)

func get_process_history(PID int) []ProcInfoR {

	history := make([]ProcInfoR, 0, 11)

	//Primer array recorre todos registro grupo de procesos
	valores := DatoProc.Values()
	for i := 0; i < len(valores); i++ {
		snap := (valores[i]).Data
		//empieza a recorrer todos los procesos para encontrar el que necesita
		for j := 0; j < len(snap); j++ {
			proc := snap[j]
			//VALIDA QUE EL PID SEA IGUAL PARA AGREGARLO
			if int(proc.PID) == PID {
				history = append(history, proc)
				break
			}

		}
	}
	//retorna todos los resultados
	return history

}

func energy_process_estimate(arr []ElmentoArr) []ElmentoArr {

	estimado := arr
	copy(estimado, arr)
	//println(len(arr))
	if len(arr) > 0 {
		//si hay data de procesos puede consultar y ordenar los procesos
		//println("Es mayor a cero ")
		ultimo_v := arr[0]
		//bloquea para copiar
		ultimo_v.mu.Lock()
		defer ultimo_v.mu.Unlock()

		procesos := ultimo_v.Data

		sort.Slice(procesos, func(i, j int) bool {
			return procesos[i].Energy > procesos[j].Energy
		})

	}

	return estimado

}
