package main

import (
	"fmt"
	"math/rand"
	randV2 "math/rand/v2"
	"runtime"
	"slices"
	"sort"
	"strings"
	"time"
)

type Profiler struct {
	startTime time.Time
	startMem  runtime.MemStats
}

type functionResult struct {
	duration time.Duration
	memory   uint64
}

const (
	Letters          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers          = "0123456789"
	AlphaNumeric     = Letters + Numbers
	stringLength     = 4
	sliceLength      = 1000
	numRange         = 32767
	showMeasurements = true
)

var (
	functionsResults = make(map[string]functionResult, 4)
)

func main() {
	var (
		needleX string
		sliceA  []string
		sliceN  []int
	)

	measure("Generating String data", func() {
		needleX = randomString(stringLength, AlphaNumeric)
		sliceA = make([]string, sliceLength)
		for i := range sliceA {
			sliceA[i] = randomString(stringLength, AlphaNumeric)
		}
		fmt.Printf(getBoldText("Generating data...\n"))
		fmt.Printf(
			"X=%q, slice A with %d values: [%v ... %v]\n",
			needleX, len(sliceA),
			strings.Join(sliceA[:3], ", "),
			strings.Join(sliceA[sliceLength-3:], ", "))
	})

	measure("Generating Int data", func() {
		sliceN = make([]int, sliceLength)
		for i := range sliceA {
			sliceN[i] = randV2.IntN(numRange)
		}
		fmt.Printf("Numeric slice with %d values: %v ... %v\n\n", len(sliceN), sliceN[:3], sliceN[sliceLength-3:])
	})

	/////////////////////////
	fmt.Println(strings.Repeat("#", 30))
	var included bool
	measure("contains()", func() {
		included = contains(sliceA, needleX)
	})
	fmt.Printf("contains() says X=%q in A: %t\n", needleX, included)

	measure("primitiveContains()", func() {
		included = primitiveContains(sliceA, needleX)
	})
	fmt.Printf("primitiveContains() says X=%q in A: %t\n", needleX, included)
	fmt.Println(strings.Repeat("#", 30))

	/////////////////////////
	var resultMax int
	measure("getMax()", func() {
		resultMax = getMax(sliceN...)
	})
	fmt.Printf("getMax() says Max element in numeric slice = %d\n", resultMax)
	/////////////////////////
	measure("primitiveGetMax()", func() {
		resultMax = primitiveGetMax(sliceN...)
	})
	fmt.Printf("primitiveGetMax() says Max element in numeric slice = %d\n", resultMax)

	if showMeasurements {
		showFuncResultsTable()
	}
}

func contains(a []string, x string) bool {
	return slices.Contains(a, x)
}

func primitiveContains(a []string, x string) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

// with library function - O(n log n)
func getMax(numbers ...int) int {
	sort.Ints(numbers)
	return numbers[len(numbers)-1]
}

// O(n)
func primitiveGetMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	maxValue := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > maxValue {
			maxValue = numbers[i]
		}
	}
	return maxValue
}

func StartProfiling() *Profiler {
	p := &Profiler{
		startTime: time.Now(),
	}
	runtime.ReadMemStats(&p.startMem)
	return p
}

func (p *Profiler) Stop() (time.Duration, uint64) {
	duration := time.Since(p.startTime)

	var endMem runtime.MemStats
	runtime.ReadMemStats(&endMem)
	memoryUsed := endMem.Alloc - p.startMem.Alloc

	return duration, memoryUsed
}

func measure(funcName string, f func()) {
	p := StartProfiling()
	f()
	duration, memory := p.Stop()
	functionsResults[funcName] = functionResult{duration, memory}
}

func getBoldText(text string) string {
	return "\033[1m" + text + "\033[0m"
}

func randomString(length int, charset string) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func showFuncResultsTable() {
	fmt.Println()
	fmt.Printf("┌%-25s┬%-12s┬%-30s┐\n", strings.Repeat("─", 25), strings.Repeat("─", 12), strings.Repeat("─", 30))
	fmt.Printf(getBoldText("│%-25s│%-12s│%-30s│\n"), "Function", "Time", "Memory")
	fmt.Printf("├%-25s┼%-12s┼%-30s┤\n", strings.Repeat("─", 25), strings.Repeat("─", 12), strings.Repeat("─", 30))

	for funcName, funcResult := range functionsResults {
		fmt.Printf("│%-25s│%-12v│%-30s│\n", funcName, funcResult.duration, fmt.Sprintf("%d bytes (%.2f KB)", funcResult.memory, float64(funcResult.memory)/1024))
	}

	fmt.Printf("└%-25s┴%-12s┴%-30s┘\n", strings.Repeat("─", 25), strings.Repeat("─", 12), strings.Repeat("─", 30))
}
