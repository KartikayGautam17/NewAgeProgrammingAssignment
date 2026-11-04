let numbers1 = [ 1; 2; 3; 4; 5; 6; 7; 8; 9; 10 ]
let numbers2 = [ -3; 0; 7; 12; 15; 20; -8 ]
let numbers3 = [ 5 ]

printfn "=== 1. List.map ==="
printfn "Square %A -> %A" numbers1 (List.map (fun x -> x * x) numbers1)
printfn "Square %A -> %A" numbers2 (List.map (fun x -> x * x) numbers2)
printfn "Add 10 to %A -> %A" numbers1 (List.map (fun x -> x + 10) numbers1)
printfn ""

printfn "=== 2. List.filter ==="
printfn "Even numbers in %A -> %A" numbers1 (List.filter (fun x -> x % 2 = 0) numbers1)
printfn "Even numbers in %A -> %A" numbers2 (List.filter (fun x -> x % 2 = 0) numbers2)
printfn "Greater than 5 in %A -> %A" numbers2 (List.filter (fun x -> x > 5) numbers2)
printfn ""

printfn "=== 3. List.fold / List.reduce ==="
printfn "Sum of %A (fold) -> %d" numbers1 (List.fold (fun acc x -> acc + x) 0 numbers1)
printfn "Sum of %A (fold) -> %d" numbers2 (List.fold (fun acc x -> acc + x) 0 numbers2)
printfn "Product of %A (fold) -> %d" numbers1 (List.fold (fun acc x -> acc * x) 1 numbers1)
printfn "Maximum of %A (reduce) -> %d" numbers1 (List.reduce (fun a b -> if a > b then a else b) numbers1)
printfn "Maximum of %A (reduce) -> %d" numbers2 (List.reduce (fun a b -> if a > b then a else b) numbers2)
printfn "Maximum of %A (reduce) -> %d" numbers3 (List.reduce (fun a b -> if a > b then a else b) numbers3)
printfn ""

printfn "=== 4. Combined: sum of squares of the even numbers ==="
printfn "%A -> %d" numbers1
    (numbers1
     |> List.filter (fun x -> x % 2 = 0)
     |> List.map (fun x -> x * x)
     |> List.fold (fun acc x -> acc + x) 0)