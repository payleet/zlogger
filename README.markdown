# zlogger
--
    import "github.com/payleet/zlogger"


## Usage

#### func  DPanic

```go
func DPanic(message string, value interface{})
```
DPanic logs at dpanic level

#### func  Debug

```go
func Debug(message string, value interface{})
```
Debug logs at debug level

#### func  Error

```go
func Error(message string, value interface{})
```
Error logs at error level

#### func  Fatal

```go
func Fatal(message string, value interface{})
```
Fatal logs at fatal level and calls os.Exit(1) after

#### func  Info

```go
func Info(message string, value interface{})
```
Info loges at info level

#### func  Panic

```go
func Panic(message string, value interface{})
```
Panic logs at panic level

#### type Fileds

```go
type Fileds map[string]interface{}
```

Fileds key/val

#### type Logger

```go
type Logger struct {
}
```

Logger is a wraper struct

#### func  NewLogger

```go
func NewLogger() *Logger
```
NewLogger Returns Logger

#### func (*Logger) DPanic

```go
func (l *Logger) DPanic(message string, val ...interface{})
```
DPanic logs at panic level

#### func (*Logger) Debug

```go
func (l *Logger) Debug(message string, val ...interface{})
```
Debug logs at debug level

#### func (*Logger) Error

```go
func (l *Logger) Error(message string, val ...interface{})
```
Error logs at error level

#### func (*Logger) Fatal

```go
func (l *Logger) Fatal(message string, val ...interface{})
```
Fatal logs at fatal level then calls os.Exit(1)

#### func (*Logger) Info

```go
func (l *Logger) Info(message string, val ...interface{})
```
Info logs at info level

#### func (*Logger) Panic

```go
func (l *Logger) Panic(message string, val ...interface{})
```
Panic logs at panic level

#### func (*Logger) Warn

```go
func (l *Logger) Warn(message string, val ...interface{})
```
Warn logs at warn level

#### func (*Logger) WithError

```go
func (l *Logger) WithError(err error) *Logger
```
WithError returns a new logger with field error set to error

#### func (*Logger) WithField

```go
func (l *Logger) WithField(key string, value interface{}) *Logger
```
WithField returns Logger with Filed key set to value
