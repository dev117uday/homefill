# Project HomeFill Backend

## Things todo (Project Level)
// TODO Create a error handling system
// TODO Use some sort of logging system

### Code Snippets

```go
config.Log.WithFields(logrus.Fields{
	"fn":  "ConnectTODB",
	"err": err.Error(),
}).Fatal("unable to connect to db")
```