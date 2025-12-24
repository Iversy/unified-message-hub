# Best Chat
curl -X POST http://localhost:8080/api/send/message \
  -H "Content-Type: application/json" \
  -d '{
    "source": 2,
    "sender": "Mannequin",
    "chat_id": 666,
    "text": "Hi my very precious friend! I have an offer. Please reply.\nMESSAGE FOR TESTING PURPOSE ONLY",
    "timestamp": "2024-01-15T10:20:00Z"
  }'
#  Horrid chat
curl -X POST http://localhost:8080/api/send/message \
  -H "Content-Type: application/json" \
  -d '{
    "source": 2,
    "sender": "goose",
    "chat_id": 111,
    "text": "Hello Dear Friend! We wasnt together for long but i hope you still remember me...",
    "timestamp": "2024-01-15T10:20:00Z"
  }'
# Rules
curl -X POST http://localhost:8080/api/send/route \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "name": "Best chat",
    "source_chat_id": 666,
    "receiver_id": 205630058,
    "keywords": [],
    "is_active": true
  }'
