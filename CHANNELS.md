Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

ch <- v    // Send v to channel ch.

v := <-ch  // Receive from ch and assign value to v.
           
(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

ch := make(chan int)

In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free. Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine. By default channel is bidirectional, means the goroutines can send or receive data through the same channel as shown in the below image:

![image](https://github.com/user-attachments/assets/d93e13ed-edfa-41b2-9131-8dc043b7e1ee)


By default channels are unbuffered, which states that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) which are ready to receive the sent value. 

Buffered channels allow to accept a limited number of values without a corresponding receiver for those values. It is possible to create a channel with a buffer. 

Buffered channel are blocked only when the buffer is full. 

Similarly receiving from a buffered channel are blocked only when the buffer will be empty.


Syntax :

ch := make(chan type, capacity)           // chan defines channel type


Output of running go run channel.go

<img width="755" alt="image" src="https://github.com/user-attachments/assets/7f2d58ff-666d-4ec1-af08-bca3730fa44c">


Explanation:
channel.go code, the write Goroutine has a for loop which writes numbers from 0 to 3 to the ch channel. The capacity of this buffered channel is 2 and hence the write Goroutine will be able to write values 0 and 1 to the ch channel immediately and then it blocks until at least one value is read from ch channel as defined below :

<img width="317" alt="image" src="https://github.com/user-attachments/assets/2818c4b5-d5d9-4190-ac6d-f7a31afed18f">

After that, the read value and then sleeps for 2 seconds again and this cycle continues until the ch is closed. So the program will print the following lines after 2 seconds, as :

<img width="224" alt="image" src="https://github.com/user-attachments/assets/de4bf48a-08a5-4325-88cc-2f297e18e997">

And this will continue until all values are written to the channel and it is closed in the write Goroutine.
