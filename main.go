package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/DenisNosik/gator/internal/config"
	"github.com/DenisNosik/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("error connecting to db: %v", err)
		os.Exit(1)
	}
	defer db.Close()
	dbQueries := database.New(db)

	st := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		cmds: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("feeds", handlerFeeds)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: not enough arguments were provided")
		os.Exit(1)
	}

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	if err := cmds.run(st, cmd); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
