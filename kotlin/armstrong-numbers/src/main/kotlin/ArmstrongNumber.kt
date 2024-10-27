object ArmstrongNumber {
    fun check(input: Int): Boolean {
        var current = input
        val digits: MutableList<Int> = mutableListOf()
        while (current > 0) {
            digits.add(current % 10)
            current /= 10
        }
        val sum = digits.fold(0) { acc, digit -> acc + pow(digit, digits.size) }
        return input == sum
    }
    fun pow(a: Int, b: Int): Int {
        var result = 1
        for (i in 1..b) {
            result *= a
        }
        return result
    }
}
