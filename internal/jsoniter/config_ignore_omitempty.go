package jsoniter

import "io"

func (cfg *frozenConfig) MarshalToStringIgnoreOmitempty(v interface{}) (string, error) {
	stream := cfg.BorrowStream(nil)
	defer cfg.ReturnStream(stream)
	stream.ignoreOmitempty = true
	stream.WriteVal(v)
	if stream.Error != nil {
		return "", stream.Error
	}
	return string(stream.Buffer()), nil
}

func (cfg *frozenConfig) MarshalIgnoreOmitempty(v interface{}) ([]byte, error) {
	stream := cfg.BorrowStream(nil)
	defer cfg.ReturnStream(stream)
	stream.ignoreOmitempty = true
	stream.WriteVal(v)
	if stream.Error != nil {
		return nil, stream.Error
	}
	result := stream.Buffer()
	copied := make([]byte, len(result))
	copy(copied, result)
	return copied, nil
}

func (cfg *frozenConfig) MarshalIndentIgnoreOmitempty(v interface{}, prefix, indent string) ([]byte, error) {
	if prefix != "" {
		panic("prefix is not supported")
	}
	for _, r := range indent {
		if r != ' ' {
			panic("indent can only be space")
		}
	}
	newCfg := cfg.configBeforeFrozen
	newCfg.IndentionStep = len(indent)
	return newCfg.frozeWithCacheReuse(cfg.extraExtensions).MarshalIgnoreOmitempty(v)
}

func (cfg *frozenConfig) NewEncoderIgnoreOmitempty(writer io.Writer) *Encoder {
	stream := NewStream(cfg, writer, 512)
	stream.ignoreOmitempty = true
	return &Encoder{stream}
}

type FrozenConfig = frozenConfig
