# The Go Programming Language

## Quick Notes

1.  Go CLI commands:

    a.  **go**: This will give you a list of the Go commands. Also, if go is not properly installed, this command will not be recognized.
    
    b.  **go env**: This will bring up a list of Go environment variables for inspection.
    
    c.  **go version**: This will display the version of Go currently being used.
    
    d.  **go get**: This command is used to download and install packages and dependencies.
    
    e.  **go run**: This command compiles and runs the Go program. It does not save an executable file.
    
    f.  **go build**: This command takes an executable file, compiles the program, reporting errors, if any, then creates an executable ==in the current folder.== After that, one can simply refer to the file name and hit enter to make it run. In a package folder with no executable file, it will not do anything.
    
    g.  **go install**: This command compiles the code, names the executable after the folder name in which the source code is held, then places the resulting executable in $GOPATH/bin.  Also, for a non executable folder, it will build an archive file in the 
    
    h.  **go clean**: This command removes the executable from the *src* directory.
    
    
2.  Basic print commands are contained in the **fmt** package, which is part of the Go standard library.  Therefore, we will need the following line:

    ```go
    import "fmt"
    ```
    
    a.  **Println( )** will send the arguments to the standard out, followed by a new-line char.
    
    b.  **Printf( )** is similar to the printf function in C. It takes formatted parameters, as follows:
    ```go
    func main() {
        fmt.Printf("This is the number %d. In binary, it is %b.\n", 42, 42)
    }
    ```
3.  Format verbs for Printf() function:

    a.  **%d** for a decimal
    
    b.  **%b** for a binary
    
    c.  **%t** for a boolean
    
    d.  **%x** for a hexadecimal (%X for uppercase letters);
    
    e.  **%#x** for a hexadecimal with preceding "0x" (%#X for preceding 0X);
    
    f.  **%q** for a UTF-8 character, with quotes;
    
    g.  **%p** for a pointer;
    
    h.  **%c** for a rune (character);
    
    i.  **%s** for a string or slice;
    
    j.  **%v** for the value of a variable, in a default format;
    
    k.  **%T** for the type of a variable;
    
    
4.  Common escape characters in formatted strings:

    a.  **\n** a new-line character
    
    b.  **\t** a tab character
<div style="page-break-after: always;"></div>

###  Common Go Terminology:

