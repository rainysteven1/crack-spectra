package entity

import "time"

type Model struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Description string `gorm:"size:256"`
	Type        string
	CreatedAt   time.Time

	Versions []ModelVersion `gorm:"foreignKey:ModelID"`
}

type ModelVersion struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	ModelID uint `gorm:"index"`

	Tag         string `gorm:"size:64;uniqueIndex"`
	Immediate   bool
	CronExpr    string // 设置时间去执行训练
	Metrics     string // 模型训练指标
	StoragePath string // 模型存储位置
	TrainParams string
	TrainedAt   time.Time
	CreatedAt   time.Time

	Model Model `gorm:"foreignKey:ModelID;constraint:OnDelete:CASCADE"`
}
