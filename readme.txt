https://github.com/flyman-liu/jiandan.git

https://github.com/JingzhaoLiu/jiandan.git
https://github.com/hunterhug/lizard
https://github.com/hunterhug/marmot.git

基于docker的分布式爬虫https://www.ctolib.com/zerg.html
1，定时任务
2，根据配置文件获取所有站点的域名
3，根据域名读数据库，获取栏目
4，根据栏目获取长尾词。
5，根据长尾词抓取文章
6，长尾词关系表
      长尾词-长尾词序号-栏目-文章标题-地址-域名-文章ID-被访问次数-是否收录
     更新以上关系表
7，长尾词表
     长尾词，权重，序号，相关词
8,  栏目-长尾词对应表
     栏目-长尾词序列 （上一次该栏目用到的长尾词序列）
9, 每个域名起一个进程进行爬取
	 gprc,yaml,nsq队列,govendor,pg数据库,ansible,crono
	 ------------------------------------------
	 相关词挖掘服务
	 站群展示
	 后台管理
	 爬虫服务
	 站群管理：
	 所有数据库创建，建站初始化等
	 爬虫客户端
	 站群管理
	 管理接口实现
