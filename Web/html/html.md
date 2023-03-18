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

### html的常用标签

h1~h6:标题标签
p:定义段落
br :换行
hr :水平分割线
iframe:嵌套外部的页面
img：图像
area:这个标签用于定义图像映射内部的区域
a :链接标签
ul :无序列表
ol :有序列表
li :列表的列表项

表格
table: 用于定义表格
tr:表示定义一行
th:表示表头
td:表示一个单元格
tfoot:表格的尾

容器元素
div:定义html文档中的一个分隔区域或者一个区域部分，标签常用于组合块级元素，以便通过css来对这些元素进行格式化
span：用于对文档中的行内块元素进行组合，标签提供了一种将文本的一部分或者一部分独立出来的方式

元素id
表示元素的身份
使用a标签跳转到某个标签时，使用#加id

元素的样式:style
脚本script
