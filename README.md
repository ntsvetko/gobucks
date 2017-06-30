# gobucks
Natalie and Vlad learn Go.

## How To Use
So this isn't actually user-friendly, it's more like an excuse to encounter as many Go concepts as we can.

* `go build`
* `./GobucksConcurrent <filepath>`

The file should look something like
```
gamble 10 nat
gamble 20 vlad
gamble 30 nat
gamble 10 justin
gamble 12 vlad
gamble 10 nat
gamble 10 nat
gamble 10 nat
```

output will look like:
```
Connected to mongo
command: gamble 10 nat
command: gamble 20 vlad
command: gamble 30 nat
command: gamble 10 justin
command: gamble 12 vlad
command: gamble 10 nat
command: gamble 10 nat
command: gamble 10 nat
justin has won 10 and now has 120!
nat has lost 30 and now has 80! :(
vlad has won 20 and now has 168!
vlad has won 12 and now has 180!
nat has won 10 and now has 90!
nat has won 10 and now has 100!
nat has won 10 and now has 110!
nat has won 10 and now has 120!
```
