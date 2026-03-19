use pyo3::prelude::*;
use std::collections::HashMap;

#[pyfunction]
fn sum_squares(numbers: Vec<i64>) -> i64 {
    numbers.iter().map(|&x| x * x).sum()
}

#[pyfunction]
fn mean(numbers: Vec<f64>) -> f64 {
    if numbers.is_empty() {
        return 0.0;
    }
    let sum: f64 = numbers.iter().sum();
    sum / numbers.len() as f64
}

#[pyfunction]
fn stats(numbers: Vec<f64>) -> PyResult<HashMap<String, f64>> {
    if numbers.is_empty() {
        return Err(pyo3::exceptions::PyValueError::new_err("empty list"));
    }
    
    let mut result = HashMap::new();
    let min = numbers.iter().fold(f64::INFINITY, |a, &b| a.min(b));
    let max = numbers.iter().fold(f64::NEG_INFINITY, |a, &b| a.max(b));
    let sum: f64 = numbers.iter().sum();
    let mean = sum / numbers.len() as f64;
    
    result.insert("min".to_string(), min);
    result.insert("max".to_string(), max);
    result.insert("sum".to_string(), sum);
    result.insert("mean".to_string(), mean);
    
    Ok(result)
}

#[pyclass]
struct Matrix {
    data: Vec<Vec<i32>>,
    rows: usize,
    cols: usize,
}

#[pymethods]
impl Matrix {
    #[new]
    fn new(data: Vec<Vec<i32>>) -> PyResult<Self> {
        if data.is_empty() {
            return Err(pyo3::exceptions::PyValueError::new_err("empty matrix"));
        }
        
        let rows = data.len();
        let cols = data[0].len();
        
        for row in &data {
            if row.len() != cols {
                return Err(pyo3::exceptions::PyValueError::new_err("invalid row length"));
            }
        }
        
        Ok(Matrix { data, rows, cols })
    }
    
    fn get(&self, row: usize, col: usize) -> PyResult<i32> {
        if row >= self.rows || col >= self.cols {
            return Err(pyo3::exceptions::PyIndexError::new_err("index out of bounds"));
        }
        Ok(self.data[row][col])
    }
    
    fn transpose(&self) -> PyResult<Matrix> {
        let mut new_data = vec![vec![0; self.rows]; self.cols];
        
        for i in 0..self.rows {
            for j in 0..self.cols {
                new_data[j][i] = self.data[i][j];
            }
        }
        
        Ok(Matrix {
            data: new_data,
            rows: self.cols,
            cols: self.rows,
        })
    }
    
    fn __repr__(&self) -> String {
        format!("Matrix({}x{})", self.rows, self.cols)
    }
}

#[pymodule]
fn fastmath(_py: Python, m: &PyModule) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(sum_squares, m)?)?;
    m.add_function(wrap_pyfunction!(mean, m)?)?;
    m.add_function(wrap_pyfunction!(stats, m)?)?;
    m.add_class::<Matrix>()?;
    Ok(())
}