package main

/* 独自定義のエラーを表す型 */
type MyError struct {
  Message string
  ErrCode int
}

/* errorインタフェースのメソッドを実装 */
func (e *MyError) Error() string {
  return e.Message
}

func RaiseError() error {
  return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

func main() {
  err := RaiseError()
  err.Error()
}
