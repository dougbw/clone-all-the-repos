Param(
    $Module = "clone-all-the-repos"
)

Clear-Host

$ErrorActionPreference = "Stop"
$ExeName = "{0}.exe" -f $Module

if (Test-Path $ExeName) {
    Remove-Item $ExeName -Force 
}

"building..."
"---------------------------------------"
go build cmd/clone-all-the-repos.go
"---------------------------------------"
if ($LASTEXITCODE -ne 0) {
    Write-Error "Build failed"
}

# .\clone-all-the-repos.exe example-configs\github.yaml
