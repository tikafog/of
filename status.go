package of

//State the state of the device
//ENUM(init,preload,loading,waiting,running,error)
type State int

func (s State) Is(other State) bool {
	return s == other
}
