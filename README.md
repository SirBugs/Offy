# Offy
offy is a tool for bugbounty hunters to save money in their EC2 instances

# Introduction & Features
I actually created this tool cuz I noticed that Amazon is getting hella money from the people who forgets their machines opened
So I decided to create this tool, Which would:
1. Keep remembering me that my instance is running
2. Stop my instance after a long command to save money
3. Stop the machine from a command without need to open EC2/AWS website

# Run
- Installation and configuration
    - Go to: https://us-east-1.console.aws.amazon.com/iamv2/home and create a user with a SecurityGroup with permission called `AmazonEC2FullAccess`, Then create your Keys !!
    - ![Alt text](https://raw.githubusercontent.com/SirBugs/Offy/main/imgs/Screen%20Shot%202023-07-12%20at%205.13.19%20AM.png)
    - If you don't know to create the Access Key ID and Secret Key ID in your IAM Panel, Ask ChatGPT or search !! Google is your friend !!
    - run: `brew install awscli` and confirm/check if it's installed by running: `aws --version`
    - a
