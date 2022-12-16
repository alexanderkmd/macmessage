package macmessage

import (
	"os"
	"path"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Connects to Users database.
// Needs proper security settings to connect.
func init() {
	dirname, _ := os.UserHomeDir()
	ChatDdPath := path.Join(dirname, "Library/Messages/chat.db")
	ConnectDatabase(ChatDdPath)
	log.Warn(ChatDdPath)
}

func TestDBConnetcion(t *testing.T) {
	log.Debug("Testing DB Connection")

	assert.NotNil(t, DB)
}

func TestBasicMessagesSelect(t *testing.T) {

	var message Message
	err := DB.Last(&message).Error
	assert.Nil(t, err)
	log.Debugf("%+v", message)
}

func TestBasicChatSelect(t *testing.T) {

	var chat Chat
	err := DB.Last(&chat).Error
	assert.Nil(t, err)
	log.Debugf("%+v", chat)
}

func TestBasicChatMessageJoinSelect(t *testing.T) {

	var cmj ChatMessageJoin
	err := DB.Last(&cmj).Error
	assert.Nil(t, err)
	log.Debugf("%+v", cmj)
}

func TestBasicHandleSelect(t *testing.T) {

	var handle Handle
	err := DB.Last(&handle).Error
	assert.Nil(t, err)
	log.Debugf("%+v", handle)
}

func TestBasicChatHandleJoinSelect(t *testing.T) {

	var chj ChatHandleJoin
	err := DB.Last(&chj).Error
	assert.Nil(t, err)
	log.Debugf("%+v", chj)
}

func TestGetChatByDisplayName(t *testing.T) {
	var err error
	var handle Handle
	err = DB.Last(&handle).Error
	require.Nil(t, err)

	var chat Chat
	// At first - test record not found
	chat, err = GetChatByDisplayName("zzzzzZZZZZzzzzzzzZZZZZZzzzzzz")
	assert.NotNil(t, err)
	assert.Equal(t, 0, chat.ROWID)
	log.Debugf("%+v", chat)

	// Test - something already in base
	chat, err = GetChatByDisplayName(handle.Uncanonicalized_id)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, chat.ROWID)
	log.Debugf("%+v", chat)
}

func TestGetMessagesByDisplayName(t *testing.T) {
	var err error
	var handle Handle
	err = DB.Last(&handle).Error
	require.Nil(t, err)

	messages, err := GetMessagesByDisplayName(handle.Uncanonicalized_id)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, messages[0].ROWID)
	log.Debugf("%+v", messages)
}
