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
	RfidUUID string `json:"rfid_uuid"`
}

type TapResponse struct {
	Status   string `json:"status"`
	Action   string `json:"action"`
	RfidUUID string `json:"rfid_uuid"`
	Message  string `json:"message"`
}
