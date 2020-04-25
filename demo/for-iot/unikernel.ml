open Lwt.Infix

module Main (KV: Mirage_kv.RO) (Time: Mirage_time.S) = struct

  let start kv _time=

    let read_from_file kv filename =
        KV.get kv (Mirage_kv.Key.v filename) >|= function
            | Error e ->
                Logs.warn (fun f -> f "Cannot find the file %a"
                KV.pp_error e)
            | Ok sensor_value ->
                Logs.info (fun f -> f "Reading from: %s -> %s" filename sensor_value);
    in
        let filename=Key_gen.filename() in
        let rec loop() =
        read_from_file kv filename >>= fun()->
        Time.sleep_ns (Duration.of_sec 2)>>= fun () ->
        loop()
        in
        loop()
end