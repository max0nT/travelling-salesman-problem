package shuffle

import "math/rand"

type T any

func Shuffle(seq []T) []T {
	// fisher-yates shuffle implementation
	rand.Shuffle(len(seq), func(i, j int) {
		seq[i], seq[j] = seq[j], seq[i]
	})
	return seq
}
