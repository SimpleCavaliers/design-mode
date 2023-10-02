package main

import "fmt"

type MediaPlayer interface {
	Play(audioType, fileName string)
}

type AdvancedMediaPlayer interface {
	PlayVlc(fileName string)
	PlayMp4(fileName string)
}

// Advanced的实体类
type VlcPlayer struct {}
type Mp4Player struct {}

func (*VlcPlayer) PlayVlc(fileName string) {
	fmt.Println("Playing vlc file. Name: " + fileName)
}

func (*VlcPlayer) PlayMp4(fileName string)  {}

func (*Mp4Player) PlayVlc(fileName string)  {}

func (*Mp4Player) PlayMp4(fileName string)  {
	fmt.Println("Playing mp4 file. Name: " + fileName)
}

// 适配器
type MediaAdapter struct {
	advancedMusicPlayer AdvancedMediaPlayer
	audioType 			string
}

func (ma *MediaAdapter) Create() {
	if ma.audioType == "vlc"{
		ma.advancedMusicPlayer = &VlcPlayer{}
	}else if ma.audioType == "mp4"{
		ma.advancedMusicPlayer = &Mp4Player{}
	}
}

func (ma *MediaAdapter) Play(audioType, fileName string) {
	if audioType == "vlc"{
		ma.advancedMusicPlayer.PlayVlc(fileName)
	}else if audioType == "mp4" {
		ma.advancedMusicPlayer.PlayMp4(fileName)
	}
}

// MediaPlayer的实体类
type AudioPlayer struct {
	mediaAdapter *MediaAdapter
}

func (ap *AudioPlayer) Play(audioType, fileName string) {
	if audioType == "mp3"{
		fmt.Println("Playing mp3 file. Name: " + fileName)
	}else if audioType == "mp4" || audioType == "vlc"{
		ap.mediaAdapter = &MediaAdapter{audioType: audioType}
		ap.mediaAdapter.Create()
		ap.mediaAdapter.Play(audioType, fileName)
	}else {
		fmt.Println("Invalid media. " + audioType + " format not supported")
	}
}

func main() {
	audioPlayer := &AudioPlayer{}

	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}
