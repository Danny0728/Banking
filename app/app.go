package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Danny0728/BankAPI/domain"
	"github.com/Danny0728/BankAPI/logger"
	"github.com/Danny0728/BankAPI/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	keys := []string{"SERVER_ADDRESS", "SERVER_PORT", "DB_USER", "DB_PASSWD", "DB_ADDR", "DB_PORT", "DB_NAME"}
	for i := 0; i < len(keys); i++ {
		if value, exists := os.LookupEnv(keys[i]); !exists || value == "" {
			logger.Error(fmt.Sprintf("%s doesn't exists or isBlank", keys[i]))
			os.Exit(0)
		}
	}
}
func Start() {

	sanityCheck()
	//creating router
	router := mux.NewRouter()

	//wiring of the whole architecture
	dbClient := getDbClient()
	customerRespositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRespositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRespositoryDb)}
	ah := AccountHandlers{service.NewAccountService(accountRespositoryDb)}
	//routes
	//customerHandlers
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//AccountHandlers
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	//TransactionHandler
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}/transaction", ah.MakeTransaction).Methods(http.MethodPost)

	//server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	logger.Info(fmt.Sprintf("Starting server on %s:%s...", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
func End() {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			case syscall.SIGINT:
				logger.Info("Ending Banking Application...")
				os.Exit(0)
			default:
				logger.Info("Application Ended Abruptly")
				os.Exit(1)
			}
		}
	}()
}
func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	err = client.Ping()
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
