package parsingv2

import (
	"testing"
)

// Success
func TestParagraphWithImg(t *testing.T) {
	input := `<p>你好哦哦哦哦哦</p><p><img src="https://res.hjcaecf.top/hjstore/images/20230911/e1b4561c88df36d1e3caa23f5dc9e3fd.jpg.txt?Expires=1694488992&Signature=PzupJZ1YQOr4cAlXtU5wMlxvaS4G~fs1cWK4h2bKRDJeDIQlYLwaVKTcVF-GPDT~kw-yWfckKe6THQbfCV9CBdwORZzUYEgrEGn1qm8O5FkSr5iwwOGesnxvD~rS9BRzYMVgtMNR4cgKU64AJHczeLz5oovk6BT7IKRucalEv6NrJNwOX6ZiCOikw6VOQeCxQjE6Vi6hAl0f~yWNd4xS~3rwoAJrZjpq3CkXmXTYUUyZib8PcXT9c8wwDk6N0vX6UlmWqU6Np2vhqhcnKn9lYzju0~KfW2z88WncSTjCRfjImhrBWp-AYfLsovT8zad8aKZjWOSZeAke8zLbvxKp8g__&Key-Pair-Id=K3HEWB9S9B6R6N" data-id="227"></p>`
	got := startParsing(input)
	if len(got) != 2 {
		t.Errorf("got %v want %v", len(got), 2)
	}

	// expected result
	p := new(AppParser)
	wants1 := p.NewTopicContentApp(contentType.Text(), "")
	wants2 := p.NewTopicContentApp(contentType.Image(), "")
	wants1.Text = "你好哦哦哦哦哦"
	wants2.AttachmentsId = "227"
	wants2.ImgUrl = "https://res.hjcaecf.top/hjstore/images/20230911/e1b4561c88df36d1e3caa23f5dc9e3fd.jpg.txt?Expires=1694488992&Signature=PzupJZ1YQOr4cAlXtU5wMlxvaS4G~fs1cWK4h2bKRDJeDIQlYLwaVKTcVF-GPDT~kw-yWfckKe6THQbfCV9CBdwORZzUYEgrEGn1qm8O5FkSr5iwwOGesnxvD~rS9BRzYMVgtMNR4cgKU64AJHczeLz5oovk6BT7IKRucalEv6NrJNwOX6ZiCOikw6VOQeCxQjE6Vi6hAl0f~yWNd4xS~3rwoAJrZjpq3CkXmXTYUUyZib8PcXT9c8wwDk6N0vX6UlmWqU6Np2vhqhcnKn9lYzju0~KfW2z88WncSTjCRfjImhrBWp-AYfLsovT8zad8aKZjWOSZeAke8zLbvxKp8g__&Key-Pair-Id=K3HEWB9S9B6R6N"

	// compare
	if got[0] != wants1 || got[1] != wants2 {
		t.Errorf("got[0] %v, want1 %v and got[1] %v, want2 %v", got[0].ToJson(), wants1.ToJson(), got[1].ToJson(), wants2.ToJson())
	}
}

