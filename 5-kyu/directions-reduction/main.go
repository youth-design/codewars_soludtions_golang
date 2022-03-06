package main

func DirReduc(arr []string) []string {
	res := make([]string, 0, len(arr))
	mapping := map[string]string{"NORTH": "SOUTH", "SOUTH": "NORTH", "EAST": "WEST", "WEST": "EAST"}

	for i := 0; i < len(arr); i++ {
		if arr[i] == "" {
			continue
		}
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == "" {
				continue
			}

			if arr[j] == mapping[arr[i]] {
				arr[j] = ""
				arr[i] = ""
				i = 0
				break
			} else if arr[j] != arr[i] {
				break
			}
		}
	}

	for _, i := range arr {
		if i != "" {
			res = append(res, i)
		}
	}

	return res
}
