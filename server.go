package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BrazenFox/WizardAppGraphQL/graph"
	"github.com/BrazenFox/WizardAppGraphQL/graph/generated"
	"github.com/go-chi/chi"
	"github.com/magiconair/properties"
	"github.com/rs/cors"
	"log"
	"net/http"
)

/*func corss(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		handler(writer, request)
	}

}*/
func GetPort() string {
	p := properties.MustLoadFile("config.conf", properties.UTF8)
	return p.GetString("port", "8081")

}
func GetIncomingUrl() string {
	p := properties.MustLoadFile("config.conf", properties.UTF8)
	return p.GetString("incomingurl", "http://localhost:3000")

}
func main() {

	port := GetPort()
	incomingurl := GetIncomingUrl()

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:" + port, "http://localhost:3000", "http://192.168.99.102:" + port, "http://192.168.99.102:3000", incomingurl},
		//AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{},
		AllowedHeaders:   []string{},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	/*srv.AddTransport(&transport.Websocket{
	Upgrader: websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return r.Host == "localhost:"+port
		}, ReadBufferSize: 1024, WriteBufferSize: 1024}})*/

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	//log.Fatal(http.ListenAndServe(":"+ port, nil))
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}

}
