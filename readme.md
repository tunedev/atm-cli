# ATM CLI
---
A little program simulating an ATM operation via the command line

The flow of the program is as follows:
- run the cli program ```go run .```
- You'd get a prompt to input your 4-digit pin, default is "0000"
- Once you enter the correct pin, you'd get a listing of the available operations
- once you enter the command for the operation, you'd get a prompt to know if you want to perform another operation
- if you respond with y, you'd get a repeat of the second step
- else the program would exit