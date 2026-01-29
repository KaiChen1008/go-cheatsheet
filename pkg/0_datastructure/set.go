package datastructure

func PairSet() {
	a := pair{1, 2}
	b := pair{1, 2}

	vis := make(map[pair]bool)

	vis[a] = true
	println(vis[b]) // true
}
