## 前言
> 无意中发现了G安全团队的JS-Forward项目，感觉这思路挺好的，自己经常也会遇到金融类比较变态的数据加密，靠这小技巧成功解决了数据加密与篡改；

> 关于工具的思路与原理请访问该项目查看：

> https://github.com/G-Security-Team/JS-Forward

> 思路挺好的，不过测试了该工具，不太适合我这个金融类的项目，因为APP内置的浏览器https跨域到http阻止了请求，所以加了个证书；再就是没有使用镜像端口转发；而是android端设置WIFI代理为Burp，js请求明文到jsForward中都会经过burp，则篡改数据就由Burp操作；

![index](https://raw.githubusercontent.com/CTF-MissFeng/jsForward/main/img/1.png)

## 使用方法
> 请参考G安全团队的JS-Forward项目说明，找到h5加解密方法，加载jquery框架，在加解密方法里插入ajax请求（如果是https的则url地址为`https://jsForward监听地址/api/request`并且安装jsForward证书、http则是`http://jsForward/api/request`）,`/api/response` 则是响应解密(为了区分请求和响应，所以设计了两个接口)

![index](https://raw.githubusercontent.com/CTF-MissFeng/jsForward/main/img/5.png)


> 真实案例:一个金融类APP的H5，可看到请求和响应都是加密的，最恶心的是秘钥是动态的（上一个请求返回包里的），并且解密了该数据包，还发现另一层加密，恶心至极；

![index](https://raw.githubusercontent.com/CTF-MissFeng/jsForward/main/img/2.png)

> 通过在加密前把明文发到到burp中进行篡改，在返回给app进行加密：

![index](https://raw.githubusercontent.com/CTF-MissFeng/jsForward/main/img/3.png)

> 同样的，响应包解密后也发送到burp中便于查看：

![index](https://raw.githubusercontent.com/CTF-MissFeng/jsForward/main/img/4.png)