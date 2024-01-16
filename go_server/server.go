package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/herizsh/grpc-forecast/forecast/pb2"

	"google.golang.org/grpc"
)

type weatherServer struct {
	pb.UnimplementedWeatherServiceServer
}

type WeatherData struct {
	Summary            string
	TemperatureCelsius float32
}

func generateWeatherForecast(city string, day int) WeatherData {
	summaries := []string{"Sunny", "Cloudy", "Rainy", "Windy", "Snowy"}
	return WeatherData{
		Summary:            summaries[rand.Intn(len(summaries))],
		TemperatureCelsius: 15 + rand.Float32()*10,
	}
}

func (s *weatherServer) StreamForecasts(req *pb.ForecastRequest, stream pb.WeatherService_StreamForecastsServer) error {
	if req.Days < 1 {
		return fmt.Errorf("request must include at least one day")
	}

	log.Printf("Received request for %d day(s) weather forecast in %s", req.Days, req.City)
	for i := 0; i < int(req.Days); i++ {
		forecastData := generateWeatherForecast(req.City, i+1)
		response := &pb.ForecastResponse{
			Date:               fmt.Sprintf("2024-01-%02d", i+1),
			Summary:            forecastData.Summary,
			TemperatureCelsius: forecastData.TemperatureCelsius,
		}

		if err := stream.Send(response); err != nil {
			log.Printf("Error sending data to stream: %v", err)
			return err
		}
		time.Sleep(1 * time.Second) // Simulating delay for each day's forecast.
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterWeatherServiceServer(grpcServer, &weatherServer{})

	log.Printf("Starting Weather Forecast server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 50051: %v", err)
	}
}
