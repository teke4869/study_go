package greeting

import(
	"fmt"
)

//nameに挨拶をするメッセージを返す
func GetGreetMsg(name string) (string,error) {
	if name == "" {
		return "", fmt.Errorf("empty name")
	}
	msg := fmt.Sprintf("Hello, %v!", name)

	return msg, nil
}