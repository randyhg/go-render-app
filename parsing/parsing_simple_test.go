package parsing

import "testing"

func TestSingleParagraphParsing(t *testing.T) {
	//define input
	input := `<p>Hello World</p>`
	//set expected result
	p := new(AppParser)
	wants := p.NewTopicContentApp(contentType.Text(), "")
	wants.Text = `Hello World`
	//call startParsing and gets result
	got := startParsing(input)
	//compare
	if got[0] != wants {
		t.Errorf("got %v want %v", got[0].ToJson(), wants.ToJson())
	}
}

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

func startParsing(input string) (apps []TopicContentApp) {
	var p = new(AppParser)
	input = p.TrimContent(input)
	apps = p.Parse(input, apps)
	return
}
