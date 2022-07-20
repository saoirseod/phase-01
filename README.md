# phase-01
JWT token generation and validation
What is expected when running:
Running the auto.sh script to generate the private key and the token.
 
If the script is prompted to run by a user who is not root, this output is generated.
 
To validate that token, run the main.go file within the validate directory.
It will use regex to compare token signatures to make sure that they match, make sure that the JWT token contains two ‘.’ And follows the expected syntax. 
 
When I altered the token.txt file and changed the token, this is the difference in the output when running the validation:
 
