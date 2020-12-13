import java.util.HashMap;
import java.util.Map;

class TwoSum
{
    public static void main(String[] args) {
        int[] nums = {1, 2, 7, 11, 34};
        int target = 9;
        int[] result = twoSum(nums, target);
        System.out.print(result);
    }

    private static int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            map.put(nums[i], i);
            if (map.containsKey(target - nums[i])) {
                return new int[]{map.get(target - nums[i]), i};
            }
        }
        return new int[]{};
    }
}