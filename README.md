# Panduan Instalasi - Shopping Cart (Go Lang)

Ini adalah panduan instalasi untuk proyek \"Shopping Cart\" yang menggunakan Go Lang. Ikuti langkah-langkah di bawah ini untuk menginstal dan menjalankan proyek ini di lingkungan pengembangan Anda.

## Persyaratan Sistem

Sebelum memulai, pastikan sistem Anda memenuhi persyaratan berikut:

- Go Lang (versi 1.16 atau yang lebih baru)
- Git (opsional)

## Langkah 1: Mengunduh Proyek

Anda dapat mengunduh proyek ini dengan menjalankan perintah berikut di terminal atau menggunakan Git:

<pre><code>git clone https://github.com/nanangNSL/shoppingcart.git</code></pre>

## Langkah 2: Menginstal Dependensi

Setelah Anda berhasil mengunduh proyek, masuk ke direktori proyek dan jalankan perintah berikut untuk menginstal dependensi yang diperlukan:

<pre><code>go mod download</code></pre>

## Langkah 3: Konfigurasi

Proyek ini menggunakan file konfigurasi `config.yaml`. Pastikan untuk menyesuaikan konfigurasi sesuai dengan kebutuhan Anda sebelum menjalankan proyek.

## Langkah 4: Menjalankan Proyek

Setelah Anda selesai mengatur konfigurasi, Anda dapat menjalankan proyek dengan perintah berikut:

<pre><code>go run main.go</code></pre>

Proyek akan dijalankan pada localhost dan port yang ditentukan dalam konfigurasi.

## Catatan Tambahan

- Pastikan Anda memiliki Go Lang yang terinstal dengan benar pada sistem Anda.
- Anda juga dapat menggunakan perintah `go build` untuk menghasilkan file biner dan menjalankannya secara terpisah.
- Jika Anda ingin menjalankan tes, Anda dapat menggunakan perintah `go test` untuk menjalankan tes unit.

Selamat menggunakan proyek \"Shopping Cart\" dengan Go Lang! Jika Anda memiliki pertanyaan atau masalah, jangan ragu untuk menghubungi pengembang.
