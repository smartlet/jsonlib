package jsonlib

import (
	"github.com/smartlet/jsonlib/internal/jsoniter"
	"io"
)

type (
	Decoder = jsoniter.Decoder
	Encoder = jsoniter.Encoder
)

var config = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       false, // will lose precession
	ObjectFieldMustBeSimpleString: true,  // do not unescape object field
}.Froze().(*jsoniter.FrozenConfig)

func MarshalToString(v any) (string, error) {
	return config.MarshalToString(v)
}
func MarshalToStringIgnoreOmitempty(v any) (string, error) {
	return config.MarshalToStringIgnoreOmitempty(v)
}

func Marshal(v any) ([]byte, error) {
	return config.Marshal(v)
}

func MarshalIgnoreOmitempty(v any) ([]byte, error) {
	return config.MarshalIgnoreOmitempty(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return config.MarshalIndent(v, prefix, indent)
}

func MarshalIndentIgnoreOmitempty(v interface{}, prefix, indent string) ([]byte, error) {
	return config.MarshalIndentIgnoreOmitempty(v, prefix, indent)
}

func NewEncoder(writer io.Writer) *Encoder {
	return config.NewEncoder(writer)
}

func NewEncoderIgnoreOmitempty(writer io.Writer) *Encoder {
	return config.NewEncoderIgnoreOmitempty(writer)
}

func Unmarshal(data []byte, v interface{}) error {
	return config.Unmarshal(data, v)
}

func NewDecoder(reader io.Reader) *Decoder {
	return config.NewDecoder(reader)
}

func ToJson(v any) string {
	s, _ := MarshalToString(v)
	return s
}
