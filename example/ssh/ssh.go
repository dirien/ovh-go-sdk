package main

import (
	"context"
	ovh "github.com/dirien/ovh-go-sdk/pkg/sdk"
	"log"
	"os"
)

func main() {

	ctx := context.Background()

	region := os.Getenv("OVH_REGION")
	serviceName := os.Getenv("OVH_SERVICENAME")
	client, err := ovh.NewOVHDefaultClient(region, serviceName)
	if err != nil {
		log.Fatal(err)
	}
	sshKeys, err := client.ListSSHKeys(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, sshKey := range sshKeys {
		log.Printf("SSH Key: %s", sshKey.Name)
	}

	log.Println("Creating SSH Key")
	key, err := client.CreateSSHKey(ctx, ovh.SSHKeyCreateOptions{
		Name:      "test",
		PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDp7SDil2G9tTt0WqVWcW0YepoFwbg/aK/+Ap0Es0OTawVi9BKjmNRVodUZc4ZIAhtiq89ynLi2bwt3/9Y1xLj+flO7t/A8EIm84oYw3mQw96PR6IYitD0nVEa4zNhwbtAqE0gEBMWstcjzlB0mxeldCzlcWTLfzghGYAlA4uGPzJUYPh21yUWsnJEOLpt8OpgbSFDJNUGH4GPPxTHsuIbebM/lr2WCEFytn66rs4/UNKJXXUsYNrMKtOfxQBKO9wj7ArH034GjYjo2/P2LyAJc4vReYgUIAQjltFjR91sT4nCrEYLPKAwgCIBQsOhNAz1q6GEqzkVV+FhyzrA2kb7Yx/dlS1FIkwBLnhVuDEQXRFUk29K4c7zTNYQKP7s+Fthv7x7QteShE1gV9g9HybZJHuoDKziAqW9QOqLYIFuWofsxnYXccQ3fAPpbilP+14Sy5We0EVVnLCWFGpwQGVRR82yx/S61+3He0x3BsdNuUj4VBSfxYVeljfj0Q60t/6s= dirien@SIT-SMBP1766",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listing SSH Keys")
	sshKeys, err = client.ListSSHKeys(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, sshKey := range sshKeys {
		log.Printf("SSH Key: %s", sshKey.Name)
	}

	log.Println("Deleting SSH Key")
	err = client.DeleteSSHKey(ctx, key.ID)
	if err != nil {
		log.Fatal(err)
	}
}
