What is the Booking App?
    - The Booking App is a GO written program which invites a user to register for a convention. By enrolling with their name, 
email, and the number of tickets they would like to purchase. This then prints a receipt for the user's confirmation.

Why is this project useful?
    - This project was built on the idea of creating a basic yet efficient form of creating bookings for a convention.
The Booking App is a great project for developers who are looking to begin their Golang journey. This application 
introduces the user to the main concepts of golang but later dives into the topics and material to which make GO 
such a strong and effiencent programming langauge.

How to build the project?
    - There are many routes one can take when building a similar project. However in this project variables 
and constants were used for user's first and last name, emails, and the number of tickets being purchased 
as well the the number of tickets avilable to be purchased. In many instances, the program also use's the 
printf formatted output which allows the users to print formatted data via trailing arguments. The Booking 
App also features data types which are used in the cases of arrays, string types, numeric types, as well as 
booleans. The program also uses Golang's user input to obtain user's name, email, and number of tickets. 
The program also contains many logics and arguments to create a user experience when booking tickets, 
validing user inputs, and encapsulated logics within a function. Finally, the program use's one of the great 
features of Golang which is the implementaion of concurrency. In this case, we use the sleep function to stop 
or block any current thread for a certain time period. This allows to have a better flowing of the project when 
executing "sendTickets" in a separate thread. In the example of the Booking App, the program is able to execute 
the function of sendTickets to mulitple user's with different sets of data.

How to run the project? 
    go run .

How to exit the program?
    CTRL + C
