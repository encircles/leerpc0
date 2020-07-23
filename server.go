package leerpc0

type Server struct {
	opts     *ServerOptions
	services map[string]Service
}
