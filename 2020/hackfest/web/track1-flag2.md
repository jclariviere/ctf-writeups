# Track 1 - Flag 2

The git repo contained the source code for the backend. There we could clearly see the 3 remaining challenges for the track.

The first one was about setting the `active` and the `admin` flag to 1 in our user account.
By looking in the source code, we find that the only place where the `admin` flag is being set is in the `crud.EditUserEntry()` function, which is routed to the `POST /api/edit-a-38` endpoint:

```
    a.Post("/api/edit-a-38", a.EditUserEntry)
```

That endpoint was hit whenever we updated our A-38 form. If we look at the source code, we can see that we cannot edit the `admin` flag directly.
However, it is part of the `userParamsArray`, which is vulnerable to an overflow vulnerability.

```go
    var userParamsArray [7]uint8

    userParamsArray[0] = uint8(crewCount)
    userParamsArray[1] = uint8(numberOfPaddles)
    userParamsArray[2] = uint8(numberOfSails)
    userParamsArray[3] = uint8(age)
    userParamsArray[5] = *user.Active
    userParamsArray[6] = *user.Admin

    var weightOfMerchandiseCasted interface{}

    if weightOfMerchandise < 0 {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Invalid weight")
        return
    }
    if weightOfMerchandise <= 255 {
        weightOfMerchandiseCasted = uint8(weightOfMerchandise)
        weightOfMerchandiseUint8, _ := weightOfMerchandiseCasted.(uint8)
        userParamsArray[4] = weightOfMerchandiseUint8
    } else if weightOfMerchandise <= 65535 {
        weightOfMerchandiseCasted = uint16(weightOfMerchandise)
        weightOfMerchandiseUint16, _ := weightOfMerchandiseCasted.(uint16)
        *(*uint16)(unsafe.Pointer(&userParamsArray[4])) = weightOfMerchandiseUint16
    } else if weightOfMerchandise <= 4294967295 {
        weightOfMerchandiseCasted = uint32(weightOfMerchandise)
        weightOfMerchandiseUint32, _ := weightOfMerchandiseCasted.(uint32)
        *(*uint32)(unsafe.Pointer(&userParamsArray[4])) = weightOfMerchandiseUint32
    } else {
        weightOfMerchandiseCasted = uint64(weightOfMerchandise)
        weightOfMerchandiseUint64, _ := weightOfMerchandiseCasted.(uint64)
        *(*uint64)(unsafe.Pointer(&userParamsArray[4])) = weightOfMerchandiseUint64
    }

    encodedIP := utils.EncodeIP(ip)

    db.Model(&user).Updates(models.User{IP: &encodedIP, Name: &name, Surname: &surname, GalleyRegistrationID: &galleyRegistrationID, CountryOfOrigin: &countryOfOrigin, MainMaterial: &mainMaterial, YearOfFabrication: &yearOfFabricationUint16, WeightOfMerchandise: &userParamsArray[4], CrewCount: &userParamsArray[0], NumberOfPaddles: &userParamsArray[1], NumberOfSails: &userParamsArray[2], Age: &userParamsArray[3], Active: &userParamsArray[5], Admin: &userParamsArray[6]})
```

Here, we can see that `weightOfMerchandise` is being casted using unsafe pointers which, given a big number, will overflow in the array and override the `active` and `admin` values.
I could have used a fancy way to find exactly such value, but I was lazy and used some trial/error.
In the end, setting the weight of merchandise parameter to `65792` did exactly that, and I got the flag by navigating to the yellow form afterwards.
