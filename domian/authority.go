package domian

// 权限校验里面会用到的数据结构,用于比较的节点
type UrlNode struct {
	Url   string
	Level int
}

// 用户的权限信息
type AuthorityUserInfo struct {
	UId int64 //用户id,某些业务会使用到
	//考虑存用户信息，其实一些特殊的业务数据可以对这个类进行扩展
	//可以考虑存角色信息，到时候自定义就行了
	UrlMap  map[string]byte //map的，主要为了命中常规路由（常规路由就是不带:的）
	UrlList []UrlNode       //主要为了命中非常规路由（带：的，可以提前处理为正则表达式），这样就避免了常规路由的重复
}
