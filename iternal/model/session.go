package model

type Session struct {
	Token    string
	LifeTime int64
	Owner    string
	FingerPrint
}

type FingerPrint struct {
	Addr    string
	Browser string
	Device  string
}
