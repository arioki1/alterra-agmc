package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getMaximumScore' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY arr
 *  2. INTEGER k
 */

func getMaximumScore(arr []int64, k int64) int64 {
	var max int64
	var op = 0
	stop := false
	for !stop {
		var maxInt int64
		var maxIndex int
		for i, v := range arr {
			if v >= maxInt {
				maxInt = v
				maxIndex = i
			}
		}
		if maxInt > 0 {
			op = op + 1
			max = max + (maxInt)
			arr[maxIndex] = int64(math.Round(float64(maxInt) / float64(3)))
		}

		if op == int(k) {
			stop = true
		}

	}
	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var arr []int64

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		arrItem := arrItemTemp
		arr = append(arr, arrItem)
	}

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := kTemp

	result := getMaximumScore(arr, k)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
