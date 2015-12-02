package main

import (
	"bufio"
	"fmt"
	//      "io"
	//"io/ioutil"
	"os"
	"strings"
	"strconv"
	"sort"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	f, err := os.Open("<GOPATH>/src/github.com/amit-thawait/file_io_ex/sample_data.txt")
	check(err)
	scanner := bufio.NewScanner(f)
	expenseMap := make(map[string]int)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		arr := strings.Split(scanner.Text(), ",")
		amount, err := strconv.Atoi(strings.TrimSpace(arr[1]))
		check(err)
		expenseMap[arr[2]] = expenseMap[arr[2]] + amount
	}
	fmt.Println(expenseMap)
	aggregatedExpense := reverseMap(expenseMap)
	fmt.Println(aggregatedExpense)
	sortedExpenseValue := sortResult(aggregatedExpense)
	fmt.Println(sortedExpenseValue)
	writeToFile(aggregatedExpense, sortedExpenseValue)
}

func reverseMap(m map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

func sortResult(m map[int]string) []int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	fmt.Println(keys)
	return keys
}


func writeToFile(expenseMap map[int]string, sortedExpenseValue []int) {
	fmt.Println(expenseMap)
	fmt.Println(sortedExpenseValue)
	outputFile, err := os.Create("/tmp/output.txt")
	check(err)
	f := bufio.NewWriter(outputFile)

	for i := 0 ; i < len(sortedExpenseValue) ; i++ {
		fmt.Println(expenseMap[sortedExpenseValue[i]])
		expenseRecord := fmt.Sprintf("%s, %v\n", expenseMap[sortedExpenseValue[i]], strconv.Itoa(sortedExpenseValue[i]))
		_, err := f.WriteString(expenseRecord)
		check(err)
	}
	f.Flush()
}

