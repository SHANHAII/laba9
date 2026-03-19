import fastmath
import pytest

def test_sum_squares():
    result = fastmath.sum_squares([1, 2, 3, 4, 5])
    assert result == 55

def test_mean():
    result = fastmath.mean([1.0, 2.0, 3.0, 4.0, 5.0])
    assert result == 3.0

def test_stats():
    stats = fastmath.stats([1.0, 2.0, 3.0, 4.0, 5.0])
    assert stats["min"] == 1.0
    assert stats["max"] == 5.0
    assert stats["sum"] == 15.0
    assert stats["mean"] == 3.0

def test_empty_stats():
    with pytest.raises(Exception):
        fastmath.stats([])

def test_matrix():
    matrix = fastmath.Matrix([[1, 2], [3, 4]])
    assert matrix.get(0, 0) == 1
    assert matrix.get(1, 1) == 4
    
    transposed = matrix.transpose()
    assert transposed.get(0, 1) == 3
    
    assert repr(matrix) == "Matrix(2x2)"