package services

import (
	"fmt"
	"rfid-gateway/internal/dto"
	"rfid-gateway/internal/repositories"
)

type GateService struct {
	GateRepo *repositories.SupabaseRepo
	AIRepo   *repositories.AIRepo
}

func NewGateService(gateRepo *repositories.SupabaseRepo, aiRepo *repositories.AIRepo) *GateService {
	return &GateService{
		GateRepo: gateRepo,
		AIRepo:   aiRepo,
	}
}

func (s *GateService) ProcessTap(req dto.TapRequest) dto.TapResponse {
	member, err := s.getMember(req)
	if err != nil {
		return s.handleMemberNotFound(req)
	}

	if req.GateType == "OUT" {
		return s.handleGateOut(req, member)
	}

	return s.handleGateIn(req, member)
}

func (s *GateService) getMember(req dto.TapRequest) (*dto.Member, error) {
	return s.GateRepo.GetMemberByUID(req.RfidUUID, req.GateType)
}

func (s *GateService) handleMemberNotFound(req dto.TapRequest) dto.TapResponse {
	s.GateRepo.InsertLog(nil, req.RfidUUID, req.GateType, nil, "DENIED", "Akses Ditolak: Kartu Tidak Terdaftar")

	return dto.TapResponse{
		Status:   "DENIED",
		Action:   "DO_NOTHING",
		RfidUUID: req.RfidUUID,
		Message:  "Kartu/Stiker tidak terdaftar",
	}
}

func (s *GateService) handleGateIn(req dto.TapRequest, member *dto.Member) dto.TapResponse {
	fmt.Println("Vehicle in via UHF, skipping AI scan.")

	keterangan := "Akses Masuk Diberikan (UHF)"
	s.GateRepo.InsertLog(&member.ID, req.RfidUUID, req.GateType, nil, "GRANTED", keterangan)

	return dto.TapResponse{
		Status:   "GRANTED",
		Action:   "OPEN_GATE",
		RfidUUID: req.RfidUUID,
		Message:  fmt.Sprintf("Welcome, %s!", member.Name),
	}
}

func (s *GateService) handleGateOut(req dto.TapRequest, member *dto.Member) dto.TapResponse {
	fmt.Println("Vehicle out, triggering AI camera...")

	plate := s.AIRepo.ScanPlate()

	if plate != nil && *plate != member.Plate {
		s.GateRepo.InsertLog(&member.ID, req.RfidUUID, req.GateType, plate, "DENIED", "Plat Nomor Tidak Cocok")

		return dto.TapResponse{
			Status:  "DENIED",
			Action:  "DO_NOTHING",
			Message: "Plat nomor tidak valid",
		}
	}

	keterangan := "Akses Keluar Diberikan (RFID)"
	s.GateRepo.InsertLog(&member.ID, req.RfidUUID, req.GateType, plate, "GRANTED", keterangan)

	return dto.TapResponse{
		Status:   "GRANTED",
		Action:   "OPEN_GATE",
		RfidUUID: req.RfidUUID,
		Message:  fmt.Sprintf("Goodbye, %s!", member.Name),
	}
}