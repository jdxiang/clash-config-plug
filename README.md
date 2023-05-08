# 简介
clash 配置文件代理插件，转换配置格式。

# 功能
* 添加 `url-test`的 proxy group。
```
    name: auto-select
    type: url-test
    proxies:
    url: 'http://www.gstatic.com/generate_204'
    interval: 300
```
* 添加转发规则,包括: 国外流量、Telegram、Youtube、️Netflix、️哔哩哔哩、️国外媒体、️苹果服务、️直接连接的转发规则。
![image](https://user-images.githubusercontent.com/132820459/236862707-90bce053-f191-46ea-a587-527116f6000f.png)

# 安装&卸载
```
# 安装
opkg install clash-config-plug_1.0.0-1_all.ipk

# 卸载
opkg remove clash-config-plug
```
# 使用
将订阅地址修改为`${openwrt_adr}:9876/clash_config/get_config?url=${subscribe_url}`。

例如若订阅地址为`https://openwrt.org/subscribe_url`,openwrt本地地址为`192.168.1.1`;

则配置订阅地址为`http://192.168.1.1:9876/clash_config/get_config?url=https://openwrt.org/subscribe_url`。

# 测试
浏览器输入修改后的订阅地址，查看配置文件是否下载成功。
