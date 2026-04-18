import cv2
import easyocr
import os
import re
# from ultralytics import YOLO

print("Loading AI Model...")
reader = easyocr.Reader(['en'])
print("AI Model Ready!")

def capture_and_scan():
    image_path = "dummy_plates/plat.jpg"

    img = cv2.imread(image_path)

    lebar_target = 600
    rasio = lebar_target / img.shape[1]
    # tinggi otomatis mengikuti rasio lebar
    tinggi_target = int(img.shape[0] * rasio)
    
    # ngecilin gambar
    img_resized = cv2.resize(img, (lebar_target, tinggi_target), interpolation=cv2.INTER_AREA)

    # grayscale
    gray = cv2.cvtColor(img_resized, cv2.COLOR_BGR2GRAY)
    
    # hitam putih pekat
    _, bw_image = cv2.threshold(gray, 100, 255, cv2.THRESH_BINARY)

    result = reader.readtext(bw_image, detail=0, paragraph=True, allowlist='ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789')
    print(f"Result bersih: {result}")

    raw_text = " ".join(result)
    clean_text = "".join(e for e in raw_text if e.isalnum()).upper()

    print(f"Hasil raw: {raw_text}")
    print(f"Hasil clean: {clean_text}")
    
    # regex plat indo 
    match = re.search(r'([A-Z]{1,2})(\d{1,5})([A-Z]{1,3})', clean_text)
    
    if match:
        valid_plate = match.group(1) + match.group(2) + match.group(3)
        print(f"Plat Lolos Regex dan valid: {valid_plate}")
        return valid_plate
    
    print(f"Teks terbaca tapi ditolak Regex: {clean_text}")
    return "UNKNOWN"