package entity

type Dataset struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(64);not null;comment:数据集名称"`
	Description string `gorm:"type:text;comment:数据集描述"`
	Path        string `gorm:"type:varchar(512);comment:数据集存储路径"`
	Size        int64  `gorm:"comment:数据集大小 (bytes)"`
	Format      string `gorm:"type:varchar(50);comment:数据集格式 (e.g., CSV, JSON, images)"`
	Version     string `gorm:"type:varchar(50);comment:数据集版本"`
	MetaData    string `gorm:"type:jsonb;comment:其他元数据"`
	UserID      uint   `gorm:"comment:创建数据集的用户ID"`
	User        User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
