package dto

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type TapRequest struct {
	GateType string `json:"gate_type"`
	RfidUUID string `json:"rfid_uuid"`
}

type TapResponse struct {
	Status   string `json:"status"`
	Action   string `json:"action"`
	RfidUUID string `json:"rfid_uuid"`
	Message  string `json:"message"`
}

type AIResponse struct {
	Status      string `json:"status"`
	PlateNumber string `json:"plate_number"`
}

type Member struct {
	ID        string `json:"id"`
	Nama      string `json:"nama"`
	PlatNomor string `json:"plat_nomor"`
}

type LogPayload struct {
	MemberID   *string `json:"member_id,omitempty"`
	ScanUID    string  `json:"scan_uid"`
	GateType   string  `json:"gate_type"`
	AiPlate    *string `json:"ai_plate,omitempty"`
	Status     string  `json:"status"`
	Keterangan string  `json:"keterangan"`
}