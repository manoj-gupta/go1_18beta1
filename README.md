# go 1.18 beta1 features

## Fuzzing

* Run test **without** Fuzzing
```
go test -v  -run=FuzzReverse ./fuzz
```

* Run test **with** Fuzzing
```
go test -v  -run=FuzzReverse -fuzz=Fuzz ./fuzz
```

* Run specific failed testcase with fuzz
```
go test -run=FuzzReverse/0a91bb34ed9b3da68c06956ada477cf71a0e51bd4427316366eb685f6c3024f9 ./fuzz
```

* Fuzz test with timeout
```
go test -fuzz=Fuzz -fuzztime 30s ./fuzz
```

## Tutorial Links
* [Generics](https://go.dev/doc/tutorial/generics)
* [Fuzzing](https://go.dev/doc/tutorial/fuzz)
