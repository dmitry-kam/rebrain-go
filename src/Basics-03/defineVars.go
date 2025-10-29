package main

import (
	"fmt"
	"unicode/utf8"
)

// Universe block - глобальные переменные пакета
var globalVar string = "global"
var myVar string = "Variable string"
var anotherVar string
var noType = 100
var a, b, c int
var d, e, f = "hello", 42, true
var (
	price       int
	qty         int
	isDeletable bool
)
var packageLevel = "package level"

const statusCode int = 200
const (
	orderStatusNew string = "new"
	baseDiscount          = 3.5
)

func main() {
	//fmt.Println(myVar) /// without \n but with %
	// Println - выводит аргументы через пробел + добавляет \n
	fmt.Println(myVar)
	fmt.Println(anotherVar)
	// Printf - форматирует строку согласно спецификаторам
	fmt.Printf("\nnoType: %d", noType)
	fmt.Printf("\nCodes of a,b,c: %c(%d) %c(%d) %c(%d)", a, a, b, b, c, c)
	a, b, c = 97, 98, 99

	fmt.Printf("\nSum a,b,c: %d\n", (a + b + c))

	fmt.Println(d, string(e), f)
	fmt.Printf("Price: %d, Qty: %d, Deletable: %t",
		price,
		qty,
		isDeletable)

	pathToFile := "path"
	fmt.Println(pathToFile)
	str, number, isExist := "new string", 42, false
	fmt.Println(str, number, isExist)

	isExist = true
	fmt.Println(isDeletable && isExist)
	fmt.Println(isDeletable || isExist)

	//str, number := "override string", 10 // static analyzer: no new variables on left side of :=

	str1, _ := "2", 3
	fmt.Println(str1)

	//////////////////
	const (
		untypedNum     = 15
		typedNum   int = 10
	)
	var c32 int32 = 30

	fmt.Println(c32 + untypedNum)
	fmt.Println(statusCode + typedNum)
	//fmt.Println(c32 + typedNum)       // static analyzer:  invalid operation: c + typedNum (mismatched types int32 and int)

	const (
		num1 = 21
		num2
		num3 = 10
		num4
	)
	fmt.Println(num1, num2, num3, num4) // 21 21 10 10

	////////////
	const (
		_       = 10 * iota
		speed10 = 10 * iota
		speed20 = 10 * iota
		speed30
		speed40
	)

	fmt.Println(speed10, speed20, speed30, speed40) // 10 20 30 40

	// bad
	const (
		CustomSpeed0  = 10 * iota // 0
		CustomSpeed8  = 8 * iota  // 8
		CustomSpeed32 = 16 * iota // 32
	)

	fmt.Println("Скорости 2:", CustomSpeed0, CustomSpeed8, CustomSpeed32)

	////////////
	fl := 4.54
	fmt.Printf("%T\n", fl) // float64

	////////
	var strChar string = "!"
	rCode, _ := utf8.DecodeRuneInString(strChar)
	fmt.Printf("%U\n", rCode) // U+0021

	var runeChar rune = '\x21'
	fmt.Println(string(runeChar)) // !

	var hello = "Hello"
	fmt.Println(hello + " from RebrainMe!")

	/////////////////////
	var greeting = "Привет!"
	fmt.Println(len(greeting))                    // 13 байт
	fmt.Println(utf8.RuneCountInString(greeting)) // 7 символов (рун)
	greeting = greeting + " from RebrainMe!"
	fmt.Println(len(greeting))
	fmt.Println(utf8.RuneCountInString(greeting))

	fmt.Println(greeting[4:6]) // и
	fmt.Println(greeting[:6])  // При
	fmt.Println(greeting[:])   // Привет!

	//greeting[0] = "Л"                   // static analyzer: cannot assign to greeting[0]
	var convGreeting = []rune(greeting)
	convGreeting[4] = 'Е'

	fmt.Println(string(convGreeting)) //  ПривЕт!

	/////////////////
	var cmplx complex128 = 1.1 + 2.1i
	cmplx3 := complex(2.1, 2)

	//результат
	//(1.1+2.1i) (2.1+2i)
	fmt.Println(cmplx, cmplx3)

	////////////
	var defArr [3]string = [3]string{"one", "two", "three"}
	fmt.Println(defArr) // [one two three]

	var array [3]int
	fmt.Println(array) // [0 0 0]

	////////////////
	dynamic := [...]bool{4: true}
	fmt.Println(dynamic)    // [false false false false true]
	fmt.Println(dynamic[2]) // false

	/////////////
	fmt.Println("------------------")
	arr := dynamic
	arr[2] = true
	fmt.Println(dynamic) // [false false false false true]
	fmt.Println(arr)     // [false false true false true]

	fmt.Println("------------------")
	arr2 := &dynamic
	arr2[2] = true
	fmt.Println(dynamic) // [false false true false true]

	fmt.Println("------------------")
	sl := dynamic[:2]
	sl[1] = true
	fmt.Println(dynamic) // [false true true false true]
	fmt.Println(sl)      // [false true]

	////////////////
	fmt.Println("------------------")
	type SettlementId string

	var cityId SettlementId = "228edf0c-0c0d-4db6-bf0e-d508e68270a3"
	var strCityId string = "228edf0c-0c0d-4db6-bf0e-d508e68270a3"

	fmt.Println(cityId == SettlementId(strCityId))

	//////////////////
	fmt.Println("\nglobalVar из main:", globalVar) // "global"
	localVar := "local in main"
	fmt.Println("\nlocalVar в main:", localVar) // "local in main"

	// Блок инструкции if
	if true {
		// Этот localVar перекрывает переменную из main
		localVar := "local in if block"
		fmt.Println("\nlocalVar в if:", localVar)    // "local in if block"
		fmt.Println("\nglobalVar из if:", globalVar) // "global"

		// Новая переменная только в блоке if
		ifOnlyVar := "only in if"
		fmt.Println("\nifOnlyVar:", ifOnlyVar) // "only in if"
	}

	// Здесь снова видна original localVar из main
	fmt.Println("\nlocalVar после if:", localVar) // "local in main"
	// fmt.Println(ifOnlyVar) // ОШИБКА: ifOnlyVar не видна здесь

	fmt.Println("\n\n=== Циклы и условия ===")

	// Блок for
	for i := 0; i < 2; i++ {
		loopVar := "in loop"
		fmt.Printf("i=%d, loopVar=%s\n", i, loopVar)
	}
	// i и loopVar не видны здесь
	// fmt.Println(i) // ОШИБКА
	// fmt.Println(loopVar) // ОШИБКА

	// Блок switch
	value := 1
	switch value {
	case 1:
		caseVar := "in case 1"
		fmt.Println("\ncaseVar:", caseVar)
		// Эта переменная видна только в этом case
	case 2:
		// caseVar не видна здесь - другой блок
		// fmt.Println(caseVar) // ОШИБКА
	}

	///////////////
	fmt.Println("\n\n=== Перекрытие глобальных переменных ===")

	fmt.Println("\nДо перекрытия:", packageLevel) // "package level"

	// Локальная переменная перекрывает глобальную
	packageLevel := "local shadow"
	fmt.Println("\nПосле перекрытия:", packageLevel) // "local shadow"

	// Чтобы обратиться к глобальной переменной, нужно использовать пакет
	// Но в этом случае она в том же пакете, поэтому напрямую нельзя
	// Вместо этого можно использовать другую функцию
	showGlobal()
}

func showGlobal() {
	fmt.Println("\nИз другой функции:", packageLevel) // "package level"
}
