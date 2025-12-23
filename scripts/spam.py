import requests
import json
from datetime import datetime, timedelta

# Define the URL and headers
url = "http://localhost:8080/api/send"
headers = {
    "Content-Type": "application/json"
}

# Function to generate and send POST requests
def send_requests(num_requests):
    for i in range(num_requests):
        # Update the data as needed for each request
        data = {
            "source": 2,
            "sender": f"a{i}",  # Change sender for each request
            "chat_id": 456 + i,  # Increment chat_id for each request
            "text": f"Hello World! Message number {i + 1}",  # Custom message
            "timestamp": (datetime.utcnow() + timedelta(seconds=i)).isoformat() + 'Z'  # Increment timestamp
        }

        # Send the POST request
        response = requests.post(url, headers=headers, data=json.dumps(data))
        
        # Print response status for tracking
        print(f"Request {i + 1}: Status Code - {response.status_code}, Response - {response.text}")

# Call the function to send 1000 requests
send_requests(10)
