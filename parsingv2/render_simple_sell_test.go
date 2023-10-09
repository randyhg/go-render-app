package parsingv2

import "testing"

func TestSellSingleImage(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><img src="./hello-world00.jpg" data-id="88">[sell]<img src="./hello-world.jpg" data-id="46">[/sell]<img src="./hello-world222.jpg" data-id="999"></body></html>`
	got := startParsing(input)
	if len(got) != 5 {
		t.Errorf("got %v want %v", len(got), 5)
	}
	//set expected result
	p := new(AppParser)
	img0 := p.NewTopicContentApp(contentType.Image(), "")
	img0.AttachmentsId = "88"
	img0.ImgUrl = "./hello-world00.jpg"
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.IsSell = "1"
	openSell.Text = "[sell]"
	openSell.IsCensored = "1"
	img := p.NewTopicContentApp(contentType.Image(), "")
	img.AttachmentsId = "46"
	img.ImgUrl = "./hello-world.jpg"
	img.IsSell = "1"
	img.IsCensored = "1"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.IsSell = "1"
	closeSell.Text = "[/sell]"
	closeSell.IsCensored = "1"
	img2 := p.NewTopicContentApp(contentType.Image(), "")
	img2.AttachmentsId = "999"
	img2.ImgUrl = "./hello-world222.jpg"
	final := []TopicContentApp{img0, openSell, img, closeSell, img2}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

func TestSellSingleAudio(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><audio src="./hello-world00.mp3" data-id="88">[sell]<audio src="./hello-world.mp3" data-id="46">[/sell]<audio src="./hello-world222.mp3" data-id="999"></body></html>`
	got := startParsing(input)
	if len(got) != 5 {
		t.Errorf("got %v want %v", len(got), 5)
	}
	//set expected result
	p := new(AppParser)
	audio0 := p.NewTopicContentApp(contentType.Audio(), "")
	audio0.AttachmentsId = "88"
	audio0.AudioUrl = "./hello-world00.mp3"
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.IsSell = "1"
	openSell.Text = "[sell]"
	openSell.IsCensored = "1"
	audio := p.NewTopicContentApp(contentType.Audio(), "")
	audio.AttachmentsId = "46"
	audio.AudioUrl = "./hello-world.mp3"
	audio.IsSell = "1"
	audio.IsCensored = "1"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.IsSell = "1"
	closeSell.Text = "[/sell]"
	closeSell.IsCensored = "1"
	audio2 := p.NewTopicContentApp(contentType.Audio(), "")
	audio2.AttachmentsId = "999"
	audio2.AudioUrl = "./hello-world222.mp3"
	final := []TopicContentApp{audio0, openSell, audio, closeSell, audio2}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

func TestSellSingleVideo(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><video src="./hello-world00.mp4" data-id="88">[sell]<video src="./hello-world.mp4" data-id="46">[/sell]<video src="./hello-world222.mp4" data-id="999"></body></html>`
	got := startParsing(input)
	if len(got) != 5 {
		t.Errorf("got %v want %v", len(got), 5)
	}
	//set expected result
	p := new(AppParser)
	video0 := p.NewTopicContentApp(contentType.Video(), "")
	video0.AttachmentsId = "88"
	video0.VideoUrl = "./hello-world00.mp4"
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.IsSell = "1"
	openSell.Text = "[sell]"
	openSell.IsCensored = "1"
	video := p.NewTopicContentApp(contentType.Video(), "")
	video.AttachmentsId = "46"
	video.VideoUrl = "./hello-world.mp4"
	video.IsSell = "1"
	video.IsCensored = "1"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.IsSell = "1"
	closeSell.Text = "[/sell]"
	closeSell.IsCensored = "1"
	video2 := p.NewTopicContentApp(contentType.Video(), "")
	video2.AttachmentsId = "999"
	video2.VideoUrl = "./hello-world222.mp4"
	final := []TopicContentApp{video0, openSell, video, closeSell, video2}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

func TestSellSingleMention(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><span data-uid="168128882007" class="at-box">@hjtestt7</span>[sell]<span data-uid="168128882007" class="at-box">@hjtestt7</span>[/sell]<span data-uid="168128882007" class="at-box">@hjtestt7</span></body></html>`
	got := startParsing(input)
	if len(got) != 5 {
		t.Errorf("got %v want %v", len(got), 5)
	}
	//set expected result
	p := new(AppParser)
	mention0 := p.NewTopicContentApp(contentType.Mention(), "")
	mention0.Uid = "168128882007"
	mention0.NickName = "@hjtestt7"
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.IsSell = "1"
	openSell.Text = "[sell]"
	openSell.IsCensored = "1"
	mention := p.NewTopicContentApp(contentType.Mention(), "")
	mention.Uid = "168128882007"
	mention.NickName = "@hjtestt7"
	mention.IsSell = "1"
	mention.IsCensored = "1"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.IsSell = "1"
	closeSell.Text = "[/sell]"
	closeSell.IsCensored = "1"
	mention2 := p.NewTopicContentApp(contentType.Mention(), "")
	mention2.Uid = "168128882007"
	mention2.NickName = "@hjtestt7"
	final := []TopicContentApp{mention0, openSell, mention, closeSell, mention2}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

func TestSellSingleBook(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p>[sell]<p><p><br/></p></p>[book]257[/book]<p><p><br/></p></p>[/sell]</body></html>`
	got := startParsing(input)
	if len(got) != 3 {
		t.Errorf("got %v want %v", len(got), 3)
	}
	//set expected result
	p := new(AppParser)
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.IsSell = "1"
	openSell.Text = "[sell]"
	openSell.IsCensored = "1"
	book := p.NewTopicContentApp(contentType.Book(), "")
	book.BookId = "[book]257[/book]"
	book.IsSell = "1"
	book.IsCensored = "1"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.IsSell = "1"
	closeSell.Text = "[/sell]"
	closeSell.IsCensored = "1"
	final := []TopicContentApp{openSell, book, closeSell}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

func TestSellSingleEmoji(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p>[sell]<p><p><br/></p></p>[emoji]002[/emoji]<p><p><br/></p></p>[/sell]</body></html>`
	got := startParsing(input)
	if len(got) != 3 {
		t.Errorf("got %v want %v", len(got), 3)
	}
	//set expected result
	p := new(AppParser)
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.IsSell = "1"
	openSell.Text = "[sell]"
	openSell.IsCensored = "1"
	emoji := p.NewTopicContentApp(contentType.Emoji(), "")
	emoji.EmoId = "[emoji]002[/emoji]"
	emoji.IsSell = "1"
	emoji.IsCensored = "1"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.IsSell = "1"
	closeSell.Text = "[/sell]"
	closeSell.IsCensored = "1"
	final := []TopicContentApp{openSell, emoji, closeSell}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

func TestSellSingleTopic(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p>[sell]<p><p><br/></p></p>[door]444327[/door]<p><p><br/></p></p>[/sell]</body></html>`
	got := startParsing(input)
	if len(got) != 3 {
		t.Errorf("got %v want %v", len(got), 3)
	}
	//set expected result
	p := new(AppParser)
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.IsSell = "1"
	openSell.Text = "[sell]"
	openSell.IsCensored = "1"
	door := p.NewTopicContentApp(contentType.Topic(), "")
	door.TopicId = "[door]444327[/door]"
	door.IsSell = "1"
	door.IsCensored = "1"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.IsSell = "1"
	closeSell.Text = "[/sell]"
	closeSell.IsCensored = "1"
	final := []TopicContentApp{openSell, door, closeSell}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

func TestSellSingleParagraph(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p>[sell]<p>lorem ipsum dolor sit amet</p>[/sell]</body></html>`
	got := startParsing(input)
	if len(got) != 3 {
		t.Errorf("got %v want %v", len(got), 3)
	}
	//set expected result
	p := new(AppParser)
	openSell := p.NewTopicContentApp(contentType.Text(), "")
	openSell.IsSell = "1"
	openSell.Text = "[sell]"
	openSell.IsCensored = "1"
	paragraph := p.NewTopicContentApp(contentType.Text(), "")
	paragraph.Text = "lorem ipsum dolor sit amet"
	paragraph.IsSell = "1"
	paragraph.IsCensored = "1"
	closeSell := p.NewTopicContentApp(contentType.Text(), "")
	closeSell.IsSell = "1"
	closeSell.Text = "[/sell]"
	closeSell.IsCensored = "1"
	final := []TopicContentApp{openSell, paragraph, closeSell}
	//compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}
