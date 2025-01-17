**Zomato Weather Union App**

A simple Go application that interacts with an external weather API to fetch weather data based on latitude and longitude inputs.

Features
	•	Fetch weather data for a given location using latitude and longitude.
	•	Handles invalid or missing input parameters gracefully.
	•	Demonstrates the use of http package, external API integration, and JSON handling in Go.

Requirements
	•	Go 1.19 or higher
	•	An active internet connection
	•	API Key for the weather API (X-Zomato-Api-Key)

Setup Instructions
	1.	Clone the repository:

git clone https://github.com/your-username/zomato-weather-app.git
cd zomato-weather-app


	2.	Update the API Key:
Open main.go and replace the placeholder apiKey value with your actual API key:

const apiKey = "your-api-key-here"


	3.	Run the application:

go run main.go


	4.	Access the server:
The application will run on port 8080. Test the /zomato endpoint with Postman or any browser:

http://localhost:8080/zomato?latitude=<latitude>&longitude=<longitude>

Example Request

URL:

http://localhost:8080/zomato?latitude=28.704060&longitude=77.102493

Response:

{
    "status": "200",
    "message": "temporarily unavailable",
    "locality_weather_data": {
        "temperature": null,
        "humidity": null,
        "wind_speed": null,
        "wind_direction": null,
        "rain_intensity": null,
        "rain_accumulation": null,
        "aqi_pm_10": null,
        "aqi_pm_2_point_5": null
    }
}

Project Structure

zomato-weather-app/
├── main.go        # Main application code
├── README.md      # Project documentation

Future Enhancements
	•	Add support for other APIs to fetch detailed weather data.
	•	Implement caching to reduce repeated API calls for the same location.
	•	Enhance error handling for various failure scenarios.

Contributing

Contributions are welcome! Please fork the repository, create a feature branch, and submit a pull request.

License

This project is licensed under the MIT License.
