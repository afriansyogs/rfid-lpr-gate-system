package repositories

import (
	"fmt"
	"rfid-gateway/internal/dto"
	"github.com/nedpals/supabase-go"
)

type SupabaseRepo struct {
	DB *supabase.Client
}

func NewSupabaseRepo(db *supabase.Client) *SupabaseRepo {
	return &SupabaseRepo{DB: db}
}

func (r *SupabaseRepo) GetMemberByUID(uid string, gateType string) (*dto.Member, error) {
	var members []dto.Member
	queryCol := "uhf_id"
	if gateType == "OUT" {
		queryCol = "rfid_id"
	}

	err := r.DB.DB.From("members").Select("*").Eq(queryCol, uid).Execute(&members)
	if err != nil || len(members) == 0 {
		return nil, fmt.Errorf("member not found")
	}
	return &members[0], nil
}

func (r *SupabaseRepo) InsertLog(memberID *string, scanUID string, gateType string, aiPlate *string, status string, keterangan string) {
	payload := dto.LogPayload{
		MemberID:   memberID,
		ScanUID:    scanUID,
		GateType:   gateType,
		AiPlate:    aiPlate,
		Status:     status,
		Keterangan: keterangan,
	}

	var result []dto.LogPayload
	if err := r.DB.DB.From("logs").Insert(payload).Execute(&result); err != nil {
		fmt.Println("Failed to insert log:", err)
	}
}