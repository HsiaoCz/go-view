### HTML

HTML:超文本标记语言(不做逻辑处理)
作用：告诉浏览器如何构造网页

```html
<p>this is element</p> 
<!-- 起始标签和结束标签 -->
```

```html
<!-- 解释文档的类型，解释这是一个html5的文档 -->
<!-- 为了兼容一些老的解释类型 -->
<!DOCTYPE html>
<head>
    <!-- head 放搜索引擎的关键字，标题等 -->
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <!-- 真正要给浏览器渲染的内容 -->
</body>
</html>
```

块级元素
- 在页面以块的形式展示
- 出现在新的一行
- 占全部宽度

内联元素
- 通常在块级元素内
- 不会导致文本换行
- 只占必要的部分宽度