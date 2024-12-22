1. `nano ~/.bashrc`

2. Add this line at the end of the file:

`export PATH=$PATH:/home/shubham/code`

3. Save and close the file. (For nano, press Ctrl+O, then Enter, and Ctrl+X.)

4. Reload the shell configuration:

`source ~/.bashrc`

5. Navigate to the directory where your CLI executable 

`cd /home/shubham/code`

6. Use the chmod command to make the file executable:

`chmod +x templ8.exe`

7. Now try running the CLI again:

`templ8.exe --help`