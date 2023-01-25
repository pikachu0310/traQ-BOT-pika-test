package commands

import (
	"example-bot/api"
	"fmt"
	"os/exec"
)

func Docker(ChannelID string, slice []string) {
	if len(slice) == 1 {
		return
	}
	if slice[1] == "compose" {
		slice = append(slice[:1], slice[2:]...)
	}
	if slice[1] == "restart" {
		msg, _ := api.PostMessageWithErr(ChannelID, "docker compose restart :loading:")
		out, err := exec.Command("docker", "compose", "restart").CombinedOutput()
		if err != nil {
			api.EditMessage(msg.Id, fmt.Sprintf("docker compose restart :caution:\n```\n%s```", out))
		} else {
			api.EditMessage(msg.Id, "docker compose restart :done:")
		}
	} else if slice[1] == "up" {
		msg, _ := api.PostMessageWithErr(ChannelID, "docker compose up :loading:")
		out, err := exec.Command("docker", "compose", "up").CombinedOutput()
		if err != nil {
			api.EditMessage(msg.Id, fmt.Sprintf("docker compose up :caution:\n```\n%s```", out))
		} else {
			api.EditMessage(msg.Id, "docker compose up :done:")
		}
	} else if slice[1] == "down" {
		msg, _ := api.PostMessageWithErr(ChannelID, "docker compose down :loading:")
		out, err := exec.Command("docker", "compose", "down").CombinedOutput()
		if err != nil {
			api.EditMessage(msg.Id, fmt.Sprintf("docker compose down :caution:\n```\n%s```", out))
		} else {
			api.EditMessage(msg.Id, "docker compose down :done:")
		}
	}
}
