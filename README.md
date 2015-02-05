# trigger
stopwatch function without benchmark test

####Import

```
import "github.com/zier/trigger"
```

####Simple use
```
tr := trigger.Start()
// do somthing
dur := tr.Stop()

fmt.Println(dur.Nanoseconds()) // Get durtaion nanasec
```

#### Reset time
```
tr.Reset()
```