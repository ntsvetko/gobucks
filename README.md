# gobucks
Natalie and Vlad learn Go, inspired by everyone's favorite slackbot, mongobucks :).

## How To Use
* make sure that you have an instance of mongo running
* install dependencies
  * run `go get github.com/tools/godep`
  * run `godep restore`
* `go build`
* choose which way you'd like to run it! (REPL Mode or File Mode)

### REPL Mode

`./gobucks <username>`

This gives you a REPL where you can gamble to your heart's content! Unfortunately, none of the gambles in this program actually mean something, but use your imagination.

Give it commands of `gamble <number>` or `balance`.

The output of `./gobucks natalie` looks something like this:

```
Connected to mongo
natalie's balance: 100
natalie> gamble 10
natalie has lost 10 and now has 90! :(
natalie> gamble 20
natalie has lost 20 and now has 70! :(
natalie> gamble 10
natalie has lost 10 and now has 60! :(
natalie> gamble 10
natalie has won 10 and now has 70! :D
natalie> gamble 20
natalie has lost 20 and now has 50! :(
```
Apparently, Natalie shouldn't gamble.

### File Mode

`./gobucks -c <filepath>`

This mode lets you enter a file with commands for multiple users pre-defined. Essentially, this allows gambles from different users to occur concurrently, but if the same user has multiple gambles in a row, the program waits for them to finish their gamble before moving on to the next one. This makes sense, because you need to know how much money you have to gamble again.

The file should look something like:

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

Output will look something like:

```
Connected to mongo
read command: gamble 10 nat
read command: gamble 20 vlad
read command: gamble 30 nat
read command: gamble 10 justin
read command: gamble 12 vlad
read command: gamble 10 nat
read command: gamble 10 nat
read command: gamble 10 nat
nat has lost 10 and now has 90! :(
justin has lost 10 and now has 90! :(
vlad has lost 20 and now has 80! :(
vlad has lost 12 and now has 68! :(
nat has won 30 and now has 120! :D
nat has won 10 and now has 130! :D
nat has lost 10 and now has 120! :(
nat has lost 10 and now has 110! :(
```

There's a slight delay (because gratification can't be *too* instant), and if you run it in this mode, you can see the commands being read in, spawning goroutines (which all complete at different times), and waiting for the previous gamble of the same user to finish before moving on to the next one.

## Uses
This literally has no uses but have fun :D.
