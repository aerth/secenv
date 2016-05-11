# Secure Environmental Variable Storage

Uses the [seconf](https://github.com/aerth/seconf) library.

Stores environmental variables in a nacl/secretbox (XSalsa20 and Poly1305) for later.

Note: While running your app using secenv yourapp.sh, the variables are going to be unencrypted (plaintext) somewhere in /proc/ directory. Just like typing `KEY=123 yourapp.sh` 

This would be possible to view if one is owner of the machine, while it is running. 

If you encounter any issues, report them on github.com/aerth/secenv/issues/new

Pull requests are welcome.

## Works with:

AWS Keys, Mandrill Keys, Sendgrid Keys, Heroku Keys, and on and on...

## Usage

Delete config file

```
secenv -d
```
Create a config file

Here you will be prompted four times. Here is an example configuration that will have a (temporarily) noticable effect on your terminal session.

```
secenv
```

Test it out (output to stdout)

```
secenv env
secenv 'echo $TEST $TEST2'
```

Run something with the variables

```
secenv '
```


Example Session:

```

bash-4.3$ ./secenv
Welcome to secenv. No config file found at ~/.secenv, would you like to create one?
yes
Enter the first environmental NAME: test1
Enter the first environmental VALUE: test1works
Enter the second environmental NAME: test2
Enter the second environmental VALUE: test2works
Create a password to encrypt config file:
Press ENTER for no password.
Config file saved at /home/aerth/.secenv 
Total size is 82 bytes.

bash-4.3$ ./secenv 'echo $test1 $test2'
test1works test2works
 

```