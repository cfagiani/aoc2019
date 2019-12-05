package intcomputer

import (
	"log"
	"strconv"
	"strings"
)

type InputSupplier interface {
	GetNextInput() int
}

type OutputConsumer interface {
	WriteOutput(int)
}

type OpCode struct {
	Code     int
	Modes    []int
	ArgCount int
	Args     []int
	Mode     int
	Name     string
	Executor func([]int, []int, int, []int, InputSupplier, OutputConsumer) int
}

var supportedCodes = map[int]OpCode{
	1: {Name: "Add", ArgCount: 3, Executor: func(args []int, mem []int, ip int, modes []int, input InputSupplier, output OutputConsumer) int {
		mem[args[2]] = getParamVal(args[0], modes[0], mem) + getParamVal(args[1], modes[1], mem)
		return ip + len(args) + 1
	}}, 2: {Name: "Multiply", ArgCount: 3, Executor: func(args []int, mem []int, ip int, modes []int, input InputSupplier, output OutputConsumer) int {
		mem[args[2]] = getParamVal(args[0], modes[0], mem) * getParamVal(args[1], modes[1], mem)
		return ip + len(args) + 1
	}}, 3: {Name: "Input", ArgCount: 1, Executor: func(args []int, mem []int, ip int, modes []int, input InputSupplier, output OutputConsumer) int {
		mem[args[0]] = input.GetNextInput()
		return ip + len(args) + 1
	}}, 4: {Name: "Output", ArgCount: 1, Executor: func(args []int, mem []int, ip int, modes []int, input InputSupplier, output OutputConsumer) int {
		if modes[0] == 0 {
			output.WriteOutput(mem[args[0]])
		} else {
			output.WriteOutput(args[0])
		}
		return ip + len(args) + 1
	}}, 5: {Name: "Jump-if-true", ArgCount: 2, Executor: func(args []int, mem []int, ip int, modes []int, input InputSupplier, output OutputConsumer) int {
		val := getParamVal(args[0], modes[0], mem)
		if val != 0 {
			return getParamVal(args[1], modes[1], mem)
		} else {
			return ip + len(args) + 1
		}
	}}, 6: {Name: "Jump-if-false", ArgCount: 2, Executor: func(args []int, mem []int, ip int, modes []int, input InputSupplier, output OutputConsumer) int {
		val := getParamVal(args[0], modes[0], mem)
		if val == 0 {
			return getParamVal(args[1], modes[1], mem)
		} else {
			return ip + len(args) + 1
		}
	}}, 7: {Name: "Less-Than", ArgCount: 3, Executor: func(args []int, mem []int, ip int, modes []int, input InputSupplier, output OutputConsumer) int {

		if getParamVal(args[0], modes[0], mem) < getParamVal(args[1], modes[1], mem) {
			mem[args[2]] = 1
		} else {
			mem[args[2]] = 0
		}
		return ip + len(args) + 1
	}}, 8: {Name: "Equals", ArgCount: 3, Executor: func(args []int, mem []int, ip int, modes []int, input InputSupplier, output OutputConsumer) int {
		if getParamVal(args[0], modes[0], mem) == getParamVal(args[1], modes[1], mem) {
			mem[args[2]] = 1
		} else {
			mem[args[2]] = 0
		}
		return ip + len(args) + 1
	}},
	99: {ArgCount: 0}}

type IntComputer struct {
	Memory             []int
	InstructionPointer int
	Program            []OpCode
	Input              InputSupplier
	Output             OutputConsumer
}

type ArrayInputSupplier struct {
	Inputs []int
	curIdx int
}

func NewInputSupplier(data []int) *ArrayInputSupplier {
	return &ArrayInputSupplier{Inputs: data, curIdx: 0}
}

type ArrayOutputSupplier struct {
	Output []int
}

func NewOutputSupplier() *ArrayOutputSupplier {
	return &ArrayOutputSupplier{}
}

func (out *ArrayOutputSupplier) WriteOutput(val int) {
	out.Output = append(out.Output, val)
}

func (in *ArrayInputSupplier) GetNextInput() int {
	val := in.Inputs[in.curIdx]
	in.curIdx++
	return val
}

func getParamVal(param int, mode int, memory []int) int {
	if mode == 0 {
		return memory[param]
	} else {
		return param
	}
}

func NewComputer(initialMemory string, noun int, verb int, setParams bool, input InputSupplier, output OutputConsumer) *IntComputer {
	tokens := strings.Split(initialMemory, ",")
	memory := make([]int, len(tokens))
	for i := 0; i < len(tokens); i++ {
		val, err := strconv.Atoi(tokens[i])
		if err != nil {
			log.Fatal("Input contained non-integer value", err)
		}
		memory[i] = val
	}
	if setParams {
		memory[1] = noun
		memory[2] = verb
	}
	return &IntComputer{Memory: memory, InstructionPointer: 0, Input: input, Output: output}
}

func (c *IntComputer) getNextInstruction() OpCode {
	packedCode := c.Memory[c.InstructionPointer]
	code := packedCode % 100
	codeString := strconv.Itoa(packedCode)
	idx := len(codeString) - 2
	var args []int
	var modes []int
	for i := 1; i <= supportedCodes[code].ArgCount; i++ {
		args = append(args, c.Memory[c.InstructionPointer+i])
		if idx-i >= 0 {
			mode, _ := strconv.Atoi(string(codeString[idx-i]))
			modes = append(modes, mode)
		} else {
			modes = append(modes, 0)
		}

	}
	return OpCode{Code: code, Args: args, Modes: modes, Executor: supportedCodes[code].Executor}
}

func (c *IntComputer) ExecuteNextInstruction() bool {
	opCode := c.getNextInstruction()

	switch opCode.Code {
	case 99:
		return true
	default:
		c.InstructionPointer = opCode.Executor(opCode.Args, c.Memory, c.InstructionPointer, opCode.Modes, c.Input, c.Output)
	}

	return false
}
