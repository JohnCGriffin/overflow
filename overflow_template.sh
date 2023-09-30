#!/bin/bash

exec > overflow_impl.go

echo "package overflow

import \"math\"

// This is generated code, created by overflow_template.sh executed
// by \"go generate\"

"


for SIZE in 8 16 32 64
do
echo "

// Abs${SIZE} performs absolute value operation on an int${SIZE} operand
// returning a result and a ok result indicating whether the operation is safe.
func Abs${SIZE}(x int${SIZE}) (int${SIZE}, bool) {
        if x == math.MinInt${SIZE} {
                return x, false
        }
        if x >= 0 {
                return x, true
        }
        return -x, true
}

// Add${SIZE} performs + operation on two int${SIZE} operands
// returning a result and a ok result indicating whether the operation is safe.
func Add${SIZE}(a, b int${SIZE}) (int${SIZE}, bool) {
        c := a + b
        if (c > a) == (b > 0) {
                return c, true
        }
        return c, false
}

// Add${SIZE}p is the unchecked panicing version of Add${SIZE}
func Add${SIZE}p(a, b int${SIZE}) int${SIZE} {
        r, ok := Add${SIZE}(a, b)
        if !ok {
                panic(\"addition overflow\")
        }
        return r
}

// UAdd${SIZE} performs + operation on two uint${SIZE} operands
// returning a result and a ok result indicating whether the operation is safe.
func UAdd${SIZE}(a, b uint${SIZE}) (uint${SIZE}, bool) {
        c := a + b
        if c >= a {
                return c, true
        }
        return c, false
}

// UAdd${SIZE}p is the unchecked panicing version of UAdd${SIZE}
func UAdd${SIZE}p(a, b uint${SIZE}) uint${SIZE} {
        r, ok := UAdd${SIZE}(a, b)
        if !ok {
                panic(\"addition overflow\")
        }
        return r
}


// Sub${SIZE} performs - operation on two int${SIZE} operands
// returning a result and a ok result indicating whether the operation is safe.
func Sub${SIZE}(a, b int${SIZE}) (int${SIZE}, bool) {
        c := a - b
        if (c < a) == (b > 0) {
                return c, true
        }
        return c, false
}

// Sub${SIZE}p is the unchecked panicing version of Sub${SIZE}
func Sub${SIZE}p(a, b int${SIZE}) int${SIZE} {
        r, ok := Sub${SIZE}(a, b)
        if !ok {
                panic(\"subtraction overflow\")
        }
        return r
}

// USub${SIZE} performs - operation on two uint${SIZE} operands
// returning a result and a ok result indicating whether the operation is safe.
func USub${SIZE}(a, b uint${SIZE}) (uint${SIZE}, bool) {
        c := a - b
        if a >= b {
                return c, true
        }
        return c, false
}

// USub${SIZE}p is the unchecked panicing version of USub${SIZE}
func USub${SIZE}p(a, b uint${SIZE}) uint${SIZE} {
        r, ok := USub${SIZE}(a, b)
        if !ok {
                panic(\"subtraction overflow\")
        }
        return r
}


// Mul${SIZE} performs * operation on two int${SIZE} operands
// returning a result and a ok result indicating whether the operation is safe.
func Mul${SIZE}(a, b int${SIZE}) (int${SIZE}, bool) {
        if a == 0 || b == 0 {
                return 0, true
        }
        c := a * b
        if (c < 0) == ((a < 0) != (b < 0)) {
                if c/b == a {
                        return c, true
                }
        }
        return c, false
}

// Mul${SIZE}p is the unchecked panicing version of Mul${SIZE}
func Mul${SIZE}p(a, b int${SIZE}) int${SIZE} {
        r, ok := Mul${SIZE}(a, b)
        if !ok {
                panic(\"multiplication overflow\")
        }
        return r
}

// UMul${SIZE} performs * operation on two uint${SIZE} operands
// returning a result and a ok result indicating whether the operation is safe.
func UMul${SIZE}(a, b uint${SIZE}) (uint${SIZE}, bool) {
        if a == 0 || b == 0 {
                return 0, true
        }
        c := a * b
        if c/b == a {
                return c, true
        }
        return c, false
}

// UMul${SIZE}p is the unchecked panicing version of UMul${SIZE}
func UMul${SIZE}p(a, b uint${SIZE}) uint${SIZE} {
        r, ok := UMul${SIZE}(a, b)
        if !ok {
                panic(\"multiplication overflow\")
        }
        return r
}


// Div${SIZE} performs / operation on two int${SIZE} operands
// returning a result and a ok result indicating whether the operation is safe.
func Div${SIZE}(a, b int${SIZE}) (int${SIZE}, bool) {
        q, _, ok := Quotient${SIZE}(a, b)
        return q, ok
}

// Div${SIZE}p is the unchecked panicing version of Div${SIZE}
func Div${SIZE}p(a, b int${SIZE}) int${SIZE} {
        r, ok := Div${SIZE}(a, b)
        if !ok {
                panic(\"division failure\")
        }
        return r
}

// Quotient${SIZE} performs + operation on two int${SIZE} operands
// returning a quotient, a remainder and a ok result indicating whether the operation is safe.
func Quotient${SIZE}(a, b int${SIZE}) (int${SIZE}, int${SIZE}, bool) {
        if b == 0 {
                return 0, 0, false
        }
        c := a / b
        status := (c < 0) == ((a < 0) != (b < 0))
        return c, a % b, status
}

// UDiv${SIZE} performs / operation on two uint${SIZE} operands
// returning a result and a ok result indicating whether the operation is safe.
func UDiv${SIZE}(a, b uint${SIZE}) (uint${SIZE}, bool) {
        q, _, ok := UQuotient${SIZE}(a, b)
        return q, ok
}

// UDiv${SIZE}p is the unchecked panicing version of UDiv${SIZE}
func UDiv${SIZE}p(a, b uint${SIZE}) uint${SIZE} {
        r, ok := UDiv${SIZE}(a, b)
        if !ok {
                panic(\"division failure\")
        }
        return r
}

// UQuotient${SIZE} performs + operation on two uint${SIZE} operands
// returning a quotient, a remainder and a ok result indicating whether the operation is safe.
func UQuotient${SIZE}(a, b uint${SIZE}) (uint${SIZE}, uint${SIZE}, bool) {
        if b == 0 {
                return 0, 0, false
        }
        c := a / b
        return c, a % b, true
}
"
done

SIZE=(64 32 16 8)

# Int -> Int
for FROM in {0..3}; do
for TO in $(seq $((FROM + 1)) 3); do
echo "

// Int${SIZE[FROM]}ToInt${SIZE[TO]} converts an int${SIZE[FROM]} value to int${SIZE[TO]}.
// returning a converted value and a ok result indicating whether the operation is safe.
func Int${SIZE[FROM]}ToInt${SIZE[TO]}(x int${SIZE[FROM]}) (int${SIZE[TO]}, bool) {
        y := int${SIZE[TO]}(x)
        if math.MinInt${SIZE[TO]} <= x && x <= math.MaxInt${SIZE[TO]} {
                return y, true
        }
        return y, false
}

"
done
done

# Uint -> Uint
for FROM in {0..3}; do
for TO in $(seq $((FROM + 1)) 3); do
echo "

// Uint${SIZE[FROM]}ToUint${SIZE[TO]} converts an uint${SIZE[FROM]} value to uint${SIZE[TO]}.
// returning a converted value and a ok result indicating whether the operation is safe.
func Uint${SIZE[FROM]}ToUint${SIZE[TO]}(x uint${SIZE[FROM]}) (uint${SIZE[TO]}, bool) {
        y := uint${SIZE[TO]}(x)
        if x <= math.MaxUint${SIZE[TO]} {
                return y, true
        }
        return y, false
}

"
done
done

# Uint -> Int
for FROM in {0..3}; do
for TO in $(seq $((FROM)) 3); do
echo "

// Uint${SIZE[FROM]}ToInt${SIZE[TO]} converts an uint${SIZE[FROM]} value to int${SIZE[TO]}.
// returning a converted value and a ok result indicating whether the operation is safe.
func Uint${SIZE[FROM]}ToInt${SIZE[TO]}(x uint${SIZE[FROM]}) (int${SIZE[TO]}, bool) {
        y := int${SIZE[TO]}(x)
        if x <= math.MaxInt${SIZE[TO]} {
                return y, true
        }
        return y, false
}

"
done
done


# Int -> Uint
for FROM in {0..3}; do
for TO in {0..3}; do
echo "

// Int${SIZE[FROM]}ToUint${SIZE[TO]} converts an int${SIZE[FROM]} value to uint${SIZE[TO]}.
// returning a converted value and a ok result indicating whether the operation is safe.
func Int${SIZE[FROM]}ToUint${SIZE[TO]}(x int${SIZE[FROM]}) (uint${SIZE[TO]}, bool) {
        y := uint${SIZE[TO]}(x)"
    if [ "$((SIZE[FROM]))" -gt "$((SIZE[TO]))" ]; then
        echo "\
        if 0 <= x && x <= math.MaxUint${SIZE[TO]} {"
    else  
        echo "\
        if 0 <= x {"
    fi
echo "\
                return y, true
        }
        return y, false
}

"
done
done
