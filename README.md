# Quizy
Aplikasi web quiz untuk sekolah,perusahaan, dll. 

## Cara Menjalankan Aplikasi
### 1. Install Golang
sebelum menjalankan aplikasi ini pastikan bahasa golang sudah terinstal di kompter, bila belum silahkan klick [disini](https://golang.org/ "Klick Disini") ikuti langkah install dari dokumentasi golang.

### 2. Setting Environmental Variabel
Silahkan set environment variabel di komputer Anda, jika menggunakan windows silahkan ikuti langkah-langkah berikut:
1.  tekan tombol *windows* pada keyboard/layar Anda
2. Ketikan kata "environment" pilih *edit environment variable for your account*
3. Klick new untuk membuat variable baru ada 3 variabel yang harus dibuat yaitu PORT, MYSQL_DSN, dan SECRET. Pastikan Anda sudah membuat database mysql terlebih dahulu, contoh pengisian variabel:
- PORT `:8081`
- MYSQL DSN `username:password@tcp(127.0.0.1:3306)/nama_database?charset=utf8mb4&parseTime=True&loc=Local`
- SECRET `Handsome_Dog`

### 3. Jalankan Aplikasi
untuk menjalankan aplikasi silahkan buka command promt/CMD, pastikan posisi directory sudah ada di directory root aplikasi lalu ketik command dibawah ini `go run *.go` atau `go run main.go` jika ingin di build masukan `go build -o nama_executable` jika system operasi yang Anda gunakan bukan windows maka anda wajib menulliskan .exe pada akhir nama executable nya.

## Packages
Berikut daftar library/package yang digunakan

|  Nama Library |  Versi |
| ------------ | ------------ |
| Driver MYSQL  | v.1.5.0  |
|  Gorilla CSRF  | v.1.7.0  |
| Gorilla MUX | v1.8.0 |
| Gorilla Sessions | v.1.2.1 |
| Crypto Bcrypt | v.0.0.0-2020 |
| GORM ORM | v.1.20.8 |
