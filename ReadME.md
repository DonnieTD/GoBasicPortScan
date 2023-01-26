How does a normal tcp handshake work?

OPEN PORT:
1 ---> SYN ---> 2
1 <--- SYN-ACK <--- 2 
1 ---> ACK ---> 2

If the port is happens a three-way handshake takes place ( the process above )

First the client sends the SYN Packet which signals the beginning of communications
The server then respons with a syn-ack packet to signal aknowledgment of the syn packet
This triggers the client to finish with an ack packet and transfer can now occur

CLOSED PORT:
1 ---> SYN ---> 2
1 <--- RST <--- 2 

If the port is closed the server will respond with a rst packet 

FILTERED PORT:
1 ---> SYN ---> /// FIREWALL /// 2

If the traffic is being filtered by a firewall the syn packet will never reach the server
theirfor no response will be sent back# GoBasicPortScan
