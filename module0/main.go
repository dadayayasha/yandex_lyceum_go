package main

import (
	"fmt"
)

func main() {
	str := []string{"asd", "123", "das"}
	fmt.Println(str)
	tmp(str)
	fmt.Println(str)

}

func tmp(str []string) {
	str[0], str[1] = str[1], str[0]
}

// func  ParseStringToTime(dateString, format string) (time.Time, error){
// 	dateTime,err:=time.Parse(format,dateString)
// 	if err!=nil{
// 		return dateTime,err
// 	}
// 	return dateTime,nil
// }

// func FormatTimeToString(timestamp time.Time, format string) string {
// 	formattime := timestamp.Format(format)
// 	return fmt.Sprint(formattime)
// }

// func AreAnagrams(str1, str2 string) bool{
// 	if len(str1)!=len(str2){
// 		return false
// 	}
// 	str1=strings.ToLower(str1)
// 	str2=strings.ToLower(str2)
// 	runes1:=[]rune(str1)
// 	runes2:=[]rune(str2)
// 	sort.Slice(runes1,func(i, j int) bool {
// 		return runes1[i] < runes1[j]
// 	})
// 	sort.Slice(runes2,func(i, j int) bool {
// 		return runes2[i] < runes2[j]
// 	})
// 	if string(runes1)==string(runes2){
// 		return true
// 	}
// 	return false
// }

// func SumTwoIntegers(a, b string) (int, error){
// 	num1,err1:=strconv.Atoi(a)
// 	num2,err2:=strconv.Atoi(b)

// 	if err1!=nil||err2!=nil{
// 		return 0,fmt.Errorf("invalid input, please provide two integers")
// 	}
// 	return num1+num2,nil
// }
// func IntToBinary(num int) (string, error){
// 	if num<0{
// 		return "",fmt.Errorf("negative numbers are not allowed")
// 	}
// 	return fmt.Sprintf("%b",num),nil
// }

// func Factorial(n int) (int, error){
// 	if n<0{
// 		return 0,fmt.Errorf("factorial is not defined for negative numbers")
// 	}
// 	res:=1
// 	for ;n>0;n--{
// 		res*=n
// 	}
// 	return res,nil
// }

// func GetCharacterAtPosition(str string, position int) (rune, error){
// 	runes:= []rune(str)
// 	if position>=len(runes){
// 		return 0,fmt.Errorf("position out of range")
// 	}
// 	return runes[position],nil
// }
// type Logger interface {
// 	Log()
// }

// type LogLevel string

// const (
// 	Error LogLevel = "ERROR"
// 	Info  LogLevel = "INFO"
// )

// type Log struct {
// 	Level LogLevel
// }

// func (l Log) Log(log string) {

// 	fmt.Printf("%s: %s", l.Level, log)
// }

// type Cat struct{

// }
// type Dog struct{

// }
// type Animal interface{
// 	MakeSound()
// }

// func (c Cat) MakeSound() {
// 	fmt.Println("Мяу!")
// }
// func (c Dog) MakeSound() {
// 	fmt.Println("Гав!")
// }

// type Note struct {
// 	title string
// 	text  string
// }

// type ToDoList struct {
// 	name  string
// 	tasks  []Task
// 	notes []Note
// }

// type Task struct {
// 	summary     string
// 	description string
// 	deadline    time.Time
// 	priority    int
// }

// func (t Task) IsOverdue() bool {

// 	if time.Now().After(t.deadline) {
// 		return true
// 	}

// 	return false
// }

// func (t Task) IsTopPriority() bool {
// 	if t.priority == 5 || t.priority == 4 {
// 		return true
// 	}
// 	return false
// }

// func (t ToDoList) TasksCount() int {
// 	return len(t.tasks)
// }

// func (t ToDoList) NotesCount() int {
// 	return len(t.notes)
// }

// func (t ToDoList) CountTopPrioritiesTasks() int {
// 	var res int
// 	for _, tk := range t.tasks {
// 		if tk.IsTopPriority() {
// 			res++
// 		}
// 	}
// 	return res
// }

// func (t ToDoList) CountOverdueTasks() int {
// 	var res int
// 	for _, tk := range t.tasks {
// 		if tk.IsOverdue() {
// 			res++
// 		}
// 	}
// 	return res
// }

