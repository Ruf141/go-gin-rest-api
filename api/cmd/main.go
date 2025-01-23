package main

import (
	"context"

	"github.com/o-ga09/api/internal/presenter"
)

func main() {
	srv := presenter.NewServer()
	if err := srv.Run(context.Background()); err != nil {
		panic(err)
	}

}
