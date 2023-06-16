# Usar uma imagem base do Python
FROM python:3.9

# Copiar o diretório de trabalho atual para /app no contêiner
COPY . /app

# Definir o diretório de trabalho
WORKDIR /app

# Instalar as dependências do script Python
RUN pip install -r requirements.txt

# Definir o comando padrão para ser executado quando o contêiner for iniciado
CMD ["python", "main.py"]
