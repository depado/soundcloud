package soundcloud

type StreamType struct {
	Protocol string
	Preset   string
}

var (
	HLSMP3         = StreamType{Protocol: "hls", Preset: "mp3_0_0"}
	ProgressiveMP3 = StreamType{Protocol: "progressive", Preset: "mp3_0_0"}
	Opus           = StreamType{Protocol: "hls", Preset: "opus_0_0"}
)
