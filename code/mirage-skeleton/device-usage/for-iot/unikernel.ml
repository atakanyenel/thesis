open Lwt.Infix

module Main (KV: Mirage_kv.RO) (Time: Mirage_time.S) = struct

  let start kv _time=

    let read_from_file kv =
        KV.get kv (Mirage_kv.Key.v "secret") >|= function
            | Error e ->
                Logs.warn (fun f -> f "Could not compare the secret against a known constant: %a"
                KV.pp_error e)
            | Ok stored_secret ->
                Logs.info (fun f -> f "Data -> %a" Format.pp_print_string stored_secret);
               
    in
        let rec loop() =
        read_from_file kv >>= fun()->
        Time.sleep_ns (Duration.of_sec 2)>>= fun () ->
        loop()
        in 
        loop()
end