# Track 1 - Flag 4

The last challenge of the first track was about setting our IP to `0.0.0.0` or `TO BE IMPLEMENTED` in the database.

```go
func IPFlag(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// ...

	if utils.DecodeIP(*user.IP) == "0.0.0.0" || utils.DecodeIP(*user.IP) == "TO BE IMPLEMENTED" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, flag4)
		return
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Sorry, your IP does not match the one stored in the database. ")
		return
	}

}
```

The IP is set in both the `crud.CreateUserEntry()` and `crud.EditUserEntry()` functions.
Both set it by first calling `utils.GetIP()`, then encoding it with `utils.EncodeIP()`.

```go
func GetIP(r *http.Request) (string, error) {
	forwardedSlice := strings.Split(r.Header.Get("X-Forwarded-For"), ", ")
	forwarded := forwardedSlice[0]
	ipRegex := regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)

	quadZeroRegex := regexp.MustCompile(`0+\.0+\.0+\.0+`)

	if quadZeroRegex.MatchString(forwarded) {
		return r.RemoteAddr, errors.New("Sorry but this is very unsafe and therefor cannot be allowed")
	}

	if ipRegex.MatchString(forwarded) {
		return forwarded, nil
	}
	return r.RemoteAddr, nil
}
```

We can see that the IP can be set by using the `X-Forwarded-For` header.
However, we cannot set our IP to `0.0.0.0` directly because it matches the `quadZeroRegex`.
We cannot set it to `TO BE IMPLEMENTED` either because it must match the `ipRegex`.

We can however set bogus IPs, which we can use to exploit the `DecodeIP()/EncodeIP()` functions.

```go
func EncodeIP(rawIP string) (decimalIP uint32) {

	ipSlice := strings.Split(rawIP, ".")

	pos := 24

	for _, octet := range ipSlice {
		castedOctet, _ := strconv.Atoi(octet)
		decimalIP += uint32(castedOctet) << pos
		pos -= 8
	}
	return
}

func DecodeIP(decimalIP uint32) (rawIP string) {

	pos := 24

	var ipArray []string

	for i := 0; i < 4; i++ {
		ipArray = append(ipArray, strconv.Itoa(int(uint8(decimalIP>>pos))))
		pos -= 8
	}

	rawIP = fmt.Sprintf(strings.Join(ipArray, "."))
	return
}
```

That part was solved by one of my teammate so I didn't go into the details, but basically you can pass `X-Forwarded-For: 256.0.0.0` and the encoding process will translate it to `0.0.0.0`.

After doing that, we can navigate to the brown form and get the flag.
