package overflow

// This is generated code, created by overflow_template.sh executed
// by "go generate"




// Add8 performs + operation on two int8 operands
// returning a result and status
func Add8(a, b int8) (int8, bool) {
        c := a + b
        if (c > a) == (b > 0) {
                return c, true
        }
        return c, false
}

// Add8p is the unchecked panicing version of Add8
func Add8p(a, b int8) int8 {
        r, ok := Add8(a, b)
        if !ok {
                panic("addition overflow")
        }
        return r
}

// UAdd8 performs + operation on two uint8 operands
// returning a result and status
func UAdd8(a, b uint8) (uint8, bool) {
        c := a + b
        if c >= a {
                return c, true
        }
        return c, false
}

// UAdd8p is the unchecked panicing version of UAdd8
func UAdd8p(a, b uint8) uint8 {
        r, ok := UAdd8(a, b)
        if !ok {
                panic("addition overflow")
        }
        return r
}


// Sub8 performs - operation on two int8 operands
// returning a result and status
func Sub8(a, b int8) (int8, bool) {
        c := a - b
        if (c < a) == (b > 0) {
                return c, true
        }
        return c, false
}

// Sub8p is the unchecked panicing version of Sub8
func Sub8p(a, b int8) int8 {
        r, ok := Sub8(a, b)
        if !ok {
                panic("subtraction overflow")
        }
        return r
}

// USub8 performs - operation on two uint8 operands
// returning a result and status
func USub8(a, b uint8) (uint8, bool) {
        c := a - b
        if a >= b {
                return c, true
        }
        return c, false
}

// USub8p is the unchecked panicing version of USub8
func USub8p(a, b uint8) uint8 {
        r, ok := USub8(a, b)
        if !ok {
                panic("subtraction overflow")
        }
        return r
}


