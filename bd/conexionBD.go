package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.d2uia.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConectarBD permite la conexión a la base de datos */
func ConectarBD() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa a la BD")
	return client

}

/* ChequeoConnection permite la conexión a la base de datos */
func ChequeoConnection() int {

	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}

	return 1

}
