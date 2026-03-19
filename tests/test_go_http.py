import subprocess
import json
import os
import pytest

class GoHTTPClient:
    def __init__(self, binary_path):
        self.binary_path = binary_path
    
    def calculate(self, numbers):
        if not os.path.exists(self.binary_path):
            pytest.skip(f"Бинарник не найден: {self.binary_path}")
        

        proc = subprocess.Popen(
            [self.binary_path],  # Без аргумента "server"
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

def test_http_sum_squares():
    binary = "E:/laba9/go-task2/http-server.exe"
    client = GoHTTPClient(binary)
    result = client.calculate([1, 2, 3, 4, 5])
    assert result == 55

def test_http_large_list():
    binary = "E:/laba9/go-task2/http-server.exe"
    client = GoHTTPClient(binary)
    numbers = list(range(1, 100))
    result = client.calculate(numbers)
    expected = sum(x*x for x in numbers)
    assert result == expected