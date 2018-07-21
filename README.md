
- udp connected
```text
WriteToUDP 方法 ？不可以
Close之后，是否可以调用WriteToUDP ？ 不可以
ReadFromUDP 方法？可以， 和Read方法相同
```

- udp unconnected

```text
Read 方法 ？  可以， 但是收不到远端地址
Write 方法？  不可以
```

```
WriteMsgUDP  通用方法
```

- UDP 组播
```text
224.0.0.255
同一个应用可以加入到多个组
多个应用可以加入到一个组
多个UDP Listener 可以监听同样的端口，加入到同一个group
```

- UDP 广播
```text
广播采用unconnected的udp connection实现
```

高并发server
```bash
$ netstat -s -u
IcmpMsg:
    InType0: 2
    InType3: 123
    OutType3: 123
    OutType8: 8
Udp:
    738329 packets received
    123 packets to unknown port received.
    778151 packet receive errors
    1516606 packets sent
    RcvbufErrors: 778151
    IgnoredMulti: 189
UdpLite:
IpExt:
    InMcastPkts: 48
    OutMcastPkts: 62
    InBcastPkts: 189
    OutBcastPkts: 10
    InOctets: 52080979
    OutOctets: 50688803
    InMcastOctets: 4000
    OutMcastOctets: 4350
    InBcastOctets: 23677
    OutBcastOctets: 416
    InNoECTPkts: 1523125
$ ethtool -g ens33
```

set proc:
```bash
$ sudo sysctl -w net.core.netdev_max_backlog=2000
$ sudo sysctl -w net.core.wmem_default=26214400
$ sudo sysctl -w net.core.wmem_max=56214400
$ sudo sysctl -w net.core.rmem_default=26214400
$ sudo sysctl -w net.core.rmem_default=56214400
```
