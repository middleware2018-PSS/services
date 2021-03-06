package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/middleware2018-PSS/Services/docs"
	"github.com/middleware2018-PSS/Services/models"
	"github.com/middleware2018-PSS/Services/repository"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/middleware2018-PSS/Services/controller"
	"github.com/phisco/hal"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	LimitError = errors.New("Limit Must Be Greater Than Zero.")
)

type (
	JwtData struct {
		Id   int
		Role string
	}
)

const (
	HAL  = "application/hal+json"
	COST = 10
)

// @title Back2School API
// @version 1.0byID
// @description These are a School management system's API .
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @host localhost:5000
func main() {
	/* p, _ := bcrypt.GenerateFromPassword([]byte("password"), 10)
	fmt.Printf("%s", p) */
	viper.SetConfigName("config")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Check database connection string
	log.Println(viper.GetString("DBparams"))

	db, err := sql.Open("postgres", viper.GetString("DBparams"))
	// err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	con := controller.NewController(repository.NewPostgresRepository(db))

	authMiddleware := jwt.GinJWTMiddleware{
		Realm:            viper.GetString("realm"),
		SigningAlgorithm: "RS256",
		PrivKeyFile:      viper.GetString("PrivateKeyFile"),
		PubKeyFile:       viper.GetString("PublicKeyFile"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		Authenticator: func(userID string, password string, c *gin.Context) (interface{}, bool) {
			id, kind, ok := con.CheckUser(userID, password)
			return JwtData{id, kind}, ok
		},
		PayloadFunc: func(data interface{}) map[string]interface{} {
			return map[string]interface{}{
				repository.USER: data.(JwtData).Id,
				repository.KIND: data.(JwtData).Role,
			}
		},
	}

	g := gin.Default()

	// If you receive a GET to root, redirect to swagger docs
	g.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, c.Request.URL.Host+"swagger/index.html")
		//c.Request.URL.Path = "/swagger/index.html"
		//g.HandleContext(c)
	})

	g.POST("/login", LoginHandler(&authMiddleware))

	api := g.Group("", authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		c.Set(repository.USER, claims[repository.USER])
		c.Set(repository.KIND, claims[repository.KIND])
	})
	api.GET("/accounts", getOffsetLimit(con.Accounts))
	api.POST("/accounts", func(c *gin.Context) {
		var user models.User
		if err := c.Bind(&user); err == nil {
			_, whokind := idKind(c)
			if err := con.CreateAccount(user.Username, user.Password, user.ID, user.Role, COST, whokind); err != nil {
				c.AbortWithStatus(http.StatusNotAcceptable)
			} else {
				c.AbortWithStatus(http.StatusCreated)
			}
		}

	})
	api.PUT("/accounts", func(c *gin.Context) {
		var user models.Account
		if err := c.Bind(&user); err == nil {
			who, whokind := idKind(c)
			if err := con.UpdateAccount(user, who, whokind, COST); err != nil {
				c.AbortWithStatus(http.StatusNotAcceptable)
			} else {
				c.AbortWithStatus(http.StatusCreated)
			}
		}

	})
	api.DELETE("/accounts/{id}", func(c *gin.Context) {
		user := c.Param("id")
		_, whoKind := idKind(c)
		res, err := con.DeleteAccount(user, whoKind)
		res, _ = ToRepresentation(res, c, checkHAL(c))
		handleErr(err, res, c)
	})
	api.GET("/refresh_token", authMiddleware.RefreshHandler)

	api.POST("/parents", func(c *gin.Context) {
		var p models.Parent
		if err := c.ShouldBind(&p); err == nil {
			who, whoKind := idKind(c)
			if id, err := con.CreateParent(p, who, whoKind); err == nil {
				p.ID = id
				c.JSON(http.StatusCreated, p)
			}
		}
	})

	parent := api.Group("/parents/:id") //, authAdminOrParent(authMiddleware.Realm))
	{
		parent.GET("", byID("id", con.ParentByID))
		parent.PUT("", func(c *gin.Context) {
			// not possible to refactor (at the best of my knowledge)
			var p models.Parent
			if err := c.ShouldBind(&p); err == nil {
				id, _ := strconv.Atoi(c.Param("id"))
				p.ID = id
				who, whoKind := idKind(c)
				if err := con.UpdateParent(p, who, whoKind); err == nil {
					c.JSON(http.StatusNoContent, p)
				}
			}
		})
		parent.DELETE("", byID("id", con.DeleteParent))
		parent.GET("/students", byIDWithOffsetAndLimit("id", con.ChildrenByParent))
		parent.GET("/appointments", byIDWithOffsetAndLimit("id", con.AppointmentsByParent))
		parent.POST("/appointments", func(c *gin.Context) {
			var a models.Appointment
			if err := c.ShouldBind(&a); err == nil {
				// TODO check parent is same parent of the appointment => isParent student
				who, whoKind := idKind(c)
				if id, err := con.CreateAppointment(a, who, whoKind); err == nil {
					a.ID = id
					c.JSON(http.StatusCreated, a)
				}
			}

		})
		parent.GET("/payments", byIDWithOffsetAndLimit("id", con.PaymentsByParent))
		parent.GET("/notifications", byIDWithOffsetAndLimit("id", con.NotificationsByParent))
	}

	// TODO add hypermedia

	api.POST("/teachers", func(c *gin.Context) {
		var t models.Teacher
		if err := c.ShouldBind(&t); err == nil {
			who, whoKind := idKind(c)
			if id, err := con.CreateTeacher(t, who, whoKind); err != nil {
				t.ID = id
				c.JSON(http.StatusCreated, t)
			}
		}
	})
	api.GET("/teachers", getOffsetLimit(con.Teachers))
	teachers := api.Group("/teachers/:id") //, authAdminOrTeacher(authMiddleware.Realm))
	{
		teachers.GET("", byID("id", con.TeacherByID))
		teachers.PUT("", func(c *gin.Context) {
			var t models.Teacher
			if err := c.ShouldBind(&t); err == nil {
				id, _ := strconv.Atoi(c.Param("id"))
				t.ID = id
				who, whoKind := idKind(c)
				if err := con.UpdateTeacher(t, who, whoKind); err != nil {
					c.JSON(http.StatusNoContent, nil)
				}
			}
		})
		teachers.DELETE("", byID("id", con.DeleteTeacher))
		teachers.GET("/lectures", byIDWithOffsetAndLimit("id", con.LecturesByTeacher))
		teachers.GET("/appointments", byIDWithOffsetAndLimit("id", con.AppointmentsByTeacher))
		teachers.GET("/notifications", byIDWithOffsetAndLimit("id", con.NotificationsByTeacher))
		teachers.GET("/subjects", byIDWithOffsetAndLimit("id", con.SubjectsByTeacher))
		teachers.GET("/subjects/:subject", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			subj := c.Param("subject")
			offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
			limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
			who, whoKind := idKind(c)
			res, err := con.ClassesBySubjectAndTeacher(id, subj, limit, offset, who, whoKind)
			handleErr(err, res, c)
		})
		teachers.GET("/classes", byIDWithOffsetAndLimit("id", con.ClassesByTeacher))
	}

	api.GET("/appointments" /*authAdmin(authMiddleware.Realm),*/, getOffsetLimit(con.Appointments))
	api.GET("/appointments/:appointment", byID("appointment", con.AppointmentByID))
	api.PUT("/appointments/:appointment", func(c *gin.Context) {
		var a models.Appointment
		if err := c.ShouldBind(&a); err == nil {
			id, _ := strconv.Atoi(c.Param("id"))
			a.ID = id
			who, whoKind := idKind(c)
			if err := con.UpdateAppointment(a, who, whoKind); err == nil {
				c.JSON(http.StatusCreated, a)
			}
		}
	})
	api.DELETE("/appointments/:appointment", byID("appointment", con.DeleteAppointment))

	api.GET("/lectures", getOffsetLimit(con.Lectures))
	api.GET("/lectures/:id", byID("id", con.LectureByID))
	api.POST("/lectures", func(c *gin.Context) {
		var s models.TimeTable
		if err := c.ShouldBind(&s); err == nil {
			who, whoKind := idKind(c)
			if id, err := con.CreateLecture(s, who, whoKind); err != nil {
				s.ID = id
				c.JSON(http.StatusCreated, s)
			}
		}
	})
	api.PUT("/lectures/:id", func(c *gin.Context) {
		var a models.TimeTable
		if err := c.ShouldBind(&a); err == nil {
			id, _ := strconv.Atoi(c.Param("id"))
			a.ID = id
			who, whoKind := idKind(c)
			if err := con.UpdateLecture(a, who, whoKind); err == nil {
				c.JSON(http.StatusCreated, a)
			}
		}
	})
	api.DELETE("/lectures/:id", byID("id", con.DeleteLecture))

	api.GET("/parents", getOffsetLimit(con.Parents))
	api.GET("/grades", getOffsetLimit(con.Grades))
	api.GET("/grades/:id", byID("id", con.GradeByID))
	api.DELETE("/grades/:id", byID("id", con.DeleteGrade))
	api.PUT("/grades/:id", func(c *gin.Context) {
		var a models.Grade
		if err := c.ShouldBind(&a); err == nil {
			id, _ := strconv.Atoi(c.Param("id"))
			a.ID = id
			who, whoKind := idKind(c)
			if err := con.UpdateGrade(a, who, whoKind); err == nil {
				c.JSON(http.StatusCreated, a)
			}
		}
	})
	api.GET("/students", getOffsetLimit(con.Students))
	api.PUT("/students/:id", func(c *gin.Context) {
		var a models.Student
		if err := c.ShouldBind(&a); err == nil {
			id, _ := strconv.Atoi(c.Param("id"))
			a.ID = id
			who, whoKind := idKind(c)
			if err := con.UpdateStudent(a, who, whoKind); err == nil {
				c.JSON(http.StatusCreated, a)
			}
		}
	})
	api.DELETE("/students/:id", byID("id", con.DeleteStudent))
	api.POST("/students", func(c *gin.Context) {
		var s models.Student
		if err := c.ShouldBind(&s); err == nil {
			who, whoKind := idKind(c)
			if id, err := con.CreateStudent(s, who, whoKind); err != nil {
				s.ID = id
				c.JSON(http.StatusCreated, s)
			}
		}
	})
	api.GET("/students/:id", byID("id", con.StudentByID))
	api.GET("/students/:id/grades", byIDWithOffsetAndLimit("id", con.GradesByStudent))
	api.GET("/notifications", getOffsetLimit(con.Notifications))
	api.POST("/notifications", func(c *gin.Context) {
		var s models.Notification
		if err := c.ShouldBind(&s); err == nil {
			who, whoKind := idKind(c)
			if id, err := con.CreateNotification(s, who, whoKind); err != nil {
				s.ID = id
				c.JSON(http.StatusCreated, s)
			}
		}
	})
	api.GET("/notifications/:id", byID("id", con.NotificationByID))
	api.DELETE("/notifications/:id", byID("id", con.DeleteNotification))
	api.PUT("/notifications/:id", func(c *gin.Context) {
		var a models.Notification
		if err := c.ShouldBind(&a); err == nil {
			id, _ := strconv.Atoi(c.Param("id"))
			a.ID = id
			who, whoKind := idKind(c)
			if err := con.UpdateNotification(a, who, whoKind); err == nil {
				c.JSON(http.StatusCreated, a)
			}
		}
	})
	api.GET("/payments", getOffsetLimit(con.Payments))
	api.POST("/payments", func(c *gin.Context) {
		var s models.Payment
		if err := c.ShouldBind(&s); err == nil {
			who, whoKind := idKind(c)
			if id, err := con.CreatePayment(s, who, whoKind); err != nil {
				s.ID = id
				c.JSON(http.StatusCreated, s)
			}
		}
	})
	api.GET("/payments/:id", byID("id", con.PaymentByID))
	api.PUT("/payments/:id", func(c *gin.Context) {
		var a models.Payment
		if err := c.ShouldBind(&a); err == nil {
			id, _ := strconv.Atoi(c.Param("id"))
			a.ID = id
			who, whoKind := idKind(c)
			if err := con.UpdatePayment(a, who, whoKind); err == nil {
				c.JSON(http.StatusCreated, a)
			}
		}
	})
	api.DELETE("/payments/:id", byID("id", con.DeletePayment))

	api.GET("/classes", getOffsetLimit(con.Classes))
	api.GET("/classes/:id", byID("id", con.ClassByID))
	api.PUT("/classes/:id", func(c *gin.Context) {
		var a models.Class
		if err := c.ShouldBind(&a); err == nil {
			id, _ := strconv.Atoi(c.Param("id"))
			a.ID = id
			who, whoKind := idKind(c)
			if err := con.UpdateClass(a, who, whoKind); err == nil {
				c.JSON(http.StatusCreated, a)
			}
		}
	})
	api.DELETE("/classes/:id", byID("id", con.DeleteClass))

	api.GET("/classes/:id/students", byIDWithOffsetAndLimit("id", con.StudentsByClass))
	api.POST("/classes", func(c *gin.Context) {
		var s models.Class
		if err := c.ShouldBind(&s); err == nil {
			who, whoKind := idKind(c)
			if id, err := con.CreateClass(s, who, whoKind); err != nil {
				s.ID = id
				c.JSON(http.StatusCreated, s)
			}
		}
	})

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.Run(":5000")
}

