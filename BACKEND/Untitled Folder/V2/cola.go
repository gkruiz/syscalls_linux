package main

import "sync"

/*Esta es la cola para trafico de red*/
type FixedQueue struct {
	data []StructTraffic
	cap  int
	mu   sync.Mutex
}

func NewFixedQueue(cap int) *FixedQueue {
	return &FixedQueue{
		data: make([]StructTraffic, 0, cap),
		cap:  cap,
	}
}

func (q *FixedQueue) Enqueue(value StructTraffic) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.data) >= q.cap {
		q.data = q.data[1:]
	}
	q.data = append(q.data, value)
}

func (q *FixedQueue) Values() []StructTraffic {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Devolver una copia para evitar condiciones de carrera externas
	result := make([]StructTraffic, len(q.data))
	copy(result, q.data)
	return result
}

/*
func main() {
	q := NewFixedQueue(11)

	for i := uint64(0); i <= 10; i++ {
		q.Enqueue(i)
	}
	fmt.Println(q.Values()) // [0 1 2 3 4 5 6 7 8 9 10]

	q.Enqueue(14)
	fmt.Println(q.Values()) // [1 2 3 4 5 6 7 8 9 10 14]
}*/

/*ESTA ES LA COLA PARA DATOS DE PROCESO */

type FixedQueueP struct {
	data []ElmentoArr
	cap  int
	mu   sync.Mutex
}

type ElmentoArr struct {
	Data *[]ProcInfoR
}
type ElmentoArrR struct {
	Data []ProcInfoR
}

func NewFixedQueueP(cap int) *FixedQueueP {
	return &FixedQueueP{
		data: make([]ElmentoArr, 0, cap),
		cap:  cap,
	}
}

func (q *FixedQueueP) Enqueue(value ElmentoArr) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.data) >= q.cap {
		q.data = q.data[1:]
	}
	q.data = append(q.data, value)
}

func (q *FixedQueueP) Values() []ElmentoArr {
	q.mu.Lock()
	defer q.mu.Unlock()

	// Devolver una copia para evitar condiciones de carrera externas
	result := make([]ElmentoArr, len(q.data))
	copy(result, q.data)
	return result
}
