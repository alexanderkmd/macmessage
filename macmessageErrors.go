package macmessage

type DbNotConnectedError struct {
}

func (dbnc *DbNotConnectedError) Error() string {
	return "DB not connected"
}
