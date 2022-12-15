// Взаимодействие с базой данных сообщений в пограмме
// Messages (Сообщения) в MacOsX
// Формально - это sqlite база в папке ~/Library/Messages/chat.db
//
// Usage:
//
// ChatDbPath := "/Users/USERNAME/Library/Messages/chat.db"
// macmessage.ConnectDatabase(ChatDbPath)
//
// ts := time.Date(2022, 12, 9, 0, 0, 0, 0, time.Local)
//
// messages, err := macmessage.GetMessagesByDisplayNameAfterDate("900", ts)
//
// Что навело на мысль с взаимодействием:
// * https://github.com/golift/imessage/blob/master/incoming.go
// * https://spin.atomicobject.com/2020/05/22/search-imessage-sql/
// В последней так же описано необходимое предоставление прав доступа
package macmessage

import (
	//sqlite "gorm.io/driver/sqlite"
	"os"
	"path"
	"time"

	sqlite "github.com/glebarez/sqlite"
	"gorm.io/gorm"

	// "crawshaw.io/sqlite"
	log "github.com/sirupsen/logrus"
)

var DB *gorm.DB

func ConnectDatabase(ChatDBPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(ChatDBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return err
}

// Connects to the chat.db in Current User's Library
func ConnectDefaultDatabase() error {
	dirname, _ := os.UserHomeDir()
	ChatDdPath := path.Join(dirname, "Library/Messages/chat.db")
	return ConnectDatabase(ChatDdPath)
}

// Returns Chat by Uncanonicalized_id - Display name in Messages App
//
// GetChatByDisplayName("+79871234567")
//
// GetChatByDisplayName("MegaFon")
func GetChatByDisplayName(chatName string) (Chat, error) {
	var err error
	var chat Chat

	if DB == nil {
		return chat, &DbNotConnectedError{}
	}

	var handle Handle
	err = DB.Where(Handle{Uncanonicalized_id: chatName}).First(&handle).Error
	if err != nil {
		return chat, err
	}
	if handle.ROWID == 0 {
		return chat, nil // TODO: как-то обозвать ошибку
	}
	var chj ChatHandleJoin
	err = DB.First(&chj, ChatHandleJoin{Handle_id: handle.ROWID}).Error
	if err != nil {
		return chat, err
	}

	err = DB.First(&chat, Chat{ROWID: chj.Chat_id}).Error
	return chat, err
}

// Returns a list of messages by phone number or
// display name in Messages App
//
// GetMessagesByDisplayName("+79871234567")
//
// GetMessagesByDisplayName("MegaFon")
func GetMessagesByDisplayName(chatName string) ([]Message, error) {
	ts := time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)
	return GetMessagesByDisplayNameAfterDate(chatName, ts)
}

// Returns a list of messages by phone number or
// display name in Messages App after the DateTime
//
// DateTime := time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)
//
// GetMessagesByDisplayNameAfterDate("+79871234567", DateTime)
//
// GetMessagesByDisplayNameAfterDate("MegaFon", DateTime)
func GetMessagesByDisplayNameAfterDate(chatName string, DateTime time.Time) ([]Message, error) {
	var err error
	var messages []Message

	if DB == nil {
		return messages, &DbNotConnectedError{}
	}

	chat, err := GetChatByDisplayName(chatName)
	if err != nil {
		return messages, err
	}
	ts := convertTimeToTimestamp(DateTime)
	err = DB.Joins(
		"JOIN chat_message_join ON chat_message_join.message_id = message.ROWID AND chat_message_join.chat_id=?",
		chat.ROWID).Order("date").Where("date >= ?", ts).Find(&messages).Error
	log.Error(err)

	return messages, err
}
