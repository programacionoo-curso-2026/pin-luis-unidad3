package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
}

func main() {
	orders := generateOrders(20)
	fmt.Print("Todas las operaciones completadas. Finalizando\n")
}
func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID: i + 1, Status: "pending",
		}
	}
	return orders
}
func processOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Procesando orden %d\n", order.ID)
	}
}
func updateOrderStatuses(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		status := []string{
			"Procesando", "Despachando", "Entregado",
		}[rand.Intn(3)]
		order.Status = status
		fmt.Printf("Actualizando orden %d con estado: %s\n",
			order.ID, status)
	}
}
func main() {
	orders := generateOrders(20)
	processOrders(orders)
	updateOrderStatuses(orders)
	reportOrderStatus(orders)
	fmt.Print("Todas las operaciones completadas. Finalizando\n")
}
func main() {

	orders := generateOrders(20)

	go processOrders(orders)

	go updateOrderStatuses(orders)

	go reportOrderStatus(orders)

	fmt.Print("Todas las operaciones completadas. Saliendo")

}
func main() {

	var wg sync.WaitGroup
	wg.Add(3)

	orders := generateOrders(20)

	go processOrders(orders)

	go updateOrderStatuses(orders)

	go reportOrderStatus(orders)

	wg.Wait()

	fmt.Print("Todas las operaciones completadas. Saliendo")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	orders := generateOrders(20)

	go func() {
		defer wg.Done()
		processOrders(orders)
	}()

	go func() {
		defer wg.Done()
		updateOrderStatuses(orders)
	}()

	go func() {
		defer wg.Done()
		reportOrderStatus(orders)
	}()

	wg.Wait()

	fmt.Print("Todas las operaciones completadas. Saliendo")
}
