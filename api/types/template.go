package types

const (
	MountValueALB   = "slb"
	MountValueNginx = "nginx"
)

const (
	ENVInt       = "nodeact.initbase" //环境初始化
	SERVICEInt   = "nodeact.initsvc"  // 服务初始化
	MountTypeSLB = "mount.slb"        //挂载slb
)

type TmpInfo struct {
	TmplName         string `json:"tmpl_name"`
	ServiceClusterId int64  `json:"service_cluster_id"`
	Describe         string `json:"describe"`
	BridgxClusname   string `json:"bridgx_clusname"`
}

type BaseEnv struct {
	IsContainer bool `json:"is_container"`
}

type ServiceEnv struct {
	ImageStorageType string `json:"image_storage_type"`
	ImageUrl         string `json:"image_url"`
	Port             int64  `json:"port"`
	Account          string `json:"account"`
	Password         string `json:"password"`
	Cmd              string `json:"cmd"`
	ServiceName      string `json:"service_name"`
}

type ParamsMount struct {
	MountType  string `json:"mount_type"`
	MountValue string `json:"mount_value"`
}
