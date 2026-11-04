open System
open System.Diagnostics
open System.Threading

let stopwatch = Stopwatch.StartNew()

let log (message: string) =
    printfn "[%5d ms | thread %2d] %s" stopwatch.ElapsedMilliseconds Thread.CurrentThread.ManagedThreadId message

let ioTask name delayMs result =
    async {
        log (sprintf "%s started (simulated I/O, %d ms)" name delayMs)
        do! Async.Sleep delayMs
        log (sprintf "%s finished" name)
        return (name, result)
    }

let countPrimesBelow limit =
    [ 2 .. limit ]
    |> List.filter (fun n -> [ 2 .. int (sqrt (float n)) ] |> List.forall (fun d -> n % d <> 0))
    |> List.length

let computeTask name limit =
    async {
        log (sprintf "%s started (compute-intensive)" name)
        let count = countPrimesBelow limit
        log (sprintf "%s finished" name)
        return (name, sprintf "%d primes below %d" count limit)
    }

let tasks =
    [ ioTask "Task 1: Download file" 1500 "2.4 MB received"
      ioTask "Task 2: Query database" 1000 "42 rows returned"
      computeTask "Task 3: Count primes" 150000
      ioTask "Task 4: Call web API" 500 "HTTP 200 OK" ]

let timed work =
    stopwatch.Restart()
    let output = work ()
    (output, stopwatch.ElapsedMilliseconds)

printfn "=== Parallel execution (Async.Parallel) ==="
printfn "All tasks are launched together; nobody waits for the previous one."
printfn ""

let (parallelResults, parallelMs) =
    timed (fun () -> tasks |> Async.Parallel |> Async.RunSynchronously)

printfn ""
printfn "Results (returned in the order the tasks were listed):"
parallelResults |> Array.iter (fun (name, result) -> printfn "  %s -> %s" name result)
printfn ""
printfn "Parallel total: %d ms" parallelMs
printfn ""

printfn "=== Sequential execution (for comparison) ==="
printfn ""

let (_, sequentialMs) =
    timed (fun () -> tasks |> List.map Async.RunSynchronously)

printfn ""
printfn "Sequential total: %d ms" sequentialMs
