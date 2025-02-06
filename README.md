# Technical Test for Golang Developer
## Soal 1: Membuat REST API dengan Golang

### Menjalankan Aplikasi
  - Pull atau clone dari remote repository
  - Jalankan `go mod tidy` untuk mendownload dan install semua dependency yang dibutuhkan
  - Jalankan `go run main.go`
  - Atau jika ingin build, gunakan `go build -o -pt_aka_tech_test`

Setup Database
  - Instal dependency go-migrate (https://github.com/golang-migrate/migrate/tree/master/cmd/migrate). Untuk Windows bisa menggunakan `choco`
  - Pada database Postgresql, buat sebuah database dengan nama `pt_aka_tech_test`
  - Lalu jalankan command `migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS -verbose up` untuk menjalankan semua migrasi `up`
  - Jika ingin menjalankan migrasi `down`, maka gunakan `migrate -database YOUR_DATABASE_URL -path PATH_TO_YOUR_MIGRATIONS -verbose down` 

### Pertanyaan
1. Jelaskan bagaimana Anda meng-handle error handling dan logging dalam kode
Anda.

  - Error seperti `duplicate email` di handle pada layer service, sebelum kemudian dilempar ke layer handler seagai object error.
  - Hal ini dikarenakan layer service dipergunakan untuk menampung business logib, sedangkan handler untuk routing

2. Bagaimana Anda memastikan kode Anda aman dari SQL injection?

  - Pada project ini, sangat kecil kemungkinan terjadi SQL Injection dikarenakan saya tidak menggunakan raw query.
  - Jika kita perlu menggunakan raw query, maka kita bisa menggunakan prepared statement atau named replacement
  - Intinya, jangan sampai input user langsung kita masukan ke dalam query

3. Jelaskan pentingnya unit testing dalam pengembangan software

   Unit testing adalah test otomatis, yang berguna untuk membantu proses QA yang dilakukan secara manual. Hal ini dikarenakan QA manual lebih mungkin untuk melakukan kesalahan atau miss saat testing

