open Mirage

let disk = generic_kv_ro "sensors"
let key =
  let doc = Key.Arg.info ~doc:"Get name of sensor" ["sensor"] in
  Key.(create "filename" Arg.(opt string "secret" doc))
let main =
  foreign
    ~keys:[Key.abstract key]
    ~packages:[package "duration"]
    "Unikernel.Main" (kv_ro @-> time @-> job)

let () =
  register "kv_ro" [main $ disk $ default_time] 