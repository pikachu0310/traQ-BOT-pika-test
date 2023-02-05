package commands

type Args struct {
	Slice     []string
	MessageID string
	ChannelID string
	UserID    string
}

type ArgsV2 struct {
	MessageText string
	MessageID   string
	ChannelID   string
	UserID      string
}
