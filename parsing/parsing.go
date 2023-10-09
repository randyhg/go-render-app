package parsing

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type TopicContentApp struct {
	Id            string `json:"id"`            //排序id
	Type          string `json:"type"`          //类型 1、text文本内容 2、图片链接 3、音频链接 4、视频链接 5、帖子id 6、书籍id 7、@人信息列表
	Uid           string `json:"uid"`           //@人用户uid
	NickName      string `json:"nickName"`      //@人昵称
	ImgUrl        string `json:"imgUrl"`        //图片链接
	VideoUrl      string `json:"videoUrl"`      //视频链接
	AudioUrl      string `json:"audioUrl"`      //音频链接
	EmoId         string `json:"emoId"`         //表情
	TopicId       string `json:"topicId"`       //帖子id
	BookId        string `json:"bookId"`        //书籍id
	Text          string `json:"text"`          //text文本内容
	AttachmentsId string `json:"attachmentsId"` //附件id
	IsSell        string `json:"isSell"`        //是否付费内容 0否1是
}

func (a *TopicContentApp) ToJson() string {
	_j, _ := json.Marshal(a)
	return string(_j)
}

type ContentType struct{}

var contentType = new(ContentType)

func (c *ContentType) Text() string {
	return "1"
}
func (c *ContentType) Image() string {
	return "2"
}
func (c *ContentType) Audio() string {
	return "3"
}
func (c *ContentType) Video() string {
	return "4"
}
func (c *ContentType) Topic() string {
	return "5"
}
func (c *ContentType) Book() string {
	return "6"
}
func (c *ContentType) Mention() string {
	return "7"
}
func (c *ContentType) Emoji() string {
	return "8"
}

var imgRegex = regexp.MustCompile(`<img src="(?P<src>[^"]*)" data-id="(?P<id>\d+)">`)
var audioRegex = regexp.MustCompile(`<audio src="(?P<src>[^"]+)"[^>]*data-id="(?P<id>\d+)"></audio>`)
var videoRegex = regexp.MustCompile(`<video src="(?P<src>[^"]+)"[^>]*data-id="(?P<id>\d+)"></video>`)
var pRegex = regexp.MustCompile(`<p>([^<]+)</p>`)
var emojiRegex = regexp.MustCompile(`\[emoji\](\d+)\[/emoji\]`)
var doorRegex = regexp.MustCompile(`\[door\](\d+)\[/door\]`)
var bookRegex = regexp.MustCompile(`\[book\](\d+)\[/book\]`)
var sellRegex = regexp.MustCompile(`\[sell\](.*)\[/sell\]`)
var mentionRegex = regexp.MustCompile(`<span data-uid="\d+" class="at-box">(@\w+)</span>`)

func main() {
	input := `<html><head></head><body><p><p><br/></p></p>[sell]<img src="" data-id="46">[/sell]<p><p><br/></p></p><span data-uid="168128882007" class="at-box">@hjtestt7</span><p><p><br/></p></p><img src="" data-id="23"></body></html>`
	var p = new(AppParser)
	input = p.TrimContent(input)
	var apps []TopicContentApp
	apps = p.Parse(input, apps)
	if len(apps) >= 1 {
		apps = apps[1:]
	}
	//print apps in json
	var _json []byte
	_json, _ = json.Marshal(apps)
	fmt.Println(string(_json))
}

type AppParser struct {
	CurrentTopicId int
}

func (p *AppParser) NewTopicContentApp(t string, attId string) TopicContentApp {
	p.CurrentTopicId++
	var app TopicContentApp
	app.Id = fmt.Sprintf("%d", p.CurrentTopicId)
	app.Type = t
	app.AttachmentsId = attId
	return app
}

func (p *AppParser) TrimContent(content string) string {
	// Sequentially removing tags
	content = strings.ReplaceAll(content, "<html>", "")
	content = strings.ReplaceAll(content, "</html>", "")
	content = strings.ReplaceAll(content, "<head>", "")
	content = strings.ReplaceAll(content, "</head>", "")
	content = strings.ReplaceAll(content, "<body>", "")
	content = strings.ReplaceAll(content, "</body>", "")
	content = strings.ReplaceAll(content, "<p><p><br/></p></p>", "")
	return content
}

