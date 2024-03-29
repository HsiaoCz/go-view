## js

解释型的语言
运行在客户端
面向对象语言

三种添加javascript的方法
1、内部的javascript
```html
<script>

</script>
```

2、外部的javascript
```html
<script src="script.js"></script>
```

3、内联的javascript

```html
<button onclick="">
    
</button>
```

### 1、js基础

为HTML添加一些动态的效果
js动态的解释语言：脚本语言，需要有个解释器，按顺序执行，直到执行的时候有保存，才会报错
静态语言：编译型语言，编译的时候会进行类型检查，语言检查，如果出错，就会报错

js是弱类型的语言：与强类型相对（强类型就是变量的类型是明确的）
弱类型的变量类型是不明确的
```javascript
var a ="test"
a=10
console.log(a)
```


javascript的运行时：
- 浏览器
- Nodejs

### 2、数据类型

1、Number

```javascript
//整数和小数都是number
var a=10.1
console.log(typeof(a))

var n=10
console.log(typeof(n))
```
2、字符串

字符串用单引号和双引号都可以表示
js中推荐使用单引号

```javascript
var a='hello'
var m="hi"
console.log(a)

// 字符串的模板
console.log(`你好,${a}`) 
```

3、布尔值


```javascript
var a=true

if(a){
    console.log(`你好,{a}`)
}
```

4、数组

js 的数组包含任意类型

```javascript
var a =[]

//往数组中添加元素

a.push(10)

console.log(a)

// 数组可以这样声明
// 相当于声明一个数组并且往里push了一套参数
var a=Array(1,2,3)

// 数组中存储的元素类型是不固定的
var a=Array(1,2,3,[],{})
console.log(a)

// 数组越界是不报错的

a[0]
a[10] //会报一个undefined

// push和pop可以添加和抛出元素 都是往尾部添加
a.push(1,2)
a.pop() // 从尾部弹出一个元素
b=a.pop()  // pop没有参数

// shift 和unshift
a.unshift(1，2)  //往头部添加两个元素
a.shift()  // 从头部弹出元素

// splice() 是修改数组的万能的方法 可以删除的同时追加元素
// 返回值是删除的元素
m=[1,2,3,4,5,5]
m.splice(2,3,10,11) //从索引为2开始往后删除三个，并追加10，11，追加从删除的索引位置开始

// sort 数组的排序操作

var arr=[100,2,15,30,7,6]
arr.sort() // 排序 按照首字母排序
arr.sort((a,b)=>a-b)
arr.reverse() //按照首字母排序

// concat用于将两个数组连在一起
// slice用于对数组的切片操作

var am=[1,2,3]
var added=am.concat(["Hello","hi"])

am.slice(1,2)  // 从索引为1开始到索引2，不包括2
am.slice(1)  // 从索引1切到最后

```

5、对象

// js 中任何事物都是对象，和python 差不多
```javascript

var ob1=new Object()  // 声明一个空的Object
ob1.name="lao pao"
obi.age=18
console.log(ob1)

// 判断是否有某个属性
console.log(ob1.hasOwnProperty("name"))

```

### 2、逻辑运算符

&& 与运算
|| 或运算
! 非运算

### 3、关系运算符

== :会自动转换数据类型再比较，很多时候，会得到非常诡异的结果
=== ：不会自动转换数据类型，如果数据类型不一致，返回false，如果一致，再比较

### 4、变量

**var声明变量**

使用这种声明会对变量的作用域进行提升，无论声明在何处，都会被提升至其所在作用域的顶部

```javascript
var f1=function(){
    console.log("hello")
    console.log(m1)
}
var m1=10
f1()  // 这里m1会报undefined

```

**let声明变量**

使用let声明变量不会提升到全局
用于声明局部的变量

常量：const

对常量操作时不被允许的



### 5、解构赋值

js里面还有这种你看不懂的操作

```javascript
let [x,[y,z]]=["hello",["hi","hello"]]

var person={
    name:"zhangsan",
    age:18,
    gender:"male",
    school: "hi school",
}

var {name,age}=person
```

### 6、javascript字符串

// 多行字符串 

```javascript
var mi =`
hello,hello,hi
`
console.log(m1)

```

字符串的大小写问题

```javascript
a="alice"
a.toLocaleLowerCase()
a.toUpperCase()
```

### 7、错误处理

**null**
null就是什么都没有

```javascript
a=[]
a.lenth()
s=null
s.lenth() // 这里就会报错
```

