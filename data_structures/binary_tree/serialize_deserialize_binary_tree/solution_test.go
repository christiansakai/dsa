package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var t1 *node = nil

	t2 := &node{
		val: 1,
		left: &node{
			val:   2,
			left:  nil,
			right: nil,
		},
		right: &node{
			val: 3,
			left: &node{
				val:   4,
				left:  nil,
				right: nil,
			},
			right: &node{
				val:   5,
				left:  nil,
				right: nil,
			},
		},
	}

	// t3 := &node{
	//   val: 5,
	//   left: &node{
	//     val: 2,
	//     left: nil,
	//     right: nil,
	//   },
	//   right: &node{
	//     val: 3,
	//     left: &node{
	//       val: 2,
	//       left: &node{
	//         val: 3,
	//         left: nil,
	//         right: nil,
	//       },
	//       right: &node{
	//         val: 1,
	//         left: nil,
	//         right: nil,
	//       },
	//     },
	//     right: &node{
	//       val: 4,
	//       left: nil,
	//       right: nil,
	//     },
	//   },
	// }

	tests := []struct {
		Data *node
	}{
		{t1},
		{t2},
		// {t3},
	}

	for _, tt := range tests {
		serialized := SolveSerialize(tt.Data)
		got := SolveDeserialize(serialized)

		if !reflect.DeepEqual(got, tt.Data) {
			t.Errorf("got %v, want %v", got, tt.Data)
		}
	}
}
