# Collector

## What is it?

Collector is a library in Go to handle operations on collections (arrays, slices and maps) in a chainable way and in O(n) complexity, with a readability like this:

```
C(collection).Filter(someLambda).Transform(someOtherLambda).Into(&c)
```

or

```
C(collection).Filter(someLambda).Each(someOtherLambda)
```

## How far along is it?

It is very much still in development, and should not be used in production for now. It is also still very much experimental, so considering Go's limitations over types, it's not even sure we can reach the above syntax.

Help out, and open a pull request!

Also, one of my key motivations here is to learn how to best package and test Go code, so feel free to comment on that.