// func main() {
// 	todo := ToDoList{name: "Gosha ToDo list", tasks: []Task{Task{summary: "Make Yandex Lyceum Task 9", deadline: time.Now().Add(-time.Hour), description: "Make Module 0, Task 9", priority: 5}, Task{summary: "Make Yandex Lyceum Task 10", deadline: time.Now().Add(time.Hour), description: "Make Module 0, Task 10", priority: 3}}, notes: []Note{Note{title: "ToDo list task", text: "ToDo list task in Yandex Lyceum is very interesting"}}}
// 	fmt.Println(todo.TasksCount())
// 	fmt.Println(todo.NotesCount())
// 	fmt.Println(todo.CountTopPrioritiesTasks())
// 	fmt.Print(todo.CountOverdueTasks())
// }
// func DeleteLongKeys(m map[string]int) map[string]int{
// 	res:= make(map[string]int)

// 	for i:=range m{
// 		if len(i)>=6{
// 			res[i]=m[i]
// 		}
// 	}
// 	return res
// }

// func CountingSort(contacts []string) map[string]int{
// 	counter:= make(map[string]int)
// 	for _,i:=range contacts{
// 		counter[i]++
// 	}
// 	return counter
// }

// func SwapKeysAndValues(m map[string]string) map[string]string{
// 	res:= make(map[string]string)
// 	for i:=range m{
// 		res[m[i]]=i
// 	}
// 	return res
// }

// func SumOfValuesInMap(m map[int]int) int{
// 	var sum int
// 	for _,i:= range m{
// 		sum+=i
// 	}
// 	return sum
// }

// func FindMaxKey(m map[int]int) int{
// 	var max int
// 	for i:= range m{
// 	max = i
// 	break
// 	}
// 	for i:= range m {
// 		if max<i{
// 			max=i
// 		}
// 	}
// 	return max
// }

// func Permutations(input string) []string{
// 	str:=[]rune(input)
// 	var res []string
// 	n:=len(input)
// 	for i:=0;i<n;i++{
// 		for j:=0;j<n-1;j++{
// 			str[j],str[j+1]=str[j+1],str[j]
// 			res = append(res,string(str))
// 		}
// 	}
// 	sort.Strings(res)
// 	return res
// }
// func IsPalindrome(input string) bool{
// 	input = strings.ToLower(input)
// 	input = strings.Replace(input," ","",-1)
// 	str:=[]rune(input)
// 	n:=len(str)
// 	for i:=0;i<n/2;i++{
// 		if str[i]!=str[n-i-1]{
// 			return false
// 		}
// 	}
// 	return true
// }

// func isLatin(input string) bool {
// 	str := []rune(input)
// 	for _, r := range str {
// 		if !unicode.Is(unicode.Latin,r){
// 			return false
// 		}
// 	}
// 	return true
// }

// func ConcatenateStrings(str1, str2 string) string{
// 	return str1+" "+str2
// }

// func StringLength(input string) int{
// 	var str []rune
// 	str = []rune(input)
// 	var res int
// 	for range str{
// 		res++
// 	}
// 	return res
// }

// func  Mix(nums []int) []int{
// 	n:=len(nums)/2
// 	res:= make([]int,0,len(nums))
// 	for i:=0;i<len(nums)/2;i++{
// 		res = append(res,nums[i],nums[n+i])

// 	}
// 	return res
// }

// func Join(nums1, nums2 []int) []int{
// 	res:= make([]int,len(nums1)+len(nums2))
// 	copy(res,nums1)
// 	copy(res[:len(nums1)],nums2)
// 	return res
// }

// func SliceCopy(nums []int) []int{
// 	var res []int
// 	return append(res,nums...)
// }

// func Clean(nums []int, x int) ([]int){
// 	for i:=0;i<len(nums);i++ {
// 		if nums[i]==x {
// 			nums = append(nums[:i],nums[i+1:]...)
// 			i--
// 		}
// 	}
// 	return nums
// }

// func UnderLimit(nums []int, limit int, n int) ([]int, error){
// 	if len(nums)==0{
// 		return nil,fmt.Errorf("slise is nil")
// 	}
// 	if n<=0{
// 		return nil,fmt.Errorf("n < 0")
// 	}
// 	var res []int
// 	for _,num:= range nums{
// 		if num<limit{
// 			res = append(res,num)
// 		}
// 	}
// 	if len(res)>n{
// 		res=res[:n]
// 	}
// 	return res,nil

// }
