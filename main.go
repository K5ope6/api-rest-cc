package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Estructuras de Tablas
type Alumno struct {
	IdAlumno  int    `json:"idAlumno"`
	Nombre    string `json:"nombre"`
	ApellidoP string `json:"apellidoP"`
	ApellidoM string `json:"apellidoM"`
	Matricula string `json:"matricula"`
	Correo    string `json:"correo"`
	Carrera   string `json:"carrera"`
	Semestre  string `json:"semestre"`
}

type Creditos struct {
	IdCredito     int    `json:"idCredito"`
	NombreCredito string `json:"nombreCredito"`
	Estado        string `json:"estado"`
}

type AlumnosCreditos struct {
	IdAlumno  int `json:"idAlumno"`
	IdCredito int `json:"idCredito"`
}

// Conexion BD
func conexionBD() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName)
	return sql.Open("mysql", dsn)

}

// Main
func main() {
	router := gin.Default() //Instancia Gin
	db, err := conexionBD() //llamada a funcion

	if err != nil {
		log.Fatal(err) //Captura el error si no se establece conexion
	}

	defer db.Close() //Cerrar conexion

	if err := db.Ping(); err != nil { //Verifica si da ping (conexion)
		log.Fatal(err)

	}

	//POST ALUMNO
	router.POST("/alumnos", func(ctx *gin.Context) { //Definir ruta HTTP, función que permite enviar respuestas
		var alumno Alumno //Variable tipo Alumno (Estructura)
		if err := ctx.ShouldBindJSON(&alumno); err != nil {
			//Los datos enviados en la solicitud se convierten en la estructura Alumno
			ctx.JSON(400, gin.H{"error": err.Error()}) //Manejo de error
			return
		}
		//Ejecuta una consulta SQL
		_, err := db.Exec("INSERT INTO Alumno (idAlumno, nombre, apellidoP, apellidoM, matricula, correo, carrera, semestre) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			alumno.IdAlumno, alumno.Nombre, alumno.ApellidoP, alumno.ApellidoM, alumno.Matricula, alumno.Correo, alumno.Carrera, alumno.Semestre)

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(201, gin.H{"message": "Alunmo creado con exito"})

	})

	//PUT ALUMNO
	router.PUT("/alumnos/:idAlumno", func(ctx *gin.Context) { //Ruta
		var alumno Alumno                 //variable de tipo Alumno, almacena información de un alumno
		idAlumno := ctx.Param("idAlumno") //Obtener id del alumno

		if err := ctx.ShouldBindJSON(&alumno); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()}) //si hay un error devuleve el código de estado 400
			return

		}
		_, err := db.Exec("UPDATE Alumno SET idAlumno=?, nombre=?, apellidoP=?, apellidoM=?, matricula=?, correo=?, carrera=?, semestre=? WHERE idAlumno=?",
			alumno.IdAlumno, alumno.Nombre, alumno.ApellidoP, alumno.ApellidoM, alumno.Matricula, alumno.Correo, alumno.Carrera, alumno.Semestre, idAlumno)

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return

		}
		ctx.JSON(200, gin.H{"message": "Alumno actualizado con exito"})

	})

	//GET ALUMNO
	router.GET("/alumnos", func(ctx *gin.Context) { //Maneja solicitudes GET en alumnos /alumnos
		//Consulta a una tabla
		rows, err := db.Query("SELECT * FROM Alumno")

		if err != nil { //Verificar si ocurrio un error
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		defer rows.Close()   //Cierra al terminar de procear las filas
		var alumnos []Alumno //Almacenar los registros de alumnos de la BD

		for rows.Next() { //Itera los registros
			var alumno Alumno

			if err := rows.Scan(&alumno.IdAlumno, &alumno.Nombre, &alumno.ApellidoP, &alumno.ApellidoM, &alumno.Matricula,
				&alumno.Correo, &alumno.Carrera, &alumno.Semestre); err != nil { //Leer los datos de cada fila y asignandolos a cada variable dentro de la estructura.
				ctx.JSON(500, gin.H{"error": err.Error()}) //Si ocurre un error al leer los datos
				return

			}

			alumnos = append(alumnos, alumno) //Agrega al alumno

		}

		ctx.JSON(200, alumnos) //Al final de leer las filas da un OK y se envia los resultados en formato JSON
	})

	//DELETE ALUMNO
	router.DELETE("/alumnos/:idAlumno", func(ctx *gin.Context) {
		idAlumno := ctx.Param("idAlumno")

		_, err := db.Exec("DELETE FROM Alumno WHERE idAlumno=?", idAlumno)

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"message": "Alumno eliminado con exito"})

	})

	//POST CREDITO
	router.POST("/creditos", func(ctx *gin.Context) {
		var credito Creditos

		if err := ctx.ShouldBindJSON(&credito); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if credito.Estado != "APROBADO" && credito.Estado != "NO APROBADO" {
			ctx.JSON(400, gin.H{"error": "Estado deber ser APROBADO o NO APROBADO"})
			return
		}

		_, err := db.Exec("INSERT INTO Creditos (idCredito, nombreCredito, estado) VALUES (?, ?, ?)",
			credito.IdCredito, credito.NombreCredito, credito.Estado)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"message": "Crédito creado con éxito"})
	})

	//GET CREDITOS
	router.GET("/creditos", func(ctx *gin.Context) {
		rows, err := db.Query("SELECT * FROM Creditos")

		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var creditos []Creditos
		for rows.Next() {
			var credito Creditos
			if err := rows.Scan(&credito.IdCredito, &credito.NombreCredito, &credito.Estado); err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
			creditos = append(creditos, credito) //Agrega el Credito
		}

		ctx.JSON(200, creditos)
	})

	//PUT CREDITOS
	router.PUT("/creditos/:idCredito", func(ctx *gin.Context) {
		var credito Creditos
		idCredito := ctx.Param("idCredito")
		if err := ctx.ShouldBindJSON(&credito); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("UPDATE Creditos SET nombreCredito=?, estado=? WHERE idCredito=?",
			credito.NombreCredito, credito.Estado, idCredito)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"message": "Crédito actualizado con éxito"})
	})

	//DELETE CREDITOS
	router.DELETE("/creditos/:idCredito", func(ctx *gin.Context) {
		idCredito := ctx.Param("idCredito")

		_, err := db.Exec("DELETE FROM Creditos WHERE idCredito=?", idCredito)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"message": "Crédito eliminado con éxito"})
	})

	//POST AlumnoCredito
	router.POST("/alumnocreditos", func(ctx *gin.Context) { //Se define la URL y ponemos una función para consultar el cuerpo del JSON
		var alumnoCredito AlumnosCreditos                                  //Variable alumnoCredito tipo estructura AlumnosCreditos
		if err := ctx.ShouldBindBodyWithJSON(&alumnoCredito); err != nil { //Se guarda el obejto con los datos en alumnocredito
			ctx.JSON(400, gin.H{"error": err.Error()}) //Captura un error si el valor es diferente de nil(nulo)
			return
		}
		_, err := db.Exec("INSERT INTO AlumnosCreditos(idAlumno, idCreditos) VALUES(?, ?)",
			alumnoCredito.IdAlumno, alumnoCredito.IdCredito)

		if err != nil {
			ctx.JSON(500, gin.H{"Error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"message": "Registro creado con exito AlumnoCreditos"})
	})

	//GET AlumnoCredito
	router.GET("/alumnocreditos/:idAlumno/:idCredito", func(ctx *gin.Context) {
		var alumnoCredito AlumnosCreditos
		idAlumno := ctx.Param("idAlumno")
		idCredito := ctx.Param("idCredito")

		row := db.QueryRow("SELECT idAlumno, idCredito FROM AlumnoCreditos WHERE idAlumno=? AND idCredito=?", idAlumno, idCredito)

		if err := row.Scan(&alumnoCredito.IdAlumno, &alumnoCredito.IdCredito); err != nil {
			ctx.JSON(404, gin.H{"error": "Registro  no encintrado"})
			return
		}
		ctx.JSON(200, alumnoCredito)
	})

	//GET AlumnoCredito todos los registros
	router.GET("/alumnocreditos", func(ctx *gin.Context) {
		rows, err := db.Query("SELECT idAlumno, idCredito FROM AlumnosCreditos")
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return

		}
		defer rows.Close()

		var listaAlumnoCredito []AlumnosCreditos
		for rows.Next() {
			var alumnoCredito AlumnosCreditos
			if err := rows.Scan(&alumnoCredito.IdAlumno, &alumnoCredito.IdCredito); err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
			listaAlumnoCredito = append(listaAlumnoCredito, alumnoCredito)
		}
		ctx.JSON(200, listaAlumnoCredito)

	})

	//PUT AlumnoCredito
	router.PUT("/alumnocreditos/:idAlumno/:idCredito", func(ctx *gin.Context) {
		//var alumnoCredito AlumnosCreditos
		//idAlumno := ctx.Param("idAlumno")
		//idCredito := ctx.Param("idCredito")

		//if err := ctx.ShouldBindBodyWithJSON(&alumnoCredito); err != nil {
		ctx.JSON(400, gin.H{"error": "No hay nada que actualizar"})
		//	return
		//}

		//_, err := db.Exec("UPDATE AlumnoCreditos SET cantidadCreditos =? WHERE idAlumno=? AND idCredito=?", idAlumno, idCredito)
		//if err != nil {
		//	ctx.JSON(500, gin.H{"error": err.Error()})
		//	return
		//}

		//ctx.JSON(200, gin.H{"message": "Creditos actualizados con exito"})

	})

	router.DELETE("/alumnoCreditos/:idAlumno/:idCredito", func(ctx *gin.Context) {
		idAlumno := ctx.Param("idAlumno")
		idCredito := ctx.Param("idCredito")

		_, err := db.Exec("DELETE FROM AlumnoCreditos WHERE idAlumno=? AND idCredito=?", idAlumno, idCredito)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"message": "Registro eliminado con exito"})

	})

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Bienvenido a Api-Rest"})
	}) //Ruta básica

	router.Run(":8080")

}
