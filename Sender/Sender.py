import stomp
# stomp.Connection([('127.0.0.1', 62613)])
conn = stomp.Connection10()

conn.start()

conn.connect()

while True:
    lusho = input("Consola Python: ")
    if lusho == "salir":
        break
    conn.send('DonL', lusho)
    print("Mensaje enviado")
conn.disconnect()
