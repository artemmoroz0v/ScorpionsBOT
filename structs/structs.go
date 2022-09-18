package structs

type Update struct {
	UpdateID int `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	Chat Chat `json:"chat"`
	Text string `json:"text"`
}

type Chat struct {
	ChatId int `json:"id"`
}

type Response struct {
	Result []Update `json:"result"`
}

type AudioMessage struct {
	ChatId int `json:"chat_id"`
	Audio string `json:"audio"`
}

type PhotoMessage struct {
	ChatId int `json:"chat_id"`
	Text string `json:"photo"`
}

type TextMessage struct {
	ChatId int `json:"chat_id"`
	Text string `json:"text"`
}
