package utils

func Diff(arr1, arr2 []string) []string {

	result := make([]string, 0, len(arr1))

	for _, v := range arr1 {

		exist := false

		for _, w := range arr2 {

			if v == w {
				exist = true

				break
			}
		}

		if !exist {
			result = append(result, v)
		}
	}

	return result
}
