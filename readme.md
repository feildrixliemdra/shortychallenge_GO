

# Ralali Golang Microservice Boilerplate

## Pendahuluan
Dengan adanya kebutuhan untuk memecah Arsitektur Ralali yang Monolitik menjadi microservice, maka hadirlan boilerplate ini yang dapat digunakan oleh internal tim ralali untuk menunjang pembangunan microservice menggunakan bahasa pemrograman Go, arsitektur pada mikroservice ini diadoptasi berdasarkan teori yang ada pada link-link berikut ini:

- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
- http://www0.cs.ucl.ac.uk/staff/A.Finkelstein/crsenotes/1B1499REQTS.pdf
- https://blog.alexellis.io/golang-writing-unit-tests/
- http://doc.gorm.io

### Panduan Best Practice
- https://talks.golang.org/2013/bestpractices.slide

### Main Open Source Library
- https://github.com/gin-gonic/gin
- https://github.com/jinzhu/gorm -> (https://github.com/go-sql-driver/mysql)
- https://github.com/joho/godotenv

### Architecture Structure
![architecture diagram](golang%20architecture%20diagram.png)
```
- rl-ms-boilerplate-go
 |- constants
 |- controllers
 |- helpers
 |- middlerware
 |- models
 |- objects
 |- repositories
 |- services
 |- storage
    |- logs
```
- *middleware:* Digunakan untuk menyimpan middleware-middleware yang akan digunakan, contoh `cors_middleware` atau `oauth_middleware`.
- *controllers:* Controller bertugas untuk menghandle HTTP Request, routing dimasukkan per-controller dan digroup berdasarkan controller, controller terhubung dengan service.
- *service:* Service bertugas untuk menghandle business logic, service memungkinkan untuk memanggil banyak repository dan atau service lain.
- *repositories:* Repository bertugas untuk menghandle query-query ke database atau storage lainnya, jangan menambahkan logic-logic programming berat pada layer ini.
- *models:* Models bertugas untuk menampung model-model representasi database schema yang dapat digunakan untuk kepentingan migration atau enkapsulasi data.
- *objects:* Objects bertugas sebagai transporter data antar layer, objects juga bertugas untuk melakukan enkapsulasi data dari HTTP request ataupun sebagai response dari sebuah request.
- *helpers:* Bertugas untuk menyimpan helpers atau libraries yang sering digunakan contohnya `error_helper` atau `redis_helper`.
- *constants:* Digunakan untuk menyimpan constant-constant seperti `error_constants` atau `configuration_constants`.
- *storage:* Storage bertugas untuk menyimpan file-file seperti log error atau temporary file storage.

## Code Versioning
versioning level dilakukan pada layer 
- `controllers` 
- `objects` 
- `repositories` 
- `services`

setiap file pada layer-layer tersebut diberi prefix version dengan format snake case, seperti pada contoh yang ada `v1_user_controller.go` yang berarti user_controller versi 1, dan pada level struct diberi prefix versi dalam bentuk upper camel case seperti pada contoh diproject ini `V1UserController` yang berarti controller `UserController` versi 1.

### Sample Case
terdapat contoh kasus pada saat update data user parameter dan response yang diterima dan diberikan oleh `v1` dan `v2` berbeda, pertama-tama, developer harus melakukan definisi DTO nya terlebih dahulu pada layer `objects`:

- v1_user_object.go
- v2_user_object.go

pada kedua file tersebut terdapat object response dan object request, setelah melakaukan devinisi DTO, developer kemudiam melakukan definisi repository pada layer `repository` yang menggunakan DTO pada masing-masing versi.

setelah melakukan definisi pada `repository`, kemudian dilakukan definisi pada layer `service` dan `controller`, perhatikan routing group pada masing masing controller harus sesuai dengan versi yang didefinisikan.    

## How to Setup Local

1. Clone repository ini dengan menggunakan command:
```bash
git clone git@github.com:ralali/rl-ms-boilerplate-go.git
```
2. Masuk kedalam directory `rl-ms-boilerplate-go` dengan menggunakan command:
```bash
cd rl-ms-boilerplate-go
```

