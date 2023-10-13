package parsingv2

import (
	"testing"
)

// Failed
// got 2 want 4
func TestSellWithParagraph(t *testing.T) {
	input := `<p>下午好</p><p>[sell]你们在干嘛啊[/sell]</p>`
	got := startParsing(input)
	if len(got) != 4 {
		t.Errorf("got %v want %v", len(got), 4)
	}

	//expected
	p := new(AppParser)
	paragraph1 := p.NewTopicContentApp(contentType.Text(), "")
	paragraph1.Text = "下午好"
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.Text = "[sell]"
	paragraph2 := p.NewTopicContentApp(contentType.Text(), "")
	paragraph2.Text = "你们在干嘛啊"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.Text = "[/sell]"
	final := []TopicContentApp{paragraph1, openSell, paragraph2, closeSell}

	// compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
		// t.Log(got[i].ToJson())
	}
}

// Failed
// got 5 want 6
func TestSellWithParagraphAndImg(t *testing.T) {
	input := `<p>下午好</p><p>[sell]大家今天吃点啥呢</p><p><img src="https://res.hjcaecf.top/hjstore/images/20230911/e1b4561c88df36d1e3caa23f5dc9e3fd.jpg.txt?Expires=1694489125&Signature=QUtZ72xPPyUPXiV0-9g2ltmtGczHRCYfysKjK6~gGxRknlhp5LVDtdkBNxEIoTb9XY6qyQqknYtUeeRNYBjnJ-uEiIefJsbqOJtxwLRySP1SqGFV8Sl~pneaHkC5gJXzZ2NH5GxF6JsNXyAN4VfqHJIMJj0ZM7zUxyiSilceqv6-vTs3ZACMYGXaxG4-N1SKcnh9yo97mCWd9SaMl2AVDjl2w2cEySijoqu5iPooUhJFfhC17WNWMN~eYjLE6j8kK6pb4YwHDc5K-7zE2EP2AzBo6lZb~r8WXl1Mb97XmgpFo1vN9QcZbiUtjfK694J8R9nD7aThjrJg-b35jFrY4A__&Key-Pair-Id=K3HEWB9S9B6R6N" data-id="227">[/sell]</p><p><img src="https://res.hjcaecf.top/hjstore/images/20230911/e1b4561c88df36d1e3caa23f5dc9e3fd.jpg.txt?Expires=1694489125&Signature=QUtZ72xPPyUPXiV0-9g2ltmtGczHRCYfysKjK6~gGxRknlhp5LVDtdkBNxEIoTb9XY6qyQqknYtUeeRNYBjnJ-uEiIefJsbqOJtxwLRySP1SqGFV8Sl~pneaHkC5gJXzZ2NH5GxF6JsNXyAN4VfqHJIMJj0ZM7zUxyiSilceqv6-vTs3ZACMYGXaxG4-N1SKcnh9yo97mCWd9SaMl2AVDjl2w2cEySijoqu5iPooUhJFfhC17WNWMN~eYjLE6j8kK6pb4YwHDc5K-7zE2EP2AzBo6lZb~r8WXl1Mb97XmgpFo1vN9QcZbiUtjfK694J8R9nD7aThjrJg-b35jFrY4A__&Key-Pair-Id=K3HEWB9S9B6R6N" data-id="227"></p>`
	got := startParsing(input)
	if len(got) != 6 {
		t.Errorf("got %v want %v", len(got), 6)
	}

	// expected
	p := new(AppParser)
	paragraph1 := p.NewTopicContentApp(contentType.Text(), "")
	paragraph1.Text = "下午好"
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.Text = "[sell]"
	paragraph2 := p.NewTopicContentApp(contentType.Text(), "")
	paragraph2.Text = "大家今天吃点啥呢"
	img1 := p.NewTopicContentApp(contentType.Image(), "")
	img1.ImgUrl = "https://res.hjcaecf.top/hjstore/images/20230911/e1b4561c88df36d1e3caa23f5dc9e3fd.jpg.txt?Expires=1694489125&Signature=QUtZ72xPPyUPXiV0-9g2ltmtGczHRCYfysKjK6~gGxRknlhp5LVDtdkBNxEIoTb9XY6qyQqknYtUeeRNYBjnJ-uEiIefJsbqOJtxwLRySP1SqGFV8Sl~pneaHkC5gJXzZ2NH5GxF6JsNXyAN4VfqHJIMJj0ZM7zUxyiSilceqv6-vTs3ZACMYGXaxG4-N1SKcnh9yo97mCWd9SaMl2AVDjl2w2cEySijoqu5iPooUhJFfhC17WNWMN~eYjLE6j8kK6pb4YwHDc5K-7zE2EP2AzBo6lZb~r8WXl1Mb97XmgpFo1vN9QcZbiUtjfK694J8R9nD7aThjrJg-b35jFrY4A__&Key-Pair-Id=K3HEWB9S9B6R6N"
	img1.AttachmentsId = "227"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.Text = "[/sell]"
	img2 := p.NewTopicContentApp(contentType.Image(), "")
	img2.ImgUrl = "https://res.hjcaecf.top/hjstore/images/20230911/e1b4561c88df36d1e3caa23f5dc9e3fd.jpg.txt?Expires=1694489125&Signature=QUtZ72xPPyUPXiV0-9g2ltmtGczHRCYfysKjK6~gGxRknlhp5LVDtdkBNxEIoTb9XY6qyQqknYtUeeRNYBjnJ-uEiIefJsbqOJtxwLRySP1SqGFV8Sl~pneaHkC5gJXzZ2NH5GxF6JsNXyAN4VfqHJIMJj0ZM7zUxyiSilceqv6-vTs3ZACMYGXaxG4-N1SKcnh9yo97mCWd9SaMl2AVDjl2w2cEySijoqu5iPooUhJFfhC17WNWMN~eYjLE6j8kK6pb4YwHDc5K-7zE2EP2AzBo6lZb~r8WXl1Mb97XmgpFo1vN9QcZbiUtjfK694J8R9nD7aThjrJg-b35jFrY4A__&Key-Pair-Id=K3HEWB9S9B6R6N"
	img2.AttachmentsId = "227"
	final := []TopicContentApp{paragraph1, openSell, paragraph2, img1, closeSell, img2}

	// compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
		// t.Log(got[i].ToJson())
	}
}

