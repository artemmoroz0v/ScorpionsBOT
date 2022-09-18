package functions

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"strconv"
	s "ScorpionsBOT/structs"
	p "ScorpionsBOT/parse"
	b "ScorpionsBOT/scorpions"
)

func GetUpdates(bot_URL string, offset int) ([]s.Update, error) {
	resp, err := http.Get(bot_URL + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //закрываем тело ответа в конце
	//так как данные с сервера приходят в байтовом формате, то приводим их в читаемый вид
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var answer s.Response
	err = json.Unmarshal(body, &answer)
	if err != nil {
		return nil, err
	}
	return answer.Result, nil
}

func Answer (bot_URL string, update s.Update) (error) {
	ChatId := update.Message.Chat.ChatId
	ClientMessage := update.Message.Text
	if ClientMessage == "/start" {
		var msg s.TextMessage
		msg.ChatId = ChatId
		msg.Text = "Привет, я ScorpionsBOT!\nВведи «/help», чтобы узнать, как со мной нужно обращаться :)"
		buf, err := json.Marshal(msg)
		_, err = http.Post(bot_URL + "/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
	} else if ClientMessage == "/help" {
		var msg s.TextMessage
		msg.ChatId = ChatId
		msg.Text = "Чтобы найти песню вместе с информацией о ней, используй только заглавные буквы, например: Scorpions - Wind Of Change"
		buf, err := json.Marshal(msg)
		_, err = http.Post(bot_URL + "/sendMessage", "application/json", bytes.NewBuffer(buf))
		if err != nil {
			return err
		}
	} else {
		song := p.GetSongName(ClientMessage)
		link, err := p.CreateURL(ClientMessage)
		if err != nil {
			var msg s.TextMessage
			msg.ChatId = ChatId
			msg.Text = err.Error()
			buf, err := json.Marshal(msg)
			_, err = http.Post(bot_URL + "/sendMessage", "application/json", bytes.NewBuffer(buf))
			if err != nil {
				return err
			}
		} else {
			p.GetHTMLCode(link)
			//for sending audio
			var audiomsg s.AudioMessage
			audiomsg.ChatId = ChatId
			audiomsg.Audio = p.GetDownloadLink()
			data, err := json.Marshal(audiomsg)
			if err != nil {
				return err
			}
			//for sending text
			var textmsg s.TextMessage
			textmsg.ChatId = ChatId
			textmsg.Text = b.SongsDataBase()[song].Description
			buf, err := json.Marshal(textmsg)
			//for sending cover
			var photomsg s.PhotoMessage
			photomsg.ChatId = ChatId
			photomsg.Text = b.SongsDataBase()[song].Cover
			photodata, err := json.Marshal(photomsg)
			if err != nil {
				return err
			}
			//sending to client
			_, err = http.Post(bot_URL + "/sendAudio", "application/json", bytes.NewBuffer(data))
			_, err = http.Post(bot_URL + "/sendPhoto", "application/json", bytes.NewBuffer(photodata))
			_, err = http.Post(bot_URL + "/sendMessage", "application/json", bytes.NewBuffer(buf))
		}
	}
	return nil
}
