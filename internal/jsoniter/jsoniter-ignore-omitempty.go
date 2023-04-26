package jsoniter

/******************************************************************************\
# jsoniter扩展

## 扩展ignoreOmitempty
- 修改config.go:[68]
```
type frozenConfig struct {
	...
	ignoreOmitempty bool
}
```
添加ignoreOmitempty属性.

- 修改reflect_struct_encoder.go:[144]
```
func (encoder *structEncoder) Encode(ptr unsafe.Pointer, stream *Stream) {
	...
	for _, field := range encoder.fields {
		if !stream.cfg.ignoreOmitempty && field.encoder.omitempty && field.encoder.IsEmpty(ptr) {
			continue
		}
	...

}
```
添加"!stream.cfg.ignoreOmitempty"作为前决条件.

- 新增jsoniter-ignore-omitempty.go
```
// FrozeIgnoreOmitempty forge API from config
func (cfg Config) FrozeIgnoreOmitempty() (API, API) {
	api := cfg.Froze().(*frozenConfig)
	return api, cloneIgnoreOmitempty(api)
}

func cloneIgnoreOmitempty(cfg *frozenConfig) *frozenConfig {
	ret := *cfg
	ret.ignoreOmitempty = true
	return &ret
}
```
生成配对的frozenConfig实现, 二者共用相同的cache与pool.


## 移除concurrent.Map
- config.go:[75]/[76]
```
type frozenConfig struct {
	...
	decoderCache                  *sync.Map
	encoderCache                  *sync.Map
	...
	ignoreOmitempty               bool
}

func (cfg *frozenConfig) initCache() {
	cfg.decoderCache = &sync.Map{}
	cfg.encoderCache = &sync.Map{}
}

var cfgCache = &sync.Map{}

```
除了ignoreOmitempty外, 将decodeCache/encoderCache改成*sync.Map. 移除concurrent依赖!
同时修改initCache()的初始逻辑!
\******************************************************************************/

// FrozeIgnoreOmitempty forge API from config
func (cfg Config) FrozeIgnoreOmitempty() (API, API) {
	api := cfg.Froze().(*frozenConfig)
	return api, cloneIgnoreOmitempty(api)
}

func cloneIgnoreOmitempty(cfg *frozenConfig) *frozenConfig {
	ret := *cfg
	ret.ignoreOmitempty = true
	return &ret
}
