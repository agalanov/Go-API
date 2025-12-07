# API Authentication Test Script

Write-Host "=== API Authentication Test ===" -ForegroundColor Cyan

# Step 1: Login and get token
Write-Host "`n1. Logging in..." -ForegroundColor Yellow
$body = @{
    username = "admin"
    password = "admin123"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri 'http://localhost:8080/api/v1/oauth/login' -Method POST -ContentType 'application/json' -Body $body
    
    # Get token
    $token = $response.access_token
    Write-Host "Success! Token received" -ForegroundColor Green
    Write-Host "Access Token: $token`n" -ForegroundColor Gray
    
    # Step 2: Call /zone with token in header
    Write-Host "2. Calling /zone with token in header..." -ForegroundColor Yellow
    $headers = @{
        'Authorization' = "Bearer $token"
    }
    $zoneResponse = Invoke-RestMethod -Uri 'http://localhost:8080/api/v1/zone' -Headers $headers
    
    Write-Host "Success! Request completed" -ForegroundColor Green
    Write-Host "`nResponse from /zone:" -ForegroundColor Cyan
    $zoneResponse | ConvertTo-Json -Depth 5
    
    # Step 3: Call /zone with token in query parameter
    Write-Host "`n3. Calling /zone with token in query parameter..." -ForegroundColor Yellow
    $zoneResponse2 = Invoke-RestMethod -Uri "http://localhost:8080/api/v1/zone?access_token=$token"
    
    Write-Host "Success! Request completed" -ForegroundColor Green
    Write-Host "`nResponse from /zone (query param):" -ForegroundColor Cyan
    $zoneResponse2 | ConvertTo-Json -Depth 5
    
}
catch {
    Write-Host "Error: $($_.Exception.Message)" -ForegroundColor Red
    if ($_.ErrorDetails.Message) {
        Write-Host "Details: $($_.ErrorDetails.Message)" -ForegroundColor Red
    }
}

Write-Host "`n=== Test completed ===" -ForegroundColor Cyan
