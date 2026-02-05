<!--markdownlint-disable-->


# TCP

Transmission Control Protocol (TCP) is a primary communication protocol of the internet, though that is changing with HTTP3 which is not build on TCP.


Data is sent over the network in packets. when the other side receives packets, they arrive potentially out of order and need to be reassembled. WIthout TCP, we cannot gurantee the order of packets is the same


# TCP VS UDP

A protoco l is a design or spec for one computer/device to communicates with another


TCP is Reliable, in order packets

TCP has a sliding window which specifies how many packets can be out in flight at a time

When a packet gets received, receiver has to send a ACK

When sender receives the ACK, we can slide the window over. Meanwhile we can continue sending packets within our window


# UDP


UDP is considered more performant because it does not wait for an ACK. In UDP we need to define how we want to break up the data, how the receiver knows it didnt receive some data.



# Files vs network

Files arnd networks are very similar. they are both just streams of bytes we can read and write to.

When we read from a file, we decide:

- how much to read
- when to read
- when to stop reading


The only difference is when we read from a network connection, the data is pushed to you by the remote server, we just have to be ready to receive it when it comes


# TCP to HTTP


TCP gurantees data is inorder and complete, but it dosent tell you what it is

HTTP gives us the ability to specify what we are sending


HTTP provides headers / field lines which provide more information about the request
Without HTTP, tcp would be useless because we dont know what is coming thru

A http message:

first line: declaration of what request it is, then destination and then the semantic version of http

GET /cats http1.1 \r\n \r\n which is CRLF and is just http's new line character

second line and beyond are field lines

field-name: value \r\n


http1 knows its done when it sees a empty line with a \r\n


TCP, and by extension HTTP is a streaming protocol meaning we receive data in chunks and should be able to parse it as it ccomes > [!IMPORTANT]
> 
Eg: GE

THEN the final T comes in 

We should be able to know this is a GET


