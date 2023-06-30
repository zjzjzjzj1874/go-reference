## [krew](https://github.com/kubernetes-sigs/krew)
* What is [Krew](https://krew.sigs.k8s.io/)?
    * Krew is the plugin manager for kubectl command-line tool.
* [Installing](https://krew.sigs.k8s.io/docs/user-guide/setup/install/)
    * For Mac:
        * 请确保Git安装
        * 执行命令
      ```shell
        (
        set -x; cd "$(mktemp -d)" &&
        OS="$(uname | tr '[:upper:]' '[:lower:]')" &&
        ARCH="$(uname -m | sed -e 's/x86_64/amd64/' -e 's/\(arm\)\(64\)\?.*/\1\2/' -e 's/aarch64$/arm64/')" &&
        KREW="krew-${OS}_${ARCH}" &&
        curl -fsSLO "https://github.com/kubernetes-sigs/krew/releases/latest/download/${KREW}.tar.gz" &&
        tar zxvf "${KREW}.tar.gz" &&
        ./"${KREW}" install krew
        )
      ```
        * 将`$HOME/.krew/bin`添加到环境变量,然后重启terminal
          `export PATH="${KREW_ROOT:-$HOME/.krew}/bin:$PATH"`
        * 执行`kubectl krew`查看
* Plugins
    * 搜索插件: `kubectl krew search`,模糊查询:`kubectl krew search 关键词`
    * 安装插件: `kubectl krew install ktop`
    * [ktop](https://github.com/vladimirvivien/ktop):适用于Kubernetes集群的类似top的工具
        * 安装:`kubectl krew install ktop`
        * 使用
            * 查看所有: `kubectl ktop`
