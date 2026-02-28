package core

var FilePatterns = map[string][]string{
	"go": {".go", ".mod", ".sum"},
	"py": {".py", ".pyw", ".ipynb"},
	"js": {".js", ".jsx", ".mjs", ".cjs", ".json"},
	"ts": {".ts", ".tsx", ".d.ts"},
}

var IgnoreDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"vendor":       true,
	"__pycache__":  true,
	"dist":         true,
	".idea":        true,
	".vscode":      true,
}
