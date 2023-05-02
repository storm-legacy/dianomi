// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/tus/tusd/pkg/filestore"
// 	tusd "github.com/tus/tusd/pkg/handler"
// )

// func main() {
// 	// app := fiber.New()

// 	store := filestore.FileStore{
// 		Path: "./uploads",
// 	}

// 	composer := tusd.NewStoreComposer()
// 	store.UseIn(composer)

// 	// Configure tusd options
// 	handler, err := tusd.NewHandler(tusd.Config{
// 		BasePath:              "/files/",
// 		StoreComposer:         composer,
// 		NotifyCompleteUploads: true,
// 	})
// 	if err != nil {
// 		panic(fmt.Errorf("unable to create handler: %s", err))
// 	}

// 	go func() {
// 		for {
// 			event := <-handler.CompleteUploads
// 			fmt.Printf("Upload %s finished\n", event.Upload.ID)
// 		}
// 	}()

// 	http.Handle("/files/", http.StripPrefix("/files/", handler))
// 	err = http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		panic(fmt.Errorf("unable to listen: %s", err))
// 	}

// 	// app.Post("/files/", func(c *fiber.Ctx) error {
// 	// 	httpHandler := fasthttpadaptor.NewFastHTTPHandlerFunc(handler.ServeHTTP)
// 	// 	httpHandler(c.Context())
// 	// 	return nil
// 	// })
// 	// panic(app.Listen(":8080"))
// }

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		// 503 On vacation!
		return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
	})

	log.Fatal(app.Listen(":3000"))
}