func (p *AppParser) IsImage(content string) bool {
	return imgRegex.MatchString(content)
}

func (p *AppParser) IsVideo(content string) bool {
	return videoRegex.MatchString(content)
}

func (p *AppParser) IsAudio(content string) bool {
	return audioRegex.MatchString(content)
}

func (p *AppParser) IsSell(s string) bool {
	return strings.Contains(s, "[sell]") || strings.Contains(s, "[/sell]")
}

func (p *AppParser) IsEmoji(content string) bool {
	return emojiRegex.MatchString(content)
}

func (p *AppParser) IsDoor(content string) bool {
	return doorRegex.MatchString(content)
}

func (p *AppParser) IsBook(content string) bool {
	return bookRegex.MatchString(content)
}

func (p *AppParser) IsMention(content string) bool {
	return mentionRegex.MatchString(content)
}

//#region extract
func (p *AppParser) GetSellContent(s string) string {
	return p.ExtractContentByTag(s, "[sell]", "[/sell]")
}

func (p *AppParser) GetImageContent(s string) string {
	return p.ExtractContentByTag(s, "<img", ">")
}

func (p *AppParser) GetAudioContent(s string) string {
	return p.ExtractContentByTag(s, "<audio", "</audio>")
}

func (p *AppParser) GetVideoContent(s string) string {
	return p.ExtractContentByTag(s, "<video", "</video>")
}

func (p *AppParser) GetEmojiContent(s string) string {
	return p.ExtractContentByTag(s, "[emoji]", "[/emoji]")
}

func (p *AppParser) GetDoorContent(s string) string {
	return p.ExtractContentByTag(s, "[door]", "[/door]")
}

func (p *AppParser) GetBookContent(s string) string {
	return p.ExtractContentByTag(s, "[book]", "[/book]")
}

func (p *AppParser) GetMentionContent(s string) string {
	return p.ExtractContentByTag(s, "<span", "</span>")
}

func (p *AppParser) GetTextContent(s string) string {
	return p.ExtractContentByTag(s, "<p>", "</p>")
}

func (p *AppParser) ExtractContentByTag(content string, openTag string, closingTag string) string {
	start := strings.Index(content, openTag)
	if start == -1 {
		return ""
	}
	end := strings.Index(content[start:], closingTag)

	if start != -1 && end != -1 {
		return content[start : start+end+len(closingTag)]
	}
	return ""
}

//#endregion

func (p *AppParser) Parse(input string, apps []TopicContentApp) []TopicContentApp {
	if len(input) == 0 {
		return apps
	}
	sell := p.GetSellContent(input)
	// if input has sell then parse sell and input = input - sell, sell may contain all elements
	if sell != "" {
		fmt.Printf("sell: %s\n", sell)
		input = strings.ReplaceAll(input, sell, "")
		sellApps := p.ParseSell(sell)
		apps = append(apps, sellApps)
		return p.Parse(input, apps)
	}
	// if input has video then parse video and input = input - video
	video := p.GetVideoContent(input)
	if video != "" {
		fmt.Printf("video: %s\n", video)
		input = strings.ReplaceAll(input, video, "")
		videoApps := p.ParseVideo(video)
		if videoApps.VideoUrl != "" {
			apps = append(apps, videoApps)
		}
		return p.Parse(input, apps)
	}

	// if input has audio then parse audio and input = input - audio
	audio := p.GetAudioContent(input)
	if audio != "" {
		fmt.Printf("audio: %s\n", audio)
		input = strings.ReplaceAll(input, audio, "")
		audioApps := p.ParseAudio(audio)
		if audioApps.AudioUrl != "" {
			apps = append(apps, audioApps)
		}
		return p.Parse(input, apps)
	}

	// if input has image then parse image and input = input - image
	image := p.GetImageContent(input)
	if image != "" {
		fmt.Printf("image: %s\n", image)
		input = strings.ReplaceAll(input, image, "")
		imageApps := p.ParseImage(image)
		if imageApps.ImgUrl != "" {
			apps = append(apps, imageApps)
		}
		return p.Parse(input, apps)

	}

	// if input has emoji then parse emoji and input = input - emoji
	emoji := p.GetEmojiContent(input)
	if emoji != "" {
		fmt.Printf("emoji: %s\n", emoji)
		input = strings.ReplaceAll(input, emoji, "")
		emojiApps := p.ParseEmoji(emoji)
		if emojiApps.EmoId != "" {
			apps = append(apps, emojiApps)
		}
		return p.Parse(input, apps)
	}

	// if input has door then parse door and input = input - door
	door := p.GetDoorContent(input)
	if door != "" {
		fmt.Printf("door: %s\n", door)
		input = strings.ReplaceAll(input, door, "")
		doorApps := p.ParseDoor(door)
		apps = append(apps, doorApps)
		return p.Parse(input, apps)
	}

	// if input has book then parse book and input = input - book
	book := p.GetBookContent(input)
	if book != "" {
		fmt.Printf("book: %s\n", book)
		input = strings.ReplaceAll(input, book, "")
		bookApps := p.ParseBook(book)
		if bookApps.BookId != "" {
			apps = append(apps, bookApps)
		}
		return p.Parse(input, apps)
	}

	// if input has mention then parse mention and input = input - mention
	mention := p.GetMentionContent(input)
	if mention != "" {
		fmt.Printf("mention: %s\n", mention)
		input = strings.ReplaceAll(input, mention, "")
		mentionApps := p.ParseMention(mention)
		apps = append(apps, mentionApps)
		return p.Parse(input, apps)
	}
	// if input has text then parse text and input = input - text
	text := p.GetTextContent(input)
	if text != "" {
		fmt.Printf("text: %s\n", text)
		input = strings.ReplaceAll(input, text, "")
		textApps := p.ParseText(text)
		apps = append(apps, textApps)
		return p.Parse(input, apps)
	}
	return apps
}

