package dummy

import (
	"fmt"
	"net"
	"route-redis/database"
	"route-redis/models"
)

func start() {
	listener, err := net.Listen("tcp", "192.168.0.121:9090")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Server listening on :9090")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr())

		//go handleClient(conn)
	}
}

func ForDummy() {
	db := database.DB

	route := models.Route{}

	err := db.Where("busroute_id = ?", 14100121).First(&route).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(route.BusRouteName)

	stations := []models.Station{}
	err = db.Where("busroute_id = ?", 14100121).Order("busstop_seq ASC").Find(&stations).Error
	if err != nil {
		panic(err)
	}
	for index, station := range stations {
		fmt.Println("index: ", index)
		fmt.Println(station.BusStopName)
		fmt.Println("seq", station.BusStopSeq)
		fmt.Println(station.GPXx)
	}
}
