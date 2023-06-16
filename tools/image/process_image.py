import base64
import os
import sys

def encode_image(image_path):
    with open(image_path, 'rb') as image:
        image_data = image.read()
    encoded_data = base64.b64encode(image_data)
    encoded_text = encoded_data.decode('utf-8')
    return encoded_text

def main():
    if len(sys.argv) < 2:
        print("Por favor, forneça o nome de uma imagem como argumento.")
        return

    image_name = sys.argv[1]
    image_path = os.path.join('plates', image_name)

    if not os.path.isfile(image_path):
        print("A imagem especificada não existe.")
        return

    encoded_image = encode_image(image_path)

    text_file_name = os.path.splitext(image_name)[0] + '.txt'
    with open(text_file_name, 'w') as text_file:
        text_file.write(encoded_image)

    print(f"A imagem '{image_name}' foi codificada e salva como '{text_file_name}'.")

if __name__ == '__main__':
    main()
