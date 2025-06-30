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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(valores)
}

var DatoRed *FixedQueue
var DatoProc *FixedQueueP

func main() {
	http.HandleFunc("/kill_process", kill_process)
	http.HandleFunc("/network_process", network_process)
	http.HandleFunc("/info_process", info_process)

	DatoRed = NewFixedQueue(20)
	go getNetworkDataTime(DatoRed)

	DatoProc = NewFixedQueueP(5)
	go getProcessInfo(DatoProc)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
