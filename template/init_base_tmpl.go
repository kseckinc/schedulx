package template

const initBaseEnvCmd = `
#!/bin/bash
res=$(docker version | grep Version)
if [ "" == "$res" ]; then
  echo "1、安装docker" >> /root/result.log
  yum install -y docker >> /root/result.log
fi
service docker start >> /root/result.log

echo "success" >> /root/result.log
echo "success"
`

func GetInitBaseCmd() string {
	return initBaseEnvCmd
}
