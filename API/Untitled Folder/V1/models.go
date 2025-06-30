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
