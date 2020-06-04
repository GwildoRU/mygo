package flowbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"runtime"
	"strconv"
	"time"
)

type Chat struct {
	id           int64
	uname        string
	bot          *FlowBot
	ch           chan *tgbotapi.Update
	lastSendTime time.Time
	tic          <-chan time.Time
}

func (c *Chat) GetIdStr() string {
	return strconv.Itoa(int(c.id))
}

func (c *Chat) GetUserNameOrId() string {
	s := c.uname
	if s == "" {
		s = strconv.Itoa(int(c.id))
	}
	return s
}

func (c *Chat) SendFile(fileName string) *tgbotapi.Message {
	doc := tgbotapi.NewDocumentUpload(c.id, fileName)
	m, _ := c.bot.Send(doc)
	return &m
}

func (c *Chat) SendMsg(rplcblMsgId int, newText string, newKbrd *Kbrd) *tgbotapi.Message {
	var m tgbotapi.Message
	if rplcblMsgId == 0 {
		msg := tgbotapi.NewMessage(c.id, newText)
		if newKbrd != nil {
			msg.ReplyMarkup = NewKbrd(newKbrd)
		}
		c.lastSendTime = time.Now()
		m, _ = c.bot.Send(msg)
	} else {
		msg := tgbotapi.NewEditMessageText(c.id, rplcblMsgId, newText)
		if newKbrd != nil {
			msg.ReplyMarkup = NewKbrd(newKbrd)
		}
		c.lastSendTime = time.Now()
		m, _ = c.bot.Send(msg)
	}
	return &m
}

func (c *Chat) SendText(rplcblMsgId int, s string) *tgbotapi.Message {
	m := c.SendMsg(rplcblMsgId, s, nil)
	return m
}

func (c *Chat) DelMsg(msgId int) {
	c.bot.DeleteMessage(tgbotapi.NewDeleteMessage(c.id, msgId))
}

func (c *Chat) DelMsgSleep(msgId int, sleepSec int) {
	go func() {
		time.Sleep(time.Duration(sleepSec) * time.Second)
		c.DelMsg(msgId)
	}()
}

func (c *Chat) WaitCallBack(msgId int) (clb *tgbotapi.CallbackQuery) {
	for {
		msg, clb := c.WaitUpdate()
		if clb != nil {
			if clb.Message.MessageID == msgId {
				return clb
			}
			c.DelMsg(clb.Message.MessageID)
		}
		if msg != nil {
			c.DelMsg(msg.MessageID)
		}
	}
}

func (c *Chat) WaitUpdate() (msg *tgbotapi.Message, clb *tgbotapi.CallbackQuery) {
	for {
		select {
		case u, ok := <-c.ch:
			if ok {
				if u.Message != nil {
					return u.Message, nil
				}
				if u.CallbackQuery != nil {
					c.bot.AnswerCallbackQuery(tgbotapi.CallbackConfig{
						CallbackQueryID: u.CallbackQuery.ID,
						Text:            "",
						ShowAlert:       false,
					})
					return nil, u.CallbackQuery
				}
			} else {
				break
			}
		case <-c.tic:
			if int(time.Since(c.lastSendTime).Seconds()) >= c.bot.timeout {
				c.SendText(0, c.bot.timeoutText)
				c.Close()
				runtime.Goexit()
			}
		}
	}
	return nil, nil
}

func (c *Chat) Close() {
	c.bot.chatStore.Del(c.id)
}

func (c *Chat) Prompt(rplcblMsgId int, text string) (int, string) {
	m := c.SendMsg(rplcblMsgId, text, nil)
	s := c.WaitText()
	return m.MessageID, s
}

func (c *Chat) Choice(rplcblMsgId int, text string, newKbrd *Kbrd) (int, string) {
	m := c.SendMsg(rplcblMsgId, text, newKbrd)
	clb := c.WaitCallBack(m.MessageID)
	return m.MessageID, clb.Data
}

func (c *Chat) WaitText() string {
	for {
		msg, clb := c.WaitUpdate()
		if msg != nil {
			if msg.Text != "" {
				return msg.Text
			}
			c.DelMsg(msg.MessageID)
		}
		if clb != nil {
			c.DelMsg(clb.Message.MessageID)
		}
	}
}

func NewChat(newid int64, userName string, b *FlowBot) *Chat {
	chat := &Chat{
		id:    newid,
		uname: userName,
		bot:   b,
		ch:    make(chan *tgbotapi.Update),
		tic:   time.Tick(time.Second),
	}
	b.chatStore.Save(chat.id, chat)
	return chat
}
