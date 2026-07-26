package internal

func (b *Base) Shuffle(n int, swap func(i, j int)) {
	b.rd.Shuffle(n, swap)
}
