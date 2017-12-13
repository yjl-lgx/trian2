# myapp 是一个简单的app登录应用。
# 页面介绍
- 1.主页显示“./”页面，简单介绍app应用。
- 2.实现用户注册“/user/signup”页面，接受用户名、密码、个人介绍参数，并写入到mysql数据库中，密码存储需要保证安全；
- 3.实现用户登录“/user/login”页面，登录完成后，跳转到“/user/profile”页面。
- 4.信息显示“/user/profile”页面，输出用户浏览器UserAgent和IP，以及当前运行myapp的主机名；并展示用户名、个人介绍。
- 5.所有页面都会显示当前产品名称以及版本号。
- 6.通过myapp-cli可以远程调用应用中的echo服务，当远程调用服务时，服务器会每分钟推送一次时间。
# 接口说明
- 1.软件名和版本号是通过在Dockerfile中设置，并用函数在内部获取。
```
// Dockerfile中
ENV OEM myapp
ENV VER 5.1.3
##################################
func myhome(){
	myhe = os.Getenv("OEM")
	myha = os.Getenv("VER")
}
```
- 2.提供错误信息打印
```
func checkErr(err error) {
	if err != nil {
    panic(err)
	}
}
```
- 3.提供主页信息
- IP：192.168.34.13:8080
```
func (c *MainController) Get()
```
- 4.登录页面的显示以及密码检索接口
- IP：192.168.34.13:8080/user/login
```
//界面显示
func (c *Zhmyapp) LogA()
//密码检索
func (c *Zhmyapp) LogB()
```
- 5.注册显示以及注册的信息录入数据库
- IP：192.168.34.13:8080:user/signup
```
//界面的显示
func (c *Sigmyapp) SigA()
//信息录入检索
func (c *Sigmyapp) SigB() 
```
- 6.主界面显示所有信息包括：
- 输出用户浏览器UserAgent和IP，以及当前运行myapp的主机名；并展示用户名、个人介绍
- IP:192.168.34.13:8080/user/profile
```
func (m *Myapp) Get()
```
# 部署myapp应用到EKOS集群上
- 1.主机上下载安装包并tar zxvf解压[1]http://192.168.1.234:8080/ekos/offline/
- 2.解压
``` 
tar -zxvf package.tgz
```
- 3.执行安装脚本 deply.sh
- 4.执行ekoslet，进入ekos命令界面
- 5.master设置 和 ceph
```
inventory init master:192.168.34.12:etcd:192.168.34.12:node:192.168.34.13
ceph init rgw:192.168.34.11:mon:192.168.34.11:osd:192.168.34.11

```
- 6.公钥配置 keygen P@ssw0rd 再次查看配置之后再执行（display）
- 7.安装ceph ceph install
- 8.最后install
- 9.将myapp文件编译生成镜像（将编译写进Makefile）执行之后生成镜像mainsver可执行文件
- 10.编写Dockerfile开放端口（8080和50051）以及设置OEM和VER环境变量。执行Dockerfile,构建镜像(使用docker build)
- 11.登录到[1]http://192.168.34.12:30000/login?referer=/ui/ 点击镜像管理--上传镜像--集群外部--选择自己的系统再配置CA证书--之后为了方便重启DOCKER--设置镜像tag--登录镜像仓库（应先拉去mysql的镜像上传再按照步骤上传myapp）
- 在EKOS中操作下面步骤
- 12.创建nfs服务器
- 13.使用创建的nfs服务器添加存储
- 14.使用上传到集群的mysql镜像，在应用管理中，根据指引创建mysql服务（将mysql的数据存储挂载到nfs存储上）
- 15.创建负载均衡，添加mysql服务的转发规则
- 16.使用上传到集群的myapp镜像，在应用管理中，根据指引创建myapp服务
- 17.在负载均衡中，添加myapp的转发规则
- 18.登录提供的网页查看功能的正常运行--以及运行客服端检查。

# 遇到的问题及解决的方法
- 1.在将镜像建好之后将push到集群时一直传不上去。再将集群重新配置之后还是同样的问题
- 解决办法将master版本换成ga版本后解决问题。
- 2.在将镜像部署到ekos上后发现运行时8080和/8080:user/profile能访问的情况下8080/user/login和8080/user/sigunp不能访问
- 解决办法如下将屏蔽掉的代码换成现在的代码即可访问
```
const (
    //address      = "127.0.0.1:60001"
    //address2     = "127.0.0.1:60002"
    address    = "192.168.34.13:60001"
    address2   = "192.168.34.13:60002"
)
```
- 3.在为myapp配置govendor时一直不能拉去到其他两个包。
- 解决办法：在老师的帮助下绕过代理后也拉去不到后，手动一个一个用 govendor sync <packename>拉去下来
暂时解决了问题。
- 4.日志激活后在系统日志正常输出的情况下仍然不能正常查看。
- 版本问题