1.  A **rune** is basically a character (a *char*), with an integer value identifying its Unicode representation.  To represent a rune, we use **single quotes** in Go, as opposed to use of double quotes (" ") or back ticks (\` \`) for strings.

    a.  A rune is an alias for **int32**. So, saying that something is of type *rune* is the same as saying it is of type *int32*.
    
    b.  If we are printing out an integer (*alias*, rune), we can see it as a character by formatting with %c.

2.  A **byte** is an 8-bit number, and is an alias for **uint8**.

3.  A string is simply a collection (*i.e*., a **slice**) of runes, and is notated with double-quotation marks.  So, printing 'A', absent formatting cues, prints out the number 65, whereas printing 'A' will result in an upper-case "A".

4.  A string literal represents a **string constant** obtained from concatenating a sequence fo runes.

    a.  A **raw string literal** is contained between back ticks (\` \`), which may contain any characters, including newlines.  This is good to use for things that need to maintain spacing, such as code snippets.
    
    b.  An **interpreted stirng literal** are character sequences between double quotes (" "), and may contain any characters except newlines and unescaped double quotes.

<div style="page-break-after: always;"></div>

### Packages in Go

1.  Note that for a file to be executable in Go, it must have a **main** function and package.

2.  **Packages** are a means of organizing code in Go. Typically, a directory will have the name of a package, and the files in that directory will begin with the statement:
    ```go
    package [packageName]
    ```
3.  A package does not need to have a *main* function or file. For example, a package could have a number of string helper methods. Those methods could be imported into another file as follows:
    ```go
    import (
        "fmt"
        "[path]/[packageName]"
    )
    
    func main() {
        
    }
    ```
    Note in the above that packages that are part of the Go standard library can be called by name, whereas self-produced or third-party packages must have a route as well as a name.
    
4.  **Methods and variables in one file of a package are available in the other files of the same package. There is no need to import.

5.  A method or variable in a package that begins with an uppercase letter can be accessed outside the package; a method in a package that begins with a lowercase letter cannot be accessed outside the package.

#### Variables

1.  There are two means of creating variables in Go:

    a. the *shorthand* method, which can only be used **inside a function**:
    
    ```go
    a := 10
    b := "golang"
    c := 3.14
    d := true
    ```
    In the above examples, the *shorthand* method is used to create the variables and assign a value to them. The *type* is automatically inferred.
    
    b. using the *var* keyword, which assigns a **zero value** to the variable.  Every type has its own *zero value*.  For example, an int would initialize to 0, a string would initialize to "", boolean is *false*.
    ```go
    var a int
    var b string
    var c float64
    var d bool
    ```
    
    The above would be assigned the values 0, "", 0, and false, respectively.
    
2.  **Initialization** of a variable is when it is declared and assigned in one statement.

3.  Types in Go:

    a.  int
    
    b.  string
    
    c.  float64
    
    d.  bool
    
    e.  
    
4.  **Blank Identifier**:   In Go, every declared variable must be used; if not, it will throw an error. In addition, a variable name can be an underscore, "_", in which case a value can be assigned to it, but cannot be read from it. This is commonly used where a function has multiple return values but we don't really care about one or more of the values. For example, in development we might not care about the error return of an http query, so we can just write something such as:
    ```go
    func main() {
        res, _ := http.Get("https://google.com")    //the error goes to the blank identifier
    }
    ```
#### Constants

1.  Constants are declared with the keyword **const**.  To declare multiple constants at once, use the following syntax:
    ```go
    const (
        PI = 3.14
        Language = "Go"
    )
    ```

2.  Note that scope rules are the same for constants, including the treatment of constants that begin with uppercase or lowercase letters.

3.  Go also has the concept of the **iota**.  The first constant that is assigned the value of "iota" will have a value of 0; the next the value of 1, etc.  So, in the following:
    ```
    const (
        a = iota
        b = iota
        c = iota
    )
    ```
    a will be 0, b will be 1, and c will be 2.
    **However**, the assignment convention only applies to the group of constants enclosed in the assignment parentheses.  A subsequent group of constants would start at 0 again.  Also, we can use a *blank identifier* if we wish to not use the zero.
    
4.  Go constants can be **typed** or **untyped**. For example:
    ```go
    const gcw = "Goodbye, cruel world!"         //an untyped constant
    const gcw string = "Goodbye, cruel world!"  //a typed constant
    ```
    A constant that does not yet have a fixed type is referred to as a **kind**.
    
5.  Go is very stronly typed, and does not automatically recast variables. For example:
    ```go
    var num1 = 50           //will be inferred to be an int by the compiler
    var num2 float64 = 25   //a float
    var num3 = num1 + num2  //will throw an error for mixing types
    ```
6.  Numbers can be recast by preceding them with the type:
    ```go
    var num1 = 50
    var num2 float64 = 25
    var num3 = num1 + int(num2)     //num3 will be 75
    ```
7.  
    
#### Scope in Go

1.  There are four levels of scope in Go:

    a.  **universal**
    
    b.  **package**: If a variable is declared outside of a block (*i.e.*, outside braces), then its scope is the entire package of which it is a part.  Remember, variable names that begin with a capital letter can be exported, but variable names that bein with a lowercase cannot, although they could be in a function that gets exported.
    
    c.  **file**: Imports are only given *file-level scope*.  They do carry over to other files in the package.
    
    d.  **block**: If a variable is declared in a block, meaning within curly braces (most often a func(){}).  In block level scope, there is no hoisting, **so order matters.**

2.  Note also that closures work in Go in a manner very similar to the way they work in javascript. For example:
    ```go
    package main
    
    import "fmt"
    
    func main() {
        x := 42
        fmt.Println(x)
        {
            fmt.Println(x)          //will print x
            y := "a new string"
            fmt.Println(y)          //will print y
        }
        fmt.Println(y)              //error, outside scope of y
    }
    ```

#### Writing Functions in Go

1.  To create a function, begin with the keyword *func*, followed by the name of the function:
    ```go
    func addUp(x int) int {
        return x + 42
    }
    ```
    In the above example, a paramet is passed in by name first, followed by type, in parentheses.  The return type is then listed before the braces enclosing the function.
    
2.  Multiple parameters can be passed in as follows:
    ```go
    func adder(x int, y int) int {
        return x + y
    }
    ```
3.  Functions in Go can be anonymous, just as in javascript. For example:

    ```go
    func main() {                   //name of function is main
        x := 0
        increment := func() int {   //anonymous, assigned to "increment"
            x++
            return x
        }
    }
    ```
    This is referred to as a function (or func) expression.
    
4.  Example of closure in Go:
    ```go
    package main
    
    import "fmt"
    
    func outer() func() int {   //the func() int is a return type
        x := 0
        return func() int {
            x++
            return x
        }
    }
    
    func main() {
        increment := outer()        //returns the outer func
        fmt.Println(increment())    //will print 1
        fmt.Println(increment())    //will print 1
    }
    ```
#### Operators in Go

1.  **+**: used to add (numbers) or concatenate (strings),

2.  **-**: used to subtract (numbers),

3.  **\***: used to multiply (numbers),

4.  **/**: used to divide numbers (no remainder),

5.  **%**: used to get remainder,

6.  **<<**: used to shift bits to the left:

    ```go
    3 << 4w  //take the number three (0b11) and shift four places to the left (i.e., add four zeros to it )
    ```

7.  **>>**: used to shift bits to the right

8.  **!**: used to reverse a boolean

9.  **&&**: logical AND operator; combines two expressions

10. **||**: logical OR operator; combines two expressions

11. **NOTE**: there is not a logical XOR operator in Go.  However, one can use, if X and Y are two boolean expressions, X != Y.

#### Evaluators in Go

1.  **==**: equal to (note that there is not, and no need for, a "===" evaluator")

2.  **!=**: not equal to

3.  **< / >**: less than / greater than

4.  **<= / >=**: less than or equal to / greater than or equal to

#### Memory Access in Go

1.  Go allows referencing of memory addresses. For example:
    ```go
    func main() {
        
        a := 40
        
        fmt.Println("a: ", a)       //will print 40
        fmt.Println("a's address: ", &a)    //will print something like 0xc42000000a298
    }
    ```
    
2.  In Go, we can use **pointers**, like in C.  Pointers are declared with the **\*** operator, as follows:
    ```go
    func main() {
        a := 42
        
        fmt.Println(a)      //prints 42
        fmt.Println(&a)     //prints address as hex
        
        var b *int = &a     //this will throw a warning, as Go will infer the type
                            //this *int is the type "point to an int"
        
    }

3.  The **\*** is also used as a **dereferencer**, when placed before the variable name.  So, the notation *\*b* would indicate, if b is a memory address, the value stored at that address.

4.  To change the value held at an address, we can to something along the following lines:
    ```go
    func main() {
        a := 50
        
        b := &a
        
        *b = 55
        
        fmt.Println(a)      //will print 55
        
    }

5.  The following is a simple example of the interaction of scope and pointers in Go:
    ```go
    func zero (y int) {
        y = 0
    }
    
    func main() {
        x := 5
        zero(x)
        fmt.Println(x)  //This outputs 5.  The assignment to y is ineffective.
    }
    ```
    This would not work, as the main function would call the zero function on the value of x, which is passed to y as a parameter, so we would say "y = 0" in the zero func, which doesn't affect x.  To make this work as intended, we need to pass in the address of x to the zero function, as follows:
    ```go
    func zero(y *int)           //parameter is type *int, a pointer to an int
        *y = 0
    }
    
    func main() {
        x := 5
        zero(&x)                //pass in the value of &x, i.e., the address of x
        fmt.Println(x)          //will print out 0, which has been assigned to the address of x
    }


