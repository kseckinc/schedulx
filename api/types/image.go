package types

// ParamsServiceEnv ImageInfo 镜像信息
type ParamsServiceEnv struct {
	ImageStorageType string `json:"image_storage_type"`
	ImageUrl         string `json:"image_url"`
	ServiceName      string `json:"service_name"`
	Port             int64  `json:"port"`
	Account          string `json:"account"`
	Password         string `json:"password"`
}
