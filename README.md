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

- Para verificar donde tenemos algún error no manejado (primero obtenerlo)
$ go get -u github.com/kisielk/errcheck
$ errcheck

- Para verificar código que podría estar mal
$ go vet

- Para ver performance go-wrk (carga)
$ go-wrk -d 5 http://localhost:8087/ariel@golang.org

- Para hacer profiling
importamos
_ "net/http/pprof"

luego hacemos correr la carga y probamos:
http://127.0.0.1:8087/debug/pprof/

Para analizar esos datos usamos (las pruebas de carga deben estar corriendo para que se haga el análisis):
$ go tool pprof -seconds=5 http://localhost:8087/debug/pprof/profile

Cuando aparezca la consola, podemos ver un listado:
(pprof) top

O mejor aún, verlo en la web:
(pprof) web

- Para ver flamegraphs
$ go get github.com/uber/go-torch

- Creamos benchamarks en el archivo de test y luego podemos hacerlo correr:
Saquemos tiempos promedio por operación
$ go test -bench .

- Hagamos algo más con los benchmarks
Saquemos un profile del cpu
$ go test -bench . -cpuprofile prof.cpu

Para nuestro caso, hacemos correr con go-torch:
$ go-torch --binaryname <nombre-archivo-test> -b <nombre-profile>
$ go-torch --binaryname tooling.test -b prof.cpu

