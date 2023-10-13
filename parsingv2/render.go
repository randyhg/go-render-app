package parsingv2

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/net/html"
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
	IsCensored    string `json:"isCensored"`    //是否需要隐藏 0否1是
}

func (a *TopicContentApp) ToJson() string {
	// _j, _ := json.Marshal(a)
	_j, _ := json.MarshalIndent(a, "", "    ")
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

type AppParser struct {
	CurrentTopicId int
	IsSell         bool
}

func (p *AppParser) NewTopicContentApp(t string, attId string) TopicContentApp {
	p.CurrentTopicId++
	var app TopicContentApp
	app.Id = fmt.Sprintf("%d", p.CurrentTopicId)
	app.Type = t
	app.AttachmentsId = attId
	return app
}

// var emojiRegex = regexp.MustCompile(`\[emoji\](\d+)\[/emoji\]`)
// var doorRegex = regexp.MustCompile(`\[door\](\d+)\[/door\]`)
// var bookRegex = regexp.MustCompile(`\[book\](\d+)\[/book\]`)
var tagRegex = regexp.MustCompile(`\[(?:(door|book|emoji|sell))\](\d+)\[\/(?:door|book|emoji|sell)\]`)

// buat regex untuk [sell]

func (p *AppParser) Render(input string) (apps []TopicContentApp) {
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}
	var contentApps []TopicContentApp
	var f func(*html.Node)
	f = func(n *html.Node) {
		// handle all media type, img, audio, video and not in p tag
		if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "audio" || n.Data == "video") && n.Parent.Data != "p" {
			contentApps = p.parseMedia(n, contentApps)
		}
		// handle mention and not in p tag
		if n.Type == html.ElementNode && n.Data == "span" && n.Parent.Data != "p" {
			contentApps = p.parseMention(n, contentApps)
		}
		// handle all p tag and its children
		if n.Type == html.ElementNode && n.Data == "p" {
			text := ""
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.ElementNode && (c.Data == "img" || c.Data == "audio" || c.Data == "video") {
					contentApps = p.parseMedia(c, contentApps)
				}
				if c.Type == html.ElementNode && c.Data == "span" {
					contentApps = p.parseMention(n, contentApps)
				}
				if c.Type == html.TextNode && c.Data != "" {
					text = c.Data
					if !tagRegex.MatchString(text) {
						// If there are no tags, add the entire text node
						if strings.Contains(text, "[sell]") {
							// openBracket := strings.Index(text, "[")
							prefix := text[:6]
							suffix := text[6:]
							if strings.Contains(suffix, "[/sell]") {
								suffix = strings.ReplaceAll(suffix, "[/sell]", "")
								suffix2 := "[/sell]"
								slice1 := []string{prefix, suffix, suffix2}
								for _, i := range slice1 {
									contentApps = p.parseParagraph(i, contentApps)
								}
							} else {
								slice1 := []string{prefix, suffix}
								for _, i := range slice1 {
									contentApps = p.parseParagraph(i, contentApps)
								}
							}
						} else {
							contentApps = p.parseParagraph(text, contentApps)
						}
					} else {
						locs := tagRegex.FindAllStringIndex(text, -1)
						fmt.Printf("locs is %v\n", locs)
						start := 0
						for _, loc := range locs {
							if loc[0] > start {
								// Add text node
								contentApps = p.parseParagraph(text[start:loc[0]], contentApps)
							}
							tags := text[loc[0]:loc[1]]
							fmt.Printf("from paragraph and tags is %v\n", tags)
							contentApps = p.parseCustomTags(tags, contentApps)
							start = loc[1]
						}
					}
				}
			}
		}

		// handle all text node not in p tag
		if n.Type == html.TextNode && n.Parent.Data != "p" {
			fmt.Printf("Sell should come here\n")
			fmt.Printf("current content is %v\n", n.Data)
			contentApps = p.parseCustomTags(n.Data, contentApps)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return contentApps
}

