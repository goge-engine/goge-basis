package gogeBasis

type Warning interface {
	// A struct of Error
	Warning() string
}

// 实现一个具体的警告类型
type warning struct {
	message string
}

func (w *warning) Warning() string {
	return w.message
}

func NewWarning(warningMessage string) Warning {
	return &warning{warningMessage}
}

func ReturnWarningForTest() Warning {
	warn := NewWarning("This is a warning message")
	return warn
}
