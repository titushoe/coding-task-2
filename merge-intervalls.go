package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type MergeSlice [][]int

var fileInput []string

func GetIntvSliceFromStdInput(text string, splitchar string) [][]int {
	inpList := make([][]int, 0)
	var iStart, iEnd int
	var s string
	text = strings.Replace(text, " ", "", -1)       // remove spaces
	ivSlices := strings.SplitAfter(text, splitchar) // split into intervall slices
	//	fmt.Printf(" iCount  iSlice\n")
	for i := range ivSlices {
		s = ivSlices[i]
		if len(s) > 0 {
			//	fmt.Printf("  %d   %s \n ", i, s)
			re := regexp.MustCompile(`[-]?\d[\d]*[\.]?[\d{2}]*`)
			//			fmt.Printf("String contains any match: %v\n", re.MatchString(s)) // True
			intvPairs := re.FindAllString(s, -1)
			//for _, IntvDigit := range intvPairs {
			//				fmt.Println(IntvDigit)
			// }
			if i, err := strconv.Atoi(intvPairs[0]); err == nil {
				iStart = i
			}
			if i, err := strconv.Atoi(intvPairs[1]); err == nil {
				iEnd = i
			}
			intvSlice := []int{iStart, iEnd}
			inpList = append(inpList, intvSlice)

		}
	}
	return inpList
}

func (s MergeSlice) Len() int           { return len(s) }
func (s MergeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s MergeSlice) Less(i, j int) bool { return s[i][0] < s[j][0] }

func merge(intervals [][]int) [][]int {
	sort.Sort(MergeSlice(intervals))
	retList := make([][]int, 0)
	size := len(intervals)

	var i int
	for i < size {
		left := intervals[i][0]
		right := intervals[i][1]
		j := i + 1
		for j < size {
			if intervals[j][0] <= right {
				right = Max(intervals[j][1], right)
				j++
			} else {
				break
			}
		}

		ret := []int{left, right}
		retList = append(retList, ret)
		i = j // After merging i=j+1 without merging i=i+1 can be expressed by i=j
	}

	return retList
}

func Max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func main() {
	if len(os.Args) == 1 {
		println("   ", filepath.Base(os.Args[0]), "will read intervalls from stdin")
		reader := bufio.NewReader(os.Stdin)
		iText, _ := reader.ReadString('\n')
		iText = strings.TrimRight(iText, "\r\n")
		intvInput := GetIntvSliceFromStdInput(iText, "]")
		fmt.Print("You entered the intervalls:\n ")
		fmt.Println(intvInput)
		fmt.Print("And they are now merged to:\n ")
		fmt.Println(merge(intvInput))
		os.Exit(0)
	} else {
		if len(os.Args) == 2 {
			println(filepath.Base(os.Args[0]), "is reading the intervalls from .csv-file", os.Args[1])
			if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
				println("  File: ", os.Args[0], "not exists")
				os.Exit(2)
			}
			f_csv, _ := os.Open(os.Args[1])
			// Create a new Scanner for the file.
			scanner := bufio.NewScanner(f_csv)
			// Loop over all lines in the file and print them.
			for scanner.Scan() {
				line := scanner.Text()
				fmt.Println(line)
				line = strings.TrimRight(line, "\r\n")
				line = strings.Replace(line, " ", "", -1) // remove spaces
				fileInput = append(fileInput, line)
			}
			intvLine := strings.Join(fileInput, ",")
			println(intvLine, "\n")
			intvInput := GetIntvSliceFromStdInput(intvLine, ",")
			fmt.Print("You entered the intervalls:\n ")
			fmt.Println(intvInput)
			fmt.Print("And they are now merged to:\n ")
			fmt.Println(merge(intvInput))
			os.Exit(0)
		}
	}
	println("  Usage: ", filepath.Base(os.Args[0]), "[<filename>]")
	os.Exit(1)
}
