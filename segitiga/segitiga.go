package segitiga

import "errors"

type TriangleCalculator struct {
	//ada 5 objek
	
	Base   float64 //ini itu dasar segitiga atau alas
	Height float64 // ini tingginya segitiga
	Side1  float64 // sisi 1
	Side2  float64 // sisi 2
	Side3  float64 // sisi 3
}

func InitSegitiga(base, height, side1, side2, side3 float64) *TriangleCalculator {
	//inibuat isi objeknya yang ada di triangle kalkulator tadi
	return &TriangleCalculator{
		Base:   base,
		Height: height,
		Side1:  side1,
		Side2:  side2,
		Side3:  side3,
	//biar bisa di kembalikan dan membuat objek
	}
}

func (t *TriangleCalculator) Area() float64 { //gitung luas
	return 0.5 * t.Base * t.Height
}

func (t *TriangleCalculator) Perimeter() float64 { //hitung kelilingg
	return t.Side1 + t.Side2 + t.Side3
}


//untuk mengatur alas, tinggi, sissis

func (t *TriangleCalculator) SetBase(base float64) { 
	t.Base = base
}

func (t *TriangleCalculator) SetHeight(height float64) {
	t.Height = height
}

func (t *TriangleCalculator) SetSides(side1, side2, side3 float64) {
	t.Side1 = side1
	t.Side2 = side2
	t.Side3 = side3
}

//validasi biar semua angka lebih besar dari nol
func (t *TriangleCalculator) ValidateTriangle() error {
	if t.Base <= 0 || t.Height <= 0 || t.Side1 <= 0 || t.Side2 <= 0 || t.Side3 <= 0 {
		return errors.New("all sides and height must be greater than zero")
	}
	return nil
}
