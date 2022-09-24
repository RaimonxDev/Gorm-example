package main

import (
	"github.com/RaimonxDev/Gorm-example/model"
	"github.com/RaimonxDev/Gorm-example/storage"
)

func main() {
	storage.New()
	// creamos las tablas en la bases de datos con AutoMigrate
	//storage.DB().AutoMigrate(
	//	&model.Product{},
	//	&model.InvoiceHeader{},
	//	&model.InvoiceItem{},
	//)

	//	CREACION DE REGISTRO EN LA BASES DE DATOS
	//product1 := model.Product{
	//	Name:  "Curso de go",
	//	Price: 120,
	//}
	// Observaciones necesita recibir un puntero, para eso creamos la variable, y luego se asigna
	// A observations y asi no tener errores con los campos not null
	//obsProduct2 := "Curso con descuento"
	//product2 := model.Product{
	//	Name:         "Curso de go avanzado",
	//	Price:        120,
	//	Observations: &obsProduct2,
	//}
	//storage.DB().Create(&product1)
	//storage.DB().Create(&product2)

	// ----------------------------------------------------------------------

	//Buscar todos los  PRODUCTS

	//creamos la variable que almacena los datos de DB
	//products := make([]model.Product, 0)

	//storage.DB().Find(&products)
	//
	//for _, product := range products {
	//	fmt.Printf("%d - %s\n", product.ID, product.Name)
	//}

	//	------------------------------------------------------------------

	//Buscar un producto por ID
	//myProduct := model.Product{}
	//
	//storage.DB().First(&myProduct, 1)
	//fmt.Println(myProduct)

	//transaciones

	invoice := model.InvoiceHeader{
		Client: "Ramon Martinez",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 60},
		},
	}

	storage.DB().Create(&invoice)

}
