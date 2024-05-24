package logic

import "github.com/cufee/shopping-list/prisma/db"

func NewDatabaseClient() (*db.PrismaClient, error) {
	client := db.NewClient()
	err := client.Connect()
	if err != nil {
		return nil, err
	}

	return client, err
}
