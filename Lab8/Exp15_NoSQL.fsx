
#r "nuget: LiteDB"

open LiteDB

type Student =
    {
        Id: int
        Name: string
        Age: int
        Course: string
    }

let db = new LiteDatabase("students.db")
let students = db.GetCollection<Student>("students")

let printStudents () =
    let records = students.FindAll() |> Seq.toList

    if List.isEmpty records then
        printfn "No students found."
    else
        records
        |> List.iter (fun s ->
            printfn "ID: %d, Name: %s, Age: %d, Course: %s"
                s.Id s.Name s.Age s.Course
        )

let student =
    {
        Id = 1
        Name = "John"
        Age = 21
        Course = "Computer Science"
    }

printfn "\nCREATE"
students.Insert(student) |> ignore
printStudents()

printfn "\nREAD"
let result = students.FindById(1)
printfn "%A" result

printfn "\nUPDATE"
let updatedStudent =
    { student with Age = 22; Course = "Embedded Systems" }

students.Update(updatedStudent) |> ignore
printStudents()

printfn "\nDELETE"
students.Delete(1) |> ignore
printStudents()

db.Dispose()
