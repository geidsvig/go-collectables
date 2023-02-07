
This package was created to approximate Scala Collections in Go.


Including the library in your project:
```
require (
    github.com/geidsvig/go-collectables latest
)
```

```
import collectables "github.com/geidsvig/go-collectables/lib"
```

Example struct to be used as a Collectable.
```
type Collectable struct {
	name  string
	value int
}
```

Example Usage:
```
T := reflect.TypeOf(Collectable{})
list := collectables.Collection{}
item1 := Collectable{name: "A", value: 1}
err := list.Append(T, item1)
```

