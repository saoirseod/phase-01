#Generate the key pair
#ed25519 is more secure than RSA etc.

#The current user's ID is stored in the $EUID variable.
#When it is equal to zero, the user is root
#User can only run  the script if they are root or they use sudo
if [ "$EUID" -ne 0 ]
        then echo 'ERROR: You must be root to run this script. Use sudo su c>
        else (
echo 'Generating key pair..'
openssl genpkey -algorithm ED25519 -outform pem -out test.ed
openssl pkey -in test.ed -pubout > test.ed.pub
#Issue the token
#add private key in ed format
echo 'Issuing the token..'
echo '-----TOKEN-----'
go run main.go > token.txt test.ed
cat token.txt > validate/token2.txt
cat token.txt
cat test.ed
#copy that output so that it can be validated at the other end
#similar to issuer but we have a public key in our validator instead
#parse the public key

)
fi

exit

