package entity

import (
	"backend/consts"
	"time"
)

type Project struct {
	ID          string `gorm:"primaryKey;size:36;comment:UUID"`
	Name        string `gorm:"size:64;uniqueIndex"`
	Description string `gorm:"size:256"`

	Tasks []Task `gorm:"foreignKey:ProjectID"`
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

	ModelVersion ModelVersion `gorm:"foreignKey:ModelVersionID"`
	Project      Project      `gorm:"foreignKey:ProjectID"`
	Creator      User         `gorm:"foreignKey:CreatorID"`
}
