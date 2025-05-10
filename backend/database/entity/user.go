package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"size:64;uniqueIndex"`
	PasswordHash string `gorm:"size:128"`
	Role         string `gorm:"size:50;index;comment:用户全局角色"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	// 关联关系
	UserProjectRelations []UserProject `gorm:"foreignKey:UserID"`
	UserRoles            []UserRole    `gorm:"foreignKey:UserID"`
}

// 角色表
type Role struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"size:64;uniqueIndex"`
	Description string `gorm:"size:256"`

	RolePermissionRelations []RolePermission `gorm:"foreignKey:RoleID"`
}

// 权限表
type Permission struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Code        string `gorm:"size:64;uniqueIndex;comment:格式 object:action"`
	Description string `gorm:"size:128"`

	RolePermissionRelations []RolePermission `gorm:"foreignKey:PermissionID"`
}

// 用户-项目关联表（带角色）
type UserProject struct {
	UserID      uint   `gorm:"primaryKey"`
	ProjectID   string `gorm:"primaryKey"`
	ProjectRole string `gorm:"size:50;index;comment:member/admin"`
	JoinedAt    time.Time

	User    User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Project Project `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// 用户-角色关联表
type UserRole struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// 角色-权限关联表
type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`

	Role       Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Permission Permission `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
