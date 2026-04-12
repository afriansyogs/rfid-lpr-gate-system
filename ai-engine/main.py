from fastapi import FastAPI
from scanner import capture_and_scan

app = FastAPI(title="LPR AI Engine")

@app.get("/api/ai/scan")
def scan_static_plate():
    plat_nomor = capture_and_scan()

    print(f"Plat nomor terdeteksi: {plat_nomor}")
    
    return {
        "status": "success",
        "plate_number": plat_nomor
    }