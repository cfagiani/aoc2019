package intcomputer

import (
	"log"
	"strconv"
	"strings"
	"sync"
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
	Executor func([]int, map[int]int, int, int, []int, InputSupplier, OutputConsumer) (int, int)
}

var supportedCodes = map[int]OpCode{
	1: {Name: "Add", ArgCount: 3, Executor: func(args []int, mem map[int]int, ip int, relBase int, modes []int, input InputSupplier, output OutputConsumer) (int, int) {
		mem[args[2]] = getParamVal(args[0], modes[0], relBase, mem) + getParamVal(args[1], modes[1], relBase, mem)
		return ip + len(args) + 1, relBase
	}}, 2: {Name: "Multiply", ArgCount: 3, Executor: func(args []int, mem map[int]int, ip int, relBase int, modes []int, input InputSupplier, output OutputConsumer) (int, int) {
		mem[args[2]] = getParamVal(args[0], modes[0], relBase, mem) * getParamVal(args[1], modes[1], relBase, mem)
		return ip + len(args) + 1, relBase
	}}, 3: {Name: "Input", ArgCount: 1, Executor: func(args []int, mem map[int]int, ip int, relBase int, modes []int, input InputSupplier, output OutputConsumer) (int, int) {
		mem[args[0]] = input.GetNextInput()
		return ip + len(args) + 1, relBase
	}}, 4: {Name: "Output", ArgCount: 1, Executor: func(args []int, mem map[int]int, ip int, relBase int, modes []int, input InputSupplier, output OutputConsumer) (int, int) {

		output.WriteOutput(getParamVal(args[0], modes[0], relBase, mem))
		return ip + len(args) + 1, relBase
	}}, 5: {Name: "Jump-if-true", ArgCount: 2, Executor: func(args []int, mem map[int]int, ip int, relBase int, modes []int, input InputSupplier, output OutputConsumer) (int, int) {
		val := getParamVal(args[0], modes[0], relBase, mem)
		if val != 0 {
			return getParamVal(args[1], modes[1], relBase, mem), relBase
		} else {
			return ip + len(args) + 1, relBase
		}
	}}, 6: {Name: "Jump-if-false", ArgCount: 2, Executor: func(args []int, mem map[int]int, ip int, relBase int, modes []int, input InputSupplier, output OutputConsumer) (int, int) {
		val := getParamVal(args[0], modes[0], relBase, mem)
		if val == 0 {
			return getParamVal(args[1], modes[1], relBase, mem), relBase
		} else {
			return ip + len(args) + 1, relBase
		}
	}}, 7: {Name: "Less-Than", ArgCount: 3, Executor: func(args []int, mem map[int]int, ip int, relBase int, modes []int, input InputSupplier, output OutputConsumer) (int, int) {

		if getParamVal(args[0], modes[0], relBase, mem) < getParamVal(args[1], modes[1], relBase, mem) {
			mem[args[2]] = 1
		} else {
			mem[args[2]] = 0
		}
		return ip + len(args) + 1, relBase
	}}, 8: {Name: "Equals", ArgCount: 3, Executor: func(args []int, mem map[int]int, ip int, relBase int, modes []int, input InputSupplier, output OutputConsumer) (int, int) {
		if getParamVal(args[0], modes[0], relBase, mem) == getParamVal(args[1], modes[1], relBase, mem) {
			mem[args[2]] = 1
		} else {
			mem[args[2]] = 0
		}
		return ip + len(args) + 1, relBase
	}}, 9: {Name: "Adj Rel Base", ArgCount: 1, Executor: func(args []int, mem map[int]int, ip int, relBase int, modes []int, input InputSupplier, output OutputConsumer) (int, int) {

		return ip + len(args) + 1, relBase + getParamVal(args[0], modes[0], relBase, mem)
	}},
	99: {ArgCount: 0}}

type IntComputer struct {
	Memory             map[int]int
	InstructionPointer int
	Program            []OpCode
	Input              InputSupplier
	Output             OutputConsumer
	RelativeBase       int
}

type ArrayInputSupplier struct {
	Inputs []int
	curIdx int
}

type ChannelIOSupplier struct {
	IOChannel chan int
}

func NewChannelIOSupplier(inputSeed []int) *ChannelIOSupplier {
	supplier := &ChannelIOSupplier{
		// buffer up to 100 values
		IOChannel: make(chan int, 100),
	}
	seedInput(supplier.IOChannel, inputSeed)
	return supplier
}

func seedInput(channel chan int, inputSeed []int) {
	for _, input := range inputSeed {
		channel <- input
	}
}

func (c *ChannelIOSupplier) GetNextInput() int {
	return <-c.IOChannel
}

func (c *ChannelIOSupplier) WriteOutput(val int) {
	c.IOChannel <- val
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

func getParamVal(param int, mode int, relBase int, memory map[int]int) int {
	if mode == 0 {
		return memory[param]
	} else if mode == 1 {
		return param
	} else {
		return memory[param+relBase]
	}
}

func NewComputer(initialMemory string, noun int, verb int, setParams bool, input InputSupplier, output OutputConsumer) *IntComputer {
	tokens := strings.Split(initialMemory, ",")
	memory := make(map[int]int)
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
		c.InstructionPointer, c.RelativeBase = opCode.Executor(opCode.Args, c.Memory, c.InstructionPointer, c.RelativeBase, opCode.Modes, c.Input, c.Output)
	}

	return false
}

func (c *IntComputer) RunToCompletion(wg *sync.WaitGroup) {
	for ; !c.ExecuteNextInstruction(); {
		// do nothing
	}
	if wg != nil {
		wg.Done()
	}
}
