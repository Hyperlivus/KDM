package parser

import (
	slices "golang.org/x/exp/slices"
	"errors"
	"unicode"
	
)

const UNREQUIRED_CHAR = "unrequired char at position"

type Cursor struct {
	index int
	code []byte
}
func (cursor *Cursor) Current() byte {
	cursor.EatSpaces()
	return cursor.code[cursor.index]
}
func (cursor *Cursor) IsEnd() bool {
	return cursor.index >= len(cursor.code)
}
func (cursor *Cursor) EatSpaces() {
	curr := cursor.code[cursor.index]
	for unicode.IsSpace(rune(curr)){
		cursor.index += 1
		curr = cursor.code[cursor.index]
	}
}
func (cursor *Cursor) RequireNext(required ...byte) (byte, error){
	cursor.EatSpaces()
	curr := cursor.code[cursor.index]
	if slices.Contains(required, curr){
		cursor.index++
		return curr, nil
	} 

	return 0, errors.New(UNREQUIRED_CHAR)
}

func NewCursor(code []byte) *Cursor {
	return &Cursor{
		code :code,
		index: 0,
	}
}