## Functions

1.  The starting point with functions in Go is that, for a program to exist, there must be a function **main**, which is the entry to the program.  Of course, there can only be a single *main* function in the program. **Main must have no arguments and no return values!**

2.  Functions are declared with the keyword **func**, followed by receivers, followed by the name of the function, parameters contained within parentheses, return types, and a body contained within brackets, as follows:
    ```go
    func [receiver] funcName() [return value] {
        //do stuff
    }

3.  Note that functions are typed by their return value. So, a function that has an *int* return value has a type of *func int*.

4.  In Go, functions can be passed in as arguments, or returned from other functions, *i.e.*, Go has **callback functions**.

4.  Always keep in mind the difference between a **function declaration** and a **function expression**. One big difference is that *a function cannot be declared within another function.* In such a case, we **must** use an anonymous function, which we can assign to a variable.

5.  Go has a keyword **defer**, which is an important Go concept. If preceding a function name, it prevents the function from being run until the end of the main function. It is often used for functions that close up a file that has been opened in the running of the program. Basically, it can be thought of as a LIFO ordering of the deferred functions (there can be multiple), so that non deferred functions run, then deferred functions run in LIFO manner.  For example:
    ```go
    func hello() {
	    fmt.Println("Hello")
    }

    func world() {
        fmt.Println("World")
    }

    func again() {
        fmt.Println("Again")
    }

    func main() {
        defer hello()
        world()
        defer again()
    }                           //will print out "World", "Again", "Hello"
    ```

5.  Functions can have multiple parameters, and if they are of the same type, they can be typed as follows:
    ```go
    func greeting(firstName, lastName string) {
        fmt.Println(firstName, lastName)
    }
    ```
5.  Return values can be named. To do this, instead of merely specifying a return type in the declaration, specify a return name and type (note the necessary parentheses):
    ```go
    function greeting(firstName string, lastName string) (s string) {
        s = fmt.Sprint(firstName, lastName)
        return
    }
    ```
6.  **Big News:** In Go, functions can have multiple return values (from a single return statement).  For example:
    ```go
    func main() {
        var w float32
        var x float32
        var y float32
        var z float32

        num1 := float32(10.0)
        num2 := float32(5)

        w, x, y, z = maker(num1, num2)

        fmt.Println(w, x, y, z)         //prints 5 15 50 2
    }

    func maker(firstNum float32, lastNum float32) (float32, float32, float32, float32) {
	    return firstNum - lastNum, firstNum + lastNum, firstNum * lastNum, firstNum / lastNum
    }

    ```
7.  **Variadic** functions can be used in Go. This is a function that can take an undetermined number of variables. The notation to implement this is a type preceded by "...".  For example:
    ```go
    package main
    import "fmt"
    func main() {
	    n := average(43, 55, 2, 89)
	    fmt.Println(n)
    }

    func average(taco ...float64) float64 {
	    total := 0.0
	    for _, v := range taco {
		    total += v
	    }
	    return total / float64(len(taco))
    }
    ```
    In the case where we would like to submit several items as arguments to a variadic function as a single group of arguments (*e.g.*, a slice), we  should follow the name of the slice with "...". For example:
    ```go
    func main() {
        data := []float64{43, 56, 87, 12, 45, 57}
        n := average(data...)
        fmt.Println(n)
    }
    ```
#### References and Pass By Value

1.  In Go, we say that data is **passed by value**, and not passed by reference, meaning that we pass into a function as arguments actual values.  For example, in the following, we are able to pass the value of the memory address of the variable age, so that we can modify it across scopes:
    ```go
    package main

    import "fmt"

    func main() {
        age := 44
        fmt.Println(&age)           //will be the same address
        changeMe(&age)
        fmt.Println(&age)           //as this
        fmt.Println(age)
    }

    func changeMe(z *int) {
        fmt.Println(z)              //as this
        fmt.Println(*z)
        *z = 22
        fmt.Println(z)              //as this
        fmt.Println(*z)
    }
    ```
    
2.  Note, however, that in Go, *slices*, *maps*, and *channels* are references to data structures and the function used to create them (**make**) returns an initialized value of type T (not *T). One cannot pass the address of a slice, for example, because it is already an address.

## Packages

### fmt

1.  **Scan(a ...interface{})** is a method that takes input from the stdin, and treats newlines as spaces. It returns the number of items successfully scanned. If the number of items is less than the number of arguments, it returns an error.  The parameters are the **addresses** of the variables that are assigned input values.
    ```go
    package main

    import "fmt"

    func main() {
        fmt.Print("Hello, what is your name: ")
	    name := ""
	    fmt.Scan(&name)
	    fmt.Println("Hello, my name is " + name)
    }
    ```

2.  **Sprint(a ...interface{})** is a method that formats its arguments and returns the resulting string. Spaces are added between arguments when neither is a string.
    ```go
    func main() {
	    fmt.Println(greet("Jordan ", 10))       //outputs "Jordan 10"
    }

    func greet(firstName string, lastName int) string {
	    return fmt.Sprint(firstName, lastName)
    }

