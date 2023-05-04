package AreTheyTheSame

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil {
		return false
	}
	count := map[int][2]int{}
	for _, v := range array2 {
		if v < 0 {
			return false
		}
		if v2, ok := count[v]; ok {
			count[v] = [2]int{v2[0], v2[1] + 1}
		} else {
			count[v] = [2]int{0, 1}
		}
	}
	for _, v := range array1 {
		v = v * v
		if v1, ok := count[v]; ok {
			count[v] = [2]int{v1[0] + 1, v1[1]}
		} else {
			count[v] = [2]int{1, 0}
		}
	}
	for _, v := range count {
		if v[0] != v[1] {
			return false
		}
	}
	return true
}

// type value struct {
// 	base        int
// 	square      int
// 	countBase   int
// 	countSquare int
// }

// func Comp(array1 []int, array2 []int) bool {
// 	sort.Ints(array1)
// 	sort.Ints(array2)
// 	fmt.Printf("%v\n%v\n", array1, array2)
// 	if array1 == nil || array2 == nil {
// 		fmt.Println(false)
// 		return false
// 	}
// 	a := []*value{}
// 	for _, v1 := range array1 {
// 		i := inArray(a, v1, true)
// 		if i != -1 {
// 			a[i].countBase++
// 			continue
// 		}
// 		a = append(a, &value{base: v1, square: v1 * v1, countBase: 1})
// 	}

// 	for _, v2 := range array2 {
// 		if v2 < 0 {
// 			fmt.Println(false)
// 			return false
// 		}
// 		i := inArray(a, v2, false)
// 		if i != -1 {
// 			a[i].countSquare++
// 			continue
// 		}
// 		a = append(a, &value{base: int(math.Sqrt(float64(v2))), square: v2, countSquare: 1})
// 	}
// 	for _, v := range a {
// 		fmt.Printf("\t%v\t|\t%v\t|\t%v\t|\t%v\n", v.base, v.square, v.countBase, v.countSquare)
// 		if v.countBase != v.countSquare {
// 			fmt.Println(false)
// 			return false
// 		}
// 	}
// 	fmt.Println(true)
// 	return true
// }

// func inArray(a []*value, t int, c bool) int {
// 	for i, v := range a {
// 		if c {
// 			if v.base == t {
// 				return i
// 			}
// 		} else {
// 			if v.square == t {
// 				return i
// 			}
// 		}
// 	}
// 	return -1
// }
