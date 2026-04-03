# RFID LPR Gate System

Sistem gerbang cerdas yang mengintegrasikan teknologi RFID dan LPR (License Plate Recognition) untuk kontrol akses otomatis.

## Fitur

- **Kontrol Akses RFID**: Otentikasi pengguna melalui kartu RFID
- **LPR (License Plate Recognition)**: Pengenalan plat nomor kendaraan
- **Gateway API**: API terpusat untuk komunikasi antar komponen
- **UI Monitoring**: Dashboard web untuk memantau status gerbang

## Teknologi yang Digunakan

- **Frontend**: SvelteKit
- **Backend**: Go (Fiber)
- **Database**: Supabase
- **AI Engine**: Python (Yolo, OpenCV, FastAPI)
- **Infrastruktur**: Docker & Docker Compose

## Instalasi

1. Clone repositori:
   ```bash
   git clone <repository-url>
   cd rfid-lpr-gate-system
   ```


## Penggunaan

- Akses dashboard di: [http://localhost:5173](http://localhost:5173)
- Akses API gateway di: [http://localhost:3000](http://localhost:3000)

## Struktur Proyek

```
rfid-lpr-gate-system/
├── ai-engine/
├── backend/          
├── client/         
├── iot/                  
├── docker-compose.yml  
└── README.md         
```