#### Loop Syntax

1.  Note the following example:
    ```go
    func main() {
        for i := 1; i <= 50; i++ {
            if i%2 == 0 {
                fmt.PrintLn(i);
            }
        }
    }
    ```
    Notice the absence of parentheses in both the *for loop* and the *if conditional* statement.  They can be written, but are removed on compilation.
    
### Program Control Flow

1.  Basic flow is from top down.

2.  Go has a concept called **inlining** which is similar to the javascript concept of **hoisting**.

#### For Loop

1.  First, note that the Go *for loop* combines features of the *do/while* loop in other languages, and there is no *while* or *do/while* loop in Go.

2.  There are three forms of *for* loops in Go. The first is the familiar form, except there is no need for parentheses:
    ```go
    total := 0                      //this will sum integers from 0 to 99
    for i := 0; i < 100; i++ {
        total += i
    }
    ```
    Note the three parts: first, there must be an initialization of the counting variable; second, there is a condition; and third, there is a statement that runs **after** what is in the parentheses occurs.
    
3.  The next form of *for* loop is the *for loop with a condition*, which is more similar to a while loop in javascript:
    ```go
    func main() {
        total := 0
        i := 0
        for total <= 100 {      //the loop continues as long as the statement evaluates to true
            total += i
            i++
        }
    }
    ```
