package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int);
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, exists := units[unit]; !exists {
		return false
	}
	bill[item] = bill[item] + units[unit];
	return true;
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, exists := bill[item]; !exists {
		return false;
	}
	if _, exists := units[unit]; !exists {
		return false;
	}
	newQuantity := bill[item] - units[unit];
	if newQuantity < 0 {
		return false;
	}
	if newQuantity == 0 {
		delete(bill, item);
		return true;
	}
	bill[item] = newQuantity;
	return true;
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item];
	return  value, exists;
}
