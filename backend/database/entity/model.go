package entity

import (
	"time"
)

type Model struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(128);not null;uniqueIndex;comment:模型名称"`
	Description string `gorm:"type:text;comment:模型描述"`
	Type        string `gorm:"type:varchar(100);comment:模型类型 (e.g., classification, detection)"`
	CreatedAt   time.Time

	Versions []ModelVersion `gorm:"foreignKey:ModelID"`
}

type ModelVersion struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	ModelID uint `gorm:"index:idx_model_tag;comment:所属模型ID"`

	Tag         string     `gorm:"size:64;uniqueIndex:idx_model_tag;comment:版本标签,模型内唯一"`
	Description string     `gorm:"type:text;comment:版本描述"`
	Immediate   bool       `gorm:"default:false;comment:是否立即训练"`
	CronExpr    string     `gorm:"type:varchar(100);comment:CRON表达式,用于定时训练"`
	Metrics     string     `gorm:"type:json;comment:模型训练指标 (JSON格式)"`
	StoragePath string     `gorm:"type:varchar(512);comment:模型文件存储路径"`
	TrainParams string     `gorm:"type:json;comment:训练参数 (JSON格式)"`
	TrainedAt   *time.Time `gorm:"comment:训练完成时间"`
	CreatedAt   time.Time

	Model Model `gorm:"foreignKey:ModelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
