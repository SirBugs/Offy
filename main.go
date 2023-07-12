package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// Constants To Set
	TelegramReporting := true
	DiscordReporting := true
	telegramBotToken := ""
	telegramChatID := ""
	discordWebhookURL := ""
	discordUsername := ""

	// to check if it's following or single command
	info, _ := os.Stdin.Stat()

	// Get the instance ID
	cmdGetInstanceID := "curl -s http://169.254.169.254/latest/meta-data/instance-id"
	outputGetInstanceID, err := exec.Command("bash", "-c", cmdGetInstanceID).Output()
	if err != nil {
		fmt.Println("Error getting instance ID:", err)
		return
	}

	instanceID := strings.TrimSpace(string(outputGetInstanceID))

	// Describe instance status
	cmdDescribeInstanceStatus := fmt.Sprintf("aws ec2 describe-instance-status --instance-ids %s --query 'InstanceStatuses[].InstanceState.Name' | jq -r '.[]'", instanceID)
	outputDescribeInstanceStatus, err := exec.Command("bash", "-c", cmdDescribeInstanceStatus).Output()
	if err != nil {
		fmt.Println("Error describing instance status:", err)
		return
	}

	instanceStatus := strings.TrimSpace(string(outputDescribeInstanceStatus))

	instanceName, err := getInstanceName(instanceID)
	if err != nil {
		fmt.Println("Error getting instance name:", err)
		return
	}

	// # //////////////////////////////////////////////////////////////////

	if len(os.Args) > 1 && os.Args[1] == "stop" {
		stopInstance(instanceID, instanceName)
		panic("Machine Stopped")
	}

	if (info.Mode() & os.ModeNamedPipe) != 0 {
		// followed command :: Stop Instance
		stopInstance(instanceID, instanceName)
	} else {
		// single command :: Notify
		for {
			if instanceStatus == "running" {
				// Send message to Telegram
				message := "Your instance: " + instanceID + " (" + instanceName + ")\nis running Till " + printDateTime() + "\nDid you forget it?"
				telemessage := strings.ReplaceAll(message, "\n", "%0A")
				discmessage := strings.ReplaceAll(message, "\n", " ")
				if TelegramReporting == true {
					sendTelegramMessage(telegramBotToken, telegramChatID, telemessage)
				}

				// Send message to Discord
				// sendDiscordMessage(discordWebhookURL, message)
				if DiscordReporting == true {
					sendDiscordMessage(discordWebhookURL, discordUsername, discmessage)
				}
			}
			time.Sleep(time.Second * 3600)
		}
	}
}

func getInstanceName(instanceID string) (string, error) {
	cmd := fmt.Sprintf("aws ec2 describe-tags --filters \"Name=resource-id,Values=%s\" \"Name=key,Values=Name\" --query 'Tags[].Value' --output text", instanceID)
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}

func stopInstance(instanceID string, instanceName string) {
	cmdStopInstance := fmt.Sprintf("aws ec2 stop-instances --instance-ids %s", instanceID)
	outputStopInstance, err := exec.Command("bash", "-c", cmdStopInstance).Output()
	if err != nil {
		fmt.Println("Error getting instance ID:", err)
		return
	}
	fmt.Println("Mahice " + instanceID + " (" + instanceName + ") Stopped: " + string(outputStopInstance))
}

func sendTelegramMessage(botToken, chatID, message string) {
	cmd := fmt.Sprintf("curl -s -X POST \"https://api.telegram.org/bot%s/sendMessage\" -d chat_id=\"%s\" -d text=\"%s\"", botToken, chatID, message)
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println("Error sending Telegram message:", err)
		return
	}

	fmt.Println("Telegram Notified !!")
}

func sendDiscordMessage(webhookURL, username, message string) {
	cmd := fmt.Sprintf("echo -e '{\"username\": \"%s\", \"content\": \"%s\"}' | curl -X POST -H \"Content-Type: application/json\" -d @- %s", username, message, webhookURL)
	output, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println("Error sending Discord message:", err)
		return
	}

	fmt.Println("Discord Notified !!")
	fmt.Println(string(output))
}

func printDateTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 03:04:05 PM")
	return formattedTime
}

// 1. Installation and configuration
// just install awscli by running: `brew install awscli`, then check If It's installed or not by running the `aws --version`
// Go to: https://us-east-1.console.aws.amazon.com/iamv2/home, Then create a user and SecurityGroup with permission "AmazonEC2FullAccess" for the selected user, Then create your Keys !!
// If you don't know to create the Access Key ID and Secret Key ID in your IAM Panel, Ask ChatGPT or search !! Google is your friend !!
// Now configure awscli by running: `aws configure`, Then submit your Key, Secret, Region, default
//
// 2. Setting Telegram and Discord
// To Configure your Telegram bot, Open @BotFather and create your own bot and set the API token of it, It's like: 6342603457:AAH6Im9kxIdDeXS3J01hKkC1lvjl9RmQoPp
// To Get your Telegram ChatID visit: https://t.me/chat_id_echo_bot and send a "/start" command
// Go search now about how to create a Webhook for your discord text channel and set the webhook URL too, It's like: https://discord.com/api/webhooks/1128499857370200126/GqDm49FpeQ-c4fdhlM5g44TrlfKd9dvyWkAoh_nVyvLc5OFgTr5FerTvHdW8s3kN3Yq
// Set your discord username too, like: `UserName#8344`
//
// 3. Saving aliases
// Save in your ~/.profile :: `alias offy="go run /home/sirbugs/Desktop/Tools/offy/main.go"`
// Save in your ~/.profile :: `alias noffy="nohup go run /home/sirbugs/Desktop/Tools/offy/main.go > output.log 2>&1 &"` as alias called `noffy` in your  ~/.profile, and each time you open your machine just run `noffy` and keep going

// To stop the machine after finishing from using it run: `offy stop`

// Why would you use offy?
// => It can keep remining you every one hour about your instance if it's running or not, just run: `noffy`
// => If you running something taking time for example `nuclei` and you want the machine to stop after this command, just run: `nuclei ...........; offy` or `nuclei ..........; offy stop`
// => You can use it directly to stop the instance after finishing using it via command: `offy stop`
