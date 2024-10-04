package rectangle // проект main на весь проект один

type Rectangle struct {
	A, B int // Эти поля экспортируемы
	color string // Это поле - нет
}

func New(newA int, newB int, newColor string) *Rectangle{
	return &Rectangle{newA, newB, newColor}
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.A + r.B)
}