启用cgo
set cgo_enabled=1

go build -buildmode=c-shared -o add.dll add.go
dumpbin.exe /exports .\add.dll /out:add.txt

> go build 里 加上 -ldflags="-w -s" 会去除 DWARF调试信息、符号信息，以减小体积
> 生成C可调用的库时，Go源代码需要以下几个注意。
> 必须导入 “C” 包
> 必须在可外部调用的函数前加上 【//export 函数名】的注释
> 必须是main包，切含有main函数，main函数可以什么都不干

gcc .\test.c -o test.exe .\add.dll