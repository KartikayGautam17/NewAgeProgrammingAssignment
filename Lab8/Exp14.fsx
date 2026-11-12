open System

type Shape(name: string) =
    member this.Name = name
    abstract Area: unit -> float
    default this.Area() = 0.0
    abstract Perimeter: unit -> float
    default this.Perimeter() = 0.0
    member this.Describe() =
        sprintf "%-10s area = %8.2f   perimeter = %8.2f" this.Name (this.Area()) (this.Perimeter())

type Circle(radius: float) =
    inherit Shape("Circle")
    member this.Radius = radius
    override this.Area() = Math.PI * radius * radius
    override this.Perimeter() = 2.0 * Math.PI * radius

type Rectangle(width: float, height: float) =
    inherit Shape("Rectangle")
    member this.Width = width
    member this.Height = height
    override this.Area() = width * height
    override this.Perimeter() = 2.0 * (width + height)

type Triangle(a: float, b: float, c: float) =
    inherit Shape("Triangle")
    override this.Area() =
        let s = (a + b + c) / 2.0
        sqrt (s * (s - a) * (s - b) * (s - c))
    override this.Perimeter() = a + b + c

printfn "=== Calling methods on individual objects ==="
let circle = Circle(5.0)
let rectangle = Rectangle(4.0, 6.0)
let triangle = Triangle(3.0, 4.0, 5.0)

printfn "Circle (radius 5)         : Area = %.2f, Perimeter = %.2f" (circle.Area()) (circle.Perimeter())
printfn "Rectangle (4 x 6)         : Area = %.2f, Perimeter = %.2f" (rectangle.Area()) (rectangle.Perimeter())
printfn "Triangle (sides 3, 4, 5)  : Area = %.2f, Perimeter = %.2f" (triangle.Area()) (triangle.Perimeter())
printfn ""

printfn "=== Polymorphism: one list of Shape, each call runs the overridden version ==="
let shapes : Shape list =
    [ 
      circle :> Shape
      rectangle :> Shape
      triangle :> Shape
      Circle(1.5) :> Shape
      Rectangle(2.0, 2.0) :> Shape ]

shapes
|> List.iter (fun shape -> printfn "%s   [runtime type: %s]" (shape.Describe()) (shape.GetType().Name))

printfn ""
printfn "Total area of all shapes: %.2f" (shapes |> List.sumBy (fun shape -> shape.Area()))
printfn "Largest shape by area   : %s" ((shapes |> List.maxBy (fun shape -> shape.Area())).Name)