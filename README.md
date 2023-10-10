# git-profile

Switches between different git profiles.

Change Github private keys on demand, write a cli tool to do it
1. ~Check if the command is running in sudo mode, if not return 1~
2. List the available profiles.
3. Lets user navigate the profile and select the appropriate one.
4. Should be as if the user has made the actual commits, commits should be signed.
5. Easiest way to do this is using SSH keys, create a cli tool to generate ssh keys (private and public) and then store the private key locally, copy the public key to clipboard.
6. User then goes to GitHub and pastes the public key.
7. Now the user should be authenticated based on the public-private key pair, all the commits should be signed before pushing to remote.

Update:
manually worked on adding ssh keys to github.com for authenticating and signing keys

Update:
To generate SSH keys -->[link](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent#generating-a-new-ssh-key)\
To add SSH keys to github account --> [link](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)\
For signing commits locally --> [link](https://docs.github.com/en/authentication/managing-commit-signature-verification/signing-commits)
