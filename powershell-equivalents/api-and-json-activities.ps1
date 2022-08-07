# Capture stdout of a cli and write it a local text file
kubectl get ns -o json | Out-File ./powershell-equivalents/kubectl.json

# Read a text file from local storage and ingest as an arry
$jsonFile = Get-Content ./powershell-equivalents/kubectl.json | ConvertFrom-Json -Depth 100

# Add a value to an array
$jsonFile.items += @{
    apiVersion = "v1"
    kind = "Namespace"
    metadata = "Something new"
}

# Call an api
$result = Invoke-WebRequest -Uri 'https://webhook.site/35bfe5d5-4530-4665-a974-6ea974ba9499' -Body $($jsonFile | ConvertTo-Json -Depth 100)