// UTIL FUNCTIONS

func byID(key string, f func(int, int, string) (interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		// TODO: check err
		id, err := strconv.Atoi(c.Param(key))
		who, whoKind := idKind(c)
		res, err := f(id, who, whoKind)
		res, _ = ToRepresentation(res, c, checkHAL(c))
		handleErr(err, res, c)
	}

}

func offsetLimit(c *gin.Context) (int, int) {
	// TODO check err
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	return offset, limit
}

func idKind(c *gin.Context) (int, string) {
	who := int(c.MustGet(repository.USER).(float64))
	kind := c.MustGet(repository.KIND).(string)
	return who, kind
}

func getOffsetLimit(f func(int, int, int, string) ([]interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		//TODO Check id and errors
		offset, limit := offsetLimit(c)
		who, whoKind := idKind(c)
		if limit > 0 {
			res, err := f(limit, offset, who, whoKind)
			if err == nil {
				h := checkHAL(c)
				for i, el := range res {
					res[i], _ = ToRepresentation(el, c, h)
				}
				if h {
					halList(c, &res, err, next(c.Request.RequestURI, offset, limit, res), prev(c.Request.RequestURI, offset, limit))
				} else {
					handleErr(err, &res, c)
				}
			} else {
				handleErr(repository.ErrNoResult, nil, c)
			}
		} else {
			handleErr(LimitError, nil, c)
		}

	}
}