3. Membuat file `src/.env` yang berisi konfigurasi environment, contoh dapat dilihat dari `src/.env.example`

4. Menjalankan `docker-compose` dengan menggunakan command:
```bash
docker-compose up --build
```

Project ini menggunakan docker-compose yang melakukan build image:
- [rll_go_boilerplate_golang] Golang 1.11.5-alpine3.9 
- [rll_go_boilerplate_golang] Mysql 5.7

Engineer dapat menggunakan database mysql pada local machine dengan konfigurasi sebagai berikut:
```
DB_HOST=172.16.231.1
DB_PORT=3309
DB_DATABASE=rll_go_boilerplate_database
DB_USERNAME=rll_go_boilerplate_username
DB_PASSWORD=rll_go_boilerplate_password
```

### Migration
Untuk melakukan migrasi database, engineer harus menjalankan docker-compose terlebih dahulu lalu menjalankan command dibawah ini:
```bash
docker exec rll_go_boilerplate_golang sh -c 'go run main.go migrate'
```
Tunggu hingga command berhasil dijalankan maka database skema berhasil dimigrasi

### Documentation
Dokumentasi project ini menggunakan swagger, untuk mengenerate doc file dari swagger dapat menggunakan command dibawah ini:
``` bash
docker exec rll_go_boilerplate_golang sh -c 'swag init'
```

### Testing
Untuk menjalankan unit test, engineer dapat menggunakan command dibawah ini:
``` bash
docker exec rll_go_boilerplate_golang sh -c 'go test ./... -v -cover'
```

### Dependency Manager
Project ini menggunakan dependency manager `Go Dep`, untuk informasi lebih detail mengenai dependency manager ini dapat mengakses link berikut ini `https://golang.github.io/dep/`.

Untuk menambahkan package baru, dapat menggunakan format command dibawah ini:
``` bash
docker exec rll_go_boilerplate_golang sh -c 'source dep_add.sh {{package-source}}'
```

Berikut adalah contoh untuk menambahkan package dari `github.com/360EntSecGroup-Skylar/excelize`:

``` bash
docker exec rll_go_boilerplate_golang sh -c 'source dep_add.sh github.com/360EntSecGroup-Skylar/excelize'
```

Jika command diatas menampilkan tampilan seperti dibawah ini:

```bash
Fetching sources...

"github.com/360EntSecGroup-Skylar/excelize" is not imported by your project, and has been temporarily added to Gopkg.lock and vendor/.
If you run "dep ensure" again before actually importing it, it will disappear from Gopkg.lock and vendor/.
```

Pesan tersebut berarti anda belum menggunakan package itu diproject anda, hal ini sering kali terjadi karena architecture kita menggunakan sub package, untuk menangani masalah ini, kita harus menambahkan package tersebut pada kolom required didalam file `src/Gopkg.toml`, seperti pada contoh dibawah ini:

```bash
required = [
    ...
    "github.com/360EntSecGroup-Skylar/excelize",
    ...
]
```

Setelah itu untuk membersihkan vendor directory setelah penambahan package baru engineer harus menjalankan perintah dibawah ini:

``` bash
docker exec rll_go_boilerplate_golang sh -c 'dep ensure -v'
```

## How to Setup Docker Image
Sudah disediakan Dockerfile pada root directory ini untuk melakukan build image untuk applikasi ini, untuk melakukan build dapat menggunakan command berikut ini:
```bash
docker build -t rll_go_boilerplate_golang .
```

Setelah docker berhasil dibuild maka image dapat di jalankan dengan menggunakan command berikut ini:
```bash
docker run rll_go_boilerplate_golang
```

### Migration 
Untuk menjalankan migrasi dapat menggunakan command dibawah ini:
```bash
docker run rll_go_boilerplate_golang sh -c 'go run main.go migrate'
```

### Testing
Untuk menjalankan testing dapat menggunakan command dibawah ini:
```bash
docker exec rll_go_boilerplate_golang sh -c 'go test ./... -v -cover'
```

## TODO
- Endpoint documentation