//#region parse
func (p *AppParser) ParseText(content string) TopicContentApp {
	app := p.NewTopicContentApp(contentType.Text(), "")
	content = strings.ReplaceAll(content, "<p>", "")
	content = strings.ReplaceAll(content, "</p>", "")
	app.Text = content
	return app
}
func (p *AppParser) ParseSell(content string) TopicContentApp {
	var app TopicContentApp
	app.Type = contentType.Text()
	app.Text = content
	return app
}

func (p *AppParser) ParseVideo(content string) TopicContentApp {
	app := p.NewTopicContentApp(contentType.Video(), "")
	//extract video src and data-id
	res := videoRegex.FindStringSubmatch(content)
	if len(res) != 3 {
		return app
	}
	app.VideoUrl = res[1]
	app.AttachmentsId = res[2]
	return app
}

func (p *AppParser) ParseAudio(content string) TopicContentApp {
	app := p.NewTopicContentApp(contentType.Audio(), "")
	//extract audio src and data-id
	res := audioRegex.FindStringSubmatch(content)
	if len(res) != 3 {
		return app
	}
	app.AudioUrl = res[1]
	app.AttachmentsId = res[2]
	return app
}

func (p *AppParser) ParseImage(content string) TopicContentApp {
	app := p.NewTopicContentApp(contentType.Image(), "")

	//extract img src and data-id
	res := imgRegex.FindStringSubmatch(content)
	//check if res has 3 elements and assign to app
	if len(res) != 3 {
		return app
	}
	app.ImgUrl = res[1]
	app.AttachmentsId = res[2]

	return app
}

func (p *AppParser) ParseEmoji(content string) TopicContentApp {
	app := p.NewTopicContentApp(contentType.Emoji(), "")
	res := emojiRegex.FindStringSubmatch(content)
	if len(res) != 2 {
		return app
	}
	app.EmoId = res[0]
	return app
}

func (p *AppParser) ParseDoor(content string) TopicContentApp {
	app := p.NewTopicContentApp(contentType.Topic(), "")
	res := doorRegex.FindStringSubmatch(content)
	if len(res) != 2 {
		return app
	}
	app.TopicId = res[0]
	return app
}

func (p *AppParser) ParseBook(content string) TopicContentApp {
	app := p.NewTopicContentApp(contentType.Book(), "")
	res := bookRegex.FindStringSubmatch(content)
	if len(res) != 2 {
		return app
	}
	app.BookId = res[0]
	return app
}

func (p *AppParser) ParseMention(content string) TopicContentApp {
	app := p.NewTopicContentApp(contentType.Mention(), "")
	app.Text = content
	return app
}

//#endregion
