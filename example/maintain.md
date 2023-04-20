```go
    // Bad:
	// The use of = instead of := can change this line completely.
	if user, err = db.UserByID(userID); err != nil {
		// ...
	}
```

```go
    year := rand.Int()
    // Bad:
    // The ! in the middle of this line is very easy to miss.
    leap := (year%4 == 0) && (!(year%100 == 0) || (year%400 == 0))
```

```go
    // Good:
    u, err := db.UserByID(userID)
    if err != nil {
    return fmt.Errorf("invalid origin user: %s", err)
    }
    user = u
```

```go
    // Good:
    // Gregorian leap years aren't just year%4 == 0.
    // See https://en.wikipedia.org/wiki/Leap_year#Algorithm.
    var (
    leap4   = year%4 == 0
    leap100 = year%100 == 0
    leap400 = year%400 == 0
    )
    leap := leap4 && (!leap100 || leap400)
```