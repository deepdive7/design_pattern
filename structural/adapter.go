package structural

import (
	"fmt"
	"strings"
)

//MediaPlayer ...
type MediaPlayer interface {
	play(string)
}

//AdvanceMediaPlayer ...
type AdvanceMediaPlayer interface {
	playVlc(string)
	playMp4(string)
}

// VlcPlayer ...
type VlcPlayer struct{}

func (vp VlcPlayer) playVlc(fileName string) {
	fmt.Println("File type is vlc, playing ", fileName)
}

func (vp VlcPlayer) playMp4(fileName string) {}

// Mp4Player ...
type Mp4Player struct{}

func (mp Mp4Player) playMp4(fileName string) {
	fmt.Println("File type is mp4, playing ", fileName)
}

func (mp Mp4Player) playVlc(fileName string) {}

//MediaAdapter ...
type MediaAdapter struct {
	amp AdvanceMediaPlayer
}

func (ma *MediaAdapter) play(fileName string) {

	switch strings.ToLower(fileName[strings.LastIndex(fileName, ".")+1:]) {
	case "vlc":
		ma.amp = &VlcPlayer{}
		ma.amp.playVlc(fileName)
	case "mp4":
		ma.amp = &Mp4Player{}
		ma.amp.playMp4(fileName)
	default:
		fmt.Println("Unsupport media type, only vlc and mp4")
	}
}

//AudioPlayer ...
type AudioPlayer struct{}

func (ap AudioPlayer) play(fileName string) {
	mediaType := strings.ToLower(fileName[strings.LastIndex(fileName, ".")+1:])
	switch mediaType {
	case "mp3":
		fmt.Println("File type is mp3, playing ", fileName)
	case "mp4", "vlc":
		ma := MediaAdapter{}
		ma.play(fileName)
	default:
		fmt.Println("Unsupport media type, only mp3, vlc and mp4")
	}
}