// Failed
// got 6 want 7
func TestSellWithParagraphImgAndVideo(t *testing.T) {
	input := `<p>武安那你</p><p>[sell]你好<img src="img1.jpg" data-id="227"></p><p><br></p><p>[/sell]</p><p><img src="img2.jpg" data-id="227"></p><p><video src="vid1.mp4" data-id="229"></video></p>`
	got := startParsing(input)
	if len(got) != 7 {
		t.Errorf("got %v want %v", len(got), 7)
	}

	// expected
	p := new(AppParser)
	paragraph1 := p.NewTopicContentApp(contentType.Text(), "")
	paragraph1.Text = "武安那你"
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.Text = "[sell]"
	paragraph2 := p.NewTopicContentApp(contentType.Text(), "")
	paragraph2.Text = "你好"
	img1 := p.NewTopicContentApp(contentType.Image(), "")
	img1.ImgUrl = "img1.jpg"
	img1.AttachmentsId = "227"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.Text = "[/sell]"
	img2 := p.NewTopicContentApp(contentType.Image(), "")
	img2.ImgUrl = "img2.jpg"
	img2.AttachmentsId = "227"
	vid := p.NewTopicContentApp(contentType.Video(), "")
	vid.VideoUrl = "vid1.mp4"
	vid.AttachmentsId = "229"
	final := []TopicContentApp{paragraph1, openSell, paragraph2, img1, closeSell, img2, vid}

	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v, want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

// Failed
// got 7 want 8
func TestSellWithParagraphImgVidAndAudio(t *testing.T) {
	input := `<p>你好你好</p><p>[sell]测试测试</p><p><img src="img1.jpg" data-id="227"></p><audio src="audio1.mp3" controls="true" controlslist="nodownload" id="showaudio" data-id="217"></audio><p>[/sell]</p><p><img src="img2.jpg" data-id="227"></p><p><video src="video1.mp4" data-id="229"></video></p>`
	got := startParsing(input)
	if len(got) != 8 {
		t.Errorf("got %v want %v", len(got), 8)
	}

	// expected
	p := new(AppParser)
	paragraph1 := p.NewTopicContentApp(contentType.Text(), "")
	paragraph1.Text = "你好你好"
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.Text = "[sell]"
	paragraph2 := p.NewTopicContentApp(contentType.Text(), "")
	paragraph2.Text = "测试测试"
	img1 := p.NewTopicContentApp(contentType.Image(), "")
	img1.ImgUrl = "img1.jpg"
	img1.AttachmentsId = "227"
	audio := p.NewTopicContentApp(contentType.Audio(), "")
	audio.AudioUrl = "audio1.mp3"
	audio.AttachmentsId = "217"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.Text = "[/sell]"
	img2 := p.NewTopicContentApp(contentType.Image(), "")
	img2.ImgUrl = "img2.jpg"
	img2.AttachmentsId = "227"
	video := p.NewTopicContentApp(contentType.Video(), "")
	video.VideoUrl = "video1.mp4"
	video.AttachmentsId = "229"
	final := []TopicContentApp{paragraph1, openSell, paragraph2, img1, audio, closeSell, img2, video}

	// compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

// func TestSellWithParagraphEmojiImgDoorAndBook(t *testing.T) {
// 	input := `<html><head></head><body><p>你好你好[emoji]020[/emoji]<p><br/></p></p><p><p><br/></p></p>[sell]<img src="" data-id="59">[door]444327[/door][book]256[/book][/sell]<p><p><br/></p></p>[door]444420[/door][book]257[/book]</body></html>`
// 	got := startParsing(input)
// 	if len(got) != 9 {
// 		t.Errorf("got %v want %v", len(got), 9)
// 	}

// 	// expected
// 	p := new(AppParser)
// 	paragraph := p.NewTopicContentApp(contentType.Text(), "")
// 	paragraph.Text = "你好你好"
// 	emoji := p.NewTopicContentApp(contentType.Emoji(), "")
// 	emoji.EmoId = "[emoji]020[/emoji]"
// 	openSell := p.NewTopicContentApp(contentType.Text(), "")
// 	openSell.Text = "[sell]"
// 	img := p.NewTopicContentApp(contentType.Image(), "")
// 	img.ImgUrl = ""
// 	img.AttachmentsId = ""
// 	vid := p.NewTopicContentApp(contentType.Video(), "")
// 	vid.VideoUrl = """
// 	vid.AttachmentsId = ""
// 	audio := p.NewTopicContentApp(contentType.Audio(), "")
// 	audio.AudioUrl = ""
// 	audio.AttachmentsId = ""
// 	door := p.NewTopicContentApp(contentType.Topic(), "")
// 	door.TopicId = "[door]444327[/door]"
// 	book := p.NewTopicContentApp(contentType.Book(), "")
// 	book.BookId = "[book]256[/book]"
// 	final := []TopicContentApp{paragraph, emoji, img, vid, audio, door, book}
// }