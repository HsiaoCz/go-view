# python 的隐式转换
a = False
b = 1
c = a + b

print(type(c))
print(c)
# 布尔遇到数字会被转换为数字类型,false为0
# true为1
aa = 0
print(bool(aa))
a1 = 2
print(bool(a1))
# 如果是强制转换 非零的数字会被转换为True,0为flase
m = "hello"
print(bool(m))
m1 = ""
print(bool(m1))
# 强制转换 非空字符串会被转换成True,空字符串被转换成flase

m2 = True
m3 = False
print(str(m2))
print(str(m3))
# bool类型的转换为str任然为True或False
