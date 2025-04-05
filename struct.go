package gogeBasis

type Warinng interface {
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

func NewWarning(warningMessage string) Warinng {
	return &warning{warningMessage}
}

func ReturnWarningForTest() Warinng {
	warn := NewWarning("This is a warning message")
	return warn
}
