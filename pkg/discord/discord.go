package discord

func NewDiscord(args ...interface{}) *Discord {
	return &Discord{
		args:        args,
		messageChan: make(chan Message, 200),
	}
}