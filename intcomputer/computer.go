package intcomputer

import (
	"log"
	"strconv"
	"strings"
)

type OpCode struct {
	Code     int
	ArgCount int
	Args     []int
	Executor func([]int, []int, int) int
}

var supportedCodes = map[int]OpCode{1: {ArgCount: 3, Executor: func(args []int, mem []int, ip int) int {
	mem[args[2]] = mem[args[0]] + mem[args[1]]
	return ip + len(args) + 1
}}, 2: {ArgCount: 3, Executor: func(args []int, mem []int, ip int) int {
	mem[args[2]] = mem[args[0]] * mem[args[1]]
	return ip + len(args) + 1
}}, 99: {ArgCount: 0}}

type IntComputer struct {
	Memory             []int
	InstructionPointer int
	Program            []OpCode
}

func NewComputer(initialMemory string, noun int, verb int) *IntComputer {
	tokens := strings.Split(initialMemory, ",")
	memory := make([]int, len(tokens))
	for i := 0; i < len(tokens); i++ {
		val, err := strconv.Atoi(tokens[i])
		if err != nil {
			log.Fatal("Input contained non-integer value", err)
		}
		memory[i] = val
	}
	memory[1] = noun
	memory[2] = verb
	return &IntComputer{Memory: memory, InstructionPointer: 0}
}

func (c *IntComputer) getNextInstruction() OpCode {
	code := c.Memory[c.InstructionPointer]
	var args []int
	for i := 1; i <= supportedCodes[code].ArgCount; i++ {
		args = append(args, c.Memory[c.InstructionPointer+i])
	}
	return OpCode{Code: code, Args: args, Executor: supportedCodes[code].Executor}
}

func (c *IntComputer) ExecuteNextInstruction() bool {
	opCode := c.getNextInstruction()
	switch opCode.Code {
	case 99:
		return true
	default:
		c.InstructionPointer = opCode.Executor(opCode.Args, c.Memory, c.InstructionPointer)
	}

	return false
}
