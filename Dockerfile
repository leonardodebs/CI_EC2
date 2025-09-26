# Use a imagem "distroless", que é mínima, mais leve e segura para binários Go.
FROM gcr.io/distroless/static-debian11

EXPOSE 8000

WORKDIR /app

# Copia o executável "programa" (renomeado no workflow) para a imagem.
COPY ./programa .

CMD [ "./programa" ]
