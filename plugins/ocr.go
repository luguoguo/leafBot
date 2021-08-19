package plugins

import (
	"github.com/3343780376/leafBot"
	"github.com/3343780376/leafBot/message"
)

func Ocr() {
	leafBot.OnCommand("/ocr").SetPluginName("图片ocr").SetBlock(false).SetWeight(10).AddRule(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) bool {
		for _, mess := range event.Message {
			if mess.Type == "image" {
				return true
			}
		}
		return false
	}).AddHandle(func(event leafBot.Event, bot *leafBot.Bot, state *leafBot.State) {
		images := event.GetImages()
		if len(images) < 1 {
			return
		}
		ocrImage := bot.OcrImage(images[0].Data["file"])
		mess := "识别结果:\n识别语言:" + ocrImage.Language
		for _, text := range ocrImage.Texts {
			mess += "\n" + text.Text

		}

		bot.Send(event, message.ReplyWithMessage(int64(event.MessageID), message.Text(mess)))
	})
}
