package go_say_hello


//git tag untuk rilis versi module
// git tag v1.1.0
// git push origin v1.1.0

//jika terjadi major changed, maka ubah lah nama module, biasanya dikasih /v2
// dan tag version diubdah drastis misalnya dari 1 menjadi 2
func SayHello(name string) string  {
	return "Hello brad" + name
}