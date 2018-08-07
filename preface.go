package log

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var pf *Preface

func init() {
	pf = NewPreface()
}

//Preface 前言结构体
type Preface struct {
	KV map[string]interface{}
}

//NewPreface new instance
func NewPreface() *Preface {
	p := new(Preface)
	p.KV = make(map[string]interface{})
	return p
}

//Reset 充值单例
func Reset() {
	pf = NewPreface()
}

//ShowPreface 展示前言(直接展示一条)
func ShowPreface(k string, v interface{}) {
	NewPreface().Set(k, v).Show()
}

//SetPreface 设置
func SetPreface(k string, v interface{}) *Preface { return pf.Set(k, v) }

//Set 设置前言
func (p *Preface) Set(k string, v interface{}) *Preface {
	p.KV[k] = v
	return p
}

//SetMultiPreface 批量设置
func SetMultiPreface(m map[string]string) *Preface { return pf.SetMulti(m) }

//SetMulti 批量设置前言
func (p *Preface) SetMulti(m map[string]string) *Preface {
	for k, v := range m {
		p.KV[k] = v
	}
	return p
}

//Show 展示前言
func (p *Preface) Show() *Preface {
	c := color.New(color.FgHiMagenta)
	c.Println(strings.Repeat("-", 64))
	for k, v := range p.KV {
		c.Println(fmt.Sprintf("%-16v |%-v", k, v))
	}
	c.Println(strings.Repeat("-", 64))
	return p
}
