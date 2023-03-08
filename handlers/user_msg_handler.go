package handlers

import (
	"github.com/869413421/wechatbot/gtp"
	"github.com/eatmoreapple/openwechat"
	"log"
	"strings"
)

var _ MessageHandlerInterface = (*UserMessageHandler)(nil)

// UserMessageHandler 私聊消息处理
type UserMessageHandler struct {
}

// handle 处理消息
func (g *UserMessageHandler) handle(msg *openwechat.Message) error {
	if msg.IsText() {
		return g.ReplyText(msg)
	}
	return nil
}

// NewUserMessageHandler 创建私聊处理器
func NewUserMessageHandler() MessageHandlerInterface {
	return &UserMessageHandler{}
}

// ReplyText 发送文本消息到群
func (g *UserMessageHandler) ReplyText(msg *openwechat.Message) error {
	// 接收私聊消息
	sender, err := msg.Sender()
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)
	//if UserService.ClearUserSessionContext(sender.ID(), msg.Content) {
	//	_, err = msg.ReplyText("上下文已经清空了，你可以问下一个问题啦。")
	//	if err != nil {
	//		log.Printf("response user error: %v \n", err)
	//	}
	//	return nil
	//}

	// 获取上下文，向GPT发起请求
	requestText := strings.TrimSpace(msg.Content)
	requestText = strings.Trim(msg.Content, "\n")
	requestText = strings.ReplaceAll(requestText, "\n", " ")
	requestText = strings.ReplaceAll(requestText, "\r", " ")

	//requestText = UserService.GetUserSessionContext(sender.ID()) + requestText
	reply, err := gtp.Completions(requestText)
	if err != nil {
		log.Printf("gtp request error: %v \n", err)
		msg.ReplyText("提问频繁，请稍后在试！")
		return err
	}
	if reply == "" {
		return nil
	}

	// 设置上下文，回复用户
	reply = strings.TrimSpace(reply)
	reply = strings.Trim(reply, "\n")
	//UserService.SetUserSessionContext(sender.ID(), requestText, reply)
	//reply = "本消息由 chatGPT Bot回复：\n" + reply
	_, err = msg.ReplyText(reply)
	if err != nil {
		log.Printf("response user error: %v \n", err)
	}
	return err
}
