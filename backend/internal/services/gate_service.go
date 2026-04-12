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
	member, err := s.GateRepo.GetMemberByUID(req.RfidUUID, req.GateType)
	if err != nil {
		s.GateRepo.InsertLog(nil, req.RfidUUID, req.GateType, nil, "DENIED", "Akses Ditolak: Kartu Tidak Terdaftar")
		return dto.TapResponse{
			Status:   "DENIED",
			Action:   "DO_NOTHING",
			RfidUUID: req.RfidUUID,
			Message:  "Kartu/Stiker tidak terdaftar",
		}
	}

	aiPlatePtr := s.AIRepo.ScanPlate()

	if req.GateType == "OUT" {
		fmt.Println("Vehicle out, triggering AI camera...")
		aiPlatePtr = s.AIRepo.ScanPlate()
		
		if aiPlatePtr != nil && *aiPlatePtr != member.PlatNomor {
			s.GateRepo.InsertLog(&member.ID, req.RfidUUID, req.GateType, aiPlatePtr, "DENIED", "Plat Nomor Tidak Cocok")
			return dto.TapResponse{ Status: "DENIED", Action: "DO_NOTHING", Message: "Plat nomor tidak valid" }
		}
	} else {
		fmt.Println("Vehicle in via UHF, skipping AI scan.")
	}

	keterangan := "Akses Masuk Diberikan (UHF)"
	if req.GateType == "OUT" {
		keterangan = "Akses Keluar Diberikan (RFID)"
	}
	s.GateRepo.InsertLog(&member.ID, req.RfidUUID, req.GateType, aiPlatePtr, "GRANTED", keterangan)

	return dto.TapResponse{
		Status:   "GRANTED",
		Action:   "OPEN_GATE",
		RfidUUID: req.RfidUUID,
		Message:  fmt.Sprintf("Welcome, %s!", member.Nama),
	}
}