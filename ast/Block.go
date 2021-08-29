package ast

import (
	"errors"

	"github.com/tusklang/tusk/tokenizer"
)

type Block struct {
	BlockType string
	Sub       []*ASTNode
}

var bmatches = map[string]string{
	"{": "}",
	"(": ")",
}

func (b *Block) Parse(lex []tokenizer.Token, i *int) (e error) {

	if lex[*i].Type != "(" && lex[*i].Type != "{" {
		return errors.New("given lex is not a group")
	}

	b.BlockType = lex[*i].Type

	gcontent := grouper(braceMatcher(lex, i, lex[*i].Type, bmatches[lex[*i].Type], true, ""))
	b.Sub, e = groupsToAST(gcontent)

	if e != nil {
		return e
	}

	return nil
}
