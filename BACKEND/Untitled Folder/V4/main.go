package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func kill_process(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido, usa POST", http.StatusMethodNotAllowed)
		return
	}

	var data StructKillProcess
	status := 0

	// Intentar decodificar el JSON del body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		status = 1
		return
	}

	// Validar campos requeridos
	if data.PID == 0 {
		http.Error(w, "El campo PID es requerido", http.StatusBadRequest)
		status = 2
		return
	}

	pid := data.PID // PID del proceso a matar
	err = KillProcessByPID(pid)
	if err != nil {
		fmt.Printf("Error al matar el proceso: %v\n", err)
		status = 3
	} else {
		fmt.Println("Proceso terminado con éxito.")
	}

	response := Response{
		Status: status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func network_process(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido, usa POST", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData

	// Intentar decodificar el JSON del body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	valores := DatoRed.Values()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(valores)
}

func info_process(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido, usa POST", http.StatusMethodNotAllowed)
		return
	}

	var data RequestData

	// Intentar decodificar el JSON del body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	valores := DatoProc.Values()

	respuesta := make([]ElmentoArr, 0, 11)
	tamano := len(valores)
	if len(valores) > 0 {
		//devuelve el ultimo valor
		DatoProc.mu.Lock()
		defer DatoProc.mu.Unlock()
		respuesta = append(respuesta, valores[tamano-1])
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func process_energy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido, usa GET", http.StatusMethodNotAllowed)
		return
	}

	valores := DatoProc.Values()

	respuesta := make([]ElmentoArr, 0, 11)
	tamano := len(valores)
	if len(valores) > 0 {
		//devuelve el ultimo valor
		DatoProc.mu.Lock()
		defer DatoProc.mu.Unlock()
		respuesta = append(respuesta, valores[tamano-1])
	}

	ordenado := energy_process_estimate(respuesta)
	print(len(ordenado))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ordenado)
}

func info_process_unique(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido, usa POST", http.StatusMethodNotAllowed)
		return
	}

	var data StructInfoProcess

	// Intentar decodificar el JSON del body
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Validar campos requeridos
	if data.PID == 0 {
		http.Error(w, "El campo PID es requerido", http.StatusBadRequest)
		return
	}

	history := get_process_history(data.PID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

var DatoRed *FixedQueue
var DatoProc *FixedQueueP

func main() {
	http.HandleFunc("/kill_process", kill_process)
	http.HandleFunc("/network_process", network_process)
	http.HandleFunc("/info_process", info_process)
	http.HandleFunc("/info_process_unique", info_process_unique)
	http.HandleFunc("/process_energy", process_energy)

	DatoRed = NewFixedQueue(20)
	go getNetworkDataTime(DatoRed)

	DatoProc = NewFixedQueueP(20)
	go getProcessInfo(DatoProc)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
