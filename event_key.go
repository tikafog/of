package of

type EventKey uint64

func (key EventKey) Key() uint64 {
	return uint64(key)
}
