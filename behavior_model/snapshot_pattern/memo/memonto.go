package memo

type Memento struct {
	state string
}

func (m *Memento) getSavedState() string {
	return m.state
}
