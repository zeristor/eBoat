package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Set up the structs

// Invoice Struct
/*
Invoice resource has to have the following fields:
- ID
- Invoice Number
- Company Name
- Total Net Amount
- Total Gross Amount (After Tax)
- Created At
- Currency
*/
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
/*
	Product resource has to have the following fields:
	- ID
	- Invoice ID
	- Name
	- Price
	- Currency
	- Invoice Price (the price in currency declared in invoice)∑∑
	- Amount
	- Created At
*/
type Product struct {
	ProductID    string  `json:"ProductID"`
	InvoiceID    string  `json:"InvoiceID"`
	Name         string  `json:"Name"`
	Price        float32 `json:"Price"`
	Currency     string  `json:"Currency"`
	InvoicePrice float32 `json:"InvoicePrice"`
	Amount       int     `json:"Amount"`
	CreatedAt    string  `json:"CreatedAt"`
}

// User struct
type User struct {
	UserID   string `json:"UserID"`
	UserName string `json:"UserName"`
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
