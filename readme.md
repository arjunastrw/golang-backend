project-root/
|-- cmd/
| |-- main.go (entry point)
|-- internal/
| |-- api/
| | |-- input/
| | | |-- input.go (handler untuk input API)
| | | |-- input_service.go (layanan untuk mengelola input data)
| | |-- output/
| | | |-- output.go (handler untuk output API)
| | | |-- output_service.go (layanan untuk mengelola output data)
| | |-- auth/
| | | |-- auth.go (handler untuk autentikasi)
| | | |-- auth_service.go (layanan untuk autentikasi)
| |-- database/
| | |-- database.go (koneksi dan operasi database)
| |-- encryption/
| | |-- encryption.go (implementasi enkripsi/dekripsi)
| |-- models/
| | |-- models.go (struktur data model)
| |-- websocket/
| |-- websocket.go (implementasi WebSocket)
|-- pkg/
| |-- utils/
| |-- utils.go (fungsi global)
|-- config/
| |-- config.go (konfigurasi aplikasi)
|-- scripts/
| |-- db_migration.sql (skrip migrasi database)
|-- go.mod
|-- go.sum
|-- README.md
