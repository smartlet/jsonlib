package jsonlib

import "io"

// Marshal json序列化
func Marshal(value any) ([]byte, error) {

}

// MarshalIgnoreOmitempty json序列化, 忽略omitempty属性.
func MarshalIgnoreOmitempty(value any) ([]byte, error) {

}

// Unmarshal json反序列化
func Unmarshal(data []byte, v any) error {

}

// Encoder 流式序列化器
type Encoder struct {
}

func NewEncoder(w io.Writer) *Encoder {

}

// Encode json序列化
func (e *Encoder) Encode(value any) error {

}

// EncodeIgnoreOmitempty json序列化, 忽略omitempty属性.
func (e *Encoder) EncodeIgnoreOmitempty(value any) error {

}

// Decoder 流式反序列化器
type Decoder struct {
}

func NewDecoder(r io.Reader) *Decoder {

}

// Decode json反序列化.
func (d *Decoder) Decode(value any) error {

}

// More json反序列化是否有下一个.
func (d *Decoder) More() bool {

}

func MarshalObject(value any) (map[string]any, error) {

}

func UnmarshalObject(object map[string]any, value any) error {

}

func MarshalObjectArray(value []any) ([]map[string]any, error) {

}
