package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func kill_process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // o "http://tu-dominio.com"
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodGet {
		print(r.Method)
		http.Error(w, "Método no permitido, usa GET", http.StatusMethodNotAllowed)
		return
	}
	pid := r.URL.Query().Get("pid")

	println(pid)

	println("adfads2")
	// Validar campos requeridos
	if pid == "" {
		println("adfads")
		http.Error(w, "El campo PID es requerido", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	err = KillProcessByPID(num)
	if err != nil {
		fmt.Printf("Error al matar el proceso: %v\n", err)
	} else {
		fmt.Println("Proceso terminado con éxito.")
	}

	response := Response{
		Status: num,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func network_process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // o "http://tu-dominio.com"
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido, usa GET", http.StatusMethodNotAllowed)
		return
	}

	valores := DatoRed.Values()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(valores)
}

func info_process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // o "http://tu-dominio.com"
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respuesta)
}

func process_energy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // o "http://tu-dominio.com"
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

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
	w.Header().Set("Access-Control-Allow-Origin", "*") // o "http://tu-dominio.com"
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodGet {
		print(r.Method)
		http.Error(w, "Método no permitido, usa GET", http.StatusMethodNotAllowed)
		return
	}
	pid := r.URL.Query().Get("pid")

	println(pid)

	println("adfads2")
	// Validar campos requeridos
	if pid == "" {
		println("adfads")
		http.Error(w, "El campo PID es requerido", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	history := get_process_history(num)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(history)
}

func get_global_info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // o "http://tu-dominio.com"
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido, usa GET", http.StatusMethodNotAllowed)
		return
	}

	valores := DatoGen.Values()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(valores)
}

func get_global_info_unique(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // o "http://tu-dominio.com"
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido, usa GET", http.StatusMethodNotAllowed)
		return
	}

	valores := DatoGen.Values()

	respuesta := make([]CM_global, 0, 11)
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

var DatoRed *FixedQueue
var DatoProc *FixedQueueP
var DatoGen *FixedQueueG

func main() {
	http.HandleFunc("/kill_process", kill_process)                     //post
	http.HandleFunc("/network_process", network_process)               //get
	http.HandleFunc("/info_process", info_process)                     //get
	http.HandleFunc("/info_process_unique", info_process_unique)       //post
	http.HandleFunc("/process_energy", process_energy)                 //get
	http.HandleFunc("/get_global_info", get_global_info)               //get
	http.HandleFunc("/get_global_info_unique", get_global_info_unique) //get

	//cola para datos genreales va de primero
	DatoGen = NewFixedQueueG(20)

	DatoRed = NewFixedQueue(20)
	go getNetworkDataTime(DatoRed)

	DatoProc = NewFixedQueueP(20)
	go getProcessInfo(DatoProc)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
