package set_matrix_zeroes

func setZeroes(matrix [][]int) {
	zeroes := map[string]map[int]struct{}{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				//row
				if val, ok := zeroes["row"]; ok {
					val[i] = struct{}{}
				} else {
					zeroes["row"] = map[int]struct{}{
						i: {},
					}
				}

				//column
				if val, ok := zeroes["column"]; ok {
					val[j] = struct{}{}
				} else {
					zeroes["column"] = map[int]struct{}{
						j: {},
					}
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if _, ok := zeroes["row"][i]; ok {
				matrix[i][j] = 0
			}

			if _, ok := zeroes["column"][j]; ok {
				matrix[i][j] = 0
			}
		}
	}
}
