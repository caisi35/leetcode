package main

import (
	"container/list"
	"fmt"
)

func main() {
	c := Constructor()
	c.AddText("leetcode")
	fmt.Println(c.DeleteText(4))
	c.AddText("practice")
	fmt.Println(c.CursorRight(3))
	fmt.Println(c.CursorLeft(8))
	fmt.Println(c.DeleteText(10))
	fmt.Println(c.CursorLeft(2))
	fmt.Println(c.CursorRight(6))
}

type TextEditor struct {
	editor *list.List
	cursor *list.Element
}

func Constructor() TextEditor {
	editor := list.New()
	editor.PushBack(nil)
	return TextEditor{editor: editor, cursor: editor.Back()}
}

func (this *TextEditor) AddText(text string) {
	for _, c := range text {
		this.editor.InsertBefore(byte(c), this.cursor)
	}
}

func (this *TextEditor) DeleteText(k int) int {
	count := 0
	for ; k > 0 && this.cursor != this.editor.Front(); k-- {
		this.editor.Remove(this.cursor.Prev())
		count++
	}
	return count
}

func (this *TextEditor) CursorLeft(k int) string {
	for ; k > 0 && this.cursor != this.editor.Front(); k-- {
		this.cursor = this.cursor.Prev()
	}
	head := this.cursor
	for i := 0; i < 10 && head != this.editor.Front(); i++ {
		head = head.Prev()
	}
	var ret []byte
	for ; head != this.cursor; head = head.Next() {
		ret = append(ret, head.Value.(byte))
	}
	return string(ret)
}

func (this *TextEditor) CursorRight(k int) string {
	for ; k > 0 && this.cursor != this.editor.Back(); k-- {
		this.cursor = this.cursor.Next()
	}
	head := this.cursor
	for i := 0; i < 10 && head != this.editor.Front(); i++ {
		head = head.Prev()
	}
	var ret []byte
	for ; head != this.cursor; head = head.Next() {
		ret = append(ret, head.Value.(byte))
	}
	return string(ret)
}
