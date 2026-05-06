package dto

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type PaginatedMemberResponse struct {
	Data  []Member `json:"data"`
	Page  int      `json:"page"`
	Limit int      `json:"limit"`
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
	ID             string `json:"id,omitempty"`
	Name           string `json:"name"`
	Plate          string `json:"plate"`
	Rfid           string `json:"rfid"`
	Uhf            string `json:"uhf"`
	TelegramChatID string `json:"telegram_chat_id,omitempty"`
}

type LogPayload struct {
	ID         string `json:"id,omitempty"`
	MemberID   *string `json:"member_id,omitempty"`
	ScanUID    string  `json:"scan_uid"`
	GateType   string  `json:"gate_type"`
	AiPlate    *string `json:"ai_plate,omitempty"`
	Status     string  `json:"status"`
	Keterangan string  `json:"keterangan"`
}

// auth
type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}