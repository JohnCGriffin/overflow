# overflow
Check for int/int64/int32 arithmetic overflow in Golang

### Synopsis

```
sum,ok := Add(a,b)
if !ok {
   panic("overflow")
}

```
For int, int64, and int32 types, provide Add, Add32, Add64, Sub, Sub32, Sub64, etc.  
Unsigned types not covered at the moment, but such additions are welcome.

