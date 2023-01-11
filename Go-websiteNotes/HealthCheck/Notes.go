// https://github.com/robzienert/http-healthcheck

package main 

import (

  "github.com/robzienert/http-healthcheck"
  "github.com/robzienert/http-healthcheck/monitor/cassandra"
  "github.com/gin-gonic/gin"


)

func GetHealthStatus(C *gin.Context) {
	status := healthcheck.FromContext(c).Status()
	resp := healthcheck.MarshalHealthStatusrResponse(status)
	if status.Healthy {
		c.IndentedJSON(http.StatusOk, resp)
	} else {
		c.IndentedJSON(http.StatusInternalServerError, resp)
	}
}



func main() {

	r := gin.Default()

	var healthProviders []healthcheck.Provider

	healthProviders = []healthcheck.Provider{
		cassandra.Newhealthprovider(gocqlSession),
	}

	healthMonitor := healthcheck.New(healthcheck.DefaultSupervisor, healthProviders...){
		defer healthMonitor.Close()
		healthMonitor.Start()
	}

	r.Get("/Status", GetHealthStatus)
	r.Run()
}


