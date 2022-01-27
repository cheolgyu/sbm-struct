package model

import "math"

const P3_type_LOW = "L"
const P3_type_HIGH = "H"

/*
P1 : 움직이는 과거의 점
P2 : 마지막일자의 점
P3 : 찾은점

P1x_Unit_type : 이동한 기간 단위 1:주, 2:월
P1x_Unit      : 이동한 기간 값
P3_type      : 찾은점 타입, 최고=h,최저점=l
P32y_percent : 찾은점과 현재점의 y 퍼센트

p3_type:
*/
type Tb52Weeks struct {
	Code_id    int
	Price_type int

	P1x_Unit_type int
	P1x_Unit      int

	P1 P
	P2 P
	P3 P

	P3_type      string
	P32y_percent float32
}

/*
X : 일자
Y : 가격
*/
type P struct {
	X int
	Y float32
}

/*
p1 왼쪽 p2 우측 점
*/
func Get_percent(p1_y float32, p2y float32) float32 {
	minus := p2y - p1_y
	chk := float32(float64(minus / p2y * 100))

	// -inf error
	if math.IsInf(float64(chk), 0) {
		return 0
	}
	return chk
}
