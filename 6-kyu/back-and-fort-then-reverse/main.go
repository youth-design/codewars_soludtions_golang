package main

func Arrange(s []int) []int {
	if s == nil {
		return nil
	}
	flag := true
	t := make([]int, 0, len(s))
	for i := 0; i <= len(s)/2-1; i += 1 {
		if flag == true {
			t = append(t, s[i])
			t = append(t, s[len(s)-i-1])
		} else {
			t = append(t, s[len(s)-i-1])
			t = append(t, s[i])

		}
		flag = !flag
	}
	if len(s)%2 == 1 {
		t = append(t, s[len(s)/2])
	}
	return t
}
