# Introduction  
**Goge Basis** is a foundational package for **Goge Engine**.  
This package provides basic structures and components.  

## Documentation  

`type Warning interface` A warning interface for returning warning messages, with the following method:  
- `Warning() string` Returns the warning message

`type Warning struct` A warning structure for recording warning messages, implementing the `type Warning interface`, with the following field:  
- `message string` The warning message

`func NewWarning(warningMessage string) Warning` Returns a `Warning` struct with the given warning message  

`func ReturnWarningForTest() Warning` Returns a `Warning` struct for testing purposes  