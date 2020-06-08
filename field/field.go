package field

// Field - game field
type Field struct {
	FieldWidth, FieldHeight int
	Head                    *Block
}

// SetSize - set new size to field
func (f *Field) SetSize(width, heigth int) {
	f.FieldWidth = width
	f.FieldHeight = heigth
}

// Init - initialize field
func (f *Field) Init() {
	const blockWidth, blockHeight int = 20, 20
	f.Head = &Block{Width: blockWidth, Height: blockHeight}
	f.Head.InitRand(f.FieldHeight-blockHeight, f.FieldWidth-blockWidth)
}
