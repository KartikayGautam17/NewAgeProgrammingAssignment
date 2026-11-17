
#r "nuget: Microsoft.Data.Sqlite"

open System
open Microsoft.Data.Sqlite

type Student =
    {
        Id: int
        Name: string
        Age: int
        Course: string
        Active: bool
    }

type StudentStatus =
    | Minor
    | Adult
    | Senior

let classifyStudent age =
    match age with
    | n when n < 18 -> Minor
    | n when n < 25 -> Adult
    | _ -> Senior

let getStatus student =
    match classifyStudent student.Age with
    | Minor -> "Minor"
    | Adult -> "Adult"
    | Senior -> "Senior"

let connectionString = "Data Source=students.db"

let openConnection () =
    async {
        let connection = new SqliteConnection(connectionString)
        do! connection.OpenAsync() |> Async.AwaitTask
        return connection
    }

let initializeDatabase () =
    async {
        use! connection = openConnection()
        use command = connection.CreateCommand()

        command.CommandText <-
            """
            CREATE TABLE IF NOT EXISTS Students (
                Id INTEGER PRIMARY KEY AUTOINCREMENT,
                Name TEXT NOT NULL,
                Age INTEGER NOT NULL,
                Course TEXT NOT NULL,
                Active INTEGER NOT NULL
            );
            """

        let! _ =
            command.ExecuteNonQueryAsync()
            |> Async.AwaitTask

        printfn "Database initialized."
    }

let addStudent student =
    async {
        use! connection = openConnection()
        use command = connection.CreateCommand()

        command.CommandText <-
            """
            INSERT INTO Students (Name, Age, Course, Active)
            VALUES (@Name, @Age, @Course, @Active);
            """

        command.Parameters.AddWithValue("@Name", student.Name) |> ignore
        command.Parameters.AddWithValue("@Age", student.Age) |> ignore
        command.Parameters.AddWithValue("@Course", student.Course) |> ignore
        command.Parameters.AddWithValue("@Active", student.Active) |> ignore

        let! result =
            command.ExecuteNonQueryAsync()
            |> Async.AwaitTask

        printfn "Inserted %d student record(s)." result
    }

let getStudents () =
    async {
        use! connection = openConnection()
        use command = connection.CreateCommand()

        command.CommandText <-
            "SELECT Id, Name, Age, Course, Active FROM Students;"

        use! reader =
            command.ExecuteReaderAsync()
            |> Async.AwaitTask

        let students = ResizeArray<Student>()
        let mutable reading = true

        while reading do
            let! hasRow =
                reader.ReadAsync()
                |> Async.AwaitTask

            if hasRow then
                students.Add
                    {
                        Id = reader.GetInt32(0)
                        Name = reader.GetString(1)
                        Age = reader.GetInt32(2)
                        Course = reader.GetString(3)
                        Active = reader.GetInt32(4) = 1
                    }
            else
                reading <- false

        return students |> Seq.toList
    }

let updateStudent id newCourse =
    async {
        use! connection = openConnection()
        use command = connection.CreateCommand()

        command.CommandText <-
            """
            UPDATE Students
            SET Course = @Course
            WHERE Id = @Id;
            """

        command.Parameters.AddWithValue("@Course", newCourse) |> ignore
        command.Parameters.AddWithValue("@Id", id) |> ignore

        let! result =
            command.ExecuteNonQueryAsync()
            |> Async.AwaitTask

        printfn "Updated %d record(s)." result
    }

let deleteStudent id =
    async {
        use! connection = openConnection()
        use command = connection.CreateCommand()

        command.CommandText <-
            "DELETE FROM Students WHERE Id = @Id;"

        command.Parameters.AddWithValue("@Id", id) |> ignore

        let! result =
            command.ExecuteNonQueryAsync()
            |> Async.AwaitTask

        printfn "Deleted %d record(s)." result
    }

let processStudents students =
    let activeStudents =
        students
        |> List.filter (fun student -> student.Active)

    let eligibleStudents =
        activeStudents
        |> List.filter (fun student -> student.Age >= 18)

    eligibleStudents
    |> List.map (fun student ->
        sprintf "%s | %s | Age: %d | %s"
            student.Name
            student.Course
            student.Age
            (getStatus student)
    )

let main =
    async {
        do! initializeDatabase()

        let sampleStudents =
            [
                {
                    Id = 0
                    Name = "Alice"
                    Age = 20
                    Course = "Computer Science"
                    Active = true
                }
                {
                    Id = 0
                    Name = "Bob"
                    Age = 17
                    Course = "Mathematics"
                    Active = true
                }
                {
                    Id = 0
                    Name = "Charlie"
                    Age = 23
                    Course = "Electronics"
                    Active = false
                }
                {
                    Id = 0
                    Name = "David"
                    Age = 26
                    Course = "Computer Science"
                    Active = true
                }
            ]

        for student in sampleStudents do
            do! addStudent student

        let! students = getStudents()

        printfn "\nAll Students:"
        students |> List.iter (printfn "%A")

        printfn "\nProcessed Students:"
        let processed = processStudents students
        processed |> List.iter (printfn "%s")

        printfn "\nUpdating student:"
        do! updateStudent 1 "Artificial Intelligence"

        let! updatedStudents = getStudents()
        updatedStudents |> List.iter (printfn "%A")

        printfn "\nDeleting student:"
        do! deleteStudent 2

        let! finalStudents = getStudents()

        printfn "\nFinal Student List:"
        finalStudents |> List.iter (printfn "%A")

        printfn "\nApplication completed."
    }

Async.RunSynchronously main