func halList(c *gin.Context, res *[]interface{}, err error, next string, prev string) {
	halRes := hal.NewResource(struct{}{}, c.Request.RequestURI)
	uri := strings.Split(c.Request.RequestURI, "/")
	embeddedLink := uri[len(uri)-1]
	for _, el := range *res {
		// el.(*hal.Resource).Links = hal.LinkRelations{"self": el.(*hal.Resource).Links["self"]}
		halRes.Embed(hal.Relation(embeddedLink), el.(*hal.Resource))
	}

	if next != "" {
		halRes.AddNewLink("next", next)
	}
	if prev != "" {
		halRes.AddNewLink("prev", prev)
	}
	handleErr(err, halRes, c)
}

func byIDWithOffsetAndLimit(id string, f func(int, int, int, int, string) ([]interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		//TODO Check id and errors (4 real)
		id, err := strconv.Atoi(c.Param(id))
		who, whoKind := idKind(c)
		offset, limit := offsetLimit(c)
		res, err := f(id, limit, offset, who, whoKind)
		h := checkHAL(c)
		for i, el := range res {
			//TODO handle err
			res[i], _ = ToRepresentation(el, c, h)
		}
		if !h {
			result := models.List{
				Self:     c.Request.RequestURI,
				Data:     res,
				Next:     next(c.Request.RequestURI, offset, limit, res),
				Previous: prev(c.Request.RequestURI, offset, limit),
			}
			handleErr(err, result, c)
		} else {
			halList(c, &res, err, next(c.Request.RequestURI, offset, limit, res), prev(c.Request.RequestURI, offset, limit))
		}
	}
}

