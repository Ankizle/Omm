package data

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type Instruction struct {
	val value.Value
}

func NewInstruction(i value.Value) *Instruction {
	return &Instruction{
		val: i,
	}
}

func (i *Instruction) LLVal(block *ir.Block) value.Value {
	return i.val
}

func (i *Instruction) Type() types.Type {
	return i.val.Type()
}

func (i *Instruction) TypeString() string {
	return i.Type().LLString()
}
