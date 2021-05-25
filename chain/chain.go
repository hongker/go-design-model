package chain

import (
	"fmt"
	"strings"
)

type Filter interface {
	Handle(content string)
	SetNext(filter Filter)
}

type AbstractFilter struct {
	filter Filter
}

func (f *AbstractFilter) SetNext(filter Filter) {
	f.filter = filter
}

func (f *AbstractFilter) Delivery(content string) {
	if f.filter != nil {
		f.filter.Handle(content)
	}
}

type AdFilter struct {
	AbstractFilter
}

func (f *AdFilter) Handle(content string) {
	fmt.Println("ad before:", content)
	replacedContent := strings.ReplaceAll(content, "ad", "**")
	fmt.Println("ad after:", replacedContent)
	f.Delivery(replacedContent)
}

type SensitiveFilter struct {
	AbstractFilter
}
func (f *SensitiveFilter) Handle(content string) {
	fmt.Println("sensitive before:", content)
	replacedContent := strings.ReplaceAll(content, "sb", "**")
	fmt.Println("sensitive after:", replacedContent)
	f.Delivery(replacedContent)
}

