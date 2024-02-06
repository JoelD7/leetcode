# Problem
https://www.geeksforgeeks.org/find-maximum-average-marks-in-student-pairs/

Given an array A[] containing N pairs in the form of (X, Y). Where X and Y denote the student’s ID and marks respectively. Then your task is to output the ID of the student that has the highest average mark.

### Example 1
    Input: A[] = {{10, 78}, {12, 85}, {10, 90}, {12, 92}, {8, 88}})
    Output: 12
    Explanation: The average marks of each student are as follows:

        10: (78+90)/2 = 168/2 = 84
        12: (85+92)/2 = 177/2 = 88.5
        8: (88)/1 = 88
        Therefore, the student with the highest average marks is 12.

If several students have the same average marks, then the student with the lower ID should be returned.

### Example 2
    Input: A[] = {{10, 78}, {12, 85}, {10, 90}, {12, 92}, {8, 88}, {8, 90}})
    Output: 12
    Explanation: The average marks of each student are as follows:

        10: (88+90)/2 = 178/2 = 89
        12: (85+92)/2 = 177/2 = 88.5
        8: (88+90)/2 = 178/2 = 89
        Both 10 and 8 have the same average marks, but 8 has a lower ID than 10. Therefore, the student with the highest average marks is 12.
