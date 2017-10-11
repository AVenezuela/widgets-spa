package sqlconn

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

type server struct {
	db *sql.DB
}

func (s *server) handler(w http.ResponseWriter, r *http.Request) {
	// slow 5 seconds query
	_, err := s.db.Exec("SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ok"))
}

func (s *server) handlerCtx(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 1*time.Second)

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("ctx ok"))
}

func (s *server) handlerDisconnect(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)

	// in case of client disconnect, cancel context
	if cn, ok := w.(http.CloseNotifier); ok {
		go func() {
			<-cn.CloseNotify()
			cancelFunc()
		}()
	}

	// slow 5 seconds query
	_, err := s.db.ExecContext(ctx, "SELECT pg_sleep(5)")
	if err != nil {
		log.Println("[ERROR]", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("disconnected ok"))
}

func main() {
	db, err := sql.Open("sqlserver", "Data Source=(localdb)\\MSSQLLocalDB;Initial Catalog=Widgets;user=widgetuser;password=widget@123;")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := server{db: db}

	http.HandleFunc("/", s.handler)
	http.HandleFunc("/ctx", s.handlerCtx)
	http.HandleFunc("/disconnect", s.handlerDisconnect)
	log.Println("Starting server on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
