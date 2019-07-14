package models

import "github.com/jinzhu/gorm"

type (
	LogModel struct {
		gorm.Model
		AppId             string `json:"appId" binding:"required"`
		DeviceId          string `json:"deviceId" binding:"required"`
		SessionId         uint   `json:"sessionId"`
		SessionLastActive int64  `json:"sessionLastActive"`
		SessionPath       string `json:"sessionPath"`
	}

	AppendToLog struct {
		gorm.Model
		SessionId  uint   `json:"sessionId" binding:"required"`
		LogMessage string `json:"message" binding:"required"`
	}
)