func (p *AppParser) parseCustomTags(content string, contentApps []TopicContentApp) []TopicContentApp {
	fmt.Printf("Parse custom tags ya where content is %v\n", content)
	matches := tagRegex.FindAllStringSubmatch(content, -1)
	if content == "[sell]" {
		// create sell open tag for content app
		openSell := p.NewTopicContentApp(contentType.Text(), "")
		openSell.IsSell = "1"
		openSell.Text = "[sell]"
		openSell.IsCensored = "1"
		p.IsSell = true
		return contentApps
	}
	if content == "[/sell]" {
		// create sell close tag for content app
		closeSell := p.NewTopicContentApp(contentType.Text(), "")
		closeSell.IsSell = "1"
		closeSell.Text = "[/sell]"
		closeSell.IsCensored = "1"
		p.IsSell = false
		return contentApps
	}
	for _, match := range matches {
		if len(match) > 2 {
			fmt.Printf("current match is %v\n", match)
			switch match[1] {
			case "door":
				contentApps = p.parseDoor(contentApps, match)
			case "book":
				contentApps = p.parseBook(contentApps, match)
			case "emoji":
				contentApps = p.parseEmoji(contentApps, match)
			default:
				continue
			}
		}
	}

	return contentApps
}

func (p *AppParser) parseDoor(contentApps []TopicContentApp, matches []string) []TopicContentApp {
	app := p.NewTopicContentApp(contentType.Topic(), "")
	app.TopicId = matches[0]
	//app.Text = matches[0]
	if p.IsSell {
		app.IsSell = "1"
		app.IsCensored = "1"
	}
	contentApps = append(contentApps, app)
	return contentApps
}

func (p *AppParser) parseEmoji(contentApps []TopicContentApp, matches []string) []TopicContentApp {
	fmt.Printf("matches: %v\n", matches)
	app := p.NewTopicContentApp(contentType.Emoji(), "")
	app.EmoId = matches[0]
	if p.IsSell {
		app.IsSell = "1"
		app.IsCensored = "1"
	}
	contentApps = append(contentApps, app)
	return contentApps
}

func (p *AppParser) parseBook(contentApps []TopicContentApp, matches []string) []TopicContentApp {
	app := p.NewTopicContentApp(contentType.Book(), "")
	app.BookId = matches[0]
	if p.IsSell {
		app.IsSell = "1"
		app.IsCensored = "1"
	}
	contentApps = append(contentApps, app)
	return contentApps
}

func (p *AppParser) parseMedia(n *html.Node, contentApps []TopicContentApp) []TopicContentApp {
	var src, id string
	for _, a := range n.Attr {
		if a.Key == "src" {
			src = a.Val
		} else if a.Key == "data-id" {
			id = a.Val
		}
	}
	contentApp := p.NewTopicContentApp("", id)
	switch n.Data {
	case "img":
		contentApp.Type = contentType.Image()
		contentApp.ImgUrl = src
		contentApp.AttachmentsId = id
	case "audio":
		contentApp.Type = contentType.Audio()
		contentApp.AudioUrl = src
		contentApp.AttachmentsId = id
	case "video":
		contentApp.Type = contentType.Video()
		contentApp.VideoUrl = src
		contentApp.AttachmentsId = id
	}
	if p.IsSell {
		contentApp.IsSell = "1"
		contentApp.IsCensored = "1"
	}
	contentApps = append(contentApps, contentApp)
	return contentApps
}

func (p *AppParser) parseMention(n *html.Node, contentApps []TopicContentApp) []TopicContentApp {
	fmt.Printf("Coming inside parseMention \n")
	var uid, class, innerText string
	for _, a := range n.Attr {
		if a.Key == "data-uid" {
			uid = a.Val
		} else if a.Key == "class" && a.Val == "at-box" {
			class = a.Val
		}
	}
	if class == "at-box" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				innerText = c.Data
				break
			}
		}
		contentApp := p.NewTopicContentApp(contentType.Mention(), "")
		contentApp.Uid = uid
		contentApp.NickName = innerText
		if p.IsSell {
			contentApp.IsSell = "1"
			contentApp.IsCensored = "1"
		}
		contentApps = append(contentApps, contentApp)
	}
	return contentApps
}

func (p *AppParser) parseParagraph(innerText string, contentApps []TopicContentApp) []TopicContentApp {
	fmt.Printf("parseParagraph: %v\n", innerText)
	contentApp := p.NewTopicContentApp(contentType.Text(), "")
	contentApp.Text = innerText
	if p.IsSell {
		contentApp.IsSell = "1"
		contentApp.IsCensored = "1"
	}
	contentApps = append(contentApps, contentApp)
	return contentApps
}
