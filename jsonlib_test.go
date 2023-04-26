package jsonlib

import (
	"fmt"
	"testing"
)

type Class struct {
	Name string `json:"name,omitempty"`
	Code int64  `json:"code,omitempty"`
}

type Student struct {
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Classes []Class `json:"classes,omitempty"`
	Score   float64 `json:"score,omitempty"`
}

var student = &Student{
	Name: "张三",
	Age:  12,
	Classes: []Class{
		{Name: "文化班", Code: 1001},
		{Name: "数学班", Code: 1002},
	},
}

const jsdata = `{"name":"张三","age":12,"classes":[{"name":"文化班","code":1001},{"name":"数学班","code":1002}]}`

func TestMarshal(t *testing.T) {
	bs, err := Marshal(student)
	if err != nil {
		panic(err)
	}
	str := string(bs)
	t.Logf("%s\n", str)
	if str != jsdata {
		t.Fail()
	}
}

func TestMarshalToString(t *testing.T) {
	ss, err := MarshalToString(student)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%s\n", ss)
	if ss != jsdata {
		t.Fail()
	}
}

func TestUnmarshal(t *testing.T) {
	var s *Student
	err := Unmarshal([]byte(jsdata), &s)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s)
}

func TestNewDecoder(t *testing.T) {

}

func TestNewEncoder(t *testing.T) {

}

func TestMarshalIgnoreOmitempty(t *testing.T) {
	bs, err := MarshalIgnoreOmitempty(student)
	if err != nil {
		panic(err)
	}
	str := string(bs)
	t.Logf("%s\n", str)
	if str != jsdata {
		t.Fail()
	}
}

func TestMarshalToStringIgnoreOmitempty(t *testing.T) {

}

func TestNewEncoderIgnoreOmitempty(t *testing.T) {

}
