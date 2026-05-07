package repositories

import (
	"fmt"
	"strings"
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
	if err := r.DB.DB.From("parking_logs").Insert(payload).Execute(&result); err != nil {
		fmt.Println("Failed to insert log:", err)
	}
}


// member 
func (r *SupabaseRepo) CreateNewMember(member *dto.Member) error {
	var result []dto.Member
	if err := r.DB.DB.From("members").Insert(member).Execute(&result); err != nil {
		return err
	}
	return nil
}

func (r *SupabaseRepo) DeleteMember(id string) error {
	var result []dto.Member
	if err := r.DB.DB.From("members").Delete().Eq("id", id).Execute(&result); err != nil {
		return err
	}
	return nil
}

func (r *SupabaseRepo) UpdateMemberByID(id string, updatedData *dto.Member) error {
	var result []dto.Member
	if err := r.DB.DB.From("members").Update(updatedData).Eq("id", id).Execute(&result); err != nil {
		return err
	}
	return nil
}

func (r *SupabaseRepo) GetFilteredMembers(name, rfid, uhf, plate, telegramChatID string) ([]dto.Member, error) {
	var results []dto.Member

	query := r.DB.DB.From("members").Select("*").Neq("id", "00000000-0000-0000-0000-000000000000")

	if name != "" {
		safeName := strings.ReplaceAll(name, " ", "%20")
		query = query.Filter("name", "ilike", "*"+safeName+"*")
	}
	if rfid != "" {
		query = query.Eq("rfid", rfid)
	}
	if uhf != "" {
		query = query.Eq("uhf", uhf)
	}
	if plate != "" {
		safePlate := strings.ReplaceAll(plate, " ", "%20")
		query = query.Filter("plate", "ilike", "*"+safePlate+"*")
	}
	if telegramChatID != "" {
		query = query.Eq("telegram_chat_id", telegramChatID)
	}

	if err := query.Execute(&results); err != nil {
		return nil, err
	}

	return results, nil
}

// log
func (r *SupabaseRepo) CreateNewLog(log *dto.LogPayload) error {
	var result []dto.LogPayload
	if err := r.DB.DB.From("parking_logs").Insert(log).Execute(&result); err != nil {
		return err
	}
	return nil
}

func (r *SupabaseRepo) DeleteLog(id string) error {
	var result []dto.LogPayload
	if err := r.DB.DB.From("parking_logs").Delete().Eq("id", id).Execute(&result); err != nil {
		return err
	}
	return nil
}

func (r *SupabaseRepo) UpdateLogByID(id string, updatedData *dto.LogPayload) error {
	var result []dto.LogPayload
	if err := r.DB.DB.From("parking_logs").Update(updatedData).Eq("id", id).Execute(&result); err != nil {
		return err
	}
	return nil
}