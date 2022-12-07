package soundcloud

type StreamType struct {
	Protocol     string
	PresetPrefix string
}

var (
	HLSMP3         = StreamType{Protocol: "hls", PresetPrefix: "mp3"}
	ProgressiveMP3 = StreamType{Protocol: "progressive", PresetPrefix: "mp3"}
	Opus           = StreamType{Protocol: "hls", PresetPrefix: "opus"}
)
