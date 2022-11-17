module path_pack

go 1.18
require (
    lib v0.0.0  //demo2引用
)

replace (
    lib => ./chapter/demo2/lib  //demo2引用
)