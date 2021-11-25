# **SchedulX**

SchedulX 是基于开源 BridgX 项目的云原生服务编排和部署解决方案，目标是让开发者在 BridgX 获取的计算资源上进行服务编排和部署。

它具有如下关键特性:

1、具备结合动态扩缩容特性进行服务部署的能力；

2、一个平台统一管理在不同云平台上的服务操作；

3、简洁易用，轻松上手；


安装部署
--------

请先按照 [BridgX上手指南](https://github.com/galaxy-future/BridgX/blob/dev/README.md) 完成BridgX和BridgX_FE的安装。要求内网部署环境，能够跟云厂商VPC连通。

* (1)源码下载
  - 后端工程：
  > `git clone git@github.com:galaxy-future/SchedulX.git`

* (2)macOS系统部署
  - 后端部署,在SchedulX目录下运行
    > `make docker-run-mac`

  - 系统运行，在浏览器中输入 http://127.0.0.1 可以看到管理控制台界面,初始用户名root和密码为123456。

* (3)Linux安装部署
  - 1）针对使用者
    - 后端部署,在SchedulX目录下运行,
      > `make docker-run-linux`
    - 系统运行，浏览器输入 `http://127.0.0.1 `可以看到管理控制台界面,初始用户名 `root`和密码为`123456`。

  - 2）针对开发者
    - 后端部署
      - SchedulX依赖mysql组件，
           - 如果使用内置的mysql，则进入SchedulX根目录，则使用以下命令：            
             > `docker-compose up -d`    //启动SchedulX <br>
             > `docker-compose down`    //停止SchedulX  <br>
           - 如果已经有了外部的mysql服务，则可以到 `cd conf` 下修改对应的ip和port配置信息,然后进入SchedulX的根目录，使用以下命令:
             > `docker-compose up -d schedulx`   //启动schedulx服务 <br>
             > `docker-compose down`     //停止SchedulX服务

      - 系统运行，浏览器输入 `http://127.0.0.1` 可以看到管理控制台界面,初始用户名 root 和密码为123456。

快速上手
------------
通过[快速上手指南](https://github.com/galaxy-future/SchedulX/blob/master/docs/getting-started.md)，可以掌握基本的服务扩缩容流程。

行为准则
------
[贡献者公约](https://github.com/galaxy-future/SchedulX/blob/master/CODE_OF_CONDUCT.md)

授权
-----

SchedulX 使用[Apache License 2.0](https://github.com/galaxy-future/SchedulX/blob/master/README.md)授权协议进行授权
