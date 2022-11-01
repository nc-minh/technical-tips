const a = 1
const b = 2

let i = 0
while (i < 1000000000) {
    i = i + 1
}

while (i < n) {
    i = i + 1
}

//-------------------------------------------

for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
        /// do something
    }
}
// f(n) = n*n = n^2, O(f(n)) = O(n^2)

//-------------------------------------------
for (let i = 0; i < n; i++) {
    for (let j = i; j < n; j++) {
        // do something
    }
}
/***
 * f(n) = n*(n + (n-1) + (n-2) + ... + 3 + 2 + 1)
 *      = n*(n + 1)/2
 *      = n^2/2 + n/2
 *      = O(n^2)
 * => O(f(n)) = O(n^2)
 */

//-------------------------------------------
const value = 1
const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
for (let i = 0; i < n; i++) {
    if (arr[i] === value) {
        return i
    }

    return -1
}

/***
 * f(n) = cn O(f(n)) = O(n)
 */

//-------------------------------------------
// arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
// value = 7
// n = 10
let low = 0, high = n - 1 // 0, 9

while (low < high) { // 0 < 9 | 6 < 9 | 6 < 7
    const mid = (high + low) / 2 // 4.5 => 5 | 9 + 6 / 2 => 7.5 => 8 | 7 + 6 / 2 => 6.5 => 7
    if (arr[mid] === value) { // 5 !== 7 | 8 !== 7 | 7 === 7
        return mid // =>>>>>>>>>>>>>>> 7
    } else if (arr[mid] < value) { // 5 < 7 | 9 < 7
        low = mid + 1 // 6
    } else {
        high = mid - 1 // 7
    }
}

return -1