4.  The final form is to just continue running (like a listener):
    ```go
    func main() {
        i:= 0
        for {
            if i > 100 {
                break
            }
            fmt.Println(i)
            i++
        }
    }
    ```
    This loop will continue running until the *break* statement is triggered.
    
5.  A similar syntax is that for iterating over an array. For example, if there is a **slice** named *sf* and we want to add all the integers contained in that slice together, we can write:
    ```go
    for ind, val := range sf {
        total += val
    }
    ```
    In the above, the "range" is a slice, and the two parameters are the index of each element and its value.
    
 6.  Go uses the keywords **break** and **continue** in the same manner as JavaScript.

#### Switch Statements

1.  The syntax for switch statements is a little different in Go than in C or JavaScript. The big difference is that the default behaviour for the cases is that the code ends after running the first applicable case, so there is no need for *break* statments.  For example:
    ```go
    func main() {
	    name := "Jordan"
	    switch name {
	    case "Bob":
            fmt.Println("Hello, Bob")
        case "Molly":
            fmt.Println("Hello, Molly")
        case "CJ":
            fmt.Println("Hello, CJ")
            fallthrough                             //explicit fallthrough statement
        case "Jordan":
            fmt.Println("You are the winner!")
            fallthrough                             //explicit fallthrough statement
        default:
            fmt.Println("default")
	    }
    }
    ```
    
2.  If fallthrough is invoked, the tests are not run on the fallen into cases, the commands are just executed.
    
3.  Multiple cases can be set up:
    ```go
    func main() {
        name := "Jorda"
        switch name {
        case "Bob", "Jordan":
            fmt.Println("Hello, Bob")
        case "Molly", "Jorda":
            fmt.Println("Hello, Molly")
        case "CJ":
            fmt.Println("Hello, CJ")
        default:
            fmt.Println("default")
        }
    }
    ```
4.  A *switch/case* statement can have no evaluator, in which case the *first (only)* case that evaluates to true is run, or the default if it gets that far.

5.  **SwitchOnType** allows for switching based on the type that a variable is.

#### Conditional Statements

1.  **If statments** evaluate a condition, then run code contained in brackets if the condition evaluates to true. The syntax is similar to JavaScript, except that parentheses are not required.
    ```go
    func main() {
        
        if true {
            fmt.Println("This is true!")        //this will run
        }
        
        if false {
            fmt.Println("This is false!")       //this will not run
        }
    }
    ```
:::danger 
Note that Go does not evaluate "truthy" and "falsey" statements as in JavaScript. 
:::

2.  *If* and *switch* statements accept initialization statments, for example:
    ```go
    if err := file.Chmod(0664); err != nil {
        log.Print(err)
        return err
    }
    ```
    This allows the scope of the variable *err* to be only within the brackets of the *if* statement.

3.  Go also has **if/else** statements: The *else* **must be** on the same line as the closing *if* bracket.

4.  Go also has **if/else if/else** statements.
    
### String Methods

1.  **len()**:
    a.  one parameter, the string to evaluate,
    b.  returns the length of the argument string,
    c.  no side effects.
    
2.  

## Data Structures in Go

### Arrays

1. An array is a collection of elements, each identified by an index, which in Go begins at zero (as in most languages).

2.  The length of an array is part of its type, and cannot be changed.  **In Go, arrays are not dynamic, *i.e.*, they cannot change in size.**

