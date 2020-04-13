package adpcm

import (
	"errors"
	"fmt"

	"github.com/qiniu/audio"
	"github.com/qiniu/audio/wav"
	"github.com/qiniu/x/bufiox"
)

var (
	errNotImpl = errors.New("not impl")
)

// -------------------------------------------------------------------------------------

type block struct {
	buf       []byte
	samples   []byte
	remaining int
}

// -------------------------------------------------------------------------------------

type decoded struct {
	src        *bufiox.Reader
	sampleRate int
	channelNum int
}

func (p *decoded) Read(b []byte) (n int, err error) {
	return 0, errNotImpl
}

func (p *decoded) Seek(offset int64, whence int) (newoff int64, err error) {
	return 0, errNotImpl
}

// Length returns the size of decoded stream in bytes.
func (p *decoded) Length() int64 {
	return 0
}

// SampleRate returns the sample rate like 44100.
func (p *decoded) SampleRate() int {
	return p.sampleRate
}

// Channels returns the number of channels. One channel is mono playback.
// Two channels are stereo playback. No other values are supported.
func (p *decoded) Channels() int {
	return p.channelNum
}

// BytesPerSample returns the number of bytes per sample per channel.
// The usual value is 2. Only values 1 and 2 are supported.
func (p *decoded) BytesPerSample() int {
	return 2
}

// -------------------------------------------------------------------------------------

func decode(src *bufiox.Reader, cfg *wav.Config) (dec audio.Decoded, err error) {
	if cfg.BitsPerSample != 4 {
		return nil, fmt.Errorf("adpcm wav: bits per sample must be 4 but was %d", cfg.BitsPerSample)
	}
	d := &decoded{
		src:        src,
		sampleRate: cfg.SampleRate,
		channelNum: cfg.Channels,
	}
	return d, nil
}

const (
	adpcmFormat = 0x11
)

func init() {
	wav.RegisterFormat(adpcmFormat, decode)
}

// -------------------------------------------------------------------------------------
