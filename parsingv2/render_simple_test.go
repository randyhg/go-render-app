package parsingv2

import "testing"

func TestSingleImage(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><img src="./hello-world.jpg" data-id="46"></body></html>`
	got := startParsing(input)
	if len(got) != 1 {
		t.Errorf("got %v want %v", len(got), 1)
	}
	//set expected result
	p := new(AppParser)
	wants := p.NewTopicContentApp(contentType.Image(), "")
	wants.AttachmentsId = "46"
	wants.ImgUrl = "./hello-world.jpg"
	//compare
	if got[0] != wants {
		t.Errorf("got %v want %v", got[0].ToJson(), wants.ToJson())
	}
}

func TestSingleAudio(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><audio src="./hello-world.mp3" data-id="46"></audio></body></html>`
	got := startParsing(input)
	if len(got) != 1 {
		t.Errorf("got %v want %v", len(got), 1)
	}
	//set expected result
	p := new(AppParser)
	wants := p.NewTopicContentApp(contentType.Audio(), "")
	wants.AttachmentsId = "46"
	wants.AudioUrl = "./hello-world.mp3"
	//compare
	if got[0] != wants {
		t.Errorf("got %v want %v", got[0].ToJson(), wants.ToJson())
	}
}

func TestSingleVideo(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><video src="./hello-world.mp4" data-id="46"></video></body></html>`
	got := startParsing(input)
	if len(got) != 1 {
		t.Errorf("got %v want %v", len(got), 1)
	}
	//set expected result
	p := new(AppParser)
	wants := p.NewTopicContentApp(contentType.Video(), "")
	wants.AttachmentsId = "46"
	wants.VideoUrl = "./hello-world.mp4"
	//compare
	if got[0] != wants {
		t.Errorf("got %v want %v", got[0].ToJson(), wants.ToJson())
	}
}

func TestSingleMention(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><span data-uid="168128882007" class="at-box">@hjtestt7</span></body></html>`
	got := startParsing(input)
	if len(got) != 1 {
		t.Errorf("got %v want %v", len(got), 1)
	}
	//set expected result
	p := new(AppParser)
	wants := p.NewTopicContentApp(contentType.Mention(), "")
	wants.Uid = "168128882007"
	wants.NickName = "@hjtestt7"
	//compare
	if got[0] != wants {
		t.Errorf("got %v want %v", got[0].ToJson(), wants.ToJson())
	}
}

func TestSingleDoor(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p>[door]444341[/door]</body></html>`
	got := startParsing(input)
	if len(got) != 1 {
		t.Errorf("got %v want %v", len(got), 1)
	}
	//set expected result
	p := new(AppParser)
	wants := p.NewTopicContentApp(contentType.Topic(), "")
	wants.TopicId = "[door]444341[/door]"
	//compare
	if got[0] != wants {
		t.Errorf("got %v want %v", got[0].ToJson(), wants.ToJson())
	}
}

func TestSingleEmoji(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p>[emoji]005[/emoji]</body></html>`
	got := startParsing(input)
	if len(got) != 1 {
		t.Errorf("got %v want %v", len(got), 1)
	}
	//set expected result
	p := new(AppParser)
	wants := p.NewTopicContentApp(contentType.Emoji(), "")
	wants.EmoId = "[emoji]005[/emoji]"
	//compare
	if got[0] != wants {
		t.Errorf("got %v want %v", got[0].ToJson(), wants.ToJson())
	}
}

func TestSingleBook(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p>[book]256[/book]</body></html>`
	got := startParsing(input)
	if len(got) != 1 {
		t.Errorf("got %v want %v", len(got), 1)
	}
	//set expected result
	p := new(AppParser)
	wants := p.NewTopicContentApp(contentType.Book(), "")
	wants.BookId = "[book]256[/book]"
	//compare
	if got[0] != wants {
		t.Errorf("got %v want %v", got[0].ToJson(), wants.ToJson())
	}
}

func TestSingleParagraph(t *testing.T) {
	input := `<html><head></head><body><p><p><br/></p></p><p>lorem ipsum dolor sit amet</p></body></html>`
	got := startParsing(input)
	if len(got) != 1 {
		t.Errorf("got %v want %v", len(got), 1)
	}
	//set expected result
	p := new(AppParser)
	wants := p.NewTopicContentApp(contentType.Text(), "")
	wants.Text = "lorem ipsum dolor sit amet"
	//compare
	if got[0] != wants {
		t.Errorf("got %v want %v", got[0].ToJson(), wants.ToJson())
	}
}
func startParsing(input string) (apps []TopicContentApp) {
	var p = new(AppParser)
	apps = p.Render(input)
	return
}
