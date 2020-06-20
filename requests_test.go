package deepl_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/go-numb/go-deepl"
	"github.com/go-numb/go-deepl/translate"
)

func TestRequest(t *testing.T) {
	key := os.Getenv("DEEPLKEY")
	if key == "" {
		t.Fatal("key undefined")
	}
	c := deepl.New(key)

	start := time.Now()
	defer func() {
		fmt.Println("exec time: ", time.Since(start))
	}()

	res := &translate.Response{}
	if err := c.Do(&translate.Request{
		Text:       []string{"translate for you", "Bitte übersetzen"},
		TargetLang: "JA",
	}, res); err != nil {
		t.Fatal(err)
	}

	res.Unescape()
	fmt.Printf("%+v\n", res)
}
