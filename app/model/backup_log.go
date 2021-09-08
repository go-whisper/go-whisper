package model

const (
	BackupTypeManual    = "manual"
	BackupTypeAutomatic = "automatic"
)

type BackupLog struct {
	ID        uint
	Type      string
	CloudKey  string
	Others    InterfaceMap
	CreatedAt string `gorm:"time"`
}

func (bl BackupLog) TableName() string {
	return "backup_logs"
}
