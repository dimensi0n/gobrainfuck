# Gobrainfuck

Just get go brainfuck:

```
go get -u github.com/dimensi0n/gobrainfuck
```

Import it:

```golang
import "github.com/dimensi0n/gobrainfuck"
```

Create a string wich contain the brainfucl code and compile it with the `Compile()` function:

```golang
string := ">+++.>-<--.>.>+"
buffer := gobrainfuck.Compile(string)
fmt.Println(buffer)
```

It will return:

```
3
1
-1
[0 1 -1 1 0 0 0 0 0 0 0 0 0 0 0]
```