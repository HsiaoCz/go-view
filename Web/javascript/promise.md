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