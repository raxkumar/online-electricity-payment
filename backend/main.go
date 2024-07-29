package main

import (
	"net/http"

	auth "com.electricity.online/auth"
	app "com.electricity.online/config"
	"com.electricity.online/controllers"
	config "com.electricity.online/db"
	eureka "com.electricity.online/eurekaregistry"
	"com.electricity.online/iam"
	"com.electricity.online/migrate"
	"github.com/asim/go-micro/v3"
	mhttp "github.com/go-micro/plugins/v3/server/http"
	"github.com/go-micro/plugins/v3/wrapper/monitoring/prometheus"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/micro/micro/v3/service/logger"
)

var configurations eureka.RegistrationVariables

func main() {
	app.Setconfig()
	migrate.MigrateAndCreateDatabase()

	// should we throughout an error if the keycloak is not running!
	auth.SetClient()

	config.InitializeDb()

	port := app.GetVal("GO_MICRO_SERVICE_PORT")
	srv := micro.NewService(
		micro.Server(mhttp.NewServer()),
	)
	microserviceOptions := []micro.Option{
		micro.Name("backend"),
		micro.Version("latest"),
		micro.Address(":" + port),

		// not working, need to check
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	}
	srv.Init(microserviceOptions...)

	// change the r to router, proper naming convention
	router := mux.NewRouter().StrictSlash(true)

	router.Use(httpHeadersMiddleware)
	registerRoutes(router)

	var handlers http.Handler = router

	// Jhipster-registry
	service_registry_url := app.GetVal("GO_MICRO_SERVICE_REGISTRY_URL")
	InstanceId := "backend:" + uuid.New().String()
	configurations = eureka.RegistrationVariables{ServiceRegistryURL: service_registry_url, InstanceId: InstanceId}
	go eureka.ManageDiscovery(configurations)

	if err := micro.RegisterHandler(srv.Server(), handlers); err != nil {
		logger.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}

	// Defer the cleanup function
	defer eureka.Cleanup(configurations)
}

func registerRoutes(router *mux.Router) {
	registerControllerRoutes(controllers.CommunicationController{}, router)
	registerControllerRoutes(controllers.EventController{}, router)
	registerControllerRoutes(controllers.ManagementController{}, router)
	registerControllerRoutes(controllers.NoteController{}, router)
	registerControllerRoutes(controllers.TestController{}, router)
	registerControllerRoutes(iam.IAMController{}, router)
	registerControllerRoutes(controllers.UserController{}, router)
	registerControllerRoutes(controllers.MonitoringController{}, router)

}

func registerControllerRoutes(controller controllers.Controller, router *mux.Router) {
	controller.RegisterRoutes(router)
}

// need to abstract the allow-origin to config file [ENV], rather then using '*'
// rename the method to something more generic
func httpHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		if request.Method == "OPTIONS" {
			writer.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
