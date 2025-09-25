package main

import (
	"bufio"
	"github-cli/github"
	"os"
	"strings"

	"fmt"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("\nWelcome to the Github User Activity CLI")

		fmt.Println("1. Fetch User Activity")

		fmt.Println("2. Exit")

		fmt.Println("Choose an option:")

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		switch input {

		case "1":

			fmt.Println("Enter Github Username: ")

			username, _ := reader.ReadString('\n')

			username = strings.TrimSpace(username)

			if username == "" {

				fmt.Println("User cannot be empty:")

				continue
			}
			fmt.Printf("\nFetching the recent activities for %s...\n\n", username)
			events, err := github.FetchUserActivity(username)

			if err != nil {
				fmt.Println("Error", err)
				continue
			}
			if len(events) == 0 {
				fmt.Println("No recent activity found")
				continue
			}

			for _, event := range events {
				switch event.Type {
				case "PushEvent":
					fmt.Printf("%s pushed to %s\n", event.Actor.Login, event.Repo.Name)
				case "Issues Event":
					fmt.Printf("%s interacted with issues in %s\n", event.Actor.Login, event.Repo.Name)
				case "WatchEvent":
					fmt.Printf("%s starred %s\n", event.Actor.Login, event.Repo.Name)
				default:
					fmt.Printf("%s did %s in %s\n", event.Actor.Login, event.Type, event.Repo.Name)
				}
			}

		case "2":
			fmt.Println("goodbye...exiting...")
			return
		default:
			fmt.Println("Invalid Option..Try Again?")

		}
	}
}
