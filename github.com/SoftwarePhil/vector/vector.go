package main

import "fmt"
import "math"
import "strconv"

func main(){
		vec := Vector{3,4}
		vec3D := Vector3D{vec,5}
		//vec3D2 := Vector3D{Vector{1,1},1}
		fmt.Println("The unit vector of", vec3D.print(), "is", vec3D.unitVector().print())
}

type Vector struct{
		x float64
		y float64
}

type Vector3D struct{
		Vector
		z float64
}

type Point struct{
		x float64
		y float64
}

type Point3D struct{
		Point
		z float64
}

func(vec Vector) length() float64{
		return math.Sqrt(math.Pow(vec.x,2) + math.Pow(vec.y,2))
} 

func(vec Vector3D) length() float64{
		return math.Sqrt(math.Pow(vec.x,2) + math.Pow(vec.y,2)+math.Pow(vec.z,2))
}

func (vec Vector3D) print() string{
		x := strconv.FormatFloat(vec.x,'f',2,64)
		y := strconv.FormatFloat(vec.y,'f',2,64)
		z := strconv.FormatFloat(vec.z,'f',2,64)
		
		return "(" + x + "," + y + "," + z + ")"
}

func (vec Vector) print() string{
		x := strconv.FormatFloat(vec.x,'f',2,64)
		y := strconv.FormatFloat(vec.y,'f',2,64)
		
		return "(" + x + "," + y +")"
}

func(vec *Vector3D) unitVector() Vector3D{
	l := vec.length()
	return Vector3D{Vector{(vec.x/l), (vec.y/l)}, (vec.z/l)}
}

func dotProduct(a,b Vector) float64{
		return (a.x * b.x) + (a.y * b.y)
}	

func dotProduct3D(a,b Vector3D) float64{
		return (a.x * b.x) + (a.y * b.y) + (a.z * b.z)
}	
