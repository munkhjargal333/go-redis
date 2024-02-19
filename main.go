package main

import (
	"context"
	"fmt"
	"route-redis/database"
	"strconv"
	"strings"
	"time"

	"route-redis/models"
	"route-redis/shared"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func main() {
	shared.LoadConfig()
	database.Connect()

	app := fiber.New()

	//runSetup()

	ticker := time.NewTicker(30 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				runSetup()
			}
		}
	}()

	setRoutes(app)

	app.Listen(":3000")
}

func runSetup() {
	clear()
	routeBuild()
	stationBuild()
	routeLineBuild()
}

func setRoutes(app *fiber.App) {
	app.Get("", HelloWrold)
	app.Get("search-route", searchRoute)
	app.Get("search-station", searchStation)

	api := app.Group("id")
	api.Get("route-stations", routeStations)
	api.Get("route-lines", routeLines)
	api.Get("station-routes", stationRoute)
}

func HelloWrold(c *fiber.Ctx) error {
	return shared.Response(c, "hello world")
}

func routeStations(c *fiber.Ctx) error {
	keyword := c.Query("id")

	for _, char := range keyword {
		if char < '0' || char > '9' {
			return shared.ResponseBadRequest(c, "Invalid keyword, must contain only digits")
		}
	}

	key := "route:" + keyword

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	route, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	if route == nil {
		panic("route id wrong")
	}

	search := "@station_route:" + keyword
	index := redisearch.NewClient("localhost:6379", "station_index")
	query := redisearch.NewQuery(search)
	query.SetSortBy("station_seq", true) // Sort by station_seq in descending order

	docs, total, err := index.Search(query)
	if err != nil {
		panic(err)
	}

	fmt.Println(docs, total)
	return shared.Response(c, docs)
}

func routeLines(c *fiber.Ctx) error {
	keyword := c.Query("id")

	for _, char := range keyword {
		if char < '0' || char > '9' {
			return shared.ResponseBadRequest(c, "Invalid keyword, must contain only digits")
		}
	}

	if len(keyword) != 8 {
		return shared.ResponseBadRequest(c, "Invalid keyword")
	}

	key := "route:" + keyword

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	route, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	if route == nil {
		return shared.ResponseBadRequest(c, "wrong ID")
	}

	search := "@line_route:" + keyword
	index := redisearch.NewClient("localhost:6379", "line_index")
	query := redisearch.NewQuery(search)

	query.SetSortBy("line_seq", true) // Sort by line_seq in descending order

	docs, total, err := index.Search(query)

	fmt.Println(total)
	return shared.Response(c, docs)
}

func stationRoute(c *fiber.Ctx) error {
	keyword := c.Query("id")

	for _, char := range keyword {
		if char < '0' || char > '9' {
			return shared.ResponseBadRequest(c, "Invalid keyword, must contain only digits")
		}
	}

	key := "station:" + keyword

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	stationRouteKey, err := rdb.HGet(ctx, key, "station_route").Result()
	if err != nil {
		panic(err)
	}
	if stationRouteKey == "" {
		return shared.ResponseBadRequest(c, "Station ID not found")
	}

	routeKey := "route:" + stationRouteKey
	route, err := rdb.HGetAll(ctx, routeKey).Result()
	if err != nil {
		panic(err)
	}

	return shared.Response(c, route)
}

func searchStation(c *fiber.Ctx) error {

	search := c.Query("search")

	specialChars := `@:!$%()""|`

	for _, char := range specialChars {
		if strings.ContainsRune(search, char) {
			return shared.ResponseBadRequest(c, "Invalid search query, must not contain \"@\", \":\", \"!\", \"$\", \"%\", or \"()\"")
		}
	}

	index := redisearch.NewClient("localhost:6379", "station_index")
	query := redisearch.NewQuery(search).Limit(0, 10)
	//query.SetInFields("@__key:station:*")

	docs, total, err := index.Search(query)

	if err != nil {
		panic(err)
	}

	fmt.Println(total)
	return shared.Response(c, docs)
}

func searchRoute(c *fiber.Ctx) error {
	search := c.Query("search")

	specialChars := `@:!$%()|`
	for _, char := range specialChars {
		if strings.ContainsRune(search, char) {
			return shared.ResponseBadRequest(c, "Invalid search query, must not contain \"@\", \":\", \"!\", \"$\", \"%\", or \"()\"")
		}
	}

	index := redisearch.NewClient("localhost:6379", "route_index")
	query := redisearch.NewQuery(search).Limit(0, 10) // Set a limit for initial search

	docs, total, err := index.Search(query)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total results: %d\n", total)

	// Adjust the limit to get at least 5 results
	if total < 5 {
		query.Limit(0, 5)
		docs, _, err = index.Search(query)
		if err != nil {
			panic(err)
		}
	}

	return shared.Response(c, docs)
}