// Failed
// got 5 want 4
func TestParagraphWithEmojiDoorAndBook(t *testing.T) {
	input := `<html><head></head><body><p>早安大家[emoji]001[/emoji][emoji]002[/emoji]<p><br/></p></p>[door]444341[/door][book]256[/book]</body></html>`
	got := startParsing(input)
	if len(got) != 4 {
		t.Errorf("got %v want %v", len(got), 4)
	}

	// expected
	p := new(AppParser)
	paragraph := p.NewTopicContentApp(contentType.Text(), "")
	paragraph.Text = "早安大家"
	emoji := p.NewTopicContentApp(contentType.Emoji(), "")
	emoji.EmoId = "[emoji]001[/emoji][emoji]002[/emoji]"
	door := p.NewTopicContentApp(contentType.Topic(), "")
	door.TopicId = "[door]444341[/door]"
	book := p.NewTopicContentApp(contentType.Book(), "")
	book.BookId = "[book]256[/book]"
	final := []TopicContentApp{paragraph, emoji, door, book}

	// compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

// Failed
// got 6 want 5
func TestParagraphWithEmojiImgDoorAndBook(t *testing.T) {
	input := `<html><head></head><body><p>你好你好[emoji]005[/emoji][emoji]005[/emoji]<p><br/></p></p><img src="https://res.hjcaecf.top/hjstore/images/20230731/a59ee5a2136ff763d6dc907b5074afb2.png.txt?Expires=1694579387&Signature=qMU7C3PPw9TK5rPiZ4sw70RR74XQJHyCp8hEX6bqz4mVzd1njpD-Wb1H1cYEKutPiWYDoG4baE2eyMiLKmbVLzz~qGdtAT5NsQr9hwT6jNwz~72WznF8DDf82LoQTx5caKvCc8Kw3Cy1PKaDUGgQAWEtmQ1RuQKFuZA841bgOXcFMzKdQBbvhQz7qNNVrNSMn9QcmadRYZ~QLNYQpYZ3NQJKLyS253B3k2nYLp2cGDY8SFhSIwM2SxDTLe3paLe6c1Yp-3y9D~fJA3FgPHWDYtSWL67iHnQEBexHhqKAJf4mbd1wUPoJsaEP64YdRabVzCgT19r~2KZZ4EcfD3Opiw__&Key-Pair-Id=K3HEWB9S9B6R6N" data-id="59">[door]444324[/door][book]256[/book]</body></html>`
	got := startParsing(input)
	if len(got) != 5 {
		t.Errorf("got %v want %v", len(got), 5)
	}

	// expected
	p := new(AppParser)
	paragraph := p.NewTopicContentApp(contentType.Text(), "")
	paragraph.Text = "你好你好"
	emoji := p.NewTopicContentApp(contentType.Emoji(), "")
	emoji.EmoId = "[emoji]005[/emoji][emoji]005[/emoji]"
	img := p.NewTopicContentApp(contentType.Image(), "")
	img.ImgUrl = "https://res.hjcaecf.top/hjstore/images/20230731/a59ee5a2136ff763d6dc907b5074afb2.png.txt?Expires=1694579387&Signature=qMU7C3PPw9TK5rPiZ4sw70RR74XQJHyCp8hEX6bqz4mVzd1njpD-Wb1H1cYEKutPiWYDoG4baE2eyMiLKmbVLzz~qGdtAT5NsQr9hwT6jNwz~72WznF8DDf82LoQTx5caKvCc8Kw3Cy1PKaDUGgQAWEtmQ1RuQKFuZA841bgOXcFMzKdQBbvhQz7qNNVrNSMn9QcmadRYZ~QLNYQpYZ3NQJKLyS253B3k2nYLp2cGDY8SFhSIwM2SxDTLe3paLe6c1Yp-3y9D~fJA3FgPHWDYtSWL67iHnQEBexHhqKAJf4mbd1wUPoJsaEP64YdRabVzCgT19r~2KZZ4EcfD3Opiw__&Key-Pair-Id=K3HEWB9S9B6R6N"
	img.AttachmentsId = "59"
	door := p.NewTopicContentApp(contentType.Topic(), "")
	door.TopicId = "[door]444324[/door]"
	book := p.NewTopicContentApp(contentType.Book(), "")
	book.BookId = "[book]256[/book]"
	final := []TopicContentApp{paragraph, emoji, img, door, book}

	// compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

// Success
func TestParagraphWithEmojiImgVidDoorAndBook(t *testing.T) {
	input := `<html><head></head><body><p>好烂[emoji]004[/emoji]<p><br/></p></p><img src="https://res.hjcaecf.top/hjstore/images/20230731/a59ee5a2136ff763d6dc907b5074afb2.png.txt?Expires=1694579823&Signature=GyRItUpK-~6~Ym241y5ZKrRq2-6YCn7VN5NausCguSlGkcans1eLJz7Ungputw3uDnp7lO4GA6msKINwszh7q7Zs9Ayl-xpFXINcVUudc9SolRFCRzXxs79S8RMCb5vWwhfbxJKO0bFnuxtw~~xf5RSsXMZP-khozf-ezaf5OoubVLdI8KX-~nU6yJdOOBDVy4ehCfxilcgAB8RAQobUZiaDxiaCrEcNBGL3ZTi8IehIpB0n2~G9MAUKYu1VhMU90YnqGeGXnn3KhIc153gdkuGPpmQny79aPb9kpMrtLitJ2i3aB4O~XyA4-DoJQ9hpXc8eUA84Z-Ap3iUfJTJOgw__&Key-Pair-Id=K3HEWB9S9B6R6N" data-id="59"><video src="https://hjcaecf.top/api/m3u8/video/20230805/38ffe945ceb1940f4b62ecc81401302d/96dAYc320s_i99bb47469a4342cbb64cb91d27671164.m3u8?Expires=1694579823&Signature=DTNwX3inTC7L3KRBvpAn4TYqVGIXYOOVgjSDCZPmNnV5yfyuKwOz5ZwjCrzjDH0P8o7TiYB5l~mDxoTRde4GkOXSL2ijxlcP14ZE0Ni~J35PAQvSAUMaPpXAJMq~yGbBs03AjTGqik~JtH0Ewhq5BeNV0MbTpSppMsoRHf5IsThrob2hYXx-DHd4mcvJH9esHeRNL29rrIa5HfgmRYPSmD5l1cBXzW2lYycSizIUeiexY~iXiR6XafqZ6alKjwVXCHs09INnvQOWVVmbc~sHjcwt~A5PC3rpj8rwMtoWyoPtNZQqkzWZLRfsh9Wxi8hX3xljArAy0cnK47CS8JhwZg__&Key-Pair-Id=K3HEWB9S9B6R6N" data-id="96"></video>[door]444327[/door][book]256[/book]<p><p><br/></p></p></body></html>`
	got := startParsing(input)
	if len(got) != 6 {
		t.Errorf("got %v want %v", len(got), 6)
	}

	// expected
	p := new(AppParser)
	paragraph := p.NewTopicContentApp(contentType.Text(), "")
	paragraph.Text = "好烂"
	emoji := p.NewTopicContentApp(contentType.Emoji(), "")
	emoji.EmoId = "[emoji]004[/emoji]"
	img := p.NewTopicContentApp(contentType.Image(), "")
	img.ImgUrl = "https://res.hjcaecf.top/hjstore/images/20230731/a59ee5a2136ff763d6dc907b5074afb2.png.txt?Expires=1694579823&Signature=GyRItUpK-~6~Ym241y5ZKrRq2-6YCn7VN5NausCguSlGkcans1eLJz7Ungputw3uDnp7lO4GA6msKINwszh7q7Zs9Ayl-xpFXINcVUudc9SolRFCRzXxs79S8RMCb5vWwhfbxJKO0bFnuxtw~~xf5RSsXMZP-khozf-ezaf5OoubVLdI8KX-~nU6yJdOOBDVy4ehCfxilcgAB8RAQobUZiaDxiaCrEcNBGL3ZTi8IehIpB0n2~G9MAUKYu1VhMU90YnqGeGXnn3KhIc153gdkuGPpmQny79aPb9kpMrtLitJ2i3aB4O~XyA4-DoJQ9hpXc8eUA84Z-Ap3iUfJTJOgw__&Key-Pair-Id=K3HEWB9S9B6R6N"
	img.AttachmentsId = "59"
	vid := p.NewTopicContentApp(contentType.Video(), "")
	vid.VideoUrl = "https://hjcaecf.top/api/m3u8/video/20230805/38ffe945ceb1940f4b62ecc81401302d/96dAYc320s_i99bb47469a4342cbb64cb91d27671164.m3u8?Expires=1694579823&Signature=DTNwX3inTC7L3KRBvpAn4TYqVGIXYOOVgjSDCZPmNnV5yfyuKwOz5ZwjCrzjDH0P8o7TiYB5l~mDxoTRde4GkOXSL2ijxlcP14ZE0Ni~J35PAQvSAUMaPpXAJMq~yGbBs03AjTGqik~JtH0Ewhq5BeNV0MbTpSppMsoRHf5IsThrob2hYXx-DHd4mcvJH9esHeRNL29rrIa5HfgmRYPSmD5l1cBXzW2lYycSizIUeiexY~iXiR6XafqZ6alKjwVXCHs09INnvQOWVVmbc~sHjcwt~A5PC3rpj8rwMtoWyoPtNZQqkzWZLRfsh9Wxi8hX3xljArAy0cnK47CS8JhwZg__&Key-Pair-Id=K3HEWB9S9B6R6N"
	vid.AttachmentsId = "96"
	door := p.NewTopicContentApp(contentType.Topic(), "")
	door.TopicId = "[door]444327[/door]"
	book := p.NewTopicContentApp(contentType.Book(), "")
	book.BookId = "[book]256[/book]"
	final := []TopicContentApp{paragraph, emoji, img, vid, door, book}

	// compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}

// Failed
// got 8 want 7
func TestParagraphWithEmojiImgVidAudioDoorAndBook(t *testing.T) {
	input := `<html><head></head><body><p>你好[emoji]002[/emoji][emoji]003[/emoji]<p><br/></p></p><img src="https://res.hjcaecf.top/hjstore/images/20230731/a59ee5a2136ff763d6dc907b5074afb2.png.txt?Expires=1694579920&Signature=fE8OMhK6vg7vrmEjZ~BGuB5Kligeo6TBnLbUjx3RNUNvnnzwZ~4ku2-OseiC6MTSAmjCMmTJvTgEjIOzCYkhgdgy8FsTYN7oomUyrGFaIPE6lx4Fw0AZBZ~WAPyvqTLbF-20zH8ELO4iJWNmJ06523niVNLNCAjT5T7T64lHxgnYlkq~idYC21AhjtYUGI7gMpAwWLzkJ2ElGXkvP8y7EE13Yok~MJy4sQYDyPmEFYsFLw~-KaJL1jMLpreDu28O9R5YbsVoqG2GSqVT90E4bcoAFrWVDLhJPSPqOJg18TgpFmZhNAODydfmysZoeq~8zVR-DY9SiNqmH1UMJ5C9KQ__&Key-Pair-Id=K3HEWB9S9B6R6N" data-id="59"><video src="https://hjcaecf.top/api/m3u8/video/20230805/d24565e283d772d00336028855673023/95nFNeZITs_i99bb47469a4342cbb64cb91d27671164.m3u8?Expires=1694579920&Signature=eUJBCyoRKkeGRcEsbFx5jm8warPb3nfnueNXClgZkZTUPZoL2JoVWCMEnz5ivoi~fbWmBKUIZqvcTTjCRm01pvXeN8HphUeYHumP6ehzVwLwdcMK36kI9QBCo6TFgnk8AeHdfnxPURS~QbzeDpS4zpqBWZ6CopeLAxEQtcu99kWT8yLGFyeuBADF399YI0rpsvFE9jR79UDV36KBNayeOdd1X1YSuKZ47u76EOTYC8KPAIK1x31mrt3m2Z54wnEX9D~UPwbc5kaBhNTDtEhu12BAUf7r2s6g9al-tBhLrYowCsxD7uPVKX3NmQpbcFsTVQTED5xeohGFjn2XI-hjqg__&Key-Pair-Id=K3HEWB9S9B6R6N" data-id="95"></video><audio src="https://res.hjcaecf.top/hjstore/audio/20230720/cf53d6f6d358a9f3e8c827bf2086b340.mp3" controls="true" controlslist="nodownload" id="showaudio" data-id="2"></audio>[door]444327[/door][book]256[/book]<p><p><br/></p><p><br/></p></p></body></html>`
	got := startParsing(input)
	if len(got) != 7 {
		t.Errorf("got %v want %v", len(got), 7)
	}

	// expected
	p := new(AppParser)
	paragraph := p.NewTopicContentApp(contentType.Text(), "")
	paragraph.Text = "你好"
	emoji := p.NewTopicContentApp(contentType.Emoji(), "")
	emoji.EmoId = "[emoji]002[/emoji][emoji]003[/emoji]"
	img := p.NewTopicContentApp(contentType.Image(), "")
	img.ImgUrl = "https://res.hjcaecf.top/hjstore/images/20230731/a59ee5a2136ff763d6dc907b5074afb2.png.txt?Expires=1694579920&Signature=fE8OMhK6vg7vrmEjZ~BGuB5Kligeo6TBnLbUjx3RNUNvnnzwZ~4ku2-OseiC6MTSAmjCMmTJvTgEjIOzCYkhgdgy8FsTYN7oomUyrGFaIPE6lx4Fw0AZBZ~WAPyvqTLbF-20zH8ELO4iJWNmJ06523niVNLNCAjT5T7T64lHxgnYlkq~idYC21AhjtYUGI7gMpAwWLzkJ2ElGXkvP8y7EE13Yok~MJy4sQYDyPmEFYsFLw~-KaJL1jMLpreDu28O9R5YbsVoqG2GSqVT90E4bcoAFrWVDLhJPSPqOJg18TgpFmZhNAODydfmysZoeq~8zVR-DY9SiNqmH1UMJ5C9KQ__&Key-Pair-Id=K3HEWB9S9B6R6N"
	img.AttachmentsId = "59"
	vid := p.NewTopicContentApp(contentType.Video(), "")
	vid.VideoUrl = "https://hjcaecf.top/api/m3u8/video/20230805/d24565e283d772d00336028855673023/95nFNeZITs_i99bb47469a4342cbb64cb91d27671164.m3u8?Expires=1694579920&Signature=eUJBCyoRKkeGRcEsbFx5jm8warPb3nfnueNXClgZkZTUPZoL2JoVWCMEnz5ivoi~fbWmBKUIZqvcTTjCRm01pvXeN8HphUeYHumP6ehzVwLwdcMK36kI9QBCo6TFgnk8AeHdfnxPURS~QbzeDpS4zpqBWZ6CopeLAxEQtcu99kWT8yLGFyeuBADF399YI0rpsvFE9jR79UDV36KBNayeOdd1X1YSuKZ47u76EOTYC8KPAIK1x31mrt3m2Z54wnEX9D~UPwbc5kaBhNTDtEhu12BAUf7r2s6g9al-tBhLrYowCsxD7uPVKX3NmQpbcFsTVQTED5xeohGFjn2XI-hjqg__&Key-Pair-Id=K3HEWB9S9B6R6N"
	vid.AttachmentsId = "95"
	audio := p.NewTopicContentApp(contentType.Audio(), "")
	audio.AudioUrl = "https://res.hjcaecf.top/hjstore/audio/20230720/cf53d6f6d358a9f3e8c827bf2086b340.mp3"
	audio.AttachmentsId = "2"
	door := p.NewTopicContentApp(contentType.Topic(), "")
	door.TopicId = "[door]444327[/door]"
	book := p.NewTopicContentApp(contentType.Book(), "")
	book.BookId = "[book]256[/book]"
	final := []TopicContentApp{paragraph, emoji, img, vid, audio, door, book}

	// compare
	for i := 0; i < len(got); i++ {
		if got[i] != final[i] {
			t.Errorf("got %v want %v", got[i].ToJson(), final[i].ToJson())
		}
	}
}
