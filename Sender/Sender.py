import stomp

conn = stomp.Connection10()

conn.start()

conn.connect()

while True:
    lusho = input("Consola Python: ")
    if lusho == "salir":
        break
    conn.send('SampleQueue', lusho)
    print("Mensaje enviado")
conn.disconnect()
