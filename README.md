# RecaptchaV3 Solver using GO

## Usage

```go
func main() {
	bg := "" //in some cases can be left empty
	getURL := "https://www.google.com/recaptcha/api2/anchor?ar=1&k=6LcR_okUAAAAAPYrPe-HK_0RULO1aZM15ENyM-Mf&co=aHR0cHM6Ly9hbnRjcHQuY29tOjQ0Mw..&hl=en&v=rKbTvxTxwcw5VqzrtN-ICwWt&size=invisible&cb=vzuqg89a4k12"
	postURL := "https://www.google.com/recaptcha/api2/reload?k=6LcR_okUAAAAAPYrPe-HK_0RULO1aZM15ENyM-Mf"
	result, err := BypassV3(getURL, postURL, bg)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
		ipscore, err := postRequest(result)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		println(ipscore)
	}
}
```
