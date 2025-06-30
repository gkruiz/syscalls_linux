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
