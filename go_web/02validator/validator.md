## 参数校验库

使用时，先new一个校验器
对结构体的校验，使用validator的Struct()方法，这个方法可以检验字段是否符合定义的约束

```go
package main

import (
  "fmt"

  "gopkg.in/go-playground/validator.v10"
)

type User struct {
  Name string `validate:"min=6,max=10"`
  Age  int    `validate:"min=1,max=100"`
}

func main() {
  validate := validator.New()

  u1 := User{Name: "lidajun", Age: 18}
  err := validate.Struct(u1)
  fmt.Println(err)

  u2 := User{Name: "dj", Age: 101}
  err = validate.Struct(u2)
  fmt.Println(err)
}
```

这里Name的长度为[6,10],age的范围在1-100之间
字符串长度和数值的范围都可以通过min和max约束

validator有丰富的约束

我们上面已经看到了使用min和max来约束字符串的长度或数值的范围，下面再介绍其它的范围约束。范围约束的字段类型有以下几种：

对于数值，则约束其值；
对于字符串，则约束其长度；
对于切片、数组和map，则约束其长度。
下面如未特殊说明，则是根据上面各个类型对应的值与参数值比较。

len：等于参数值，例如len=10；
max：小于等于参数值，例如max=10；
min：大于等于参数值，例如min=10；
eq：等于参数值，注意与len不同。对于字符串，eq约束字符串本身的值，而len约束字符串长度。例如eq=10；
ne：不等于参数值，例如ne=10；
gt：大于参数值，例如gt=10；
gte：大于等于参数值，例如gte=10；
lt：小于参数值，例如lt=10；
lte：小于等于参数值，例如lte=10；
oneof：只能是列举出的值其中一个，这些值必须是数值或字符串，以空格分隔，如果字符串中有空格，将字符串用单引号包围，例如oneof=red green。

**跨字段约束**

eqfiled可以约束同一结构体的字段
```go
type RegisterForm struct {
  Name      string `validate:"min=2"`
  Age       int    `validate:"min=18"`
  Password  string `validate:"min=10"`
  Password2 string `validate:"eqfield=Password"`
}
```

**字符串约束**

validator中关于字符串的约束有很多，这里介绍几个：

contains=：包含参数子串，例如contains=email；
containsany：包含参数中任意的 UNICODE 字符，例如containsany=abcd；
containsrune：包含参数表示的 rune 字符，例如containsrune=☻；
excludes：不包含参数子串，例如excludes=email；
excludesall：不包含参数中任意的 UNICODE 字符，例如excludesall=abcd；
excludesrune：不包含参数表示的 rune 字符，excludesrune=☻；
startswith：以参数子串为前缀，例如startswith=hello；
endswith：以参数子串为后缀，例如endswith=bye

```go
type User struct {
  Name string `validate:"containsrune=☻"`
  Age  int    `validate:"min=18"`
}
```

**唯一性约束**
使用unqiue来指定唯一性约束，对不同类型的处理如下：

对于数组和切片，unique约束没有重复的元素；
对于map，unique约束没有重复的值；
对于元素类型为结构体的切片，unique约束结构体对象的某个字段不重复，通过unqiue=field指定这个字段名。

```go
type User struct {
  Name    string   `validate:"min=2"`
  Age     int      `validate:"min=18"`
  Hobbies []string `validate:"unique"`
  Friends []User   `validate:"unique=Name"`
}

```

**email** 限制必须是邮件格式
```go
type User struct {
  Name  string `validate:"min=2"`
  Age   int    `validate:"min=18"`
  Email string `validate:"email"`
}
```

**特殊约束**

有一些比较特殊的约束
-：跳过该字段，不检验；
|：使用多个约束，只需要满足其中一个，例如rgb|rgba；
required：字段必须设置，不能为默认值；
omitempty：如果字段未设置，则忽略它。

**VarWithValue 方法**

有时候我们仅仅只是希望检验两个变量，并不想定义一个结构体，我们可以使用VarWithValue()方法

```go
func main() {
  name1 := "dj"
  name2 := "dj2"

  validate := validator.New()
  fmt.Println(validate.VarWithValue(name1, name2, "eqfield"))

  fmt.Println(validate.VarWithValue(name1, name2, "nefield"))
}
```

**自定义约束**

除了使用validator提供的约束外，还可以定义自己的约束。例如现在有个奇葩的需求，产品同学要求用户必须使用回文串作为用户名，我们可以自定义这个约束：

```go
type RegisterForm struct {
  Name string `validate:"palindrome"`
  Age  int    `validate:"min=18"`
}

func reverseString(s string) string {
  runes := []rune(s)
  for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
    runes[from], runes[to] = runes[to], runes[from]
  }

  return string(runes)
}

func CheckPalindrome(fl validator.FieldLevel) bool {
  value := fl.Field().String()
  return value == reverseString(value)
}

func main() {
  validate := validator.New()
  validate.RegisterValidation("palindrome", CheckPalindrome)

  f1 := RegisterForm{
    Name: "djd",
    Age:  18,
  }
  err := validate.Struct(f1)
  if err != nil {
    fmt.Println(err)
  }

  f2 := RegisterForm{
    Name: "dj",
    Age:  18,
  }
  err = validate.Struct(f2)
  if err != nil {
    fmt.Println(err)
  }
}
```

**错误处理**

在上面的例子中，校验失败时我们仅仅只是输出返回的错误。其实，我们可以进行更精准的处理。validator返回的错误实际上只有两种，一种是参数错误，一种是校验错误。参数错误时，返回InvalidValidationError类型；校验错误时返回ValidationErrors，它们都实现了error接口。而且ValidationErrors是一个错误切片，它保存了每个字段违反的每个约束信息：

```go
type InvalidValidationError struct {
  Type reflect.Type
}

// Error returns InvalidValidationError message
func (e *InvalidValidationError) Error() string {
  if e.Type == nil {
    return "validator: (nil)"
  }

  return "validator: (nil " + e.Type.String() + ")"
}

type ValidationErrors []FieldError

func (ve ValidationErrors) Error() string {
  buff := bytes.NewBufferString("")
  var fe *fieldError

  for i := 0; i < len(ve); i++ {
    fe = ve[i].(*fieldError)
    buff.WriteString(fe.Error())
    buff.WriteString("\n")
  }
  return strings.TrimSpace(buff.String())
}
```

所以validator校验返回的结果只有 3 种情况：

nil：没有错误；
InvalidValidationError：输入参数错误；
ValidationErrors：字段违反约束。
我们可以在程序中判断err != nil时，依次将err转换为InvalidValidationError和ValidationErrors以获取更详细的信息：

```go
func processErr(err error) {
  if err == nil {
    return
  }

  invalid, ok := err.(*validator.InvalidValidationError)
  if ok {
    fmt.Println("param error:", invalid)
    return
  }

  validationErrs := err.(validator.ValidationErrors)
  for _, validationErr := range validationErrs {
    fmt.Println(validationErr)
  }
}

func main() {
  validate := validator.New()

  err := validate.Struct(1)
  processErr(err)

  err = validate.VarWithValue(1, 2, "eqfield")
  processErr(err)
}
```