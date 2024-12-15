<div align="center">
	<img src="./korm.png" width="auto" style="margin:0 auto 0 auto;"/>
</div>
<br>
<div align="center">
	<img src="https://img.shields.io/github/go-mod/go-version/kamalshkeir/korm" width="auto" height="20px">
	<img src="https://img.shields.io/github/languages/code-size/kamalshkeir/korm" width="auto" height="20px">
	<img src="https://img.shields.io/badge/License-BSD%20v3-blue.svg" width="auto" height="20px">
	<img src="https://img.shields.io/github/v/tag/kamalshkeir/korm" width="auto" height="20px">
	<img src="https://img.shields.io/github/stars/kamalshkeir/korm?style=social" width="auto" height="20px">
	<img src="https://img.shields.io/github/forks/kamalshkeir/korm?style=social" width="auto" height="20px">
</div>


<div align="center">
	<a href="https://www.youtube.com/watch?v=KMnnwly3Mpc" style="display:flex;justify-content:center;align-items:center;gap:10px;">
	<img src="https://user-images.githubusercontent.com/54605903/217871012-9c5dc1da-25bd-47d5-ac9e-c3acee7178d5.svg" width="auto" height="50px">
	<span><strong> Simple Example video with bus </strong></span>
	</a>
</div>



<svg height="50px" width="50px">
	<g>
		<path style="fill:#F61C0D;" d="M365.257,67.393H95.744C42.866,67.393,0,110.259,0,163.137v134.728
			c0,52.878,42.866,95.744,95.744,95.744h269.513c52.878,0,95.744-42.866,95.744-95.744V163.137
			C461.001,110.259,418.135,67.393,365.257,67.393z M300.506,237.056l-126.06,60.123c-3.359,1.602-7.239-0.847-7.239-4.568V168.607
			c0-3.774,3.982-6.22,7.348-4.514l126.06,63.881C304.363,229.873,304.298,235.248,300.506,237.056z"/>
	</g>
</svg>

<br>
<div align="center">
	<a href="https://kamalshkeir.dev" target="_blank">
		<img src="https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white" width="auto" height="32px">
	</a>
	<a href="https://www.linkedin.com/in/kamal-shkeir/">
		<img src="https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white" width="auto" height="30px">
	</a>
	<a href="https://www.buymeacoffee.com/kamalshkeir" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" width="auto" height="32px" ></a>

	
</div>

---
### Introducing Korm - the elegant, lightning-fast ORM/Framework for all your needs, see [benchmarks](#benchmarks-vs-gorm). Inspired by the highly popular Django Framework, Korm offers similar functionality with the added bonus of performance

### It is also composable, allowing for integration with a network websocket PubSub using [WithBus](#example-with-bus-between-2-korm) when you want to synchronise your data between multiple Korm or [WithDashboard](#example-with-dashboard-you-dont-need-kormwithbus-with-it-because-withdashboard-already-call-it-and-return-the-server-bus-for-you) to have a complete setup of server bus and Admin Dashboard.

#### Why settle for less when you can have the best ?
- Django become very hard to work with when you need concurrency and async, you will need django channels and a server like daphne or uvicorn, Go have the perfect implementation
- Django can handle at most 300 request per second, Go handle 44,000 requests per second (benchmarks done on my machine)
- The API is also more user-friendly and less verbose than Django's
- Deploying an executable binary file using Korm , with automatic TLS Let's encrypt, a built-in Admin Dashboard, Interactive Shell, Eventbus to communicate between multiple Korm applications is pretty neat
- Additionally, its caching system uses goroutines and channels to efficiently to clean the cache when rows or tables are created, updated, deleted, or dropped

### It Has :
- <strong>New:</strong>  When using korm.WithDashboard, now you have access to all logs in realtime (websockets) from admin dashboard when you log using lg pkg. By default only 10 last logs are keeped in memory, you can increase it using lg.SaveLogs(50) for keeping last 50 logs

- <strong>New:</strong>  Automatic check your structs (schema) against database tables, prompt you with changes, and so it can add or remove columns by adding or removing fields to the struct, it is Disabled by default, use `korm.EnableCheck()` to enable it

