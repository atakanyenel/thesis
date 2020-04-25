open Lwt.Infix

module Hello (Time : Mirage_time.S) = struct

  let start _time =

    let hello = Key_gen.hello () in

    let rec loop () =
        let _= Sys.command @@  "say " ^ hello in
        Time.sleep_ns (Duration.of_sec 1) >>= fun () ->
        loop ()
    in
    loop ()

end
