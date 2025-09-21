package errcode

const Success = 200000

// todo 格式待定義 (外部模組代號+內部服務代號+自定義碼)
// (內)game 10
// (內)member 20
// (內)txn 30
// (內)report 40
// (外)gateway 10
// (外)thirdparty-game 20
// (外)thirdparty-hall 30
// (外)notify 50
const (
	// 公用錯誤代碼
	CommonParamError        = 200001 + iota //參數錯誤
	CommonUnKnowError                       //未知錯誤
	CommonGRPCError                         //GRPC錯誤
	CommonConvertError                      //資料格式錯誤
	CommonTokenValidError                   //token驗證錯誤
	CommonDataNotFoundError                 //查無資料
	CommonUploadError                       //上傳檔案錯誤
	CommonUploadDataError                   //上傳檔案內容錯誤
	DefaultFaultError       = 400000 + iota

	//通知失敗
	CommonNotifyError     = 500000 //TODO 目前更新機器人配置用
	CommonGameStatusError = 500001 //通知遊戲開關錯誤

	//權限驗證
	AuthGroupNotFoundError  = 600000
	InsufficientPermissions = 60001
)