- <strong>New:</strong>  [Handle Nested or Embeded structs](#example-nested-or-embeded-structs) and slice of structs through joins, like sqlx, but sqlx doesn't handle slice of structs

- <strong>New:</strong> korm.QueryNamed, QueryNamedS, korm.ExecNamed, korm.ExecContextNamed and WhereNamed(query string, args map[string]any) like :Where("email = :email",map[string]any{"email":"abc@mail.com"}) 

- <strong>New:</strong>  korm.LogsQueries() that log statements and time tooked by sql queries 

- <a href="#swagger-documentation">Auto Docs with Model API and video tutoriel
</a><a href="https://www.youtube.com/watch?v=r7rbMrTkVek">
	<img src="https://user-images.githubusercontent.com/54605903/217871012-9c5dc1da-25bd-47d5-ac9e-c3acee7178d5.svg" width="auto" height="50px">
</a>

- <a href="#swagger-documentation">Swagger Documentation and tutorial
</a><a href="https://www.youtube.com/watch?v=RupARTkPzf4">
	<img src="https://user-images.githubusercontent.com/54605903/217871012-9c5dc1da-25bd-47d5-ac9e-c3acee7178d5.svg" width="auto" height="50px">
</a>

- [PPROF](#pprof) Go profiling tool and [Metrics Prometheus](#metrics-prometheus)

- [Logs Middleware](#logs-middleware)

- [Admin dashboard](#example-with-dashboard-you-dont-need-kormwithbus-with-it-because-withdashboard-already-call-it-and-return-the-server-bus-for-you) with ready offline and installable PWA (using /static/sw.js and /static/manifest.webmanifest). All statics mentionned in `sw.js` will be cached and served by the service worker, you can inspect the Network Tab in the browser to check it

- Shared Network Bus allowing you to send and recv data in realtime using pubsub websockets between your ORMs, so you can decide how you data will be distributed between different databases, see [Example](#example-with-bus-between-2-korm) 

- [Built-in Authentication](#auth-middleware-example) using `korm.Auth` , `korm.Admin` or `korm.BasicAuth` middlewares, whenever Auth and Admin middlewares are used, you get access to the `.User` model and variable `.IsAuthenticated` from any template html like this example [admin_nav.html](#example-admin-and-auth-user-model-and-isauthenticated) 

- [Interactive Shell](#interactive-shell), to CRUD in your databases from command line, use `korm.WithShell()`

- [AutoMigrate](#automigrate) directly from struct

- Compatible with official database/sql,  so you can do your queries yourself using sql.DB  `korm.GetConnection()``, and overall a painless integration of your existing codebases using database/sql

- [Router/Mux](https://github.com/kamalshkeir/ksmux) accessible from the serverBus after calling `korm.WithBus(...opts)` or `korm.WithDashboard(addr, ...opts)`

- [Hooks](#hooks) : OnInsert OnSet OnDelete and OnDrop

- [many to many](#manytomany-relationships-example) relationships

- [GENERATED ALWAYS AS](#example-generated-tag) tag added (all dialects)

- [Concatination and Length](#example-concat-and-len-from-korm_testgo) support for `Where` and for tags: `check` and `generated` (all dialects)

- Support for foreign keys, indexes , checks,... [See all](#automigrate)

- [Kenv](#example-not-required-load-config-from-env-directly-to-struct-using-kenv) load env vars to struct

- [Python Bus Client](#python-bus-client-example) `pip install ksbus`

####  All drivers concurrent safe read and write
#### Supported databases:
- Sqlite
- Mysql
- Maria
- Postgres
- Cockroach


---
# Installation

```sh
go get -u github.com/kamalshkeir/korm@v1.94.8
```

# Drivers moved outside this package to not get them all in your go.mod file
```sh
go get -u github.com/kamalshkeir/sqlitedriver@latest
go get -u github.com/kamalshkeir/pgdriver@latest
go get -u github.com/kamalshkeir/mysqldriver@latest
```

### Global Vars
```go
// Debug when true show extra useful logs for queries executed for migrations and queries statements
Debug = false
// FlushCacheEvery execute korm.FlushCache() every 10 min by default, you should not worry about it, but useful that you can change it
FlushCacheEvery = 10 * time.Minute
// SetCacheMaxMemory set max size of each cache cacheAllS AllM ...
korm.SetCacheMaxMemory(megaByte int) // default maximum of 50 Mb , cannot be lower
// Connection pool
MaxOpenConns = 20
MaxIdleConns = 20
MaxLifetime = 30 * time.Minute
MaxIdleTime = 15 * time.Minute
```

### Connect to a database
```go
// sqlite
// go get github.com/kamalshkeir/sqlitedriver
err := korm.New(korm.SQLITE, "dbName", sqlitedriver.Use()) // Connect
// postgres, cockroach
// go get github.com/kamalshkeir/pgdriver
err := korm.New(korm.POSTGRES,"dbName", pgdriver.Use(), "user:password@localhost:5432") // Connect
// mysql, maria
// go get github.com/kamalshkeir/mysqldriver
err := korm.New(korm.MYSQL,"dbName", mysqldriver.Use(), "user:password@localhost:3306") // Connect

korm.Shutdown(databasesName ...string) error
```

### Hello world example

```go
package main

import (
	"fmt"
	"time"

	"github.com/kamalshkeir/lg"
	"github.com/kamalshkeir/korm"
	"github.com/kamalshkeir/sqlitedriver"
)

type Class struct {
	Id       uint `korm:"pk"`
	Name     string
	Students []Student
}

type Student struct {
	Id      uint `korm:"pk"`
	Name    string
	Class   uint `korm:"fk:classes.id:cascade:cascade"`
	Classes Class
}

func main() {
	err := korm.New(korm.SQLITE, "db", sqlitedriver.Use())
	if lg.CheckError(err) {
		return
	}
	defer korm.Shutdown()

	server := korm.WithDashboard(":9313")
	korm.WithShell()

	err = korm.AutoMigrate[Class]("classes")
	lg.CheckError(err)

	err = korm.AutoMigrate[Student]("students")
	lg.CheckError(err)

	// go run main.go shell to createsuperuser
	// connect to admin and create some data to query

	// nested structs with joins, scan the result to the channel directly after each row
	// so instead of receiving a slice, you will receive data on the channel[0] of the passed slice
	studentsChan := []chan Student{make(chan Student)}
	go func() {
		for s := range studentsChan[0] {
			fmt.Println("chan students:", s)
		}
	}()
	err = korm.To(&studentsChan).Query("select students.*,classes.id as 'classes.id',classes.name as 'classes.name'  from students join classes where classes.id = students.class")
	lg.CheckError(err)
	fmt.Println()

	// Named with nested (second argument of 'To') filled automatically from join, support nested slices and structs
	classes := []Class{}
	query := "select classes.*, students.id as 'students.id',students.name as 'students.name' from classes join students on students.class = classes.id order by :order_here"
	err = korm.To(&classes, true).Named(query,map[string]any{
		"order_here": "classes.id",
	})
	lg.CheckError(err)
	for _, s := range classes {
		fmt.Println("class:", s)
	}
	fmt.Println()

	// // not nested, only remove second arg true from 'To' method
	students := []Student{}
	err = korm.To(&students, true).Query("select students.*,classes.id as 'classes.id',classes.name as 'classes.name'  from students join classes where classes.id = students.class")
	lg.CheckError(err)
	for _, s := range students {
		fmt.Println("student:", s)
	}
	fmt.Println()

	maps := []map[string]any{}
	err = korm.To(&maps).Query("select * from students")
	lg.CheckError(err)
	fmt.Println("maps =", maps)
	fmt.Println()

	names := []*string{}
	err = korm.To(&names).Query("select name from students")
	lg.CheckError(err)
	fmt.Println("names =", names)
	fmt.Println()

	ids := []int{}
	err = korm.To(&ids).Query("select id from students")
	lg.CheckError(err)
	fmt.Println("ids =", ids)
	fmt.Println()

	bools := []bool{}
	err = korm.To(&bools).Query("select is_admin from users")
	lg.CheckError(err)
	fmt.Println("bools =", bools)
	fmt.Println()

	times := []time.Time{}
	err = korm.To(&times).Query("select created_at from users")
	lg.CheckError(err)
	fmt.Println("times =", times)

	server.Run()
}

// OUTPUT
// chan students: {1 student-1 1 {1 Math []}}
// chan students: {2 student-2 2 {2 French []}}
// chan students: {3 student-3 1 {1 Math []}}
// chan students: {4 student-4 2 {2 French []}}

// class: {1 Math [{1 student-1 0 {0  []}} {3 student-3 0 {0  []}}]}
// class: {2 French [{2 student-2 0 {0  []}} {4 student-4 0 {0  []}}]}

// student: &{1 student-1 1 {1 Math []}}
// student: &{2 student-2 2 {2 French []}}
// student: &{3 student-3 1 {1 Math []}}
// student: &{4 student-4 2 {2 French []}}

// maps = [map[class:1 id:1 name:student-1] map[class:2 id:2 name:student-2] map[class:1 id:3 name:student-3] map[class:2 id:4 name:student-4]]

// names = [student-1 student-2 student-3 student-4]

// ids = [1 2 3 4]

// bools = [true]

// times = [2023-04-30 19:19:32 +0200 CEST]
```


### AutoMigrate 

[Available Tags](#available-tags-by-struct-field-type) (SQL)

SQL:
```go
korm.AutoMigrate[T comparable](tableName string, dbName ...string) error 

err := korm.AutoMigrate[User]("users")
err := korm.AutoMigrate[Bookmark ]("bookmarks")

type User struct {
	Id        int       `korm:"pk"` // AUTO Increment ID primary key
	Uuid      string    `korm:"size:40"` // VARCHAR(50)
	Email     string    `korm:"size:50;iunique"` // insensitive unique
	Password  string    `korm:"size:150"` // VARCHAR(150)
	IsAdmin   bool      `korm:"default:false"` // DEFAULT 0
	Image     string    `korm:"size:100;default:''"`
	CreatedAt time.Time `korm:"now"` // auto now
    Ignored   string    `korm:"-"`
}

type Bookmark struct {
	Id      uint   `korm:"pk"`
	UserId  int    `korm:"fk:users.id:cascade:setnull"` // options cascade,donothing/noaction, setnull/null, setdefault/default
	IsDone	bool   
	ToCheck string `korm:"size:50; notnull; check: len(to_check) > 2 AND len(to_check) < 10; check: is_done=true"`  // column type will be VARCHAR(50)
	Content string `korm:"text"` // column type will be TEXT not VARCHAR
	UpdatedAt time.Time `korm:"update"` // will update when model updated, handled by triggers for sqlite, cockroach and postgres, and on migration for mysql
	CreatedAt time.Time `korm:"now"` // now is default to current timestamp and of type TEXT for sqlite
}

all, _ := korm.Model[User]()
                   .Where("id = ?",id) 
                   .Select("item1","item2")
                   .OrderBy("created")
				   .Limit(8)
				   .Page(2)
                   .All()
```


### API
```go
korm.New(dbType Dialect, dbName string, dbDriver driver.Driver, dbDSN ...string) error
korm.LogQueries()
korm.GetConnection(dbName ...string) *sql.DB
korm.To[T any](dest *[]T, nestedSlice ...bool) *Selector[T] // scan query to any type slice, even channels and slices with nested structs and joins
(sl *Selector[T]) Ctx(ct context.Context) *Selector[T]
(sl *Selector[T]) Query(statement string, args ...any) error
(sl *Selector[T]) Named(statement string, args map[string]any, unsafe ...bool) error
korm.WithBus(...opts) *ksbus.Server // Usage: WithBus(...opts) or share an existing one
korm.WithDashboard(address, ...opts) *ksbus.Server
korm.WithShell()
korm.WithDocs(generateJsonDocs bool, outJsonDocs string, handlerMiddlewares ...func(handler kmux.Handler) kmux.Handler) *ksbus.Server
korm.WithEmbededDocs(embeded embed.FS, embededDirPath string, handlerMiddlewares ...func(handler kmux.Handler) kmux.Handler) *ksbus.Server
korm.WithMetrics(httpHandler http.Handler) *ksbus.Server
korm.WithPprof(path ...string) *ksbus.Server
korm.Transaction(dbName ...string) (*sql.Tx, error)
korm.Exec(dbName, query string, args ...any) error
korm.ExecContext(ctx context.Context, dbName, query string, args ...any) error
korm.ExecNamed(query string, args map[string]any, dbName ...string) error
korm.ExecContextNamed(ctx context.Context, query string, args map[string]any, dbName ...string) error
korm.BeforeServersData(fn func(data any, conn *ws.Conn))
korm.BeforeDataWS(fn func(data map[string]any, conn *ws.Conn, originalRequest *http.Request) bool)
korm.GetAllTables(dbName ...string) []string
korm.GetAllColumnsTypes(table string, dbName ...string) map[string]string
korm.GetMemoryTable(tbName string, dbName ...string) (TableEntity, error)
korm.GetMemoryTables(dbName ...string) ([]TableEntity, error)
korm.GetMemoryDatabases() []DatabaseEntity
korm.GetMemoryDatabase(dbName string) (*DatabaseEntity, error)
korm.Shutdown(databasesName ...string) error
korm.FlushCache()
korm.DisableCache() 
korm.ManyToMany(table1, table2 string, dbName ...string) error // add table relation m2m 
```
#### Builder `Struct`:
```go
korm.Exec(dbName, query string, args ...any) error
korm.Transaction(dbName ...string) (*sql.Tx, error)
// Model is a starter for Buider
func Model[T comparable](tableName ...string) *BuilderS[T]
// Database allow to choose database to execute query on
func (b *BuilderS[T]) Database(dbName string) *BuilderS[T]
// Insert insert a row into a table and return inserted PK
func (b *BuilderS[T]) Insert(model *T) (int, error)
// InsertR add row to a table using input struct, and return the inserted row
func (b *BuilderS[T]) InsertR(model *T) (T, error)
// BulkInsert insert many row at the same time in one query
func (b *BuilderS[T]) BulkInsert(models ...*T) ([]int, error)
// AddRelated used for many to many, and after korm.ManyToMany, to add a class to a student or a student to a class, class or student should exist in the database before adding them
func (b *BuilderS[T]) AddRelated(relatedTable string, whereRelatedTable string, whereRelatedArgs ...any) (int, error)
// DeleteRelated delete a relations many to many
func (b *BuilderS[T]) DeleteRelated(relatedTable string, whereRelatedTable string, whereRelatedArgs ...any) (int, error)
// GetRelated used for many to many to get related classes to a student or related students to a class
func (b *BuilderS[T]) GetRelated(relatedTable string, dest any) error
// JoinRelated same as get, but it join data
func (b *BuilderS[T]) JoinRelated(relatedTable string, dest any) error
// Set used to update, Set("email,is_admin","example@mail.com",true) or Set("email = ? AND is_admin = ?","example@mail.com",true)
func (b *BuilderS[T]) Set(query string, args ...any) (int, error)
// Delete data from database, can be multiple, depending on the where, return affected rows(Not every database or database driver may support affected rows)
func (b *BuilderS[T]) Delete() (int, error)
// Drop drop table from db
func (b *BuilderS[T]) Drop() (int, error)
// Select usage: Select("email","password")
func (b *BuilderS[T]) Select(columns ...string) *BuilderS[T]
// Where can be like : Where("id > ?",1) or Where("id",1) = Where("id = ?",1)
func (b *BuilderS[T]) Where(query string, args ...any) *BuilderS[T]
// Limit set limit
func (b *BuilderS[T]) Limit(limit int) *BuilderS[T]
// Context allow to query or execute using ctx
func (b *BuilderS[T]) Context(ctx context.Context) *BuilderS[T]
// Page return paginated elements using Limit for specific page
func (b *BuilderS[T]) Page(pageNumber int) *BuilderS[T]
// OrderBy can be used like: OrderBy("-id","-email") OrderBy("id","-email") OrderBy("+id","email")
func (b *BuilderS[T]) OrderBy(fields ...string) *BuilderS[T]
// Debug print prepared statement and values for this operation
func (b *BuilderS[T]) Debug() *BuilderS[T]
// All get all data
func (b *BuilderS[T]) All() ([]T, error)
// One get single row
func (b *BuilderS[T]) One() (T, error)

Examples:
korm.Model[models.User]().Select("email","uuid").OrderBy("-id").Limit(PAGINATION_PER).Page(1).All()

// INSERT
uuid,_ := korm.GenerateUUID()
hashedPass,_ := argon.Hash(password)
korm.Model[models.User]().Insert(&models.User{
	Uuid: uuid,
	Email: "test@example.com",
	Password: hashedPass,
	IsAdmin: false,
	Image: "",
	CreatedAt: time.Now(),
})

//if using more than one db
korm.Database[models.User]("dbNameHere").Where("id = ? AND email = ?",1,"test@example.com").All() 

// where
korm.Model[models.User]().Where("id = ? AND email = ?",1,"test@example.com").One() 

// delete
korm.Model[models.User]().Where("id = ? AND email = ?",1,"test@example.com").Delete()

// drop table
korm.Model[models.User]().Drop()

// update
korm.Model[models.User]().Where("id = ?",1).Set("email = ?","new@example.com")
```
#### Builder `map[string]any`:
```go
// BuilderM is query builder map string any
type BuilderM struct
// Table is a starter for BuiderM
func Table(tableName string) *BuilderM
// Database allow to choose database to execute query on
func (b *BuilderM) Database(dbName string) *BuilderM
// Select select table columns to return
func (b *BuilderM) Select(columns ...string) *BuilderM
// Where can be like: Where("id > ?",1) or Where("id",1) = Where("id = ?",1)
func (b *BuilderM) Where(query string, args ...any) *BuilderM
// Limit set limit
func (b *BuilderM) Limit(limit int) *BuilderM
// Page return paginated elements using Limit for specific page
func (b *BuilderM) Page(pageNumber int) *BuilderM
// OrderBy can be used like: OrderBy("-id","-email") OrderBy("id","-email") OrderBy("+id","email")
func (b *BuilderM) OrderBy(fields ...string) *BuilderM
// Context allow to query or execute using ctx
func (b *BuilderM) Context(ctx context.Context) *BuilderM
// Debug print prepared statement and values for this operation
func (b *BuilderM) Debug() *BuilderM
// All get all data
func (b *BuilderM) All() ([]map[string]any, error)
// One get single row
func (b *BuilderM) One() (map[string]any, error)
// Insert add row to a table using input map, and return PK of the inserted row
func (b *BuilderM) Insert(rowData map[string]any) (int, error)
// InsertR add row to a table using input map, and return the inserted row
func (b *BuilderM) InsertR(rowData map[string]any) (map[string]any, error)
// BulkInsert insert many row at the same time in one query
func (b *BuilderM) BulkInsert(rowsData ...map[string]any) ([]int, error)
// Set used to update, Set("email,is_admin","example@mail.com",true) or Set("email = ? AND is_admin = ?","example@mail.com",true)
func (b *BuilderM) Set(query string, args ...any) (int, error)
// Delete data from database, can be multiple, depending on the where, return affected rows(Not every database or database driver may support affected rows)
func (b *BuilderM) Delete() (int, error)
// Drop drop table from db
func (b *BuilderM) Drop() (int, error)
// AddRelated used for many to many, and after korm.ManyToMany, to add a class to a student or a student to a class, class or student should exist in the database before adding them
func (b *BuilderM) AddRelated(relatedTable string, whereRelatedTable string, whereRelatedArgs ...any) (int, error)
// GetRelated used for many to many to get related classes to a student or related students to a class
func (b *BuilderM) GetRelated(relatedTable string, dest *[]map[string]any) error
// JoinRelated same as get, but it join data
func (b *BuilderM) JoinRelated(relatedTable string, dest *[]map[string]any) error
// DeleteRelated delete a relations many to many
func (b *BuilderM) DeleteRelated(relatedTable string, whereRelatedTable string, whereRelatedArgs ...any) (int, error)


Examples:

sliceMapStringAny,err := korm.Table("users")
							.Select("email","uuid")
							.OrderBy("-id")
							.Limit(PAGINATION_PER)
							.Page(1)
							.All()

// INSERT
uuid,_ := korm.GenerateUUID()
hashedPass,_ := argon.Hash("password") // github.com/kamalshkeir/argon

korm.Table("users").Insert(map[string]any{
	"uuid":uuid,
	"email":"test@example.com",
	 ...
})

//if using more than one db
korm.Database("dbNameHere").Table("tableName").Where("id = ? AND email = ?",1,"test@example.com").All() 

// where
Where("id = ? AND email = ?",1,"test@example.com") // this work
Where("id,email",1,"test@example.com") // and this work

korm.Table("tableName").Where("id = ? AND email = ?",1,"test@example.com").One() 

// delete
korm.Table("tableName").Where("id = ? AND email = ?",1,"test@example.com").Delete() 

// drop table
korm.Table("tableName").Drop()

// update
korm.Table("tableName").Where("id = ?",1).Set("email = ?","new@example.com") 
korm.Table("tableName").Where("id",1).Set("email","new@example.com") 
```

### Dashboard defaults you can set
```go
korm.PaginationPer      = 10
korm.DocsUrl           = "docs"
korm.EmbededDashboard   = false
korm.MediaDir           = "media"
korm.AssetsDir          = "assets"
korm.StaticDir          = path.Join(AssetsDir, "/", "static")
korm.TemplatesDir       = path.Join(AssetsDir, "/", "templates")
korm.RepoUser           = "kamalshkeir"
korm.RepoName           = "korm-dash"
korm.adminPathNameGroup = "/admin" // korm.SetAdminPath("/another")
// so you can create a custom dashboard, upload it to your repos and change like like above korm.RepoUser and korm.RepoName
```

### Example With Dashboard (you don't need korm.WithBus with it, because WithDashboard already call it and return the server bus for you)

```go
package main

import (
	"github.com/kamalshkeir/lg"
	"github.com/kamalshkeir/ksmux"
	"github.com/kamalshkeir/korm"
	"github.com/kamalshkeir/sqlitedriver"
)

func main() {
	err := korm.New(korm.SQLITE, "db", sqlitedriver.Use())
	lg.CheckError(err)



	serverBus := korm.WithDashboard("localhost:9313")
	korm.WithShell()
	// you can overwrite Admin and Auth middleware used for dashboard (dash_middlewares.go) 
	//korm.Auth = func(handler ksmux.Handler) ksmux.Handler {}
	//korm.Admin = func(handler ksmux.Handler) ksmux.Handler {}

	// and also all handlers (dash_views.go)
	//korm.LoginView = func(c *ksmux.Context) {
	//	c.Html("admin/new_admin_login.html", nil)
	//}

	// add extra static directory if you want
	//serverBus.App.LocalStatics("assets/mystatic","myassets") // will be available at /myassets/*
	//serverBus.App.LocalTemplates("assets/templates") // will make them available to use with c.Html

	// serve HTML 
	// serverBus.App.Get("/",func(c *ksmux.Context) {
	// 	c.Html("index.html", map[string]any{
	// 		"data": data,
	// 	})
	// })
	serverBus.Run()
	// OR run https if you have certificates
	serverBus.RunTLS(cert string, certKey string)

	// OR generate certificates let's encrypt for a domain name, check https://github.com/kamalshkeir/ksbus for more infos
	serverBus.RunAutoTLS(subDomains ...string)
}
```
Then create admin user to connect to the dashboard
```sh
go run main.go shell

createsuperuser
```

Then you can visit `/admin`


### Auth middleware example
```go
func main() {
	err := korm.New(korm.SQLITE, "db", sqlitedriver.Use())
	if lg.CheckError(err) {
		return
	}
	defer korm.Shutdown()
	
	srv := korm.WithDashboard("localhost:9313")
	korm.WithShell()
	lg.Printfs("mgrunning on http://localhost:9313\n")
	app := srv.App

	app.Get("/", korm.Auth(func(c *ksmux.Context) { // work with korm.Admin also
		// c.IsAuthenticated also return bool
		if v, ok := c.User(); ok {
			c.Json(map[string]any{
				"msg": "Authenticated",
				"v":   v.(korm.User).Email,
			})
		} else {
			c.Json(map[string]any{
				"error": "not auth",
			})
		}
	}))

	srv.Run()
}
```

### Example Admin Auth User and IsAuthenticated
```html
{{define "admin_nav"}}
<header id="admin-header">
  <nav>
    <a href="/">
      <h1>KORM</h1>
    </a> 
    
    <ul>
        <li>
          <a {{if eq .Request.URL.Path "/" }}class="active"{{end}} href="/">Home</a>
        </li>

        <li>
          <a {{if contains .Request.URL.Path .admin_path }}class="active"{{end}} href="{{.admin_path}}">Admin</a>
        </li>

        {{if .IsAuthenticated}}
            <li>
              <a href="{{.admin_path}}/logout">Logout</a>
            </li>
            
            {{if .User.Email}}
              <li>
                <span>Hello {{.User.Email}}</span>
              </li>
            {{end}}
        {{end}}
    </ul>
  </nav>
</header>
{{end}}


```

### Admin middlewares

```go
// dash_middlewares.go
package korm

import (
	"context"
	"net/http"

	"github.com/kamalshkeir/aes"
	"github.com/kamalshkeir/ksmux"
)

var Auth = func(handler ksmux.Handler) ksmux.Handler {
	return func(c *ksmux.Context) {
		session, err := c.GetCookie("session")
		if err != nil || session == "" {
			// NOT AUTHENTICATED
			c.DeleteCookie("session")
			handler(c)
			return
		}
		session, err = aes.Decrypt(session)
		if err != nil {
			handler(c)
			return
		}
		// Check session
		user, err := Model[User]().Where("uuid = ?", session).One()
		if err != nil {
			// session fail
			handler(c)
			return
		}

		// AUTHENTICATED AND FOUND IN DB
		c.SetKey("korm-user", user)
		handler(c)
	}
}

var Admin = func(handler ksmux.Handler) ksmux.Handler {
	return func(c *ksmux.Context) {
		session, err := c.GetCookie("session")
		if err != nil || session == "" {
			// NOT AUTHENTICATED
			c.DeleteCookie("session")
			c.Status(http.StatusTemporaryRedirect).Redirect(adminPathNameGroup + "/login")
			return
		}
		session, err = aes.Decrypt(session)
		if err != nil {
			c.Status(http.StatusTemporaryRedirect).Redirect(adminPathNameGroup + "/login")
			return
		}
		user, err := Model[User]().Where("uuid = ?", session).One()

		if err != nil {
			// AUTHENTICATED BUT NOT FOUND IN DB
			c.Status(http.StatusTemporaryRedirect).Redirect(adminPathNameGroup + "/login")
			return
		}

		// Not admin
		if !user.IsAdmin {
			c.Status(403).Text("Middleware : Not allowed to access this page")
			return
		}
		c.SetKey("korm-user", user)
		handler(c)
	}
}

var BasicAuth = func(handler ksmux.Handler) ksmux.Handler {
	return ksmux.BasicAuth(handler, BASIC_AUTH_USER, BASIC_AUTH_PASS)
}

```

### Example With Bus between 2 KORM
KORM 1:

```go
package main

import (
	"net/http"

	"github.com/kamalshkeir/lg"
	"github.com/kamalshkeir/ksmux"
	"github.com/kamalshkeir/ksmux/ws"
	"github.com/kamalshkeir/korm"
	"github.com/kamalshkeir/ksbus"
	"github.com/kamalshkeir/sqlitedriver"
)

func main() {
	err := korm.New(korm.SQLITE,"db1", sqlitedriver.Use())
	if lg.CheckError(err) {return}

	korm.WithShell()
	serverBus := korm.WithBus(ksbus.ServerOpts{
		ID              string
		Address         string
		Path            string
		OnWsClose       func(connID string)
		OnDataWS        func(data map[string]any, conn *ws.Conn, originalRequest *http.Request) error
		OnServerData    func(data any, conn *ws.Conn)
		OnId            func(data map[string]any)
		OnUpgradeWs     func(r *http.Request) bool
		WithOtherRouter *ksmux.Router
		WithOtherBus    *Bus
	})
	// handler authentication	
	korm.BeforeDataWS(func(data map[string]any, conn *ws.Conn, originalRequest *http.Request) bool {
		lg.Info("handle authentication here")
		return true
	})
	// handler data from other KORM
	korm.BeforeServersData(func(data any, conn *ws.Conn) {
		lg.Info("recv orm:", "data", data)
	})

	// built in router to the bus, check it at https://github.com/kamalshkeir/ksbus
	serverBus.App.Get("/",func(c *ksmux.Context) {
		serverBus.SendToServer("localhost:9314",map[string]any{
			"msg":"hello from server 1",
		})
		c.Text("ok")
	})

	
	serverBus.Run("localhost:9313")
	// OR run https if you have certificates
	serverBus.RunTLS(addr string, cert string, certKey string)
	// OR generate certificates let's encrypt for a domain name, check https://github.com/kamalshkeir/ksbus for more details
	serverBus.RunAutoTLS(domainName string, subDomains ...string)
}
```
KORM 2:
```go
package main

import (
	"net/http"

	"github.com/kamalshkeir/lg"
	"github.com/kamalshkeir/ksmux"
	"github.com/kamalshkeir/ksmux/ws"
	"github.com/kamalshkeir/korm"
	"github.com/kamalshkeir/sqlitedriver"
)

func main() {
	err := korm.New(korm.SQLITE,"db2",sqlitedriver.Use())
	if lg.CheckError(err) {return}

	korm.WithShell() // if dashboard used, this line should be after it
	serverBus := korm.WithBus(ksbus.ServerOpts{
		ID              string
		Address         string
		Path            string
		OnWsClose       func(connID string)
		OnDataWS        func(data map[string]any, conn *ws.Conn, originalRequest *http.Request) error
		OnServerData    func(data any, conn *ws.Conn)
		OnId            func(data map[string]any)
		OnUpgradeWs     func(r *http.Request) bool
		WithOtherRouter *ksmux.Router
		WithOtherBus    *Bus
	})

	korm.BeforeServersData(func(data any, conn *ws.Conn) {
        lg.Info("recv", "data", data)
	})

	// built in router to the bus, check it at https://github.com/kamalshkeir/ksbus
	serverBus.App.GET("/",func(c *ksmux.Context) {
		serverBus.SendToServer("localhost:9314",map[string]any{
			"msg":"hello from server 2",
		})
		c.Status(200).Text("ok")
	})


    // Run Server Bus
	serverBus.Run("localhost:9314")

	// OR run https if you have certificates
	serverBus.RunTLS(addr string, cert string, certKey string)

	// OR generate certificates let's encrypt for a domain name, check https://github.com/kamalshkeir/ksbus for more infos
	serverBus.RunAutoTLS(domainName string, subDomains ...string)
}
```

### Example generated tag
```go
// generated example using concatination and length
type TestUser struct {
	Id        *uint   `korm:"pk"`
	Uuid      string  `korm:"size:40;iunique"`
	Email     *string `korm:"size:100;iunique"`
	Gen       string  `korm:"size:250;generated: concat(uuid,'working',len(password))"`
	Password  string
	IsAdmin   *bool
	CreatedAt time.Time `korm:"now"`
	UpdatedAt time.Time `korm:"update"`
}

func TestGeneratedAs(t *testing.T) {
	u, err := Model[TestUser]().Limit(3).All()
	if err != nil {
		t.Error(err)
	}
	if len(u) != 3 {
		t.Error("len not 20")
	}
	if u[0].Gen != u[0].Uuid+"working"+fmt.Sprintf("%d", len(u[0].Password)) {
		t.Error("generated not working:", u[0].Gen)
	}
}
```


### Example concat and len from korm_test.go
```go
// Where example
func TestConcatANDLen(t *testing.T) {
	groupes, err := Model[Group]().Where("name = concat(?,'min') AND len(name) = ?", "ad", 5).Debug().All()
	// translated to select * from groups WHERE name = 'ad' || 'min'  AND  length(name) = 5 (sqlite)
	// translated to select * from groups WHERE name = concat('ad','min')  AND  char_length(name) = 5 (postgres, mysql)
	if err != nil {
		t.Error(err)
	}
	if len(groupes) != 1 || groupes[0].Name != "admin" {
		t.Error("len(groupes) != 1 , got: ", groupes)
	}
}
```




## Router/Mux 
Learn more about [Ksmux](https://github.com/kamalshkeir/ksmux)
```go
func main() {
	err := korm.New(korm.SQLITE, "db", sqlitedriver.Use())
	if err != nil {
		log.Fatal(err)
	}

	serverBus := korm.WithDashboard("localhost:9313")
	korm.WithShell()
	mux := serverBus.App
	// add global middlewares
	mux.Use((midws ...func(http.Handler) http.Handler))
	...
}

```

### Pprof
```go

serverBus := korm.WithDashboard("localhost:9313")
// or srv := korm.WithBus()
serverBus.WithPprof(path ...string) // path is 'debug' by default

will enable:
	- /debug/pprof
	- /debug/profile
	- /debug/heap
	- /debug/trace
```
To execute profile cpu: `go tool pprof -http=":8000" pprofbin http://localhost:9313/debug/profile?seconds=18`
To execute profile memory: `go tool pprof -http=":8000" pprofbin http://localhost:9313/debug/heap?seconds=18`
To execute generate trace: go to endpoint `http://localhost:9313/debug/trace?seconds=18` from browser , this will download the trace of 18 seconds
Then to see the trace : `go tool trace path/to/trace`

### Metrics Prometheus
```go
// or srv := korm.WithBus()
//srv.WithMetrics(httpHandler http.Handler, path ...string) path default to 'metrics'
srv.WithMetrics(promhttp.Handler())

will enable:
	- /metrics
```

### Logs middleware
```go
// or srv := korm.WithBus()
//srv.WithMetrics(httpHandler http.Handler, path ...string) path default to 'metrics'
srv.App.Use(ksmux.Logs()) // it take an optional callback executed on each request if you want to add log to a file or send
srv.App.Use(ksmux.Logs(func(method, path, remote string, status int, took time.Duration) {
	// save somewhere
}))

will enable:
	- /metrics
```

# Hooks
```go
korm.OnInsert(func(database, table string, data map[string]any) error {
	fmt.Println("inserting into", database, table, data)
	// if error returned, it will not insert
	return nil
})

korm.OnSet(func(database, table string, data map[string]any) error {
	fmt.Println("set into", database, table, data)
	return nil
})

korm.OnDelete(func(database, table, query string, args ...any) error {})

korm.OnDrop(func(database, table string) error {})
```


## Python bus client example
```sh
pip install ksbus==1.1.0
# if it doesn't work , execute it again 
```
```py
from ksbus import Bus


# onOpen callback that let you know when connection is ready, it take the bus as param
def OnOpen(bus):
    print("connected")
    # bus.autorestart=True
    # Publish publish to topic
    bus.Publish("top", {
        "data": "hello from python"
    })
    # Subscribe, it also return the subscription
    bus.Subscribe("python", pythonTopicHandler)
    # SendToNamed publish to named topic
    bus.SendToNamed("top:srv", {
        "data": "hello again from python"
    })
    # bus.Unsubscribe("python")
    print("finish everything")


# pythonTopicHandler handle topic 'python'
def pythonTopicHandler(data, subs):
    print("recv on topic python:", data)
    # Unsubscribe
    #subs.Unsubscribe()

if __name__ == "__main__":
    Bus("localhost:9313", onOpen=OnOpen) # blocking
    print("prorgram exited")
```

# ManyToMany Relationships Example

```go
type Class struct {
	Id          uint   `korm:"pk"`
	Name        string `korm:"size:100"`
	IsAvailable bool
	CreatedAt   time.Time `korm:"now"`
}

type Student struct {
	Id        uint      `korm:"pk"`
	Name      string    `korm:"size:100"`
	CreatedAt time.Time `korm:"now"`
}

// migrate
func migrate() {
	err := korm.AutoMigrate[Class]("classes")
	if lg.CheckError(err) {
		return
	}
	err = korm.AutoMigrate[Student]("students")
	if lg.CheckError(err) {
		return
	}
	err = korm.ManyToMany("classes", "students")
	if lg.CheckError(err) {
		return
	}
}

// korm.ManyToMany create relation table named m2m_classes_students

// then you can use it like so to get related data

// get related to map to struct
std := []Student{}
err = korm.Model[Class]().Where("name = ?", "Math").Select("name").OrderBy("-name").Limit(1).GetRelated("students", &std)

// get related to map
std := []map[string]any{}
err = korm.Table("classes").Where("name = ?", "Math").Select("name").OrderBy("-name").Limit(1).GetRelated("students", &std)

// join related to map
std := []map[string]any{}
err = korm.Table("classes").Where("name = ?", "Math").JoinRelated("students", &std)

// join related to strcu
cu := []JoinClassUser{}
err = korm.Model[Class]().Where("name = ?", "Math").JoinRelated("students", &cu)

// to add relation
_, err = korm.Model[Class]().AddRelated("students", "name = ?", "hisName")
_, err = korm.Model[Student]().AddRelated("classes", "name = ?", "French")
_, err = korm.Table("students").AddRelated("classes", "name = ?", "French")

// delete relation
_, err = korm.Model[Class]().Where("name = ?", "Math").DeleteRelated("students", "name = ?", "hisName")
_, err = korm.Table("classes").Where("name = ?", "Math").DeleteRelated("students", "name = ?", "hisName")

```


### Swagger documentation

<img src="docs.png">
<div style="display:flex;justify-content:center;align-items:center;gap:20px;margin:20px 0">
<a href="https://www.youtube.com/watch?v=RupARTkPzf4">
	<img src="https://user-images.githubusercontent.com/54605903/217871012-9c5dc1da-25bd-47d5-ac9e-c3acee7178d5.svg" width="auto" height="50px">
</a>
</div>


```go
korm.DocsUrl = "docs" // default endpoint '/docs' 
korm.BASIC_AUTH_USER = "test"
korm.BASIC_AUTH_PASS = "pass"
korm.WithDocs(generate, dirPath, korm.BasicAuth)
korm.WithDocs(true, "", korm.BasicAuth) // dirPath default to 'assets/static/docs'
korm.WithEmbededDocs(embeded embed.FS, dirPath, korm.BasicAuth)
// dirPath default to 'assets/static/docs' if empty
```


### Interactive shell
```shell
Commands :  
[databases, use, tables, columns, migrate, createsuperuser, createuser, query, getall, get, drop, delete, clear/cls, q/quit/exit, help/commands]
  'databases':
	  list all connected databases

  'use':
	  use a specific database

  'tables':
	  list all tables in database

  'columns':
	  list all columns of a table
	  (accept but not required extra param like : 'columns' or 'columns users')

  'migrate':
	  migrate or execute sql file

  'createsuperuser': (only with dashboard)
	  create a admin user
  
  'createuser': (only with dashboard)
	  create a regular user

  'query': 
	  query data from database 
	  (accept but not required extra param like : 'query' or 'query select * from users where ...')


  'getall': 
	  get all rows given a table name
	  (accept but not required extra param like : 'getall' or 'getall users')

  'get':
	  get single row 
	  (accept but not required extra param like : 'get' or 'get users email like "%anything%"')

  'delete':
	  delete rows where field equal_to
	  (accept but not required extra param like : 'delete' or 'delete users email="email@example.com"')

  'drop':
	  drop a table given table name
	  (accept but not required extra param like : 'drop' or 'drop users')

  'clear / cls':
	  clear shell console

  'q / quit / exit / q!':
	  exit shell

  'help':
	  show this help message
```


# Example, not required, Load config from env directly to struct using Kenv
```go
import "github.com/kamalshkeir/kenv"

type EmbedS struct {
	Static    bool `kenv:"EMBED_STATIC|false"`
	Templates bool `kenv:"EMBED_TEMPLATES|false"`
}

type GlobalConfig struct {
	Host       string `kenv:"HOST|localhost"` // DEFAULT to 'localhost': if HOST not found in env
	Port       string `kenv:"PORT|9313"`
	Embed 	   EmbedS
	Db struct {
		Name     string `kenv:"DB_NAME|db"` // NOT REQUIRED: if DB_NAME not found, defaulted to 'db'
		Type     string `kenv:"DB_TYPE"` // REEQUIRED: this env var is required, you will have error if empty
		DSN      string `kenv:"DB_DSN|"` // NOT REQUIRED: if DB_DSN not found it's not required, it's ok to stay empty
	}
	Smtp struct {
		Email string `kenv:"SMTP_EMAIL|"`
		Pass  string `kenv:"SMTP_PASS|"`
		Host  string `kenv:"SMTP_HOST|"`
		Port  string `kenv:"SMTP_PORT|"`
	}
	Profiler   bool   `kenv:"PROFILER|false"`
	Docs       bool   `kenv:"DOCS|false"`
	Logs       bool   `kenv:"LOGS|false"`
	Monitoring bool   `kenv:"MONITORING|false"`
}


kenv.Load(".env") // load env file

// Fill struct from env loaded before:
Config := &GlobalConfig{}
err := kenv.Fill(Config) // fill struct with env vars loaded before
```

# Example nested or embeded structs

```go
package main

import (
	"fmt"
	"time"

	"github.com/kamalshkeir/lg"
	"github.com/kamalshkeir/korm"
	"github.com/kamalshkeir/sqlitedriver"
)

type Class struct {
	Id       uint `korm:"pk"`
	Name     string
	Students []Student
}

type Student struct {
	Id      uint `korm:"pk"`
	Name    string
	Class   uint `korm:"fk:classes.id:cascade:cascade"`
	Classes Class
}

func main() {
	err := korm.New(korm.SQLITE, "db", sqlitedriver.Use())
	if lg.CheckError(err) {
		return
	}
	defer korm.Shutdown()

	server := korm.WithDashboard("localhost:9313")
	korm.WithShell()

	err = korm.AutoMigrate[Class]("classes")
	lg.CheckError(err)

	err = korm.AutoMigrate[Student]("students")
	lg.CheckError(err)

	// go run main.go shell to createsuperuser
	// connect to admin and create some data to query

	// nested structs with joins, scan the result to the channel directly after each row
	// so instead of receiving a slice, you will receive data on the channel[0] of the passed slice
	studentsChan := []chan Student{make(chan Student)}
	go func() {
		for s := range studentsChan[0] {
			fmt.Println("chan students:", s)
		}
	}()
	err = korm.To(&studentsChan).Query("select students.*,classes.id as 'classes.id',classes.name as 'classes.name'  from students join classes where classes.id = students.class")
	lg.CheckError(err)
	fmt.Println()

	// nested (second argument of 'Scan') filled automatically from join, support nested slices and structs
	classes := []Class{}
	err = korm.To(&classes, true).Query("select classes.*, students.id as 'students.id',students.name as 'students.name' from classes join students on students.class = classes.id order by classes.id")
	lg.CheckError(err)
	for _, s := range classes {
		fmt.Println("class:", s)
	}
	fmt.Println()

	// // not nested, only remove second arg true from Scan method
	students := []Student{}
	err = korm.To(&students, true).Query("select students.*,classes.id as 'classes.id',classes.name as 'classes.name'  from students join classes where classes.id = students.class")
	lg.CheckError(err)
	for _, s := range students {
		fmt.Println("student:", s)
	}
	fmt.Println()

	maps := []map[string]any{}
	err = korm.To(&maps).Query("select * from students")
	lg.CheckError(err)
	fmt.Println("maps =", maps)
	fmt.Println()

	names := []*string{}
	err = korm.To(&names).Query("select name from students")
	lg.CheckError(err)
	fmt.Println("names =", names)
	fmt.Println()

	ids := []int{}
	err = korm.To(&ids).Query("select id from students")
	lg.CheckError(err)
	fmt.Println("ids =", ids)
	fmt.Println()

	bools := []bool{}
	err = korm.To(&bools).Query("select is_admin from users")
	lg.CheckError(err)
	fmt.Println("bools =", bools)
	fmt.Println()

	times := []time.Time{}
	err = korm.To(&times).Query("select created_at from users")
	lg.CheckError(err)
	fmt.Println("times =", times)

	server.Run()
}

// OUTPUT
// chan students: {1 student-1 1 {1 Math []}}
// chan students: {2 student-2 2 {2 French []}}
// chan students: {3 student-3 1 {1 Math []}}
// chan students: {4 student-4 2 {2 French []}}

// class: {1 Math [{1 student-1 0 {0  []}} {3 student-3 0 {0  []}}]}
// class: {2 French [{2 student-2 0 {0  []}} {4 student-4 0 {0  []}}]}

// student: &{1 student-1 1 {1 Math []}}
// student: &{2 student-2 2 {2 French []}}
// student: &{3 student-3 1 {1 Math []}}
// student: &{4 student-4 2 {2 French []}}

// maps = [map[class:1 id:1 name:student-1] map[class:2 id:2 name:student-2] map[class:1 id:3 name:student-3] map[class:2 id:4 name:student-4]]

// names = [student-1 student-2 student-3 student-4]

// ids = [1 2 3 4]

// bools = [true]

// times = [2023-04-30 19:19:32 +0200 CEST]

```


# Benchmark vs Tarantool, Pgx, Gorm

[https://github.com/kamalshkeir/korm-vs-gorm-vs-tarantool-vs-pgx](https://github.com/kamalshkeir/korm-vs-gorm-vs-tarantool-vs-pgx)

# Benchmarks vs Gorm
```sh
goos: windows
goarch: amd64
pkg: github.com/kamalshkeir/korm/benchmarks
cpu: Intel(R) Core(TM) i5-7300HQ CPU @ 2.50GHz
```

To execute these benchmarks on your machine, very easy :

- git clone https://github.com/kamalshkeir/korm.git
- cd korm
- uncomment commented code at benchmarks/bench-test.go and Save
- go mod tidy
- go test -bench ^ .\benchmarks\ -benchmem

```go
type TestTable struct {
	Id        uint `korm:"pk"`
	Email     string
	Content   string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time `korm:"now"`
	UpdatedAt time.Time `korm:"update"`
}

type TestTableGorm struct {
	Id        uint `gorm:"primarykey"`
	Email     string
	Content   string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
////////////////////////////////////////////  query 7000 rows  //////////////////////////////////////////////
BenchmarkGetAllS_GORM-4                       19          56049832 ns/op        12163316 B/op     328790 allocs/op
BenchmarkGetAllS-4                       2708934               395.3 ns/op           224 B/op          1 allocs/op
BenchmarkGetAllM_GORM-4                       18          62989567 ns/op        13212278 B/op     468632 allocs/op
BenchmarkGetAllM-4                       4219461               273.5 ns/op           224 B/op          1 allocs/op
BenchmarkGetRowS_GORM-4                    12188             96988 ns/op            5930 B/op        142 allocs/op
BenchmarkGetRowS-4                       1473164               805.1 ns/op           336 B/op          7 allocs/op
BenchmarkGetRowM_GORM-4                    11402            101638 ns/op            7408 B/op        203 allocs/op
BenchmarkGetRowM-4                       1752652               671.9 ns/op           336 B/op          7 allocs/op
BenchmarkPagination10_GORM-4                7714            153304 ns/op           19357 B/op        549 allocs/op
BenchmarkPagination10-4                  1285722               934.5 ns/op           400 B/op          7 allocs/op
BenchmarkPagination100_GORM-4               1364            738934 ns/op          165423 B/op       4704 allocs/op
BenchmarkPagination100-4                 1278724               956.5 ns/op           400 B/op          7 allocs/op
BenchmarkQueryS-4                        5781499               207.7 ns/op             4 B/op          1 allocs/op
BenchmarkQueryM-4                        4643155               227.2 ns/op             4 B/op          1 allocs/op
BenchmarkGetAllTables-4                 47465865                25.48 ns/op            0 B/op          0 allocs/op
BenchmarkGetAllColumns-4                23657019                42.82 ns/op            0 B/op          0 allocs/op
////////////////////////////////////////////  query 5000 rows  //////////////////////////////////////////////
BenchmarkGetAllS_GORM-4                       24          43247546 ns/op         8796840 B/op     234784 allocs/op
BenchmarkGetAllS-4                       2854401               426.8 ns/op           224 B/op          1 allocs/op
BenchmarkGetAllM_GORM-4                       24          46329242 ns/op         9433050 B/op     334631 allocs/op
BenchmarkGetAllM-4                       4076317               283.4 ns/op           224 B/op          1 allocs/op
BenchmarkGetRowS_GORM-4                    11445            101107 ns/op            5962 B/op        142 allocs/op
BenchmarkGetRowS-4                       1344831               848.4 ns/op           336 B/op          7 allocs/op
BenchmarkGetRowM_GORM-4                    10000            100969 ns/op            7440 B/op        203 allocs/op
BenchmarkGetRowM-4                       1721742               688.5 ns/op           336 B/op          7 allocs/op
BenchmarkPagination10_GORM-4                7500            156208 ns/op           19423 B/op        549 allocs/op
BenchmarkPagination10-4                  1253757               952.3 ns/op           400 B/op          7 allocs/op
BenchmarkPagination100_GORM-4               1564            749408 ns/op          165766 B/op       4704 allocs/op
BenchmarkPagination100-4                 1236270               957.5 ns/op           400 B/op          7 allocs/op
BenchmarkGetAllTables-4                 44399386                25.43 ns/op            0 B/op          0 allocs/op
BenchmarkGetAllColumns-4                27906392                41.45 ns/op            0 B/op          0 allocs/op
////////////////////////////////////////////  query 1000 rows  //////////////////////////////////////////////
BenchmarkGetAllS_GORM-4                      163           6766871 ns/op         1683919 B/op      46735 allocs/op
BenchmarkGetAllS-4                       2882660               399.0 ns/op           224 B/op          1 allocs/op
BenchmarkGetAllM_GORM-4                      140           8344988 ns/op         1886922 B/op      66626 allocs/op
BenchmarkGetAllM-4                       3826730               296.5 ns/op           224 B/op          1 allocs/op
BenchmarkGetRowS_GORM-4                    11940             97725 ns/op            5935 B/op        142 allocs/op
BenchmarkGetRowS-4                       1333258               903.0 ns/op           336 B/op          7 allocs/op
BenchmarkGetRowM_GORM-4                    10000            106079 ns/op            7408 B/op        203 allocs/op
BenchmarkGetRowM-4                       1601274               748.2 ns/op           336 B/op          7 allocs/op
BenchmarkPagination10_GORM-4                7534            159991 ns/op           19409 B/op        549 allocs/op
BenchmarkPagination10-4                  1153982              1022 ns/op             400 B/op          7 allocs/op
BenchmarkPagination100_GORM-4               1468            766269 ns/op          165876 B/op       4705 allocs/op
BenchmarkPagination100-4                 1000000              1016 ns/op             400 B/op          7 allocs/op
BenchmarkGetAllTables-4                 56200297                25.36 ns/op            0 B/op          0 allocs/op
BenchmarkGetAllColumns-4                25478679                41.30 ns/op            0 B/op          0 allocs/op
////////////////////////////////////////////  query 300 rows  //////////////////////////////////////////////
BenchmarkGetAllS_GORM-4                      558           2046830 ns/op          458475 B/op      13823 allocs/op
BenchmarkGetAllS-4                       2798872               411.5 ns/op           224 B/op          1 allocs/op
BenchmarkGetAllM_GORM-4                      428           2605646 ns/op          567011 B/op      19721 allocs/op
BenchmarkGetAllM-4                       4093662               287.9 ns/op           224 B/op          1 allocs/op
BenchmarkGetRowS_GORM-4                    12182             97764 ns/op            5966 B/op        142 allocs/op
BenchmarkGetRowS-4                       1347084               886.4 ns/op           336 B/op          7 allocs/op
BenchmarkGetRowM_GORM-4                    10000            105311 ns/op            7440 B/op        203 allocs/op
BenchmarkGetRowM-4                       1390363               780.0 ns/op           336 B/op          7 allocs/op
BenchmarkPagination10_GORM-4                7502            155949 ns/op           19437 B/op        549 allocs/op
BenchmarkPagination10-4                  1000000              1046 ns/op             400 B/op          7 allocs/op
BenchmarkPagination100_GORM-4               1479            779700 ns/op          165679 B/op       4705 allocs/op
BenchmarkPagination100-4                 1000000              1054 ns/op             400 B/op          7 allocs/op
BenchmarkGetAllTables-4                 52255704                26.00 ns/op            0 B/op          0 allocs/op
BenchmarkGetAllColumns-4                29292368                42.09 ns/op            0 B/op          0 allocs/op
////////////////////////////////////////////    MONGO       //////////////////////////////////////////////
BenchmarkGetAllS-4               3121384               385.6 ns/op           224 B/op          1 allocs/op
BenchmarkGetAllM-4               4570059               264.2 ns/op           224 B/op          1 allocs/op
BenchmarkGetRowS-4               1404399               866.6 ns/op           336 B/op          7 allocs/op
BenchmarkGetRowM-4               1691026               722.6 ns/op           336 B/op          7 allocs/op
BenchmarkGetAllTables-4         47424489                25.34 ns/op            0 B/op          0 allocs/op
BenchmarkGetAllColumns-4        27039632                42.22 ns/op            0 B/op          0 allocs/op
//////////////////////////////////////////////////////////////////////////////////////////////////////////
```



---
### Available Tags by struct field type:

# String Field:
<table>
<tr>
<th>Without parameter&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</th>
<th>With parameter&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</th>
</tr>
<tr>
<td>
 
```
*  	text (create column as TEXT not VARCHAR)
*  	notnull
*  	unique
*   iunique // insensitive unique
*  	index, +index, index+ (INDEX ascending)
*  	index-, -index (INDEX descending)
*  	default (DEFAULT '')
```
</td>
<td>

```
* 	default:'any' (DEFAULT 'any')
*	mindex:...
* 	uindex:username,Iemail // CREATE UNIQUE INDEX ON users (username,LOWER(email)) 
	// 	email is lower because of 'I' meaning Insensitive for email
* 	fk:...
* 	size:50  (VARCHAR(50))
* 	check:...
```

</td>
</tr>
</table>


---



# Int, Uint, Int64, Uint64 Fields:
<table>
<tr>
<th>Without parameter&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</th>
</tr>
<tr>
<td>
 
```
*   -  			 (To Ignore a field)
*   autoinc, pk  (PRIMARY KEY)
*   notnull      (NOT NULL)
*  	index, +index, index+ (CREATE INDEX ON COLUMN)
*  	index-, -index(CREATE INDEX DESC ON COLUMN)     
*   unique 		 (CREATE UNIQUE INDEX ON COLUMN) 
*   default		 (DEFAULT 0)
```
</td>
</tr>

<tr><th>With parameter&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</th></tr>
<tr>
<td>

```
Available 'on_delete' and 'on_update' options: cascade,(donothing,noaction),(setnull,null),(setdefault,default)

*   fk:{table}.{column}:{on_delete}:{on_update} 
*   check: len(to_check) > 10 ; check: is_used=true (You can chain checks or keep it in the same CHECK separated by AND)
*   mindex: first_name, last_name (CREATE MULTI INDEX ON COLUMN + first_name + last_name)
*   uindex: first_name, last_name (CREATE MULTI UNIQUE INDEX ON COLUMN + first_name + last_name) 
*   default:5 (DEFAULT 5)
```

</td>
</tr>
</table>

---


# Bool : bool is INTEGER NOT NULL checked between 0 and 1 (in order to be consistent accross sql dialects)
<table>
<tr>
<th>Without parameter&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</th>
<th>With parameter&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</th>
</tr>
<tr>
<td>
 
```
*  	index, +index, index+ (CREATE INDEX ON COLUMN)
*  	index-, -index(CREATE INDEX DESC ON COLUMN)  
*   default (DEFAULT 0)
```
</td>
<td>

```
*   default:1 (DEFAULT 1)
*   mindex:...
*   fk:...
```

</td>
</tr>
</table>

---

# time.Time :
<table>
<tr>
<th>Without parameter&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</th>
<th>With parameter</th>
</tr>
<tr>
<td>
 
```
*  	index, +index, index+ (CREATE INDEX ON COLUMN)
*  	index-, -index(CREATE INDEX DESC ON COLUMN)  
*   now (NOT NULL and defaulted to current unix timestamp)
*   update (NOT NULL DEFAULT UNIX_TIMESTAMP ON UPDATE UNIX_TIMESTAMP)
```
</td>
<td>

```
*   fk:...
*   check:...
```

</td>
</tr>
</table>

---

# Float64 :
<table>
<tr>
<th>Without parameter&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</th>
<th>With parameter&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</th>
</tr>
<tr>
<td>
 
```
*   notnull
*  	index, +index, index+ (CREATE INDEX ON COLUMN)
*  	index-, -index(CREATE INDEX DESC ON COLUMN)  
*   unique
*   default
```
</td>
<td>

```
*   default:...
*   fk:...
*   mindex:...
*   uindex:...
*   check:...
```

</td>
</tr>
</table>



# JSON

```go
type JsonOption struct {
	As       string
	Dialect  string
	Database string
	Params   []any
}
func JSON_EXTRACT(dataJson string, opt ...JsonOption) string
func JSON_REMOVE(dataJson string, opt ...JsonOption) string
func JSON_SET(dataJson string, opt ...JsonOption) string
func JSON_ARRAY(values []any, as string, dialect ...string) string
func JSON_OBJECT(values []any, as string, dialect ...string) string
func JSON_CAST(value string, as string, dialect ...string) string

// create query json
q := korm.JSON_EXTRACT(`{"a": {"c": 3}, "b": 2}`, korm.JsonOption{
	As:     "data",
	Params: []any{"a.c", "b"},
})
fmt.Println("q ==", q) // q == JSON_EXTRACT('{"a": {"c": 3}, "b": 2}','$.a.c','$.b') AS data

var data []map[string]any
err := korm.To(&data).Query("SELECT " + q)
lg.CheckError(err)

fmt.Println("data=", data) // data= [map[data:[3,2]]]
```


---


# 🔗 Links
[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://kamalshkeir.dev/) [![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/kamal-shkeir/)


---

# Licence
Licence [BSD-3](./LICENSE)
