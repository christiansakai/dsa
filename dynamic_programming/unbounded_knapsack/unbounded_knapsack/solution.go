package solution

import "math"

type Fruit struct {
	Name   string
	Weight int
	Profit int
}

func TopDown(fruits []Fruit, capacity int) int {
	// if len(fruits) == 0 {
	//   return 0
	// }

	// cache := map[int]int{}

	// return recurse(fruits, capacity, index, cache)
}

func recurse(fruits []Fruit, capacity, index int, cache map[int]int) int {
	// if capacity == 0 {
	//   return 0
	// }

	// if result, ok := cache[capacity]; ok {
	//   return result
	// }

	// var max float64 = 0
	// for _, f := range fruits {
	//   if f.Weight <= capacity {
	//     selection := f.Profit + recurse(fruits, capacity - f.Weight, cache)
	//     max = math.Max(max, float64(selection))
	//   }
	// }

	// result := int(max)

	// cache[capacity] = result

	// return result
}
