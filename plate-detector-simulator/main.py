from dotenv import load_dotenv
import os
import paho.mqtt.client as mqtt

load_dotenv()

print("Loading environment variables...")

broker_address = os.getenv("MQTT_BROKER_ADDRESS")
broker_port = int(os.getenv("MQTT_BROKER_PORT"))
username = os.getenv("MQTT_USERNAME")
password = os.getenv("MQTT_PASSWORD")

test_plate = os.getenv("TEST_PLATE")

def on_connect(client, userdata, flags, rc):
  if rc == 0:
    print("Connected successfully!")
    for i in range(1, 3):
      client.subscribe(f"camera/{i}")
  else:
    print(f"Failing connecting to broker. Return code: {rc}")

def on_message(client, userdata, msg):
  if msg.topic.startswith("camera"):
    id = msg.topic.split("/")[1]

    print("New message received from camera " + id)
    print("Publishing result...")
    client.publish(f"camera/{id}/result", test_plate)

print("Initializing MQTT client...")
client = mqtt.Client()

print("Setup callbacks...")
client.on_connect = on_connect
client.on_message = on_message

if username and password:
    client.username_pw_set(username, password)

print("Connecting to broker...")
client.connect(broker_address, broker_port, 60)

print("Starting loop...")
client.loop_forever()
