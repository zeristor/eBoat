package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Set up the structs

// Invoice Struct
type Invoice struct {
	InvoiceID     string  `json:"InvoiceID"`
	InvoiceNumber int     `json:"InvoiceNumber"`
	CompanyName   string  `json:"CompanyName"`
	UserID        string  `json:"UserID"`
	TotalNet      float32 `json:"TotalNet"`
	TotalGross    float32 `json:"TotalGross"`
	CreatedAt     string  `json:"CreatedAt"`
	Currency      string  `json:"Currency"`
}

// Product struct
type Product struct {
	ProductID    string  `json:"ProductID"`
	InvoiceID    string  `json:"InvoiceID"`
	Name         string  `json:"Name"`
	Price        float64 `json:"Price"`
	Currency     string  `json:"Currency"`
	InvoicePrice float64 `json:"InvoicePrice"`
	Amount       int     `json:"Amount"`
	CreatedAt    string  `json:"CreatedAt"`
}

// CurrencyRatescd Struct
type CurrencyRates struct {
	Success   bool     `json:"success"`
	Timestamp int      `json:"timestamp"`
	Base      string   `json:"base"`
	Date      string   `json:"date"`
	Rates     []*Rates `json:"rates"`
}

// Rates struct for, each currency
type Rates struct {
	CurrencyID string  `json:"CurrencyID"`
	Rate       float64 `json:"Rate"`
}

var invoices []Invoice

// Set up the handlers

func createInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func listInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}

func getInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, invoice := range invoices {
		if invoice.InvoiceID == params["invoiceID"] {
			json.NewEncoder(w).Encode(invoice)
		}
	}
}

func deleteInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

// Main function
func main() {
	fmt.Println("eBoat invoice system")

	handleRequests()
}

func handleRequests() {
	r := mux.NewRouter()

	/*
	   The API should have 6 different endpoints:
	   - Create new Invoice
	   - Get list of created Invoices
	   - Get single Invoice
	   - Delete Invoice
	   - Create Product
	   - Delete Product
	*/
	r.HandleFunc("/invoice", createInvoice).Methods("POST")
	r.HandleFunc("/listInvoices", listInvoices).Methods("GET")
	r.HandleFunc("/invoice/{id}", getInvoice).Methods("GET")
	r.HandleFunc("invoice/{id}", deleteInvoice).Methods("DELETE")

	r.HandleFunc("/product", createProduct).Methods("POST")
	r.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8765", r))
}

func updateCurrencyRates() {

}
