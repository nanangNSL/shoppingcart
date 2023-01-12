package controllers

import (
	"API-CART/config"
	"API-CART/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCart(t *testing.T) {
	config.DB = config.InitTestDB()
	defer config.DB.AutoMigrate()

	// Prepare test data
	config.DB.Create(&models.Cart{
		KodeProduk: 1,
		NamaProduk: "Product 1",
		Kuantitas: "5",
	})
	config.DB.Create(&models.Cart{
		KodeProduk: 2,
		NamaProduk: "Product 2",
		Kuantitas: "3",
	})

	// Prepare request and response
	router := gin.Default()
	router.GET("/", GetCart)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	// Assert response
	var carts []models.Cart
	config.DB.Find(&carts)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, carts, w.Body.Bytes())
}

func TestGetFilterProduct(t *testing.T) {
	config.DB = config.InitTestDB()
	defer config.DB.AutoMigrate()

	// Prepare test data
	config.DB.Create(&models.Cart{
		KodeProduk: 1,
		NamaProduk: "Product 1",
		Kuantitas: "5",
	})
	config.DB.Create(&models.Cart{
		KodeProduk: 2,
		NamaProduk: "Product 2",
		Kuantitas: "3",
	})

	// Prepare request and response
	router := gin.Default()
	router.GET("/", GetFilterProduct)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/?namaProduk=Product 1&kuantitas=5", nil)
	router.ServeHTTP(w, req)
	var carts []models.Cart
config.DB.Where("nama_produk = ? AND kuantitas = ?", "Product 1", 5).Find(&carts)
assert.Equal(t, 200, w.Code)
assert.Equal(t, carts, w.Body.Bytes())

}

// func TestCreateCart(t *testing.T) {
// 	config.DB = config.InitTestDB()
// 	defer config.DB.AutoMigrate()

// // Prepare test data
// product := models.Cart{
// 	KodeProduk: 1,
// 	NamaProduk: "Product 1",
// 	Kuantitas: "5",
// }

// // Prepare request and response
// router := gin.Default()
// router.POST("/", CreateCart)
// w := httptest.NewRecorder()
// req, _ := http.NewRequest("POST", "/", nil)
// req.Header.Set("Content-Type", "application/json")
// req.Body = json.Marshal(product)
// router.ServeHTTP(w, req)

// // Assert response
// var cart models.Cart
// config.DB.Where("kode_produk = ?", "P001").First(&cart)
// assert.Equal(t, 200, w.Code)
// assert.Equal(t, cart, w.Body.Bytes())

// }

// func encodeJSON(product models.Cart) {
// 	panic("unimplemented")
// }

func TestDeleteCart(t *testing.T) {
	config.DB = config.InitTestDB()
	defer config.DB.AutoMigrate()

// Prepare test data
config.DB.Create(&models.Cart{
	KodeProduk: 1,
	NamaProduk: "Product 1",
	Kuantitas: "5",
})

// Prepare request and response
router := gin.Default()
router.DELETE("/:id", DeleteCart)
w := httptest.NewRecorder()
req, _ := http.NewRequest("DELETE", "/P001", nil)
router.ServeHTTP(w, req)

// Assert response
var cart models.Cart
config.DB.Where("kode_produk = ?", "P001").First(&cart)
assert.Equal(t, 200, w.Code)
assert.Equal(t, "Product terhapus", w.Body.String())

}
