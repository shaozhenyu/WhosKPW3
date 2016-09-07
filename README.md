# WhosKPW3

接口规格

接口方法1： 请求：

http://host/set?key=<key>&value=<value>
(body为空)
期望返回(body)：

OK
接口方法2： 请求：

http://host/get?key=<key>
(body为空)
期望返回(body)

value=<value>
其中key和value为不包含空格、等号的字符串，长度为1-257字节随机分布


目前性能最佳 fasthttp

优化想法：
1.nginx
2.修改fasthttp源码
