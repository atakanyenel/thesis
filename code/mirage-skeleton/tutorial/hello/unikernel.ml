
module Hello (Time : Mirage_time_lwt.S) = struct

  let start _time =

    let rec loop = function
      | 0 -> Lwt.return_unit
      | n ->
        Logs.info (fun f -> f "hello");
        loop (n-1)
    in
    loop 1

end