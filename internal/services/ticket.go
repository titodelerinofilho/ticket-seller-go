package services

import "fmt"

func ProcessOrder(order map[string]interface{}) {
	fmt.Printf("Processando pedido: %+v\n", order)
}
