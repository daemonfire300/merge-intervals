package merge

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMergeRandomPermutation(t *testing.T) {
	expected := [][]int{{5, 9}, {15, 19}, {20, 24}, {28, 38}, {39, 49}, {54, 69}, {70, 79}, {80, 84}, {90, 112}, {114, 120}, {126, 129}, {131, 138}, {139, 142}, {144, 145}, {151, 153}, {155, 168}, {169, 170}, {173, 178}, {182, 190}, {192, 197}, {201, 212}, {221, 239}, {259, 265}, {267, 272}, {275, 280}, {297, 306}, {307, 328}, {333, 347}, {356, 358}, {362, 367}, {368, 369}, {371, 375}, {376, 383}, {386, 391}, {394, 403}, {409, 421}, {432, 443}, {452, 452}, {458, 464}, {465, 467}, {475, 480}, {481, 488}}
	in := [][]int{{5, 9}, {15, 19}, {20, 24}, {28, 38}, {39, 49}, {54, 69}, {70, 79}, {80, 84}, {90, 112}, {114, 120}, {126, 129}, {131, 138}, {139, 142}, {144, 145}, {151, 153}, {155, 168}, {169, 170}, {173, 178}, {182, 190}, {192, 197}, {201, 212}, {221, 239}, {259, 265}, {267, 272}, {275, 280}, {297, 306}, {307, 328}, {333, 347}, {356, 358}, {362, 367}, {368, 369}, {371, 375}, {376, 383}, {386, 391}, {394, 403}, {409, 421}, {432, 443}, {452, 452}, {458, 464}, {465, 467}, {475, 480}, {481, 488}}

	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprintf("run_%d", i), func(t *testing.T) {
			src := rand.NewSource(time.Now().Unix())
			rnd := rand.New(src)
			sort.Slice(in, func(i, j int) bool {
				return rnd.Int()%2 == 0
			})
			out := Merge(in)
			require.Equal(t, expected, out)
		})
	}
}

func TestMergeFixedData(t *testing.T) {
	expected := [][]int{{5, 9}, {15, 19}, {20, 24}, {28, 38}, {39, 49}, {54, 69}, {70, 79}, {80, 84}, {90, 112}, {114, 120}, {126, 129}, {131, 138}, {139, 142}, {144, 145}, {151, 153}, {155, 168}, {169, 170}, {173, 178}, {182, 190}, {192, 197}, {201, 212}, {221, 239}, {259, 265}, {267, 272}, {275, 280}, {297, 306}, {307, 328}, {333, 347}, {356, 358}, {362, 367}, {368, 369}, {371, 375}, {376, 383}, {386, 391}, {394, 403}, {409, 421}, {432, 443}, {452, 452}, {458, 464}, {465, 467}, {475, 480}, {481, 488}}
	in := [][]int{{20, 24}, {376, 383}, {70, 79}, {114, 120}, {54, 69}, {126, 129}, {297, 306}, {394, 403}, {15, 19}, {475, 480}, {275, 280}, {458, 464}, {182, 190}, {192, 197}, {155, 168}, {169, 170}, {139, 142}, {481, 488}, {371, 375}, {144, 145}, {368, 369}, {386, 391}, {409, 421}, {267, 272}, {5, 9}, {39, 49}, {173, 178}, {356, 358}, {362, 367}, {201, 212}, {333, 347}, {432, 443}, {452, 452}, {221, 239}, {259, 265}, {80, 84}, {151, 153}, {90, 112}, {307, 328}, {465, 467}, {131, 138}, {28, 38}}
	out := Merge(in)
	require.Equal(t, expected, out)
}

func TestMergeSmallData(t *testing.T) {
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	in := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	out := Merge(in)
	require.Equal(t, expected, out)
}

func TestIntersects(t *testing.T) {
	cases := []struct {
		a      []int
		b      []int
		should bool
	}{
		{
			a:      []int{1, 3},
			b:      []int{2, 6},
			should: true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d_%v_%v", i, tc.a, tc.b), func(t *testing.T) {
			intersection := intersects(tc.a, tc.b)
			require.Equal(t, tc.should, intersection)
		})
		t.Run(fmt.Sprintf("%d_%v_%v_reverse", i, tc.a, tc.b), func(t *testing.T) {
			intersection := intersects(tc.b, tc.a)
			require.Equal(t, tc.should, intersection)
		})
	}
}

func TestMergedInterval(t *testing.T) {
	cases := []struct {
		a      []int
		b      []int
		should []int
	}{
		{
			a:      []int{1, 3},
			b:      []int{2, 6},
			should: []int{1, 6},
		},
		{
			a:      []int{1, 3},
			b:      []int{3, 6},
			should: []int{1, 6},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d_%v_%v", i, tc.a, tc.b), func(t *testing.T) {
			res := mergedInterval(tc.a, tc.b)
			require.Equal(t, tc.should, res)
		})
	}
}
