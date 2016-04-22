# mockgen -source "対象のインターフェース" \
		-destination "出力するモックのファイルパス" \
		-package "出力するパッケージ名"

# 例：
mockgen -source ${GOPATH}/src/github.com/aws/aws-sdk-go/service/kinesis/kinesisiface/interface.go \
	-destination mock/kinesismock.go \
	-package awsmock

