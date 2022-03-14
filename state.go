package of

//State the state of the device
//ENUM(init,preload,loading,waiting,running,syncing,stopping,stopped,error,max)
type State int

func (x State) Is(other State) bool {
	return x == other
}
