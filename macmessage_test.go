package macmessage

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Connects to Users database.
// Needs proper security settings to connect.
func init() {
	ml := NewMessagesLoader()
	err := ml.ConnectDatabase()

	if err != nil {
		log.Error(err)
	}
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

func TestGetHandlesByDisplayName(t *testing.T) {
	var err error
	var handle Handle
	err = DB.Last(&handle).Error
	require.Nil(t, err)

	ml := NewMessagesLoader()

	var handles []Handle
	// At first - test record not found
	handles, err = ml.GetHandlesByKey("zzzzzZZZZZzzzzzzzZZZZZZzzzzzz")
	assert.NotNil(t, err)
	log.Error(err)
	assert.Equal(t, 0, len(handles))
	log.Debugf("%+v", handles)

	// Test - something already in base
	handles, err = ml.GetHandlesByKey(handle.Uncanonicalized_id)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, handles[0].ChatHandleJoin.Chat_id)
	assert.NotEqual(t, 0, handles[0].ChatId())
	log.Debugf("%+v", handles[0])
}

func TestGetMessagesByDisplayName(t *testing.T) {
	var err error
	var handle Handle
	err = DB.Last(&handle).Error
	require.Nil(t, err)

	ml := NewMessagesLoader()

	// At first - test record not found
	messages, err := ml.GetMessagesByDisplayName("zzzzzZZZZZzzzzzzzZZZZZZzzzzzz")
	assert.NotNil(t, err)
	log.Error(err)
	assert.Equal(t, 0, len(messages))
	log.Debugf("%+v", messages)

	messages, err = ml.GetMessagesByDisplayName(handle.Uncanonicalized_id)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, messages[0].ROWID)
	log.Debugf("%+v", messages)
}

func TestGetChatByDisplayName(t *testing.T) {
	var err error
	var handle Handle
	err = DB.Last(&handle).Error
	require.Nil(t, err)

	ml := NewMessagesLoader()
	chats, err := ml.GetChatsByDisplayName(handle.Uncanonicalized_id)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, chats[0].ROWID)
	log.Debugf("%+v", chats)
}
