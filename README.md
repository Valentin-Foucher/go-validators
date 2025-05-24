# Go-Validators

Go-Validators is a library written in Go allowing to validate any field of a struct.

## Usage

### Basic

```go
// define a public object to validate
type ObjectToBeValidated struct {
    First string
    Second int
    Third *bool
}

// validators definition function
// the object passed as parameter needs to be a pointer
func getValidators(object *ObjectToBeValidated) []validators.Validator {
	return []validators.Validator{
        validators.String(object, "first").
            Default("fallback").
            OneOf("fallback", "valid"),
        validators.Int(object, "second").
            Gt(1).
            Lte(10),
        validators.Pointer(object, "third", func(field, value bool) validators.Validator {
            return validators.BoolFromValue(field, value).Equal("true")
        })}
}

func main() {
    third := true
    o := ObjectToBeValidated{
        First: "any value", // nok
        Second: 3, // ok
        Third: &third // ok
    }

    // validation step
    if err := validators.Validate(getValidators(&o)); err != nil {
        panic(err)
    }

    ...
}
```

### Conditional

It is also possible to create conditional validation statements with `Or` and `And` functions.

You can use them in the following way:
```go
func validate(object *ObjectToBeValidated) []validators.Validator {
	return []validators.Validator{
		validators.Or(
			validators.And(
				validators.String(object, "first").Equal("first"),
                validators.Int(object, "second").Equal(2),
			),
			validators.And(
				validators.String(object, "first").Equal("first"),
                validators.Int(object, "second").Equal(2),
			),
		),
	}
}
```

`And` is simple syntactic sugar that becomes useful when dealing with complex validations. It is not mandatory as it is the default 
validation behavior.

## Rules

There are a few rules that are worth to notice:
- Implementation impose validated to be exposed publicly (must start with a capital letter). This is particularly important when using the `Default` method as it will update the input object and that will only work if it is a pointer
- Parameters in `validators.String` or `validators.Int` or any other validator factories are:
  1. the address of the object to be validated
  2. the name of the field to validate (case insensitive) as a string

Besides, primitive types, there are also validators for time.Time, maps, slices, emails (and pointers).

You will find below the validation methods by data type:

### Strings

- Equal
- NotEqual
- IsNotZeroValue
- IsZeroValue
- OneOf
- Default
- UpdateBeforeValidation
- MatchRegex
- Contains
- DoesNotContain
- StartsWith
- DoesNotStartWith
- EndsWith
- DoesNotEndWith
- MaxSize
- MinSize
- MinMaxSize

### Ints

- Equal
- NotEqual
- IsNotZeroValue
- IsZeroValue
- OneOf
- Default
- UpdateBeforeValidation
- Gt
- Lt
- Gte
- Lte

### Floats

- Equal
- NotEqual
- IsNotZeroValue
- IsZeroValue
- OneOf
- Default
- UpdateBeforeValidation
- Gt
- Lt
- Gte
- Lte

### Bools

- Equal
- NotEqual
- IsNotZeroValue
- IsZeroValue
- OneOf
- Default
- UpdateBeforeValidation

### time.Time

- After
- Before
- Equal
- NotEqual
- IsNotZeroValue
- IsZeroValue
- OneOf
- Default
- UpdateBeforeValidation

### Maps

- ContainsKey
- ContainsValue
- DoesNotContainKey
- DoesNotContainValue
- Default
- UpdateBeforeValidation
- MaxSize
- MinSize
- MinMaxSize

### Slices

- Contains
- DoesNotContain
- Default
- UpdateBeforeValidation
- MaxSize
- MinSize
- MinMaxSize

### Pointers

- IsDefined
- IsNotDefined
- Default
- UpdateBeforeValidation

It is also possible to validate the value addressed by the pointer (see previous example).

### Emails

- Equal
- NotEqual
- IsNotZeroValue
- IsZeroValue
- OneOf
- Default
- UpdateBeforeValidation
- MatchRegex
- Contains
- DoesNotContain
- StartsWith
- DoesNotStartWith
- EndsWith
- DoesNotEndWith
- MaxSize
- MinSize
- MinMaxSize
- IsValid
