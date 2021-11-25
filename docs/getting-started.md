# 概览
本文是SchedulX的快速入门指南，典型的用户包括三个步骤：<br>
1. 添加云厂商账户；<br>
2. 在集群管理模块创建需要的集群模板；<br>
3. 在服务部署模块，进行服务扩缩容；<br>

# 前置条件
1、已经有云厂商账户，并且获得了AccessKey和AccessKey Secret信息，如果没有，请提前申请；<br>
[阿里云申请入口链接](https://help.aliyun.com/document_detail/53045.html)<br>

2、已经获取了网址信息，以及用户名和密码信息；<br>
如果是初次部署，可在浏览器中打开 http://127.0.0.1:80 登录系统,采用的认证方式是默认配置可用以下帐号登录<br>
默认用户名:root<br>
默认密码为:123456<br>

# 第一步：添加云厂商账户

1. 进入云厂商账户，点击添加云账户
![image](https://user-images.githubusercontent.com/94337797/142158688-a3a17da1-a068-4396-81fb-cf1f1270f184.png)

2. 填写云厂商的ak和sk信息，并保存；
![image](https://user-images.githubusercontent.com/94337797/142158808-19166f17-9ed6-4f5e-9ffe-65f698bbe7ed.png)


# 第二步：创建机型模板

1. 集群管理->创建集群
![image](https://user-images.githubusercontent.com/94337797/143417597-28f11c05-4fb0-4d82-83fc-3e8eaf4b5307.png)


2. 进入云厂商配置页面，进行配置：
![image](https://user-images.githubusercontent.com/94337797/143417734-fa659a8b-234d-4b2a-b95a-ddde5388c46a.png)


3. 点击下一步进入网络配置页面：
![image](https://user-images.githubusercontent.com/94337797/143417837-bdad0940-d97e-44c9-99f5-8eea31f961a8.png)


4. 网络配置完成后，点击下一步，进行机器规格配置
![image](https://user-images.githubusercontent.com/94337797/143417976-210e8ba9-4a6c-40c4-acbf-a7d2f103eb0d.png)


5. 机器规格配置完成后，点击下一步，进行系统配置，配置完成后，点击提交，创建集群成功
![image](https://user-images.githubusercontent.com/94337797/143418056-c1220cf0-2fa0-45da-836b-43c05554d6c5.png)


6. 点击提交，则页面跳转到集群列表，有最新的创建的集群信息，则表示创建成功
![image](https://user-images.githubusercontent.com/94337797/143418186-cecc4cb5-4b86-46ee-a6db-4cd7ddd7abd0.png)



# 第三步：服务扩缩容

1. 服务部署->创建服务；
![image](https://user-images.githubusercontent.com/94337797/143418424-d82216df-7b88-4ddb-b3bf-25b301e1dd16.png)


2. 填写相关的服务信息后，保存；
![image](https://user-images.githubusercontent.com/94337797/143418932-c2010f43-765a-4dee-86b5-b8a6f0efadd4.png)


3. 提交成功后，页面跳转到服务列表，显示刚才创建的服务
![image](https://user-images.githubusercontent.com/94337797/143419062-5e1d4ae1-8766-4cd7-a134-3869362d45dd.png)


4. 对服务进行扩缩容，在操作列表点击相应服务的扩缩容；
![image](https://user-images.githubusercontent.com/94337797/143419363-9293d820-cf3b-40e5-b88a-d5e7d727cd96.png)

5. 在扩缩容页面，根据需求进行操作；
![image](https://user-images.githubusercontent.com/94337797/143419451-2811560b-905b-4e92-bc78-55f69f51dcd5.png)

6. 如果需要查看详情，可以在操作列表，点击扩缩容历史查看；
![image](https://user-images.githubusercontent.com/94337797/143419744-eaeee07b-6d25-4cc0-a8d1-88018e1dea70.png)

7. 服务扩缩容执行详情；
![image](https://user-images.githubusercontent.com/94337797/143419885-4f9c93d7-a96a-40c2-b37c-e061d7d95d56.png)



