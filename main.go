package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Order model
type Order struct {
	ID         string      `json:"id"`
	Title      string      `json:"title"`
	Restaurant *Restaurant `json:"restaurant"`
	ValidUntil string      `json:"validUntil"`
}

// Restaurant model
type Restaurant struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	Name    string `json:"name"`
}

// INit orders slice using order model
var orders []Order

// Get setall orders
func getOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

//  Get order by id
func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get parameters

	// Loop through orders and find the order with id
	for _, item := range orders {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Order{})
}

// Create a new order
func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order Order
	_ = json.NewDecoder(r.Body).Decode(&order)
	order.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID
	orders = append(orders, order)
	json.NewEncoder(w).Encode(order)

}

// Update a order
func updateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get parameters

	// Loop through orders and find the order with id
	for index, item := range orders {
		if item.ID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			var order Order
			_ = json.NewDecoder(r.Body).Decode(&order)
			order.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID
			orders = append(orders, order)
			json.NewEncoder(w).Encode(order)
			return
		}
	}
	json.NewEncoder(w).Encode(orders)
}

// Delete a order
func deleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get parameters

	// Loop through orders and find the order with id
	for index, item := range orders {
		if item.ID == params["id"] {
			orders = append(orders[:index], orders[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(orders)
}

func main() {
	fmt.Println("app running")
	// Init router
	router := mux.NewRouter()

	// Mock data - @todo - implement CRUD
	orders = append(orders, Order{ID: "1", Title: "Order A", ValidUntil: "2019/12/12", Restaurant: &Restaurant{
		ID: "1", Name: "Paddington", Address: "No 34, Galle Road, Colombo 2"}})
	orders = append(orders, Order{ID: "2", Title: "Order B", ValidUntil: "2019/12/22", Restaurant: &Restaurant{
		ID: "2", Name: "Bake by Bella", Address: "No 14, Maradana Road, Colombo 4"}})

	// Route handlers /End points
	router.HandleFunc("/api/orders", getOrders).Methods("GET")
	router.HandleFunc("/api/orders/{id}", getOrder).Methods("GET")
	router.HandleFunc("/api/orders", createOrder).Methods("POST")
	router.HandleFunc("/api/orders/{id}", updateOrder).Methods("PUT")
	router.HandleFunc("/api/orders/{id}", deleteOrder).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}
