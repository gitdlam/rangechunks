// rangechunks
package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {

	limit := 3000

	snip1 := "ph.shipto between '"
	snip2 := "' and '"
	snip3 := "'"
	snip4 := "ph.shipto = '"

	var text string
	entries := strings.Split(data, "\n")
	var numbers []int
	for _, v := range entries {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("cannot convert " + v)
		}
		numbers = append(numbers, n)

	}

	sort.Ints(numbers)
	N := len(numbers)
	paired := 1
	for i, v := range numbers {
		if i == 0 {
			text = snip1 + strconv.Itoa(v) + snip2
			paired = 1
			continue
		}

		if paired == 2 && numbers[i+1]-v <= limit {
			text = text + " or " + snip1 + strconv.Itoa(v) + snip2
			paired = 1
			continue
		}

		if paired == 2 && numbers[i+1]-v > limit {
			text = text + " or " + snip4 + strconv.Itoa(v) + snip3
			paired = 2
			continue
		}

		if i+1 == N && paired == 1 {
			text = text + strconv.Itoa(v) + snip3
			paired = 2
			continue
		}

		if numbers[i+1]-v > limit && paired == 1 {
			text = text + strconv.Itoa(v) + snip3
			paired = 2
			continue

		}
	}

	fmt.Println(text)

}

var data = `10105147
10106497
10101311
10101312
10101313
10107375
10105813
10101083
10101124
10101134
10101135
10101136
10101137
10101138
10101139
10101140
10101142
10101143
10600997
10601303
10601711`
