package config

/*
You should touch a file: "$lark_bot/etc/lark_secret.yml"
	AppSecret : xxxxxxxxxxxxxxxxxxx
	EncryptKey: xxxxxxxxxxxxxxxxxxx
	VerificationToken: xxxxxxxxxxxxxxxxxxx
*/

var LSecret LarkSecret

type LarkSecret struct {
	AppSecret         string `yaml:"AppSecret"`
	EncryptKey        string `yaml:"EncryptKey"`
	VerificationToken string `yaml:"VerificationToken"`
}
