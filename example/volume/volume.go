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
	volumes, err := client.ListVolumes(ctx)
	for _, volume := range volumes {
		log.Printf("Volume: %s", volume.Name)
		log.Printf("Region: %s", volume.Region)
	}
	volume, err := client.CreateVolume(ctx, ovh.VolumeCreateOptions{
		Name:   "test",
		Region: region,
		Type:   ovh.VolumeClassic,
		Size:   10,
	})
	if err != nil {
		log.Fatal(err)
	}
	volumes, err = client.ListVolumes(ctx)
	for _, volume := range volumes {
		log.Printf("Volume: %s", volume.Name)
		log.Printf("Region: %s", volume.Region)
	}
	err = client.DeleteVolume(ctx, volume.ID)
	if err != nil {
		log.Fatal(err)
	}

}
