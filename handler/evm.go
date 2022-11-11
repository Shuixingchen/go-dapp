package handler

import (
	"fmt"
	"strings"

	"github.com/Shuixingchen/go-dapp/utils"
	"github.com/ethereum/go-ethereum/core/asm"
)

type Disassemble struct {
	ByteCode       []byte
	Instructions   []*Instruction
	FunctionsHashs map[string]string
}
type Instruction struct {
	PC  string // 指令的地址
	Op  string // opcode
	Arg string // 参数
}

// 通过反编译获取function
func CheckContractFunction(contractAddr string) {
	d := NewDisassemble()
	d.Start(GetCode(contractAddr))
	funcs := d.GetFunctions()
	isERC20 := matchFunctions(funcs, utils.ERC20FunctionSig)
	fmt.Println(isERC20)
}

// byteCode to opcode
func NewDisassemble() *Disassemble {
	var d Disassemble
	d.FunctionsHashs = utils.LoadJsonFile("functionHashes")
	return &d
}

func (d *Disassemble) Start(byteCode []byte) {
	d.ByteCode = byteCode
	d.Instructions = make([]*Instruction, 0)
	it := asm.NewInstructionIterator(d.ByteCode)
	for it.Next() {
		var inst Instruction
		inst.PC = fmt.Sprintf("%05x", it.PC())
		inst.Op = fmt.Sprintf("%v", it.Op())
		if it.Arg() != nil && 0 < len(it.Arg()) {
			inst.Arg = fmt.Sprintf("%x", it.Arg())
		}
		d.Instructions = append(d.Instructions, &inst)
	}
}

func (d *Disassemble) GetFunctions() map[string]bool {
	res := make(map[string]bool, 0)
	for _, instruction := range d.Instructions {
		if strings.EqualFold(instruction.Op, "PUSH4") && instruction.Arg != "" {
			res[instruction.Arg] = true
		}
	}
	return res
}

func matchFunctions(funcs map[string]bool, mustFuncs []string) bool {
	for _, f := range mustFuncs {
		if _, ok := funcs[f]; !ok {
			return false
		}
	}
	return true
}
