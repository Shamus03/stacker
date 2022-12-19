# stacker

# ⚠ Don't use this project.  Go supports generics now - do that instead. ⚠

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
