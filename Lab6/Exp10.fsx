open System
open System.Numerics
open System.Text.RegularExpressions

let reverse (s: string) =
    s |> Seq.rev |> Seq.map string |> String.concat ""

let normalize (s: string) =
    s
    |> Seq.filter Char.IsLetterOrDigit
    |> Seq.map (Char.ToLower >> string)
    |> String.concat ""

let isPalindrome (s: string) =
    let cleaned = normalize s
    cleaned = reverse cleaned

let isAllLetters (s: string) = Seq.forall Char.IsLetter s

let isAllDigits (s: string) = Seq.forall (fun c -> c >= '0' && c <= '9') s

let transform (s: string) =
    match s with
    | _ when String.IsNullOrEmpty s -> s
    | _ when isAllLetters s -> s.ToUpper()
    | _ when isAllDigits s -> (BigInteger.Parse s * 2I).ToString()
    | _ -> s

let emailPattern = @"^[A-Za-z0-9._%+-]+@[A-Za-z0-9-]+(\.[A-Za-z0-9-]+)*\.[A-Za-z]{2,}$"

let isValidEmail (s: string) = Regex.IsMatch(s, emailPattern)

let demo title describe inputs =
    printfn "%s" title
    inputs
    |> List.iter (fun input ->
        printfn "  %-30s -> %s" (sprintf "\"%s\"" input) (describe input))
    printfn ""

demo
    "1. Reverse and palindrome check"
    (fun s -> sprintf "reversed: \"%s\", palindrome: %b" (reverse s) (isPalindrome s))
    [ "racecar"; "Hello"; "A man, a plan, a canal: Panama"; "12321" ]

demo
    "2. Transform"
    (fun s -> sprintf "\"%s\"" (transform s))
    [ "hello"; "Functional"; "12345"; "007"; "abc123"; "F#4" ]

demo
    "3. Email validation"
    (fun s -> if isValidEmail s then "Valid" else "Invalid")
    [ "user@example.com"; "first.last@mail.co.in"; "invalid@"; "no-at-sign.com"; "a@b" ]