package models

import (
	"time"
)

// AuthUsers ....
type AuthUser struct {
	Id            int        `json:"user_id" gorm:"column:id;autoincrement;primaryKey"`
	Name          string     `json:"name" gorm:"column:name;type:varchar"`
	Email         *string    `json:"account" gorm:"column:email;type:varchar;unique;check:(email ~* '^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$')"`
	Password      *string    `json:"-" gorm:"column:password;type:varchar"`
	IdNo          *string    `json:"-" gorm:"column:idno;type:varchar;unique;"`
	TpUuid        *string    `json:"-" gorm:"column:uuid;type:varchar;unique;"`
	TpAccount     *string    `json:"TpAccount" gorm:"column:tp_account;type:varchar"`
	TpMemberType  *string    `json:"-" gorm:"column:member_type;type:varchar"`                           // 會員類別
	TpVerifyLevel *string    `json:"-" gorm:"column:verify_level;type:varchar"`                          // 會員驗證等級
	IsAdmin       bool       `json:"is_admin"   gorm:"column:is_admin;type:boolean;default:false"`       // 系統管理者
	IsActive      bool       `json:"is_active" gorm:"column:is_active;type:boolean;default:true"`        // 啟用
	IsWhitelist   bool       `json:"is_whitelist" gorm:"column:is_whitelist;type:boolean;default:false"` // 白名單
	IsBlacked     bool       `json:"is_blacked" gorm:"column:is_blacked;type:boolean;default:false"`     // 黑名單
	ExpiredAt     *time.Time `json:"expired_at" gorm:"column:expired_at;type:timestamp with time zone;"` // 停用時間
	CreatedAt     time.Time  `json:"created_at" gorm:"column:created_at;type:timestamp with time zone;"`
	LoginAt       time.Time  `json:"login_at" gorm:"column:login_at;type:timestamp with time zone;"`
	// Roles       []Role       `json:"roles" gorm:"many2many:email_user_roles;"`
	// Groups      []Group      `json:"groups" gorm:"many2many:email_user_groups;"`
}

// // IssoUsers ....
// type IssoUser struct {
// 	Id          int          `json:"user_id"       gorm:"column:id;autoincrement;primaryKey"`
// 	Uuid        string       `json:"-"             gorm:"column:uuid;type:varchar;unique;"`
// 	IdNo        string       `json:"-"             gorm:"column:idno;type:varchar;unique;"`
// 	Name        string       `json:"name"          gorm:"column:name;type:varchar"`
// 	TpAccount   string       `json:"account"       gorm:"column:tp_account;type:varchar"`
// 	MemberType  string       `json:"-"             gorm:"column:member_type;type:varchar"`                // 會員類別
// 	VerifyLevel string       `json:"-"             gorm:"column:verify_level;type:varchar"`               // 會員驗證等級
// 	IsActive    sql.NullBool `json:"is_active"     gorm:"column:is_active;type:boolean;default:true"`     // 啟用
// 	IsWhitelist sql.NullBool `json:"is_whitelist"  gorm:"column:is_whitelist;type:boolean;default:false"` // 白名單
// 	IsBlacked   sql.NullBool `json:"is_blacked"    gorm:"column:is_blacked;type:boolean;default:false"`   // 黑名單
// 	CreatedAt   time.Time    `json:"created_at"    gorm:"column:created_at;type:timestamp with time zone;"`
// 	LoginAt     time.Time    `json:"login_at"      gorm:"column:login_at;type:timestamp with time zone;"`
// 	Roles       []Role       `json:"roles"         gorm:"many2many:isso_user_roles;"`
// 	Groups      []Group      `json:"groups"        gorm:"many2many:isso_user_groups;"`
// }

// Role represents a user role in the system.
type Role struct {
	Id            int    `json:"role_id"         gorm:"column:id;autoincrement;primaryKey"`
	Name          string `json:"role_name"       gorm:"column:name;type:varchar"`
	AccessControl bool   `json:"access_control"  gorm:"column:access_control;type:boolean;default:false"`
	Modify        bool   `json:"modify"          gorm:"column:modify;type:boolean;default:false"`
	Read          bool   `json:"read"            gorm:"column:read;type:boolean;default:false"`
}

// Groups ....
type Group struct {
	Id         int      `json:"group_id"    gorm:"column:id;autoincrement;primaryKey"`
	Name       string   `json:"group_name"  gorm:"column:name;type:varchar"`
	IsPersonal bool     `json:"is_personal" gorm:"column:is_personal;type:boolean;default:false"`
	CreateBy   int      `json:"create_by"   gorm:"column:create_by;"`
	AuthUser   AuthUser `gorm:"foreignKey:create_by"`
}

type AuthUserGroupRole struct {
	AuthUserID int      `json:"user_id"     gorm:"column:auth_user_id;primaryKey"`
	GroupID    int      `json:"group_id"    gorm:"column:group_id;primaryKey;"`
	RoleID     int      `json:"role_id"     gorm:"column:role_id;primaryKey"`
	AuthUser   AuthUser `gorm:"foreignKey:AuthUserID"`
	Group      Group    `gorm:"foreignKey:GroupID"`
	Role       Role     `gorm:"foreignKey:RoleID"`
}

// // EmailUserRole represents the relationship between email users and roles.
// type EmailUserRole struct {
// 	UserID    int       `json:"user_id"     gorm:"column:email_user_id;primaryKey"`
// 	RoleID    int       `json:"role_id"     gorm:"column:role_id;primaryKey"`
// 	EmailUser EmailUser `gorm:"foreignKey:UserID"`
// 	Role      Role      `gorm:"foreignKey:RoleID"`
// }

// // EmailUserGroup represents the relationship between email users and groups.
// type EmailUserGroup struct {
// 	UserID    int       `json:"user_id"   gorm:"column:email_user_id;primaryKey;"`
// 	GroupID   int       `json:"group_id"  gorm:"column:group_id;primaryKey;"`
// 	EmailUser EmailUser `gorm:"foreignKey:UserID"`
// 	Group     Group     `gorm:"foreignKey:GroupID"`
// }

// // IssoUserRole represents the relationship between isso users and roles.
// type IssoUserRole struct {
// 	UserID   int      `json:"user_id"     gorm:"column:isso_user_id;primaryKey"`
// 	RoleID   int      `json:"role_id"     gorm:"column:role_id;primaryKey"`
// 	IssoUser IssoUser `gorm:"foreignKey:UserID"`
// 	Role     Role     `gorm:"foreignKey:RoleID"`
// }

// // IssoUserGroup represents the relationship between isso users and groups.
// type IssoUserGroup struct {
// 	UserID   int      `json:"user_id"   gorm:"column:isso_user_id;primaryKey;"`
// 	GroupID  int      `json:"group_id"  gorm:"column:group_id;primaryKey;"`
// 	IssoUser IssoUser `gorm:"foreignKey:UserID"`
// 	Group    Group    `gorm:"foreignKey:GroupID"`
// }
