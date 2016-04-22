func main() {

	cfg := profile.Config{
		MemProfile:     true,
		CPUProfile:     true,
		NoShutdownHook: false,
	}

	// deferなのでmain関数の最後まで到達すればプロファイル結果が出力される。
	// Web APIなどOSのシグナルで終了されmain関数が最後まで到達しない可能性が高いものは、
	// 上記のNoShutdownHookをfalseにする
	defer profile.Start(&cfg).Stop()

	// webサーバスタート
	...
}
