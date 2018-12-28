package structural

import "testing"

func TestAdapter(t *testing.T) {
	mediaAdapter := &MediaAdapter{}
	vlcMedia := "hhh.vlc"
	mp4Media := "aaaa.mp4"
	mkvMedia := "dfg.mkv"
	mediaAdapter.play(vlcMedia)
	mediaAdapter.play(mp4Media)
	mediaAdapter.play(mkvMedia)
}
func TestAudioPlayer(t *testing.T) {
	ap := AudioPlayer{}
	mp3Media := "1234.mp3"
	vlcMedia := "hhh.vlc"
	mp4Media := "aaaa.mp4"
	mkvMedia := "dfg.mkv"
	ap.play(mp3Media)
	ap.play(mp4Media)
	ap.play(vlcMedia)
	ap.play(mkvMedia)
}
