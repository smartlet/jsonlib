package jsonlib

import (
	"io"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func UnsafeString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func UnsafeBytes(s string) (b []byte) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return
}

func NewStringBuffer(cap int) *StringBuffer {
	return &StringBuffer{
		cur: 0,
		len: 0,
		cap: cap,
		buf: make([]byte, cap),
	}
}

type StringBuffer struct {
	cur int // 读下标
	len int // 写下标
	cap int
	buf []byte
}

func (b *StringBuffer) Close() error {
	return nil
}

func (b *StringBuffer) Remain() int {
	return b.len - b.cur
}

func (b *StringBuffer) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		b.cur = int(offset)
	case io.SeekCurrent:
		b.cur += int(offset)
	case io.SeekEnd:
		b.cur = b.len - int(offset)
	}
	if b.cur < 0 {
		b.cur = 0
	} else if b.cur > b.len {
		b.cur = b.len
	}
	return int64(b.cur), nil
}

// 返回安全的string
func (b *StringBuffer) String() string {
	return string(b.buf[b.cur:b.len])
}

func (b *StringBuffer) UnsafeString() string {
	buf := b.buf[b.cur:b.len]
	return *(*string)(unsafe.Pointer(&buf))
}

func (b *StringBuffer) UnsafeBytes() []byte {
	return b.buf[b.cur:b.len]
}

func (b *StringBuffer) Cur() int { return b.cur }

func (b *StringBuffer) Len() int { return b.len }

func (b *StringBuffer) Cap() int { return b.cap }

func (b *StringBuffer) Reset() *StringBuffer {
	b.cur = 0
	b.len = 0
	return b
}

func (b *StringBuffer) Grow(n int) {
	b.cap += n
	buf := make([]byte, b.cap)
	copy(buf, b.buf)
	b.buf = buf
}

func (b *StringBuffer) Unread(n int) {
	if b.len > n {
		b.len -= n
	} else {
		b.len = 0
	}
}

func (b *StringBuffer) Read(p []byte) (int, error) {
	if b.cur < b.len {
		n := copy(p, b.buf[b.cur:b.len])
		b.cur += n
		return n, nil
	}
	return 0, io.EOF
}

func (b *StringBuffer) ReadByte() (byte, error) {
	if b.cur < b.len {
		r := b.buf[b.cur]
		b.cur++
		return r, nil
	}
	return 0, io.EOF
}

// Write appends the contents of p to b's buffer.
// Write always returns len(p), nil.
func (b *StringBuffer) Write(p []byte) (int, error) {
	n := len(p)
	// 注意: 此处是小于(无等于)
	if b.cap < b.len+n {
		b.Grow(b.cap + n)
	}
	copy(b.buf[b.len:], p)
	b.len += n
	return n, nil
}

// WriteByte appends the byte c to b's buffer.
// The returned error is always nil.
func (b *StringBuffer) WriteByte(c byte) error {
	// 注意: 此处是小于或等于
	if b.cap <= b.len {
		b.Grow(b.cap + 1)
	}
	b.buf[b.len] = c
	b.len++
	return nil
}

// WriteRune appends the UTF-8 encoding of Unicode code point r to b's buffer.
// It returns the length of r and a nil error.
func (b *StringBuffer) WriteRune(r rune) (int, error) {
	if r < utf8.RuneSelf {
		if b.cap < b.len+1 {
			b.Grow(b.cap + 1)
		}
		b.buf[b.len] = byte(r)
		b.len++
		return 1, nil
	}
	if b.cap < b.len+utf8.UTFMax {
		b.Grow(b.cap + utf8.UTFMax)
	}
	n := utf8.EncodeRune(b.buf[b.len:b.len+utf8.UTFMax], r)
	b.len += n
	return n, nil
}

func (b *StringBuffer) ReadFrom(r io.Reader) (int64, error) {
	n, err := r.Read(b.buf[b.len:b.cap])
	if err != nil && err != io.EOF {
		return 0, err
	}
	b.len += n
	return int64(n), nil
}

func (b *StringBuffer) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(b.buf[b.cur:b.len])
	if err != nil && err != io.EOF {
		return 0, err
	}
	return int64(n), nil
}

// WriteString appends the contents of s to b's buffer.
// It returns the length of s and a nil error.
func (b *StringBuffer) WriteString(s string) (int, error) {
	n := len(s)
	if b.cap < b.len+n {
		b.Grow(b.cap + n)
	}
	copy(b.buf[b.len:], s)
	b.len += n
	return n, nil
}

var _ io.ReadWriter = (*StringBuffer)(nil)
var _ io.ByteReader = (*StringBuffer)(nil)
var _ io.ByteWriter = (*StringBuffer)(nil)
var _ io.Seeker = (*StringBuffer)(nil)
var _ io.Closer = (*StringBuffer)(nil)
