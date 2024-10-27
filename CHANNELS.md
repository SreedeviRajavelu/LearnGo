Output of running go run channel.go

<img width="755" alt="image" src="https://github.com/user-attachments/assets/7f2d58ff-666d-4ec1-af08-bca3730fa44c">


Explanation:
channel.go code, the write Goroutine has a for loop which writes numbers from 0 to 3 to the ch channel. The capacity of this buffered channel is 2 and hence the write Goroutine will be able to write values 0 and 1 to the ch channel immediately and then it blocks until at least one value is read from ch channel as defined below :

<img width="317" alt="image" src="https://github.com/user-attachments/assets/2818c4b5-d5d9-4190-ac6d-f7a31afed18f">

After that, the read value and then sleeps for 2 seconds again and this cycle continues until the ch is closed. So the program will print the following lines after 2 seconds, as :

<img width="224" alt="image" src="https://github.com/user-attachments/assets/de4bf48a-08a5-4325-88cc-2f297e18e997">

And this will continue until all values are written to the channel and it is closed in the write Goroutine.
