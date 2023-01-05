# debug-go
The golang version of remote debug tool :dizzy:. 

```textmate
██████╗ ███████╗██████╗ ██╗   ██╗ ██████╗ 
██╔══██╗██╔════╝██╔══██╗██║   ██║██╔════╝ 
██║  ██║█████╗  ██████╔╝██║   ██║██║  ███╗
██║  ██║██╔══╝  ██╔══██╗██║   ██║██║   ██║
██████╔╝███████╗██████╔╝╚██████╔╝╚██████╔╝
╚═════╝ ╚══════╝╚═════╝  ╚═════╝  ╚═════╝                                           
```


```textmate

Usage:	debug [Options] COMMAND

A fast remote debugging environment configurator.

Options:
      --config string      Location of client config files (default "/Users/apple/.docker")
  -D, --debug              Enable debug mode
      --help               Print usage
  -H, --host list          Daemon socket(s) to connect to
  -l, --log-level string   Set the logging level ("debug"|"info"|"warn"|"error"|"fatal") (default "info")
  -v, --version            Print version information and quit

Commands(PHP):
  xdebug      在远程服务器上开启xdebug插件
  dbgpProxy   在远程服务器上开启DBGp Proxy代理
  
Commands(Golang):
  dlv         Start a dlv process
  rsync       Start a Rsync Server waiting for local binary files
 
Commands(Python):
  attach      Attach local standard input, output, and error streams to a running container
  build       Build an image from a Dockerfile

Commands(C/C++):
  attach      Attach local standard input, output, and error streams to a running container
  build       Build an image from a Dockerfile
```