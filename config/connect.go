/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package config

import (
	"context"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/getsentry/sentry-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
	"log"
	"net/url"
	"path/filepath"
	"time"
)

const (
	configPath =
	//"/home/pato/go/src/Mbunge-Core/config/config.ini"
	"config/config.ini"
	//"/var/www/go/Mbunge-Core/config/config.ini" //

)

var dbUrl, sentryKey, mqtturl, id, username, mqttpassword, mqttport string
var err error

func GetKey(section string, key string) (string, error) {

	abs, err := filepath.Abs(configPath)
	cfg, err := ini.Load(abs)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		log.Panic(err)
		return "", err
		//os.Exit(1)
	}
	return cfg.Section(section).Key(key).String(), nil
}

func init() {

	// Mongodb
	dbUrl, err = GetKey("mongodb", "url")
	if err != nil {
		panic(err)
	}

	//Sentry
	sentryKey, err = GetKey("sentry", "key")
	if err != nil {
		panic(err)
	}

	// MQTT
	mqtturl, err = GetKey("MQTT", "address")
	if err != nil {
		panic(err)
	}
	id, err = GetKey("MQTT", "id")
	if err != nil {
		panic(err)
	}
	username, err = GetKey("MQTT", "username")
	if err != nil {
		panic(err)
	}
	mqttpassword, err = GetKey("MQTT", "password")
	if err != nil {
		panic(err)
	}
	mqttport, err = GetKey("MQTT", "port")
	if err != nil {
		panic(err)
	}
}

// ErrorReporter connection
func ErrorReporter(report string) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryKey,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage(report)

}

// ConnectDB ..
func ConnectDB() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")
	return client
}

// Connect func connecting to the MQTT
func ConnectMqtt() mqtt.Client {
	mqttUrl := fmt.Sprintf("http://%s:%s", mqtturl, mqttport)
	mqttURI, err := url.Parse(mqttUrl)
	if err != nil {
		panic(err)
	}
	opts := createClientOptions(id, mqttURI)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(10 * time.Second) {
	}
	if err := token.Error(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Token: %s\n", token)
	return client
}
func createClientOptions(clientID string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	//opts.SetUsername(username)
	//opts.SetPassword(mqttpassword)
	opts.SetClientID(clientID)
	return opts
}
