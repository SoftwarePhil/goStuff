package main

import "fmt"
import "math"
import "strconv"

func main(){
		vec := Vector{3,4}
		vec3D := Vector3D{vec,5}
		vec3D2 := Vector3D{Vector{1,1},1}
		fmt.Println("The unit vector of", vec3D.print(), "is", vec3D.unitVector().print())
		fmt.Println("The dot product of", vec3D.print(), "and", vec3D2.print(), "is", vec3D.dotProduct(vec3D2))
		fmt.Println("Vector", vec3D.print(), "times scalar 4 is", vec3D.scalarMult(4).print())
		fmt.Println("Vector", vec3D.print(), "cross", vec3D2.print(), "equals", vec3D.crossProduct(vec3D2).print())
		
		a := Vector3D{Vector{0,3},1}
		b := Vector3D{Vector{-3,7},8}
		fmt.Println("Scalar Mult", a.scalarMult(a.dotProduct(b)).print())
		fmt.Println("cross product a cross a", a.crossProduct(a).print())
		
		c := Vector3D{Vector{1,1},1}
		d := Vector3D{Vector{1,-1},-1}
		fmt.Println("Angle between", c.print(), "and", d.print(), "is", c.angle(d))
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

func (vec Vector) print() string{
		x := strconv.FormatFloat(vec.x,'f',2,64)
		y := strconv.FormatFloat(vec.y,'f',2,64)
		
		return "(" + x + "," + y +")"
}
func (vec Vector3D) print() string{
		x := strconv.FormatFloat(vec.x,'f',2,64)
		y := strconv.FormatFloat(vec.y,'f',2,64)
		z := strconv.FormatFloat(vec.z,'f',2,64)
		
		return "(" + x + "," + y + "," + z + ")"
}

func(vec Vector) unitVector() Vector{
		l := vec.length()
	
		return Vector{(vec.x/l),(vec.y/l)}
}

func(vec Vector3D) unitVector() Vector3D{
		l := vec.length()
	
		return Vector3D{Vector{(vec.x/l), (vec.y/l)}, (vec.z/l)}
}

func(a Vector) dotProduct(b Vector) float64{
		return (a.x * b.x) + (a.y * b.y)
}	

func(a Vector3D) dotProduct(b Vector3D) float64{
		return (a.x * b.x) + (a.y * b.y) + (a.z * b.z)
}

func(vec Vector) scalarMult(x float64) Vector{
		return Vector{(vec.x*x),(vec.y*x)}
}

func(vec Vector3D) scalarMult(x float64) Vector3D{ 
		return Vector3D{Vector{(vec.x*x),(vec.y*x)},(vec.z*x)}
}

func(a Vector) crossProduct(b Vector) float64{		
		x := (a.x*b.y)-(a.y*b.x)	
		
		return x
}	

func(a Vector3D) crossProduct(b Vector3D) Vector3D{
		x := (a.y*b.z)-(a.z*b.y)
		y := (a.z*b.x)-(a.x*b.z)
		z := (a.x*b.y)-(a.y*b.x)
	
		return Vector3D{Vector{x,y},z}
}

func(a Vector3D) angle(b Vector3D) float64{
		angle := (a.dotProduct(b))/((a.length())*(b.length()))
		
		return math.Acos(angle)*180/math.Pi
}

/*
 * how pointers work, this will change the value of any vector who calls 
 * into a zero vector
func(v *Vector) test(){
		*v = Vector{0,0}
}	
*/
