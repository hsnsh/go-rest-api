package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"

	. "go-rest-api/helpers"
	. "go-rest-api/models"
)

var productStore = make(map[string]ProductDto)
var id int = 0

// HTTP GET - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []ProductDto

	// Get all products from the store
	for _, product := range productStore {
		products = append(products, product)
	}

	// Convert Dto object to Json object
	data, err := json.Marshal(products)
	CheckError(err)

	// Return Json data object
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HTTP GET - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {

	// Get variables from request url
	variables := mux.Vars(r)

	// Get ID parameter
	key := variables["id"]

	// Check Id from exist in Store
	productDto, exist := productStore[key]
	if !exist {
		// Return Not Found
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	// Convert Dto object to Json object
	data, err := json.Marshal(productDto)
	CheckError(err)

	// Return Json data object
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HTTP POST - /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {

	var productCreateDto ProductCreateDto

	// Get create dto from request body
	err := json.NewDecoder(r.Body).Decode(&productCreateDto)
	CheckError(err)

	// Create store entity
	var productDto ProductDto
	id++
	productDto.ID = id
	productDto.CreationTime = time.Now()
	productDto.UpdateTime = productDto.CreationTime
	productDto.Name = productCreateDto.Name
	productDto.Description = productCreateDto.Description
	key := strconv.Itoa(id)

	// Add store entity to store
	productStore[key] = productDto

	// Convert Dto object to Json object
	data, err := json.Marshal(productDto)
	CheckError(err)

	// Return Json data object
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// HTTP PUT - /api/products/{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {

	// Get variables from request url
	variables := mux.Vars(r)

	// Get ID parameter
	key := variables["id"]

	// Check Id from exist in the store
	productDto, exist := productStore[key]
	if !exist {
		// Return Not Found
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	// Get update dto from request body
	var productUpdateDto ProductUpdateDto
	err := json.NewDecoder(r.Body).Decode(&productUpdateDto)
	CheckError(err)

	// Update store entity object
	productDto.UpdateTime = time.Now()
	productDto.Name = productUpdateDto.Name
	productDto.Description = productUpdateDto.Description

	// Update the store
	productStore[key] = productDto

	// Convert Dto object to Json object
	data, err := json.Marshal(productDto)
	CheckError(err)

	// Return Json data object
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// HTTP DELETE - /api/products/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	// Get variables from request url
	variables := mux.Vars(r)

	// Get ID parameter
	key := variables["id"]

	// Check Id from exist in Store
	_, exist := productStore[key]
	if !exist {
		// Return Not Found
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
		return
	}

	// Remove data from store
	delete(productStore, key)

	// Return No Content
	w.WriteHeader(http.StatusNoContent)
}
