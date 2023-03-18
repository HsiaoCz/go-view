## 浏览器基础

先抛出两个问题
- 如何适配屏幕大小?
- 如何兼容多种浏览器?

js可以获取浏览器提供的很多对象，并进行操作

### 浏览器对象

网页是通过浏览器加载出来的，而浏览器本身也有很多功能

- 浏览器本身版本信息
- 当前窗口大小
- 历史记录

而这些数据都是绑定在window对象上的，所以你可以认为对浏览器加载的网页window就表示浏览器本身

### 浏览器窗口大小

window的如下4个属性控制这个浏览器的窗口大小，
innerHeight:页面视图的高
innerWidth:页面视图的宽
outerWidth:浏览器整个窗口的宽
outerHeight：浏览器整个窗口的高

### 浏览器的基本信息

我们在页面的时候，如何知道用户使用的

window.navigator

常用的属性：
navigator.appName:浏览器的名称
navigator.appVersion：浏览器的版本
navigator.language:浏览器设置的语言
navigator.platform：操作系统类型
navigator.userAgent：浏览器设定的User-Agent字符串
navigator.userAgentData:useragent的相关信息

如何判断用户使用的是pc还是手机？
userAgentData里面有一个字段:mobile,还有platform，通过这两个字段可以判断用户使用的是否是手机

### 屏幕相关的信息

window.screen:可以获取屏幕的信息

### 访问的网址

window.location
这个很有用，为了保证每个用户访问URL都呈现

一些常用的属性:
location.protocol:// "http"
location.host // 'www.example.com'
location.port // '8080'
location.pathname://"/path/index.html"
location.search:// "?a=1&b=2"
location.hash:// "TOP"

需要加载一个新页面，可以调用location.assign()，如果要重新加载当前页面，调用location.reload()方法非常方便

### AJAx

ajax不是javascript的规范，ajax意思是用javascript执行异步网络请求
简单来说就是浏览器中的http client
在现代浏览器上写Ajax，主要依赖XMLHttpRequest

```javascript
// 新建XMLHttpRequest
var request=new XMLHttpRequest();

// 发送请求
request.open("GET","https://www.baidu.com");
request.send();
```

如何就这样我们是获取不到请求结果，由于js都是异步的，我们需要定义回调来处理返回

```javascript
request.onreadystatechange=function(){
    console.log(request.status)
    console.log(request.responseTest)
}
```

对于Ajax请求特别需要注意跨域的问题:CORS

要理解跨域需要先理解两个概念:

**简单请求**
满足下面3个条件的就是简单请求：
- HTTP Method： GET、HEAD和POST
- Content-Type：application/x-www-form-urlencoded,muttipart/form-data和text/explain
- 不能出现任何自定义头

通常能满足90%的需求
控制其跨域的关键头来自于服务端设置的Access-Control-Allow-Origin
对于简单请求，浏览器直接发送CORS请求，具体来说，就是在头信息之中，增加一个Origin字段
- 客户端发送请求时:Origin.xxx
- 服务端响应请求时，设置Access-Control-Allow-Origin.xxx
如果服务端不允许该域，就跨域失败
如果允许特定的header，服务端也可以通过添加Access-Control-Expose-Headers来进行控制

**复杂请求**

复杂请求

- HTTP Method:PUT,DELETE
- Content-Type：其他类型如application
- 含自定义头，比如我们后面携带的用于认证的X-OAUTH-TOKEN头

简单请求的CORS请求，会在正式通信之前，增加一次HTTP查询请求，称为预检请求(preflight)

preflight请求：
- Method:OPTIONS
- Header:Access-Control-Request-Method：列出浏览器的CORS请求会用到哪些HTTP方法
- Header:Access-Control-Request-Headers:该字段是以一个逗号分隔的字符串，指定浏览器CORS请求会额外发送的头信息字段

preflight响应：
- Header:Access-Control-Allow-Origin,允许的域
- Header:Access-Control-Allow-Methods:服务端允许的方法(所有方法，一般大于请求的该头)
- Header:Access-Control-Allow-Headers:服务队，允许自定义Header
- Header:Access-Control-Max-Age:用来指定本次预检请求的有效期，避免多次请求，该字段可选

**axios**

### js如何操作DOM

当前浏览器的page，在浏览器叫document
由于HTML在浏览器中以DOM形式表示为树形结构，document对象就是整个DOM树的根节点
比如我们获取当前Document中的一些属性，比如title document的title属性是从HTML文档中的<title>xxx</title>读取的，但是可以动态改变

```document
document.title

// 可以动态修改
doucment.title="测试方法"
```

很多js库，比如JQuery都是通过动态操作DOM实现很多高级功能的，这些是上层库的基石

**DOM查询**

要查找DOM树的某个节点，需要从Document对象开始查找，最常用的查找是根据ID和TAG Name以及ClassName

document.getElementsById("hello")
document.getElementsByTagName("hello")
document.getElementsByClassName("hello")

当然我们也可以自己组合使用
document.getElementById("hello").getElementByTagName('11')
document.getElementById("hello").children
document.getElementById("hello").firstElementChild
document.getElementById("hello").lastElementChild
document.getElementById("hello").parentElement

**更新dom**

获取到元素后，我们可以通过元素的如下2个方法，修改元素

- innerHTML:修改html标签
- innerText:修改文本

**插入DOM**

很多响应式框架都会根据数据新增，动态的创建一个DOM元素，并植入到指定的位置
我们使用createElement来创建一个DOM元素，比如创建一个a标签

```javascript
// 创建一个a标签
var newlink=document.createElement('a')

// 修改a标签的属性
newlink.herf="www.baidu.com"
newlink.innerText="百度"

document.getElementByTagName('u1').appendChild(newlink)

// 插入到coffee之前
lm.insertBefore(newlink,cf)
```

总结，有两种方式可以插入一个DOM元素
- appendChild把一个子节点添加到父节点的最后一个子节点
- insertBefore插入到某个元素之前

删除DOM
删除一个DOM节点就比插入要容易的多
要删除一个节点，首先要获得该节点本身以及它的父节点，然后调用父节点的removeChild把自己删掉

removeChild的语法如下：
parent.removeChild(childNode);
比如我们删除刚才添加的那个元素
lm.removeChild(newlink)