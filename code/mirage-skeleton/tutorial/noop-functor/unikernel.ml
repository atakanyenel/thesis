open Lwt.Infix

module Main = struct

  let start =
    Logs.info (fun f -> f "hello");
    Lwt.return_unit

end
