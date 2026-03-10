$baseUrl = "http://localhost:8080/subscriptions"

# 1. Создать подписку
$sub = @{
    service_name = "Netflix"
    price = 500
    user_id = "60601fee-2bf1-4721-ae6f-7636e79a0cba"
    start_date = "03-2026"
} | ConvertTo-Json

$response = Invoke-RestMethod -Uri $baseUrl -Method Post -Body $sub -ContentType "application/json"
$id = $response.id
Write-Host "[1] Created subscription ID: $id"

# 2. Список подписок
$list = Invoke-RestMethod -Uri "$baseUrl/list?user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba&service_name=Netflix"
Write-Host "[2] List subscriptions:"
$list | Format-Table id, service_name, price, user_id, start_date

# 3. Получить по ID
$subById = Invoke-RestMethod -Uri "$baseUrl/get?id=$id"
Write-Host "[3] Get subscription by ID:"
$subById | Format-Table id, service_name, price, user_id, start_date

# 4. Total
$total = Invoke-RestMethod -Uri "$baseUrl/total?user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba&service_name=Netflix&from=03-2026&to=03-2026"
Write-Host "[4] Total price:"
$total.total_price

# 5. Delete
Invoke-RestMethod -Uri "$baseUrl/delete?id=$id" -Method Delete
Write-Host "[5] Deleted subscription $id"

# 6. Список после удаления
$listAfter = Invoke-RestMethod -Uri "$baseUrl/list?user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba&service_name=Netflix"
Write-Host "[6] List after deletion:"
$listAfter | Format-Table id, service_name, price, user_id, start_date