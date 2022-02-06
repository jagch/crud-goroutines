package bd

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"github.com/jagch/crud-goroutines/configuracion"

	_ "github.com/lib/pq"
)

//BD variable global que tiene la conexión a la DB.
var BD = Conectar()
var once sync.Once
var bd *sql.DB

//Conectar establece la conexión.
func Conectar() *sql.DB {
	var err error
	once.Do(func() { //singleton: se ejecuta una única vez
		dsn := "postgres://" + configuracion.Configuracion("DB_USER") + ":" + configuracion.Configuracion("DB_PASSWORD") + "@" + configuracion.Configuracion("DB_HOST") + ":" + configuracion.Configuracion("DB_PORT") + "/" + configuracion.Configuracion("DB_NAME") + "?sslmode=" + configuracion.Configuracion("DB_SSL_MODE") + ""
		//log.Println(dsn)

		bd, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Fatal(err)
		}
	})
	//actualmente las líneas siguientes ralentizan las consultas, sin conocer el porqué aún
	//bd.SetMaxIdleConns(0)//no limit in the pool
	//bd.SetMaxOpenConns(0)//no limit
	return bd
}

//VerificarConexion hace ping a la DB.
func VerificarConexion() bool {

	err := BD.PingContext(context.TODO())
	if err != nil {
		return true
	}
	return false
}
