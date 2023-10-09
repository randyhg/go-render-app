package parsingv2

import (
	"fmt"
	"testing"
)

func TestMediaCombination(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><img src="./hello-world.jpg" data-id="40401"><audio src="./hello-world.mp3" data-id="80401"></audio><img src="./hello-world2.jpg" data-id="40402"><video src="./hello-world.mp4" data-id="98888"></video></body></html>`
	//expected should be 4 with these ordering img, audio, img, video
	got := startParsing(input)
	if len(got) != 4 {
		t.Errorf("got %v want %v", len(got), 4)
	}
	//set expected result
	p := new(AppParser)
	img1 := p.NewTopicContentApp(contentType.Image(), "")
	img1.AttachmentsId = "40401"
	img1.ImgUrl = "./hello-world.jpg"
	audio := p.NewTopicContentApp(contentType.Audio(), "")
	audio.AttachmentsId = "80401"
	audio.AudioUrl = "./hello-world.mp3"
	img2 := p.NewTopicContentApp(contentType.Image(), "")
	img2.AttachmentsId = "40402"
	img2.ImgUrl = "./hello-world2.jpg"
	video := p.NewTopicContentApp(contentType.Video(), "")
	video.AttachmentsId = "98888"
	video.VideoUrl = "./hello-world.mp4"
	final := []TopicContentApp{img1, audio, img2, video}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

func TestCustomTagCombination(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p>[door]444327[/door][book]257[/book]<p><p><br/></p><p><br/></p></p><p>[emoji]002[/emoji][emoji]003[/emoji]</p>[door]444999[/door]</body></html>`
	//expected should be 5 with these ordering door, book, emoji, emoji, door
	got := startParsing(input)
	if len(got) != 5 {
		t.Errorf("got %v want %v", len(got), 5)
	}
	//set expected result
	p := new(AppParser)
	door1 := p.NewTopicContentApp(contentType.Topic(), "")
	door1.TopicId = "[door]444327[/door]"
	book := p.NewTopicContentApp(contentType.Book(), "")
	book.BookId = "[book]257[/book]"
	emoji1 := p.NewTopicContentApp(contentType.Emoji(), "")
	emoji1.EmoId = "[emoji]002[/emoji]"
	emoji2 := p.NewTopicContentApp(contentType.Emoji(), "")
	emoji2.EmoId = "[emoji]003[/emoji]"
	door2 := p.NewTopicContentApp(contentType.Topic(), "")
	door2.TopicId = "[door]444999[/door]"
	final := []TopicContentApp{door1, book, emoji1, emoji2, door2}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

func TestCustomTagAndMediaCombination(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><img src="./hello-world.jpg" data-id="40401">[door]444327[/door]<audio src="./hello-world.mp3" data-id="80401"></audio>[book]257[/book]<p><p><br/></p><p><br/></p></p><p>[emoji]002[/emoji]<video src="./hello-world.mp4" data-id="98888"></video>[emoji]003[/emoji]</p>[door]444999[/door]</body></html>`
	//expected should be 8 with these ordering img, door, audio, book, emoji, video, emoji, door
	got := startParsing(input)
	if len(got) != 8 {
		t.Errorf("got %v want %v", len(got), 8)
	}
	fmt.Printf("got: %v\n", got)
	//set expected result
	p := new(AppParser)
	img := p.NewTopicContentApp(contentType.Image(), "")
	img.AttachmentsId = "40401"
	img.ImgUrl = "./hello-world.jpg"
	door1 := p.NewTopicContentApp(contentType.Topic(), "")
	door1.TopicId = "[door]444327[/door]"
	audio := p.NewTopicContentApp(contentType.Audio(), "")
	audio.AttachmentsId = "80401"
	audio.AudioUrl = "./hello-world.mp3"
	book := p.NewTopicContentApp(contentType.Book(), "")
	book.BookId = "[book]257[/book]"
	emoji1 := p.NewTopicContentApp(contentType.Emoji(), "")
	emoji1.EmoId = "[emoji]002[/emoji]"
	video := p.NewTopicContentApp(contentType.Video(), "")
	video.AttachmentsId = "98888"
	video.VideoUrl = "./hello-world.mp4"
	emoji2 := p.NewTopicContentApp(contentType.Emoji(), "")
	emoji2.EmoId = "[emoji]003[/emoji]"
	door2 := p.NewTopicContentApp(contentType.Topic(), "")
	door2.TopicId = "[door]444999[/door]"
	final := []TopicContentApp{img, door1, audio, book, emoji1, video, emoji2, door2}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}
