# Python
- Easy and intuitive programming language
- Free and Open Source
- Can be widely used for a variety of tasks

## Functions
A part of your code that's used to cause an effect or evaluate a value.

Can come from:
- Python (built-in functions)
- Modules
- Your own code

### Function Execution
1. Check function name
2. Check arguments passed
3. Jumps into the function
4. Executes the function
5. Returns to your code
6. Resumes execution

## Literals
A literal is a value that is written into the code. For example, `print("Hello World")` has the literal `"Hello World"`.

Literal types:
1. Integers
   1. Octal numbers (base 8)
   2. Hexadecimal numbers (base 16)
2. Floating-point numbers
3. Strings
4. Booleans

### Numbers
- Integers: Ex - 123
- Octal: Ex - 0o123
- Hexadecimal: Ex - 0x123
- Floating-point: Ex - 123.45

### Strings
- Single quotes: Ex - 'Hello World'
- Double quotes: Ex - "Hello World"
- Use quotes within strings: Ex - "Hello 'World'"
- Triple quotes: Ex - """Hello World"""
- Escape sequences: Ex - "Hello \"World\""
- Raw strings: Ex - r"Hello \n World"
- Byte strings: Ex - b"Hello World"
- Unicode strings: Ex - u"Hello World"
- Formatted strings: Ex - f"Hello {name}"

### Booleans
- True
- False
- 0
- 1

## Operators
Operators are used to perform operations on variables and values.

### Arithmetic Operators
- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`
- Floor Division: `//`
- Modulus: `%`
- Exponentiation: `**`

### Assignment Operators
- `=`
- `+=`
- `-=`
- `*=`
- `/=`
- `//=`
- `%=`
- `**=`
- `&=`
- `|=`
- `^=`

### Comparison Operators
- Equal: `==`
- Not Equal: `!=`
- Greater Than: `>`
- Less Than: `<`
- Greater Than or Equal To: `>=`
- Less Than or Equal To: `<=`
- Identity: `is`
- Not Identity: `is not`
- Membership: `in`
- Not Membership: `not in`

### Logical Operators
- Boolean AND: `and`
- Boolean OR: `or`
- Boolean NOT: `not`

### Bitwise Operators
- Bitwise AND: `&`
- Bitwise OR: `|`
- Bitwise XOR: `^`
- Bitwise NOT: `~`
- Bitwise Left Shift: `<<`
- Bitwise Right Shift: `>>`
- Bitwise Zero Fill Left Shift: `<<<`
- Bitwise Zero Fill Right Shift: `>>>`
- Bitwise AND Assignment: `&=`
- Bitwise OR Assignment: `|=`
- Bitwise XOR Assignment: `^=`
- Bitwise Left Shift Assignment: `<<=`
- Bitwise Right Shift Assignment: `>>=`
- Bitwise Zero Fill Left Shift Assignment: `<<<=`
- Bitwise Zero Fill Right Shift Assignment: `>>>=`

### Other Operators
- `+` (Unary Plus)
- `-` (Unary Minus)
- `~` (Bitwise NOT)

## Variables
Variables are containers for storing data values. They can be used to store any type of data.

- Variables allow you to store values
- A variable has a valid name (letters, digists, underscore, not a reserved keyword)
- Python is dynamically typed: variables can be reassigned to different data types
- Variables are case-sensitive 
- We can use shortcut operators in order to clearly redeclare a variable

Example code:
```python
amount_of_apples = 2
cost_of_apple = 5
total_cost = amount_of_apples * cost_of_apple
print(total_cost)
```
Output:
```bash
10
```

### Variable Names
- Must start with a letter or underscore
- Can only contain alphanumeric characters and underscores
- Case-sensitive
- Cannot be a keyword
- Cannot start with a number
- Can be any length
- Can contain any type of data
- Can be reassigned
- Can be global or local
- Can be used in functions, classes, and modules
- Can not be named the same as reserved words (keywords) or built-in functions

## Comments
Comments are used to explain code when necessary. They are ignored by the interpreter.

### Single Line Comments
```python
# This is a comment
```

### Multi Line Comments
```python
"""
This is a comment
written in
more than just one line
"""
```
