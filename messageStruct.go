// Схема данных таблицы сообщений в chat.db
// Вязта из репо: https://github.com/golift/imessage/blob/master/message_schema.txt
package macmessage

import (
	"time"
)

type Tabler interface {
	TableName() string
}

type Chat struct {
	ROWID                          int    `gorm:"column:ROWID;primaryKey"` // INTEGER PRIMARY KEY AUTOINCREMENT,
	Guid                           string // TEXT UNIQUE NOT NULL,
	Style                          int    //INTEGER,
	State                          int    // INTEGER,
	Account_id                     string //TEXT,
	Properties                     string // BLOB,
	Chat_identifier                string // TEXT,
	Service_name                   string //TEXT,
	Room_name                      string // TEXT,
	Account_login                  string // TEXT,
	Is_archived                    bool   // INTEGER DEFAULT 0,
	Last_addressed_handle          string //TEXT,
	Display_name                   string // TEXT,
	Group_id                       string //TEXT,
	Is_filtered                    int    // INTEGER,
	Successful_query               int    //INTEGER,
	Engram_id                      string //TEXT,
	Server_change_token            string //TEXT,
	Ck_sync_state                  int    //INTEGER DEFAULT 0,
	Original_group_id              string //TEXT,
	Last_read_message_timestamp    int    //INTEGER DEFAULT 0,
	Ck_record_system_property_blob string // BLOB,
	Sr_server_change_token         string //TEXT,
	Sr_ck_sync_state               int    //INTEGER DEFAULT 0,
	Cloudkit_record_id             string //TEXT DEFAULT NULL,
	Sr_cloudkit_record_id          string //TEXT DEFAULT NULL,
	Last_addressed_sim_id          string //TEXT DEFAULT NULL,
	Is_blackholed                  int    //INTEGER DEFAULT 0
}

func (Chat) TableName() string {
	return "chat"
}

type Message struct {
	ROWID                             int    `gorm:"column:ROWID;primaryKey"` // INTEGER PRIMARY KEY AUTOINCREMENT,
	Guid                              string // TEXT UNIQUE NOT NULL,
	Text                              string // TEXT,
	Replace                           bool   // INTEGER DEFAULT 0,
	Service_center                    string // TEXT,
	Handle_id                         int    // INTEGER DEFAULT 0,
	Subject                           string // TEXT,
	Country                           string // TEXT,
	attributedBody                    string // BLOB,
	Version                           int    // INTEGER DEFAULT 0,
	Type                              int    // INTEGER DEFAULT 0,
	Service                           string // TEXT,
	Account                           string // TEXT,
	Account_guid                      string // TEXT,
	Error                             int    // INTEGER DEFAULT 0,
	Date                              int    // INTEGER,
	Date_read                         int    // INTEGER,
	Date_delivered                    int    // INTEGER,
	Is_delivered                      int    // INTEGER DEFAULT 0,
	Is_finished                       int    // INTEGER DEFAULT 0,
	Is_emote                          int    // INTEGER DEFAULT 0,
	Is_from_me                        int    // INTEGER DEFAULT 0,
	Is_empty                          int    // INTEGER DEFAULT 0,
	Is_delayed                        int    // INTEGER DEFAULT 0,
	Is_auto_reply                     int    // INTEGER DEFAULT 0,
	Is_prepared                       int    // INTEGER DEFAULT 0,
	Is_read                           int    // INTEGER DEFAULT 0,
	Is_system_message                 int    // INTEGER DEFAULT 0,
	Is_sent                           int    // INTEGER DEFAULT 0,
	Has_dd_results                    int    // INTEGER DEFAULT 0,
	Is_service_message                int    // INTEGER DEFAULT 0,
	Is_forward                        int    // INTEGER DEFAULT 0,
	Was_downgraded                    int    // INTEGER DEFAULT 0,
	Is_archive                        int    // INTEGER DEFAULT 0,
	Cache_has_attachments             int    // INTEGER DEFAULT 0,
	Cache_roomnames                   string // TEXT,
	Was_data_detected                 int    // INTEGER DEFAULT 0,
	Was_deduplicated                  int    // INTEGER DEFAULT 0,
	Is_audio_message                  int    // INTEGER DEFAULT 0,
	Is_played                         int    // INTEGER DEFAULT 0,
	Date_played                       int    // INTEGER,
	Item_type                         int    // INTEGER DEFAULT 0,
	Other_handle                      int    // INTEGER DEFAULT 0,
	Group_title                       string // TEXT,
	Group_action_type                 int    // INTEGER DEFAULT 0,
	Share_status                      int    // INTEGER DEFAULT 0,
	Share_direction                   int    // INTEGER DEFAULT 0,
	Is_expirable                      int    // INTEGER DEFAULT 0,
	Expire_state                      int    // INTEGER DEFAULT 0,
	Message_action_type               int    // INTEGER DEFAULT 0,
	Message_source                    int    // INTEGER DEFAULT 0,
	Associated_message_guid           string // TEXT,
	Associated_message_type           int    // INTEGER DEFAULT 0,
	Balloon_bundle_id                 string // TEXT,
	Payload_data                      string // BLOB,
	Expressive_send_style_id          string // TEXT,
	Associated_message_range_location int    // INTEGER DEFAULT 0,
	Associated_message_range_length   int    // INTEGER DEFAULT 0,
	Time_expressive_send_played       int    // INTEGER,
	Message_summary_info              string // BLOB,
	Ck_sync_state                     int    // INTEGER DEFAULT 0,
	Ck_record_id                      string // TEXT DEFAULT NULL,
	Ck_record_change_tag              string // TEXT DEFAULT NULL,
	Destination_caller_id             string // TEXT DEFAULT NULL,
	Sr_ck_sync_state                  int    // INTEGER DEFAULT 0,
	Sr_ck_record_id                   string // TEXT DEFAULT NULL,
	Sr_ck_record_change_tag           string // TEXT DEFAULT NULL
}

// TableName overrides the table name used by User to `profiles`
func (Message) TableName() string {
	return "message"
}

// Returns message Date in go Time format
func (m *Message) DateTime() time.Time {
	return convertTimestampToTime(m.Date)
}

type ChatMessageJoin struct {
	Chat_id      int // INTEGER REFERENCES chat (ROWID) ON DELETE CASCADE,
	Message_id   int // INTEGER REFERENCES message (ROWID) ON DELETE CASCADE,
	Message_date int // INTEGER DEFAULT 0,
	// PRIMARY KEY (chat_id, message_id)
}

func (ChatMessageJoin) TableName() string {
	return "chat_message_join"
}

type Handle struct {
	ROWID              int            `gorm:"column:ROWID;primaryKey"` //  INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
	Id                 string         // TEXT NOT NULL,
	Country            string         // TEXT,
	Service            string         // TEXT NOT NULL,
	Uncanonicalized_id string         // TEXT,
	Person_centric_id  string         // TEXT DEFAULT NULL
	ChatHandleJoin     ChatHandleJoin `gorm:"foreignKey:Handle_id"`
}

func (Handle) TableName() string {
	return "handle"
}

// Returns chat_id for the specified handle
func (h Handle) ChatId() int {
	return h.ChatHandleJoin.Chat_id
}

type ChatHandleJoin struct {
	Chat_id   int // INTEGER REFERENCES chat (ROWID) ON DELETE CASCADE,
	Handle_id int //INTEGER REFERENCES handle (ROWID) ON DELETE CASCADE
}

func (ChatHandleJoin) TableName() string {
	return "chat_handle_join"
}
