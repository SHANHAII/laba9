import requests
import pytest

BASE_URL = "http://localhost:8080"

def test_health_check():
    response = requests.get(f"{BASE_URL}/health")
    assert response.status_code == 200
    assert response.json()["status"] == "alive"

def test_process_endpoint():

    payload = {"data": "hello world"}
    response = requests.post(f"{BASE_URL}/process", json=payload)
    
    assert response.status_code == 200
    assert "Processed: hello world" in response.json()["result"]
    assert response.json()["length"] == 11

def test_invalid_payload():
   
    response = requests.post(f"{BASE_URL}/process", json={})
    assert response.status_code == 400