如果在一个函数内部发生了错误，它自身没有捕获，错误就会被抛到外层调用函数，如果外层调用函数也没有捕获，该错误会一直沿着调用链一直往上抛出，直到被javascript引擎捕获

try..catch语法捕获异常

```javascript
try{
    var s=null
    s.length
}catch(error){
   console.log(error)
}finally{
    console.log("hello")
}
```

js抛出异常：throw new Error("抛出异常")

**错误类型**

error也是js中的一种类型

```javascript
// 自定义一个error

err=error("hello")
```

### 8、函数

```javascript
function abs(s){
    if (x>=0){
        return x;
    }else{
        return -x;
    }
}
```

- function 用于定义函数
- abs是函数名称
- (x)括号内列出函数的参数，多个参数以,号分割
- {...}之间的代码是函数体

### 9、方法

比如我们定义了一个对象

```javascript
var person={name : "小明",age=18}

// 我们可以为这个对象添加方法
person.greet=function(){
    console.log(`hello,my name is ${this.name}`)
}

person.greet()
```

绑定到对象上的函数称之为方法，和普通的函数也没什么区别，但是它在内部使用了一个this 关键字
在一个方法内部，this是一个特殊变量，它始终指向当前对象

this的坑：
如果this没有绑定在任何对象上，那么this指的是浏览器的window对象

```javascript
var person={name:"zhangsan",age:22}
person.greet=function(){
   let that=this  // 使用that 来指向persion 将that赋值给this
    return function(){
        console.log(`hello,my name is ${this.name}`)
    }
}

person.greet() // 这里会输出undefined
// 因为this指向的是上层，上层为一个什么都没有的函数
// 所以什么都没输出
// 想要避免这种情况，可以使用
```

// arrow function
```javascript
function hello(x,y){return x+y}
console.log(sum(1,2))
```

箭头函数
```javascript
let f3=(x,y)=>{return x+y}
console.log(f3(10,10))
```

箭头函数看上去是匿名函数的一种简写，但实际上，箭头函数和匿名函数有个明显的区别，箭头函数内部的this是词法作用域，由上下文确定

什么叫词法作用域？

```javascript
var persion={name:"张三",age:22}
persion.greet=function(){
    // 这个时候this指向的就是persion
    // =>会帮我们确定上下文的this
    return ()=>{
        console.log(`hello,my name is ${this.name}`)
    }
}

person.greet()()
```

### 10、命名空间

不在任何函数内定义的变量就具有全局作用域，实际上，javaScript默认有一个全局对象window，全局作用域的变量实际上被绑定到window的一个属性

对于浏览器而言默认的全局对象就是window，如果直接在编辑器里执行，默认的全局对象是Object,

可以认为不在花括号内声明的代码就是全局作用变量

**全局作用域与window**

由于函数定义有两种方式，以变量的方式var foo=function()定义的函数实际上也是一个全局变量
因此，顶层函数有多个全局变量，并绑定到window对象

```javascript
var a=10
a // 10
window.a //10
```

甚至我们可以覆盖掉浏览器的内置方法:

```javascript
alert=()=>(console.log("覆盖掉alert方法"))
alert()
```

为了避免不同的程序员将相同的变量放在同一个对象上，所以有一个类似于go package的包管理方式

**export**

在js中一个模块就是一个独立的文件，该文件内部的所以变量，外部无法获取，如果你希望外部能够获取模块内部的某个变量

```javascript
export var firstName="Micheal";
var lastName ="jackson";
var year =1982;

// 通过export 将对象暴露到外部
// 外部通过import导入
export {firstName,lastName,year}

// 或者直接使用export来暴露文件
```

当我们在一个包中暴露了一些变量
可以在另一个包中导入这些变量

```javascript
import {firstName as fName,lastName,age} from './profile.mjs'
console.log(firstName,lastName,age)

// 这里要注意一点，引入的文件后缀得是.mjs
// 被引入得文件的后缀也是.mjs
// as 这这里是一个起别名的用途,主要是为了避免冲突
```

我们可以将所有的变量放到一个对象上
然后将这个对象暴露出去

```javascript
var MYAPP={}

MYAPP.firstName="hello";
MYAPP.LastName="hello";
MYAPP.year=18;

export MYAPP;
```
使用这种暴露好像还会出一些问题

可以这样
```javascript
MYAPP.firstName="hello";
MYAPP.LastName="hello";
MYAPP.year=18;

export var MyApp={firstName,lastName,year}

```

**默认导出**

