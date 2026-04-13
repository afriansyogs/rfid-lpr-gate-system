#include <SPI.h>
#include <MFRC522.h>
#include <WiFi.h>
#include <HTTPClient.h>
#include <ArduinoJson.h>
#include <Wiegand.h>

// module UHF
WIEGAND wg;
#define PIN_WG_D0 4
#define PIN_WG_D1 5

// module RFID
#define SS_PIN 21
#define RST_PIN 22
MFRC522 rfid(SS_PIN, RST_PIN);

// module servo
#define RELAY_GATE_IN 26
#define RELAY_GATE_OUT 27

const char* ssid = "NAMA_WIFI";
const char* password = "PASSWORD_WIFI";
// Ip laptop for local development
const char* serverUrl = "http://192.168.1.xxx:3000/api/gate/tap"; 

void setup() {
  Serial.begin(115200);

  pinMode(RELAY_GATE_IN, OUTPUT);
  pinMode(RELAY_GATE_OUT, OUTPUT);
  digitalWrite(RELAY_GATE_IN, HIGH);
  digitalWrite(RELAY_GATE_OUT, HIGH);

  WiFi.begin(ssid, password);
  Serial.print("Connecting to WiFi");
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\nConnected to WiFi!");

  wg.begin(PIN_WG_D0, PIN_WG_D1);
  Serial.println("UHF Wiegand initialized");

  SPI.begin();
  rfid.PCD_Init();
  Serial.println("RFID initialized");
}

void loop() {
  if (WiFi.status() != WL_CONNECTED) {
    return;
  }

  // GATE IN UHF
  if (wg.available()) {
    String uid_uhf = String(wg.getCode(), HEX);
    uid_uhf.toUpperCase();
    Serial.println("UHF Sticker Detected! UID: " + uid_uhf);

    sendToBackend("IN", uid_uhf);
  }

  // GATE OUT RFID
  if (rfid.PICC_IsNewCardPresent() && rfid.PICC_ReadCardSerial()) {
    String uid_tap = "";
    
    for (byte i = 0; i < rfid.uid.size; i++) {
      uid_tap += String(rfid.uid.uidByte[i] < 0x10 ? "0" : "");
      uid_tap += String(rfid.uid.uidByte[i], HEX);
    }

    uid_tap.toUpperCase();
    Serial.println("Card Detected! UID: " + uid_tap);

    sendToBackend("OUT", uid_tap);

    rfid.PICC_HaltA();
  }
}

void sendToBackend(String gateType, String uid) {
  HTTPClient http;
  
  Serial.println("test connect to backend...");
  http.begin(serverUrl);
  http.addHeader("Content-Type", "application/json");

  // raw JSON
  String payload = "{\"gate_type\":\"" + gateType + "\", \"rfid_uuid\":\"" + uid + "\"}";
  
  int httpResponseCode = http.POST(payload);

  if (httpResponseCode > 0) {
    String response = http.getString();
    Serial.print("HTTP Response code: ");
    Serial.println(httpResponseCode);
    Serial.println("Response from backend: " + response);

    if (response.indexOf("\"action\":\"OPEN_GATE\"") > 0) {
      Serial.println("Akses Diterima. Membuka gerbang " + gateType + "...");
      bukaGerbang(gateType);
    } else {
      Serial.println("Akses DITOLAK.");
    }
  } else {
    Serial.print("Error POST API: ");
    Serial.println(httpResponseCode);
  }

  http.end();
}

void bukaGerbang(String gateType) {
  int relayPin = (gateType == "IN") ? RELAY_GATE_IN : RELAY_GATE_OUT;
  
  digitalWrite(relayPin, LOW); 
  delay(500);                 
  digitalWrite(relayPin, HIGH); 
}