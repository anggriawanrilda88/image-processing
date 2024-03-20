package entities

import (
	"time"

	"github.com/ubersnap-test/src/infra/constants"
)

type BaseEntity struct {
	isNew         bool
	isModified    bool
	markAsDeleted bool
	createdAt     time.Time
	updatedAt     time.Time
}

func CreateNewBaseEntity() BaseEntity {
	now := time.Now()
	return BaseEntity{
		isNew:         true,
		isModified:    false,
		markAsDeleted: false,
		createdAt:     now,
		updatedAt:     now,
	}
}

func (b *BaseEntity) IsNew() bool {
	return b.isNew
}

func (b *BaseEntity) IsModified() bool {
	return b.isModified
}

func (b *BaseEntity) IsMarkedAsDeleted() bool {
	return b.markAsDeleted
}

func (b *BaseEntity) GetCreatedAt() time.Time {
	return b.createdAt
}

func (m *BaseEntity) GetCreatedAtAsISOString() string {
	return m.createdAt.UTC().Format(constants.ISODateTimeFormat)
}

func (b *BaseEntity) GetUpdatedAt() time.Time {
	return b.updatedAt
}

func (m *BaseEntity) GetUpdatedAtISOString() string {
	return m.updatedAt.UTC().Format(constants.ISODateTimeFormat)
}

func (b *BaseEntity) RestoreFromDeleted() {
	b.markAsDeleted = false
}

func (b *BaseEntity) ResetAsNew() {
	b.isNew = true
	b.markAsUnmodified()
	b.resetTime()
}

func (b *BaseEntity) resetTime() {
	now := time.Now()

	b.createdAt = now
	b.updatedAt = now
}

func (b *BaseEntity) markAsUnmodified() {
	b.isModified = false
}
