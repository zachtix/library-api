package service

import (
	"library/core/domain"
	"library/port/inport"
	"library/port/outport"

	"github.com/google/uuid"
)

type memberServiceImpl struct {
	repo outport.MemberRepository
}

func NewMemberService(repo outport.MemberRepository) inport.MemberService {
	return &memberServiceImpl{
		repo: repo,
	}
}

func (s *memberServiceImpl) Create(member domain.Member) (domain.Member, error) {
	member.ID = uuid.New()
	member.Status = domain.MemberStatusActive

	created, err := s.repo.Save(member)
	if err != nil {
		return domain.Member{}, err
	}
	return created, nil
}
func (s *memberServiceImpl) Get(id uuid.UUID) (domain.Member, error) {
	member, err := s.repo.FindByID(id)
	if err != nil {
		return domain.Member{}, err
	}
	return member, nil
}
func (s *memberServiceImpl) UpdateStatus(id uuid.UUID, status domain.MemberStatus) (domain.Member, error) {
	if !status.Valid() {
		return domain.Member{}, domain.ErrInvalidStatus
	}
	member, err := s.repo.SetStatus(id, status)
	if err != nil {
		return domain.Member{}, err
	}
	return member, nil
}
