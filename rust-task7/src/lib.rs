#[derive(Debug, PartialEq)] // Чтобы тесты могли сравнивать объекты
pub struct CalculationResult {
    pub value: f64,
    pub description: String,
}

pub fn calculate_area(radius: f64) -> Result<CalculationResult, String> {
    if radius < 0.0 {
        return Err("Radius cannot be negative".to_string());
    }
    
    let area = std::f64::consts::PI * radius.powi(2);
    
    Ok(CalculationResult {
        value: area,
        description: format!("Area for radius {}", radius),
    })
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_valid_calculation() {
        let res = calculate_area(1.0).unwrap();
        assert!((res.value - std::f64::consts::PI).abs() < 1e-10);
    }

    #[test]
    fn test_negative_radius() {
        assert!(calculate_area(-1.0).is_err());
    }
}