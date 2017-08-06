# sbg
Ssh Banner Generator

# description
OpenSSH provides advanced cryptographic operations for authentication at various levels.
One of the basic verification is to avoid MITM between you and the remote server using server host keys

Common verification goes with hash values, however OpenSSH provides a visual way
to verify you are speaking with your host.

You can display it at client connection time using the following in your ssh
client configuration:
```
VisualHostKey yes
```

However you need to compare it to something... 
may be I missed something.

This small tools helps quickly generate banners that are displayed pre-authentication in order to compare with your visual hostkey at connection time.

```
$ ssh host
Host key fingerprint is SHA256:PUI2nXLoM93QX8C9Lt1DuO6gmyhH6a8MFsj4KHVXkx8
+---[ECDSA 256]---+
|             ... |
|         + o  ...|
|        X E . . o|
|  o .  = O + o + |
| ..o... S = . * .|         <== this is from my ssh client "VisualHostKey yes"
| .o. ..o + . o +.|
|.. . oo    .. . .|
|.   ..oo. o ..   |
|      o+o=. ..   |
+----[SHA256]-----+

+---[DSA 1024]----+  +---[ECDSA 256]---+  +--[ED25519 256]--+  +---[RSA 2048]----+
|       .. ... .+.|  |             ... |  |                 |  |         .oo.E.. |
|   . oE. . +..+ O|  |         + o  ...|  |   .             |  |         .o  ..o.|
|    + o . +.o  @o|  |        X E . . o|  |. . o .          |  |        .o o +o..|
|     . o o =o . =|  |  o .  = O + o + |  |+.   = o        E|  |       .+ = =..  |
|        S =. + .o|  | ..o... S = . * .|  |oo.   = S       .|  |        SO + .+  |
|       . o..= o +|  | .o. ..o + . o +.|  |o .. o. ooo     .|  |       =o++ o=   | <== generated banner 
|        .  + = +.|  |.. . oo    .. . .|  | . .o. o.B.     .|  |      + =..oo    |
|            + * o|  |.   ..oo. o ..   |  |  .=*oo =+..  o .|  |     . O o=.     |
|           . o o+|  |      o+o=. ..   |  |  =**X=*Oo  .o . |  |      o *=**.    |
+----[SHA256]-----+  +----[SHA256]-----+  +----[SHA256]-----+  +----[SHA256]-----+

SHA256:W5mDVO36dmOhNPbVxeiw/18TMbSFdydnoCalqQY43IU (DSA)
SHA256:PUI2nXLoM93QX8C9Lt1DuO6gmyhH6a8MFsj4KHVXkx8 (ECDSA)
SHA256:GxgO6F/w+mLveIqbbVdPsqWeqqwDgKP9U2/uuVt/dlU (ED25519)
SHA256:ox5bnM7L/HCgRZWUvjsE/UiyGdecq2v76lVaRtkpVkg (RSA2048)

```

# installation

```
go get github.com/unix4fun/sbg
```

# usage

```
sbg -h
```

Example:

```
$ sbg /etc/ssh/ssh_host_*.pub > /etc/ssh/myhostbanner.txt
```

then add to your sshd_config:
```
Banner /etc/ssh/myhostbanner.txt
```


