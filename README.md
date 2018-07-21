
- connected
WriteToUDP 方法 ？不可以
Close之后，是否可以调用WriteToUDP ？ 不可以
ReadFromUDP 方法？可以， 和Read方法相同


- unconnected
Read 方法 ？  可以， 但是收不到远端地址
Write 方法？  不可以

WriteMsgUDP  通用方法

- UDP 组播
224.0.0.255
同一个应用可以加入到多个组
多个应用可以加入到一个组
多个UDP Listener 可以监听同样的端口，加入到同一个group

- UDP 广播
广播采用unconnected的udp connection实现

