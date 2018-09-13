# Tooling

## Comandos útiles
- Para construir para el sistema operativo actual
$ go build

- Para construir para otro OS (windows en este caso)
$ GOOS=windows go build

- Para instalar (en $GOPATH/bin/nombre) el binario ejecutable
$ go install

- Dado un paquete, lo busca, lo descarga, lo compila y lo instala en bin (obviamente dentro de gopath)
$ go get https://github.com/example/name

- Para mostrar el origen del paquete actual
$ go list

- Podemos obtener infomración expecífica
$ go list -f '{{ .Name }}: {{ .Doc }}'

- Puedo ver los imports del paquete actual
$ go list -f '{{ .Imports }}'

- Puedo ver los imports de un paquete específico
$ go list -f '{{ .Imports }}' fmt
$ go list -f '{{ join .Imports "\n" }}' fmt

- Para obtener la documentación
$ go doc
$ go doc fmt
$ go doc fmt Printf

- Podemos publicar toda la documentación q tenemos
$ godoc -http :6060


