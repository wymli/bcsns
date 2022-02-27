package logx

type Config struct {
	Env         string `json:",options=[test,dev,prod]"`
	Level       string `json:",default=debug,options=[debug,info,error]"`
	ServiceName string `json:",optional"`
}
