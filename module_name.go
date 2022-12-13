package of

var emtpyName struct{}

var preloadModules = map[uint64]struct{}{
	0:                    emtpyName,
	18275347644319189965: emtpyName,
	7852121090758537921:  emtpyName,
	2239877576446804528:  emtpyName,
	14909113260478095105: emtpyName,
	18108201700337443927: emtpyName,
	14580695937169233336: emtpyName,
	16813717775905332386: emtpyName,
	10629397049073578029: emtpyName,
	2061340576534510319:  emtpyName,
	15590474080070586392: emtpyName,
	3635930719169311595:  emtpyName,
	6483280667709159949:  emtpyName,
}

func IsPreloadModule(name Name) bool {
	if _, ok := preloadModules[name.ID()]; ok {
		return true
	}
	return false
}
