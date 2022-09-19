package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'getMaximumEvenSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts INTEGER_ARRAY val as parameter.
 */

func getMaximumEvenSum(val []int32) int64 {
	//maximum sum val
	var maxSum int64
	var totalVal = len(val)
	for i, _ := range val {
		var sum int64
		for j := i; j < totalVal; j++ {
			sum = sum + int64(val[j])
			if sum%2 == 0 {
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}

	return maxSum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	valCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var val []int32

	for i := 0; i < int(valCount); i++ {
		valItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		valItem := int32(valItemTemp)
		val = append(val, valItem)
	}

	result := getMaximumEvenSum(val)

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
