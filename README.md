# git-profile
Switches between different git profiles.

Change Github private keys on demand, write a cli tool to do it 
1. List the available profiles.
2. Lets user navigate the profile and select the appropriate one. 
3. Should be as if the user has made the actual commits, commits should be signed.
4. Easiest way to do this is using SSH keys, create a cli tool to generate ssh keys (private and public) and then store the private key locally, copy the public key to clipboard. 
5. User then goes to GitHub and pastes the public key.
6. Now the user should be authenticated based on the public-private key pair, all the commits should be signed before pushing to remote.
