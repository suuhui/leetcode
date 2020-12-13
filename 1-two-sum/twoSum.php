<?php
function twoSum(array $nums, int $target) {
    $numsMap = [];
    foreach ($nums as $idx => $num) {
        $numsMap[$num] = $idx;
        $diffrence = $target - $num;
        if (isset($numsMap[$diffrence])) {
            return [$numsMap[$diffrence], $idx];
        }
    }
    return [];
}

$nums = [11, 2, 6, 7, 9];
$target = 9;
var_dump(twoSum($nums, $target));