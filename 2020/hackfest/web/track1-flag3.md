# Track 1 - Flag 3

The next flag was about bypassing the MFA validation. It is being checked in this part of `flags.go`:

```go
func MFAFlag(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
    mfaOK := false

    defer utils.MFAVerdict(w, &mfaOK)

    var mfaType *string

    formatsArray := [3]string{"compact", "space", "dash"}
    formatsRegex := make(map[string]*regexp.Regexp)

    formatsRegex["main"] = regexp.MustCompile(`^\d{3}[\s-]?\d{3}[\s-]?\d{3}$`)
    formatsRegex["compact"] = regexp.MustCompile(`^\d{9}$`)
    formatsRegex["space"] = regexp.MustCompile(`^\d{3}\s\d{3}\s\d{3}$`)
    formatsRegex["dash"] = regexp.MustCompile(`^\d{3}-\d{3}-\d{3}$`)

    input := r.FormValue("mfa_string")

    if !formatsRegex["main"].MatchString(input) {
        return
    }

    if formatsRegex["compact"].MatchString(input) {
        mfaType = &formatsArray[0]
    } else if formatsRegex["space"].MatchString(input) {
        mfaType = &formatsArray[1]
    } else if formatsRegex["dash"].MatchString(input) {
        mfaType = &formatsArray[2]
    }

    utils.ValidateMFA(input, mfaType, &mfaOK)

}
```

This function first defers the `utils.MFAVerdict(w, &mfaOK)` function, then it sets some format regexes and validates the MFA code.

I honestly got a bit lucky here, I simply started by playing around with the regexes and this value gave me the flag: `123 456-789`.

The reason why this worked is because by using a value that matches the `formatsRegex["main"]` regex but not the others, the `mfaType` is never being set, resulting in a nil pointer.
This nil pointer is being passed to `utils.ValidateMFA()`.

```go
func ValidateMFA(mfaInput string, mfaType *string, mfaOK *bool) {

	*mfaOK = true

	var mfaCode [9]int

	for i := 0; i < 9; i++ {
		mfaCode[i] = rand.Intn(9)
	}

	if *mfaType == "compact" {
    // ...
```

In this function, `mfaOK` is first set to true, and since it is passed by pointer, this will also set `mfaOK` in the rest of the code execution.

Then, after creating the `mfaCode`, the `mfaType` is evaluated. Since this is a nil pointer, the function crashes with a `nil pointer dereference`.

The "defered" function `utils.MFAVerdict()` is then executed:

```go
func MFAVerdict(w http.ResponseWriter, mfaOK *bool) {

	flag3 := os.Getenv("FLAG3")

	if *mfaOK {
		recover()
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, flag3)
		return
	}
// ...
```

Since the `mfaOK` value was set to true earlier, the code recovers and prints the flag.
