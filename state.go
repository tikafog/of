package of

//State the state of the device
//Event can only handler from loading to max
//ENUM(invalid,init,preload,loading,waiting,running,syncing,syncing_online,syncing_offline,stopping,stopped,error,max)
type State uint32

func (x State) Is(other State) bool {
	return x == other
}
