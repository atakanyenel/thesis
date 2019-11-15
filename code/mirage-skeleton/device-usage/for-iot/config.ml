open Mirage

let disk = generic_kv_ro "t"

let main =
  foreign
    ~packages:[package "duration"]
    "Unikernel.Main" (kv_ro @-> time @-> job)

let () =
  register "kv_ro" [main $ disk $ default_time] 