// Mul8 performs * operation on two int8 operands
// returning a result and status
func Mul8(a, b int8) (int8, bool) {
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

// Mul8p is the unchecked panicing version of Mul8
func Mul8p(a, b int8) int8 {
        r, ok := Mul8(a, b)
        if !ok {
                panic("multiplication overflow")
        }
        return r
}

// UMul8 performs * operation on two uint8 operands
// returning a result and status
func UMul8(a, b uint8) (uint8, bool) {
        if a == 0 || b == 0 {
                return 0, true
        }
        c := a * b
        if c/b == a {
                return c, true
        }
        return c, false
}

// UMul8p is the unchecked panicing version of UMul8
func UMul8p(a, b uint8) uint8 {
        r, ok := UMul8(a, b)
        if !ok {
                panic("multiplication overflow")
        }
        return r
}


// Div8 performs / operation on two int8 operands
// returning a result and status
func Div8(a, b int8) (int8, bool) {
        q, _, ok := Quotient8(a, b)
        return q, ok
}

// Div8p is the unchecked panicing version of Div8
func Div8p(a, b int8) int8 {
        r, ok := Div8(a, b)
        if !ok {
                panic("division failure")
        }
        return r
}

// Quotient8 performs + operation on two int8 operands
// returning a quotient, a remainder and status
func Quotient8(a, b int8) (int8, int8, bool) {
        if b == 0 {
                return 0, 0, false
        }
        c := a / b
        status := (c < 0) == ((a < 0) != (b < 0))
        return c, a % b, status
}

// UDiv8 performs / operation on two uint8 operands
// returning a result and status
func UDiv8(a, b uint8) (uint8, bool) {
        q, _, ok := UQuotient8(a, b)
        return q, ok
}

// UDiv8p is the unchecked panicing version of UDiv8
func UDiv8p(a, b uint8) uint8 {
        r, ok := UDiv8(a, b)
        if !ok {
                panic("division failure")
        }
        return r
}

// UQuotient8 performs + operation on two uint8 operands
// returning a quotient, a remainder and status
func UQuotient8(a, b uint8) (uint8, uint8, bool) {
        if b == 0 {
                return 0, 0, false
        }
        c := a / b
        return c, a % b, true
}



// Add16 performs + operation on two int16 operands
// returning a result and status
func Add16(a, b int16) (int16, bool) {
        c := a + b
        if (c > a) == (b > 0) {
                return c, true
        }
        return c, false
}

// Add16p is the unchecked panicing version of Add16
func Add16p(a, b int16) int16 {
        r, ok := Add16(a, b)
        if !ok {
                panic("addition overflow")
        }
        return r
}

// UAdd16 performs + operation on two uint16 operands
// returning a result and status
func UAdd16(a, b uint16) (uint16, bool) {
        c := a + b
        if c >= a {
                return c, true
        }
        return c, false
}

// UAdd16p is the unchecked panicing version of UAdd16
func UAdd16p(a, b uint16) uint16 {
        r, ok := UAdd16(a, b)
        if !ok {
                panic("addition overflow")
        }
        return r
}


// Sub16 performs - operation on two int16 operands
// returning a result and status
func Sub16(a, b int16) (int16, bool) {
        c := a - b
        if (c < a) == (b > 0) {
                return c, true
        }
        return c, false
}

// Sub16p is the unchecked panicing version of Sub16
func Sub16p(a, b int16) int16 {
        r, ok := Sub16(a, b)
        if !ok {
                panic("subtraction overflow")
        }
        return r
}

// USub16 performs - operation on two uint16 operands
// returning a result and status
func USub16(a, b uint16) (uint16, bool) {
        c := a - b
        if a >= b {
                return c, true
        }
        return c, false
}

// USub16p is the unchecked panicing version of USub16
func USub16p(a, b uint16) uint16 {
        r, ok := USub16(a, b)
        if !ok {
                panic("subtraction overflow")
        }
        return r
}


// Mul16 performs * operation on two int16 operands
// returning a result and status
func Mul16(a, b int16) (int16, bool) {
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

// Mul16p is the unchecked panicing version of Mul16
func Mul16p(a, b int16) int16 {
        r, ok := Mul16(a, b)
        if !ok {
                panic("multiplication overflow")
        }
        return r
}

// UMul16 performs * operation on two uint16 operands
// returning a result and status
func UMul16(a, b uint16) (uint16, bool) {
        if a == 0 || b == 0 {
                return 0, true
        }
        c := a * b
        if c/b == a {
                return c, true
        }
        return c, false
}

// UMul16p is the unchecked panicing version of UMul16
func UMul16p(a, b uint16) uint16 {
        r, ok := UMul16(a, b)
        if !ok {
                panic("multiplication overflow")
        }
        return r
}


// Div16 performs / operation on two int16 operands
// returning a result and status
func Div16(a, b int16) (int16, bool) {
        q, _, ok := Quotient16(a, b)
        return q, ok
}

// Div16p is the unchecked panicing version of Div16
func Div16p(a, b int16) int16 {
        r, ok := Div16(a, b)
        if !ok {
                panic("division failure")
        }
        return r
}

// Quotient16 performs + operation on two int16 operands
// returning a quotient, a remainder and status
func Quotient16(a, b int16) (int16, int16, bool) {
        if b == 0 {
                return 0, 0, false
        }
        c := a / b
        status := (c < 0) == ((a < 0) != (b < 0))
        return c, a % b, status
}

// UDiv16 performs / operation on two uint16 operands
// returning a result and status
func UDiv16(a, b uint16) (uint16, bool) {
        q, _, ok := UQuotient16(a, b)
        return q, ok
}

// UDiv16p is the unchecked panicing version of UDiv16
func UDiv16p(a, b uint16) uint16 {
        r, ok := UDiv16(a, b)
        if !ok {
                panic("division failure")
        }
        return r
}

// UQuotient16 performs + operation on two uint16 operands
// returning a quotient, a remainder and status
func UQuotient16(a, b uint16) (uint16, uint16, bool) {
        if b == 0 {
                return 0, 0, false
        }
        c := a / b
        return c, a % b, true
}



// Add32 performs + operation on two int32 operands
// returning a result and status
func Add32(a, b int32) (int32, bool) {
        c := a + b
        if (c > a) == (b > 0) {
                return c, true
        }
        return c, false
}

// Add32p is the unchecked panicing version of Add32
func Add32p(a, b int32) int32 {
        r, ok := Add32(a, b)
        if !ok {
                panic("addition overflow")
        }
        return r
}

// UAdd32 performs + operation on two uint32 operands
// returning a result and status
func UAdd32(a, b uint32) (uint32, bool) {
        c := a + b
        if c >= a {
                return c, true
        }
        return c, false
}

// UAdd32p is the unchecked panicing version of UAdd32
func UAdd32p(a, b uint32) uint32 {
        r, ok := UAdd32(a, b)
        if !ok {
                panic("addition overflow")
        }
        return r
}


// Sub32 performs - operation on two int32 operands
// returning a result and status
func Sub32(a, b int32) (int32, bool) {
        c := a - b
        if (c < a) == (b > 0) {
                return c, true
        }
        return c, false
}

// Sub32p is the unchecked panicing version of Sub32
func Sub32p(a, b int32) int32 {
        r, ok := Sub32(a, b)
        if !ok {
                panic("subtraction overflow")
        }
        return r
}

// USub32 performs - operation on two uint32 operands
// returning a result and status
func USub32(a, b uint32) (uint32, bool) {
        c := a - b
        if a >= b {
                return c, true
        }
        return c, false
}

// USub32p is the unchecked panicing version of USub32
func USub32p(a, b uint32) uint32 {
        r, ok := USub32(a, b)
        if !ok {
                panic("subtraction overflow")
        }
        return r
}


// Mul32 performs * operation on two int32 operands
// returning a result and status
func Mul32(a, b int32) (int32, bool) {
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

// Mul32p is the unchecked panicing version of Mul32
func Mul32p(a, b int32) int32 {
        r, ok := Mul32(a, b)
        if !ok {
                panic("multiplication overflow")
        }
        return r
}

// UMul32 performs * operation on two uint32 operands
// returning a result and status
func UMul32(a, b uint32) (uint32, bool) {
        if a == 0 || b == 0 {
                return 0, true
        }
        c := a * b
        if c/b == a {
                return c, true
        }
        return c, false
}

// UMul32p is the unchecked panicing version of UMul32
func UMul32p(a, b uint32) uint32 {
        r, ok := UMul32(a, b)
        if !ok {
                panic("multiplication overflow")
        }
        return r
}


// Div32 performs / operation on two int32 operands
// returning a result and status
func Div32(a, b int32) (int32, bool) {
        q, _, ok := Quotient32(a, b)
        return q, ok
}

// Div32p is the unchecked panicing version of Div32
func Div32p(a, b int32) int32 {
        r, ok := Div32(a, b)
        if !ok {
                panic("division failure")
        }
        return r
}

// Quotient32 performs + operation on two int32 operands
// returning a quotient, a remainder and status
func Quotient32(a, b int32) (int32, int32, bool) {
        if b == 0 {
                return 0, 0, false
        }
        c := a / b
        status := (c < 0) == ((a < 0) != (b < 0))
        return c, a % b, status
}

// UDiv32 performs / operation on two uint32 operands
// returning a result and status
func UDiv32(a, b uint32) (uint32, bool) {
        q, _, ok := UQuotient32(a, b)
        return q, ok
}

// UDiv32p is the unchecked panicing version of UDiv32
func UDiv32p(a, b uint32) uint32 {
        r, ok := UDiv32(a, b)
        if !ok {
                panic("division failure")
        }
        return r
}

// UQuotient32 performs + operation on two uint32 operands
// returning a quotient, a remainder and status
func UQuotient32(a, b uint32) (uint32, uint32, bool) {
        if b == 0 {
                return 0, 0, false
        }
        c := a / b
        return c, a % b, true
}



// Add64 performs + operation on two int64 operands
// returning a result and status
func Add64(a, b int64) (int64, bool) {
        c := a + b
        if (c > a) == (b > 0) {
                return c, true
        }
        return c, false
}

// Add64p is the unchecked panicing version of Add64
func Add64p(a, b int64) int64 {
        r, ok := Add64(a, b)
        if !ok {
                panic("addition overflow")
        }
        return r
}

// UAdd64 performs + operation on two uint64 operands
// returning a result and status
func UAdd64(a, b uint64) (uint64, bool) {
        c := a + b
        if c >= a {
                return c, true
        }
        return c, false
}

// UAdd64p is the unchecked panicing version of UAdd64
func UAdd64p(a, b uint64) uint64 {
        r, ok := UAdd64(a, b)
        if !ok {
                panic("addition overflow")
        }
        return r
}


// Sub64 performs - operation on two int64 operands
// returning a result and status
func Sub64(a, b int64) (int64, bool) {
        c := a - b
        if (c < a) == (b > 0) {
                return c, true
        }
        return c, false
}

// Sub64p is the unchecked panicing version of Sub64
func Sub64p(a, b int64) int64 {
        r, ok := Sub64(a, b)
        if !ok {
                panic("subtraction overflow")
        }
        return r
}

// USub64 performs - operation on two uint64 operands
// returning a result and status
func USub64(a, b uint64) (uint64, bool) {
        c := a - b
        if a >= b {
                return c, true
        }
        return c, false
}

// USub64p is the unchecked panicing version of USub64
func USub64p(a, b uint64) uint64 {
        r, ok := USub64(a, b)
        if !ok {
                panic("subtraction overflow")
        }
        return r
}


// Mul64 performs * operation on two int64 operands
// returning a result and status
func Mul64(a, b int64) (int64, bool) {
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

// Mul64p is the unchecked panicing version of Mul64
func Mul64p(a, b int64) int64 {
        r, ok := Mul64(a, b)
        if !ok {
                panic("multiplication overflow")
        }
        return r
}

// UMul64 performs * operation on two uint64 operands
// returning a result and status
func UMul64(a, b uint64) (uint64, bool) {
        if a == 0 || b == 0 {
                return 0, true
        }
        c := a * b
        if c/b == a {
                return c, true
        }
        return c, false
}

// UMul64p is the unchecked panicing version of UMul64
func UMul64p(a, b uint64) uint64 {
        r, ok := UMul64(a, b)
        if !ok {
                panic("multiplication overflow")
        }
        return r
}


// Div64 performs / operation on two int64 operands
// returning a result and status
func Div64(a, b int64) (int64, bool) {
        q, _, ok := Quotient64(a, b)
        return q, ok
}

// Div64p is the unchecked panicing version of Div64
func Div64p(a, b int64) int64 {
        r, ok := Div64(a, b)
        if !ok {
                panic("division failure")
        }
        return r
}

// Quotient64 performs + operation on two int64 operands
// returning a quotient, a remainder and status
func Quotient64(a, b int64) (int64, int64, bool) {
        if b == 0 {
                return 0, 0, false
        }
        c := a / b
        status := (c < 0) == ((a < 0) != (b < 0))
        return c, a % b, status
}

// UDiv64 performs / operation on two uint64 operands
// returning a result and status
func UDiv64(a, b uint64) (uint64, bool) {
        q, _, ok := UQuotient64(a, b)
        return q, ok
}

// UDiv64p is the unchecked panicing version of UDiv64
func UDiv64p(a, b uint64) uint64 {
        r, ok := UDiv64(a, b)
        if !ok {
                panic("division failure")
        }
        return r
}

// UQuotient64 performs + operation on two uint64 operands
// returning a quotient, a remainder and status
func UQuotient64(a, b uint64) (uint64, uint64, bool) {
        if b == 0 {
                return 0, 0, false
        }
        c := a / b
        return c, a % b, true
}

