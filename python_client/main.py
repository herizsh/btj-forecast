import grpc

import weather_pb2
import weather_pb2_grpc


def run():
    # Set up a connection to the server.
    channel = grpc.insecure_channel("localhost:50051")
    stub = weather_pb2_grpc.WeatherServiceStub(channel)

    # Get user input for city and number of forecast days.
    city = input("Enter city: ")
    days = int(input("Enter number of forecast days: "))
    if days < 1:
        print("Invalid number of days. Please enter a positive integer.")
        return

    # Create a request.
    request = weather_pb2.ForecastRequest(city=city, days=days)

    # Stream forecasts from the server.
    try:
        for forecast in stub.StreamForecasts(request):
            print(
                f"Date: {forecast.date}, Weather: {forecast.summary}, Temperature: {forecast.temperatureCelsius:.2f}°C"
            )
    except grpc.RpcError as e:
        print(f"RPC error: {e}")


run()
