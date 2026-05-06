package services

import (
	"rfid-gateway/internal/dto"
	"rfid-gateway/internal/repositories"
)

type MemberService struct {
	Repo *repositories.SupabaseRepo
}

func (s *MemberService) FetchMembers(page, limit int, name, rfid, uhf, plate, telegramChatID string) (dto.PaginatedMemberResponse, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	allResults, err := s.Repo.GetFilteredMembers(name, rfid, uhf, plate, telegramChatID)
	if err != nil {
		return dto.PaginatedMemberResponse{}, err
	}

	if allResults == nil {
		allResults = []dto.Member{}
	}

	totalData := len(allResults)
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	if startIndex > totalData {
		startIndex = totalData
	}
	if endIndex > totalData {
		endIndex = totalData
	}

	paginatedResults := allResults[startIndex:endIndex]

	return dto.PaginatedMemberResponse{
		Data:  paginatedResults,
		Page:  page,
		Limit: limit,
	}, nil
}

func (s *MemberService) CreateMember(data *dto.Member) error {
	return s.Repo.CreateNewMember(data)
}

func (s *MemberService) UpdateMember(id string, data *dto.Member) error {
	return s.Repo.UpdateMemberByID(id, data)
}

func (s *MemberService) DeleteMember(id string) error {
	return s.Repo.DeleteMember(id)
}