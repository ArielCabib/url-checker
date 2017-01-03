# url-checker

check a list of urls for http status codes

**example usage**

Let's say bla.csv looks like this:

```
http://google.com
http:/google.com
httpawesfd://google.com
```


run:
```sh
go build
./url-checker bla.csv
```

and you'll get a nice csv describing the statuses of all input lines in the same order.
I didn't use go routines because it caused lots of 429. The sites were sure they are DoS'ed.

## Compiling for different OS's

Look here for all possibilities: https://golang.org/doc/install/source#environment
For example, compile for windows:

```
GOOS=windows GOARCH=amd64 go build
```
