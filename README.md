## EdGrep 
Is a simple clone of the unix tool grep codes probs a bit gnarly but learning Go so who cares. This is based off the John Crickett [coding challenge](https://codingchallenges.fyi/challenges/challenge-grep).

## Setup:
```
go mod download
go build -o /usr/local/bin/edgrep
```
## args:
```
       pattern
       filepath
[-i    case sensitivity flag]
[-r    recursive flag]
[-v    exclude a word]
```
**Usage**
basic usage:
```
edgrep "" <filepath>
```
recursive usage:
```
edgrep -r "<MY_PATTERN>" <DIR_PATH>
```

## testing:
- unit tests:
```
go test -v -tags=unit 
```
- integration tests:
```
go test -v -tags=integration
```

## Extension:
- switch to using cobra for easier CLI development lots of hand cranking
- better error handling for the basic grepfile logic
