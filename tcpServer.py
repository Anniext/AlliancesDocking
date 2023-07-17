import socket
import time

# 创建TCP套接字
server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

# 绑定服务器的主机名和端口号
server_host = '127.0.0.1'
server_port = 8888
server_socket.bind((server_host, server_port))

# 监听连接
server_socket.listen(5)

print("Server is listening on {}:{}".format(server_host, server_port))

while True:
    # 接受客户端连接
    client_socket, client_address = server_socket.accept()
    print("Client {} connected".format(client_address))

    # 接收数据
    data = client_socket.recv(1024)
    print("Received data:", data.decode())

    if data.decode() == "GET_CHECKET7a60fc3e9880d9ba768fa45f89dae774Sunmnet":
        client_socket.sendall("GET_CHECKETAuthorization success !Sunmnet".encode())
        time.sleep(3)
        while True:
            data = client_socket.recv(1024)
            if data != b'':
                print("Received data:", data.decode())
            client_socket.sendall("KEEPALIVESunmnet".encode())
            time.sleep(3)
    else:
        # 3秒后关闭连接
        time.sleep(3)
        client_socket.close()