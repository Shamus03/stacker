# stacker

Install:

```bash
go get github.com/shamus03/stacker
```

Use (via `go generate`):

```go
// This will generate a type `coolThingStack`
//go:generate stacker -type coolThing

type coolThing struct {
    howCool int
}
```
