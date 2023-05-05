promise 表示一个异步操作的最终完成(成功或者失败，以及最终返回的结果值)

可以绑定回调函数，避免回调函数作为参数
链式调用，避免回调地狱

写一个以前没有promise对象时，是怎么进行异步操作的

```javascript
const success=msg=>console.log(msg)
const failed=msg=>console.log(msg)

const wait=(wait,success,fail)=>{
 setTimeout(()=>{
    if (true){
        success("success")
    }else{
        failed("fail")
    }
 },wait);

}
wait(5000,success,fail);
```
如何使用promise

```javascript
const wait=ms=>new Promise((resolve,reject)=>setTimeout(resolve,ms))
wait(3000).then(()=>console.log("success")).catch(()=>console.log("fail"))
```

在javascript中有两种实现异步的方式

1、回调函数

```javascript
setTimeout(()=>{
    console.log("兄弟你好");
},3000)

console.log("你会立刻看到我");
```
这个函数本身会立刻返回，程序会紧接着执行之后的代码
而我们传入的回调函数则会等到预定的时间才会执行

js是单线程的，即便看上去回调函数和主程序在并发执行，但它们都运行在同一个主线程中
回调函数的缺点，当依次执行多个异步操作，程序会层层嵌套
这种情况被称为函数的回调地狱

promise：
promise代表一个"承诺",承诺在未来会返回某个数据
然后调用then方法，传递一个回调函数，如果这个请求在未来成功完成，那么回调函数会被调起，请求的结果也会以参数的形式传递进来
promise的优点在于它可以用一种链式结构将多个异步操作串联起来
```javascript
fetch("https://jsonplcaeholder.typicode.com/posts/1").then((response)=>{
    //
})
// 代表未来的某个时刻将返回的数据用json的格式表示
// 如果我们希望在它完成之后在追加别的操作，我们可以在之后在追加then
// promise可以在then之后添加一个catch()来捕获之前任意阶段出现的错误
// 如果之前任意节点出现了错误，那么之后的then将不会被执行
// promise还提供finally方法，会在promise链结束之后调用，无论失败与否
fetch("https://jsonplcaeholder.typicode.com/posts/1").then((response)=>response.json())
.then((json)=>console.log(json));
```

async和await
简单来说，这两个是基于promise之上的一个语法糖,可以让异步操作更加简单明了
首先使用async关键字将函数标记为异步函数，异步函数这里指的是返回值为promise对象的函数
在异步函数中，可以调用其他的异步函数，不过不需要再使用then，只需要加上一个await
await会等待promise完成之后直接返回最终的结果
```javascript
async function f(){
    // 这里response就是返回的数据
   const response=await fetch("https://...");
    // await看上去会暂停函数的执行
    // 但是实际上，在等待的过程中,js同样可以处理其他的任务
    // await底层是基于promise和时间循环机制实现的
};

f(); // 注：这个函数返回值永远是promise对象
```

注意await的这几个陷阱

1、
```javascript
async function f(){
    // 这样写会打破这两个fetch的并行
    // 这里会等到第一个任务执行完成之后才开始执行第二个任务
    // 这里更高效的做法是将所有promise用promise.all组合起来,然后再去await
    // 修改后的程序运行效率也会直接提升一倍
    const a=await fetch("https://../post/1")
    const b=await fetch("https://../post/2")

    const PromiseA=await fetch("https://../post/1")
    const PromiseB=await fetch("https://../post/2")

    const [a,b]=await Promise.all([promiseA,promiseB]);
}
```

2、
在循环中执行异步操作，是不能够直接调用forEach或者map这一类的方法的
尽管在回调函数中写了await
但是forEach会立刻返回，它并不会暂停等到所有的异步操作都执行完毕

如果希望等待循环中的异步操作都一一完成之后才继续执行
那我们还是应该使用传统的for循环
```javascript
async function f(){
    [1,2,3].forEach(async (i)=>{
        await someAsyncOperation();
    });

    console.log("done");
}

async function f(){
    for (let i of [1,2,3]){
        await someAsyncOperation();
    }

    console.log("done");
}
f();
```

如果我们希望循环中的所有操作都并发执行
一种更炫酷的写法是使用for await，这里的for 循环依然会等待所有的异步操作都完成之后才继续向后执行

```javascript
async function f(){
    const promises=[
        someAsyncOperation(),
        someAsyncOperation(),
        someAsyncOperation(),
    ];

    for await (let result of promise){
        // 
    }
    console.log("done");
}

f();
```

3、不能在全局或者普通函数中直接使用await关键字,await只能被用在异步函数中

如果想要在外层中使用await，需要先定义一个异步函数，然后在函数体中使用await

```javascript
async function f(){
    await someAsyncOperation();
}

f()

// 或者更简易的代码
(async ()=>{
    await someAsyncOperation();
})();
```

使用async和await可以写出更清晰、更容易理解的异步代码
