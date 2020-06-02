# go-restapi-boilerplate

This is an explanation of how I write "json rest api server" in go. I wrote this to share with my friends who are not familiar with go, but thought it would be great to get feedbacks from the community also. 
Newbie in go myself, I won't say that this is way to go. But after spending hours to write rest api that is simple, easy to manage, I'm quite satisfied with the result.

This includes extremely simple boilerplate and example for
- openapi v3 integration(server gen, documents)
- contextual, structured logging
- orm
- access control
- metrics
- graceful shutdown
- configuration through flag, env. (no files, but viper also can load configuration from files)
- docker image: [acadx0/go-restapi-boilerplate](https://hub.docker.com/repository/docker/acadx0/go-restapi-boilerplate)
- kubernetes deploy with kustomize
- ...

Usage
---
- **Write spec first** (check details below)
- ```go generate ./...``` 
- write your code.
```
go run cmd/api/main.go
```
Boilerplate includes CRUD for user as example. Check [spec/swagger.yaml](spec/swagger.yaml)
```
 » http 'localhost:8080/api/v1/user' user_name=aca user_id=acadx0
{
    "id": 1,
    "user_id": "acadx0",
    "user_name": "aca"
}

 » http 'localhost:8080/api/v1/user/acadx0'
{
    "id": 1,
    "user_id": "acadx0",
    "user_name": "aca"
}
```

---
Define kubernetes spec with kustomize, check [deploy](deploy).
```
kustomize build prod | kubectl apply -f -
```


Structure
---
```
├── api // put all handlers here.
│   ├── api.gen.go
│   ├── api.go 
│   ├── config.go 
│   ├── error.go
│   ├── metrics.go
│   └── user.go // I usually separate file by resource.
├── cmd // put any executable here
│   └── api
├── pkg // library code.
├── ent
│   └── schema // write your ent schema here, I often put other models in here.
└── spec
    └── swagger.yaml

```

Libraries
---
Here are the libraries I chose. Some of them are relatively new and may not be mature compared to the competitors. I had to make multiple patches to the listed library to satisfy my use cases. But at least for me, they really helped me to simplify the process. Hope they all have more users and contributors. 

- **Openapi integration for documents, client/server codegen.** [oapi-codegen](https://github.com/deepmap/oapi-codegen), [kin-openapi](https://github.com/getkin/kin-openapi)  
  API service is meaningless without document. But managing documents and code separately gets really messy when your service grows.  

  "OpenAPI Specification" defines the standard way to manage your REST API service.
  Instead of writing code first, write openapi spec(which is extremely simple) with [swagger editor](https://editor.swagger.io/) first and verify your api. Then, generate code from spec. For client code generation, I recommend [openapi-generator](https://github.com/OpenAPITools/openapi-generator). I used it to generate typescript client for my react app. It's just perfect.

  For server code generation, [go-swagger](https://github.com/go-swagger/go-swagger) seems popular and mature. But It doesn't support openapi spec v3(*might not matter that much*), and I find it too complicated. [oapi-codegen](https://github.com/deepmap/oapi-codegen) on the otherhand, generates extremely simple go code that you probably won't even have to read the document. It just generates types, and server interface in the name of "OperationID" in your spec.

  In this boilerplate, I generate "openapi components" from ent model. 
  And based on updated [swagger.yaml](./spec/swagger.yaml), generate chi-server (but with no types because I just directly use ent generated model).  
  So when [path.yaml](spec/path.yaml) or ent/schema is updated, run ```go generate ./...``` to generate code.


- **Web Framework.** net/http with [chi](https://github.com/go-chi/chi)  
  I don't want to learn another framework because I never thought I need it. I believe net/http is enough and complete when writing http service in go. Only thing it lacks is probably router and few helper functions / middlewares. chi has 100% compatibility with net/http. gorilla/mux is more famous for this but, as oapi-codegen only supports chi, I just use chi. 

  If you use standard net/http handler, It is extremely simple to integrate with third party middlewares like [rs/cors](https://github.com/rs/cors), [zerolog](https://github.com/rs/zerolog).


- **ORM.**  [facebookincubator/ent](https://github.com/facebookincubator/ent)  
  There's lots of ORM in go out there. But If you ever felt that there's something wrong with it, try [ent](https://github.com/facebookincubator/ent). 
  ```
    db.Where("name <> ?", "aca").Find(&users) // gorm
    db.User.Query().Where(user.NameEQ("aca")).All(ctx) // ent
  ```
   Ent generates 100% statically typed go code. With ent, I was able to write code which I felt much more solid. It's amazing project that changed my mind on "orm in go", you should definitely check. 

  
- **Contextual, structured logging.** [zerolog](https://github.com/rs/zerolog)  
  zerolog offers simplest api. It also offers helper library(hlog) that can be used with standard ```http.Handler```, and ```context.Context``` integration is amazing. 
  ``` 
  ctx := log.With().Str("component", "module").Logger().WithContext(ctx)

  // ... somewhere in your function with context
  log.Ctx(ctx).Info().Msg("hello world")
  // Output: {"component":"module","level":"info","message":"hello world"} 
  ```

- **Configuration.** [spf13/viper](https://github.com/spf13/viper)  

- **Metrics.** - [Prometheus](https://github.com/prometheus/client_golang)  
  It is extremely easy to add custom metrics to your server, check [api/metrics.go](api/metrics.go).
  
- **Access Control.** [ Open Policy Agent ](https://github.com/open-policy-agent/opa)  
  It is not included in this boilerplate. But I recommend opa. I write rules in [REGO](https://play.openpolicyagent.org/), and add opa middlewares between authentication middleware, and handler. It was not easy to learn new "OPA", "REGO" thing. But I use it to define complex rules and control my api without changing my code. The point is to decouple access control with business logic.

