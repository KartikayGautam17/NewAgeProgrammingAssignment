open System

let readMarks (subject: string) =
    let rec loop () =
        printf "Enter marks for %s (0-100): " subject
        match Double.TryParse(Console.ReadLine()) with
        | true, m when m >= 0.0 && m <= 100.0 -> m
        | _ ->
            printfn "Invalid input. Please enter a number between 0 and 100."
            loop ()
    loop ()

let division (percentage: float) =
    match percentage with
    | p when p >= 60.0 -> "First Division"
    | p when p >= 50.0 -> "Second Division"
    | p when p >= 40.0 -> "Third Division"
    | _ -> "Fail"

let subjects = [ "Subject 1"; "Subject 2"; "Subject 3"; "Subject 4" ]
let marks = subjects |> List.map readMarks
let total = List.sum marks
let percentage = total / float subjects.Length

printfn "Total Marks : %.2f / %d" total (subjects.Length * 100)
printfn "Percentage  : %.2f%%" percentage
printfn "Result      : %s" (division percentage)