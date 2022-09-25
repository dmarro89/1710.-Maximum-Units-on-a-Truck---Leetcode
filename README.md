# 1710. Maximum Units on a Truck
## This is a go solution for the problem - Check the main.go file

## **Problem Description**
You are assigned to put some amount of boxes onto one truck. You are given a 2D array boxTypes, where boxTypes[i] = [numberOfBoxesi, numberOfUnitsPerBoxi]:

1. numberOfBoxesi is the number of boxes of type i.
2. numberOfUnitsPerBoxi is the number of units in each box of the type i.

You are also given an integer truckSize, which is the maximum number of boxes that can be put on the truck. You can choose any boxes to put on the truck as long as the number of boxes does not exceed truckSize.

Return the maximum total number of units that can be put on the truck.

Example 1:

Input: boxTypes = [[1,3],[2,2],[3,1]], truckSize = 4
Output: 8

Explanation: 

There are:
- 1 box of the first type that contains 3 units.
- 2 boxes of the second type that contain 2 units each.
- 3 boxes of the third type that contain 1 unit each.
You can take all the boxes of the first and second types, and one box of the third type.
The total number of units will be = (1 * 3) + (2 * 2) + (1 * 1) = 8.
Example 2:

Input: boxTypes = [[5,10],[2,5],[4,7],[3,9]], truckSize = 10
Output: 91
 

Constraints:

1 <= boxTypes.length <= 1000

1 <= numberOfBoxesi, numberOfUnitsPerBoxi <= 1000

1 <= truckSize <= 106

## **Solution**
To maximize the number of units to put on the track, the boxes should be considered in order of descending units.

For example, if this is the boxTypes matrix :
> [[5,10],[2,5],[4,7],[3,9]]

It should be ordered by descending units.
>[[5,**10**],[3,**9**],[4,**7**],[2,**5**]]

Then, the first **truckSize** boxes should be picked up.
So, if the truckSize is equal to ten, then:
1. The first 5 boxes of the first array - truckSize = 10- **5** = 5
2. The first 3 boxes of the second array - truckSize = 5 - **3** = 2
3. The first 2 boxes of the third array - truckSize = 2 - **2** = 0

At this point, the total output should be the sum of the first **truckSize** multiplied for their units:

* (5 * 10) + (3 * 9) + (2 * 7)
 = 50 + 27 + 14 = **91**

 ## Go Implementation
For simplicity, in this case it will be used a native golang sorting function (**sort.SliceStable**).
The performance of the solution can be improved using a more efficient sorting alghoritm.
After the slice has been sorted, a loop starts on the **typeBoxes** matrix. Until the **truckSize** is not equal to 0, the total units of the current box is calculated (as the multiplication between the first element of the array and the second element of the array) and added to the previous values calculated. Of course, at each iteration the **truckSize** considered is decremented by the number of the boxes picked up.
The value returned is the sum of the multiplication done for each iteration.

#
# Any suggestion to improve the code and/or the performance?