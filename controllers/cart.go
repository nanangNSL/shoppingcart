package controllers

import (
	"API-CART/config"
	"API-CART/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context){
	var Cart []models.Cart
	config.DB.Find(&Cart)
	c.JSON(200, &Cart)
}

func GetFilterProduct(c *gin.Context) {
    namaProduk := c.Query("namaProduk")
    kuantitas := c.Query("kuantitas")

    var carts []models.Cart

    if namaProduk != "" && kuantitas != "" {
        kuantitasInt, _ := strconv.Atoi(kuantitas)
        config.DB.Where("nama_produk = ? AND kuantitas >= ?", namaProduk, kuantitasInt).Find(&carts)
    } else if namaProduk != "" {
        config.DB.Where("nama_produk = ?", namaProduk).Find(&carts)
    } else if kuantitas != "" {
        kuantitasInt, _ := strconv.Atoi(kuantitas)
        config.DB.Where("kuantitas >= ?", kuantitasInt).Find(&carts)
    } else {
        config.DB.Find(&carts)
    }

    c.JSON(200, &carts)
}





func CreateCart(c *gin.Context) {
    var cart models.Cart
    c.BindJSON(&cart)
    
    // Check if the product already exists in the cart
    existingCart := models.Cart{}
    config.DB.Where("kode_produk = ?", cart.KodeProduk).First(&existingCart)

    if existingCart.KodeProduk != 0 {
        // Product already exists in the cart, update the quantity
        existingCart.Kuantitas += cart.Kuantitas
        config.DB.Save(&existingCart)
        c.JSON(200, &existingCart)
    } else {
        // Product does not exist in the cart, create a new entry
        config.DB.Create(&cart)
        c.JSON(200, &cart)
    }
}

func DeleteCart(c *gin.Context){
	var Cart models.Cart
	config.DB.Where("kode_produk = ?", c.Param("id")).Delete(&Cart)
	c.JSON(200, "Product terhapus")
}

