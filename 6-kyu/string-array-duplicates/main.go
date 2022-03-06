package main

func Dup(s []string) []string {
	res := make([]string, len(s))
	for ind, i := range s {
		na := make([]rune, 0, len([]rune(i)))
		for _, c := range i {
			if len(na) == 0 || na[len(na)-1] != c {
				na = append(na, c)
			}
		}
		res[ind] = string(na)
	}
	return res
}
