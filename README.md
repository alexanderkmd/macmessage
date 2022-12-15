# MacMessage

A simple library to connect to MacOsX chat.db of iMessage App and read messages from it.

```go

err:= macmessage.ConnectDefaultDatabase()
if err != nil{
    log.Fatal(err)
}

ts := time.Date(2022, 12, 9, 0, 0, 0, 0, time.Local)
messages, err := macmessage.GetMessagesByDisplayNameAfterDate("900", ts)

```

## Connecting to Chat.db

You **have to set permissions** to your terminal or any other app to access the FileSystem.

The way it can be done is described here: [https://spin.atomicobject.com/2020/05/22/search-imessage-sql/]

```macmessage.ConnectDefaultDatabase()``` Connects to current user's db i.e. "~/Library/Messages/chat.db"

```macmessage.ConnectDatabase(ChatDBPath string)``` You can define the exact chat.db  file to connect to.

## Reading the messages

Message list is read by DisplayName in iMessage app. The display name can be:

* phone number: "+79871234567"
* some readable name for services: "900" or "MegaFon"

### Read messages after date

```golang
// Set beginning timestamp to 2022-12-09T00:00:00 Local TZ
ts := time.Date(2022, 12, 9, 0, 0, 0, 0, time.Local)

// Read the messages after set Timestamp with display name "900"
messages, err := macmessage.GetMessagesByDisplayNameAfterDate("900", ts)
```

### Read all messages

```golang
macmessage.GetMessagesByDisplayName("900")
```

is fully equivalent to:

```golang
ts := time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)
messages, err := macmessage.GetMessagesByDisplayNameAfterDate("900", ts)
```