使用import命令的时候，用户需要知道所要加载的变量名或函数名，否则无法加载成功

比如
```javascript
import (MYAPP) from './export'
```

如果我们想要直接导入模块，再看该模块下有哪些变量可用，就像golang一样
我们可以使用export default 命令

```javascript
export default MYAPP
```

与export 的区别是，其他模块加载此模块时，可以随意起名字
相当于自动帮你完成了一个default as pkg的操作


```javascript
import myapp from './profile'
```

**注意版本**

这里需要注意的是，ES版本不同，有两种导入语法
- require/exports 是Commonjs/AMD中为了解决模块化语法而引入的ES5语法，nodejs使用该规范
- import/export是ES6引入的规范，因为兼容问题，需要在node中用babel将ES6语法编译成ES5语法

commonjs的语法，一个模块想要对外暴露变量，可以使用module.exports=variable,一个模块想要引用其他模块暴露的变量，使用var ref=require(module_name)就拿到了引用模块的变量

这里的机制其实就是为每个文件准备一个module对象，把它装进去

```javascript
var module={
    id :'hello',
    exports:{}
};
```

### 11、条件判断

主要是条件判断的时候多加一个括号

```javascript
var a=10;
if (a<=10){

}eles if(a>=10){

}else{

}
```

### 12 for 循环

主要注意条件要加括号

```javascript
var x=0;
var i;
for (i =1;i<=1000;i++){
    x=x+1;
}

```

for in循环，将一个对象的所有属性取出来

```javascript
var a={
    name:"hello",
    age:18,
};
for(var key in a){
    console.log(key,a[key])
}

var s=[1,2,3,4]
for (let i in s){
    console.log(i)
}
```

for in有一个问题，当动态的为一个map或数组绑定了一个值之后
for in也会遍历出来

于是出现了`for of`;

```javascript
for (let item of a){
    console.log(item)
}
```

for of遍历的是集合中的元素，并不单是key
那么怎么使用这种方式遍历到key

```javascript
for (let key of Object.key(a)){
    console.log(key);
}
```

**for Each**

任何一个array都有一个forEach方法

```javascript
a.forEach(element=>{
    console.log(element);
})
```

### 13、Pormise对象

在javascript的世界里，所有的代码都是单线程执行的
由于这个缺陷，导致js的所有网络操作，浏览器事件，都必须是异步执行，javascript通过回调函数实现异步的，这是js的一大特色

单线程异步模型

```javascript
function callback(){
    console.log("Deav")
}

console.log("before setTimeout()");
setTimeout(callback,5000);
console.log("after setTimeout()");
```
这里会立即返回，5秒后执行callback，并不是在这个阻塞5s
js通过回调函数实现异步逻辑

**Promise**

如何编写类似于setTimeout的函数，是获取数据setData，要求返回，不能阻塞
通过一种特殊的写法:success(data) error(err)
```javascript
var success=(data)=>(console.log(data))

var reject=(err)=>(console.log(err))

function getData(success,reject){
    //业务逻辑非常耗时，比如2s服务端才返回数据
    setTimeout(()=>{
        // 成功调用success，回调
        success({code:0,data:{name:"laowang",age:18}})
        // 失败者调用reject回调
        reject({code:500,data:{message:"hello",age:18}})
    })
}

getData((data)=>{console.log(data)},(error)=>{console.log(error)})
```

getData函数只关注自身的逻辑，并不关心具体的resolve和reject将如何处理结果
js把这种编程方式抽象成了一种对象:Promise

采用pormise,相比于我们自己传递回调函数，添加了一个语法糖，做成了一个链式调用
```javascript
var p1 =new Promise(getData)

p1.then((resp)=>{
  console.log(resp)
}).catch((err)=>{
  console.log(err)
}).finally(()={
    console.log("finally")
})
```

**Async函数+Promise组合**

从回调函数，到promise对象，再到Generator函数(协程方案的一种过度形态),javaScript异步编程解决方案现在到了Async/await

async函数由内置的执行器进行执行，这和go func()差不多

凡是async标明的函数都会起一个runtime来执行
如果函数内部阻塞了，也不会阻塞主线程

```javascript
async function getData(){
    var p1=new Promise(testResultCallbackFunc)
    try {
        var rest=await p1
        console.log(result)
    }catch(err){
        console.log(err)
    }
}
getData()
```

这里getData就是一个异步函数，他执行的时候，是交给js的携程处理器的，而await关键字就是告诉执行器，当p1执行完之后，主动通知我