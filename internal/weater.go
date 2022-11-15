package internal

type Weater struct {
	Wind        int
	Temperature int
	Snow        bool
	Rain        bool
	//Drive       func() (bool, error)
}

func (w *Weater) CheckDrive() (bool, error) {
	if w.Snow {
		return false, nil
	}

	return true, nil
}

func NewWeater(wind int, temperature int, snow bool, rain bool) *Weater {
	return &Weater{
		Wind:        wind,
		Temperature: temperature,
		Snow:        snow,
		Rain:        rain,
	}
}