3.  An array can be declared as follows:
    ```go
    var myArray [12]string
    ```
    An array **must** have a number in the brackets, which defines its length. This is in contrast to a slice, which is not given an initial size.
    
4.  Arrays are not used a great deal in Go, but they underlie the *slice*, which is used more often in programming.

5.  To determine the length of an array, use the built-in **len** function, for example:
    ```go
    var x [10]string
    fmt.Println(len(x))         //will print out "10"
    ```
    Note that the declaration of the array causes it to be initialized with "zero values", in this case an empty string, in each space in the array.
    
6.  Values can be assigned to, or viewed from, the array directly using the element number:
    ```go
    var x [10]string        //initializes an array of 10 empty strings
    x[0] = "Now"
    x[1] = "is"
    Println(x[0])           //will print "Now"

### Slices

1.  A **slice** is, basically, a list of items.  It is a descriptor for a contiguous section of an underlying array.

2.  The value of an unitialized slice is **nil**.

3.  Like an array, the length of a slice can be determined with the function *len()*.

4.  Once initialized, a slice is *always associated with an underlying array that holds its elements.*

5.  Slices are *dynamic*, and not of static length; in this manner they are different from arrays.

6.  **Capacity** is the property of a slice obtained by adding the length of the slice together with the length of the underlying array beyond the slice. The *capacity* of a slice can be determined by the function **cap()**.

7.  To create a slice, we can use the keyword **make**, which takes the slice type and parameters specifying length and capacity (optional), or use literal notation, for example:
    ```go
    package main

    import "fmt"

    func main() {
        mySlice := []int{1, 3, 5, 7, 9, 11,}    //NOTE: last item requires comma
        fmt.Println(mySlice)
    }
    ```
    Notice that the brackets are empty when it is initialized; *i.e.*, in contrast to arrays, this has no specified length. So, the type []string would be read "a slice of strings", whereas the type [4]stirng is "an array of four strings".
    
8.  We can also create a slice by declaration, as follows:
    ```go
    var student []string
    ```
8.  A slice can be made from another slice (or array), with the following notation:
    ```go
    myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 33, 214}
    mySlice := myArray[5:8]
    fmt.Printf("%T\n", mySlice)         //will print [6, 7, 8]
    ```
9.  Example of a slice of slices:
    ```go
    func main() {
        records := make([][]string, 0)
        //student 1's records
        student1 := make([]string, 4)
        student1[0] = "Jordan"
        student1[1] = "Ball"
        student1[2] = "100%"
        student1[3] = "80%"
        records = append(records, student1)
    }
    ```
    Note that we must use the append method for records in the above example, because it is set up with no length, so we cannot assign Student1 to records[0].
    
9.  **Making a New Slice**: By using the *make* keyword, we will also allocate a new, hidden array to which the returned slice value refers.  Thus, the following two produce the same result:
    ```go
    make([]int, 50, 100)    //a slice of ints, with length of 50 and capacity of 100
    
    new([100]int)[0:50]     //taking a new array of 100 ints and slicing the first 50
    ```
    If only one number is included, it is both the length and capacity.
    
10. If the slice exceeds the capacity of the underlying array, the computer creates a new array of twice the size, transfers the elements to the new array, tosses the old array, and updates the reference of the slice to the new array.

11. To add an element to a slice, we use the function **append**, which takes two+ parameters: the name of the slice and the values to add.  It gets added to the end of the slice, and returns the slice, as increased.
    ```go
    mySlice = append(mySlice, 50, 1)  //add 50 and 1 to the end of the slice
    
    mySlice = append(mySlice, otherSlice...)  //appends the elements of the 2nd slice.
    ```
    Notice that the second example above would not work without the ... appended because of the type differences.

12. To remove items from a slice, use *append* and leave out unwanted items:
    ```go
    mySlice = append(mySlice[:2], mySlice[3:]...)
    ```
12. To designate slices of a slice, we use the following notation: **[x:y]**, where x denotes where to begin the slice (inclusive) and y denotes where to end the slice (exclusive).  If no x entry, then start at 0; if no y entry, then start at the x index and go until the end (inclusive). If neither x nor y, then the slice is all of the initial slice.

13. A slice is a **reference type** (along with *maps* and *channels*).  It is defined by three things: a pointer to its location, its length, and its capacity.

### Maps

1.  A key/value storage structure

### Structs