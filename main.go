package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Expense struct{
	ID string `json:"id"`
	Amount string `json:"amount"`
	Name string `json:"name"`
	Category *Category `json:"category"`
}

type Category struct{
	CategoryName string `json:"categoryname`
}

var expenses []Expense

func getExpenses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}

func deleteExpense(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range expenses {
		if item.ID == params["id"]{
			expenses = append(expenses[:index], expenses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(expenses)
}

func getExpense(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range expenses {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createExpense(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var expense Expense
	_ = json.NewDecoder(r.Body).Decode(&expense)
	expense.ID = strconv.Itoa(rand.Intn(10000000))
	expenses = append(expenses, expense)
	json.NewEncoder(w).Encode(expense)
}

func updateExpense(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range expenses {
		if item.ID == params["id"]{
			expenses = append(expenses[:index], expenses[index+1:]...)
			var expense Expense
			_ = json.NewDecoder(r.Body).Decode(&expense)
			expense.ID = params["id"]
			expenses = append(expenses, expense)
			json.NewEncoder(w).Encode(expense)
		}
	}
} 

func main(){
	r := mux.NewRouter()

	expenses = append(expenses, Expense{ID: "1", Amount: "12.99", Name: "Test Expense", Category: &Category{CategoryName:"Transport"}})
	expenses = append(expenses, Expense{ID: "2", Amount: "14.50", Name: "Test Expense 2", Category: &Category{CategoryName:"Entertainment"}})

	r.HandleFunc("/expenses", getExpenses).Methods("GET")
	r.HandleFunc("/expenses/{id}", getExpense).Methods("GET")
	r.HandleFunc("/expenses", createExpense).Methods("POST")
	r.HandleFunc("/expenses/{id}", updateExpense).Methods("POST")
	r.HandleFunc("/expenses/{id}", deleteExpense).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}