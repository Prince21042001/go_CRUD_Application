package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Car struct {
	ID    int    `json:"id"`
	Company  string `json:"company"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

var (
	cars       = make(map[int]Car) 
	nextID     = 1                 
	carsMux    sync.Mutex   
)

func createCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newCar Car
	err := json.NewDecoder(r.Body).Decode(&newCar)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	carsMux.Lock()
	defer carsMux.Unlock()

	newCar.ID = nextID
	nextID++
	cars[newCar.ID] = newCar

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCar)
}

func getCarHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/cars/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	carsMux.Lock()
	defer carsMux.Unlock()

	car, exists := cars[id]
	if !exists {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func updateCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/cars/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	var updatedCar Car
	err = json.NewDecoder(r.Body).Decode(&updatedCar)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	carsMux.Lock()
	defer carsMux.Unlock()

	_, exists := cars[id]
	if !exists {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	updatedCar.ID = id
	cars[id] = updatedCar

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCar)
}

func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/cars/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	carsMux.Lock()
	defer carsMux.Unlock()

	_, exists := cars[id]
	if !exists {
		http.Error(w, "Car not found", http.StatusNotFound)
		return
	}

	delete(cars, id)

	w.WriteHeader(http.StatusNoContent)
}

func listCarsHandler(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	carsMux.Lock()
	defer carsMux.Unlock()

	var carList []Car
	for _, car := range cars {
		carList = append(carList, car)
	}

	json.NewEncoder(w).Encode(carList)
}

func main() {
	http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			listCarsHandler(w)
		case http.MethodPost:
			createCarHandler(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/cars/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCarHandler(w, r)
		case http.MethodPut:
			updateCarHandler(w, r)
		case http.MethodDelete:
			deleteCarHandler(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on port 8060...")
	log.Fatal(http.ListenAndServe(":8060", nil))
}