func prev(uri string, offset int, limit int) string {
	if offset == 0 {
		return ""
	} else if n := strings.Index(uri, "?"); n >= 0 {
		uri = uri[:n]
	}
	if prev := offset - limit; prev < 0 {
		offset = 0
	} else {
		offset = prev
	}
	return strings.Join([]string{uri, fmt.Sprintf("?offset=%d&limit=%d", offset, limit)}, "")
}

func next(uri string, offset int, limit int, input []interface{}) string {
	if l := len(input); l < limit {
		return ""
	}
	if n := strings.Index(uri, "?"); n >= 0 {
		uri = uri[:n]
	}
	return strings.Join([]string{uri, fmt.Sprintf("?offset=%d&limit=%d", offset+limit, limit)}, "")
}

func unauthorized(c *gin.Context) {
	c.Header("WWW-Authenticate", viper.GetString("Realm"))
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
}

func negotiate(data interface{}, c *gin.Context, status int) {
	if !checkHAL(c) {
		c.Negotiate(status, gin.Negotiate{
			Offered: []string{gin.MIMEJSON, gin.MIMEXML},
			Data:    data,
		})
	} else {
		c.JSON(status, data)
	}
}

func handleErr(err error, res interface{}, c *gin.Context) {
	if res != nil {
		switch err {
		case nil:
			negotiate(res, c, http.StatusOK)
		case repository.ErrNoResult:
			negotiate(gin.H{"error": err.Error()}, c, http.StatusNotFound)
		default:
			negotiate(gin.H{"error": err.Error()}, c, http.StatusBadRequest)
		}
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}

func checkHAL(c *gin.Context) bool {
	return strings.HasPrefix(c.GetHeader("Accept"), HAL)

}

func ToRepresentation(res interface{}, c *gin.Context, halF bool) (interface{}, error) {
	if r, ok := res.(models.Repr); ok {
		return r.GetRepresentation(halF)
	} else {
		return res, nil
	}
}

// @Summary Authenticate
// @Param Account body models.Login true "data"
// @Tags Accounts
// @Success 200 {object} models.Token
// @Router /login [post]
func LoginHandler(authMiddleware *jwt.GinJWTMiddleware) func(*gin.Context) {
	return authMiddleware.LoginHandler
}
