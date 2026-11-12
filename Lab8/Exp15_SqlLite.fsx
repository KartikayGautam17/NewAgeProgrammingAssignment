#r "nuget: Microsoft.Data.Sqlite, 8.0.0"

open System
open Microsoft.Data.Sqlite

type Student = { Id: int; Name: string; Age: int; Course: string }

let connectionString = "Data Source=students.db"

let buildCommand (connection: SqliteConnection) (sql: string) (parameters: (string * obj) list) =
    let command = connection.CreateCommand()
    command.CommandText <- sql
    parameters
    |> List.iter (fun (name, value) -> command.Parameters.AddWithValue(name, value) |> ignore)
    command

let execute connection sql parameters =
    use command = buildCommand connection sql parameters
    command.ExecuteNonQuery()

let query connection sql parameters =
    use command = buildCommand connection sql parameters
    use reader = command.ExecuteReader()
    [ while reader.Read() do
        yield
            { Id = reader.GetInt32 0
              Name = reader.GetString 1
              Age = reader.GetInt32 2
              Course = reader.GetString 3 } ]

let selectAll = "SELECT Id, Name, Age, Course FROM Students"

let readAll connection =
    query connection (selectAll + " ORDER BY Id") []

let readById connection (id: int) =
    query connection (selectAll + " WHERE Id = @id") [ ("@id", box id) ]

let readByCourse connection (course: string) =
    query connection (selectAll + " WHERE Course = @course ORDER BY Id") [ ("@course", box course) ]

let createStudent connection (name: string, age: int, course: string) =
    execute
        connection
        "INSERT INTO Students (Name, Age, Course) VALUES (@name, @age, @course)"
        [ ("@name", box name); ("@age", box age); ("@course", box course) ]
    |> ignore

let updateStudent connection (id: int) (age: int) (course: string) =
    execute
        connection
        "UPDATE Students SET Age = @age, Course = @course WHERE Id = @id"
        [ ("@age", box age); ("@course", box course); ("@id", box id) ]

let deleteStudent connection (id: int) =
    execute connection "DELETE FROM Students WHERE Id = @id" [ ("@id", box id) ]

let printState (label: string) (students: Student list) =
    printfn "--- %s ---" label
    match students with
    | [] -> printfn "  (no records)"
    | _ -> students |> List.iter (fun s -> printfn "  %d | %-8s | %2d | %s" s.Id s.Name s.Age s.Course)
    printfn ""

use connection = new SqliteConnection(connectionString)
connection.Open()
printfn "Connected to SQLite database: %s" connectionString
printfn ""

execute connection "DROP TABLE IF EXISTS Students" [] |> ignore

execute
    connection
    "CREATE TABLE Students (Id INTEGER PRIMARY KEY AUTOINCREMENT, Name TEXT NOT NULL, Age INTEGER NOT NULL, Course TEXT NOT NULL)"
    []
|> ignore

printState "Initial state" (readAll connection)

printfn "=== CREATE ==="
[ ("Asha", 20, "Computer Science")
  ("Ravi", 21, "Mathematics")
  ("Meera", 19, "Computer Science") ]
|> List.iter (createStudent connection)
printState "After inserting 3 students" (readAll connection)

printfn "=== READ ==="
printState "All students" (readAll connection)
printState "Student with Id = 2" (readById connection 2)
printState "Students in Computer Science" (readByCourse connection "Computer Science")

printfn "=== UPDATE ==="
let updated = updateStudent connection 2 22 "Data Science"
printfn "Rows updated: %d" updated
printfn ""
printState "After updating student 2 (Age 22, Course Data Science)" (readAll connection)

printfn "=== DELETE ==="
let deleted = deleteStudent connection 3
printfn "Rows deleted: %d" deleted
printfn ""
printState "After deleting student 3" (readAll connection)