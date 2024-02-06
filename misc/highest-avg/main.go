package highest_avg

func getHighestGradedStudent(grades [][]int) int {
	avgByStudent := make(map[int][]int)

	var student, grade, maxAvg, maxAvgStudent, curAvg, size int

	for _, studentGrade := range grades {
		student, grade = studentGrade[0], studentGrade[1]
		curAvg, size = grade, 1

		data, ok := avgByStudent[student]
		if ok {
			curAvg, size = data[0], data[1]
			size++
			curAvg = (curAvg + grade) / size
		}

		avgByStudent[student] = []int{curAvg, size}
		if curAvg > maxAvg {
			maxAvg = curAvg
			maxAvgStudent = student
		}

		if curAvg == maxAvg {
			maxAvgStudent = getLowestIDStudent(maxAvgStudent, student)
		}
	}

	return maxAvgStudent
}

func getLowestIDStudent(stdA, stdB int) int {
	if stdA < stdB {
		return stdA
	}

	return stdB
}
