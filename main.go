package main

type application struct {
	*ServiceLogger
}

func main() {
	//addr := flag.String("addr", ":4000", "HTTP network address")
	//flag.Parse()
	//
	//slog := NewServiceLogger()
	//app := &application{
	//	slog,
	//}
	//
	//srv := &http.Server{
	//	Addr:     *addr,
	//	ErrorLog: app.errorLog,
	//	Handler:  app.routes(),
	//}
	//
	//app.infoLog.Printf("Starting server on %s", *addr)
	//err := srv.ListenAndServe()
	//app.errorLog.Fatal(err)

	jsonUnmarshal()

}
