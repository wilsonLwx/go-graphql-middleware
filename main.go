package main

import (
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)



// 定义日志response的状态码和状态内容结构体
type LoggerResponseWriter struct {
	http.ResponseWriter
	statusCode int
	statusText string
}

// 结构体LoggerResponseWriter实现http.ResponseWriter接口的WriteHeader方法
func (lw *LoggerResponseWriter) WriteHeader(code int) {
	lw.statusCode = code
	lw.statusText = http.StatusText(code)
	lw.ResponseWriter.WriteHeader(code)
}


func NewLoggerResponseWriter(w http.ResponseWriter) *LoggerResponseWriter{
	return &LoggerResponseWriter{w, http.StatusOK, http.StatusText(http.StatusOK)}
}


type NewHandler  struct {
    handler http.Handler
}

//ServeHTTP 通过将原请求handler传入然后调用，实现对请求的封装
//处理并打印请求的细节信息
func (l *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
	lw := NewLoggerResponseWriter(w)
    l.handler.ServeHTTP(lw, r)
    log.Printf("\"%s %s %s\" %v %d %d", r.Method, r.URL.Path, r.Proto, time.Since(start), lw.statusCode, r.ContentLength)
}

//Logger返回NewHandler实例
func Logger(h http.Handler) *NewHandler {
    return &NewHandler{h}
}

func main() {
	//定义返回参数
	var studentsType = graphql.NewObject(
		graphql.ObjectConfig{
		 Name: "students",
		 Fields: graphql.Fields{
		   "id": &graphql.Field{
			Type: graphql.Int,
		   },
		   "name": &graphql.Field{
			Type: graphql.String,
		   },
		 },
		},
	  )
	
	resolve_students := func (p graphql.ResolveParams) (interface{}, error) {
		// 类型断言
		id, _ := p.Args["id"].(int)
		name, _ := p.Args["name"].(string)
		
		ret := []map[string] interface{}{}
		ret = append(ret, map[string]interface{}{
			"id": 1,
			"name": "zhangsan",
		}, map[string]interface{}{
			"id": 2,
			"name": "lisi",
		})

		for _, item := range(ret){
			if item["id"] == id && item["name"] == name {
				return item, nil
			}
			
		}
		return nil, nil	
	}


	// 定义请求体
	fields := graphql.Fields{
		"students": &graphql.Field{
			// 将返回参数类型加入请求体
			Type:  studentsType,
			// Type:  graphql.NewList(studentsType),
			// 定义请求参数
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			// 定义接口函数 业务逻辑在接口函数中实现
			Resolve: resolve_students,
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "Query", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	
	// 用graphql-go/handler中间件处理graphql的http请求
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		GraphiQL: true,
	})

	http.Handle("/graphql", Logger(h))
	http.ListenAndServe(":9090", nil)
	
	// mux := http.NewServeMux()
	// mux.Handle("/graphql", h)
	// wrappedMux := NewLogger(mux)
	// log.Fatal(http.ListenAndServe(":9090", wrappedMux))
	
}
