package module2

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

//Hello call name
func Hello(name string) string {
	return fmt.Sprintf("Hello %v", name)
}
