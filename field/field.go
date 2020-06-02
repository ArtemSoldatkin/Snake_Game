package field

// Field - game field
type Field struct {
	Width, Height int
}

// SetSize - set new size to field
func (f *Field) SetSize(width, heigth int) {
	f.Width = width
	f.Height = heigth
}
