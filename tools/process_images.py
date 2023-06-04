import os
import base64
from PIL import Image
import json

directory = 'plates'

json_list = []

for filename in os.listdir(directory):
  if filename.endswith(".jpg") or filename.endswith(".png"):
    plate_code = os.path.splitext(filename)[0]

    image_path = os.path.join(directory, filename)

    with Image.open(image_path) as image:
      image_data = base64.b64encode(image.tobytes()).decode('utf-8')

      json_obj = {
        "plate": plate_code,
        "image_coded": image_data
      }

      json_list.append(json_obj)

with open('plates_coded.json', 'w') as json_file:
  json.dump(json_list, json_file)
