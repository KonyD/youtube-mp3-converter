package main

import (
	"fmt"
	"os"
	"os/exec"
)

// convert yt video to mp3
func main() {
	videoUrl := "https://www.youtube.com/watch?v=8aC6B_5aFCE&list=RD8aC6B_5aFCE&start_radio=1"
	videoFile := "video2.mp4"
	audioFile := "video2.mp3"

	err := downloadVideo(videoUrl, videoFile)
	if err != nil {
		fmt.Printf("Error downloading video: %v\n", err)
		return
	}

	err = convertVideoToAudio(videoFile, audioFile)
	if err != nil {
		fmt.Printf("Error converting video to audio: %v\n", err)
		return
	}

	fmt.Printf("Audio file created: %s\n", audioFile)
}

func downloadVideo(videoUrl, videoFile string) error {
	cmd := exec.Command("./yt-dlp", "-f", "bestaudio/best", "-o", videoFile, "--downloader", "aria2c", videoUrl)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func convertVideoToAudio(videoFile, audioFile string) error {
	cmd := exec.Command("ffmpeg", "-i", videoFile, "-vn", "-ab", "192k", "-ar", "44100", "-y", audioFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
