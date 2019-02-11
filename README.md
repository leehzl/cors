# cors
总结了跨域的一些问题，并根据[CORS跨域原理浅析](https://zhuanlan.zhihu.com/p/29980092) 文章，自己实现了go版本，复习下关于跨域访问的一些知识点。

# 简介
本例子将client端和server端部署在不同端口实现跨域，其中client在8000端口，server在8080端口

# 静态资源服务器
在Mac下，用nginx作为静态资源服务器，并将nginx的配置文件nginx.conf设置为
```
server {
        listen       8000;
        server_name  localhost;
        location / {
            root   /Users/leehzl/cors/client;
            index  index.html index.htm;
        }
}   
```

具体Mac下的nginx命令参考[Mac利用Nginx搭建静态文件服务器](https://www.jianshu.com/p/2083ccf2ade6)