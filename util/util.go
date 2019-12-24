package util // Package util provides utility functions that are not specific to any single problem from AoC.
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

//Checks the error argument and, if it is not nil, it will log the msg passed in. If isFatal is true, the log will be
//written as Fatal which will cause exit(1) to be called.
func CheckError(err error, msg string, isFatal bool) bool {
	if err != nil {
		if isFatal {
			log.Fatal(msg, err)
		} else {
			log.Println(msg)
		}
		return true
	}
	return false
}

//returns the difference between arr1 and arr2
func FilterArray(arr1 []string, arr2 []string) []string {
	var result []string
	for _, v := range arr1 {
		if !IsStringInSlice(v, arr2) {
			result = append(result, v)
		}
	}
	return result
}

// Returns the full contents of a file as a string. If the file cannot be read, it will log a Fatal error and exit the program.
func ReadFileAsString(fname string) string {
	dat, err := ioutil.ReadFile(fname)
	CheckError(err, "Could not read file", true)
	return string(dat)
}

// Reads all lines in a file and returns them as an array of strings. If an error is encountered, panic.
func ReadAllLines(fname string) []string {
	file, err := os.Open(fname)
	var results []string
	if !CheckError(err, "cannot open file", true) {
		reader := bufio.NewReader(file)
		line, err := Readln(reader)
		for err == nil {
			results = append(results, line)
			line, err = Readln(reader)
		}
	}
	return results
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

//returns true if a is in the list passed in
func IsIntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func PrintByteArray(state [][]byte) {
	fmt.Printf("\n")
	for i := 0; i < len(state); i++ {
		for j := 0; j < len(state[i]); j++ {
			fmt.Printf("%s", string(state[i][j]))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// Converts a string containing integers into a int array
func GetStringAsIntArray(input string) []int {
	output := make([]int, len(input))
	for i, c := range input {
		n, _ := strconv.Atoi(string(c))
		output[i] = n
	}
	return output
}

// Returns an array of slices such that all permutations of the input array are represented.
// Uses Heap's algorithm (https://en.wikipedia.org/wiki/Heap%27s_algorithm)
func AllPermutations(arr []int) [][]int {
	var permutationGenerator func([]int, int)
	var allPermutations [][]int

	permutationGenerator = func(arr []int, n int) {
		if n == 1 {
			// if there is only 1 element, there is nothing to do
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			allPermutations = append(allPermutations, tmp)
		} else {
			for i := 0; i < n; i++ {
				permutationGenerator(arr, n-1)
				if n%2 == 1 {
					// n is odd
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					// n is even
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	permutationGenerator(arr, len(arr))
	return allPermutations
}
