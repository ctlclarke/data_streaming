package main

import (
    "fmt"
    "os"
    "os/signal"
    "math/rand"
    "github.com/yosssi/gmq/mqtt"
    "github.com/yosssi/gmq/mqtt/client"
)

func main() {
    // Set up channel on which to send signal notifications.
    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc, os.Interrupt, os.Kill)

    // Create an MQTT Client.
    cli := client.New(&client.Options{
        // Define the processing of the error handler.
        ErrorHandler: func(err error) {
            fmt.Println(err)
        },
    })

    // Terminate the Client.
    defer cli.Terminate()

    // Connect to the MQTT Server.
    err := cli.Connect(&client.ConnectOptions{
        Network:  "tcp",
        Address:  "172.28.42.32:15675",
        ClientID: []byte("dylon"),
    })
    if err != nil {
        panic(err)
    }

    for i := 0; i < 100; i++ {
        go func(x int, y int) {
             // Publish a message.
            err = cli.Publish(&client.PublishOptions{
                QoS:       mqtt.QoS0,
                TopicName: []byte("bar/baz"),
                Message:   []byte("testMessage"),
            })
            if err != nil {
                panic(err)
            }
            fmt.Println("Sent")
        }(rand.Intn(100), rand.Intn(100))
    }

    // Disconnect the Network Connection.
    if err := cli.Disconnect(); err != nil {
        panic(err)
    }
}