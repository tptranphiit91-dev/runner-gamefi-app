#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
RED='\033[0;31m'
NC='\033[0m' # No Color

BASE_URL="http://localhost:8080"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  Booking Service API Test Script${NC}"
echo -e "${BLUE}========================================${NC}\n"

# Test 1: Health Check
echo -e "${GREEN}1. Testing Health Check...${NC}"
curl -s -X GET "$BASE_URL/health" | jq '.'
echo -e "\n"

# Test 2: Create User 1
echo -e "${GREEN}2. Creating User 1...${NC}"
USER1=$(curl -s -X POST "$BASE_URL/api/v1/users" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john.doe@example.com",
    "username": "johndoe",
    "password": "password123",
    "full_name": "John Doe",
    "phone": "+1234567890"
  }')
echo "$USER1" | jq '.'
USER1_ID=$(echo "$USER1" | jq -r '.data.id')
echo -e "\n"

# Test 3: Create User 2
echo -e "${GREEN}3. Creating User 2...${NC}"
USER2=$(curl -s -X POST "$BASE_URL/api/v1/users" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "jane.smith@example.com",
    "username": "janesmith",
    "password": "securepass456",
    "full_name": "Jane Smith",
    "phone": "+9876543210"
  }')
echo "$USER2" | jq '.'
USER2_ID=$(echo "$USER2" | jq -r '.data.id')
echo -e "\n"

# Test 4: List All Users
echo -e "${GREEN}4. Listing All Users...${NC}"
curl -s -X GET "$BASE_URL/api/v1/users" | jq '.'
echo -e "\n"

# Test 5: Get User by ID
echo -e "${GREEN}5. Getting User by ID ($USER1_ID)...${NC}"
curl -s -X GET "$BASE_URL/api/v1/users/$USER1_ID" | jq '.'
echo -e "\n"

# Test 6: Update User
echo -e "${GREEN}6. Updating User ($USER1_ID)...${NC}"
curl -s -X PUT "$BASE_URL/api/v1/users/$USER1_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "John Doe Updated",
    "phone": "+1111111111"
  }' | jq '.'
echo -e "\n"

# Test 7: Get Updated User
echo -e "${GREEN}7. Getting Updated User ($USER1_ID)...${NC}"
curl -s -X GET "$BASE_URL/api/v1/users/$USER1_ID" | jq '.'
echo -e "\n"

# Test 8: List Users with Filter
echo -e "${GREEN}8. Listing Users with Filter (limit=1)...${NC}"
curl -s -X GET "$BASE_URL/api/v1/users?limit=1&is_active=true" | jq '.'
echo -e "\n"

# Test 9: Delete User
echo -e "${GREEN}9. Deleting User ($USER2_ID)...${NC}"
curl -s -X DELETE "$BASE_URL/api/v1/users/$USER2_ID" | jq '.'
echo -e "\n"

# Test 10: Verify Deletion
echo -e "${GREEN}10. Verifying Deletion (should return error)...${NC}"
curl -s -X GET "$BASE_URL/api/v1/users/$USER2_ID" | jq '.'
echo -e "\n"

# Test 11: List Remaining Users
echo -e "${GREEN}11. Listing Remaining Users...${NC}"
curl -s -X GET "$BASE_URL/api/v1/users" | jq '.'
echo -e "\n"

echo -e "${BLUE}========================================${NC}"
echo -e "${BLUE}  All Tests Completed!${NC}"
echo -e "${BLUE}========================================${NC}"

