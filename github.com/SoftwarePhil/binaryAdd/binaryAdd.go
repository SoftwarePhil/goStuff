package main

import "fmt"
import "math"
import "strconv"

func main(){
	rec := aRectangle{67.5,100,170,87}
	p1,p2,p3,p4 := rec.generateFourPoints()
	fmt.Println(p1.getStringPoint(),p2.getStringPoint(),p3.getStringPoint(),p4.getStringPoint())
	fmt.Println("The distance between point", p1.getStringPoint(), "and point", p2.getStringPoint(), "is", getDistance(p1,p2))
	//a := Point{5,5}
	//b := Point{10,14}
	fmt.Println(findSlope(p1,p4))
}

type aRectangle struct{
	
	rightTopX float64
	rightTopY float64
	length float64
	width float64

}

type Point struct{
	
	x float64
	y float64
}

func(p Point) getStringPoint()string{
	x := strconv.FormatFloat(p.x, 'f', 2, 64)
	y := strconv.FormatFloat(p.y, 'f', 2, 64)
	cor := "("+ x+ ","+ y+ ")"
	return cor
}

func(rect aRectangle) generateFourPoints()(Point, Point, Point, Point){
	x2 := rect.rightTopX -rect.length
	x1 := rect.rightTopX
	y1 := rect.rightTopY
	y2 := rect.rightTopY -rect.width
	
	point1 := Point{x1,y1}
	point2 := Point{x1,y2}
	point3 := Point{x2,y1}
	point4 := Point{x2,y2}

	return point1, point2, point3, point4
}

func getDistance(p1, p2 Point) float64{
	distance :=math.Sqrt(math.Pow((p2.y-p1.y),2) + math.Pow((p2.x-p1.x),2))
	return distance
}

func findSlope(p1,p2 Point) float64{
	return (p2.y-p1.y)/(p2.x-p1.x)
} 