func routeBuild() {
	routes := make([]models.Route, 0)
	db := database.DB

	tx := db.Model(&models.Route{}).Order("id desc")
	routes = make([]models.Route, 0)

	//err2 := tx.Select("*").Where("id = 108")
	err := tx.Select("*").Find(&routes).Error
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	rdb := database.RDB
	//fmt.Println(routes[0])

	// Set some fields.
	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		for _, route := range routes {
			key := "route:" + shared.IntToString(route.BusRouteID)
			rdb.HSet(ctx, key, "route_name", route.BusRouteName, "route_nomer", route.BusRouteNo, "id", route.ID)
			//fmt.Println(route)
			//fmt.Println(route.BusRouteNo)
		}
		return nil
	}); err != nil {
		panic(err)
	}

	// Create a RediSearch index
	indexDefinition := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("route_name", redisearch.TextFieldOptions{})).
		AddField(redisearch.NewTextFieldOptions("route_nomer", redisearch.TextFieldOptions{}))

	// Create the index
	//index := redisearch.NewClient("route_index", client)
	index := redisearch.NewClient("localhost:6379", "route_index")

	if err := index.CreateIndex(indexDefinition); err != nil {
		panic(err)
	}

	fmt.Println("route index created successfully")
}

func stationBuild() {
	db := database.DB

	tx := db.Model(&models.Station{}).Order("id desc")
	stations := make([]models.Station, 0)

	//err2 := tx.Select("*").Where("id = 108")
	err := tx.Select("*").Find(&stations).Error
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	rdb := database.RDB
	//fmt.Println(routes[0])

	// Set some fields.
	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		for _, station := range stations {
			key := "station:" + shared.IntToString(int(station.ID))
			rdb.HSet(ctx, key, "station_route", station.BusRouteID, "station_name", station.BusStopName, "station_id", station.ID, "station_seq", station.BusStopSeq)
			//fmt.Println(route)
			//fmt.Println(route.BusRouteNo)
		}
		return nil
	}); err != nil {
		panic(err)
	}
	// Create a RediSearch index
	indexDefinition := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("station_route", redisearch.TextFieldOptions{})).
		AddField(redisearch.NewTextFieldOptions("station_name", redisearch.TextFieldOptions{})).
		AddField(redisearch.NewTextFieldOptions("station_id", redisearch.TextFieldOptions{})).
		AddField(redisearch.NewNumericFieldOptions("station_seq", redisearch.NumericFieldOptions{Sortable: true}))

	// Create the index
	index := redisearch.NewClient("localhost:6379", "station_index")

	if err := index.CreateIndex(indexDefinition); err != nil {
		panic(err)
	}

	fmt.Println("station index created successfully")
}

func routeLineBuild() {
	db := database.DB

	tx := db.Model(&models.RouteLine{}).Order("seq desc")
	routeLines := make([]models.RouteLine, 0)

	err := tx.Select("*").Find(&routeLines).Error
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	rdb := database.RDB

	if _, err := rdb.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		for _, routeLine := range routeLines {
			key := "line:" + shared.UintToString(routeLine.ID)
			rdb.HSet(ctx, key, "line_route", routeLine.BusRouteID, "line_seq", routeLine.Seq, "gpx_x", routeLine.GPXx, "gpx_y", routeLine.GPXy)
		}
		return nil
	}); err != nil {
		panic(err)
	}

	indexDefinition := redisearch.NewSchema(redisearch.DefaultOptions).
		AddField(redisearch.NewTextFieldOptions("line_route", redisearch.TextFieldOptions{})).
		AddField(redisearch.NewSortableNumericField("line_seq"))

	index := redisearch.NewClient("localhost:6379", "line_index")

	if err := index.CreateIndex(indexDefinition); err != nil {
		panic(err)
	}

	fmt.Println("line index created successfully")
}

func clear() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: ":6379",
	})
	_ = rdb.FlushDB(ctx).Err()
}

func getRouteFromRedis(key string) (map[string]string, error) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	result, err := rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func testFind() {
	key := "route:108"
	route, err := getRouteFromRedis(key)
	if err != nil {
		panic(err)
	}

	// Print the values
	for field, value := range route {
		fmt.Printf("%s: %s\n", field, value)
	}
}

func stationBuild2() {
	indexName := "stations"

	client := redisearch.NewClient("localhost:6379", indexName)

	// Search for stations by ID
	db := database.DB
	tx := db.Model(&models.Station{}).Order("id desc")
	stations := make([]models.Station, 0)

	err := tx.Select("*").Find(&stations).Error
	if err != nil {
		panic(err)
	}

	// Index stations in Redisearch
	for _, station := range stations {
		doc := redisearch.NewDocument(strconv.Itoa(int(station.ID)), 1.0).
			Set("id", station.ID).
			Set("bus_route_id", station.BusRouteID).
			Set("bus_stop_name", station.BusStopName)

		if err := client.Index(doc); err != nil {
			panic(err)
		}
	}

	// Perform a search query
	query := redisearch.NewQuery("@bus_route_id:108").Limit(0, 10)
	res, totalResults, err := client.Search(query)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Total results: %d\n", totalResults)
	for _, doc := range res {
		fmt.Printf("ID: %s, BusRouteID: %s, BusStopName: %s\n", doc.Id, doc.Properties["bus_route_id"], doc.Properties["bus_stop_name"])
	}
}
