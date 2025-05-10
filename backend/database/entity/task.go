package entity

import (
	"backend/consts"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID          string `gorm:"primaryKey;size:36;comment:UUID"`
	Name        string `gorm:"size:64;uniqueIndex"`
	Description string `gorm:"size:256"`

	Tasks []Task `gorm:"foreignKey:ProjectID"`
}

func (p *Project) BeforeCreate(tx *gorm.DB) (err error) {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	return
}

type Task struct {
	ID             string `gorm:"primaryKey;size:36;comment:UUID"`
	ModelVersionID uint   `gorm:"index"`
	ProjectID      string `gorm:"index;size:36;comment:UUID"`
	CreatorID      uint   `gorm:"index"`

	Description string            `gorm:"size:256"`
	InputPath   string            `gorm:"size:512;comment:MinIO路径"`
	OutputPath  string            `gorm:"size:512;comment:MinIO路径"`
	Status      consts.TaskStatus `gorm:"type:ENUM('pending','running','completed','failed');index"`

	CreatedAt  time.Time
	FinishedAt *time.Time

	ModelVersion ModelVersion `gorm:"foreignKey:ModelVersionID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Project      Project      `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Creator      User         `gorm:"foreignKey:CreatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == "" {
		t.ID = uuid.NewString()
	}
	return
}
