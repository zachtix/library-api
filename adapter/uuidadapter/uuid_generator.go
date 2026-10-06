package uuidadapter

import (
	"library/port/outport"

	"github.com/google/uuid"
)

type UUIDGenerator struct{}

func NewUUIDGenerator() outport.IDGenerator {
	return UUIDGenerator{}
}

func (UUIDGenerator) NewID() uuid.UUID {
	return uuid.New()
}
