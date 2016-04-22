// 標準のエラーをwrapする
if err != nil {
	return errors.Wrap(err)
}
// 自分でエラーを定義する
if num < 10 {
	return errors.Wrap("numが10未満だよん")
} else {
	// 書式付きメッセージもできる
	return errors.Errorf("残念 %d もダメでしたー", num)
}
