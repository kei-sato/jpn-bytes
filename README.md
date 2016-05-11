# getting started

```
$ go get github.com/kei-sato/jpn-bytes
$ jpn-bytes 日本語
     UTF8 : e697a5e69cace8aa9e
    EUCJP : c6fccbdcb8ec
ISO2022JP : 1b2442467c4b5c386c1b2842
 ShiftJIS : 93fa967b8cea
$ echo -n '日本語' | od -t x1 | cut -c 10- | tr -d ' ' | tr -d '\n'
e697a5e69cace8aa9e
```
