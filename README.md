# ziptz

US zip code to IANA timezone lookup.

## Install

```bash
go get github.com/kjbreil/ziptz
```

## Usage

```go
import "github.com/kjbreil/ziptz"

ziptz.Lookup("60601")        // "America/Chicago"
ziptz.Location("60601")      // *time.Location
ziptz.Abbreviation("60601")  // "CST" or "CDT"
ziptz.Offset("60601")        // "-06:00"
ziptz.IsDST("60601")         // true/false
```

Invalid or unknown zip codes return empty string (or error for `Location`).

## Update Data

```bash
./scripts/update-data.sh
```

Data source: [seanpianka/Zipcodes](https://github.com/seanpianka/Zipcodes) (MIT License)
