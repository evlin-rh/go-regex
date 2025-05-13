# go-regex



Initialise a Go Module for the project
    - go.mod file is created
    - go.mod file is used to manage dependencies and module path

Role of go.mod
    - tells Go that our project is a module called "newPackage"
    - Important for importting any files to your project

    i.e
    project/
    - go.mod
    - main.go
    - regex/
        -advancedRegex.go

# Importing new package 
    `import "newPackage/regex"

# Calling exported functions
    ` regex.ExportedFunctions `

