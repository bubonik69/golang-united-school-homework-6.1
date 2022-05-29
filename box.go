package golang_united_school_homework

import "fmt"

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) < b.shapesCapacity  {
		b.shapes = append(b.shapes, shape)
		return nil
	}
	//panic("implement me")
	return fmt.Errorf("added more than max elements")
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= 0 && i < len(b.shapes) {
		return b.shapes[i], nil
	}
	//panic("implement me")
	return nil, fmt.Errorf("Shape with index %d not exist ", i)
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= 0 && i < len(b.shapes) {
		removedShape := b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return removedShape, nil
	}
	//panic("implement me")
	return nil, fmt.Errorf("Shape with index %d not exist ", i)
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= 0 && i < len(b.shapes) {
		removedShape := b.shapes[i]
		b.shapes[i] = shape
		return removedShape, nil
	}
	//panic("implement me")
	return nil, fmt.Errorf("Shape with index %d not exist ", i)
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for i := 0; i < len(b.shapes); i++ {
		sum += b.shapes[i].CalcPerimeter()
	}
	return sum
	//panic("implement me")

}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for i := 0; i < len(b.shapes); i++ {
		sum += b.shapes[i].CalcArea()
	}
	return sum
	//panic("implement me")

}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	var sumCircle int =0
	var NewShape []Shape
	for i := 0; i < len(b.shapes); i++ {
		switch b.shapes[i].(type){
		case *Circle:
			sumCircle += 1
			continue
		}
		NewShape= append(NewShape,b.shapes[i])
	}
	if sumCircle == 0 {
		return fmt.Errorf("Have no one circle!")
	}
	b.shapes=NewShape
	return nil


}
