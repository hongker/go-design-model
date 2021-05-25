package template

import "fmt"

type DocSuper struct {
	GetContent func() string
}

func (doc DocSuper) Read() {
	fmt.Println("read:", doc.GetContent())
}

type LocalDoc struct {
	DocSuper
}

func (doc LocalDoc) GetContent() string {
	return "local doc"
}

func NewLocalDoc() *LocalDoc {
	doc := new(LocalDoc)
	doc.DocSuper.GetContent = doc.GetContent
	return doc
}

type NetworkDoc struct {
	DocSuper
}

func (doc NetworkDoc) GetContent() string {
	return "local doc"
}

func NewNetworkDoc() *NetworkDoc {
	doc := new(NetworkDoc)
	doc.DocSuper.GetContent = doc.GetContent
	return doc
}