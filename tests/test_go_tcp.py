import subprocess
import json
import os
import pytest
import socket
import time

class GoTCPClient:
    def __init__(self, binary_path):
        self.binary_path = binary_path
    
    def calculate_via_subprocess(self, numbers):
        if not os.path.exists(self.binary_path):
            pytest.skip(f"Бинарник не найден: {self.binary_path}")
        
        proc = subprocess.Popen(
            [self.binary_path],
            stdin=subprocess.PIPE,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE
        )
        
        input_data = json.dumps({"numbers": numbers})
        stdout, stderr = proc.communicate(input_data.encode(), timeout=5)
        
        if stderr:
            print("STDERR:", stderr.decode())
        
        result = json.loads(stdout)
        return result["sum"]
    
    def calculate_via_socket(self, numbers, port=9000):
        try:
            sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
            sock.connect(("localhost", port))
            
            data = json.dumps({"numbers": numbers})
            sock.send(data.encode())
            
            response = sock.recv(1024)
            result = json.loads(response)
            return result["sum"]
        finally:
            sock.close()

def test_tcp_subprocess():
    binary = "E:/laba9/go-task5/tcp-server.exe"
    client = GoTCPClient(binary)
    result = client.calculate_via_subprocess([1, 2, 3, 4, 5])
    assert result == 55

def test_tcp_socket():
    binary = "E:/laba9/go-task5/tcp-server.exe"
    client = GoTCPClient(binary)
    
    try:
        result = client.calculate_via_socket([1, 2, 3, 4, 5])
        assert result == 55
    except ConnectionRefusedError:
        pytest.skip("Сервер не запущен")