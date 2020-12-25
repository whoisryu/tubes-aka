package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func binaryIterative(arrayParams [][]string, searchVariable int) int {
	totalRecord := len(arrayParams) - 1
	start,_ := strconv.Atoi(arrayParams[1][0]) 
	for start <= totalRecord {
		mid := start + (totalRecord-start)/2

		arrMid,_ := strconv.Atoi(arrayParams[mid][0])

		if arrMid == searchVariable {
			return mid
		}
		if arrMid < searchVariable {
			start = mid + 1
		} else {
			totalRecord = mid - 1
		}
	}
	return -1
}

func binarySearchRecursive(numList [][]string, key int) int {

	low := 1
	high := len(numList) - 1

	if low <= high {

		mid := ((high + low) / 2)
		arrMid ,_:= strconv.Atoi(numList[mid][0])

		if arrMid> key {

			return binarySearchRecursive(numList[:mid], key)

		} else if arrMid< key {

			return binarySearchRecursive(numList[mid+1:], key)

		} else {

			return mid
		}
	}

	return -1
}




func main(){
	f, err := excelize.OpenFile("angka_1-10000.xlsx")
    if err != nil {
        fmt.Println(err)
        return
	}
	rows:= f.GetRows("angka")
	
	var cari int
	fmt.Print("masukkan angka yang dicari:")
	fmt.Scanln(&cari)

	hasil := binarySearchRecursive(rows,cari)

	if hasil == -1 {
		fmt.Println("data tidak ditemukan")
	} else{
		fmt.Println("data ditemukan")
	}

	start := time.Now()
	elapsed := time.Since(start)

    fmt.Printf("Binary search took %s\n", elapsed)
}

