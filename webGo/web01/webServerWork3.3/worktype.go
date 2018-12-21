package main

//web工作方式的几个概念：
//request :客户端请求的信息，用来解析用户的请求信息，包括post get cookie url等信息
//response :服务器需要反馈给客户端的信息
//conn :用户每次请求链接
//Handle :处理请求和生成反馈信息的处理逻辑



//http包运行机制
//1：创建listen socket，监听指定的端口，等待客户短请求到来；
//2：listen socket 接受客户端请求，得到client socket ，接下来通过client socket与客户端通信；
//3：处理客户端请求，首先从client socket 中读取http请求的协议头，
// 如果是post方法，还可能读出客户端提交的数据，然后交给相应的handler处理请求，
// handler处理完毕准备好客户端需要的数据，通过client socket写给客户端；


//要清楚三件事：
//1：如何监听端口
//2：如何接受客户端请求？
//3：如何分配handler

//go通过一个函数listenAndServer来处理这些事情
//底层的处理是：初始化一个server对象，然后调用net.listen("tcp",addr),也就是用tcp搭建一个服务，然后监听我们设置的端口；（问题1）

//监听好端口后，会调用srv.Server(net.listener)函数，这个Server函数用来处理客户端请求，里边起了一个for{}，首先accept()接受请求，
//然后创建一个Conn，单独开一个goroutine，把这个请求的数据当作参数扔给这个conn去服务：go c.serve();这就是高并发体现了，
//用户的每一个请求都是在一个新的goroutine去服务器，互不影响；（问题2）

//conn首先会解析request：c.readRequest()，然后回去相应的handler：handler:=c.server.Handler,也就是ListenAndServe时候的第二
//个参数，我们前边例子穿的是nil，也就是空，那么默认获取handler=defaultServeMux，这个变量是什么能？
//就是一个路由器，用来匹配url跳转 到其相应的handler函数，我们在哪设置过呢？http.handlerFunc("/",sayhelloName)，
//这个作用就是注册请求/的路由规划，请求为/路由就会跳转到sayhellowName，defaultServerMux会调用ServerHTTP方法，这个方法内部就是say...
//本身，最后通过写入reponse的信息反馈给客户端；

func main() {
	
}
