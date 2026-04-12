package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"rfid-gateway/internal/dto"
)

type AIRepo struct {
	BaseURL string
}

func NewAIRepo(baseURL string) *AIRepo {
	return &AIRepo{BaseURL: baseURL}
}

func (r *AIRepo) ScanPlate() *string {
	resp, err := http.Get(r.BaseURL + "/api/ai/scan")
	if err != nil {
		fmt.Println("Failed to connect Python AI")
		return nil
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var aiData dto.AIResponse
	
	if jsonErr := json.Unmarshal(bodyBytes, &aiData); jsonErr != nil || aiData.Status != "success" {
		return nil
	}
	return &aiData.PlateNumber
}