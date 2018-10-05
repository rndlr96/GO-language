package Rectangle

type Rectangle struct {
	width  int
	height int
}

func rectangleScaleA(rect *Rectangle, factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func rectangleScaleB(rect Rectangle, factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (rect *Rectangle) ScaleA(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}

func (rect Rectangle) ScaleB(factor int) {
	rect.width = rect.width * factor
	rect.height = rect.height * factor
}
