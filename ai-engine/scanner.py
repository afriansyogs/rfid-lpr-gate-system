import cv2
import easyocr
import os
import re
from ultralytics import YOLO

print("Loading AI Model...")
reader = easyocr.Reader(['en'])
print("AI Model Ready!")

def capture_and_scan():
    image_path = "dummy_plates/plat.jpg"

    img = cv2.imread(image_path)
    result = reader.readtext(img)

    raw_text = "".join([text for (bbox, text, prob) in result])
    clean_text = "".join(e for e in raw_text if e.isalnum()).upper()
    
    # regex plat indo 
    match = re.search(r'([A-Z]{1,2}\d{1,4}[A-Z]{1,3})', clean_text)

    if match:
        valid_plate = match.group(1)
        return valid_plate
    return "UNKNOWN"