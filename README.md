# phase-01
JWT token generation and validation
What is expected when running:
Running the auto.sh script to generate the private key and the token.

![Picture1](https://user-images.githubusercontent.com/106681090/180015773-535a6dd8-4678-4aa7-ba39-31bf0a7d43d4.png)

If the script is prompted to run by a user who is not root, this output is generated.

 ![Picture2](https://user-images.githubusercontent.com/106681090/180015823-8073eab3-dbff-42e3-91dc-fca63dddf4de.png)

To validate that token, run the main.go file within the validate directory.
It will use regex to compare token signatures to make sure that they match, make sure that the JWT token contains two ‘.’ And follows the expected syntax. 

![Picture3](https://user-images.githubusercontent.com/106681090/180015877-3f84c43f-a86a-4282-a7ef-e620f91def13.png)
 
When I altered the token.txt file and changed the token, this is the difference in the output when running the validation:

 ![Picture4](https://user-images.githubusercontent.com/106681090/180015935-1383f540-8547-4778-b613-5d84ffd5a021.png)

