module go-admin

go 1.15

require (
	git.nonolive.co/FE/mildom-video v0.0.0-20210714134605-c5988881925d
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/alibaba/sentinel-golang v0.6.1
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190307165228-86c17b95fcd5
	github.com/bytedance/go-tagexpr/v2 v2.7.10
	github.com/casbin/casbin/v2 v2.25.1
	github.com/douyu/jupiter v0.2.7
	github.com/gin-gonic/gin v1.7.1
	github.com/go-admin-team/go-admin-core v1.3.5-rc.6.0.20210629084023-dc3d0627688c
	github.com/go-admin-team/go-admin-core/sdk v1.3.5-rc.6.0.20210629084023-dc3d0627688c
	github.com/google/uuid v1.2.0
	github.com/lib/pq v1.8.0 // indirect
	github.com/mssola/user_agent v0.5.2
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.11.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/shirou/gopsutil v3.21.5+incompatible
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/objx v0.3.0 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	github.com/tklauser/go-sysconf v0.3.6 // indirect
	github.com/unrolled/secure v1.0.8
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	gorm.io/driver/mysql v1.0.4-0.20201206014609-ae5fd10184f6
	gorm.io/driver/postgres v1.0.6-0.20201208020313-1ed927cfab53
	gorm.io/driver/sqlite v1.1.5-0.20201206014648-c84401fbe3ba
	gorm.io/gorm v1.21.10
)

replace (
	//git.nonolive.co/FE/mildom-video   => ../mildom-video
	github.com/go-admin-team/go-admin-core => ../go-admin-core
	github.com/go-admin-team/go-admin-core/sdk => ../go-admin-core/sdk
	google.golang.org/grpc v1.31.0 => google.golang.org/grpc v1.26.0
)
