# Web4WoL-SimpleExample

This is a program for users who use web(such as browser, curl, and so on...) to trun on the computer.(BASE ON WOL)

But this just a example, it's my learning object for Golang, but it can work on the instance.

## Build
So, you just only alter the `ip:port` ,`port` and `0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF`(this is your net-card's MAC) which are your likes.

Then use `go build` to build it. Of course you need to install golang enviroment.

Congratulation! You can use it now!

## How to use
Run your program, then use browser or others tool which can connect the website.

Input your url into the tool.

Example:`helloworld.com:1234/openpower`

    helloworld.com //your domain, also you can use ip.
    :1234          //your host's port.
    /openpower     //the route. Do Not Forget to input it!

So, this exmaple is this.

    ip(or domain):port/openpower

Then hold your machine trun on.

--------

自制的一个使用WoL的方法，通过输入Web域名或者IP到浏览器实现远程开机的一个小工具的代码，前提是电脑要支持WoL.

用的Golang写，其实是一个自己学Golang然后尝试做的一个作品，代码很简单，以后有机会再维护扩大吧。

## 编译
把源文件里面的`ip:port` ,`port` and `0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF`改成你自己得，后面那一串0x是你的MAC地址，自己去查网卡的MAC地址，对应填上去就行。

然后用`go build`编译成文件就行了，环境什么的就不用我多说了吧，也可以试试Github的Action

## 怎么用
直接跑就是了，运行软件，然后打开浏览器，输入你的IP+端口，或者域名，反正你上面改了啥就输入啥，格式如下

    ip(or domain):port/openpower
    
`/openpower`这个不能删不能删不能删，这个是一个路由。删掉会出现404，不理解无所谓，就是考这个路径执行命令，知道就行了。

亲测能用，不能用的话检查你的网络，特别是防火墙那些，记得端口放行，用之前先ping一下要做远程开机的电脑测试连通性。
