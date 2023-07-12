# Offy
offy is a tool for bugbounty hunters to save money in their EC2 instances

# Introduction
I actually created this tool cuz I noticed that Amazon is getting hella money from the people who forgets their machines opened
So I decided to create this tool, Which would:
1. Keep remembering me that my instance is running
2. Stop my instance after a long command to save money
3. Stop the machine from a command without need to open EC2/AWS website

# Run
1. Installation and configuration
    - Go to: https://us-east-1.console.aws.amazon.com/iamv2/home and create a user with a SecurityGroup with permission called `AmazonEC2FullAccess`, Then create your Keys !!
    - ![Alt text](https://raw.githubusercontent.com/SirBugs/Offy/main/imgs/Screen%20Shot%202023-07-12%20at%205.13.19%20AM.png)
    - If you don't know to create the Access Key ID and Secret Key ID in your IAM Panel, Ask ChatGPT or search !! Google is your friend !!
    - ![Alt text](https://raw.githubusercontent.com/SirBugs/Offy/main/imgs/Screen%20Shot%202023-07-12%20at%204.48.24%20AM.png)
    - run: `brew install awscli` and confirm/check if it's installed by running: `aws --version`
    - Now configure awscli by running: `aws configure`, Then submit your generated Key, Secret, Region, default
2. Configuring Telegram and Discord
    - To Configure your Telegram bot, Open @BotFather and create your own bot and set the API token of it, It's like: `6342603457:AAH6Im9kxIdDeXS3J01hKkC1lvjl9RmQoPp`
    - To Get your Telegram ChatID visit: https://t.me/chat_id_echo_bot and send a `/start` message
    - Go search now about how to create a Webhook for your discord text channel and set the webhook URL too, It's like: `https://discord.com/api/webhooks/1128499857370200126/GqDm49FpeQ-c4fdhlM5g44TrlfKd9dvyWkAoh_nVyvLc5OFgTr5FerTvHdW8s3kN3Yq`
    - Set your discord username too, like: `UserName#8344`
    - ![Alt text](https://raw.githubusercontent.com/SirBugs/Offy/main/imgs/Screen%20Shot%202023-07-12%20at%205.37.25%20AM.png)
    - ![Alt text](https://raw.githubusercontent.com/SirBugs/Offy/main/imgs/Screen%20Shot%202023-07-12%20at%205.38.10%20AM.png)
3. Saving aliases
    - Save in your ~/.profile :: `alias offy="go run /home/sirbugs/Desktop/Tools/offy/main.go"`
    - Save in your ~/.profile :: `alias noffy="nohup go run /home/sirbugs/Desktop/Tools/offy/main.go > output.log 2>&1 &"` as alias called `noffy` in your  ~/.profile, and each time you open your machine just run `noffy` and keep going

# Features
- It can keep remining you every one hour about your instance if it's running or not, just run: `noffy`
- If you running something taking time for example `nuclei` and you want the machine to stop after this command, just run: `nuclei ...........; offy` or `nuclei ..........; offy stop`
- You can use it directly to stop the instance after finishing using it via command: `offy stop`

# Updates
- (1.0.0) :: Published

# Credits
This tool was written in Golang 1.19.4, Made with all love in Egypt! <3

Twitter@SirBagoza , Github@SirBugs
