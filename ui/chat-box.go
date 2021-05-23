package ui

import (
	"chat-ui/core/models"
	"fmt"
	"time"
)

type MessageBox struct {
	User       models.User
	MessageStr string
	Time       time.Time
}

func AddMessageBox(box MessageBox) (msg string) {
	msg = fmt.Sprintf("%s %s:\n %s\n", box.Time.Format("2006-01-02 15:04:05"), box.User.Name, box.MessageStr)
	return
}
