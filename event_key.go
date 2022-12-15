package of

type EventKey uint64

type EventKeyAble interface {
	Key(str string) EventKey
	KeyName(EventKey) string
}

func (key EventKey) Key() uint64 {
	return uint64(key)
}
