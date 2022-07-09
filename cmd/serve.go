package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"

	"github.com/rifatikbal/E-Com-PMS/internal/config"
	"github.com/rifatikbal/E-Com-PMS/internal/conn"
	productHttpDelivery "github.com/rifatikbal/E-Com-PMS/product/delivery"
	productRepo "github.com/rifatikbal/E-Com-PMS/product/repository"
	productUseCase "github.com/rifatikbal/E-Com-PMS/product/usecase"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve will serve the product apis",
	Long:  `serve will serve the product apis`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.LoadDBCfg()
		config.LoadAppCfg()
		if err := conn.ConnectDB(config.DB()); err != nil {
			log.Println(err)
			return err
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.App()
		productRepo := productRepo.New(conn.GetDB())
		productUC := productUseCase.New(productRepo)

		r := chi.NewRouter()

		apiRouter := chi.NewRouter()
		r.Mount("/api", apiRouter)

		productHttpDelivery.New(apiRouter, productUC)

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)

		server := http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler: r,
		}

		go func() {
			log.Println("server started on : 8081")
			if err := server.ListenAndServe(); err != nil {
				log.Println("info shutting down server")
			}
		}()
		<-quit

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
