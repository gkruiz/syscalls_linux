package main

// Modelos temporales
type RequestData struct {
	Text   string `json:"text"`
	Number int    `json:"number"`
}

// Estructura de RESPUESTA
type Response struct {
	Status int `json:"status"`
}

//MODELO PARA MATAR PROCESO

type StructKillProcess struct {
	PID int `json:"pid"`
}

//MODELO PARA GUARDAR TRAFICO RED

type StructTraffic struct {
	RX uint64 `json:"rx"`
	TX uint64 `json:"tx"`
}

// MODELO PARA INFO PROCESO INGRESO PETICION
type StructInfoProcess struct {
	PID int `json:"pid"`
}

/*Modelo para respuesta de procesos*/
type ProcInfoR struct {
	PID        int32
	Name       string
	UID        uint32
	RamUsageKB uint64
	Priority   int32
	CPUUsage   uint64
	StartTime  uint64
	Energy     uint64
}
