package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BrazenFox/WizardAppGraphQL/graph"
	"github.com/BrazenFox/WizardAppGraphQL/graph/generated"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8081"

/*func corss(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		handler(writer, request)
	}

}*/
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		//AllowedOrigins:   []string{"http://localhost:8081","http://localhost:3000"},
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods: []string{},
		AllowedHeaders: []string{},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
		return r.Host == "localhost:8081"
	}, ReadBufferSize: 1024, WriteBufferSize: 1024}})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground",  port)
	//log.Fatal(http.ListenAndServe(":"+ port, nil))
	err:=http.ListenAndServe(":8081", router)
	if err!=nil{
		panic(err)
	}

}
