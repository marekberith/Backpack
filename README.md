# Backpack
Backpack Calculator in Golang + QT

Backpack is a GoLang app used to calculate the maximum value of a backpack.

Example of an input:

(1, 1); (2, 2); (4, 3); (4, 4); (5, 5) //items added into the backpack(price, weight)

10 //maximum weight of the backpack

5 //number of items

Output:

((4, 10), (3, 6), (2, 3), (1, 1))


## How to use?

You can either build the app or use the Windows binary(will be added in few days).

If you want to build the app, you'll need to install GO and the qt binding into your GO root. Please, use the following package: https://github.com/therecipe/qt.
Then clone the repository, open the src folder and build the app using the following command:

`go build`
