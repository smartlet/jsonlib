package jsonlib

import (
	"github.com/smartlet/jsonlib/internal/jsoniter"
	"io"
)

type (
	Config  = jsoniter.Config
	Decoder = jsoniter.Decoder
	Encoder = jsoniter.Encoder
)

var iterator, iteratorIgnoreOmitempty = Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       false, // will lose precession
	ObjectFieldMustBeSimpleString: true,  // do not unescape object field
}.FrozeIgnoreOmitempty()

func Marshal(v any) ([]byte, error) {
	return iterator.Marshal(v)
}

func MarshalToString(v any) (string, error) {
	return iterator.MarshalToString(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return iterator.Unmarshal(data, v)
}

func NewEncoder(writer io.Writer) *Encoder {
	return iterator.NewEncoder(writer)
}

func NewDecoder(reader io.Reader) *Decoder {
	return iterator.NewDecoder(reader)
}

func MarshalIgnoreOmitempty(v any) ([]byte, error) {
	return iteratorIgnoreOmitempty.Marshal(v)
}

func MarshalToStringIgnoreOmitempty(v any) (string, error) {
	return iteratorIgnoreOmitempty.MarshalToString(v)
}

func NewEncoderIgnoreOmitempty(writer io.Writer) *Encoder {
	return iteratorIgnoreOmitempty.NewEncoder(writer)
}

func ToJson(v any) string {
	s, _ := MarshalToString(v)
	return s
}
