package model

import "time"

const (
	BackupTypeManual    = "manual"
	BackupTypeAutomatic = "automatic"
)

type BackupLog struct {
	ID        uint
	Type      string
	CloudKey  string
	Others    InterfaceMap
	CreatedAt time.Time
}

func (bl BackupLog) TableName() string {
	return "backup_logs"
}
