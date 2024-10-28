// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func Calc(expression string) (float64, error) {
// 	var operations []rune
// 	for _, r := range expression {
// 		if r == '+' || r == '-' || r == '*' || r == '/' {
// 			operations = append(operations, r)
// 		}
// 	}
// 	var tmp float64
// 	nums := readNum(expression)


// 	res, _ := strconv.ParseFloat(nums[0], 64)
// 	for i := range operations {
// 		switch {
// 		case operations[i] == '+':
// 			tmp, _ = strconv.ParseFloat(nums[i+1], 64)
// 			res += tmp
// 		case operations[i] == '-':
// 			tmp, _ = strconv.ParseFloat(nums[i+1], 64)
// 			res -= tmp
// 		case operations[i] == '*':
// 			tmp, _ = strconv.ParseFloat(nums[i+1], 64)
// 			res /= tmp
// 		case operations[i] == '/':
// 			tmp, _ = strconv.ParseFloat(nums[i+1], 64)
// 			res /= tmp
// 		}
// 	}
// 	return res, nil
// }

// func merge(expression string, nums []string, operations []rune) []string{
// 	var splitExprrssion []string
// 	for n:=0;_,o:= range operoperations;n++{

// 	}
// }

// // func merge(expression string, nums []string, operations []rune) []string{
// // 	runes := []rune(expression)
// // 	for {
// // 		if i := strings.Index(string(operations), "("); i != -1 {
// // 			j := strings.Index(string(operations), ")")
// // 			//runes[i] = ' '
// // 			//runes[j] = ' '



// // 			continue
// // 		}
// // 		if i := strings.Index(string(runes), "-"); i >= 0 {
// // 			runes[i] = ' '
// // 			continue
// // 		}
// // 		if i := strings.Index(string(runes), "*"); i >= 0 {
// // 			runes[i] = ' '
// // 			continue
// // 		}
// // 		if i := strings.Index(string(runes), "/"); i >= 0 {
// // 			runes[i] = ' '
// // 			continue
// // 		}
// // 		if i := strings.Index(string(runes), "("); i >= 0 {
// // 			runes[i] = ' '
// // 			continue
// // 		}
// // 		if i := strings.Index(string(runes), ")"); i >= 0 {
// // 			runes[i] = ' '
// // 			continue
// // 		}
// // 		break
// // 	}

// // 	res := strings.Split(string(runes), " ")
// // }

// func readNum(expression string) []string {

// 	splitFunc := func(r rune) bool {
// 		return strings.ContainsRune("*+-/() ", r)
// 	}

// 	res := strings.FieldsFunc(expression, splitFunc)

// 	return res
// }

// func main() {
// 	arr := "10+2*(3-2)"
// 	//fmt.Println(arr[6])
// 	fmt.Println(Calc(arr))
// }
