// package red

// import (
// 	"context"
// 	"fmt"
// 	"route-redis/database"
// 	"route-redis/models"
// 	"strconv"

// 	//"route-redis/dummy"
// 	"route-redis/models"

// 	"github.com/RediSearch/redisearch-go/redisearch"
// 	"github.com/redis/go-redis/v9"
// )

// func stationBuild() {
// 	// Connect to Redis server
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 	})

// 	//c := redisearch.NewClient("localhost:6379", "myIndex")

// 	// Connect to Redisearch index
// 	indexName := "stations"
// 	client := redisearch.NewClient("localhost:6379", indexName)

// 	// Search for stations by ID
// 	db := database.DB
// 	tx := db.Model(&models.Station{}).Order("id desc")
// 	stations := make([]models.Station, 0)

// 	err := tx.Select("*").Find(&stations).Error
// 	if err != nil {
// 		panic(err)
// 	}

// 	ctx := context.Background()

// 	// Index stations in Redisearch
// 	for _, station := range stations {
// 		doc := redisearch.NewDocument(strconv.Itoa(int(station.ID)), 1.0).
// 			Set("id", station.ID).
// 			Set("bus_route_id", station.BusRouteID).
// 			Set("bus_stop_name", station.BusStopName)

// 		if err := client.Index(doc); err != nil {
// 			panic(err)
// 		}
// 	}

// 	// Perform a search query
// 	query := redisearch.NewQuery("bus_route_id:108").Limit(0, 10)
// 	res, err := client.Search(query)
// 	if err != nil {
// 		panic(err)
// 	}

//		fmt.Printf("Total results: %d\n", res.TotalResults)
//		for _, doc := range res.Docs {
//			fmt.Printf("ID: %s, BusRouteID: %s, BusStopName: %s\n", doc.ID, doc.Properties["bus_route_id"], doc.Properties["bus_stop_name"])
//		}
//	}
package red
