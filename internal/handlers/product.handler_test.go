package handlers

import (
	"coffeeshop/config"
	"coffeeshop/internal/repository"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoProductMock = repository.RepoMock{}
var reqBody = `{
	"photo_product": "http://localhost:3001/user/image/veggie_tomato_mix.jpg",
	"product_name": "Veggie Tomato Mix",
	"price": 34000,
	"description": "A hazelnut latte is a delicious coffee beverage that combines espresso, steamed milk, and hazelnut flavoring. It's typically made by adding a shot of espresso to a cup or glass, then mixing in steamed milk infused with hazelnut syrup or flavoring.",
	"size": ["250 gr","300 gr","500 gr"],
	"delivery_method": ["Home Delivery", "Dine In"],
	"start_hour": "13:00:00",
	"end_hour": "19:00:00",
	"stock": 50,
	"product_type": "food"
}`

// Create Product
func TestCreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()

	handler := NewProduct(&repoProductMock)
	expectedResult := &config.Result{Message: "1 data product created"}
	repoProductMock.On("CreateProduct", mock.Anything).Return(expectedResult, nil)

	r.POST("/", handler.PostProduct)
	req := httptest.NewRequest("POST", "/", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, `{"status": "Created", "description": "1 data product created"}`, w.Body.String())
}

// Get Product
func TestGetProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()
	handler := NewProduct(&repoProductMock)

	products := []map[string]interface{}{
		{
			"product_id":      "bd691506-3fce-468f-9757-9f62da056db1",
			"PhotoUpload":     nil,
			"photo_product":   "http://localhost:3001/user/image/veggie_tomato_mix.jpg",
			"product_name":    "Veggie Tomato Mix",
			"price":           34000,
			"description":     "Veggie with Tomato Mix",
			"size":            []string{"250 gr", "300 gr", "500 gr"},
			"delivery_method": []string{"Home Delivery", "Dine In"},
			"start_hour":      "0000-01-01T13:00:00Z",
			"end_hour":        "0000-01-01T17:00:00Z",
			"stock":           50,
			"product_type":    "food",
			"created_at":      "2024-05-06T23:49:00.537108Z",
			"updated_at":      nil,
		},
	}

	meta := map[string]interface{}{
		"next":  nil,
		"prev":  nil,
		"total": 1,
	}

	repoProductMock.On("SearchProduct", mock.Anything, mock.Anything, mock.Anything).Return(&config.Result{Data: products, Meta: meta}, nil)

	expectedJSONResult := map[string]interface{}{
		"status": "OK",
		"data":   products,
		"meta":   meta,
	}

	expectedJSON, err := json.Marshal(expectedJSONResult)
	if err != nil {
		t.Errorf("Error marshaling expectedJSON: %v", err)
	}

	r.GET("/", handler.GetProduct)
	req := httptest.NewRequest("GET", "/?search=v&page=1", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, string(expectedJSON), w.Body.String())
}

// Update Product
func TestUpdateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()

	handler := NewProduct(&repoProductMock)
	expectedResult := &config.Result{Message: "1 data product updated"}
	repoProductMock.On("UpdateProduct", mock.Anything, mock.Anything).Return(expectedResult, nil)

	var reqBodyUpdate = `{
		"product_name": "Veggie Tomato",
		"price": 35000
	}`

	r.PATCH("/:id", handler.PatchProduct)
	req := httptest.NewRequest("PATCH", "/bd691506-3fce-468f-9757-9f62da056db1", strings.NewReader(reqBodyUpdate))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status": "OK", "description": "1 data product updated"}`, w.Body.String())
}

// Delete Product
func TestDeleteProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	w := httptest.NewRecorder()

	handler := NewProduct(&repoProductMock)
	expectedResult := &config.Result{Message: "1 data product deleted"}
	repoProductMock.On("RemoveProduct", mock.Anything).Return(expectedResult, nil)

	r.DELETE("/:id", handler.DeleteProduct)
	req := httptest.NewRequest("DELETE", "/bd691506-3fce-468f-9757-9f62da056db1", nil)
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status": "OK", "description": "1 data product deleted"}`, w.Body.String())
}
