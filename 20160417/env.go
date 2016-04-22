type Env interface {
	GetExternalAPI() map[string]string
}
type StgEnv struct {)  // ステージング環境
type PrdEnv struct {}  // 本番環境

func (c *StgEnv) GetExternalAPI() map[string]string {
	return map[string]string{
		"URL": "http://stg.hoge.com/", "Key": "xxxxxxxxxxxxxxxxxxxxxxx",
	}
}
func (c *PrdEnv) GetExternalAPI() map[string]string {
	return map[string]string{
		"URL": "http://hoge.com/", "Key": "yyyyyyyyyyyyyyyyyyyyyyy",
	}
}
func init() {
	switch os.Getenv("ENV") {
	case "production":
		cfg = &PrdEnv{}
	case "staging":
		cfg = &StgEnv{}
	}
}
