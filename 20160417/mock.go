// テスト本体
func TestHoge(t *testing.T) {
	cfg = &TestEnv{ // ここを環境ごとに初期化できると良い
		MockCtrl: gomock.NewController(t),
	}
	initMock()
	// テスト対象関数の呼び出し
	...
}
func initMock() {
	// 先述の環境ごとの設定でテストの時だけモックを取得するようにする
	mock := cfg.GetKinesisService().(*awsmock.MockKinesisAPI)
	mock.EXPECT().CreateStream(
		... // 検証したい引数を定義
	).Return(
		... // 戻り値を定義
	)
}
func (e *TestEnv) GetKinesisService() kinesisiface.KinesisAPI {
	return awsmock.NewMockKinesisAPI(e.MockCtrl)
}
