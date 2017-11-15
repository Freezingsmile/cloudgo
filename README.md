# cloudgo
cloudgo
 
### 结构框架 

- Martini 

由于构建Web应用应该使用相对于来说较为合适的构架，而Martini经过了相当多的使用者的推荐，且对于常见的问题与步骤有着较为详尽的描述与引导，
所以选择它来作为框架的结构。

- go run main.go -p9090
 
此时启动服务器，并进入，通过在浏览器中输入http://localhost:9090/hello/Freezingsmile(本人用户名),即可访问当前服务器的状况。


- curl -v http://localhost:9090/hello/Freezingsmile

此时通过使用curl观测服务器状态

- 完成cloudgo 的 Martini框架
