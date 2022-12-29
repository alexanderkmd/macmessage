// Взаимодействие с базой данных сообщений в пограмме
// Messages (Сообщения) в MacOsX
// Формально - это sqlite база в папке ~/Library/Messages/chat.db
//
// Usage:
//
// ml := macmessage.NewMessagesLoader()
// ml.DbPath = "Path to non default DB"
//
// ts := time.Date(2022, 12, 9, 0, 0, 0, 0, time.Local)
// ml.From = ts
// Default is from 2001-01-01T00:00:00
//
// messages, err := ml.GetMessagesByDisplayName("900")
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
	"gorm.io/gorm/logger"

	// "crawshaw.io/sqlite"
	log "github.com/sirupsen/logrus"
)

var DB *gorm.DB

type MessagesLoader struct {
	Limit  int       // Limit number of messages to return
	From   time.Time // Starting timestamp to return messages from
	DbPath string    // Path to DB to connect to
}

func NewMessagesLoader() *MessagesLoader {
	dirname, _ := os.UserHomeDir()
	return &MessagesLoader{
		Limit:  0,
		From:   time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
		DbPath: path.Join(dirname, "Library/Messages/chat.db"), // Path to default user's DB in MacOsX
	}
}

func (ml MessagesLoader) ConnectDatabase() error {
	var err error
	DB, err = gorm.Open(sqlite.Open(ml.DbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return err
}

// returns timestamp of a From field for Database
func (ml MessagesLoader) timestamp() int {
	return convertTimeToTimestamp(ml.From)
}

// Returns a list of messages by phone number or
// display name in Messages App
//
// GetMessagesByDisplayName("+79871234567")
//
// GetMessagesByDisplayName("MegaFon")
func (ml MessagesLoader) GetMessagesByDisplayName(key string) ([]Message, error) {
	var messages []Message
	var err error

	if DB == nil {
		err = ml.ConnectDatabase()
		if err != nil {
			return messages, err
		}
	}

	handles, err := ml.GetHandlesByKey(key)
	if err != nil {
		return messages, err
	}

	var chatIds []int
	for _, handle := range handles {
		chatIds = append(chatIds, handle.ChatId())
	}

	return ml.GetMessagesByChatIds(chatIds)
}

// Returns all handles and it's chatIDs for the specified key, also looking by person_centric_id
// and Uncanonicalized_id - Display name in Messages App
//
// GetHandlesByDisplayName("+79871234567")
//
// chatId := GetHandlesByDisplayName("MegaFon")[0].ChatHandleJoin.Chat_id
// chatId := GetHandlesByDisplayName("MegaFon")[0].ChatId()
// They are equivalent
func (ml MessagesLoader) GetHandlesByKey(key string) ([]Handle, error) {
	var handle Handle
	var handles []Handle
	var err error

	if DB == nil {
		err = ml.ConnectDatabase()
		if err != nil {
			return handles, err
		}
	}

	// Find any mention of the key in handles table
	err = DB.Where(Handle{Uncanonicalized_id: key}).Or(Handle{Id: key}).First(&handle).Error
	if err != nil {
		log.Debug(err)
		return handles, err
	}

	log.Debugf("%+v", handle)

	// Try to find all the mentions by Person_centric_id or just key variants
	if handle.Person_centric_id != "" {
		//log.Warn(DB.Session(&gorm.Session{DryRun: true}).Where(Handle{Person_centric_id: handle.Person_centric_id}).Joins("ChatHandleJoin").Find(&handles).Statement.SQL.String())
		err = DB.Where(Handle{Person_centric_id: handle.Person_centric_id}).Joins(
			"ChatHandleJoin").Find(&handles).Error
	} else {
		err = DB.Where(Handle{Id: handle.Id}).Joins("ChatHandleJoin").Find(&handles).Error
	}
	if err != nil {
		log.Debug(err)
		return handles, err
	}

	return handles, nil
}

// Get all messages, specified by chat's chatIDs
func (ml MessagesLoader) GetMessagesByChatIds(chatIds []int) ([]Message, error) {
	var messages []Message
	var err error

	if DB == nil {
		err = ml.ConnectDatabase()
		if err != nil {
			return messages, err
		}
	}

	tmp := DB.Joins("JOIN chat_message_join ON chat_message_join.message_id = message.ROWID").Where(
		"chat_message_join.chat_id in ?", chatIds).Where("message_date > ?", ml.timestamp())

	if ml.Limit > 0 {
		// If Limit is passed in config
		tmp.Limit(ml.Limit)
	}

	err = tmp.Find(&messages).Error

	if err != nil {
		log.Error(err)
	}

	log.Info(len(messages))
	if err != nil {
		log.Error(err)
		return messages, err
	}
	return messages, err
}

// Returns Chats (not messages of this chats) by Display name
func (ml MessagesLoader) GetChatsByDisplayName(key string) ([]Chat, error) {
	var err error
	var chats []Chat

	if DB == nil {
		err = ml.ConnectDatabase()
		if err != nil {
			return chats, err
		}
	}

	var handles []Handle
	handles, err = ml.GetHandlesByKey(key)
	if err != nil {
		return chats, err
	}

	var chatIds []int
	for _, handle := range handles {
		chatIds = append(chatIds, handle.ChatId())
	}

	err = DB.Where("ROWID in ?", chatIds).Find(&chats).Error
	return chats, err
}
