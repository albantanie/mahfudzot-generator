# Mahfudzot Generator

Mahfudzot Generator adalah REST API yang dapat digunakan oleh semua kalangan untuk menghasilkan kutipan-kutipan dalam bahasa Arab dari para cendekiawan terdahulu. Dengan menggunakan API ini, pengguna dapat dengan mudah mendapatkan inspirasi dan kebijaksanaan dari kata-kata para pemikir besar dalam bahasa Arab.

## Features

- ✨ REST API untuk mengakses kutipan-kutipan Arab (mahfudzot)
- 🔀 Endpoint untuk mendapatkan kutipan acak
- 📚 Pencarian kutipan berdasarkan penulis dan kategori
- 🌐 CORS support untuk integrasi web
- 🐳 Docker support untuk deployment mudah
- 📖 Mode demo dengan data contoh

## Quick Start

### Menjalankan dengan Go

```bash
# Clone repository
git clone https://github.com/albantanie/mahfudzot-generator.git
cd mahfudzot-generator

# Install dependencies
go mod tidy

# Jalankan aplikasi
go run main.go
```

### Menjalankan dengan Docker

```bash
# Build dan jalankan dengan Docker Compose
docker-compose up -d

# Atau build manual
docker build -t mahfudzot-generator .
docker run -p 8080:8080 mahfudzot-generator
```

## API Endpoints

### Health Check
```
GET /health
```

### Quotes

#### Mendapatkan semua kutipan
```
GET /api/v1/quotes?limit=10&page=1
```

#### Mendapatkan kutipan acak
```
GET /api/v1/quotes/random
```

#### Mendapatkan kutipan berdasarkan ID
```
GET /api/v1/quotes/{id}
```

#### Mendapatkan kutipan berdasarkan penulis
```
GET /api/v1/quotes/author/{author}?limit=10&page=1
```

#### Mendapatkan kutipan berdasarkan kategori
```
GET /api/v1/quotes/category/{category}?limit=10&page=1
```

## Response Format

### Success Response
```json
{
  "success": true,
  "data": {
    "id": 1,
    "text_arabic": "العلم نور",
    "text_latin": "Al-'ilmu nur",
    "translation": "Knowledge is light",
    "author": "Imam Ali",
    "category": "Knowledge",
    "source": "Nahj al-Balagha",
    "created_at": "2025-07-18T13:03:08.016283+07:00",
    "updated_at": "2025-07-18T13:03:08.016283+07:00"
  }
}
```

### Error Response
```json
{
  "success": false,
  "error": "Quote not found",
  "message": "Quote with id 999 not found"
}
```

## Configuration

Aplikasi menggunakan environment variables untuk konfigurasi:

```bash
# Server
PORT=8080
HOST=localhost

# Database (PostgreSQL)
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=mahfudzot
DB_SSLMODE=disable
```

Copy `.env.example` ke `.env` dan sesuaikan dengan konfigurasi Anda.

## Database Setup

### PostgreSQL

1. Buat database PostgreSQL
2. Jalankan migration script:
```bash
psql -U postgres -d mahfudzot -f migrations/001_create_quotes_table.sql
```

### Demo Mode

Jika database tidak tersedia, aplikasi akan otomatis berjalan dalam mode demo dengan data contoh.

## Development

### Project Structure
```
.
├── main.go                 # Entry point
├── internal/
│   ├── config/            # Configuration
│   ├── database/          # Database layer
│   ├── handlers/          # HTTP handlers
│   └── models/            # Data models
├── migrations/            # Database migrations
├── docker-compose.yml     # Docker Compose config
├── Dockerfile            # Docker config
└── README.md
```

### Testing

```bash
# Test health endpoint
curl http://localhost:8080/health

# Test random quote
curl http://localhost:8080/api/v1/quotes/random

# Test quotes list
curl "http://localhost:8080/api/v1/quotes?limit=5"

# Test quotes by author
curl "http://localhost:8080/api/v1/quotes/author/Prophet%20Muhammad"

# Test quotes by category
curl "http://localhost:8080/api/v1/quotes/category/Knowledge"

# Test specific quote
curl http://localhost:8080/api/v1/quotes/1
```

## Database Seeding

Aplikasi ini dilengkapi dengan **68+ kutipan mahfudzot** dari 39 ulama dan cendekiawan Islam terkemuka.

### Menjalankan Seeder

```bash
# Menggunakan command line tool
go run cmd/seeder/main.go

# Dengan force flag (override data yang ada)
go run cmd/seeder/main.go -force

# Menggunakan SQL migration
psql -U postgres -d mahfudzot -f migrations/002_seed_quotes.sql
```

### Data yang Tersedia

- **68+ kutipan** dari ulama seperti Nabi Muhammad SAW, Imam Ali, Al-Ghazali, Ibn Sina, Al-Mutanabbi, Ibn Khaldun, dan banyak lagi
- **Transliterasi Latin** untuk setiap kutipan Arab
- **Terjemahan bahasa Inggris**
- **Kategorisasi** berdasarkan topik (Knowledge, Wisdom, Spirituality, dll.)
- **Referensi sumber** yang akurat

Lihat [SEEDER.md](SEEDER.md) untuk detail lengkap tentang data yang tersedia.

## Contributing

1. Fork repository
2. Buat feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push ke branch (`git push origin feature/amazing-feature`)
5. Buat Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

- GitHub: [@albantanie](https://github.com/albantanie)
- Repository: [mahfudzot-generator](https://github.com/albantanie/mahfudzot-generator)
