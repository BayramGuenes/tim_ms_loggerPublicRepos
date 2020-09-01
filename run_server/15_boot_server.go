package run_server

import (
	"net/http"
	"tim_presse/tim_ms_loggerPublicRepos/utils"

	"github.com/gin-gonic/gin"
)

func UseGinAsRouter() {
	Router = gin.New()

}

//https://github.com/gin-gonic/gin/issues/1799
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		//fmt.utils.StdOut(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			//fmt.utils.StdOut("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func ListenAndServe() {
	utils.StdOut("Listening on " + TimMicroservices.NameTimLogServer + ":" + TimMicroservices.PortTimLogServer)
	http.ListenAndServe(":"+TimMicroservices.PortTimLogServer, Router)
	/*utils.StdOut("utils.StdOut: Listening on :4080")
	log.Printf("log.Printf: Listening on :4080")
	http.ListenAndServe(":4080", Router)*/
}
