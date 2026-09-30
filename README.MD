# **Project Folder Structure**

Tes
Deskripsi tentang struktur folder untuk proyek ini.

## **Struktur Folder**

### **Root Level**

- **`build/`**
  - Folder ini digunakan untuk menyimpan artefak build, seperti file binary hasil kompilasi atau file konfigurasi untuk deployment.
- **`cmd/`**
  - Berisi entry point untuk aplikasi, seperti file `main.go` untuk HTTP/Websocket/Microservice.

- **`configs/`**
  - Folder untuk menyimpan file konfigurasi aplikasi, seperti pengaturan database, API keys, atau environment variables.

- **`db/`**
  - Folder ini digunakan untuk interaksi dengan database, seperti migrasi, konektivitas database.

- **`internal/`**
  - Folder yang berisi logika bisnis atau modul aplikasi. Konten di dalam folder ini tidak boleh diakses langsung oleh modul di luar proyek.
  - **Subfolder:**
    - `modules/`: Berisi modul atau fitur spesifik dari aplikasi.
    - `routes/`: Berisi Root Route untuk mengatur endpoint API dari Route berbagai Modules.

- **`pkg/`**
  - Berisi paket-paket utilitas yang bisa digunakan kembali di berbagai bagian aplikasi.
  - **Subfolder:**
    - `error/`: Berisi Standard Error Response.
    - `redis/`: Fungsi atau helper terkait Redis.
    - `success/`: Berisi respon atau utilitas untuk manajemen success response.
    - `util/`: Berisi fungsi, helper, atau integrasi dengan API lainnya.
    - `validation/`: Logika untuk validasi input.

## **Cara Kerja Tiap Komponen**

1. **API**
   - `api/main.go` adalah entry point untuk API.
   - Endpoint didefinisikan di dalam `internal/routes/`.
   - Data awal atau dummy dimasukkan melalui folder `seeders/`.

2. **Database**
   - Segala hal yang berhubungan dengan database, seperti migrasi dan skema, dikelola di folder `db/`.

3. **Business Logic**
   - Logika inti aplikasi ditangani di dalam folder `internal/module/`.

4. **Utils**
   - Fungsi-fungsi yang digunakan di seluruh aplikasi berada di dalam `pkg/util/`.

5. **Error Handling**
   - Semua error ditangani secara konsisten melalui `pkg/error/`.

6. **Redis Integration**
   - Untuk caching dan manajemen Redis, gunakan helper yang ada di `pkg/redis/`.

7. **Hot Reload**
   - Gunakan `.air.toml` untuk konfigurasi hot-reloading selama pengembangan.

# Run Application

1. **Install the Dependencies**

```bash
go mod tidy && go mod download
```

2. Copy .env.example

```bash
cp .env.example .env
```

3. Run Migration

- Install Goose CLI

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

- Run Migration

```bash
goose up
```

4. Run Seeders

```bash
go run cmd/seeders/main.go
```

5. Run Application

```bash
go run cmd/api/main.go
```

or

```bash
air
```
