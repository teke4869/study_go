package greeting

import(
	"testing"
)

func TestGetGreetMsg(t *testing.T) {
    testCases := []struct {
        name     string
        expected string
    }{
        {"shuya", "Hello, shuya!"},
        {"goto", "Hello, goto!"},
    }

    for _, tc := range testCases {
        result, err := GetGreetMsg(tc.name)
        if err != nil {
            t.Errorf("GetGreetMsg(%s) = %s; expected %s", tc.name, result, tc.expected)
        }
        
        if result != tc.expected {
            t.Errorf("GetGreetMsg(%s) = %s; expected %s", tc.name, result, tc.expected)
        }
    }

    //  空文字を入れるとエラーになることを確認
    result, err := GetGreetMsg("")

    if err == nil {
        t.Errorf("GetGreetMsg(\"\") = %s; expected error", result)
    }
}