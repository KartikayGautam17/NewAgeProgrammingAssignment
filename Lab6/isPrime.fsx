open System

let isPrime n =
    let rec hasNoDivisorFrom d =
        d * d > n || (n % d <> 0 && hasNoDivisorFrom (d + 1))
    n > 1 && hasNoDivisorFrom 2

let rec readInt prompt =
    printf "%s" prompt
    match Int32.TryParse(Console.ReadLine()) with
    | true, value -> value
    | _ ->
        printfn "Invalid input. Please enter an integer."
        readInt prompt

let describe n =
    sprintf "%d -> %s" n (if isPrime n then "Prime" else "Not prime")

[ 1; 2; 3 ]
|> List.map (fun i -> readInt (sprintf "Enter integer %d: " i))
|> List.map describe
|> List.iter (printfn "%s")