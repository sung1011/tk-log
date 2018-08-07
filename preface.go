package log

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

//Preface 前言结构体
type Preface struct {
	KV map[string]interface{}
}

//NewPreface new instance
func NewPreface() *Preface {
	return new(Preface)
}

//ShowPreface 展示前言(直接展示一条)
func ShowPreface(k string, v interface{}) {
	NewPreface().Set(k, v).Show()
}

//Set 设置前言
func (p *Preface) Set(k string, v interface{}) *Preface {
	//PrefKV Preface值的单例
	p.KV = make(map[string]interface{})
	p.KV[k] = v
	return p
}

//MultSet 批量设置前言
func (p *Preface) MultSet(ks, vs []string) *Preface {
	if len(ks) != len(vs) {
		Erro("preface参数个数不相等")
	}
	p.KV = make(map[string]interface{})
	for k, v := range vs {
		p.KV[ks[k]] = v
	}
	return p
}

//Show 展示前言
func (p *Preface) Show() *Preface {
	c := color.New(color.FgHiMagenta)
	c.Println(strings.Repeat("-", 64))
	for k, v := range p.KV {
		c.Println(fmt.Sprintf("%s : \t%s", k, v))
	}
	c.Println(strings.Repeat("-", 64))
	return p
}
