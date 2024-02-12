package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: <program> <mongo_uri> <db_name>")
		os.Exit(1)
	}
	mongo_uri := os.Args[1]
	db_name := os.Args[2]

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_uri))
	if err != nil {
		fmt.Printf("Failed to connect to MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(context.Background())

	db := client.Database(db_name)
	players_collection := db.Collection("players")
	packs_collection := db.Collection("packs")
	games_collection := db.Collection("games")
	rounds_collection := db.Collection("rounds")
	responses_collection := db.Collection("responses")

	gofakeit.Seed(0)

	player_ids := insert_players(players_collection, 10)
	pack_ids, err := get_random_pack_ids(packs_collection, 3)
	if err != nil {
		fmt.Println("Error fetching pack IDs:", err)
		return
	}
	game_id, err := insert_game(games_collection, player_ids, pack_ids)
	if err != nil {
		fmt.Println("Error inserting game:", err)
		return
	}
	err = insert_rounds(rounds_collection, responses_collection, game_id, player_ids, 10)
	if err != nil {
		fmt.Println("Error inserting rounds:", err)
		return
	}

	fmt.Println("Data insertion complete.")
}

func insert_players(collection *mongo.Collection, count int) []string {
	player_ids := make([]string, count)
	for i := 0; i < count; i++ {
		player := bson.M{
			"name":       gofakeit.Name(),
			"created_at": time.Now(),
			"updated_at": time.Now(),
		}
		result, err := collection.InsertOne(context.Background(), player)
		if err != nil {
			panic(err)
		}
		player_ids[i] = result.InsertedID.(primitive.ObjectID).Hex()
	}
	return player_ids
}

func get_random_pack_ids(collection *mongo.Collection, count int) ([]string, error) {
	// Define the pipeline with the $sample stage
	pipeline := mongo.Pipeline{
		{{"$sample", bson.D{{"size", count}}}},
	}

	cursor, err := collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}
	var packs []bson.M
	if err = cursor.All(context.Background(), &packs); err != nil {
		return nil, err
	}
	pack_ids := make([]string, len(packs))
	for i, pack := range packs {
		pack_ids[i] = pack["_id"].(primitive.ObjectID).Hex() // Convert ObjectID to string
	}
	return pack_ids, nil
}

func insert_game(collection *mongo.Collection, player_ids []string, pack_ids []string) (string, error) {
	created_by_index := rand.Intn(len(player_ids))
	game := bson.M{
		"name":       "New Game",
		"packs":      pack_ids,
		"started_at": time.Now(),
		"ended_at":   time.Now().Add(1 * time.Hour),
		"max_rounds": 5,
		"players":    player_ids,
		"banned":     []interface{}{},
		"winner":     nil,
		"created_by": player_ids[created_by_index],
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}
	result, err := collection.InsertOne(context.Background(), game)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func insert_rounds(rounds_collection *mongo.Collection, responses_collection *mongo.Collection, game_id string, player_ids []string, total_rounds int) error {
	rand.Seed(time.Now().UnixNano()) // Ensure randomness

	for round_num := 1; round_num <= total_rounds; round_num++ {
		// Placeholder: Simulate fetching a black card ID and its pick count (1 or 2)
		black_card_id := primitive.NewObjectID().Hex()
		pick_count := rand.Intn(2) + 1

		// Insert the round document without responses
		round_doc := bson.M{
			"game_id":    game_id,
			"dealer":     player_ids[round_num%len(player_ids)], // Rotate dealer each round
			"challenge":  black_card_id,
			"winner":     "", // Temporarily empty, will be updated later
			"started_at": time.Now(),
			"ended_at":   time.Now().Add(time.Duration(rand.Intn(10)) * time.Minute),
			"num":        round_num,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		}
		round_result, err := rounds_collection.InsertOne(context.Background(), round_doc)
		if err != nil {
			return err
		}
		round_id := round_result.InsertedID.(primitive.ObjectID).Hex()

		// Placeholder winner logic; actual winner determination logic should be implemented
		winner_index := rand.Intn(10)

		// Insert responses as separate documents
		for i := 0; i < 10; i++ { // Ensuring at least 10 responses
			player_index := i % len(player_ids)
			white_card_ids := make([]string, pick_count)
			for j := range white_card_ids {
				// Placeholder: Simulate fetching a white card ID
				white_card_ids[j] = primitive.NewObjectID().Hex()
			}

			response_doc := bson.M{
				"round_id":  round_id,
				"player_id": player_ids[player_index],
				"cards":     white_card_ids,
			}
			_, err := responses_collection.InsertOne(context.Background(), response_doc)
			if err != nil {
				return err
			}

			// Update the winner in the round document based on placeholder logic
			if i == winner_index {
				_, err := rounds_collection.UpdateOne(
					context.Background(),
					bson.M{"_id": round_result.InsertedID},
					bson.M{"$set": bson.M{"winner": player_ids[player_index]}},
				)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
