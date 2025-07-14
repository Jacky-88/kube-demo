package config

const (
	// 监听地址
	ListenAddr = "0.0.0.0:9090"
	WsAddr     = "0.0.0.0:8081"
	//Kubeconfig路径
	//Kubeconfigs = `{"TEST1":"C:\\custom\\project\\config"}` //windows路径
	Kubeconfigs = `{"TEST1":"/Users/jacky/.kube/aliyun-jlc-ops-test-config"}` //mac路径
	//账号密码
	AdminUser = "root"
	AdminPwd  = "luyijian"
)
