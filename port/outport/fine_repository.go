package outport

import (
	"library/core/domain"

	"github.com/google/uuid"
)

type FineRepository interface {
	Save(fine domain.Fine) (domain.Fine, error)
	HasUnpaidByMember(memberID uuid.UUID) (bool, error)
}
