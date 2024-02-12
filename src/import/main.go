package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	if len(os.Args) < 4 {

		fmt.Println("Usage: import ../../resources/data/cards.json mongodb://user:pass@localhost:27017/db_name db_name")

		os.Exit(1)
	}

	json_file, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Printf("Error reading JSON file: %v\n", err)
		return
	}

	var packs []struct {
		Name  string `json:"name"`
		White []struct {
			Text string `json:"text"`
			Pack int    `json:"pack"`
		} `json:"white"`
		Black []struct {
			Text string `json:"text"`
			Pick int    `json:"pick"`
			Pack int    `json:"pack"`
		} `json:"black"`
	}

	err = json.Unmarshal(json_file, &packs)

	if err != nil {

		fmt.Printf("Error unmarshalling JSON file: %v\n", err)

		return
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Args[2]))

	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		return
	}

	defer client.Disconnect(context.Background())

	packs_collection := client.Database(os.Args[3]).Collection("packs")
	cards_collection := client.Database(os.Args[3]).Collection("cards")

	total_inserts := len(packs)

	fmt.Println("Starting inserts...")

	for i, pack := range packs {

		pack_id_result, err := packs_collection.InsertOne(context.TODO(), bson.M{
			"en":         pack.Name,
			"es":         nil,
			"fr":         nil,
			"de":         nil,
			"it":         nil,
			"pt":         nil,
			"total":      0,
			"created_by": nil,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		})

		if err != nil {

			fmt.Printf("Error inserting pack: %v\n", err)

			continue
		}

		pack_id_str := pack_id_result.InsertedID.(primitive.ObjectID).Hex()
		total_cards := 0

		for _, card := range pack.White {

			_, err = cards_collection.InsertOne(context.TODO(), bson.M{
				"pack_id":    pack_id_str,
				"type":       "white",
				"en":         card.Text,
				"es":         nil,
				"fr":         nil,
				"de":         nil,
				"it":         nil,
				"pt":         nil,
				"created_by": nil,
				"created_at": time.Now(),
				"updated_at": time.Now(),
			})

			if err != nil {

				fmt.Printf("Error inserting white card: %v\n", err)

			} else {

				total_cards++

			}
		}

		for _, card := range pack.Black {

			card_text := strings.ReplaceAll(card.Text, " _", " _____")

			_, err = cards_collection.InsertOne(context.TODO(), bson.M{
				"pack_id":    pack_id_str,
				"type":       "black",
				"en":         card_text,
				"es":         nil,
				"fr":         nil,
				"de":         nil,
				"it":         nil,
				"pt":         nil,
				"pick":       card.Pick,
				"created_by": nil,
				"created_at": time.Now(),
				"updated_at": time.Now(),
			})

			if err != nil {

				fmt.Printf("Error inserting black card: %v\n", err)

			} else {

				total_cards++

			}
		}

		_, err = packs_collection.UpdateOne(
			context.TODO(),
			bson.M{"_id": pack_id_result.InsertedID},
			bson.M{"$set": bson.M{"total": total_cards}},
		)

		if err != nil {

			fmt.Printf("Error updating pack with total card count: %v\n", err)

		}

		if (i+1)%((total_inserts/100)+1) == 0 {

			fmt.Print(".")

		}
	}

	fmt.Println("\nInserts completed.")
}
