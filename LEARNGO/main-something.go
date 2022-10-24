package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Solution() {
	// 파일 오픈
	file, _ := os.Open("./sample01.csv")
	file2, _ := os.Open("./sample02.csv")

	// csv reader 생성
	rdr := csv.NewReader(bufio.NewReader(file))
	rdr2 := csv.NewReader(bufio.NewReader(file2))

	// csv 내용 모두 읽기
	rows, _ := rdr.ReadAll()
	rows2, _ := rdr2.ReadAll()
	var m = make(map[string][]string)

	// 행,열 읽기
	for i, row := range rows {
		for j := range row {
			if j == 0 {
				continue
			} else {
				m[rows[i][0]] = append(m[rows[i][0]], rows[i][j])
			}
		}
	}
	for i, row2 := range rows2 {
		for j := range row2 {
			if j == 0 {
				continue
			} else {
				m[rows2[i][0]] = append(m[rows2[i][0]], rows2[i][j])
			}
		}
	}
	file3, err := os.Create("./result.csv")
	if err != nil {
		panic(err)
	}

	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file3))

	//key값으로 정렬
	sortKeys := make([]string, 0, len(m))
	for k := range m {
		sortKeys = append(sortKeys, k)
	}
	sort.Strings(sortKeys) //keys로 정렬을 함

	//정렬한 keys 값으로 데이터를 출력함
	for _, k := range sortKeys {
		// csv 내용 쓰기
		wr.Write([]string{k, m[k][0], m[k][1], m[k][2], m[k][3]})
	}
	// csv 내용 쓰기
	/*	for s := range m {
		wr.Write([]string{s, m[s][0], m[s][1], m[s][2], m[s][3]})
	}*/
	wr.Flush()
}
func main() {
	/*	name := "Mark"
		name = "Lynn"
		fmt.Println(name)
		fmt.Println(multiply(2, 2))*/
	/*	totalLength, upperName := lenAndUpper("James")
		fmt.Println(totalLength, upperName)
		repeatMe("A", "B", "C", "D", "E")
		total := superAdd(1, 2, 3, 4, 5, 6, 7)
		fmt.Println(total)
		fmt.Println(canIDrink(15))*/
	//pointer
	/*	a := 2
		b := &a
		a = 5
		*b += 20
		fmt.Println(a)*/

	//array
	/*	names := []string{"nico", "lynn", "dal"}
		names = append(names, "fly")
		fmt.Println(names)*/

	//Maps
	/*	nico := map[string]string{"name": "nico", "age": "18"}
		for key, value := range nico {
			fmt.Println(key, value)
		}*/
	Solution()

}

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, upperCase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	upperCase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func canIDrink(age int) bool {
	// 변수룰 따로 생성할수도 있지만 이렇게 생성하면
	// 아 이 변수는 조건물을 위해 사용했구나라고 파악하기 용이.
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}
