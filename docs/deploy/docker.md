# 使用docker部署项目

docker 部署一般使用如下文件
````
- docker-entrypoint
- Dockerfile
- Makefile
````

docker-entrypoint 可以让镜像表现的像一个可执行程序一样. 借助了 docker 的 ENTRYPOINT 功能.

Dockerfile 用于构建docker镜像. 需要对docker有一定的了解.

Makefile 用于自动化编译, 此处不再多说, 感兴趣的同学可以了解下 make/makefile.

需要注意的是, 当使用不同的基础镜像时, 其内置工具链等环境很可能时不相同的. 如在如下示例中, 当使用 alpine 作为基础镜像时,
需要将go程序静态编译(alpine 镜像内没有相应的工具链). 具体编译命令见 `Makefile/docker-build`,
解决过程参考: [docker 部署调试记录](https://github.com/everywan/note/blob/master/logs/20190327-docker.md)
