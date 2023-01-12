package main

import (
	"API-CART/config"
	"API-CART/routes"
	"fmt"
	"math"
	"github.com/gin-gonic/gin"
)

func JawabanSatu(price float64) map[string]int {
	price = math.Ceil(price)
	denominations := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	result := make(map[string]int)
	for _, value := range denominations {
		if int(price) >= value {
			count := int(price) / value
			price -= float64(count * value)
			result[fmt.Sprintf("Rp. %d", value)] = count
		}
	}
	if int(price) >= 100 {
		result["Rp. 100"] = int(price) / 100
		price -= float64(int(price) / 100 * 100)
	}
	if int(price) != 0 {
		result["Rp. 50"] = int(price) / 50
		price -= float64(int(price) / 50 * 50)
	}
	return result
}

func JawabanDua(s1, s2 string) bool {
    if len(s1) > len(s2) {
        s1, s2 = s2, s1
    }
    if len(s2) - len(s1) > 1 {
        return false
    }
    var i, j, diff int
    for i < len(s1) && j < len(s2) {
        if s1[i] != s2[j] {
            diff++
            if len(s1) == len(s2) {
                i++
            }
        } else {
            i++
        }
        j++
    }
    diff += len(s2) - i
    if len(s1) != len(s2) {
		diff--
	}
    return diff <= 1
}

func main() {
	price := 145000.0
	result := JawabanSatu(price)
	fmt.Println(result)
	// Output: map[Rp. 100000:1 Rp. 20000:2 Rp. 5000:1]
	
	price = 2050.0
	result = JawabanSatu(price)
	fmt.Println(result)
	// Output: map[Rp. 2000:1 Rp. 100:1]

	fmt.Println(JawabanDua("telkom", "telecom"))  // Output: false
    fmt.Println(JawabanDua("telkom", "tlkom")) // Output: true

	// Jawaban tiga
	// Pada baris terakhir, perintah LISTEN 80 tidak berlaku dalam Dockerfile. 
	// Perintah yang benar digunakan adalah EXPOSE untuk menentukan port yang digunakan oleh aplikasi. Anda perlu menggantinya dengan EXPOSE 80
    // Pada baris pertama, FROM golang tidak menunjukkan nama image yang valid dan mungkin tidak akan ditemukan pada sistem anda. Anda perlu menggunakan versi yang valid dari golang, seperti golang:latest atau golang:1.x.x untuk versi tertentu.
    // Sebaiknya menambahkan CMD atau ENTRYPOINT pada baris terakhir untuk menjalankan aplikasi saat container dijalankan.
	// yang harus diperbaiki ialah
	// FROM golang:latest
	// ADD . /go/src/github.com/telkomdev/indihome/backend
	// WORKDIR /go/src/github.com/telkomdev/indihome
	// RUN go get github.com/tools/godep
	// RUN godep restore
	// RUN go install github.com/telkomdev/indihome
	// EXPOSE 80
	// CMD ["/go/bin/indihome"]

	// Jawaban nomor empat (4)
	// Tujuan utama dari penggunaan microservices adalah untuk membuat arsitektur aplikasi yang fleksibel, skalabel, dan mudah diatur. Microservices memungkinkan tim pengembangan untuk mengembangkan, di-deploy, dan mengelola setiap bagian dari aplikasi secara independen satu sama lain. Ini memungkinkan perusahaan untuk mencapai skalabilitas, kelancaran, dan kecepatan yang dibutuhkan untuk mengejar inovasi dan menanggapi perubahan pasar dengan cepat.

	// Jawaban nomor lima (5)
	// Index adalah fitur yang digunakan dalam database untuk mempercepat proses pencarian data. Cara kerja index adalah dengan menyimpan sebuah copy dari data yang ada dalam tabel, yang disusun secara terurut menurut nilai dari kolom yang ditentukan sebagai index. Ini memungkinkan database untuk menemukan data yang dicari dengan cepat tanpa perlu melakukan pencarian linier pada seluruh tabel
	
	router := gin.New()
	config.Connect()
	routes.CartRoutes(router)
	router.Run(":8080")
}