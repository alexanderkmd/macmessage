# MacMessage

A simple library to connect to MacOsX chat.db of iMessage/Messages App and read messages from it.

## Connecting to Chat.db

You **have to set permissions** to your terminal or any other app to access the FileSystem.

The way it can be done is described here: [https://spin.atomicobject.com/2020/05/22/search-imessage-sql/]

By default it connects to current user's db i.e. "~/Library/Messages/chat.db"

You can define the exact chat.db  file to connect to:

```golang
ml := macmessage.NewMessagesLoader()
ml.DbPath = "/Path/to/non/default/DB"
ml.ConnectDatabase() // Optional
```

## Reading the messages

Message list is read by DisplayName in iMessage app. The display name can be:

* phone number: "+79871234567"
* some readable name for services: "900" or "MegaFon"

### Read all messages

```golang
ml := macmessage.NewMessagesLoader()
messages, err := ml.GetMessagesByDisplayName("900")
```

### Read messages after date

```golang
ml := macmessage.NewMessagesLoader()

// Set beginning timestamp to 2022-12-09T00:00:00 Local TZ
ml.From = time.Date(2022, 12, 9, 0, 0, 0, 0, time.Local)

// Read the messages after set Timestamp with display name "900"
messages, err := ml.GetMessagesByDisplayName("900")
```
