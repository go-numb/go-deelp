# go-deepl
This package is wrapper DeepL Translation API, with golang.

# Usage
```go
package main

func main() {
    // Subscription to DeepL
	key := os.Getenv("DEEPLKEY")
	if key == "" {
		t.Fatal("key undefined")
    }
    // Created client
	c := deepl.New(key)

    // Response and Processing time
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

    // Unscape query for ja.
	res.Unescape()
	fmt.Printf("%+v\n", res)
}
```


## Author
[@_numbP](https://twitter.com/_numbp)