package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	totalStudents, studentsPerRoom := 0, 0
	fmt.Fscan(reader, &totalStudents, &studentsPerRoom)

	m := make(map[int]map[int]int)
	for i := 0; i < totalStudents; i++ {
		gender, grade := 0, 0
		fmt.Fscan(reader, &gender, &grade)
		if m[grade] == nil {
			m[grade] = make(map[int]int)
		}
		m[grade][gender]++
	}

	cnt := 0
	for _, grades := range m {
		for _, genders := range grades {
			if genders <= studentsPerRoom {
				cnt++
			} else {
				if genders%studentsPerRoom == 0 {
					cnt += genders / studentsPerRoom
				} else {
					cnt += genders/studentsPerRoom + 1
				}
			}
		}
	}

	writer.WriteString(strconv.Itoa(cnt))
}
